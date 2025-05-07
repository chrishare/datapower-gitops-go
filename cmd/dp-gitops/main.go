package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/chrishare/datapower-gitops-go/internal/dpobjects"
)

func main() {
	log.Println("Starting main")

	// TODO: Do TLS verification. Disabling SSL just for development purposes.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	server := dpobjects.Server{
		Hostname: "localhost:5554",
		Username: "admin",
		Password: "password",
	}
	serverStatus := dpobjects.TestConnection(server)
	if serverStatus != dpobjects.SUCCESS {
		log.Println("Got an error:", serverStatus)
		os.Exit(1)
	}
	//var baseObject dpobjects.BaseObject
	//_ = baseObject

	//var logger dpobjects.LogTarget
	//_ = logger
	loggerData := map[string]interface{}{
		"name":        "test_policy",
		"mAdminState": "enabled",
		"Type":        "file",
		"Size":        500,
		"LocalFile":   "logtemp:///temp.log",
		"LogEvents":   "blah",
	}

	//var logTarget dpobjects.LogTarget = dpobjects.NewLogTarget(loggerData)
	//baseObject = logTarget
	//log.Println(logTarget.ToPrettyJSON())
	//log.Println(baseObject.ToPrettyJSON())
	//log.Println(baseObject.ToPrettyYAML())

	logCat := dpobjects.LogTarget{DPObjectCommon: dpobjects.DPObjectCommon{InternalState: loggerData}}
	var dpBase dpobjects.DPObject = logCat
	log.Println(dpBase.ToPrettyJSON())

}
