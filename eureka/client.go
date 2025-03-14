package eureka

import "github.com/ArthurHlt/go-eureka-client/eureka"

type Client interface {
	Register(appId string, instanceInfo *eureka.InstanceInfo) error
	SendHeartbeat(appId string, instanceId string) error
	CreateInstance() *eureka.InstanceInfo
}
