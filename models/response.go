package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	rw          http.ResponseWriter
}

type ResponseToken struct {
	Token string `json:"token"`
}

func createDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		contentType: "application/json",
		rw:          rw,
	}
}

func (resp *Response) send() {
	resp.rw.WriteHeader(resp.Status)
	resp.rw.Header().Add("Content-Type", resp.contentType)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.rw, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}) {
	resp := createDefaultResponse(rw)
	resp.Data = data
	resp.send()
}

func (resp *Response) notFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource not found"
}

func SendNotFound(rw http.ResponseWriter) {
	resp := createDefaultResponse(rw)
	resp.notFound()
	resp.send()
}

func (resp *Response) unprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "Resource unprocessable entity"
}

func SendUnprocessableEntity(rw http.ResponseWriter) {
	resp := createDefaultResponse(rw)
	resp.unprocessableEntity()
	resp.send()
}

func (resp *Response) unauthorized() {
	resp.Status = http.StatusUnauthorized
	resp.Message = "Unauthorized"
}

func SendUnauthorized(rw http.ResponseWriter) {
	resp := createDefaultResponse(rw)
	resp.unauthorized()
	resp.send()
}
