swagger:=quay.io/goswagger/swagger
swagger_file:=https://raw.githubusercontent.com/InterNetX/internetx-swagger-files/master/src/domainrobot.json
project=$(HOME)/Documents/workspaces/internetx-autodns-go

.PHONY: clean
clean:
	rm -f domainrobot.json

domainrobot.json:
	curl -sS -o ./domainrobot.json $(swagger_file)

.PHONY: generate
generate: domainrobot.json
	mkdir -p $(project)
	docker run --rm -it -e GOPATH=$(HOME)/go:/go -v $(HOME):$(HOME) -w $(CURDIR) $(swagger) \
		generate client \
		-f ./domainrobot.json \
		-A internetx-autodns-go \
		--skip-validation \
		--target=$(project)
