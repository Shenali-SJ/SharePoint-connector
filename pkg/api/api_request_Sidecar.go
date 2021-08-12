package api

import (
	"context"
	"reflect"
	arfsidecar "sharepoint-connector/pkg/rule/ARF_Sidecar"
	util "sharepoint-connector/pkg/util"
)

import (
	"github.com/jinzhu/copier"
	"sharepoint-connector/pkg/dto"
	"sharepoint-connector/sidecar"
)

func Sidecar(request *sidecar.RequestDto, ctx context.Context) (*sidecar.ResponseDto, error) {

	mRequest := dto.Request{}
	copier.Copy(&mRequest, &request)
	cntxt := util.CustomContext{}
	cntxt.GoContext = ctx
	cntxt.AppGoContext = context.Background()
	config := make(map[string]interface{})
	result := arfsidecar.ARF_Sidecar(&mRequest, &cntxt, config)
	if reflect.TypeOf(result) == reflect.TypeOf(dto.Response{}) || reflect.TypeOf(result) == reflect.TypeOf(&dto.Response{}) {
		var data *sidecar.ResponseDto
		copier.Copy(&data, result.(*dto.Response))
		return data, nil
	} else {
		var data sidecar.ResponseDto
		return &data, nil
	}
}
