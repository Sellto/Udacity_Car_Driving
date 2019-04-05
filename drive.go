package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/googollee/go-socket.io"
	"strconv"
	// "encoding/base64"
	// "image"
	// "strings"
	// _ "image/jpeg"
	)

//Send data in right format to the car simulation
func send(s socketio.Conn,steer float64,throttle float64){
	s.Emit("steer",map[string]string{"steering_angle": strconv.FormatFloat(steer, 'E', -1, 64), "throttle": strconv.FormatFloat(throttle, 'E', -1, 64)})
}


func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	//Action made by the client at the connection
	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		send(s,0,0)
		return nil
	})

	//Error handler
	server.OnError("/", func(e error) {
		fmt.Println("meet error:", e)
	})

	//Disconnect handler
	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})

	//path used by the car to send data.
	server.OnEvent("/", "telemetry", func(s socketio.Conn, msg map[string]string) {
		send(s,0,10)
		//reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(msg["image"]))
		//Image,_,_ := image.Decode(reader)
		//im := GetImageFeature(Image)
	})

	go server.Serve()
	defer server.Close()

	// http.HandleFunc("/view", func (w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, test) } )
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:4567...")
	log.Fatal(http.ListenAndServe(":4567", nil))
}
