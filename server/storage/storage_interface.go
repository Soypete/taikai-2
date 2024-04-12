package storage

import (
	"context"
	meetupreplacementprojectv1 "github.com/forgeutah/meetup-replacement-project/protos/gen/go/meetupreplacementproject/v1"
)

var Storage StorageInterface

type StorageInterface interface {
	Initialize() (shutdown func(), err error)
	Ready() bool
	UpsertHellos(ctx context.Context, request *meetupreplacementprojectv1.UpsertHellosRequest) ([]*meetupreplacementprojectv1.Hello, error)
	DeleteHellos(ctx context.Context, ids []string) error
	ListHellos(ctx context.Context, request *meetupreplacementprojectv1.ListRequest) ([]*meetupreplacementprojectv1.Hello, error)
	GetHellos(ctx context.Context, request *meetupreplacementprojectv1.GetRequest) ([]*meetupreplacementprojectv1.Hello, error)
	GetHello(ctx context.Context, id string) (*meetupreplacementprojectv1.Hello, error)
}
