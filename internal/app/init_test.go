// Code generated by mtgroup-generator.
package app

import (
		"time"
		"github.com/go-openapi/strfmt"

	"github.com/google/uuid"
		"demo/internal/types"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

var (
	isolatedEntityID = uuid.New().String()
	profileID = uuid.New().String()
	profile          = Profile{
		ID:    profileID,
		Authn: true,
		Authz: Authz{
			User:    true,
			Admin:   true,
			Manager: true,
		},
		IsolatedEntityID: isolatedEntityID,
	}
	profileUnAuth = Profile{
		ID:    profileID,
		Authn: false,
		Authz: Authz{
			User:    false,
			Admin:   false,
			Manager: false,
		},
		IsolatedEntityID: isolatedEntityID,
	}
	testDeploy1 = &Deploy{
		Config: "vitae",
		CpuUnits: types.NewDecimal(960.19),
		CreatedAt: mustParseTime("1974-02-25T22:41:47.590Z"),
		ID: uuid.New().String(),
		MemoryMb: 1376152209,
		Name: "est",
		Price: types.NewDecimal(172.91),
		State: "tempora",
		StorageMb: 1315016055,
		User: testUser1,
	}
		testDeploy2 = &Deploy{
			Config: "magni",
			CpuUnits: types.NewDecimal(459.38),
			CreatedAt: mustParseTime("1990-08-01T01:51:53.264Z"),
			ID: uuid.New().String(),
			MemoryMb: -475944003,
			Name: "quos",
			Price: types.NewDecimal(910.98),
			State: "consequatur",
			StorageMb: -755813181,
			User: testUser2,
		}
		testDeploys = []*Deploy{testDeploy1, testDeploy2}
	testExample1 = &Example{
		CpuUnits: types.NewDecimal(417.35),
		ID: uuid.New().String(),
		MemoryMb: 1166824988,
		Name: "ad",
		Price: types.NewDecimal(658.81),
		StorageMb: -1452982485,
		User: testUser1,
	}
		testExample2 = &Example{
			CpuUnits: types.NewDecimal(84.64),
			ID: uuid.New().String(),
			MemoryMb: -1976806566,
			Name: "at",
			Price: types.NewDecimal(525.84),
			StorageMb: -767506881,
			User: testUser2,
		}
		testExamples = []*Example{testExample1, testExample2}
	testTransaction1 = &Transaction{
		Amount: types.NewDecimal(632.25),
		CreatedAt: mustParseTime("1928-03-29T03:01:38.925Z"),
		ID: uuid.New().String(),
		Info: "quidem",
		Type: "voluptate",
		User: testUser1,
	}
		testTransaction2 = &Transaction{
			Amount: types.NewDecimal(942.19),
			CreatedAt: mustParseTime("1934-08-10T00:47:19.409Z"),
			ID: uuid.New().String(),
			Info: "sunt",
			Type: "id",
			User: testUser2,
		}
		testTransactions = []*Transaction{testTransaction1, testTransaction2}
	testUser1 = &User{
		Balance: types.NewDecimal(665.03),
		ID: uuid.New().String(),
		Name: "qui",
		Role: "soluta",
	}
		testUser2 = &User{
			Balance: types.NewDecimal(985.92),
			ID: uuid.New().String(),
			Name: "ipsum",
			Role: "alias",
		}
		testUsers = []*User{testUser1, testUser2}
	listParams = &ListParams{
		Offset: 0,
		Limit: 5,
	}
)
	func mustParseTime(t string) *time.Time {
		date, err := time.Parse(time.RFC3339, t)
		if err != nil {
			date, _ = time.Parse(strfmt.RFC3339FullDate, t)
		}
		return &date
	}