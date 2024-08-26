// +build integration  

package dal

import (
	"testing"

	"github.com/powerman/check"
)
func TestGetExampleSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		getted, err := testRepo.GetExample(testExample1.ID, isolatedEntityID)
		t.Nil(err)
		t.DeepEqual(getted, testExample1)

	t.Nil(testRepo.truncate())
}
func TestAddExampleSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestDeleteExampleSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		t.Nil(testRepo.DeleteExample(testExample1.ID, profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestEditExampleSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		t.Nil(testRepo.EditExample(testExample1.ID, isolatedEntityID, testExample2))

	t.Nil(testRepo.truncate())
}
func TestListExampleSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		list, _, _,  err := testRepo.ListExample(isolatedEntityID, listParams)
		t.Nil(err)
		t.DeepEqual(list, testExamples)

	t.Nil(testRepo.truncate())
}