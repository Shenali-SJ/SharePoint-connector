package ARF_GetListItemsSubrule

import util "sharepoint-connector/pkg/util"

import (
	"sharepoint-connector/pkg/functions"
	"sharepoint-connector/pkg/model"
)

type CFGContext struct {
	Requestdata  *model.Requestdata
	listName     string
	items        []*model.Listitemdata
	_context     *util.CustomContext
	_ruleConfig  map[string]interface{}
	_returnValue interface{}
	_errorValue  interface{}
}

func NewCFGContext(pRequestdata *model.Requestdata, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Requestdata: pRequestdata,
		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_GetListItemsSubrule(pRequestdata *model.Requestdata, pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pRequestdata, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiHybridFunctionNodeM01()
	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM01() error {
	var err error
	cfg.items, err = functions.GetItems(cfg.listName)
	cfg._returnValue = cfg.items
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.listName = cfg.Requestdata.Listname

	return nil
}
