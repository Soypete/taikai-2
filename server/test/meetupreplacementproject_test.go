package test

import (
	meetupreplacementprojectv1 "github.com/forgeutah/meetup-replacement-project/protos/gen/go/meetupreplacementproject/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
	"time"
)

type MeetupreplacementprojectSuite struct {
	suite.Suite
}

func TestMeetupreplacementprojectSuite(t *testing.T) {
	suite.Run(t, new(MeetupreplacementprojectSuite))
}

func (s *MeetupreplacementprojectSuite) SetupSuite() {
	initializeApiClient(5 * time.Second)
	loadTestData(s.T())
}

func assertProtoEqualitySortById(t *testing.T, expected, actual interface{}, opts ...cmp.Option) {
	theOpts := []cmp.Option{
		cmpopts.SortSlices(func(x, y *meetupreplacementprojectv1.Hello) bool {
			return *x.Id < *y.Id
		}),
		protocmp.Transform(),
		protocmp.SortRepeated(func(x, y *meetupreplacementprojectv1.Hello) bool {
			return *x.Id < *y.Id
		}),
	}
	theOpts = append(theOpts, opts...)
	diff := cmp.Diff(expected, actual, theOpts...)
	require.Empty(t, diff)
}
