package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"motorcycle-rent-api/app/constant"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)

	return r.ResponseWriter.Write(b)
}

func InboundLogger(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record the start time to calculate response time
		startTime := time.Now()

		// Capture the request body
		var requestBodyBuffer bytes.Buffer
		requestBodyTeeReader := io.TeeReader(c.Request.Body, &requestBodyBuffer)
		rawRequestBody, err := io.ReadAll(requestBodyTeeReader)
		if err != nil {
			logger.Warn("Failed to read request body:", err)
		}
		// Restore the request body for further processing
		c.Request.Body = io.NopCloser(&requestBodyBuffer)

		// Wrap the response writer to capture the response body
		responseBodyBuffer := &responseBodyWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = responseBodyBuffer

		// Process the request
		c.Next()

		// Calculate the total response time
		responseTime := time.Since(startTime)

		// Prepare the request data for logging
		var requestData any
		if c.Request.Method == http.MethodGet {
			// For GET requests, log the raw query string
			requestData = string(rawRequestBody)
		} else {
			// For other methods, check if the request body is valid JSON
			if json.Valid(rawRequestBody) {
				requestData = json.RawMessage(rawRequestBody)
			} else {
				// Fallback for non-JSON request bodies
				requestData = string(rawRequestBody)
			}
		}

		// Prepare the response body for logging
		var responseData any
		if json.Valid(responseBodyBuffer.body.Bytes()) {
			// If the response body is valid JSON, log it as-is
			responseData = json.RawMessage(responseBodyBuffer.body.Bytes())
		} else {
			// Fallback for non-JSON responses
			responseData, _ = json.Marshal(responseBodyBuffer.body.String())
		}

		// Log all request/response details using structured fields
		logger.WithFields(logrus.Fields{
			"api_call_id":      c.GetString(constant.RequestIDKey),
			"function":         c.HandlerName(),
			"http_method":      c.Request.Method,
			"request_path":     c.Request.RequestURI,
			"query_parameters": c.Request.URL.Query(),
			"request_headers":  c.Request.Header,
			"request_body":     requestData,
			"client_ip":        c.ClientIP(),
			"user_agent":       c.Request.UserAgent(),
			"response_time":    responseTime.String(),
			"status_code":      c.Writer.Status(),
			"response_body":    responseData,
			"type":             "Inbound",
		}).Info("Inbound Request processed")
	}
}
