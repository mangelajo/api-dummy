package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostName, _ := os.Hostname()
	podName := os.Getenv("POD_NAME")
	clusterName := os.Getenv("CLUSTER_NAME")
	fmt.Fprintln(w, "api-dummy responder 🤖:")
	fmt.Fprintf(w, " @ pod     : %s\n", podName)
	fmt.Fprintf(w, " @ cluster : %s\n", clusterName)
	fmt.Fprintf(w, " @ ip      : %s\n", getLocalIP())
	fmt.Fprintf(w, " @ hostname: %s\n", hostName)
}

func handleRequests() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	handleRequests()
}

func getLocalIP() string {
	return getLocalIPForDestination("8.8.8.8")
}

func getLocalIPForDestination(dst string) string {
	conn, err := net.Dial("udp", dst+":53")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
