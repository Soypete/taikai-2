package test

import (
	meetupreplacementprojectv1 "github.com/forgeutah/meetup-replacement-project/protos/gen/go/meetupreplacementproject/v1"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/testing/protocmp"
)

func (s *MeetupreplacementprojectSuite) TestListHellos() {
	Hellos, err := ListHellos(1000, 0, "")
	require.NoError(s.T(), err)
	Hellos = lo.Filter(Hellos, func(item *meetupreplacementprojectv1.Hello, index int) bool {
		return LoadedTestData.HelloIds.Contains(lo.FromPtr(item.Id))
	})
	assertProtoEqualitySortById(s.T(), LoadedTestData.Hellos, Hellos)
}

func (s *MeetupreplacementprojectSuite) TestGetHellosById() {
	ids := []string{}
	for _, id := range LoadedTestData.HelloIds.Values() {
		ids = append(ids, id.(string))
	}
	Hellos, err := GetHellosById(ids)
	require.NoError(s.T(), err)
	assertProtoEqualitySortById(s.T(), LoadedTestData.Hellos, Hellos)
}

func (s *MeetupreplacementprojectSuite) TestUpdateHellos() {
	hellos := CreateRandomNumHellos(s.T())
	randomizedHellos := randomizeHellos(s.T(), hellos)
	updatedHellos, err := UpsertHellos(randomizedHellos)
	require.NoError(s.T(), err)

	assertProtoEqualitySortById(s.T(), randomizedHellos, updatedHellos,
		protocmp.IgnoreFields(&meetupreplacementprojectv1.Hello{}, "id", "updated_at", "hello_type", "person_name"),
	)
}

func (s *MeetupreplacementprojectSuite) TestDeleteHellos() {
	Hellos := CreateRandomNumHellos(s.T())
	ids := lo.Map(Hellos, func(item *meetupreplacementprojectv1.Hello, index int) string {
		return lo.FromPtr(item.Id)
	})
	err := DeleteHellos(ids)
	require.NoError(s.T(), err)
	Hellos, err = GetHellosById(ids)
	require.NoError(s.T(), err)
	require.Len(s.T(), Hellos, 0)
}
