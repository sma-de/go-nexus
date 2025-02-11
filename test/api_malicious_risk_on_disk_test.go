/*
Nexus Repository Manager REST API

Testing MaliciousRiskOnDiskAPIService

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

func Test_nexus_MaliciousRiskOnDiskAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MaliciousRiskOnDiskAPIService GetEnabledRegistries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.MaliciousRiskOnDiskAPI.GetEnabledRegistries(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MaliciousRiskOnDiskAPIService GetMaliciousRiskOnDiskCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.MaliciousRiskOnDiskAPI.GetMaliciousRiskOnDiskCount(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
