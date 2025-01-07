/*
Nexus Repository Manager REST API

Testing ManageSonatypeHTTPSystemSettingsAPIService

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

func Test_nexus_ManageSonatypeHTTPSystemSettingsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ManageSonatypeHTTPSystemSettingsAPIService GetHttpSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ManageSonatypeHTTPSystemSettingsAPI.GetHttpSettings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManageSonatypeHTTPSystemSettingsAPIService ResetHttpSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ManageSonatypeHTTPSystemSettingsAPI.ResetHttpSettings(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManageSonatypeHTTPSystemSettingsAPIService UpdateHttpSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ManageSonatypeHTTPSystemSettingsAPI.UpdateHttpSettings(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
