# Golang-GRPC-Work

set-up grpc
```
steps for intstallling grpc:
  1.brew install protobuf
  2.check: protoc --version
  3. Protocol buffer compiler, protoc, version 3.
  4. For installation instructions, see Protocol Buffer Compiler Installation.
  5. $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  6. $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
  7. run bwlow cmd
     a. export GOPATH="$HOME/go"
     b. PATH="$GOPATH/bin:$PATH"
     c. export PATH="$PATH:$(go env GOPATH)/bin"
  8. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
Incase error come up in 

run blog

```
1. make blog
2. ./bin/blog/server
3. ./bin/blog/client
```

run grpcTutorial

```
1. make greet
2. ./bin/blog/server
3. ./bin/blog/client
```

run grpcTutorialSumApi

```
1. make sumapi
2. ./bin/blog/server
3. ./bin/blog/client
```
