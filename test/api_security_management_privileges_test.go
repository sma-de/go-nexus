/*
Nexus Repository Manager REST API

Testing SecurityManagementPrivilegesAPIService

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

func Test_nexus_SecurityManagementPrivilegesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SecurityManagementPrivilegesAPIService CreatePrivilege", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementPrivilegesAPI.CreatePrivilege(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService CreatePrivilege1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementPrivilegesAPI.CreatePrivilege1(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService CreatePrivilege2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementPrivilegesAPI.CreatePrivilege2(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService CreatePrivilege3", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementPrivilegesAPI.CreatePrivilege3(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService CreatePrivilege4", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementPrivilegesAPI.CreatePrivilege4(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService CreatePrivilege5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementPrivilegesAPI.CreatePrivilege5(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService DeletePrivilege", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var privilegeName string

		httpRes, err := apiClient.SecurityManagementPrivilegesAPI.DeletePrivilege(context.Background(), privilegeName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService GetPrivilege", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var privilegeName string

		resp, httpRes, err := apiClient.SecurityManagementPrivilegesAPI.GetPrivilege(context.Background(), privilegeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService GetPrivileges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SecurityManagementPrivilegesAPI.GetPrivileges(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService UpdatePrivilege", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var privilegeName string

		httpRes, err := apiClient.SecurityManagementPrivilegesAPI.UpdatePrivilege(context.Background(), privilegeName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService UpdatePrivilege1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var privilegeName string

		httpRes, err := apiClient.SecurityManagementPrivilegesAPI.UpdatePrivilege1(context.Background(), privilegeName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService UpdatePrivilege2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var privilegeName string

		httpRes, err := apiClient.SecurityManagementPrivilegesAPI.UpdatePrivilege2(context.Background(), privilegeName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService UpdatePrivilege3", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var privilegeName string

		httpRes, err := apiClient.SecurityManagementPrivilegesAPI.UpdatePrivilege3(context.Background(), privilegeName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService UpdatePrivilege4", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var privilegeName string

		httpRes, err := apiClient.SecurityManagementPrivilegesAPI.UpdatePrivilege4(context.Background(), privilegeName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementPrivilegesAPIService UpdatePrivilege5", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var privilegeName string

		httpRes, err := apiClient.SecurityManagementPrivilegesAPI.UpdatePrivilege5(context.Background(), privilegeName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
