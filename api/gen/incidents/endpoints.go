// Code generated by goa v3.5.4, DO NOT EDIT.
//
// Incidents endpoints
//
// Command:
// $ goa gen github.com/lawrencejones/goa-example/api/design -o api

package incidents

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "Incidents" service endpoints.
type Endpoints struct {
	Create goa.Endpoint
}

// NewEndpoints wraps the methods of the "Incidents" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Create: NewCreateEndpoint(s),
	}
}

// Use applies the given middleware to all the "Incidents" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "Create" of service "Incidents".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreatePayload)
		return s.Create(ctx, p)
	}
}