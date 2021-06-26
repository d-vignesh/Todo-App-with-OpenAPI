.PHONY: openapi
openapi-codegen:
	oapi-codegen -generate types -o openapi_types.gen.go -package main todo.yml
	oapi-codegen -generate chi-server -o openapi_server.gen.go -package main todo.yml