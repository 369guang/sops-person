package serializers

import (
	"errors"
	"gorm.io/gorm"
	"person/apps/user/models"
	"person/core"
	"person/core/database"
	"person/core/errno"
	"person/core/serializer"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

func List(u *User) (interface{}, int64, error) { //查询

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
		data := serializer.ModelToSerializers(user)
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

	data := serializer.ModelToSerializers(user)

	return data, nil
}

func Create(u *models.User) error { // 创建
	if !errors.Is(core.DATABASE.Where("username = ?", u).First(&models.User{}).Error, gorm.ErrRecordNotFound) {
		return errno.ErrUserDuplicate
	}
	if err := core.DATABASE.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func Update(id int, u interface{}, method string) (interface{}, error) { //更新

	return nil, nil
}

func Destroy(id int) error { // 删除

	return nil
}
