package api

import (
	"fmt"
	"net/http"
	"testing"
		"time"

	"demo/internal/api/restapi/models"
	"demo/internal/app"
	"demo/internal/def"
	extauthapi "github.com/mtgroupit/mt-mock-extauthapi"
		"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
	"github.com/powerman/gotest/testinit"
	"github.com/google/uuid"
	"github.com/phayes/freeport"
)

var (
	isolatedEntityID = uuid.New().String()
	profileID           = uuid.New().String()
	sess             = extauthapi.SessionCookieName + "=sess"
	profile          = &extauthapi.Profile{
		ID:    extauthapi.MustParseID(profileID),
		Authn: true,
		IsolatedEntityID: extauthapi.MustParseID(isolatedEntityID),
	}
	testDeployID1  = uuid.New().String()
	testDeploy1 = &models.Deploy{
		Config: "accusantium",
		CpuUnits: 96.34,
		CreatedAt: toDateTime("1970-08-25T18:25:19.838Z"),
		ID: testDeployID1,
		MemoryMb: -655446824,
		Name: "dolorem",
		Price: 69.84,
		State: "reiciendis",
		StorageMb: 1182685215,
		User: testUser1,
	}
		testDeployID2  = uuid.New().String()
		testDeploy2 = &models.Deploy{
			Config: "distinctio",
			CpuUnits: 575.82,
			CreatedAt: toDateTime("1978-04-09T19:00:53.901Z"),
			ID: testDeployID1,
			MemoryMb: 2088246653,
			Name: "et",
			Price: 538.88,
			State: "rerum",
			StorageMb: 1581178902,
			User: testUser2,
		}
		testDeploys = []*models.Deploy{testDeploy1, testDeploy2}
	testAddDeploy1 = &models.DeployAdd{
		Config: "voluptatem",
		CpuUnits: 396.43,
		MemoryMb: -942493560,
		Name: "non",
		Price: 958.82,
		State: "aspernatur",
		StorageMb: -1297255494,
		User: testUser1.ID,
	}
	testAddDeploy2 = &models.DeployAdd{
		Config: "consequuntur",
		CpuUnits: 296.93,
		MemoryMb: -62650477,
		Name: "quos",
		Price: 735.66,
		State: "eos",
		StorageMb: -257966444,
		User: testUser2.ID,
	}
	testExampleID1  = uuid.New().String()
	testExample1 = &models.Example{
		CpuUnits: 573.01,
		ID: testExampleID1,
		MemoryMb: 433666317,
		Name: "adipisci",
		Price: 362.58,
		StorageMb: -145194971,
		User: testUser1,
	}
		testExampleID2  = uuid.New().String()
		testExample2 = &models.Example{
			CpuUnits: 632.20,
			ID: testExampleID1,
			MemoryMb: -1730258525,
			Name: "dolorum",
			Price: 758.65,
			StorageMb: -1454335385,
			User: testUser2,
		}
		testExamples = []*models.Example{testExample1, testExample2}
	testAddExample1 = &models.ExampleAdd{
		CpuUnits: 495.78,
		MemoryMb: 346207303,
		Name: "aperiam",
		Price: 689.35,
		StorageMb: 509440509,
		User: testUser1.ID,
	}
	testAddExample2 = &models.ExampleAdd{
		CpuUnits: 185.96,
		MemoryMb: -751079460,
		Name: "aut",
		Price: 98.93,
		StorageMb: -1450983758,
		User: testUser2.ID,
	}
	testTransactionID1  = uuid.New().String()
	testTransaction1 = &models.Transaction{
		Amount: 686.55,
		CreatedAt: toDateTime("1930-11-17T23:56:28.214Z"),
		ID: testTransactionID1,
		Info: "minus",
		Type: "sit",
		User: testUser1,
	}
		testTransactionID2  = uuid.New().String()
		testTransaction2 = &models.Transaction{
			Amount: 118.92,
			CreatedAt: toDateTime("1933-04-25T22:19:02.159Z"),
			ID: testTransactionID1,
			Info: "nobis",
			Type: "est",
			User: testUser2,
		}
		testTransactions = []*models.Transaction{testTransaction1, testTransaction2}
	testAddTransaction1 = &models.TransactionAdd{
		Amount: 38.69,
		Info: "doloremque",
		Type: "deserunt",
		User: testUser1.ID,
	}
	testAddTransaction2 = &models.TransactionAdd{
		Amount: 574.57,
		Info: "id",
		Type: "ea",
		User: testUser2.ID,
	}
	testUserID1  = uuid.New().String()
	testUser1 = &models.User{
		Balance: 569.34,
		ID: testUserID1,
		Name: "distinctio",
		Role: "dolorum",
	}
		testUserID2  = uuid.New().String()
		testUser2 = &models.User{
			Balance: 464.16,
			ID: testUserID1,
			Name: "fugit",
			Role: "quasi",
		}
		testUsers = []*models.User{testUser1, testUser2}
	testAddUser1 = &models.UserAdd{
		Balance: 501.28,
		Name: "deserunt",
		Role: "eum",
	}
	testAddUser2 = &models.UserAdd{
		Balance: 580.63,
		Name: "magnam",
		Role: "temporibus",
	}


	offset int64 = 0
	limit  int64 = 5
	
	testList = &models.ListParams{
		Offset: &offset,
		Limit:  limit,
	}

)

func TestMain(m *testing.M) { testinit.Main(m) }

func testNewServer(t *check.C) (string, func(), *app.MockApp, *MockAuthSvc) {
	t.Helper()
	ctrl := gomock.NewController(t)

	mockApp := app.NewMockApp(ctrl)
	mockExtAuthSvc := NewMockAuthSvc(ctrl)

	port , err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	
	server, err := NewServer(mockApp, mockExtAuthSvc, Config{
		Host:     "localhost",
		Port:     port,
		BasePath: def.APIBasePath,
	})
	t.Nil(err, "NewServer")
	t.Nil(server.Listen(), "server.Listen")
	errc := make(chan error, 1)
	go func() { errc <- server.Serve() }()

	shutdown := func() {
		t.Helper()
		t.Nil(server.Shutdown(), "server.Shutdown")
		t.Nil(<-errc, "server.Serve")
		ctrl.Finish()
	}

	url := fmt.Sprintf("localhost:%d", server.Port)

	return url, shutdown, mockApp, mockExtAuthSvc
}

type matchCookie string // Implements gomock.Matcher.

func (m matchCookie) String() string { return string(m) }
func (m matchCookie) Matches(x interface{}) bool {
	for _, c := range (&http.Request{Header: map[string][]string{"Cookie": {x.(string)}}}).Cookies() {
		if c.String() == string(m) {
			return true
		}
	}
	return false
}
	func fromDateTime(dt strfmt.DateTime) time.Time {
		return time.Time(dt)
	}

	func toDateTime(date interface{}) (*strfmt.DateTime) {
		if date == nil {
			return nil
		}
		var dt strfmt.DateTime
		dt.Scan(date)
		return &dt
	}
	
	func fromDate(d strfmt.Date) time.Time {
		return time.Time(d)
	}

	func toDate(date interface{}) (*strfmt.Date) {
		if date == nil {
			return nil
		}
		var dt strfmt.Date
		dt.Scan(date)
		return &dt
	}

