package dpobjects

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Hostname string
	Username string
	Password string
}

type ServerStatus int

const (
	SUCCESS = iota
	REQUEST_CREATE_FAILED
	BAD_TCP_CONNECT
	RESPONSE_READ
	BAD_AUTH
	OBJECT_NOT_FOUND
	BAD_REQUEST
)

func (s ServerStatus) String() string {
	switch s {
	case SUCCESS:
		return "SUCCESS"
	case REQUEST_CREATE_FAILED:
		return "REQUEST_CREATE_FAILED"
	case BAD_TCP_CONNECT:
		return "BAD_TCP_CONNECT"
	case RESPONSE_READ:
		return "RESPONSE_READ"
	case OBJECT_NOT_FOUND:
		return "OBJECT_NOT_FOUND"
	case BAD_REQUEST:
		return "BAD_REQUEST"
	default:
		return "Unknown"
	}
}

func TestConnection(server Server) ServerStatus {

	requestUrl := fmt.Sprintf("https://%s", server.Hostname)
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Println("Could not create server request: ", err)
		return REQUEST_CREATE_FAILED
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Bad TCP connect: ", err)
		return BAD_TCP_CONNECT
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Connected but could not read the response: ", err)
		return RESPONSE_READ
	}
	log.Printf("client: response body: %s\n", resBody)
	return SUCCESS
}

type DPObject interface {
	Create(Server)
	Update(Server)
	Set(Server)
	Get(Server, string, string, string) error
	ToPrettyJSON() string
	ToPrettyYAML() string
	GetType() string
}

func HandleHTTPStatusCode(statusCode int) ServerStatus {
	switch statusCode {
	case 404:
		return OBJECT_NOT_FOUND
	case 401:
		return BAD_AUTH
	case 400:
		return BAD_REQUEST
	default:
		return SUCCESS
	}
}

type DPObjectCommon struct {
	InternalState map[string]interface{}
	ObjectType    string
}

func (dpObjectCommon DPObjectCommon) Create(server Server) {
	// Do nothing
}

func (dpObjectCommon DPObjectCommon) Update(server Server) {
	// Do nothing
}

func (dpObjectCommon DPObjectCommon) Set(server Server) {
	// Do nothing
}

func (dpObjectCommon *DPObjectCommon) Get(server Server, domain string, name string) ServerStatus {
	requestUrl := fmt.Sprintf("https://%s/mgmt/config/%s/%s/%s", server.Hostname, domain, dpObjectCommon.ObjectType, name)
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Println("Could not create server request: ", err)
		return REQUEST_CREATE_FAILED
	}
	basicAuthValue := b64.StdEncoding.EncodeToString([]byte(server.Username + ":" + server.Password))
	req.Header.Set("Authorization", "Basic "+basicAuthValue)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Bad TCP connect: ", err)
		return BAD_TCP_CONNECT
	}
	log.Println(res.Status)
	log.Println(res.StatusCode)
	serverStatus := HandleHTTPStatusCode(res.StatusCode)
	if serverStatus != SUCCESS {
		return serverStatus
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Connected but could not read the response: ", err)
		return RESPONSE_READ
	}
	log.Printf("client: response body: %s\n", resBody)
	var data map[string]interface{} = nil
	json.Unmarshal(resBody, &data)
	links := data["_links"].(map[string]interface{})
	self := links["self"].(map[string]interface{})
	href := self["href"].(string)
	var fetchedState map[string]interface{} = data[dpObjectCommon.ObjectType].(map[string]interface{})
	dpObjectCommon.InternalState = fetchedState
	log.Println(href)
	return SUCCESS
}

func (dpObjectCommon DPObjectCommon) ToPrettyJSON() string {
	json, err := json.MarshalIndent(dpObjectCommon.InternalState, "", "  ")
	if err != nil {
		panic(err)
	} else {
		return string(json[:])
	}
}

func (dpObjectCommon DPObjectCommon) ToPrettyYAML() string {
	yaml, err := yaml.Marshal(dpObjectCommon.InternalState)
	if err != nil {
		panic(err)
	}
	return string(yaml[:])
}
