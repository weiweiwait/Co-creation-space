package config

import "my_project/project-common/kk"

var kw *kk.KafkaWriter

func InitKafkaWriter() func() {
	kw = kk.GetWriter("localhost:9092")
	return kw.Close
}

func SendLog(data []byte) {
	kw.Send(kk.LogData{
		Topic: "msproject_log",
		Data:  data,
	})
}
