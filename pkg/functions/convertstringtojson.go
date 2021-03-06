package functions

import (
	"encoding/json"
	"sharepoint-connector/pkg/model"
)

func ConvertStringToJson(requeststring string) (*model.Requestdata, error) {

	var jsonModel *model.Requestdata
	err := json.Unmarshal([]byte(requeststring), jsonModel)

	return jsonModel, err
}
