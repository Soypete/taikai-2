package test

import (
	"context"
	"time"

	meetupreplacementprojectv1 "github.com/forgeutah/meetup-replacement-project/protos/gen/go/meetupreplacementproject/v1"
	"google.golang.org/grpc"
)

var ApiClient = initializeApiClient(5 * time.Second)

func initializeApiClient(timeout time.Duration) meetupreplacementprojectv1.ApiClient {
	connectTo := "127.0.0.1:6000"
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx, connectTo, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return meetupreplacementprojectv1.NewApiClient(conn)
}

func UpsertHellos(hellos []*meetupreplacementprojectv1.Hello) ([]*meetupreplacementprojectv1.Hello, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.UpsertHellos(ctx, &meetupreplacementprojectv1.UpsertHellosRequest{Hellos: hellos})
	if err != nil {
		return nil, err
	} else if response == nil {
		return nil, nil
	} else {
		return response.Hellos, err
	}
}

func DeleteHellos(ids []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := ApiClient.DeleteHellos(ctx, &meetupreplacementprojectv1.DeleteRequest{Ids: ids})
	return err
}

func ListHellos(limit, offset int, orderBy string) ([]*meetupreplacementprojectv1.Hello, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.ListHellos(ctx, &meetupreplacementprojectv1.ListRequest{Limit: int32(limit), Offset: int32(offset), OrderBy: orderBy})
	if err != nil {
		return nil, err
	}
	return response.Hellos, err
}

func GetHellosById(ids []string) ([]*meetupreplacementprojectv1.Hello, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.GetHellos(ctx, &meetupreplacementprojectv1.GetRequest{Ids: ids})
	return response.Hellos, err
}
