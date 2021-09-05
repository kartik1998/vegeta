hello:
	echo "building vegeta binary.."
build:
	go build -o vegeta vegeta.go
run:
	go run vegeta.go