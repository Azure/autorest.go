// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azkeys_test

import (
	"azkeys"
	"azkeys/fake"
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/stretchr/testify/require"
)

func TestFakeBackupKey(t *testing.T) {
	const fakeKeyName = "fakeKey"
	fakeKeyBlob := []byte{1, 2, 3}
	server := fake.KeyVaultServer{
		BackupKey: func(ctx context.Context, keyName string, options *azkeys.KeyVaultClientBackupKeyOptions) (resp azfake.Responder[azkeys.KeyVaultClientBackupKeyResponse], errResp azfake.ErrorResponder) {
			if keyName != fakeKeyName {
				errResp.SetError(fmt.Errorf("bad fake key name %s", keyName))
				return
			}
			resp.SetResponse(http.StatusOK, azkeys.KeyVaultClientBackupKeyResponse{
				BackupKeyResult: azkeys.BackupKeyResult{
					Value: fakeKeyBlob,
				},
			}, nil)
			return
		},
	}

	client, err := azkeys.NewKeyVaultClient("https://contoso.com/fake/vault", &azcore.ClientOptions{
		Transport: fake.NewKeyVaultServerTransport(&server),
	})
	require.NoError(t, err)

	resp, err := client.BackupKey(context.Background(), fakeKeyName, nil)
	require.NoError(t, err)
	require.Equal(t, fakeKeyBlob, resp.Value)
}