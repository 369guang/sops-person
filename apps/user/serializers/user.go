package serializers

import (
	"errors"
	"gorm.io/gorm"
	"person/apps/user/models"
	"person/core"
	"person/core/auth"
	"person/core/database"
	"person/core/errno"
	"person/core/serializer"
)

type LoginFields struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

func List(u User) (interface{}, int64, error) { //查询

	users := make([]models.User, 0)
	db := core.DATABASE
	where := database.CombinationConditions(u)
	var total int64
	db.Model(&models.User{}).Where(where).Count(&total)

	sPage, pageSize, err := serializer.Partition(total, u.Page, u.PageSize)
	if err != nil {
		return nil, 0, err
	}

	if err := db.Where(where).Offset(sPage).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	dataset := make([]interface{}, 0)
	for _, user := range users {
		user.Password = ""
		data := serializer.ModelToSerializers(&user)
		dataset = append(dataset, data)
	}

	return dataset, total, nil
}

func Retrieve(id int) (interface{}, error) { // 详情

	user := new(models.User)
	err := core.DATABASE.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	user.Password = ""
	data := serializer.ModelToSerializers(user)

	return data, nil
}

func Create(u *models.User) error { // 创建
	if !errors.Is(core.DATABASE.Where("username = ?", u.Username).First(&models.User{}).Error, gorm.ErrRecordNotFound) {
		return errno.ErrUserDuplicate
	}
	if err := core.DATABASE.Create(&u).Error; err != nil {
		core.LOGGER.Error("create user error:  " + err.Error())
		return err
	}

	// 创建用户角色
	key, url := auth.CreateKey(u.Username)
	authTotp := models.AuthToTp{
		UserID: u.ID,
		Keys:   key,
		Urls:   url,
	}
	core.DATABASE.Create(&authTotp)
	return nil
}

func Update(id int, u *models.User, method string) error { //更新
	if errors.Is(core.DATABASE.Where("id = ?", id).First(&models.User{}).Error, gorm.ErrRecordNotFound) {
		return errno.ErrRecordNotExist
	}

	updateData := database.ExcludeStructToMap(&u, "username", "last_login", "login_ip")
	if _, ok := updateData["password"]; ok {
		if updateData["password"] == "" {
			delete(updateData, "password")
		} else {
			updateData["password"] = u.Password
		}
	}

	if err := core.DATABASE.Model(&models.User{}).Where("id = ?", id).Updates(updateData).Error; err != nil {
		core.LOGGER.Error("update user error:  " + err.Error())
		return err
	}

	return nil
}

func Destroy(id int) error { // 删除
	if errors.Is(core.DATABASE.Where("id = ?", id).First(&models.User{}).Error, gorm.ErrRecordNotFound) {
		return errno.ErrRecordNotExist
	}

	if err := core.DATABASE.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		core.LOGGER.Error("delete user error:  " + err.Error())
		return err
	}
	return nil
}

func Login(u *LoginFields) (string, error) { // 登录
	user := new(models.User)
	fUser := core.DATABASE.Where("username = ?", u.Username).First(&user)
	if errors.Is(fUser.Error, gorm.ErrRecordNotFound) {
		return "", errno.ErrUserNotFound
	}

	if user.Status == 0 || user.Status == 2 {
		return "", errno.ErrUserDisabled
	}

	token, err := auth.Sign(auth.Context{
		ID:       user.ID,
		Username: user.Username,
	}, "")
	if err != nil {
		return "", err
	}

	return token, nil
}

func Info(c *auth.Context) (interface{}, error) { // 获取用户信息

	user := new(models.User)
	fUser := core.DATABASE.Where("id = ?", c.ID).First(&user)
	if errors.Is(fUser.Error, gorm.ErrRecordNotFound) {
		return "", errno.ErrUserNotFound
	}
	user.Password = ""

	return user, nil
}

func UpdateSettings(c *auth.Context, u *models.User) error { // 更新用户设置
	updateData := database.ExcludeStructToMap(&u, "username", "last_login", "login_ip")
	if _, ok := updateData["password"]; ok {
		if updateData["password"] == "" {
			delete(updateData, "password")
		} else {
			updateData["password"] = u.Password
		}
	}

	if err := core.DATABASE.Model(&models.User{}).Where("id = ?", c.ID).Updates(updateData).Error; err != nil {
		core.LOGGER.Error("update settings error:  " + err.Error())
		return err
	}
	return nil
}
