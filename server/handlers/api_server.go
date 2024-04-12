package handlers

import (
	"context"
	"errors"
	meetupreplacementprojectv1 "github.com/forgeutah/meetup-replacement-project/protos/gen/go/meetupreplacementproject/v1"
	"github.com/forgeutah/meetup-replacement-project/server/storage"
)

type ApiServer struct{}

func (a ApiServer) Healthy(ctx context.Context, empty *meetupreplacementprojectv1.Empty) (*meetupreplacementprojectv1.Empty, error) {
	return &meetupreplacementprojectv1.Empty{}, nil
}

func (a ApiServer) Ready(ctx context.Context, empty *meetupreplacementprojectv1.Empty) (*meetupreplacementprojectv1.Empty, error) {
	if storage.Storage.Ready() {
		return &meetupreplacementprojectv1.Empty{}, nil
	}
	return nil, errors.New("")
}
