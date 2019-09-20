## Needed /Tech stacks
    + go
    
## To get started follow this checklist:
    + in dir project run "go run service/service-user/main.go"

## How to :
    + compile file proto : protoc --go_out=. *.proto
    + compile file proto output grpc : protoc --go_out=plugins=grpc:. *.proto
