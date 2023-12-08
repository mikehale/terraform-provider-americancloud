objects = main americancloud.gen.go
all: $(objects)

default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

americancloud.gen.go: americancloud.api.json
	oapi-codegen -package=main -generate=client,types -o ./americancloud.gen.go americancloud.api.json

main: americancloud.gen.go main.go
	go build ./main.go ./americancloud.gen.go

clean:
	rm -f $(objects)
