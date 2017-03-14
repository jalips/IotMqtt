package main

import (
	"fmt"
	"os"
	"os/signal"

	"encoding/json"
	"time"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
	"github.com/franela/goreq"
	"./common"
)

/*
Subscribe to the api and publish a message to the main
This publish can be switched by anything while the sensor is subscribed and the message is publish on the good topic
 */
func main() {
	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	sensor := &common.Sensor{
		DisplayName: "Rasp-Sensor",
		Vendor: "Raspberry Foundation",
		Product: "Pi",
		Version: 3,
	}

	jsonsensor, err := json.Marshal(sensor)
	if err != nil {
		fmt.Println(err)
		return
	}

	request := goreq.Request{
		Method: "POST",
		Uri: "http://" + common.IpApiServ + "/api/sensors?sender=go",
		Accept: "application/json",
		ContentType: "application/json",
		UserAgent: "goreq",
		Body: string(jsonsensor),
	}
	resguid, err := request.Do()
	if err != nil {
		panic(err)
	}

	fmt.Println(resguid.Header)
	fmt.Println(resguid.Body.ToString())

	var guid interface{}
	resguid.Body.FromJsonTo(guid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(guid)

	// Create an MQTT Client.
	cli := client.New(&client.Options{
		// Define the processing of the error handler.
		ErrorHandler: func(err error) {
			fmt.Println(err)
		},
	})

	// Terminate the Client.
	defer cli.Terminate()

	// Connect to the MQTT Server.
	err = cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address:  common.IpMosquitoServ,
		ClientID: []byte("rasp-client"),
	})
	if err != nil {
		panic(err)
	}

	data := &common.SensorData{
		SensorName:  "truc",
		Measurement: "temp",
		Time:        time.Now().UnixNano(),
		Value:       "OVER 9000",
	}

	jsonitem, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Publish a message.
	err = cli.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS0,
		TopicName: []byte("temp"),
		Message:   []byte(jsonitem),
	})
	if err != nil {
		panic(err)
	}

	// Wait for receiving a signal.
	<-sigc

	// Disconnect the Network Connection.
	if err := cli.Disconnect(); err != nil {
		panic(err)
	}
}
