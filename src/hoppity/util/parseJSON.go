package util

import (
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"
)

func ParseBodyJsonObj(reader *io.ReadCloser, obj interface{}) (error) {
    defer (*reader).Close()
    jsonBytes, err := ioutil.ReadAll(*reader)
    if err == nil {
	    err = json.Unmarshal(jsonBytes, obj)
    }
    return err
}

func ParseJsonObj(req *http.Request, obj interface{}) (error) {
    return ParseBodyJsonObj(&req.Body, obj)
}
