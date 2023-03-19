
to create proto 

**protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  proto/pushnotification.proto**

to start grpc server and notification api

**go run main.go**

to start client and subscribe to notification

**go run client/client.go**

to send notification to server

 localhost:8080/notify

-> verify client console log to see push notification.
