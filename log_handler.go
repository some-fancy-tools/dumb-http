package main

import (
	"crypto/subtle"
	"fmt"
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
		duration:       time.Duration(0),
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

	// For getting duration
	// startTime := time.Now()
	username, password, ok := r.BasicAuth()
	if user == "" || pass == "" {
		SuccessfulResponse(logRecord, h.handler, h.out, r)
		return
	}
	if user != "" && pass != "" && (!ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1) {
		UnsuccessfulResponse(logRecord, h.handler, h.out, w, r)
		return
	}
	SuccessfulResponse(logRecord, h.handler, h.out, r)
}

func SuccessfulResponse(logRecord *LogRecord, h http.Handler, out io.Writer, r *http.Request) {
	startTime := time.Now()
	h.ServeHTTP(logRecord, r)
	endTime := time.Now()
	logRecord.time = endTime.UTC()
	logRecord.duration = endTime.Sub(startTime) / 1000.0
	logRecord.Log(out)
}

func UnsuccessfulResponse(logRecord *LogRecord, h http.Handler, out io.Writer, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
	logRecord.statusCode = http.StatusUnauthorized
	w.WriteHeader(logRecord.statusCode)
	fmt.Fprintln(w, "Unauthorized")
	endTime := time.Now()
	logRecord.time = endTime.UTC()
	logRecord.duration = endTime.Sub(startTime) / 1000.0
	logRecord.Log(out)
}
