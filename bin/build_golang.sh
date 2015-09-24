BINARY_NAME="a.out"
GOOS="linux"
GOARCH="arm"
GOARM="7"

if [ -n "$1" ]; then
    BINARY_NAME="$1"
fi

env GOOS=$GOOS GOARCH=$GOARCH GOARM=$GOARM go build -v -o $BINARY_NAME *.go
