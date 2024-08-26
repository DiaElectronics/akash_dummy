package api

import (
	"testing"

	"demo/internal/api/restapi/client"
	example "demo/internal/api/restapi/client/example"
	
	cl "github.com/go-openapi/runtime/client"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)


func TestGetExample(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := example.NewGetExampleParams()
			params.Body.ID = testExampleID1
		mockApp.EXPECT().GetExample(gomock.Any(), gomock.Any()).Return(appExample(testExample1, true), nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			geted, err := c.Example.GetExample(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(geted.Payload, apiExample(appExample(testExample1, true)))
		})
}

func TestAddExample(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := example.NewAddExampleParams()	
		params.Body = testAddExample1
		mockApp.EXPECT().AddExample(gomock.Any(), gomock.Any()).Return(appExample(testExample1, true), nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			geted, err := c.Example.AddExample(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(geted.Payload, apiExample(appExample(testExample1, true)))
		})
}

func TestDeleteExample(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := example.NewDeleteExampleParams()
			params.Body.ID = testExample1.ID
		mockApp.EXPECT().DeleteExample(gomock.Any(), gomock.Any()).Return(nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			_, err := c.Example.DeleteExample(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
		})
}

func TestEditExample(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := example.NewEditExampleParams()
			params.Body.Data = testAddExample1
				params.Body.ID = testExample1.ID
		mockApp.EXPECT().EditExample(gomock.Any(), gomock.Any(), gomock.Any()).Return(appExample(testExample1, true), nil)
		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			geted, err := c.Example.EditExample(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
				t.DeepEqual(geted.Payload, apiExample(appExample(testExample1, true)))
		})
}

func TestListExample(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := example.NewListExampleParams()
		params.Body = testList
		mockApp.EXPECT().ListExample(gomock.Any(), gomock.Any()).Return(appExamples(testExamples, true), 2, []string{}, nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			list, err := c.Example.ListExample(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(list.Payload.Items, apiExamples(appExamples(testExamples, true)))
		})
}
