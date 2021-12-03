package coreplugins

import (
	"encoding/json"
	"io/ioutil"

    "lab2/src/interfaces"
)

func Env()interfaces.Env{
    fileenv := "./src/config/dev.json"
    var data []byte
    data, _ = ioutil.ReadFile(fileenv)

    // dataJson
    var dataJson interfaces.Env
    _ = json.Unmarshal(data, &dataJson)

    return dataJson
}