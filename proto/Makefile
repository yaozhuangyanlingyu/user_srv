.PHONY: proto

# makefile 绝对路径
MAKEFILE_PATH = $(abspath $(lastword $(MAKEFILE_LIST)))
# 当前目录
CURRENT_DIR = $(dir $(MAKEFILE_PATH))
# 父级目录
PARENT_DIR = $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../)

ENV_SUFFIX =
# 可执行文件目录
BIN_DIR = $(CURRENT_DIR)/local_bin

# protoc可执行文件
PROTOC = $(BIN_DIR)/protoc$(ENV_SUFFIX)

# grpc_php_plugin可执行文件
GRPC_PHP_PLUGIN = $(BIN_DIR)/grpc_php_plugin

# protoc-gen-micro可执行文件
MICRO_PLUGIN = $(BIN_DIR)/protoc-gen-micro

# protoc-gen-go可执行文件
GO_PLUGIN = $(BIN_DIR)/protoc-gen-go

SYSTEM = $(shell uname)

GO_TEMPLATE = $(PROTOC)  --proto_path=$(CURRENT_DIR) --micro_out=$(CURRENT_DIR)/go --go_out=$(CURRENT_DIR)/go --plugin=$(GO_PLUGIN) --plugin=$(MICRO_PLUGIN)

define generate_proto_fn_go
	$(call download_file)
	$(GO_TEMPLATE) $(CURRENT_DIR)/$(1)/*.proto
	@echo "generate go proto OK.                                                 [  OK  ]"
endef

define download_file
	$(shell mkdir -p $(CURRENT_DIR)/php)
	$(shell mkdir -p $(CURRENT_DIR)/go)

	$(shell mkdir -p $(BIN_DIR))
	$(shell ! command -v wget >/dev/null 2>&1 && echo 没有呢)
	$(shell [ ! -e  $(GRPC_PHP_PLUGIN) ] && wget -O $(GRPC_PHP_PLUGIN) https://files.aplum.com/grpc_tools/grpc_php_plugin$(ENV_SUFFIX) && chmod +x $(GRPC_PHP_PLUGIN))
	$(shell [ ! -e  $(PROTOC) ] && wget -O $(PROTOC) https://files.aplum.com/grpc_tools/protoc$(ENV_SUFFIX) && chmod +x $(PROTOC))
	$(shell [ ! -e  $(MICRO_PLUGIN) ] && wget -O $(MICRO_PLUGIN) https://files.aplum.com/grpc_tools/protoc-gen-micro$(ENV_SUFFIX) && chmod +x $(MICRO_PLUGIN))
	$(shell [ ! -e  $(GO_PLUGIN) ] && wget -O $(GO_PLUGIN) https://files.aplum.com/grpc_tools/protoc-gen-go$(ENV_SUFFIX) && chmod +x $(GO_PLUGIN))
endef

proto_user:
	$(call generate_proto_fn_go,user)

build:
	docker build . -t user_srv:latest

