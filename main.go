package main

import (
	"context"
	trfmain "sharepoint-connector/pkg/rule/TRF_Main"
	util "sharepoint-connector/pkg/util"
)

import (
	_ "sharepoint-connector/docs"
)
import (
	"sharepoint-connector/pkg/middleware"
)
import (
	"log"
	"net"
	"sharepoint-connector/pkg/controller"
	"sharepoint-connector/pkg/xiLogger"
	"sharepoint-connector/sidecar"
)

func main() {

	cntxt := util.CustomContext{}
	cntxt.AppGoContext = context.Background()
	config := make(map[string]interface{})
	trfmain.TRF_Main(&cntxt, config)

	RunServer()
}

func RunServer() {
	port := middleware.GetGrpcPort()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := middleware.GetGrpcServerConfigs()

	xiLogger.Log.Info("\n Server listening on port %v \n", ":"+port)
	sidecar.RegisterServiceServer(s, &controller.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
