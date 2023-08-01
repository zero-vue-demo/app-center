
run:
	@echo ""
	@echo "-----------------------------------------"
	@echo "make init  | 初始化项目"
	@echo "-----------------------------------------"
	@echo "make dev   | 初始化开发环境"
	@echo "make stop  | 关闭所有后台程序"
	@echo "-----------------------------------------"
	@echo "make jwtx  | 后台运行 jwtx rpc 服务"
	@echo "-----------------------------------------"
	@echo ""

.PHONY:init
init:
	sh shell/install-docker.sh
	sh shell/install-goctl.sh
	sh shell/install-gopls.sh
	sh shell/install-vscode-extension.sh

.PHONY:dev
dev:
	cd work/zero-auth/private/jwtx/rpc && go build jwtx.go && mv jwtx ../../../../../dev/jwtx-rpc
	cp work/zero-auth/public/jwtx/deploy/jwtx-rpc.yaml dev/jwtx-rpc.yaml

.PHONY:stop
stop:
	cd dev && sh stop jwtx-rpc

.PHONY:jwtx
jwtx:
	cd dev && sh start jwtx-rpc -f jwtx-rpc.yaml
