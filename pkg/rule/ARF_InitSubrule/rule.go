package ARF_InitSubrule

import util "sharepoint-connector/pkg/util"

import (
	"sharepoint-connector/pkg/env"
	"sharepoint-connector/pkg/fn"
	"sharepoint-connector/pkg/model"
	"sharepoint-connector/pkg/xiLogger"
)

type CFGContext struct {
	Requestdata  *model.Requestdata
	siteUrl      string
	clientId     string
	clientSecret string
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
func ARF_InitSubrule(pRequestdata *model.Requestdata, pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pRequestdata, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	if err := cfg.xiLambdaFunctionNodeM01(); err != nil {
		xiLogger.Log.Info("Could not authenticate")
	}
	return nil
}

func (cfg *CFGContext) xiLambdaFunctionNodeM01() error {
	var err error
	env.Initreturnvalue, err = fn.Init(cfg.siteUrl, cfg.clientId, cfg.clientSecret)
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.siteUrl = cfg.Requestdata.Siteurl
	cfg.clientId = cfg.Requestdata.Clientid
	cfg.clientSecret = cfg.Requestdata.Clientsecret

	return nil
}
