package main

import (
	"fmt"
	"os"
	"os/signal"
	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
	"github.com/franela/goreq"
	"github.com/jalips/IotMqtt/common"
	//"time"
	"encoding/json"
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

	// Connect to the MQTT Server.
	err := cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address: common.IpMosquitoServ,
		ClientID: []byte("vm-client"),
	})
	if err != nil {
		panic(err)
	}

	// Subscribe to topics.
	err = cli.Subscribe(&client.SubscribeOptions{
		SubReqs: []*client.SubReq{
			&client.SubReq{
				TopicFilter: []byte("temp"),
				QoS:         mqtt.QoS0,
				// Define the processing of the message handler.
				Handler: sensorDataHandler,
			},
			&client.SubReq{
				TopicFilter: []byte("hydro"),
				QoS:         mqtt.QoS1,
				Handler: sensorDataHandler,
			},
			&client.SubReq{
				TopicFilter: []byte("valve"),
				QoS:         mqtt.QoS2,
				Handler: sensorDataHandler,
			},
		},
	})
	if err != nil {
		panic(err)
	}

	// Wait for receiving a signal.
	<-sigc

	// Unsubscribe from topics.
	err = cli.Unsubscribe(&client.UnsubscribeOptions{
		TopicFilters: [][]byte{
			[]byte("temp"),
		},
	})
	if err != nil {
		panic(err)
	}

	// Disconnect the Network Connection.
	if err := cli.Disconnect(); err != nil {
		panic(err)
	}

}

func sensorDataHandler(topicName, message []byte) {
	fmt.Println(string(topicName), string(message))

	// If the message publish on the good topic
	/*
	if (string(topicName) != "sensor/temp") {
		panic("wrong message")
	}
	*/

	data := &common.SensorData{
		Data:  string(message),
		StatisticType: string(topicName),
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
		Uri: "http://autoyas.jalips-test.fr/app.php/statistics/"+string(message)+"/"+string(topicName)+"/new",
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

}