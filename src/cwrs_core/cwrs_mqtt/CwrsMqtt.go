package cwrs_mqtt

import (
	"bytes"
	"cwrs_go_server/src/cwrs_core/cwrs_viper"
	"cwrs_go_server/src/cwrs_core/cwrs_zap_logger"
	"encoding/binary"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

var log = cwrs_zap_logger.ZapLogger
var client mqtt.Client

// 初始化MQTT服务
//func init() {
//	InitMqttClient()
//}

func InitMqttClient() {
	var host = cwrs_viper.GlobalViper.GetString("mqtt.path")
	var port = cwrs_viper.GlobalViper.GetString("mqtt.port")
	var userName = cwrs_viper.GlobalViper.GetString("mqtt.user-name")
	var password = cwrs_viper.GlobalViper.GetString("mqtt.password")
	//var clientId = time.Now().UnixNano()
	var clientId = cwrs_viper.GlobalViper.GetString("mqtt.clientId") + "-" + userName
	if client == nil {
		opts := mqtt.NewClientOptions()
		//服务器地址格式: wss://域名或IP/mqtt/:port
		opts.AddBroker(fmt.Sprintf("%s:%s", host, port)) // 这个中转服务器不需要任何账号密码
		opts.SetClientID(fmt.Sprintf(clientId))
		opts.SetUsername(userName)
		opts.SetPassword(password)
		opts.OnConnect = func(c mqtt.Client) {
			log.Info("MQTT连接成功!")
			//批量订阅消息
			BatchSubscribe()

			tempTest(client)
		}
		opts.OnConnectionLost = func(c mqtt.Client, err error) {
			log.Error("MQTT断开连接:", zap.Error(err))
			fmt.Println("MQTT尝试重新连接！")
			InitMqttClient()
		}
		opts.SetDefaultPublishHandler(mqttMessageHandler)
		client = mqtt.NewClient(opts)
	}

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Error("MQTT连接失败:", zap.Error(token.Error()))
		panic(token.Error())
	}

}

// 发布消息
func ClientPublish(topic string, qos byte, retained bool, payload interface{}) {
	token := client.Publish(topic, qos, retained, payload)
	token.Wait()
	if token.Error() != nil {
		log.Error("MQTT消息发布失败!", zap.Error(token.Error()))
	}
}

// 订阅消息
func ClientSubscribe(topic string, qos byte) {
	token := client.Subscribe(topic, qos, nil)
	token.Wait()
	if token.Error() != nil {
		log.Error("MQTT订阅消息失败!", zap.Error(token.Error()))
	}
}

func tempTest(client mqtt.Client) {
	var topic = "/ceshi/swg"

	nb := bytes.NewBuffer([]byte{})
	binary.Write(nb, binary.BigEndian, byte(1))
	binary.Write(nb, binary.BigEndian, byte(98))

	client.Publish(topic, 0, true, nb.Bytes())

}

// 获取mqttClient
func GetSysMqttClient() mqtt.Client {
	return client
}

// 接收消息 具体见receiveMqttMessage方法
var mqttMessageHandler mqtt.MessageHandler = receiveMqttMessage
