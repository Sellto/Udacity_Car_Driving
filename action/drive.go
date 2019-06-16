package action

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
	 "../lib"
	 "strings"
	)

	//Send data in right format to the car simulation
	func send(s socketio.Conn,steer float64,throttle float64){
		s.Emit("steer",map[string]string{"steering_angle": strconv.FormatFloat(steer, 'E', -1, 64), "throttle": strconv.FormatFloat(throttle, 'E', -1, 64)})
	}



	func RunServer(mode int,model string) {

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
      steer := 0.0
      listdata := []float64{}
			test := [][]float64{}
      switch mode {
			  case 0:
				 test = cp.ScanImage(HOG64,64,24,Image,true,model)
				 if test[0][0] == -1 && test[0][1] == -1 && test[0][2] == -1 {
					 steer = 0.05
				 }
				 if test[0][0] == -1 && test[0][1] == -1 && test[0][2] == 1 {
					 steer = 0.05
				 }
				 if test[0][0] == -1 && test[0][1] == 1 && test[0][2] == 1 {
					 steer = -0.0
				 }
				 if test[0][0] == -1 && test[0][1] == 1 && test[0][2] == -1 {
					 steer = -0.05
				 }
				 if test[0][0] == 1 && test[0][1] == 1 && test[0][2] == -1 {
					 steer = -0.05
				 }
				 if test[0][0] == 1 && test[0][1] == -1 && test[0][2] == -1 {
					 steer = -0.1
				 }
        case 2:
          listdata,_ = cp.ComputeHoG4All(Image,HOGALL)
					convert := make(map[int]float64)
					for i,v := range listdata {
						convert[i]= float64(int(v))
					}
					steer = cp.PredictFrom(convert,model)/1000
        case 3:
          listdata = cp.GetImageGSFeature(Image)
					convert := make(map[int]float64)
					for i,v := range listdata {
						convert[i]= float64(int(v))
					}
					steer = cp.PredictFrom(convert,model)/1000
        case 4:
          listdata = cp.GetImageCustGSFeature(Image)
					convert := make(map[int]float64)
					for i,v := range listdata {
						convert[i]= float64(int(v))
					}
					steer = cp.PredictFrom(convert,model)/1000
        case 5:
          listdata = cp.GetImageBWFeature(Image)
					convert := make(map[int]float64)
					for i,v := range listdata {
						convert[i]= float64(int(v))
					}
					steer = cp.PredictFrom(convert,model)/1000
	       default:
            fmt.Println(mode)
	    }
			fmt.Println(test,steer)
			sp,_ := strconv.ParseFloat(msg["speed"],64)
			if sp > 10{
				send(s,steer,0)
			} else {
				send(s,steer,0.2)
			}

		})

		go server.Serve()
		defer server.Close()
		http.Handle("/socket.io/", server)
		http.Handle("/", http.FileServer(http.Dir("./asset")))
		fmt.Println("Serving at localhost:4567")
    fmt.Println("Configuration mode : ",mode)
    fmt.Println("Model used : ",model)
		log.Fatal(http.ListenAndServe(":4567", nil))
	}
