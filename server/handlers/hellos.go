package handlers

import (
	"context"
	meetupreplacementprojectv1 "github.com/forgeutah/meetup-replacement-project/protos/gen/go/meetupreplacementproject/v1"
	"github.com/forgeutah/meetup-replacement-project/server/storage"
)

func (a ApiServer) UpsertHellos(ctx context.Context, request *meetupreplacementprojectv1.UpsertHellosRequest) (*meetupreplacementprojectv1.Hellos, error) {
	upsertedHellos, err := storage.Storage.UpsertHellos(ctx, request)
	return &meetupreplacementprojectv1.Hellos{Hellos: upsertedHellos}, err
}

func (a ApiServer) DeleteHellos(ctx context.Context, request *meetupreplacementprojectv1.DeleteRequest) (*meetupreplacementprojectv1.DeleteResponse, error) {
	return &meetupreplacementprojectv1.DeleteResponse{}, storage.Storage.DeleteHellos(ctx, request.Ids)
}

func (a ApiServer) ListHellos(ctx context.Context, request *meetupreplacementprojectv1.ListRequest) (*meetupreplacementprojectv1.Hellos, error) {
	if request.Limit == 0 {
		request.Limit = 100
	}
	hellos, err := storage.Storage.ListHellos(ctx, request)
	return &meetupreplacementprojectv1.Hellos{Hellos: hellos}, err
}

func (a ApiServer) GetHellos(ctx context.Context, request *meetupreplacementprojectv1.GetRequest) (*meetupreplacementprojectv1.Hellos, error) {
	hellos, err := storage.Storage.GetHellos(ctx, request)
	return &meetupreplacementprojectv1.Hellos{Hellos: hellos}, err
}
