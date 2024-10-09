/*
Treasury Management API

Testing ToolsDayEndAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ToolsDayEndAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ToolsDayEndAPIService OfficesOfficeIdDayBeginEndStatusGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId int32

		resp, httpRes, err := apiClient.ToolsDayEndAPI.OfficesOfficeIdDayBeginEndStatusGet(context.Background(), officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolsDayEndAPIService OfficesOfficeIdDayBeginEndStatusPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId int32

		resp, httpRes, err := apiClient.ToolsDayEndAPI.OfficesOfficeIdDayBeginEndStatusPut(context.Background(), officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
