

check_swagger:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_swagger
	@echo Ensure you have the swagger CLI or this command will fail.
	@echo You can install the swagger CLI with: go get -u github.com/go-swagger/go-swagger/cmd/swagger
	@echo ....

	swagger generate spec -o ./swagger.json --scan-models


docker run --rm -it -v ${PWD}:/app -w /app quay.io/goswagger/swagger
docker run --rm -it -v ${PWD}:/app -w /app quay.io/goswagger/swagger generate spec -o ./swagger.json --scan-models

   