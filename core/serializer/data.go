package serializer

import (
	"reflect"
)

func ModelToSerializers(model interface{}) map[string]interface{} {
	modelT := reflect.TypeOf(model)
	modelV := reflect.ValueOf(model)

	data := make(map[string]interface{})

	for i := 0; i < modelT.Elem().NumField(); i++ {
		if modelT.Elem().Field(i).Name == "BaseModel" {
			baseModel := reflect.ValueOf(modelV.Elem().Field(i).Interface())
			data["id"] = baseModel.FieldByName("ID").Interface()
			data["created"] = baseModel.FieldByName("CreatedAt").Interface()
			data["updated"] = baseModel.FieldByName("UpdatedAt").Interface()
			continue
		}
		index := modelT.Elem().Field(i).Tag.Get("json")
		data[index] = modelV.Elem().Field(i).Interface()
	}
	return data
}
