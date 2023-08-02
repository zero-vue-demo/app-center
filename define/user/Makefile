
run:
	@echo ""
	@echo "this is [ user ] service"
	@echo ""
	@echo "make tidy      |  更新 Golang 依赖"
	@echo "-----------------------------------------"
	@echo "make ts        |  生成 TypeScript 脚本"
	@echo "make doc       |  生成 api 文档"
	@echo "-----------------------------------------"
	@echo "make sql       |  合并 sql 文件"
	@echo "make dao       |  生成 dao 文件"
	@echo "-----------------------------------------"
	@echo "make api       |  生成 api 文件"
	@echo "make api-run   |  运行 api 服务"
	@echo ""

.PHONY:tidy
tidy:
	cd ../../app && go mod tidy

.PHONY:ts
ts:
	goctl api plugin -p goctl-ap="ts" --api api/platform/client.api  --dir ../../work/app-center-public/ts/client
	goctl api plugin -p goctl-ap="ts" --api api/platform/manager.api --dir ../../work/app-center-public/ts/manager

.PHONY:doc
doc:
	@goctl api plugin -p goctl-ap="swagger -f swagger.json" --api api/app/app.api --dir ../../app/user/api/doc
	@echo "http://localhost:8888/doc/swagger\n"

.PHONY:sql
sql:
	cd db && > mysql.sql && cat mysql/*.sql >> mysql.sql

.PHONY:dao
dao:
	cd ../../app/user/db && go run db.go -f ../api/etc/user.yaml

.PHONY:api
api:
	goctl api go --api api/app/app.api --dir ../../app/user/api --style goZero \
	--remote https://github.com/5-say/go-zero-template --branch custom

.PHONY:api-run
api-run:
	cd ../../app/user/api && go run user.go -f ./etc/user.yaml
