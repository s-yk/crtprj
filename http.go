package main

import (
	"io/ioutil"
	"net/http"
)

// getProjectTemplate Get zip from github
func getProjectTemplate(fileName string) error {

	url := "https://github.com/s-yk/crtprj/archive/master.zip"

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, byteArray, 0777)
	if err != nil {
		return err
	}

	return nil
}
