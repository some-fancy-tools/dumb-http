package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// LogPattern is format for writing the logRecord
	LogPattern = "%s - - [%s] \"%s\" %d %d %s \"%s\" \"%s\"\n"
	// DateTimeFormat is for logging Date and Time of request
	DateTimeFormat = "02/Jan/2006 15:04:05" // Minimal
	// DateTimeFormat = "Mon, 02 Jan 2006 15:04:05 MST" // Full
)

// LogRecord struct extends http.ResponseWriter interface which has different methods included in it.
type LogRecord struct {
	http.ResponseWriter
	method        string
	path          string
	clientIP      string
	time          time.Time
	contentLength int64
	protocol      string
	statusCode    int
	referer       string
	userAgent     string
	duration      time.Duration
}

// Log method to be called for logging to "out"
func (l *LogRecord) Log(out io.Writer) {
	timeFormatted := l.time.Format(DateTimeFormat)
	requestLine := l.method + " " + l.path + " " + l.protocol
	fmt.Fprintf(out, LogPattern, l.clientIP, timeFormatted, requestLine,
		l.statusCode, l.contentLength, l.duration.String(), l.referer, l.userAgent)
}

// WriteHeader method has been extended to record status code from previous handler.
func (l *LogRecord) WriteHeader(statusCode int) {
	l.statusCode = statusCode
	l.ResponseWriter.WriteHeader(statusCode)
}

// Write method has been extended to record the bytes written or the content length
func (l *LogRecord) Write(p []byte) (int, error) {
	written, err := l.ResponseWriter.Write(p)
	l.contentLength += int64(written)
	return written, err
}
