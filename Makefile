BASE = $(GOPATH)/src/github.com/girikuncoro/godc

.PHONY: cli
cli: ; $(info building cli...) @ ## Build cli
	./hack/build-cli.sh

.PHONY: generate
generate: ; $(info generating swagger server...) @ ## Generate swagger server
	cd $(BASE)/gen && swagger generate server \
		-A godc -f $(BASE)/swagger/swagger.yaml -c godc --skip-validation
