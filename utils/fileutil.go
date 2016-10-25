package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

func DownloadFile(basedir string, urlstr string, reload bool) (filepath string, err error) {

	filename, err := getFilename(urlstr)
	if err != nil {
		return "", err
	}
	filepath = fmt.Sprintf("%s/%s", basedir, filename)

	// check if non zero size file exists, download if not exist or reload true
	finfo, err := os.Stat(filepath)
	if err == nil && finfo != nil && finfo.Size() > 0 && reload == false {
		return filepath, nil
	}
	// Create the file
	//fmt.Printf("filepath %s", filepath)
	out, err := os.Create(filepath)
	if err != nil {
		return filepath, err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(urlstr)
	if err != nil {
		return filepath, err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return filepath, err
	}

	return filepath, nil
}

func WriteFile(basedir string, filename string, data []byte, rewrite bool) (filepath string, err error) {
	filepath = fmt.Sprintf("%s/%s", basedir, filename)
	// Create the file
	fp, err := os.Create(filepath)
	if err != nil {
		return filepath, err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		return filepath, err
	}
	return filepath, nil

}

func getFilename(urlstring string) (string, error) {
	u, err := url.Parse(urlstring)
	if err != nil {
		return "", err
	}
	return path.Base(u.Path), nil
}

func EncodeSlice(inmap []string) string {
	mapbyte, err := json.Marshal(inmap)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	return string(mapbyte)
}
func DecodeToMap(mapstr string) []string {
	var outmap []string
	if mapstr == "" {
		return outmap
	}
	err := json.Unmarshal([]byte(mapstr), &outmap)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	return outmap
}
