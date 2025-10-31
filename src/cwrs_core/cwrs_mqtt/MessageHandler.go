package cwrs_mqtt

import (
	"bytes"
	"encoding/binary"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"strings"
	"time"
)

// 接收MQTT消息
func receiveMqttMessage(client mqtt.Client, msg mqtt.Message) {
	//添加日志
	log.Info("MQTT接收消息:", zap.String("主题", msg.Topic()), zap.String("消息", string(msg.Payload())))

	//解析示例如下：
	//解析主题
	topicArr := strings.Split(msg.Topic(), "/")

	//解析payload
	payload := msg.Payload()
	nb := bytes.NewBuffer(payload)

	var test uint16
	err := binary.Read(nb, binary.BigEndian, &test)
	if err != nil {
		log.Error(fmt.Sprintf("解析失败：主题:%s error:%+v", msg.Topic(), err))
		return
	}

	fmt.Printf("topicArr:%s,payload：%s,test:%d\n", topicArr, payload, test)

	fmt.Println(" mqtt 消息结束---------------------", time.Now().Format("2006-01-02 15:04:05"))
}
