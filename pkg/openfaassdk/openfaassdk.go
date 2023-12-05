package openfaassdk

import (
	"fmt"

	"github.com/tommoulard/side/pkg/files"
)

func ReadSecret(secretName string) (string, error) {
	fName := "/var/openfaas/secrets/" + secretName

	content, err := files.GetFileContent(fName)
	if err != nil {
		return "", fmt.Errorf("failed to read secret %s: %w", secretName, err)
	}

	return content, nil
}
