test:
	go test -v

cover:
	rm -rf *.coverprofile
	go test -coverprofile=favicon.coverprofile
	gover
	go tool cover -html=favicon.coverprofile