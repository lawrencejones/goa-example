.PHONY: api/gen

api/gen:
	goa gen github.com/lawrencejones/goa-example/api/design -o api
