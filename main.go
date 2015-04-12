package main

import (
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/on", ledOn)
	router.HandleFunc("/off", ledOff)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.Handle("/", router)
	http.ListenAndServe(":"+port, nil)
}

func ledOn(res http.ResponseWriter, req *http.Request) {
	sendCommand("led-on")

}

func ledOff(res http.ResponseWriter, req *http.Request) {
	sendCommand("ledoff")
}

func sendCommand(command string) {
	ip, port := "192.168.1.69", "333"
	token := os.Getenv("HOMEBASE_TOKEN")
	// ip, port := "10.0.1.215", "333"

	tvAddr, err := net.ResolveTCPAddr("tcp", ip+":"+port)
	if err != nil {
		println("Could not resolve", ip, "on", port)
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tvAddr)
	if err != nil {
		println("Could not connect to", ip, "on", port)
		os.Exit(1)
	}

	var recv []byte
	conn.Read(recv)
	conn.Write([]byte(token + command + "\r\n"))
}
