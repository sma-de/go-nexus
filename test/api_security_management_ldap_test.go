/*
Nexus Repository Manager REST API

Testing SecurityManagementLDAPAPIService

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

func Test_nexus_SecurityManagementLDAPAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SecurityManagementLDAPAPIService ChangeOrder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementLDAPAPI.ChangeOrder(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementLDAPAPIService CreateLdapServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementLDAPAPI.CreateLdapServer(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementLDAPAPIService DeleteLdapServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.SecurityManagementLDAPAPI.DeleteLdapServer(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementLDAPAPIService GetLdapServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.SecurityManagementLDAPAPI.GetLdapServer(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementLDAPAPIService GetLdapServers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.SecurityManagementLDAPAPI.GetLdapServers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityManagementLDAPAPIService UpdateLdapServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.SecurityManagementLDAPAPI.UpdateLdapServer(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
