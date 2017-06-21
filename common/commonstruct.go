package common

type SensorData struct {
	Data string `json:"data"`
	StatisticType string `json:"statisticType"`
}

type Sensor struct {
	DisplayName string `json:"displayname"`
	Vendor string `json:"vendor"`
	Product string `json:"product"`
	Version int `json:"version"`
}