ifndef version
$(error version must be defined. make version=someVersion)
endif

staging:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/soyfr main.go
	docker build -t soyfr .
	@-docker tag -f soyfr soyfr:$(version) || true 
	@-docker stop soyfr || true
	@-docker rm soyfr || true
	@-docker tag -f soyfr soyfr:latest || true 
