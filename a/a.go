package a

import (
	"encoding/hex"
	"strings"

	"github.com/smartcontractkit/chainlink-common/pkg/types/chains/evm"
)

func hexToAddress(hexValue string) (evm.Address, error) {
	// Strip the 0x prefix if it exists
	if strings.HasPrefix(hexValue, "0x") {
		hexValue = hexValue[2:]
	}

	addressBytes, err := hex.DecodeString(hexValue)
	if err != nil {
		return evm.Address{}, err
	}
	return evm.Address(addressBytes), nil
}

func addressToHex(address evm.Address) string {
	// Add the 0x prefix
	return "0x" + hex.EncodeToString(address[:])
}

func Hello() string {
	address, err := hexToAddress("0x1234567890123456789012345678901234567890")
	if err != nil {
		return "Error: " + err.Error()
	}

	return "Address: " + addressToHex(address) + "; Hello"
}
