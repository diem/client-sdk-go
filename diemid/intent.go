// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

// This file implemenets Diem Intent Identifier proposal
// https://github.com/diem/lip/blob/master/lips/lip-5.md

package diemid

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	DiemScheme        = "diem"
	CurrencyParamName = "c"
	AmountParamName   = "am"
)

// Params for Intent
type Params struct {
	Currency string
	Amount   *uint64
}

// Intent captures all parts of intent identifier
type Intent struct {
	Account Account
	Params  Params
}

// DecodeToIntent decode given intent string to `Intent`.
// Given `networkPrefix` is used to validate intent account identifier network prefix.
func DecodeToIntent(networkPrefix NetworkPrefix, intent string) (*Intent, error) {
	u, err := url.ParseRequestURI(intent)
	if err != nil {
		return nil, fmt.Errorf("invalid intent identifier: %s", err.Error())
	}
	if u.Scheme != DiemScheme {
		return nil, fmt.Errorf("invalid intent scheme: %s", u.Scheme)
	}
	account, err := DecodeToAccount(networkPrefix, u.Host)
	if err != nil {
		return nil, fmt.Errorf("invalid account identifier: %s", err.Error())
	}
	return &Intent{
		Account: *account,
		Params: Params{
			Currency: u.Query().Get(CurrencyParamName),
			Amount:   toIntPtr(u.Query().Get(AmountParamName)),
		},
	}, nil
}

func (i *Intent) Encode() (string, error) {
	encoded, err := i.Account.Encode()
	if err != nil {
		return "", fmt.Errorf("encode account identifier failed: %s", err.Error())
	}
	u := url.URL{Scheme: DiemScheme, Host: encoded}
	q := u.Query()
	if i.Params.Currency != "" {
		q.Add(CurrencyParamName, i.Params.Currency)
	}
	if i.Params.Amount != nil {
		q.Add(AmountParamName, strconv.FormatUint(*i.Params.Amount, 10))
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func toIntPtr(str string) *uint64 {
	ret, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return nil
	}
	return &ret
}
