#### 安装 protobuf
grpc使用protobuf作为IDL(interface descriton language)，且要求protobuf 3.0以上，这里我们直接选用当前最新版本 3.8。

选择操作系统对应的版本下载，这里我们直接使用已经编译好的protoc可执行文件（或者下载安装包编译安装）。
```
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.8.0-rc1/protoc-3.8.0-rc-1-linux-x86_64.zip
tar -zxvf protoc-3.8.0-rc-1-linux-x86_64.zip
mv protoc-3.8.0-rc-1-linux-x86_64 /usr/local/protoc
ln -s /usr/local/protoc/bin/protoc /usr/local/bin/protoc
查看 protoc
protoc --version
```

#### 安装 protoc-gen-go
protoc-gen-go是Go的protoc编译插件，protobuf内置了许多高级语言的编译器，但没有Go的。
```
# 运行 protoc -h 命令可以发现内置的只支持以下语言
protoc -h
...
--cpp_out=OUT_DIR           Generate C++ header and source.
--csharp_out=OUT_DIR        Generate C# source file.
--java_out=OUT_DIR          Generate Java source file.
--js_out=OUT_DIR            Generate JavaScript source.
--objc_out=OUT_DIR          Generate Objective C header and source.
--php_out=OUT_DIR           Generate PHP source file.
--python_out=OUT_DIR        Generate Python source file.
--ruby_out=OUT_DIR          Generate Ruby source file.
...
```

所以我们使用protoc编译生成Go版的grpc时，需要先安装此插件。

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

#### 安装 grpc-go 库
```
cd $GOPATH/src/
go install google.golang.org/grpc
```

#### 生成pb文件
```
cd user 
protoc -I. --go_out=plugins=grpc:. user.proto
```

[grpc](https://segmentfault.com/a/1190000019216566)