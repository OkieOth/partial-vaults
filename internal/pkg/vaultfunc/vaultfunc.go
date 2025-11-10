package vaultfunc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/okieoth/pvault/pkg/types"
	vault "github.com/sosedoff/ansible-vault-go"
)

const ansibleVaultStr string = "$ANSIBLE_VAULT"

func Encrypt(value any, valueType types.ValueType, password string) (string, error) {
	var strValue string
	switch valueType {
	case types.BOOL:
		strValue = fmt.Sprintf("%v", value)
	case types.NUMBER:
		strValue = fmt.Sprintf("%f", value)
	case types.INTEGER:
		strValue = fmt.Sprintf("%d", value)
	case types.STRING:
		strValue = fmt.Sprintf("%s", value)
	default:
		return "", fmt.Errorf("value type isn't supported for encryption: %s", valueType)
	}
	if strings.Contains(strValue, ansibleVaultStr) {
		return strValue, nil
	}
	return vault.Encrypt(strValue, password)
}

func Decrypt(value, password string) (any, types.ValueType, error) {
	separator := "$ANSIBLE_VAULT;"
	index := strings.Index(value, separator)
	if index == -1 {
		return "", types.STRING, fmt.Errorf("value doesn't contain Ansible vault prefix")
	}
	valToDecode := value[index:]
	decrypted, err := vault.Decrypt(valToDecode, password)
	if err != nil {
		return "", types.STRING, fmt.Errorf("error while decrypting value: %v", err)
	}
	v, vt := InputType(decrypted)
	return v, vt, nil
}

func InputType(value string) (any, types.ValueType) {
	if i, err := strconv.ParseInt(value, 10, 64); err == nil {
		return i, types.INTEGER
	}
	if b, err := strconv.ParseBool(value); err == nil {
		return b, types.BOOL
	}
	if f, err := strconv.ParseFloat(value, 64); err == nil {
		return f, types.NUMBER
	}
	return value, types.STRING

}
