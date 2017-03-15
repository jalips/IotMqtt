IotMqqt
=======================

A Go project created on February 21, 2017.

## Description

* Subscribe a sensor via IotApi
* Push a message via mqtt
* Receive a message via mqtt
* Post a requet to IotApi and register a message

## Features

* Go 1.7
* Mosquitto 1.4
* [yosssi/gmq](https://github.com/yosssi/gmq)
* [franela/goreq](https://github.com/franela/goreq)

## Install

Clone the project:
```bash
go get github.com/miroufff/IotMqtt
```

Change Ip in commonvariables.go:
```bash
var IpServ string = "0.0.0.0"
```

## Usage

Run:
```bash
go run main.go
```
