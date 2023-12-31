# app-center

服务中心

> Docker 环境下的 go-zero 开发示例

## 使用方法

- 执行命令 `docker dev create -o --name zero-vue-demo-app-center git@github.com:zero-vue-demo/app-center.git` 构建开发环境
- 在根目录下执行命令 `make init` 初始化项目
- 在根目录下执行 `make doc` 启动 api 文档容器
- 在根目录下执行 `make dev` 初始化开发环境
- 运行基础权限服务
  - 利用 vscode 插件创建数据库、导入数据结构 `dev/service/jwtx.sql`
  - 编辑配置文件 `dev/service/jwtx-rpc.yaml`
  - 在根目录下执行 `make jwtx` 后台运行 jwtx rpc 服务

## 开发流程

> 以 admin 服务为例\
服务定义目录为 `define/admin`

### 服务创建

- 初始化“*服务定义目录*”
    - 在根目录下执行 `make new` 生成服务初始化文件
- 数据结构定义
  - 在 `define/admin/db/mysql` 目录下分别定义各表数据结构
  - 在 `define/admin` 目录下执行 `make sql` 合并 sql 文件
  - 利用 vscode 插件创建数据库、导入数据结构
  - 在 `define/admin` 目录下执行 `make dao` 生成 dao 文件
- 服务 api 定义
  - 在 `define/admin/api` 目录下定义 api 结构
  - 在 `define/admin` 目录下执行 `make doc` 生成 api 文档
  - 在 `define/admin` 目录下执行 `make api` 生成 api 基础代码
  - 在 `define/admin` 目录下执行 `make ts` 生成 api TypeScript 脚本
  - 文件目录说明
    - `define/admin/api/app/app.api` 为生成 api 基础代码的入口文件
    - `define/admin/api/app/platform/*.api` 为生成 api TypeScript 脚本的入口文件
  - 文件命名规范
    - `平台名称_权限分组.api`

### 服务开发

> 命令行进入 `define/admin` 目录，执行 `make` 命令查看支持的快捷指令

- 编写业务逻辑
  - 配置文件路径 `app/admin/api/etc`
  - 逻辑文件路径 `app/admin/api/internal/logic`
- 运行 api 服务
  - 在 `define/admin` 目录下执行 `make api-run` 运行 api 服务
- 项目相关的通用代码
  - 目录 `common`
  - 仅供内部调用
- 业务无关的通用代码
  - 目录 `work` 为项目工作区
  - *工作区内的代码是跨项目使用的，切记不要混入业务逻辑！*
    - 通用工具封装
      - 目录 `work/go-tools`
    - 通用权限服务
      - 目录 `work/zero-auth`
