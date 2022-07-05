package database

import (
	"encoding/json"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type BaseModel struct {
	ID        int            `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

//CombinationConditions 组合条件
func CombinationConditions(whereStruct interface{}) map[string]interface{} {
	whereT := reflect.TypeOf(whereStruct)
	whereV := reflect.ValueOf(whereStruct)
	data := make(map[string]interface{})
	for i := 0; i < whereT.NumField(); i++ {
		if whereT.Field(i).Tag.Get("json") == "page" || whereT.Field(i).Tag.Get("json") == "page_size" {
			continue
		}
		params := whereV.Field(i).Interface()
		if whereV.Field(i).IsZero() {
			continue
		}
		data[whereT.Field(i).Tag.Get("json")] = params
	}

	return data
}

// ExcludeStructToMap 排除更新字段
func ExcludeStructToMap(source interface{}, exclude ...string) map[string]interface{} {
	// 排除base model
	defaultExclude := []string{"id", "created_at", "updated_at", "deleted_at"}
	m := make(map[string]interface{})
	data, _ := json.Marshal(source)
	json.Unmarshal(data, &m)

	for _, de := range defaultExclude {
		if _, ok := m[de]; ok {
			delete(m, de)
		}
	}

	if len(exclude) > 0 {
		for _, de := range exclude {
			if _, ok := m[de]; ok {
				delete(m, de)
			}
		}
	}

	return m
}

func ExcludeMap(data map[string]interface{}, exclude ...string) {
	defaultExclude := []string{"id", "created_at", "updated_at", "deleted_at"}
	for _, de := range defaultExclude {
		if _, ok := data[de]; ok {
			delete(data, de)
		}
	}

	if len(exclude) > 0 {
		for _, de := range exclude {
			if _, ok := data[de]; ok {
				delete(data, de)
			}
		}
	}
}
