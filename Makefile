hello:
	echo "building vegeta binary.."
compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/vegeta-freebsd-386 vegeta.go
	GOOS=linux GOARCH=386 go build -o bin/vegeta-linux-386 vegeta.go
	GOOS=windows GOARCH=386 go build -o bin/vegeta-windows-386 vegeta.go
run:
	go run vegeta.go