package middleware

import (
	"Ramdoni007/Learn-Golang-RestFullAPI/helper"
	"Ramdoni007/Learn-Golang-RestFullAPI/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler)*AuthMiddleware  {
	return &AuthMiddleware{Handler: handler}
}



func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		//ok
		middleware.Handler.ServeHTTP(writer,request)
	}else {
		//error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "HEY BRO ! UNAUTHORIZED",
		}

		helper.WriterToResponseBody(writer, webResponse)


	}
}
