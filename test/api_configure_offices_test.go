/*
Treasury Management API

Testing ConfigureOfficesAPIService

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

func Test_openapi_ConfigureOfficesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConfigureOfficesAPIService OfficeTransactionsTransactionIdBankCreditLimitsPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.ConfigureOfficesAPI.OfficeTransactionsTransactionIdBankCreditLimitsPut(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigureOfficesAPIService OfficeTransactionsTransactionIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.ConfigureOfficesAPI.OfficeTransactionsTransactionIdDelete(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigureOfficesAPIService OfficesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigureOfficesAPI.OfficesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigureOfficesAPIService OfficesOfficeIdBankCreditLmitsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId int32

		resp, httpRes, err := apiClient.ConfigureOfficesAPI.OfficesOfficeIdBankCreditLmitsGet(context.Background(), officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigureOfficesAPIService OfficesOfficeIdConfigsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId int32

		resp, httpRes, err := apiClient.ConfigureOfficesAPI.OfficesOfficeIdConfigsGet(context.Background(), officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigureOfficesAPIService OfficesOfficeIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId int32

		resp, httpRes, err := apiClient.ConfigureOfficesAPI.OfficesOfficeIdGet(context.Background(), officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigureOfficesAPIService OfficesOfficeIdLinkedConfiguredGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var officeId int32

		resp, httpRes, err := apiClient.ConfigureOfficesAPI.OfficesOfficeIdLinkedConfiguredGet(context.Background(), officeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigureOfficesAPIService OfficesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigureOfficesAPI.OfficesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
