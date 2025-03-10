package main

import (
	"crypto/tls"
	"log"
	"net/http"

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
	dpobjects.TestConnection(server)

	var baseObject dpobjects.BaseObject
	_ = baseObject

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
	logger := dpobjects.NewLogTarget(loggerData)
	_ = logger

	log.Println(dpobjects.ToPrettyJSON(logger))
	log.Println(dpobjects.ToPrettyYAML(logger))
}
