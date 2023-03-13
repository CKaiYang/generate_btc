### 环境

* go version go1.20.1 darwin/arm64

### 配置

btc.env 默认为 main 可选 main test 切换不同的网络
server.active 可指定子配置文件 配置文件必须是 conf/app-{}.yaml 格式 当配置文件存在时 会自动覆盖父配置文件 注意是全覆盖
server.port 指定启动端口

### seed

本项目 seed 采用当前主流方法 助记词生成