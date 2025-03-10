package dpobjects

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Server struct {
	Hostname string
	Username string
	Password string
}

func TestConnection(server Server) {
	requestUrl := fmt.Sprintf("https://%s", server.Hostname)
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Println("Bad connect")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Bad request: ", err)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Bad read")
	}
	log.Printf("client: response body: %s\n", resBody)
}

type BaseObject interface {
	Create(Server)
	Update(Server)
	Set(Server)
	Get(Server)
	ToPrettyJSON() string
	ToPrettyYAML() string
}
