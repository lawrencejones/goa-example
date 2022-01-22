package api

import (
	"context"

	"github.com/lawrencejones/goa-example/api/gen/incidents"

	"github.com/google/uuid"
)

func NewIncidents() incidents.Service {
	return &incidentsService{}
}

type incidentsService struct {
}

func (svc *incidentsService) Create(ctx context.Context, payload *incidents.CreatePayload) (*incidents.Incident, error) {
	inc := incidents.Incident{
		ID:   uuid.NewString(),
		Name: payload.Name,
	}

	return &inc, nil
}
