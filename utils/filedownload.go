package utils

import (
    "os"
    "io"
    "path"
    "fmt"
    "net/http"
    "net/url"
)

func DownloadFile(basedir string, urlstr string) (filename string, err error) {

  filename, err = getFilename(urlstr)
  if err != nil {
    return "", err
  }
  filepath := fmt.Sprintf("%s/%s",basedir,filename)
 
  // Create the file
  out, err := os.Create(filepath)
  if err != nil  {
    return filename, err
  }
  defer out.Close()

  // Get the data
  resp, err := http.Get(urlstr)
  if err != nil {
    return filename, err
  }
  defer resp.Body.Close()

  // Writer the body to file
  _, err = io.Copy(out, resp.Body)
  if err != nil  {
    return filename, err
  }

  return filename, nil
}

func getFilename(urlstring string) (string, error){
	u, err := url.Parse(urlstring)
	if err != nil {
		return "", err
	}
  return path.Base(u.Path), nil
}
