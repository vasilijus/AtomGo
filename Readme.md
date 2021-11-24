# Install
Required plugins for atom
go ide
plarformio terminal
emmet

in terminal run
```
go run ./main.go
go build
```

## Adding dependency 
withing project dir

Example usage:
        'go mod init example.com/m/v2' to initialize a v2 module

go: creating new go.mod: module golocal.local
go: to add module requirements and sums:
        go mod tidy

PS C:\Users\SErgio\Documents\Programming\Go\AtomGo> go mod tidy 
go: finding module for package github.com/go-sql-driver/mysql
go: found github.com/go-sql-driver/mysql in github.com/go-sql-driver/mysql v1.6.

This will add *go.mod / go.sum* files