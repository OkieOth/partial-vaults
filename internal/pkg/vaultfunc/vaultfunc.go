package vaultfunc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/okieoth/pvault/pkg/types"
	vault "github.com/sosedoff/ansible-vault-go"
)

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
	return vault.Encrypt(strValue, password)
}

func Decrypt(value, password string) (any, types.ValueType, error) {
	seperator := "$ANSIBLE_VAULT;"
	index := strings.Index(value, seperator)
	if index == -1 {
		return "", types.STRING, fmt.Errorf("value doesn't contain Ansible vault prefix")
	}
	valToDecode := value[index:]
	decrypted, err := vault.Decrypt(valToDecode, password)
	if err != nil {
		return "", types.STRING, fmt.Errorf("error while decrypting value")
	}
	if i, err := strconv.ParseInt(decrypted, 10, 64); err == nil {
		return i, types.INTEGER, nil
	}
	if b, err := strconv.ParseBool(decrypted); err == nil {
		return b, types.BOOL, nil
	}
	if f, err := strconv.ParseFloat(decrypted, 64); err == nil {
		return f, types.NUMBER, nil
	}
	return decrypted, types.STRING, nil
}
