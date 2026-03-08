gen:
	protoc -I . \
		-I ./proto \
		--go_out=. --go_opt=module=github.com/goginenibhavani2000/task-management-service \
		--go-grpc_out=. --go-grpc_opt=module=github.com/goginenibhavani2000/task-management-service \
		--grpc-gateway_out=. --grpc-gateway_opt=module=github.com/goginenibhavani2000/task-management-service \
		proto/task.proto