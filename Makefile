
run:
	@echo ""
	@echo "-----------------------------------------"
	@echo "make init  | 初始化项目"
	@echo "-----------------------------------------"
	@echo "make user  | 启动 swagger 文档"
	@echo "-----------------------------------------"
	@echo ""

.PHONY:init
init:
	sh shell/install-docker.sh
	sh shell/install-goctl.sh
	sh shell/install-gopls.sh
	sh shell/install-vscode-extension.sh

.PHONY:jwtx-rpc
jwtx-rpc:
	cd work/zero-services/private/jwtx/rpc && go run jwtx.go -f ./etc/jwtx.yaml

.PHONY:admin-rpc
admin-rpc:
	cd app/admin/rpc && go run admin.go -f ./etc/admin-rpc.yaml

.PHONY:user-rpc
user-rpc:
	cd app/user/rpc && go run user.go -f ./etc/user-rpc.yaml
