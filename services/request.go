package services

import (
	"connector/conf"
	"connector/lib"
	"connector/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

var infoLog lib.InfoLogger
var debugLog lib.DebugLogger
var errorLog lib.ErrorLogger

type GatewayData struct {
	Data      map[string]interface{}   `json:"data"`
	DataList  []map[string]interface{} `json:"data"`
	ErrCode   string                   `json:"errCode"`
	CompanyId string                   `json:"companyId"`
}

//InitGateway InitGateway
func InitGateway() {
	managerAddr := conf.ManagerAddr
	url := managerAddr + `/manager/api/gateway?companyId=` + viper.GetString("server.company_id")
	debugLog.Println("GET manager url is:", url)

	response, err := http.Get(url)
	debugLog.Println("/manager/api/gateway response", response)
	if err != nil || response.StatusCode != http.StatusOK {
		errorLog.Println("get manager err !", url, err)
		return
	}

	//change body type to []byte
	bytebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("change response.Body type to []byte err, %v\n", err)
		return
	}

	//get json body into data
	var data GatewayData
	if err = json.Unmarshal(bytebody, &data); err != nil {
		errorLog.Println("Unmarshal err:", err)
		return
	}

	cID := data.DataList[0]["companyId"]
	debugLog.Println("data.Data is:", data.DataList[0])
	debugLog.Println("data.companyId is:", cID)
	debugLog.Println("data.errCode is:", data.ErrCode)

	dataString, err := utils.Map2String(data.DataList[0])
	if err != nil {
		errorLog.Println(err)
		return
	}

	cIDString, err := utils.Map2String(cID)
	if err != nil {
		errorLog.Println(err)
		return
	}

	err = utils.WriteFile(string(cIDString), []byte(dataString))
	if err != nil {
		errorLog.Println(err)
		return
	}

}
