package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/guptaachin/tracing-with-open-telemetry/ping-ack-api/docs"
	"github.com/swaggo/http-swagger"
)

// -------------------- Handlers -----------------

// Ping Sends a simple ack message on ping request
// @Summary Ping
// @Description Ping checks API availability
// @Tags ping
// @ID ping
// @Produce plain
// @Success 200 {string} string "Ack"
// @Router /ping [get]
func Ping(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Ack")
}

// DelayPing Sends a simple ack message on ping request
// @Summary DelayPing
// @Description DelayPing sends ping requests to return after random wait time
// @Tags delay
// @ID delay
// @Produce plain
// @Success 200 {string} string "Delay Ack"
// @Router /delay [get]
func DelayPing(w http.ResponseWriter, _ *http.Request) {
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

	fmt.Fprintf(w, "Delayed Ack")
}

// -------------------- URL Mappings -----------------
func urlMappings() {
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/delay", DelayPing)
}

// @title Simple Ping Ack API
// @version 1.0
// @description Simple Ping Ack API to explain Swagger and Observability

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7997
// @BasePath /
func main() {
	// Run server
	fmt.Println("Starting server...........")
	urlMappings()
	http.ListenAndServe(":7997", nil)
}
