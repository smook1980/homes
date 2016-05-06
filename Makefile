homes: test
	go install

test:
	go test $(glide novendor)
