
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

.PHONY:user
user:
	goctl api plugin -p goctl-ap="swagger -f swagger.json" --api api/platform/api.api --dir ../../app/user/api/doc
	docker run --rm --name="doc-user-service" -p 7888:8080 -e SWAGGER_JSON_URL=http://localhost:8888/doc/swagger swaggerapi/swagger-ui
