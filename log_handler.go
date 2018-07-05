package main

import (
	"io"
	"net/http"
	"strings"
	"time"
)

// LoggingHandler is to be used as wrapper of mux.
type LoggingHandler struct {
	handler http.Handler
	out     io.Writer
}

// NewLoggingHandler for creating a new Logging Handler
func NewLoggingHandler(handler http.Handler, out io.Writer) http.Handler {
	return &LoggingHandler{
		handler: handler,
		out:     out,
	}
}

func (h LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ipPort := strings.Split(r.RemoteAddr, ":")

	logRecord := &LogRecord{
		ResponseWriter: w,
		time:           time.Time{},
		method:         r.Method,
		statusCode:     http.StatusOK,
		protocol:       r.Proto,
		path:           r.RequestURI,
		clientIP:       strings.Join(ipPort[:len(ipPort)-1], ":"),
		referer:        "-",
		userAgent:      "-",
		totalTime:      time.Duration(0),
	}

	// For logging Real IP Address
	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
		logRecord.clientIP = realIP
	}

	// For logging Referer URL
	if referer := r.Header.Get("Referer"); referer != "" {
		logRecord.referer = referer
	}

	// For logging User Agent
	if userAgent := r.Header.Get("User-Agent"); userAgent != "" {
		logRecord.userAgent = userAgent
	}
	startTime := time.Now()
	h.handler.ServeHTTP(logRecord, r)
	endTime := time.Now()
	logRecord.time = endTime.UTC()
	logRecord.totalTime = endTime.Sub(startTime) / 1000.0
	logRecord.Log(h.out)
}
