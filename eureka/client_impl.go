package eureka

import (
	"log"
	"time"

	"github.com/ArthurHlt/go-eureka-client/eureka"
)

type clientImpl struct {
	client *eureka.Client
}

type configEureka struct {
	ServiceUrls           []string
	App                   string
	InstanceID            string
	HostName              string
	Port                  *eureka.Port
	RenewalIntervalInSecs int
	DurationInSecs        int
}

func NewClient(machines []string) Client {
	return &clientImpl{
		client: eureka.NewClient(machines),
	}
}

func (c *clientImpl) CreateInstance() *eureka.InstanceInfo {
	instance := eureka.NewInstanceInfo("test.com", "test", "69.172.200.235", 80, 30, false) //Create a new instance to register
	instance.Metadata = &eureka.MetaData{
		Map: make(map[string]string),
	}
	instance.Metadata.Map["foo"] = "bar" //add metadata for example

	return instance
}

func (c *clientImpl) Register(appId string, instanceInfo *eureka.InstanceInfo) error {
	return c.client.RegisterInstance(appId, instanceInfo)
}

func (c *clientImpl) SendHeartbeat(appId string, instanceId string) error {
	return c.client.SendHeartbeat(appId, instanceId)
}

func RegisterWithEureka() error {
	config := configEureka{
		App:                   "api-gateway",
		InstanceID:            "api-gateway-1",
		HostName:              "localhost",
		Port:                  &eureka.Port{Port: 8080, Enabled: true},
		RenewalIntervalInSecs: 30,
		DurationInSecs:        90,
		ServiceUrls:           []string{"http://localhost:8761/eureka/"},
	}

	client := NewClient(config.ServiceUrls)
	instance := client.CreateInstance()
	err := client.Register(instance.App, instance)
	if err != nil {
		log.Fatalf("Failed to register with Eureka: %v", err)
	}

	go func() {
		for {
			time.Sleep(time.Duration(config.RenewalIntervalInSecs) * time.Second)
			err := client.SendHeartbeat(instance.App, instance.HostName)
			if err != nil {
				log.Printf("Failed to send heartbeat to Eureka: %v", err)
			}
		}
	}()
	return nil
}
