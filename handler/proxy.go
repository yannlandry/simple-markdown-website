package handler

import (
	"bytes"
	"io"
	"net/http"
)

func (this *HandlerFactory) Proxy(forward string) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		// Read body of incoming request.
		requestBody, err := io.ReadAll(request.Body)
		if err != nil {
			Error(this.errors, this.configuration, 500, response, request)
			return
		}

		// Create proxied request with header and shallow-copy headers.
		proxyRequest, err := http.NewRequest(request.Method, forward, bytes.NewReader(requestBody))
		proxyRequest.Header = request.Header

		// Send request to next server.
		client := &http.Client{}
		proxyResponse, err := client.Do(proxyRequest)
		if err != nil {
			Error(this.errors, this.configuration, 502, response, request)
			return
		}
		defer proxyResponse.Body.Close()

		// Send response to client.
		responseBody, err := io.ReadAll(proxyResponse.Body)
		response.WriteHeader(proxyResponse.StatusCode)
		for key, values := range proxyResponse.Header {
			for _, value := range values {
				response.Header().Set(key, value)
			}
		}
		response.Write(responseBody)
	}
}
