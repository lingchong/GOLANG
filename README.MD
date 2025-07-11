golang 标准项目


#### 配置阿里镜像源
* go env -w GO111MODULE=on
* go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/

#### 初始化模块
* go mod init p1 (初始化 p1模块)

* 手动执行导入 github.com/gin-gonic/gin 后. 执行以下命令自动下载并添加依赖
> go mod tidy

* 检查可用更新
> go list -u -m all

* 更新所有依赖
> go get -u ./...
