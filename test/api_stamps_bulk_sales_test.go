/*
Treasury Management API

Testing StampsBulkSalesAPIService

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

func Test_openapi_StampsBulkSalesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StampsBulkSalesAPIService StampBulkSalesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StampsBulkSalesAPI.StampBulkSalesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StampsBulkSalesAPIService StampBulkSalesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StampsBulkSalesAPI.StampBulkSalesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StampsBulkSalesAPIService StampBulkSalesTransactionIdApprovePut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.StampsBulkSalesAPI.StampBulkSalesTransactionIdApprovePut(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StampsBulkSalesAPIService StampBulkSalesTransactionIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var transactionId string

		resp, httpRes, err := apiClient.StampsBulkSalesAPI.StampBulkSalesTransactionIdGet(context.Background(), transactionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
