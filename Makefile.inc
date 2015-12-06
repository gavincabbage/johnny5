# johnny5 inc

TARGET_DIR=dist
TARGET=dist.tgz

GO_ENV=env GOOS=linux GOARCH=arm GOARM=7
GO_FLAGS=-o
GO=$(GO_ENV) go build $(GO_FLAGS)

TARZIP=tar czvf
