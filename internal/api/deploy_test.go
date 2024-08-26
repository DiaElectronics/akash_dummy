package api

import (
	"testing"

	"demo/internal/api/restapi/client"
	deploy "demo/internal/api/restapi/client/deploy"
	
	cl "github.com/go-openapi/runtime/client"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)


func TestGetDeploy(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := deploy.NewGetDeployParams()
			params.Body.ID = testDeployID1
		mockApp.EXPECT().GetDeploy(gomock.Any(), gomock.Any()).Return(appDeploy(testDeploy1, true), nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			geted, err := c.Deploy.GetDeploy(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(geted.Payload, apiDeploy(appDeploy(testDeploy1, true)))
		})
}

func TestAddDeploy(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := deploy.NewAddDeployParams()	
		params.Body = testAddDeploy1
		mockApp.EXPECT().AddDeploy(gomock.Any(), gomock.Any()).Return(appDeploy(testDeploy1, true), nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			geted, err := c.Deploy.AddDeploy(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(geted.Payload, apiDeploy(appDeploy(testDeploy1, true)))
		})
}

func TestDeleteDeploy(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := deploy.NewDeleteDeployParams()
			params.Body.ID = testDeploy1.ID
		mockApp.EXPECT().DeleteDeploy(gomock.Any(), gomock.Any()).Return(nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			_, err := c.Deploy.DeleteDeploy(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
		})
}

func TestEditDeploy(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := deploy.NewEditDeployParams()
			params.Body.Data = testAddDeploy1
				params.Body.ID = testDeploy1.ID
		mockApp.EXPECT().EditDeploy(gomock.Any(), gomock.Any(), gomock.Any()).Return(appDeploy(testDeploy1, true), nil)
		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			geted, err := c.Deploy.EditDeploy(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
				t.DeepEqual(geted.Payload, apiDeploy(appDeploy(testDeploy1, true)))
		})
}

func TestListDeploy(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := deploy.NewListDeployParams()
		params.Body = testList
		mockApp.EXPECT().ListDeploy(gomock.Any(), gomock.Any()).Return(appDeploys(testDeploys, true), 2, []string{}, nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			list, err := c.Deploy.ListDeploy(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(list.Payload.Items, apiDeploys(appDeploys(testDeploys, true)))
		})
}
