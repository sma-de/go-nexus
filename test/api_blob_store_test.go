/*
Nexus Repository Manager REST API

Testing BlobStoreAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package nexus

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sma-de/go-nexus"
)

func Test_nexus_BlobStoreAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BlobStoreAPIService ConvertBlobStoreToGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var newNameForOriginal string

		resp, httpRes, err := apiClient.BlobStoreAPI.ConvertBlobStoreToGroup(context.Background(), name, newNameForOriginal).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService CreateBlobStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BlobStoreAPI.CreateBlobStore(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService CreateBlobStore1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BlobStoreAPI.CreateBlobStore1(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService CreateFileBlobStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BlobStoreAPI.CreateFileBlobStore(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService CreateGroupBlobStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BlobStoreAPI.CreateGroupBlobStore(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService DeleteBlobStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.BlobStoreAPI.DeleteBlobStore(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService GetBlobStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.BlobStoreAPI.GetBlobStore(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService GetBlobStore1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.BlobStoreAPI.GetBlobStore1(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService GetFileBlobStoreConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.BlobStoreAPI.GetFileBlobStoreConfiguration(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService GetGroupBlobStoreConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.BlobStoreAPI.GetGroupBlobStoreConfiguration(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService ListBlobStores", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlobStoreAPI.ListBlobStores(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService QuotaStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.BlobStoreAPI.QuotaStatus(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService UpdateBlobStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.BlobStoreAPI.UpdateBlobStore(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService UpdateBlobStore1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.BlobStoreAPI.UpdateBlobStore1(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService UpdateFileBlobStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.BlobStoreAPI.UpdateFileBlobStore(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlobStoreAPIService UpdateGroupBlobStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.BlobStoreAPI.UpdateGroupBlobStore(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
