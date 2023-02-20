package controllers

import (
	"MobileIDAPI/models"
	"MobileIDAPI/services"
	"MobileIDAPI/utils/validate"
	"encoding/json"
	"github.com/golang/glog"
	"net/http"
)

func IdentityByISDNAndIPDestination(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	requestBody := new(models.IdentityRequest)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&requestBody)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		switch err.(type) {
		case *json.SyntaxError:
			glog.Error("Json syntax error:" + err.Error())
			w.Write([]byte(""))
		case *json.UnmarshalTypeError:
			glog.Error("Json Unmarshal Type Error:" + err.Error())
			w.Write([]byte(""))
		default:
			glog.Error("Unknown Error:" + err.Error())
			w.Write([]byte(""))
		}
	} else {
		isValidPhone := validate.ValidatePhoneNumber(requestBody.Msisdn)
		if requestBody.Msisdn == "" || requestBody.Address == "" || !isValidPhone {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(""))
		} else {
			status, res := services.IdentifyPhoneNumber(requestBody)
			w.WriteHeader(status)
			w.Write(res)
		}

	}
}
