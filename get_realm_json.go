package main

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"
)

func getRealmData(realm string) (string, error)  {
	fileUrl := "https://www.wowprogress.com/export/ranks/"+realm+"_tier25.json.gz"
	return downloadFile(fileUrl)
}

func downloadFile(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return unzip(*resp)
}

func unzip(resp http.Response) (string, error) {
	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return "", nil
	}
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", nil
	}

	return string(body), nil
}
