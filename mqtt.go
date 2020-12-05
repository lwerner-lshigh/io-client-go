package adafruitio

import (
	"github.com/jeffallen/mqtt"
	"fmt"
	"runtime"
	"net"
)

type MQTTClient struct {
	// Base MQTT client to talk to io.adafruit.com
	mqttclient *mqtt.ClientConn

	// Base URL for API requests. Defaults to public adafruit io URL
	BaseURL string

	//Authentication
	APIUsername string
	APIKey string
	

	// Services that make up adafruit io.
	Data  *DataService
	Feed  *FeedService
	Group *GroupService
}

func NewMQTTClient(username string, key string) *MQTTClient {
	c := &MQTTClient{APIUsername: username}
	c.APIKey = key
	c.userAgent = fmt.Sprintf("AdafruitIO-go/%v (%v %v)", Version, runtime.GOOS, runtime.Version())
	conn, _ := net.Dial("tcp", c.BaseURL)
	c.mqttclient := mqtt.NewClientConn(conn)
	
}
