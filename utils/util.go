package utils

import (
	"connector/conf"
	"encoding/json"
	"io/ioutil"
	"path"
	"reflect"
)

func Map2String(data interface{}) (string, error) {
	str, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(str), err
}

//结构体转为map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func String2Map(data string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func WriteFile(companyId string, data []byte) error {
	file := path.Join(conf.GatewayCompanyConfigDir, companyId+`_gateway.json`)
	err := ioutil.WriteFile(file, data, 0777)
	if err != nil {
		return err
	}
	return nil
}
