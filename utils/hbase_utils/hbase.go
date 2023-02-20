package hbase_utils

import (
	"MobileIDAPI/models"
	"MobileIDAPI/utils/settings"
	"context"
	"github.com/golang/glog"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"strconv"
)

var client gohbase.Client

func GetHBaseClient() (gohbase.Client, error) {
	if client == nil {
		// create a new client if one doesn't exist
		address := settings.GetHbaseConfig().Host + ":" + strconv.Itoa(settings.GetHbaseConfig().Port)
		client = gohbase.NewClient(address)
	}
	// return the singleton client
	return client, nil
}
func GetInfoIndentityByRowKey(clientHbase gohbase.Client, identitySchema settings.Identity, rowKey string) (error, models.Identity) {
	identity := models.Identity{
		Timestamp: "",
		Phone:     "",
	}
	tableName := identitySchema.NameTable
	columnFamilyGet := identitySchema.ColumFamily1Identity.Name
	getReq, err := hrpc.NewGetStr(context.Background(), tableName, rowKey, hrpc.Families(map[string][]string{columnFamilyGet: nil}))
	if err != nil {
		glog.Error("Error get record identity by rowkey: ", rowKey, err)
		return err, identity
	}
	// send the Get request to HBase
	getRsp, err := clientHbase.Get(getReq)
	if err != nil {
		glog.Error("Error get record radius by rowkey: ", rowKey, err)
		return err, identity
	}
	cells := getRsp.Cells
	if len(cells) < 1 {
		return nil, identity
	}
	identity.Phone = string(cells[0].Value)
	identity.Timestamp = string(cells[1].Value)
	//for _, cell := range getRsp.Cells {
	//	glog.Info("column family: %s, column: %s, value: %s\n", string(cell.Family), string(cell.Qualifier), string(cell.Value))
	//}
	return nil, identity
}
