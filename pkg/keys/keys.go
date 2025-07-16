package keys

import (
	"fmt"
	"strings"

	"github.com/okieoth/pvault/pkg/types"
)

func IsEncrypted(input any, vt types.ValueType) (bool, string, error) {
	if vt != types.STRING {
		return false, "", nil
	}

	valueToDecrypt, ok := input.(string)
	if !ok {
		return false, "", fmt.Errorf("error while casting value to decrypt to string")
	}

	separator := "$ANSIBLE_VAULT;"
	index := strings.Index(valueToDecrypt, separator)
	if index == -1 {
		return false, valueToDecrypt, nil
	}
	valueToDecrypt = valueToDecrypt[index:]
	return true, valueToDecrypt, nil
}

func testForEncryptedImpl(input any, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
	encrypted, _, err := IsEncrypted(input, vt)
	if !encrypted {
		return input, vt, types.HANDLING_SKIP, err
	}
	return input, vt, types.HANDLING_PROCESS, nil
}

func TestEncryptedProcessor() types.ProcessFunc {
	return func(input any, vt types.ValueType, keyPath string) (any, types.ValueType, types.ProcessHandling, error) {
		return testForEncryptedImpl(input, vt, keyPath)
	}
}
