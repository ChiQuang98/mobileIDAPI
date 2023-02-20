package services

import (
	"MobileIDAPI/models"
	"MobileIDAPI/utils/hbase_utils"
	"MobileIDAPI/utils/http_utils"
	"MobileIDAPI/utils/settings"
	"github.com/golang/glog"
	"net/http"
)

func IdentifyPhoneNumber(identityRequest *models.IdentityRequest) (int, []byte) {
	clientHbase, err := hbase_utils.GetHBaseClient()
	if err != nil {
		glog.Error("Error get client Hbase ", err)
		code, resJson := http_utils.ResponseError("ERROR", err, http.StatusInternalServerError)
		return code, resJson
	}
	shemaIdentity := settings.GetSchemaHbase().Identity
	rowKey := identityRequest.Address + "|" + identityRequest.Msisdn
	err, identity := hbase_utils.GetInfoIndentityByRowKey(clientHbase, shemaIdentity, rowKey)
	//Neu tim thay, co data
	if err == nil && identity.Phone != "" {
		code, resJson := http_utils.ResponseOK(identity, http.StatusOK, "OK")
		return code, resJson
	} else if err == nil {
		code, resJson := http_utils.ResponseOK(identity, http.StatusOK, "NOT_MATCHED")
		return code, resJson
	}
	glog.Error(err)
	code, resJson := http_utils.ResponseError("ERROR", err, http.StatusInternalServerError)
	return code, resJson
}
