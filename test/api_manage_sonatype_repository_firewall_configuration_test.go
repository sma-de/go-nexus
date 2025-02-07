/*
Nexus Repository Manager REST API

Testing ManageSonatypeRepositoryFirewallConfigurationAPIService

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

func Test_nexus_ManageSonatypeRepositoryFirewallConfigurationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ManageSonatypeRepositoryFirewallConfigurationAPIService DisableIq", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.DisableIq(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManageSonatypeRepositoryFirewallConfigurationAPIService EnableIq", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.EnableIq(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManageSonatypeRepositoryFirewallConfigurationAPIService GetAllAuditStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.GetAllAuditStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManageSonatypeRepositoryFirewallConfigurationAPIService GetAuditStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var repositoryName string

		httpRes, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.GetAuditStatus(context.Background(), repositoryName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManageSonatypeRepositoryFirewallConfigurationAPIService GetConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.GetConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManageSonatypeRepositoryFirewallConfigurationAPIService ManageAudit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.ManageAudit(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManageSonatypeRepositoryFirewallConfigurationAPIService UpdateConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.UpdateConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManageSonatypeRepositoryFirewallConfigurationAPIService VerifyConnection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.VerifyConnection(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
