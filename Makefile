
build:	winbuild macbuild linuxbuild
winbuild:
	GOOS=windows GOARCH=amd64 go build -o bin/dss-amd64-win.exe cmd/apiserver/main.go

macbuild:
	GOOS=darwin GOARCH=amd64 go build -o bin/dss-amd64-darwin cmd/apiserver/main.go

linuxbuild:
	GOOS=linux GOARCH=amd64 go build -o bin/dss-amd64-linux cmd/apiserver/main.go

createkafkatopic:
	kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic first_kafka_topic

viewtopicnames:
	kafka-topics.sh --list --zookeeper zookeeper:2181

startapiserver:
	go run cmd/apiserver/main.go

startKafkaconsumer:
	go run cmd/kafkadriver/consumer/main.go

gotest:
	go test ./...
