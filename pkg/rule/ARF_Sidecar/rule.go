package ARF_Sidecar

import util "sharepoint-connector/pkg/util"

import (
	"sharepoint-connector/pkg/dto"
	"sharepoint-connector/pkg/functions"
	"sharepoint-connector/pkg/model"
	"sharepoint-connector/pkg/rule/ARF_GetListItemsSubrule"
	"sharepoint-connector/pkg/rule/ARF_InitSubrule"
	"sharepoint-connector/pkg/xiLogger"
)

type CFGContext struct {
	Request          *dto.Request
	operation        string
	requestString    string
	requestJsonModel *model.Requestdata
	listItems        []*model.Listitemdata
	_context         *util.CustomContext
	_ruleConfig      map[string]interface{}
	_returnValue     interface{}
	_errorValue      interface{}
}

func NewCFGContext(pRequest *dto.Request, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Request:     pRequest,
		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_Sidecar(pRequest *dto.Request, pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pRequest, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	if err := cfg.xiHybridFunctionNodeM01(); err != nil {
		xiLogger.Log.Info("Could not convert from string to JSON")
	}

	cfg.xiModelPropertyNodeM12()

	cfg.xiConstantNodeM23()

	cfg.xiSwitchNodeM34()
	return nil
}

func (cfg *CFGContext) xiSwitchNodeM34() error {

	switch _input := cfg.operation; {

	case _input == initOperation:
		cfg.branchM41()
	case _input == getListItemsOperation:
		cfg.branchM42()

	}
	return nil
}

const initOperation = "init"
const getListItemsOperation = "getListItems"

func (cfg *CFGContext) xiConstantNodeM23() error {
	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM12() error {
	cfg.operation = cfg.requestJsonModel.Operation

	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM01() error {
	var err error
	cfg.requestJsonModel, err = functions.ConvertStringToJson(cfg.requestString)
	cfg._returnValue = cfg.requestJsonModel
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.requestString = cfg.Request.Data
	cfg.requestJsonModel = &model.Requestdata{}

	return nil
}
func (cfg *CFGContext) branchM41() error {

	cfg.xiSubRuleNodeM5()
	return nil
}

func (cfg *CFGContext) xiSubRuleNodeM5() error {
	ARF_InitSubrule.ARF_InitSubrule(cfg.requestJsonModel, cfg._context, cfg._ruleConfig)
	return nil
}
func (cfg *CFGContext) branchM42() error {

	cfg.xiModelPropertyNodeM7()

	cfg.xiSubRuleNodeM76()
	return nil
}

func (cfg *CFGContext) xiSubRuleNodeM76() error {
	cfg.listItems = ARF_GetListItemsSubrule.ARF_GetListItemsSubrule(cfg.requestJsonModel, cfg._context, cfg._ruleConfig).([]*model.Listitemdata)
	cfg._returnValue = cfg.listItems
	return nil
}

func (cfg *CFGContext) xiModelPropertyNodeM7() error {

	return nil
}
