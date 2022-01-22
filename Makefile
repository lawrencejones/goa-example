.PHONY: api/gen openapi-generator.jar clients/typescript

api/gen:
	goa gen github.com/lawrencejones/goa-example/api/design -o api

openapi-generator-cli.jar:
	curl https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.2.0/openapi-generator-cli-5.2.0.jar \
		--output $@

clients/typescript:
	rm -rfv $@
	java -jar openapi-generator-cli.jar \
		generate \
			--generator-name typescript-fetch \
			--input-spec api/gen/http/openapi.json \
			--skip-validate-spec \
			--additional-properties npmName=goa-example,typescriptThreePlus=true,modelPropertyNaming=original \
			--output $@
