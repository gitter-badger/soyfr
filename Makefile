ifndef version
$(error version must be defined. make version=someVersion)
endif

scheme:
	java -jar bin/mongeezer-1.0-SNAPSHOT-jar-with-dependencies.jar -d soyfr_development -h 127.0.0.1 -p 27017 -l changesets/bootstrap.xml

staging:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GOPATH=`godep path`:$(GOPATH) go build -o build/soyfr main.go
	npm install
	./node_modules/.bin/bower install
	./node_modules/.bin/grunt
	docker build -t soyfr .
	@-docker tag -f soyfr soyfr:$(version) || true 
	@-docker stop soyfr || true
	@-docker rm soyfr || true
	@-docker tag -f soyfr soyfr:latest || true 
