// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libraclient

import (
	"github.com/libra/libra-client-sdk-go/librajsonrpctypes"
)

// Amount represents currency amount
type Amount = librajsonrpctypes.Amount

// CurrencyInfo is get_currencies response
type CurrencyInfo = librajsonrpctypes.CurrencyInfo

// AccountRole represents role specific data for account
type AccountRole = librajsonrpctypes.AccountRole

// Account is get_account method response
type Account = librajsonrpctypes.Account

// EventData is event type specific data
type EventData = librajsonrpctypes.EventData

// Event data
type Event = librajsonrpctypes.Event

// Metadata is get_metadata method response
type Metadata = librajsonrpctypes.Metadata

// Script represents decoded transaction script arguments
type Script = librajsonrpctypes.Script

// TransactionData include specific type transaction details
type TransactionData = librajsonrpctypes.TransactionData

// VmStatus represents transaction execution result and error info
type VmStatus = librajsonrpctypes.VMStatus

// Transaction represents executed / failed transaction
type Transaction = librajsonrpctypes.Transaction

// StateProof is get_state_proof response
type StateProof = librajsonrpctypes.StateProof

// AccountStateProof represents account state blob proof
type AccountStateProof = librajsonrpctypes.AccountStateProof

// AccountStateWithProof is get_account_state_with_proof response
type AccountStateWithProof = librajsonrpctypes.AccountStateWithProof
