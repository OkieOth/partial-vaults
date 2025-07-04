package vaultfunc_test

import (
	"io"
	"os"
	"testing"

	vault "github.com/sosedoff/ansible-vault-go"
	"github.com/stretchr/testify/require"
)

func TestDecrypt(t *testing.T) {
	inputFile := "../../../resources/tests/example_encrypted.json"
	inputFile2 := "../../../resources/tests/example.json"
	testFile := "../../../temp/example_encrypted.json"
	src, err := os.Open(inputFile)
	require.Nil(t, err)

	dest, err := os.Create(testFile)
	require.Nil(t, err)

	_, err = io.Copy(dest, src)
	require.Nil(t, err)
	src.Close()
	dest.Close()
	decoded, err := vault.DecryptFile(testFile, "secretpassword")
	require.Nil(t, err)

	inputContent, err := os.ReadFile(inputFile2)
	require.Nil(t, err)

	inputStr := string(inputContent)
	require.Equal(t, inputStr, decoded)
}
