package main

import (
	"encoding/json"
)

type Notify struct {
	Repository string
	Branch     string
}

func ParseBitBucket(body []byte) (Notify, error) {

	notify := Notify{}
	result := make(map[string]interface{})

	err := json.Unmarshal([]byte(body), &result)
	if err != nil {
		return notify, err
	}

	notify.Repository = result["repository"].(map[string]interface{})["full_name"].(string)
	notify.Branch = result["push"].(map[string]interface{})["changes"].([]interface{})[0].(map[string]interface{})["new"].(map[string]interface{})["name"].(string)

	return notify, nil
}
