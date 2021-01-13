package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Jeffail/gabs/v2"
)

type LockedUser struct {
	ID      string
	Channel string
}

type TempLockedUsers struct {
	Users []LockedUser `json:"Users"`
}

func (users *TempLockedUsers) Write() error {
	jsonObj := gabs.New()
	jsonObj.Array("Users")
	jsonObj.ArrayConcat(users, "Users")
	_ = ioutil.WriteFile("tmp/temp_locked.json", jsonObj.Bytes(), 0644)
	return nil
}

func (users *TempLockedUsers) Read() (*gabs.Container, error) {
	jsonFile, err := os.Open("tmp/temp_locked.json")
	if err != nil {
		return nil, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonParsed, err := gabs.ParseJSON(byteValue)
	if err != nil {
		return nil, err
	}

	result := jsonParsed.S("Users")

	defer jsonFile.Close()
	return result, nil
}

func (users *TempLockedUsers) Update(user *gabs.Container) error {
	existingData, err := users.Read()
	if err != nil {
		return err
	}
	user.Merge(existingData)
	_ = ioutil.WriteFile("tmp/temp_locked.json", user.Bytes(), 0644)
	return nil
}

func (users *TempLockedUsers) Remove(ID string) error {
	newData, err := users.Read()
	if err != nil {
		return err
	}
	for i, user := range newData.S("Users").Children() {
		if user.S("ID").Data().(string) == ID {
			newData.ArrayRemove(i, "Users")
		}
	}
	_ = ioutil.WriteFile("tmp/temp_locked.json", newData.Bytes(), 0644)
	return nil
}

type PermLockedUsers struct {
	Users []string
}

func (users *PermLockedUsers) Write() error {
	file, _ := json.MarshalIndent(users, "", " ")
	_ = ioutil.WriteFile("tmp/perm_locked.json", file, 0644)
	return nil
}

func UserIsLocked(a string, list *gabs.Container) (bool, LockedUser) {
	for _, b := range list.Children() {
		if b.S("ID").Data().(string) == a {
			return true, LockedUser{
				ID:      b.S("ID").Data().(string),
				Channel: b.S("Channel").Data().(string),
			}
		}
	}
	return false, LockedUser{}
}
