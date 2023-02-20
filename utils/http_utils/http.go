package http_utils

import (
	"MobileIDAPI/models"
	"encoding/json"
	"github.com/golang/glog"
	"net/http"
)

func ResponseError(msg string, err error, code int) (int, []byte) {
	res := models.ResponseHTTP{
		Status: msg,
		Data:   &models.MessageData{Message: err.Error()},
	}
	resJson, _ := json.Marshal(res)
	return code, resJson
}
func ResponseOK(identity models.Identity, codeRes int, status string) (int, []byte) {
	res := models.ResponseHTTP{
		Status: status,
		Time:   identity.Timestamp,
		Data:   nil,
	}
	resJson, err := json.Marshal(res)
	if err != nil {
		glog.Error(err)
		code, resJson := ResponseError("ERROR", err, http.StatusInternalServerError)
		return code, resJson
	}
	return codeRes, resJson
}
