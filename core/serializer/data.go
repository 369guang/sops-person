package serializer

import (
	"reflect"
	"strings"
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
		source := modelT.Elem().Field(i).Tag.Get("serializers")
		index := modelT.Elem().Field(i).Tag.Get("json")
		if source != "" && strings.ContainsAny(source, "write") {
			continue
		}
		data[index] = modelV.Elem().Field(i).Interface()
	}
	return data
}
