# steps

From this dir (cd ~/src/Rpc/helloworld)
- Run protoc -I helloworld/ helloworld.proto --go_out=plugins=grpc:helloworld in this dir.
  - this will generate helloworld.pb.go in ~/helloworld dir.

- Run go run greeter_server/main.go

- In another terminal, run greeter_client/main.go

: The client will see me
