package logger

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

// InitLogger initializes a logger that writes to a file and console
func InitLogger() {
	file, err := os.OpenFile("gateway.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.Println("Logger initialized")
}

// loggingResponseWriter is a wrapper to capture status code
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the HTTP status code
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// LoggerMiddleware logs details of each HTTP request
func LoggerMiddleware(next http.Handler) http.Handler {
	fmt.Println("logger middle inainte de return")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("logger middle dupa return")
		// Wrap the ResponseWriter to capture the status code
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(lrw, r) // Call the next handler

		duration := time.Since(start)

		log.Printf("[API GATEWAY] Method: %s | Path: %s | Status: %d | Duration: %v",
			r.Method, r.URL.Path, lrw.statusCode, duration)
	})
}

func addCallerData(message string) string {
	var file string
	var line int
	var ok bool

	_, file, line, ok = runtime.Caller(2)

	if ok == false {
		return "Unable to get caller data for message: " + message
	}

	return fmt.Sprintf("%s:%d %s", file, line, message)
}

func writeToFile(message string) {
	log.Println(message)
}

func Info(message string) {
	logData := addCallerData("INFO " + message)

	writeToFile(logData)
}

func Error(err error) {
	logData := addCallerData("ERROR " + err.Error())

	writeToFile(logData)
}

func Panic(err error) {
	logData := addCallerData("PANIC " + err.Error())

	writeToFile(logData)
	panic(err)
}

func Warning(message string) {
	logData := addCallerData("WARNING " + message)

	writeToFile(logData)
}
