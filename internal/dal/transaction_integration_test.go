// +build integration  

package dal

import (
	"testing"

	"github.com/powerman/check"
)
func TestGetTransactionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		getted, err := testRepo.GetTransaction(testTransaction1.ID, isolatedEntityID)
		t.Nil(err)
		t.DeepEqual(getted, testTransaction1)

	t.Nil(testRepo.truncate())
}
func TestAddTransactionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestDeleteTransactionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		t.Nil(testRepo.DeleteTransaction(testTransaction1.ID, profID1, isolatedEntityID))

	t.Nil(testRepo.truncate())
}
func TestEditTransactionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		t.Nil(testRepo.EditTransaction(testTransaction1.ID, isolatedEntityID, testTransaction2))

	t.Nil(testRepo.truncate())
}
func TestListTransactionSmoke(tt *testing.T) {
	t := check.T(tt)

	t.Nil(testRepo.AddTestData(profID1, isolatedEntityID))
		list, _, _,  err := testRepo.ListTransaction(isolatedEntityID, listParams)
		t.Nil(err)
		t.DeepEqual(list, testTransactions)

	t.Nil(testRepo.truncate())
}