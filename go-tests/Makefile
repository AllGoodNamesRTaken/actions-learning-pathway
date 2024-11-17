test:
	go mod download
	go mod tidy
	go test -cover -v

build: 
	docker build -t joserracloud/devops:my-go-app .
push: 
	docker push joserracloud/devops:my-go-app