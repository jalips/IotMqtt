package main

import (
	"fmt"
	"os"
	"os/signal"
	//"encoding/json"
	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
	"github.com/franela/goreq"
	"github.com/jalips/IotMqtt/common"
	//"time"
)

/**
Listen mqtt messages and push to the IotSymfonyApi

TODO : Get the uuid of the sensor
 */
func main() {
	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	// Create an MQTT Client.
	cli := client.New(&client.Options{
		// Define the processing of the error handler.
		ErrorHandler: func(err error) {
			fmt.Println(err)
		},
	})

	// Terminate the Client.
	cli.Terminate()

	/************************************
				SENSOR 1
	*************************************/
	// Connect to the MQTT Server.
	err := cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address: common.IpMosquitoServ,
		ClientID: []byte("vm-client"),
	})
	if err != nil {
		panic(err)
	}

	/************************************
				SENSOR 1
	*************************************/
	// Subscribe to topics.
	err = cli.Subscribe(&client.SubscribeOptions{
		SubReqs: []*client.SubReq{
			&client.SubReq{
				TopicFilter: []byte("sensor/temp"),
				QoS:         mqtt.QoS0,
				// Define the processing of the message handler.
				Handler: sensorDataHandlerTemp,
			},
		},
	})
	if err != nil {
		panic(err)
	}

	/************************************
				SENSOR 2
	*************************************/
	// Connect to the MQTT Server.
	err2 := cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address: common.IpMosquitoServ,
		ClientID: []byte("vm-client"),
	})
	if err2 != nil {
		panic(err2)
	}

	/************************************
				SENSOR 2
	*************************************/
	// Subscribe to topics.
	err2 = cli.Subscribe(&client.SubscribeOptions{
		SubReqs: []*client.SubReq{
			&client.SubReq{
				TopicFilter: []byte("sensor/hydro"),
				QoS:         mqtt.QoS0,
				// Define the processing of the message handler.
				Handler: sensorDataHandlerHydro,
			},
		},
	})
	if err2 != nil {
		panic(err2)
	}

	// Wait for receiving a signal.
	<-sigc

	/************************************
				SENSOR 1
	*************************************/
	// Unsubscribe from topics.
	err = cli.Unsubscribe(&client.UnsubscribeOptions{
		TopicFilters: [][]byte{
			[]byte("sensor/temp"),
		},
	})
	if err != nil {
		panic(err)
	}

	// Disconnect the Network Connection.
	if err := cli.Disconnect(); err != nil {
		panic(err)
	}

	/************************************
				SENSOR 2
	*************************************/
	err2 = cli.Unsubscribe(&client.UnsubscribeOptions{
		TopicFilters: [][]byte{
			[]byte("sensor/hydro"),
		},
	})
	if err2 != nil {
		panic(err2)
	}

	// Disconnect the Network Connection.
	if err2 := cli.Disconnect(); err2 != nil {
		panic(err2)
	}
}

func sensorDataHandlerTemp(topicName, message []byte) {
	fmt.Println(string(topicName), string(message))
/*
	// If the message publish on the good topic
	if (string(topicName) != "/temperature") {
		panic("wrong message")
	}

	data := &common.SensorData{
		SensorName:  "c10d2fc4-9361-4c24-91f4-c355379cbf44",
		Measurement: "temp",
		Time:        time.Now().UnixNano(),
		Value:       string(message),
	}

	jsonitem, err := json.Marshal(data)

	fmt.Println(string(jsonitem))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Post request to IotApi api
	request := goreq.Request{
		Method: "POST",
		Uri: "http://" + common.IpApiServ + "/api/datasensors?sender=go",
		Accept: "application/json",
		ContentType: "application/json",
		UserAgent: "goreq",
		Body: string(jsonitem),
	}
	res, err := request.Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Header)
	fmt.Println(res.Body.ToString())
*/
}


func sensorDataHandlerHydro(topicName, message []byte) {
	fmt.Println(string(topicName), string(message))
	/*
		// If the message publish on the good topic
		if (string(topicName) != "/temperature") {
			panic("wrong message")
		}

		data := &common.SensorData{
			SensorName:  "c10d2fc4-9361-4c24-91f4-c355379cbf44",
			Measurement: "temp",
			Time:        time.Now().UnixNano(),
			Value:       string(message),
		}

		jsonitem, err := json.Marshal(data)

		fmt.Println(string(jsonitem))
		if err != nil {
			fmt.Println(err)
			return
		}

		// Post request to IotApi api
		request := goreq.Request{
			Method: "POST",
			Uri: "http://" + common.IpApiServ + "/api/datasensors?sender=go",
			Accept: "application/json",
			ContentType: "application/json",
			UserAgent: "goreq",
			Body: string(jsonitem),
		}
		res, err := request.Do()
		if err != nil {
			panic(err)
		}

		fmt.Println(res.Header)
		fmt.Println(res.Body.ToString())
	*/
}