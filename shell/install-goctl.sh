go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go install github.com/zeromicro/go-zero/tools/goctl@v1.5.2
goctl env check -i -f --verbose