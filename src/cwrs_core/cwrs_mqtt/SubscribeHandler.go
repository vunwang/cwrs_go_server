package cwrs_mqtt

// 对所需主题进行批量订阅
func BatchSubscribe() {

	//订阅消息
	ClientSubscribe("/+/test", 1)

}
