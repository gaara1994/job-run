test:
	go mod tidy && go build -o main && ./main --mode=succeed --seconds=5

build: ## Build the docker image
	docker build -t register/job-run:latest .

push: ## Push the docker image
	docker push register/job-run:latest
