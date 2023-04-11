package utils

import (
	"errors"
	"os"
	"strconv"
)

var (
	ErrEnvVarEmpty = errors.New("getenv: environment variable empty")
	sdkEnabled     = true
)

func GetenvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, ErrEnvVarEmpty
	}
	return v, nil
}

func GetenvBool(key string) (bool, error) {
	s, err := GetenvStr(key)
	if err != nil {
		return false, err
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return v, nil
}

func IsSDKEnabled() bool {
	return sdkEnabled
}

func DisableSDK() {
	sdkEnabled = false
}

func EnableSDK() {
	sdkEnabled = true
}
