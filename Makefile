.PHONY: api/gen openapi-generator.jar clients/typescript docs

# Generate backend code and API specification from Goa design.
api/gen:
	goa gen github.com/lawrencejones/goa-example/api/design -o api

# Download openapi-generator, which we use to generate clients.
openapi-generator-cli.jar:
	curl https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.2.0/openapi-generator-cli-5.2.0.jar \
		--output $@

# Generate a TypeScript client from the OpenAPI specification.
clients/typescript:
	rm -rfv $@
	java -jar openapi-generator-cli.jar \
		generate \
			--generator-name typescript-fetch \
			--input-spec api/gen/http/openapi.json \
			--skip-validate-spec \
			--additional-properties npmName=goa-example,typescriptThreePlus=true,modelPropertyNaming=original \
			--output $@

# Expose documentation for the HTTP API using go-swagger.
docs:
	docker run --platform=linux/amd64 -p 4000:4000 --rm -v "$$(pwd)/api:/api" -it quay.io/goswagger/swagger:v0.28.0 \
		serve --no-open --port=4000 --host=0.0.0.0 /api/gen/http/openapi.json
