package common

type SensorData struct {
	data string `json:"data"`
	statisticType string `json:"type"`
}

type Sensor struct {
	DisplayName string `json:"displayname"`
	Vendor string `json:"vendor"`
	Product string `json:"product"`
	Version int `json:"version"`
}