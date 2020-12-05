package adafruitio

import (
	"fmt"
	"net"

	"github.com/jeffallen/mqtt"
	proto "github.com/huin/mqtt"
)

type MQTTClient struct {
	// Base MQTT client to talk to io.adafruit.com
	mqttclient *mqtt.ClientConn
	conn       *net.Conn

	// Base URL for API requests. Defaults to public adafruit io URL
	BaseURL string

	//Authentication
	APIUsername string
	APIKey      string

	// Services that make up adafruit io.
	Data  *DataService
	Feed  *FeedService
	Group *GroupService
}

func NewMQTTClient(username string, key string) *MQTTClient {
	c := &MQTTClient{APIUsername: username}
	c.APIKey = key
	c.APIUsername = username
	conn, _ := net.Dial("tcp", c.BaseURL)
	c.mqttclient = mqtt.NewClientConn(conn)
	if err := c.mqttclient.Connect(c.APIUsername, c.APIKey) {
		panic(err)
	}
	c.Data = &DataService{}
	c.Feed = &FeedService{}
	c.Group = &GroupService{}
	return c
}

func (c *MQTTClient) SetFeed(feed *Feed) {
	c.Feed.CurrentFeed = feed
}

func (c *MQTTClient) checkFeed() error {
	if c.Feed.CurrentFeed == nil {
		return fmt.Errorf("CurrentFeed must be set")
	}
	return nil
}

func (c *MQTTClient) Publish(datatype something) error {
	//TODO
	return nil
}

func (c *MQTTClient) Subscribe(channel string, callback func) error {
	//TODO
	return nil
}