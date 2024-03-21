
.PHONY: build

SERVICE := GLOWER
CUR_PWD := C:/Users/"Zhuiri Xiao"/Desktop/glower


SERVER_DEMO_PATH := $(CUR_PWD)/examples/zinx_server
CLIENT_DEMO_PATH := $(CUR_PWD)/examples/zinx_client
SERVER_DEMO_BIN := $(SERVER_DEMO_PATH)/server.exe
CLIENT_DEMO_BIN := $(CLIENT_DEMO_PATH)/client.exe



export GO111MODULE=on

LD_FLAGS='-X "$(SERVICE)/version.TAG=$(TAG)" -X "$(SERVICE)/version.VERSION=$(VERSION)" -X "$(SERVICE)/version.AUTHOR=$(AUTHOR)" -X "$(SERVICE)/version.BUILD_INFO=$(BUILD_INFO)" -X "$(SERVICE)/version.BUILD_DATE=$(BUILD_DATE)"'

TEST_FILES := znet/acceptdelay_test.go znet/acceptdelay.go

default: build

build:
	go build  -ldflags $(LD_FLAGS) -gcflags "-N"  -o $(SERVER_DEMO_BIN) $(SERVER_DEMO_PATH)/main.go
	go build  -ldflags $(LD_FLAGS) -gcflags "-N"  -o $(CLIENT_DEMO_BIN) $(CLIENT_DEMO_PATH)/main.go

test:
	go test -v -cover $(TEST_FILES)
clean:
	rm $(SERVER_DEMO_BIN)
	rm $(CLIENT_DEMO_BIN)
