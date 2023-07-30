# 获取参数：容器名称
CONTAINER_NAME=$1
# 判断容器是否存在
if docker container inspect "$CONTAINER_NAME" >/dev/null 2>&1; then
    # 查询容器状态
    CONTAINER_STATUS=$(docker container inspect -f '{{.State.Status}}' "$CONTAINER_NAME")
    if [ "$CONTAINER_STATUS" = "exited" ]; then
        echo "Container $CONTAINER_NAME exist and stopped."
        # 容器存在，未运行
        exit 0
    else
        # 容器存在，已运行
        echo "Container $CONTAINER_NAME is running."
        exit 1
    fi
else
    #容器不存在
    echo "Container $CONTAINER_NAME does not exist."
    docker run --rm --name="zero" -p 7888:8080 -e SWAGGER_JSON_URL=http://localhost:8888/doc/swagger swaggerapi/swagger-ui
    exit 1
fi

