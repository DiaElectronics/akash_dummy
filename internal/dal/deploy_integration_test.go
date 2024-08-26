// +build integration  

package dal

import (
	"testing"

	"github.com/powerman/check"
)
func TestGetDeploySmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		getted, err := testRepo.GetDeploy(testDeploy1.ID, isolatedEntityID)
		t.Nil(err)
		t.DeepEqual(getted, testDeploy1)

	t.Nil(testRepo.truncate())
}
func TestAddDeploySmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestDeleteDeploySmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		t.Nil(testRepo.DeleteDeploy(testDeploy1.ID, profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestEditDeploySmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		t.Nil(testRepo.EditDeploy(testDeploy1.ID, isolatedEntityID, testDeploy2))

	t.Nil(testRepo.truncate())
}
func TestListDeploySmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		list, _, _,  err := testRepo.ListDeploy(isolatedEntityID, listParams)
		t.Nil(err)
		t.DeepEqual(list, testDeploys)

	t.Nil(testRepo.truncate())
}