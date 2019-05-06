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
		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(msg["image"]))
		Image,_,_ := image.Decode(reader)
		listpix := cp.GetImageCropCenterFeature(Image)
		sp,_ := strconv.ParseInt(msg["speed"],10,8)
		listpix = append(listpix,uint8(sp))
		convert := make(map[int]float64)
		for i,v := range listpix {
			convert[i]= float64(v)
		}
		prediction := cp.PredictFrom(convert,"data5.model_POLY")
		fmt.Println(prediction)
		steer, thro := cp.DecodeLabel(prediction)
		fmt.Println(steer,":",thro)
		send(s,steer,0.2)
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
