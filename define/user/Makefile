
run:
	@echo ""
	@echo "make tidy      |  更新 Golang 依赖"
	@echo "-----------------------------------------"
	@echo "make ts        |  生成 TypeScript 脚本"
	@echo "-----------------------------------------"
	@echo "make sql       |  合并 sql 文件"
	@echo "make dao       |  生成 dao 文件"
	@echo "-----------------------------------------"
	@echo "make api       |  生成 api 文件"
	@echo "make api-run   |  运行 api 服务"
	@echo "make api-build |  编译 api 服务"
	@echo ""

.PHONY:tidy
tidy:
	cd ../../app && go mod tidy

.PHONY:ts
ts:
	goctl api plugin -plugin goctl-ts ts --api api/platform/client.api --dir ../../frontend/ts/client
	goctl api plugin -plugin goctl-ts ts --api api/platform/manager.api --dir ../../frontend/ts/manager

.PHONY:sql
sql:
	cd db && > mysql.sql && cat mysql/*.sql >> mysql.sql

.PHONY:dao
dao:
	cd ../../app/user/db && go run db.go -f ../api/etc/user.yaml

.PHONY:api
api:
	goctl api go --api api/platform/api.api --dir ../../app/user/api --style goZero
