PACKAGE=github.com/rebuy-de/terraform-provider-graylog

include golang.mk

testacc: format
	TF_ACC=1 go test $(GOPKGS) -v $(TESTARGS) -timeout 120m
