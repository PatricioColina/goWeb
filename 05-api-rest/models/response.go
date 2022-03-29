package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//estructura que responde al cliente
type Response struct {
	Status      int         `json:status`
	Data        interface{} `json:data`
	Message     string      `json:message`
	contentType string
	respWrithe  http.ResponseWriter
}

func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		respWrithe:  rw,
		contentType: "aplication/jason",
	}
}

func (resp *Response) Send() {
	resp.respWrithe.Header().Set("Content-Type", resp.contentType)
	resp.respWrithe.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.respWrithe, string(output))
}

//responde al cliente
func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

func (resp *Response) NoFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource No Found"
}

//errores
//responde no found
func SendNoFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NoFound()
	response.Send()
}

func (resp *Response) StatusUnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "StatusUnprocessableEntity No Found"
}

func SendStatusUnprocessableEntity(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.StatusUnprocessableEntity()
	response.Send()
}
