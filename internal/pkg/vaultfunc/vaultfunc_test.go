package vaultfunc_test

import (
	"io"
	"os"
	"testing"

	"github.com/okieoth/pvault/internal/pkg/vaultfunc"
	"github.com/okieoth/pvault/pkg/types"
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

func TestEncryptDecryptInt(t *testing.T) {
	password := "_strengGeheim_"
	tests := []int64{42, -45, 0}
	for _, v := range tests {
		encrypted, err := vaultfunc.Encrypt(v, types.INTEGER, password)
		require.Nil(t, err)
		decrypted, valueType, err := vaultfunc.Decrypt(encrypted, password)
		require.Nil(t, err)
		require.Equal(t, types.INTEGER, valueType)
		require.Equal(t, v, decrypted)
	}
}

func TestEncryptDecryptBool(t *testing.T) {
	password := "_strengGeheim_"
	tests := []bool{true, false}
	for _, v := range tests {
		encrypted, err := vaultfunc.Encrypt(v, types.BOOL, password)
		require.Nil(t, err)
		decrypted, valueType, err := vaultfunc.Decrypt(encrypted, password)
		require.Nil(t, err)
		require.Equal(t, types.BOOL, valueType)
		require.Equal(t, v, decrypted)
	}
}

func TestEncryptDecryptFloat(t *testing.T) {
	password := "_strengGeheim_"
	tests := []float64{42.78, -13.8, 0}
	for _, v := range tests {
		encrypted, err := vaultfunc.Encrypt(v, types.NUMBER, password)
		require.Nil(t, err)
		decrypted, valueType, err := vaultfunc.Decrypt(encrypted, password)
		require.Nil(t, err)
		require.Equal(t, types.NUMBER, valueType)
		require.Equal(t, v, decrypted)
	}
}

func TestEncryptDecryptString(t *testing.T) {
	password := "_strengGeheim_"
	tests := []string{"Some words with äoß", "1234x", ""}
	for _, v := range tests {
		encrypted, err := vaultfunc.Encrypt(v, types.STRING, password)
		require.Nil(t, err)
		decrypted, valueType, err := vaultfunc.Decrypt(encrypted, password)
		require.Nil(t, err)
		require.Equal(t, types.STRING, valueType)
		require.Equal(t, v, decrypted)
	}
}
