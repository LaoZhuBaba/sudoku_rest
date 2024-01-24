.PHONY: build_cmd
build_cmd: sudokurest

default: sudokurest

test:
	@echo Running test for all packages
	go test ./...

deploy:
	@echo Deploying Cloud Function using ../gcf-deploy.sh script
	bash ../gcf-deploy.sh

run_cmd:
	go run cmd/cmd.go

sudokurest: 
	go build -o ./sudokurest cmd/cmd.go
	
