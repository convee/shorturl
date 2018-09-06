package util

import (
	"encoding/json"
	"net/http"
)

type Json struct {
	Error int         `json:"error"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func JsonReturn(w http.ResponseWriter, raw interface{}) {
	jsonStr, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}
