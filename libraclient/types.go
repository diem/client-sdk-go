// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package libraclient

// Address represents account address hex-encoded string
type Address = string

// Amount represents currency amount
type Amount struct {
	Amount   uint64 `json:"amount"`
	Currency string `json:"currency"`
}

// CurrencyInfo is get_currencies response
type CurrencyInfo struct {
	BurnEventsKey               string  `json:"burn_events_key"`
	CancelBurnEventsKey         string  `json:"cancel_burn_events_key"`
	Code                        string  `json:"code"`
	ExchangeRateUpdateEventsKey string  `json:"exchange_rate_update_events_key"`
	FractionalPart              uint64  `json:"fractional_part"`
	MintEventsKey               string  `json:"mint_events_key"`
	PreburnEventsKey            string  `json:"preburn_events_key"`
	ScalingFactor               uint64  `json:"scaling_factor"`
	ToLbrExchangeRate           float64 `json:"to_lbr_exchange_rate"`
}

// AccountRole represents role specific data for account
type AccountRole struct {
	Type string `json:"type"`
	// child_vasp
	ParentVaspAddress string `json:"parent_vasp_address"`

	// parent_vasp / designated_dealer
	HumanName      string `json:"human_name"`
	BaseUrl        string `json:"base_url"`
	ExpirationTime uint64 `json:"expiration_time"`
	ComplianceKey  string `json:"compliance_key"`

	// parent_vasp
	NumChildren uint64 `json:"num_children"`

	// designated_dealer
	PreburnBalances       []Amount `json:"preburn_balances"`
	ReceivedMintEventsKey string   `json:"received_mint_events_key"`
}

// Account is get_account method response
type Account struct {
	AuthenticationKey              string   `json:"authentication_key"`
	Balances                       []Amount `json:"balances"`
	DelegatedKeyRotationCapability bool     `json:"delegated_key_rotation_capability"`
	DelegatedWithdrawalCapability  bool     `json:"delegated_withdrawal_capability"`
	IsFrozen                       bool     `json:"is_frozen"`
	ReceivedEventsKey              string   `json:"received_events_key"`
	// Role                           AccountRole `json:"role"`
	Role           interface{} `json:"role"`
	SentEventsKey  string      `json:"sent_events_key"`
	SequenceNumber uint64      `json:"sequence_number"`
}

// EventData is event type specific data
type EventData struct {
	Type string `json:"type"`
	// burn / cancelburn / preburn / mint / receivedpayment / sentpayment / receivedmint
	Amount Amount `json:"amount"`
	// burn / cancelburn / preburn
	PreburnAddress string `json:"preburn_address"`

	// to_lbr_exchange_rate_update
	CurrencyCode         string  `json:"currency_code"`
	NewToLbrExchangeRate float32 `json:"new_to_lbr_exchange_rate"`

	// receivedpayment / sentpayment
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Metadata string `json:"metadata"`

	// upgrade
	WriteSet string `json:"write_set"`

	// newepoch
	Epoch uint64 `json:"epoch"`

	// newblock
	Round        uint64 `json:"round"`
	Proposer     string `json:"proposer"`
	ProposedTime uint64 `json:"proposed_time"`

	// receivedmint
	DestinationAddress string `json:"destination_address"`
}

// Event data
type Event struct {
	Data               EventData `json:"data"`
	Key                string    `json:"key"`
	SequenceNumber     uint64    `json:"sequence_number"`
	TransactionVersion uint64    `json:"transaction_version"`
}

// Metadata is get_metadata method response
type Metadata struct {
	Timestamp uint64 `json:"timestamp"`
	Version   uint64 `json:"version"`
}

// Script represents decoded transaction script arguments
type Script struct {
	Type string `json:"type"`
	// peer_to_peer_transaction / mint_transaction
	Receiver string `json:"receiver"`
	// peer_to_peer_transaction / mint_transaction
	Amount uint64 `json:"amount"`
	// peer_to_peer_transaction / mint_transaction
	Currency string `json:"currency"`
	// peer_to_peer_transaction
	Metadata string `json:"metadata"`
	// peer_to_peer_transaction
	MetadataSignature string `json:"metadata_signature"`
	// mint_transaction
	AuthKeyPrefix string `json:"auth_key_prefix"`
}

// TransactionData include specific type transaction details
type TransactionData struct {
	Type string `json:"type"`
	// blockmetadata
	TimestampUsecs uint64 `json:"timestamp_usecs"`
	// user
	Sender                  string `json:"sender"`
	SignatureScheme         string `json:"signature_scheme"`
	Signature               string `json:"signature"`
	PublicKey               string `json:"public_key"`
	SequenceNumber          uint64 `json:"sequence_number"`
	ChainID                 uint8  `json:"chain_id"`
	MaxGasAmount            uint64 `json:"max_gas_amount"`
	GasUnitPrice            uint64 `json:"gas_unit_price"`
	GasCurrency             string `json:"gas_currency"`
	ExpirationTimestampSecs uint64 `json:"expiration_timestamp_secs"`
	ScriptHash              string `json:"script_hash"`
	ScriptBytes             string `json:"script_bytes"`
	Script                  Script `json:"script"`
}

// VmStatus represents transaction execution result and error info
type VmStatus struct {
	Type string `json:"type"`

	// execution_failure / move_abort
	Location string `json:"location"`
	// move_abort
	AbortCode uint64 `json:"abort_code"`
	// execution_failure
	FunctionIndex uint16 `json:"function_index"`
	CodeOffset    uint16 `json:"code_offset"`
}

// Transaction represents executed / failed transaction
type Transaction struct {
	Bytes       string          `json:"bytes"`
	Events      []Event         `json:"events"`
	GasUsed     uint64          `json:"gas_used"`
	Hash        string          `json:"hash"`
	Transaction TransactionData `json:"transaction"`
	Version     uint64          `json:"version"`
	VmStatus    interface{}     `json:"vm_status"`
	// VmStatus    VmStatus        `json:"vm_status"`
}

// StateProof is get_state_proof response
type StateProof struct {
	EpochChangeProof         string `json:"epoch_change_proof"`
	LedgerConsistencyProof   string `json:"ledger_consistency_proof"`
	LedgerInfoWithSignatures string `json:"ledger_info_with_signatures"`
}

// AccountStateProof represents account state blob proof
type AccountStateProof struct {
	LedgerInfoToTransactionInfoProof string `json:"ledger_info_to_transaction_info_proof"`
	TransactionInfo                  string `json:"transaction_info"`
	TransactionInfoToAccountProof    string `json:"transaction_info_to_account_proof"`
}

// AccountStateWithProof is get_account_state_with_proof response
type AccountStateWithProof struct {
	Blob    interface{}       `json:"blob,omitempty"`
	Proof   AccountStateProof `json:"proof"`
	Version uint64            `json:"version"`
}
