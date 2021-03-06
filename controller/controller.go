package controller

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"log"
)

func setCORSHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
}

type JSONRequestWrapper map[string]interface{}

type RESTWrapper map[string]interface{}

type Links struct {
	Self string `json:"self"`
	Next string `json:"next"`
	Last string `json:"last"`
}

func NewWrapper() (*RESTWrapper) {
	restWrapper := RESTWrapper(map[string]interface{} {
		"links": new(Links),
	})
	return &restWrapper
}

func (wrapper *JSONRequestWrapper) Content(key string, data interface{}) {
	wrapperMap := map[string]interface{}(*wrapper)
	wrapperMap[key] = data
}

func (wrapper *RESTWrapper) setSelf(value string) {
	wrapperMap := map[string]interface{}(*wrapper)
	links := wrapperMap["links"].(*Links)
	links.Self = value
}

func (wrapper *RESTWrapper) Self() string {
	wrapperMap := map[string]interface{}(*wrapper)
	links := wrapperMap["links"].(*Links)
	return links.Self
}

func (wrapper *RESTWrapper) setData(key string, data interface{}) {
	wrapperMap := map[string]interface{}(*wrapper)
	wrapperMap[key] = data
}

func HandleOptionsRequest(c *gin.Context)  {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "POST,PUT,DELETE")
	c.String(http.StatusOK, "")
}

func BadRequestResponse(c *gin.Context, err error) {
	log.Printf("There are errs: %s", err)
	c.JSON(http.StatusBadRequest, err)
}

func assembleQuery(c *gin.Context) (query map[string]interface{}, ok bool) {
	query = make(map[string]interface{})
	if c.Query("busID") != "" {
		query["bus_id"] = c.Query("busID")
		ok = true
	}

	return query, ok
}

func wrapperResponseJson(url, jsonName string, data interface{}) RESTWrapper {
	wrapper := NewWrapper();
	wrapper.setSelf(url)
	wrapper.setData(jsonName, data)
	return *wrapper
}

