install:
	go get github.com/onsi/ginkgo/ginkgo github.com/onsi/gomega

run:
	go run main.go

test:
	ginkgo -r -v
