package main

// Server communicating with autonomous car simulator client
// task - Connect to the client (socketio)
// Task - Receiving camera image and speed value
// Task - Sending throttle and steer command

import (
	"fmt"
	"log"
	"net/http"
	"github.com/googollee/go-socket.io"
	"strconv"
	 "encoding/base64"
	 "image"
	 _ "image/jpeg"
	 "./lib"
	 "strings"
	 //"time"
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
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(msg["image"]))
		Image,_,_ := image.Decode(reader)
		listpix := cp.GetImageFeature(Image)
		sp,_ := strconv.ParseFloat(msg["speed"],64)
		convert := make(map[int]float64)
		for i,v := range listpix {
			convert[i]= float64(v)
		}
		steer := cp.PredictFrom(convert,"sample_steer.model")/100
		//thro  := cp.PredictFrom(convert,"sample_thro.model")/10
		fmt.Println(steer)
		if sp > 15 {
			send(s,steer,0)
		} else {
			send(s,steer,0.4)
		}

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
