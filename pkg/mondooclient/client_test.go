// Copyright (c) Mondoo, Inc.
// SPDX-License-Identifier: BUSL-1.1

package mondooclient_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/yaml"

	"go.mondoo.com/mondoo-operator/pkg/mondooclient"
	"go.mondoo.com/mondoo-operator/pkg/mondooclient/fakeserver"
)

var webhookPayload = mustRead("../../tests/data/webhook-payload.json")

func TestScanner(t *testing.T) {
	testserver := fakeserver.FakeServer()
	url := testserver.URL
	token := ""

	// To test with a real client, just set
	// url := "http://127.0.0.1:8990"
	// token := "<token here>"

	// do client request
	mClient := mondooclient.NewClient(mondooclient.ClientOptions{
		ApiEndpoint: url,
		Token:       token,
	})

	// Run Health Check
	healthResp, err := mClient.HealthCheck(context.Background(), &mondooclient.HealthCheckRequest{})
	require.NoError(t, err)
	assert.True(t, healthResp.Status == "SERVING")

	request := admission.Request{}
	err = yaml.Unmarshal(webhookPayload, &request)
	require.NoError(t, err)

	result, err := mClient.RunAdmissionReview(context.Background(), &mondooclient.AdmissionReviewJob{
		Labels: map[string]string{
			"k8s.mondoo.com/author":     request.UserInfo.Username,
			"k8s.mondoo.com/operation":  string(request.Operation),
			"k8s.mondoo.com/cluster-id": "SOME-ID-HERE",
		},
	})
	require.NoError(t, err)
	require.NotNil(t, result)

	// check if the scan passed
	if assert.NotNil(t, result.WorstScore, "nil WorstScore") {
		passed := result.WorstScore.Type == 2 && result.WorstScore.Value == 100
		assert.True(t, passed)
	}
}

func TestScanner_ScanKubernetesResources(t *testing.T) {
	testserver := fakeserver.FakeServer()
	url := testserver.URL
	token := ""

	// To test with a real client, just set
	// url := "http://127.0.0.1:8989"
	// token := "abcdefgh"

	// do client request
	mClient := mondooclient.NewClient(mondooclient.ClientOptions{
		ApiEndpoint: url,
		Token:       token,
	})

	// Run Health Check
	healthResp, err := mClient.HealthCheck(context.Background(), &mondooclient.HealthCheckRequest{})
	require.NoError(t, err)
	assert.True(t, healthResp.Status == "SERVING")

	result, err := mClient.ScanKubernetesResources(context.Background(), &mondooclient.ScanKubernetesResourcesOpts{})
	require.NoError(t, err)
	require.NotNil(t, result)

	// check if the scan passed
	if assert.NotNil(t, result.WorstScore, "nil WorstScore") {
		passed := result.WorstScore.Type == 2 && result.WorstScore.Value == 100
		assert.True(t, passed)
	}
}

func mustRead(file string) []byte {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic("couldn't read in file")
	}
	return bytes
}
