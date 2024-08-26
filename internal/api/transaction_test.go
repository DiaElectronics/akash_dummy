package api

import (
	"testing"

	"demo/internal/api/restapi/client"
	transaction "demo/internal/api/restapi/client/transaction"
	
	cl "github.com/go-openapi/runtime/client"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)


func TestGetTransaction(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := transaction.NewGetTransactionParams()
			params.Body.ID = testTransactionID1
		mockApp.EXPECT().GetTransaction(gomock.Any(), gomock.Any()).Return(appTransaction(testTransaction1, true), nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			geted, err := c.Transaction.GetTransaction(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(geted.Payload, apiTransaction(appTransaction(testTransaction1, true)))
		})
}

func TestAddTransaction(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := transaction.NewAddTransactionParams()	
		params.Body = testAddTransaction1
		mockApp.EXPECT().AddTransaction(gomock.Any(), gomock.Any()).Return(appTransaction(testTransaction1, true), nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			geted, err := c.Transaction.AddTransaction(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(geted.Payload, apiTransaction(appTransaction(testTransaction1, true)))
		})
}

func TestDeleteTransaction(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := transaction.NewDeleteTransactionParams()
			params.Body.ID = testTransaction1.ID
		mockApp.EXPECT().DeleteTransaction(gomock.Any(), gomock.Any()).Return(nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			_, err := c.Transaction.DeleteTransaction(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
		})
}

func TestEditTransaction(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := transaction.NewEditTransactionParams()
			params.Body.Data = testAddTransaction1
				params.Body.ID = testTransaction1.ID
		mockApp.EXPECT().EditTransaction(gomock.Any(), gomock.Any(), gomock.Any()).Return(appTransaction(testTransaction1, true), nil)
		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			geted, err := c.Transaction.EditTransaction(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
				t.DeepEqual(geted.Payload, apiTransaction(appTransaction(testTransaction1, true)))
		})
}

func TestListTransaction(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := transaction.NewListTransactionParams()
		params.Body = testList
		mockApp.EXPECT().ListTransaction(gomock.Any(), gomock.Any()).Return(appTransactions(testTransactions, true), 2, []string{}, nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			list, err := c.Transaction.ListTransaction(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(list.Payload.Items, apiTransactions(appTransactions(testTransactions, true)))
		})
}
