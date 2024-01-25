LOCATION=australia-southeast1
GCF_NAME=sudokurest2

.PHONY: build_cmd
build_cmd: sudokurest

default: sudokurest

test:
	@echo Running test for all packages
	go test ./...

deploy:
	@echo Deploying Cloud Function using ../gcf-deploy.sh script
	cd ../ && bash gcf-deploy.sh $(GCF_NAME) $(LOCATION)

run_cmd:
	go run cmd/cmd.go

sudokurest: 
	go build -o ./sudokurest cmd/cmd.go
	
test_gcf:
	@echo Testing the GCF with curl
	bash ../testgcf.sh $(GCF_NAME) $(LOCATION) ../sample.json

test_local:
	@echo Testing a local server with curl
	@echo ensure that you have started the server locally first!
	bash ../testlocal.sh ../sample.json