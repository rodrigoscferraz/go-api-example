package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

const(
	home = "/"
	healthz = "/healthz"
	schedule = "/schedule"
	ping = "/ping"
)

func main (){
	mux := http.NewServeMux()
	mux.Handle(home, &homeHandler{})
	mux.Handle(healthz, &healthzHandler{})
	mux.Handle(ping, &pingHandler{})

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}


	
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !checkPath(home, r.URL.Path){
		http.Error(w, "404: URL Not Found",http.StatusNotFound )	
		return
	}
	ip := GetOutboundIP()
	w.Write([]byte(fmt.Sprintf("Local IP: %s", ip)))
}

type healthzHandler struct{}

func (h *healthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !checkPath(home, r.URL.Path){
		http.Error(w, "404: URL Not Found",http.StatusNotFound )	
		return
	}
	w.Write([]byte("OK"))
}

type pingHandler struct{}

func (h *pingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !checkPath(ping, r.URL.Path){
		http.Error(w, "404: URL Not Found",http.StatusNotFound )	
		return
	}
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	ctx := context.Background()
	err := client.Incr(ctx,"API_COUNT").Err()
	if err != nil{
		fmt.Println(err)
	}
	w.Write([]byte(fmt.Sprintf(
		"Requester IP: %v\nTime: %s\nAPI COUNT: %s",strings.Split(r.RemoteAddr, ":"), time.Now().Format(time.RFC3339), client.Get(ctx, "API_COUNT"))))
}

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)
    return localAddr.IP
}
