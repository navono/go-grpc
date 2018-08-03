# go-grpc
使用`golang`演示`grpc`的功能

# development
## 安装go插件
`grpc`依赖`protobuff3`，同时针对`golang`，还有一个`protobuff`的插件。因此需要安装以下模块：
- protobuff
- protoc-gen-go
- grpc

获取源码：
> go get -u github.com/golang/protobuf/protoc-gen-go

安装：
> go install github.com/golang/protobuf/protoc-gen-go

## 编译.proto文件
如果`GOPATH`的bin没有加入到环境变量：
> protoc --plugin=protoc-gen-go=$env:GOPATH\bin\protoc-gen-go.exe -I . helloworld.proto --go_out=.

如果加入了环境变量：
<br>
根目录下运行：
> protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld

helloworld下运行：
> protoc -I helloworld/ helloworld.proto --go_out=plugins=grpc:helloworld


- dep ensure -add google.golang.org/grpc
- dep ensure -add github.com/golang/protobuf/protoc-gen-go