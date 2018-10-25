BASE = $(GOPATH)/src/github.com/girikuncoro/godc

.PHONY: cli
cli: ; $(info building cli...) @ ## Build cli
	./hack/build-bin.sh godc

.PHONY: exposure
exposure: ; $(info building godc-exposure bin...) @ ## Build cli
	./hack/build-bin.sh godc-exposure

.PHONY: generate
generate: ; $(info generating swagger server...) @ ## Generate swagger server
	cd $(BASE)/gen && swagger generate server \
		-A godc -f $(BASE)/swagger/swagger.yaml -c godc --skip-validation

.PHONY: generate-kopral
generate-kopral: ; $(info generating swagger server for kopral...) @ ## Kopral server
	cd $(BASE) && swagger generate server \
		-A Kopral -t ./pkg/kopral/gen -f $(BASE)/swagger/kopral.yaml
