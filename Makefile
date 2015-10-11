# export GOROOT=/d/apps/go1.5.1.windows-amd64.msi~/Go
# export PATH=$PATH:$GOROOT/bin
# export GOPATH=/g/golib

clean:
    -rm bcfs
build:
    go build src/cfs.go
