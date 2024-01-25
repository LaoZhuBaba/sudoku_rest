LOCATION=australia-southeast1
GCF_NAME=sudokurest

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
	@echo 4x4:
	bash ../testgcf.sh $(GCF_NAME) $(LOCATION) ../sample2.json
	@echo ""
	@echo 9x9:
	bash ../testgcf.sh $(GCF_NAME) $(LOCATION) ../sample3.json
	@echo ""
	@echo 16x16:
	bash ../testgcf.sh $(GCF_NAME) $(LOCATION) ../sample4.json

test_local:
	@echo Testing a local server with curl
	@echo ensure that you have started the server locally first!
	@echo ""
	@echo 4x4:
	bash ../testlocal.sh ../sample2.json
	@echo ""
	@echo 9x9:
	bash ../testlocal.sh ../sample3.json
	@echo ""
	@echo 16x16:
	bash ../testlocal.sh ../sample4.json

lookup_table:
	cd autogen && python3 make_lookup_table.py > lookup_table.go 