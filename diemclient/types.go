// Copyright (c) The Diem Core Contributors
// SPDX-License-Identifier: Apache-2.0

package diemclient

import (
	"github.com/diem/client-sdk-go/diemjsonrpctypes"
)

// Amount represents currency amount
type Amount = diemjsonrpctypes.Amount

// CurrencyInfo is get_currencies response
type CurrencyInfo = diemjsonrpctypes.CurrencyInfo

// AccountRole represents role specific data for account
type AccountRole = diemjsonrpctypes.AccountRole

// Account is get_account method response
type Account = diemjsonrpctypes.Account

// EventData is event type specific data
type EventData = diemjsonrpctypes.EventData

// Event data
type Event = diemjsonrpctypes.Event

// Metadata is get_metadata method response
type Metadata = diemjsonrpctypes.Metadata

// Script represents decoded transaction script arguments
type Script = diemjsonrpctypes.Script

// TransactionData include specific type transaction details
type TransactionData = diemjsonrpctypes.TransactionData

// VmStatus represents transaction execution result and error info
type VmStatus = diemjsonrpctypes.VMStatus

// Transaction represents executed / failed transaction
type Transaction = diemjsonrpctypes.Transaction

// StateProof is get_state_proof response
type StateProof = diemjsonrpctypes.StateProof

// AccountStateProof represents account state blob proof
type AccountStateProof = diemjsonrpctypes.AccountStateProof

// AccountStateWithProof is get_account_state_with_proof response
type AccountStateWithProof = diemjsonrpctypes.AccountStateWithProof
