/*
Treasury Management API

Testing MiscTransactionsAPIService

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

func Test_openapi_MiscTransactionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MiscTransactionsAPIService MiscTransactionsAccountCodesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MiscTransactionsAPI.MiscTransactionsAccountCodesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MiscTransactionsAPIService MiscTransactionsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MiscTransactionsAPI.MiscTransactionsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MiscTransactionsAPIService MiscTransactionsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MiscTransactionsAPI.MiscTransactionsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MiscTransactionsAPIService MiscTransactionsTransactionIdApprovePut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.MiscTransactionsAPI.MiscTransactionsTransactionIdApprovePut(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MiscTransactionsAPIService MiscTransactionsTransactionIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.MiscTransactionsAPI.MiscTransactionsTransactionIdGet(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}