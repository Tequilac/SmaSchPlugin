BIN_DIR=output/bin

mkdir -p ${BIN_DIR}
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o=${BIN_DIR}/ssp-scheduler ./scheduler
docker buildx build --platform linux/arm64 -t tequilac/ssp-scheduler -f docker/Dockerfile --push .
