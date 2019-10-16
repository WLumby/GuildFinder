package main

import (
	"encoding/json"
)

func unmarshalRealm(jsonFile string) (realmList []Guild, err error) {
	err = json.Unmarshal([]byte(jsonFile), &realmList)
	if err != nil {
		return []Guild{}, nil
	}

	return
}
