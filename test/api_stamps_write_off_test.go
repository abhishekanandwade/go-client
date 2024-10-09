/*
Treasury Management API

Testing StampsWriteOffAPIService

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

func Test_openapi_StampsWriteOffAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StampsWriteOffAPIService OfficesOfficeIdStampsWriteOffGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId string

		resp, httpRes, err := apiClient.StampsWriteOffAPI.OfficesOfficeIdStampsWriteOffGet(context.Background(), officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StampsWriteOffAPIService OfficesOfficeIdStampsWriteOffPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId int32

		resp, httpRes, err := apiClient.StampsWriteOffAPI.OfficesOfficeIdStampsWriteOffPost(context.Background(), officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StampsWriteOffAPIService StampsWriteoffApprovePut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StampsWriteOffAPI.StampsWriteoffApprovePut(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}