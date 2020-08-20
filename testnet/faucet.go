// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package testnet

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// MustMint mints coins with retry, and panics if all retries failed.
func MustMint(authKey string, amount uint64, currencyCode string) (ret int) {
	retry := 3
	var err error
	for i := 0; i < retry; i++ {
		ret, err = Mint(authKey, amount, currencyCode)
		if err != nil {
			time.Sleep(1100 * time.Millisecond)
			continue
		}
		return
	}
	panic(fmt.Sprintf("mint coins failed with retry: %s", err.Error()))
}

// Mint mints coints once without retry
func Mint(authKey string, amount uint64, currencyCode string) (int, error) {
	url := fmt.Sprintf(
		"http://faucet.testnet.libra.org/?amount=%d&auth_key=%s&currency_code=%s",
		amount, authKey, currencyCode)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte{}))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("Non 200 response: %s", string(body))
	}
	return strconv.Atoi(string(body))
}
