package stdlib


import (
	"fmt"
	"github.com/libra/libra-client-sdk-go/libratypes"
)


// Structured representation of a call into a known Move script.
type ScriptCall interface {
	isScriptCall()
}

// Add a `Currency` balance to `account`, which will enable `account` to send and receive
// `Libra<Currency>`.
// Aborts with NOT_A_CURRENCY if `Currency` is not an accepted currency type in the Libra system
// Aborts with `LibraAccount::ADD_EXISTING_CURRENCY` if the account already holds a balance in
// `Currency`.
type ScriptCall__AddCurrencyToAccount struct {
	Currency libratypes.TypeTag
}

func (*ScriptCall__AddCurrencyToAccount) isScriptCall() {}

// Add the `KeyRotationCapability` for `to_recover_account` to the `RecoveryAddress` resource under `recovery_address`.
//
// ## Aborts
// * Aborts with `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` if `account` has already delegated its `KeyRotationCapability`.
// * Aborts with `RecoveryAddress:ENOT_A_RECOVERY_ADDRESS` if `recovery_address` does not have a `RecoveryAddress` resource.
// * Aborts with `RecoveryAddress::EINVALID_KEY_ROTATION_DELEGATION` if `to_recover_account` and `recovery_address` do not belong to the same VASP.
type ScriptCall__AddRecoveryRotationCapability struct {
	RecoveryAddress libratypes.AccountAddress
}

func (*ScriptCall__AddRecoveryRotationCapability) isScriptCall() {}

// Append the `hash` to script hashes list allowed to be executed by the network.
type ScriptCall__AddToScriptAllowList struct {
	Hash []byte
	SlidingNonce uint64
}

func (*ScriptCall__AddToScriptAllowList) isScriptCall() {}

// Add `new_validator` to the validator set.
// Fails if the `new_validator` address is already in the validator set
// or does not have a `ValidatorConfig` resource stored at the address.
// Emits a NewEpochEvent.
type ScriptCall__AddValidatorAndReconfigure struct {
	SlidingNonce uint64
	ValidatorName []byte
	ValidatorAddress libratypes.AccountAddress
}

func (*ScriptCall__AddValidatorAndReconfigure) isScriptCall() {}

// Permanently destroy the `Token`s stored in the oldest burn request under the `Preburn` resource.
// This will only succeed if `account` has a `MintCapability<Token>`, a `Preburn<Token>` resource
// exists under `preburn_address`, and there is a pending burn request.
// sliding_nonce is a unique nonce for operation, see sliding_nonce.move for details
type ScriptCall__Burn struct {
	Token libratypes.TypeTag
	SlidingNonce uint64
	PreburnAddress libratypes.AccountAddress
}

func (*ScriptCall__Burn) isScriptCall() {}

// Burn transaction fees that have been collected in the given `currency`
// and relinquish to the association. The currency must be non-synthetic.
type ScriptCall__BurnTxnFees struct {
	CoinType libratypes.TypeTag
}

func (*ScriptCall__BurnTxnFees) isScriptCall() {}

// Cancel the oldest burn request from `preburn_address` and return the funds.
// Fails if the sender does not have a published `BurnCapability<Token>`.
type ScriptCall__CancelBurn struct {
	Token libratypes.TypeTag
	PreburnAddress libratypes.AccountAddress
}

func (*ScriptCall__CancelBurn) isScriptCall() {}

// Create a `ChildVASP` account for sender `parent_vasp` at `child_address` with a balance of
// `child_initial_balance` in `CoinType` and an initial authentication_key
// `auth_key_prefix | child_address`.
// If `add_all_currencies` is true, the child address will have a zero balance in all available
// currencies in the system.
// This account will a child of the transaction sender, which must be a ParentVASP.
//
// ## Aborts
// The transaction will abort:
//
// * If `parent_vasp` is not a parent vasp with error: `Roles::EINVALID_PARENT_ROLE`
// * If `child_address` already exists with error: `Roles::EROLE_ALREADY_ASSIGNED`
// * If `parent_vasp` already has 256 child accounts with error: `VASP::ETOO_MANY_CHILDREN`
// * If `CoinType` is not a registered currency with error: `LibraAccount::ENOT_A_CURRENCY`
// * If `parent_vasp`'s withdrawal capability has been extracted with error:  `LibraAccount::EWITHDRAWAL_CAPABILITY_ALREADY_EXTRACTED`
// * If `parent_vasp` doesn't hold `CoinType` and `child_initial_balance > 0` with error: `LibraAccount::EPAYER_DOESNT_HOLD_CURRENCY`
// * If `parent_vasp` doesn't at least `child_initial_balance` of `CoinType` in its account balance with error: `LibraAccount::EINSUFFICIENT_BALANCE`
type ScriptCall__CreateChildVaspAccount struct {
	CoinType libratypes.TypeTag
	ChildAddress libratypes.AccountAddress
	AuthKeyPrefix []byte
	AddAllCurrencies bool
	ChildInitialBalance uint64
}

func (*ScriptCall__CreateChildVaspAccount) isScriptCall() {}

// Create an account with the DesignatedDealer role at `addr` with authentication key
// `auth_key_prefix` | `addr` and a 0 balance of type `Currency`. If `add_all_currencies` is true,
// 0 balances for all available currencies in the system will also be added. This can only be
// invoked by an account with the TreasuryCompliance role.
type ScriptCall__CreateDesignatedDealer struct {
	Currency libratypes.TypeTag
	SlidingNonce uint64
	Addr libratypes.AccountAddress
	AuthKeyPrefix []byte
	HumanName []byte
	AddAllCurrencies bool
}

func (*ScriptCall__CreateDesignatedDealer) isScriptCall() {}

// Create an account with the ParentVASP role at `address` with authentication key
// `auth_key_prefix` | `new_account_address` and a 0 balance of type `currency`. If
// `add_all_currencies` is true, 0 balances for all available currencies in the system will
// also be added. This can only be invoked by an Association account.
// `sliding_nonce` is a unique nonce for operation, see sliding_nonce.move for details.
type ScriptCall__CreateParentVaspAccount struct {
	CoinType libratypes.TypeTag
	SlidingNonce uint64
	NewAccountAddress libratypes.AccountAddress
	AuthKeyPrefix []byte
	HumanName []byte
	AddAllCurrencies bool
}

func (*ScriptCall__CreateParentVaspAccount) isScriptCall() {}

// Extract the `KeyRotationCapability` for `recovery_account` and publish it in a
// `RecoveryAddress` resource under  `account`.
// ## Aborts
// * Aborts with `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` if `account` has already delegated its `KeyRotationCapability`.
// * Aborts with `RecoveryAddress::ENOT_A_VASP` if `account` is not a ParentVASP or ChildVASP
type ScriptCall__CreateRecoveryAddress struct {
}

func (*ScriptCall__CreateRecoveryAddress) isScriptCall() {}

// Create a validator account at `new_validator_address` with `auth_key_prefix`and human_name.
type ScriptCall__CreateValidatorAccount struct {
	SlidingNonce uint64
	NewAccountAddress libratypes.AccountAddress
	AuthKeyPrefix []byte
	HumanName []byte
}

func (*ScriptCall__CreateValidatorAccount) isScriptCall() {}

// Create a validator operator account at `new_validator_address` with `auth_key_prefix`and human_name.
type ScriptCall__CreateValidatorOperatorAccount struct {
	SlidingNonce uint64
	NewAccountAddress libratypes.AccountAddress
	AuthKeyPrefix []byte
	HumanName []byte
}

func (*ScriptCall__CreateValidatorOperatorAccount) isScriptCall() {}

// Freeze account `address`. Initiator must be authorized.
// `sliding_nonce` is a unique nonce for operation, see sliding_nonce.move for details.
type ScriptCall__FreezeAccount struct {
	SlidingNonce uint64
	ToFreezeAccount libratypes.AccountAddress
}

func (*ScriptCall__FreezeAccount) isScriptCall() {}

// Mint `amount_lbr` LBR from the sending account's constituent coins and deposits the
// resulting LBR into the sending account.
type ScriptCall__MintLbr struct {
	AmountLbr uint64
}

func (*ScriptCall__MintLbr) isScriptCall() {}

// Transfer `amount` coins of type `Currency` from `payer` to `payee` with (optional) associated
// `metadata` and an (optional) `metadata_signature` on the message
// `metadata` | `Signer::address_of(payer)` | `amount` | `DualAttestation::DOMAIN_SEPARATOR`.
// The `metadata` and `metadata_signature` parameters are only required if `amount` >=
// `DualAttestation::get_cur_microlibra_limit` LBR and `payer` and `payee` are distinct VASPs.
// However, a transaction sender can opt in to dual attestation even when it is not required (e.g., a DesignatedDealer -> VASP payment) by providing a non-empty `metadata_signature`.
// Standardized `metadata` LCS format can be found in `libra_types::transaction::metadata::Metadata`.
//
// ## Events
// When this script executes without aborting, it emits two events:
// `SentPaymentEvent { amount, currency_code = Currency, payee, metadata }`
// on `payer`'s `LibraAccount::sent_events` handle, and
//  `ReceivedPaymentEvent { amount, currency_code = Currency, payer, metadata }`
// on `payee`'s `LibraAccount::received_events` handle.
//
// ## Common Aborts
// These aborts can in occur in any payment.
// * Aborts with `LibraAccount::EINSUFFICIENT_BALANCE` if `amount` is greater than `payer`'s balance in `Currency`.
// * Aborts with `LibraAccount::ECOIN_DEPOSIT_IS_ZERO` if `amount` is zero.
// * Aborts with `LibraAccount::EPAYEE_DOES_NOT_EXIST` if no account exists at the address `payee`.
// * Aborts with `LibraAccount::EPAYEE_CANT_ACCEPT_CURRENCY_TYPE` if an account exists at `payee`, but it does not accept payments in `Currency`.
//
// ## Dual Attestation Aborts
// These aborts can occur in any payment subject to dual attestation.
// * Aborts with `DualAttestation::EMALFORMED_METADATA_SIGNATURE` if `metadata_signature`'s is not 64 bytes.
// * Aborts with `DualAttestation:EINVALID_METADATA_SIGNATURE` if `metadata_signature` does not verify on the message `metadata` | `payer` | `value` | `DOMAIN_SEPARATOR` using the `compliance_public_key` published in the `payee`'s `DualAttestation::Credential` resource.
//
// ## Other Aborts
// These aborts should only happen when `payer` or `payee` have account limit restrictions or
// have been frozen by Libra administrators.
// * Aborts with `LibraAccount::EWITHDRAWAL_EXCEEDS_LIMITS` if `payer` has exceeded their daily
// withdrawal limits.
// * Aborts with `LibraAccount::EDEPOSIT_EXCEEDS_LIMITS` if `payee` has exceeded their daily deposit limits.
// * Aborts with `LibraAccount::EACCOUNT_FROZEN` if `payer`'s account is frozen.
type ScriptCall__PeerToPeerWithMetadata struct {
	Currency libratypes.TypeTag
	Payee libratypes.AccountAddress
	Amount uint64
	Metadata []byte
	MetadataSignature []byte
}

func (*ScriptCall__PeerToPeerWithMetadata) isScriptCall() {}

// Preburn `amount` `Token`s from `account`.
// This will only succeed if `account` already has a published `Preburn<Token>` resource.
type ScriptCall__Preburn struct {
	Token libratypes.TypeTag
	Amount uint64
}

func (*ScriptCall__Preburn) isScriptCall() {}

// (1) Rotate the authentication key of the sender to `public_key`
// (2) Publish a resource containing a 32-byte ed25519 public key and the rotation capability
//     of the sender under the sender's address.
// Aborts if the sender already has a `SharedEd25519PublicKey` resource.
// Aborts if the length of `new_public_key` is not 32.
type ScriptCall__PublishSharedEd25519PublicKey struct {
	PublicKey []byte
}

func (*ScriptCall__PublishSharedEd25519PublicKey) isScriptCall() {}

// Set validator's config locally.
// Does not emit NewEpochEvent, the config is NOT changed in the validator set.
type ScriptCall__RegisterValidatorConfig struct {
	ValidatorAccount libratypes.AccountAddress
	ConsensusPubkey []byte
	ValidatorNetworkIdentityPubkey []byte
	ValidatorNetworkAddress []byte
	FullnodesNetworkIdentityPubkey []byte
	FullnodesNetworkAddress []byte
}

func (*ScriptCall__RegisterValidatorConfig) isScriptCall() {}

// Removes a validator from the validator set.
// Fails if the validator_address is not in the validator set.
// Emits a NewEpochEvent.
type ScriptCall__RemoveValidatorAndReconfigure struct {
	SlidingNonce uint64
	ValidatorName []byte
	ValidatorAddress libratypes.AccountAddress
}

func (*ScriptCall__RemoveValidatorAndReconfigure) isScriptCall() {}

// Rotate the sender's authentication key to `new_key`.
// `new_key` should be a 256 bit sha3 hash of an ed25519 public key.
// * Aborts with `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` if the `KeyRotationCapability` for `account` has already been extracted.
// * Aborts with `0` if the key rotation capability held by the account doesn't match the sender's address.
// * Aborts with `LibraAccount::EMALFORMED_AUTHENTICATION_KEY` if the length of `new_key` != 32.
type ScriptCall__RotateAuthenticationKey struct {
	NewKey []byte
}

func (*ScriptCall__RotateAuthenticationKey) isScriptCall() {}

// Rotate `account`'s authentication key to `new_key`.
// `new_key` should be a 256 bit sha3 hash of an ed25519 public key. This script also takes
// `sliding_nonce`, as a unique nonce for this operation. See sliding_nonce.move for details.
type ScriptCall__RotateAuthenticationKeyWithNonce struct {
	SlidingNonce uint64
	NewKey []byte
}

func (*ScriptCall__RotateAuthenticationKeyWithNonce) isScriptCall() {}

// Rotate `account`'s authentication key to `new_key`.
// `new_key` should be a 256 bit sha3 hash of an ed25519 public key. This script also takes
// `sliding_nonce`, as a unique nonce for this operation. See sliding_nonce.move for details.
type ScriptCall__RotateAuthenticationKeyWithNonceAdmin struct {
	SlidingNonce uint64
	NewKey []byte
}

func (*ScriptCall__RotateAuthenticationKeyWithNonceAdmin) isScriptCall() {}

// Rotate the authentication key of `account` to `new_key` using the `KeyRotationCapability`
// stored under `recovery_address`.
//
// ## Aborts
// * Aborts with `RecoveryAddress::ENOT_A_RECOVERY_ADDRESS` if `recovery_address` does not have a `RecoveryAddress` resource
// * Aborts with `RecoveryAddress::ECANNOT_ROTATE_KEY` if `account` is not `recovery_address` or `to_recover`.
// * Aborts with `LibraAccount::EMALFORMED_AUTHENTICATION_KEY` if `new_key` is not 32 bytes.
// * Aborts with `RecoveryAddress::ECANNOT_ROTATE_KEY` if `account` has not delegated its `KeyRotationCapability` to `recovery_address`.
type ScriptCall__RotateAuthenticationKeyWithRecoveryAddress struct {
	RecoveryAddress libratypes.AccountAddress
	ToRecover libratypes.AccountAddress
	NewKey []byte
}

func (*ScriptCall__RotateAuthenticationKeyWithRecoveryAddress) isScriptCall() {}

// Rotate `account`'s base URL to `new_url` and its compliance public key to `new_key`.
// Aborts if `account` is not a ParentVASP or DesignatedDealer
// Aborts if `new_key` is not a well-formed public key
type ScriptCall__RotateDualAttestationInfo struct {
	NewUrl []byte
	NewKey []byte
}

func (*ScriptCall__RotateDualAttestationInfo) isScriptCall() {}

// (1) Rotate the public key stored in `account`'s `SharedEd25519PublicKey` resource to
// `new_public_key`
// (2) Rotate the authentication key using the capability stored in `account`'s
// `SharedEd25519PublicKey` to a new value derived from `new_public_key`
// Aborts if `account` does not have a `SharedEd25519PublicKey` resource.
// Aborts if the length of `new_public_key` is not 32.
type ScriptCall__RotateSharedEd25519PublicKey struct {
	PublicKey []byte
}

func (*ScriptCall__RotateSharedEd25519PublicKey) isScriptCall() {}

// Set validator's config and updates the config in the validator set.
// NewEpochEvent is emitted.
type ScriptCall__SetValidatorConfigAndReconfigure struct {
	ValidatorAccount libratypes.AccountAddress
	ConsensusPubkey []byte
	ValidatorNetworkIdentityPubkey []byte
	ValidatorNetworkAddress []byte
	FullnodesNetworkIdentityPubkey []byte
	FullnodesNetworkAddress []byte
}

func (*ScriptCall__SetValidatorConfigAndReconfigure) isScriptCall() {}

// Set validator's operator
type ScriptCall__SetValidatorOperator struct {
	OperatorName []byte
	OperatorAccount libratypes.AccountAddress
}

func (*ScriptCall__SetValidatorOperator) isScriptCall() {}

// Set validator operator as 'operator_account' of validator owner 'account' (via Admin Script).
// `operator_name` should match expected from operator account. This script also
// takes `sliding_nonce`, as a unique nonce for this operation. See `Sliding_nonce.move` for details.
type ScriptCall__SetValidatorOperatorWithNonceAdmin struct {
	SlidingNonce uint64
	OperatorName []byte
	OperatorAccount libratypes.AccountAddress
}

func (*ScriptCall__SetValidatorOperatorWithNonceAdmin) isScriptCall() {}

// Mint 'mint_amount' to 'designated_dealer_address' for 'tier_index' tier.
// Max valid tier index is 3 since there are max 4 tiers per DD.
// Sender should be treasury compliance account and receiver authorized DD.
// `sliding_nonce` is a unique nonce for operation, see sliding_nonce.move for details.
type ScriptCall__TieredMint struct {
	CoinType libratypes.TypeTag
	SlidingNonce uint64
	DesignatedDealerAddress libratypes.AccountAddress
	MintAmount uint64
	TierIndex uint64
}

func (*ScriptCall__TieredMint) isScriptCall() {}

// Unfreeze account `address`. Initiator must be authorized.
// `sliding_nonce` is a unique nonce for operation, see sliding_nonce.move for details.
type ScriptCall__UnfreezeAccount struct {
	SlidingNonce uint64
	ToUnfreezeAccount libratypes.AccountAddress
}

func (*ScriptCall__UnfreezeAccount) isScriptCall() {}

// Unmints `amount_lbr` LBR from the sending account into the constituent coins and deposits
// the resulting coins into the sending account.
type ScriptCall__UnmintLbr struct {
	AmountLbr uint64
}

func (*ScriptCall__UnmintLbr) isScriptCall() {}

// Update the dual attesation limit to `new_micro_lbr_limit`.
type ScriptCall__UpdateDualAttestationLimit struct {
	SlidingNonce uint64
	NewMicroLbrLimit uint64
}

func (*ScriptCall__UpdateDualAttestationLimit) isScriptCall() {}

// Update the on-chain exchange rate to LBR for the given `currency` to be given by
// `new_exchange_rate_numerator/new_exchange_rate_denominator`.
type ScriptCall__UpdateExchangeRate struct {
	Currency libratypes.TypeTag
	SlidingNonce uint64
	NewExchangeRateNumerator uint64
	NewExchangeRateDenominator uint64
}

func (*ScriptCall__UpdateExchangeRate) isScriptCall() {}

// Update Libra version.
// `sliding_nonce` is a unique nonce for operation, see sliding_nonce.move for details.
type ScriptCall__UpdateLibraVersion struct {
	SlidingNonce uint64
	Major uint64
}

func (*ScriptCall__UpdateLibraVersion) isScriptCall() {}

// Allows--true--or disallows--false--minting of `currency` based upon `allow_minting`.
type ScriptCall__UpdateMintingAbility struct {
	Currency libratypes.TypeTag
	AllowMinting bool
}

func (*ScriptCall__UpdateMintingAbility) isScriptCall() {}

// Build a Libra `Script` from a structured object `ScriptCall`.
func EncodeScript(call ScriptCall) libratypes.Script {
	switch call := call.(type) {
	case *ScriptCall__AddCurrencyToAccount:
		return EncodeAddCurrencyToAccountScript(call.Currency)
	case *ScriptCall__AddRecoveryRotationCapability:
		return EncodeAddRecoveryRotationCapabilityScript(call.RecoveryAddress)
	case *ScriptCall__AddToScriptAllowList:
		return EncodeAddToScriptAllowListScript(call.Hash, call.SlidingNonce)
	case *ScriptCall__AddValidatorAndReconfigure:
		return EncodeAddValidatorAndReconfigureScript(call.SlidingNonce, call.ValidatorName, call.ValidatorAddress)
	case *ScriptCall__Burn:
		return EncodeBurnScript(call.Token, call.SlidingNonce, call.PreburnAddress)
	case *ScriptCall__BurnTxnFees:
		return EncodeBurnTxnFeesScript(call.CoinType)
	case *ScriptCall__CancelBurn:
		return EncodeCancelBurnScript(call.Token, call.PreburnAddress)
	case *ScriptCall__CreateChildVaspAccount:
		return EncodeCreateChildVaspAccountScript(call.CoinType, call.ChildAddress, call.AuthKeyPrefix, call.AddAllCurrencies, call.ChildInitialBalance)
	case *ScriptCall__CreateDesignatedDealer:
		return EncodeCreateDesignatedDealerScript(call.Currency, call.SlidingNonce, call.Addr, call.AuthKeyPrefix, call.HumanName, call.AddAllCurrencies)
	case *ScriptCall__CreateParentVaspAccount:
		return EncodeCreateParentVaspAccountScript(call.CoinType, call.SlidingNonce, call.NewAccountAddress, call.AuthKeyPrefix, call.HumanName, call.AddAllCurrencies)
	case *ScriptCall__CreateRecoveryAddress:
		return EncodeCreateRecoveryAddressScript()
	case *ScriptCall__CreateValidatorAccount:
		return EncodeCreateValidatorAccountScript(call.SlidingNonce, call.NewAccountAddress, call.AuthKeyPrefix, call.HumanName)
	case *ScriptCall__CreateValidatorOperatorAccount:
		return EncodeCreateValidatorOperatorAccountScript(call.SlidingNonce, call.NewAccountAddress, call.AuthKeyPrefix, call.HumanName)
	case *ScriptCall__FreezeAccount:
		return EncodeFreezeAccountScript(call.SlidingNonce, call.ToFreezeAccount)
	case *ScriptCall__MintLbr:
		return EncodeMintLbrScript(call.AmountLbr)
	case *ScriptCall__PeerToPeerWithMetadata:
		return EncodePeerToPeerWithMetadataScript(call.Currency, call.Payee, call.Amount, call.Metadata, call.MetadataSignature)
	case *ScriptCall__Preburn:
		return EncodePreburnScript(call.Token, call.Amount)
	case *ScriptCall__PublishSharedEd25519PublicKey:
		return EncodePublishSharedEd25519PublicKeyScript(call.PublicKey)
	case *ScriptCall__RegisterValidatorConfig:
		return EncodeRegisterValidatorConfigScript(call.ValidatorAccount, call.ConsensusPubkey, call.ValidatorNetworkIdentityPubkey, call.ValidatorNetworkAddress, call.FullnodesNetworkIdentityPubkey, call.FullnodesNetworkAddress)
	case *ScriptCall__RemoveValidatorAndReconfigure:
		return EncodeRemoveValidatorAndReconfigureScript(call.SlidingNonce, call.ValidatorName, call.ValidatorAddress)
	case *ScriptCall__RotateAuthenticationKey:
		return EncodeRotateAuthenticationKeyScript(call.NewKey)
	case *ScriptCall__RotateAuthenticationKeyWithNonce:
		return EncodeRotateAuthenticationKeyWithNonceScript(call.SlidingNonce, call.NewKey)
	case *ScriptCall__RotateAuthenticationKeyWithNonceAdmin:
		return EncodeRotateAuthenticationKeyWithNonceAdminScript(call.SlidingNonce, call.NewKey)
	case *ScriptCall__RotateAuthenticationKeyWithRecoveryAddress:
		return EncodeRotateAuthenticationKeyWithRecoveryAddressScript(call.RecoveryAddress, call.ToRecover, call.NewKey)
	case *ScriptCall__RotateDualAttestationInfo:
		return EncodeRotateDualAttestationInfoScript(call.NewUrl, call.NewKey)
	case *ScriptCall__RotateSharedEd25519PublicKey:
		return EncodeRotateSharedEd25519PublicKeyScript(call.PublicKey)
	case *ScriptCall__SetValidatorConfigAndReconfigure:
		return EncodeSetValidatorConfigAndReconfigureScript(call.ValidatorAccount, call.ConsensusPubkey, call.ValidatorNetworkIdentityPubkey, call.ValidatorNetworkAddress, call.FullnodesNetworkIdentityPubkey, call.FullnodesNetworkAddress)
	case *ScriptCall__SetValidatorOperator:
		return EncodeSetValidatorOperatorScript(call.OperatorName, call.OperatorAccount)
	case *ScriptCall__SetValidatorOperatorWithNonceAdmin:
		return EncodeSetValidatorOperatorWithNonceAdminScript(call.SlidingNonce, call.OperatorName, call.OperatorAccount)
	case *ScriptCall__TieredMint:
		return EncodeTieredMintScript(call.CoinType, call.SlidingNonce, call.DesignatedDealerAddress, call.MintAmount, call.TierIndex)
	case *ScriptCall__UnfreezeAccount:
		return EncodeUnfreezeAccountScript(call.SlidingNonce, call.ToUnfreezeAccount)
	case *ScriptCall__UnmintLbr:
		return EncodeUnmintLbrScript(call.AmountLbr)
	case *ScriptCall__UpdateDualAttestationLimit:
		return EncodeUpdateDualAttestationLimitScript(call.SlidingNonce, call.NewMicroLbrLimit)
	case *ScriptCall__UpdateExchangeRate:
		return EncodeUpdateExchangeRateScript(call.Currency, call.SlidingNonce, call.NewExchangeRateNumerator, call.NewExchangeRateDenominator)
	case *ScriptCall__UpdateLibraVersion:
		return EncodeUpdateLibraVersionScript(call.SlidingNonce, call.Major)
	case *ScriptCall__UpdateMintingAbility:
		return EncodeUpdateMintingAbilityScript(call.Currency, call.AllowMinting)
	}
	panic("unreachable")
}

// Try to recognize a Libra `Script` and convert it into a structured object `ScriptCall`.
func DecodeScript(script *libratypes.Script) (ScriptCall, error) {
	if helper := script_decoder_map[string(script.Code)]; helper != nil {
		val, err := helper(script)
                return val, err
	} else {
		return nil, fmt.Errorf("Unknown script bytecode: %s", string(script.Code))
	}
}

// Add a `Currency` balance to `account`, which will enable `account` to send and receive
// `Libra<Currency>`.
// Aborts with NOT_A_CURRENCY if `Currency` is not an accepted currency type in the Libra system
// Aborts with `LibraAccount::ADD_EXISTING_CURRENCY` if the account already holds a balance in
// `Currency`.
func EncodeAddCurrencyToAccountScript(currency libratypes.TypeTag) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), add_currency_to_account_code...),
		TyArgs: []libratypes.TypeTag{currency},
		Args: []libratypes.TransactionArgument{},
	}
}

// Add the `KeyRotationCapability` for `to_recover_account` to the `RecoveryAddress` resource under `recovery_address`.
//
// ## Aborts
// * Aborts with `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` if `account` has already delegated its `KeyRotationCapability`.
// * Aborts with `RecoveryAddress:ENOT_A_RECOVERY_ADDRESS` if `recovery_address` does not have a `RecoveryAddress` resource.
// * Aborts with `RecoveryAddress::EINVALID_KEY_ROTATION_DELEGATION` if `to_recover_account` and `recovery_address` do not belong to the same VASP.
func EncodeAddRecoveryRotationCapabilityScript(recovery_address libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), add_recovery_rotation_capability_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{recovery_address}},
	}
}

// Append the `hash` to script hashes list allowed to be executed by the network.
func EncodeAddToScriptAllowListScript(hash []byte, sliding_nonce uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), add_to_script_allow_list_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U8Vector{hash}, &libratypes.TransactionArgument__U64{sliding_nonce}},
	}
}

// Add `new_validator` to the validator set.
// Fails if the `new_validator` address is already in the validator set
// or does not have a `ValidatorConfig` resource stored at the address.
// Emits a NewEpochEvent.
func EncodeAddValidatorAndReconfigureScript(sliding_nonce uint64, validator_name []byte, validator_address libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), add_validator_and_reconfigure_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__U8Vector{validator_name}, &libratypes.TransactionArgument__Address{validator_address}},
	}
}

// Permanently destroy the `Token`s stored in the oldest burn request under the `Preburn` resource.
// This will only succeed if `account` has a `MintCapability<Token>`, a `Preburn<Token>` resource
// exists under `preburn_address`, and there is a pending burn request.
// sliding_nonce is a unique nonce for operation, see sliding_nonce.move for details
func EncodeBurnScript(token libratypes.TypeTag, sliding_nonce uint64, preburn_address libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), burn_code...),
		TyArgs: []libratypes.TypeTag{token},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__Address{preburn_address}},
	}
}

// Burn transaction fees that have been collected in the given `currency`
// and relinquish to the association. The currency must be non-synthetic.
func EncodeBurnTxnFeesScript(coin_type libratypes.TypeTag) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), burn_txn_fees_code...),
		TyArgs: []libratypes.TypeTag{coin_type},
		Args: []libratypes.TransactionArgument{},
	}
}

// Cancel the oldest burn request from `preburn_address` and return the funds.
// Fails if the sender does not have a published `BurnCapability<Token>`.
func EncodeCancelBurnScript(token libratypes.TypeTag, preburn_address libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), cancel_burn_code...),
		TyArgs: []libratypes.TypeTag{token},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{preburn_address}},
	}
}

// Create a `ChildVASP` account for sender `parent_vasp` at `child_address` with a balance of
// `child_initial_balance` in `CoinType` and an initial authentication_key
// `auth_key_prefix | child_address`.
// If `add_all_currencies` is true, the child address will have a zero balance in all available
// currencies in the system.
// This account will a child of the transaction sender, which must be a ParentVASP.
//
// ## Aborts
// The transaction will abort:
//
// * If `parent_vasp` is not a parent vasp with error: `Roles::EINVALID_PARENT_ROLE`
// * If `child_address` already exists with error: `Roles::EROLE_ALREADY_ASSIGNED`
// * If `parent_vasp` already has 256 child accounts with error: `VASP::ETOO_MANY_CHILDREN`
// * If `CoinType` is not a registered currency with error: `LibraAccount::ENOT_A_CURRENCY`
// * If `parent_vasp`'s withdrawal capability has been extracted with error:  `LibraAccount::EWITHDRAWAL_CAPABILITY_ALREADY_EXTRACTED`
// * If `parent_vasp` doesn't hold `CoinType` and `child_initial_balance > 0` with error: `LibraAccount::EPAYER_DOESNT_HOLD_CURRENCY`
// * If `parent_vasp` doesn't at least `child_initial_balance` of `CoinType` in its account balance with error: `LibraAccount::EINSUFFICIENT_BALANCE`
func EncodeCreateChildVaspAccountScript(coin_type libratypes.TypeTag, child_address libratypes.AccountAddress, auth_key_prefix []byte, add_all_currencies bool, child_initial_balance uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), create_child_vasp_account_code...),
		TyArgs: []libratypes.TypeTag{coin_type},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{child_address}, &libratypes.TransactionArgument__U8Vector{auth_key_prefix}, &libratypes.TransactionArgument__Bool{add_all_currencies}, &libratypes.TransactionArgument__U64{child_initial_balance}},
	}
}

// Create an account with the DesignatedDealer role at `addr` with authentication key
// `auth_key_prefix` | `addr` and a 0 balance of type `Currency`. If `add_all_currencies` is true,
// 0 balances for all available currencies in the system will also be added. This can only be
// invoked by an account with the TreasuryCompliance role.
func EncodeCreateDesignatedDealerScript(currency libratypes.TypeTag, sliding_nonce uint64, addr libratypes.AccountAddress, auth_key_prefix []byte, human_name []byte, add_all_currencies bool) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), create_designated_dealer_code...),
		TyArgs: []libratypes.TypeTag{currency},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__Address{addr}, &libratypes.TransactionArgument__U8Vector{auth_key_prefix}, &libratypes.TransactionArgument__U8Vector{human_name}, &libratypes.TransactionArgument__Bool{add_all_currencies}},
	}
}

// Create an account with the ParentVASP role at `address` with authentication key
// `auth_key_prefix` | `new_account_address` and a 0 balance of type `currency`. If
// `add_all_currencies` is true, 0 balances for all available currencies in the system will
// also be added. This can only be invoked by an Association account.
// `sliding_nonce` is a unique nonce for operation, see sliding_nonce.move for details.
func EncodeCreateParentVaspAccountScript(coin_type libratypes.TypeTag, sliding_nonce uint64, new_account_address libratypes.AccountAddress, auth_key_prefix []byte, human_name []byte, add_all_currencies bool) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), create_parent_vasp_account_code...),
		TyArgs: []libratypes.TypeTag{coin_type},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__Address{new_account_address}, &libratypes.TransactionArgument__U8Vector{auth_key_prefix}, &libratypes.TransactionArgument__U8Vector{human_name}, &libratypes.TransactionArgument__Bool{add_all_currencies}},
	}
}

// Extract the `KeyRotationCapability` for `recovery_account` and publish it in a
// `RecoveryAddress` resource under  `account`.
// ## Aborts
// * Aborts with `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` if `account` has already delegated its `KeyRotationCapability`.
// * Aborts with `RecoveryAddress::ENOT_A_VASP` if `account` is not a ParentVASP or ChildVASP
func EncodeCreateRecoveryAddressScript() libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), create_recovery_address_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{},
	}
}

// Create a validator account at `new_validator_address` with `auth_key_prefix`and human_name.
func EncodeCreateValidatorAccountScript(sliding_nonce uint64, new_account_address libratypes.AccountAddress, auth_key_prefix []byte, human_name []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), create_validator_account_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__Address{new_account_address}, &libratypes.TransactionArgument__U8Vector{auth_key_prefix}, &libratypes.TransactionArgument__U8Vector{human_name}},
	}
}

// Create a validator operator account at `new_validator_address` with `auth_key_prefix`and human_name.
func EncodeCreateValidatorOperatorAccountScript(sliding_nonce uint64, new_account_address libratypes.AccountAddress, auth_key_prefix []byte, human_name []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), create_validator_operator_account_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__Address{new_account_address}, &libratypes.TransactionArgument__U8Vector{auth_key_prefix}, &libratypes.TransactionArgument__U8Vector{human_name}},
	}
}

// Freeze account `address`. Initiator must be authorized.
// `sliding_nonce` is a unique nonce for operation, see sliding_nonce.move for details.
func EncodeFreezeAccountScript(sliding_nonce uint64, to_freeze_account libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), freeze_account_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__Address{to_freeze_account}},
	}
}

// Mint `amount_lbr` LBR from the sending account's constituent coins and deposits the
// resulting LBR into the sending account.
func EncodeMintLbrScript(amount_lbr uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), mint_lbr_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{amount_lbr}},
	}
}

// Transfer `amount` coins of type `Currency` from `payer` to `payee` with (optional) associated
// `metadata` and an (optional) `metadata_signature` on the message
// `metadata` | `Signer::address_of(payer)` | `amount` | `DualAttestation::DOMAIN_SEPARATOR`.
// The `metadata` and `metadata_signature` parameters are only required if `amount` >=
// `DualAttestation::get_cur_microlibra_limit` LBR and `payer` and `payee` are distinct VASPs.
// However, a transaction sender can opt in to dual attestation even when it is not required (e.g., a DesignatedDealer -> VASP payment) by providing a non-empty `metadata_signature`.
// Standardized `metadata` LCS format can be found in `libra_types::transaction::metadata::Metadata`.
//
// ## Events
// When this script executes without aborting, it emits two events:
// `SentPaymentEvent { amount, currency_code = Currency, payee, metadata }`
// on `payer`'s `LibraAccount::sent_events` handle, and
//  `ReceivedPaymentEvent { amount, currency_code = Currency, payer, metadata }`
// on `payee`'s `LibraAccount::received_events` handle.
//
// ## Common Aborts
// These aborts can in occur in any payment.
// * Aborts with `LibraAccount::EINSUFFICIENT_BALANCE` if `amount` is greater than `payer`'s balance in `Currency`.
// * Aborts with `LibraAccount::ECOIN_DEPOSIT_IS_ZERO` if `amount` is zero.
// * Aborts with `LibraAccount::EPAYEE_DOES_NOT_EXIST` if no account exists at the address `payee`.
// * Aborts with `LibraAccount::EPAYEE_CANT_ACCEPT_CURRENCY_TYPE` if an account exists at `payee`, but it does not accept payments in `Currency`.
//
// ## Dual Attestation Aborts
// These aborts can occur in any payment subject to dual attestation.
// * Aborts with `DualAttestation::EMALFORMED_METADATA_SIGNATURE` if `metadata_signature`'s is not 64 bytes.
// * Aborts with `DualAttestation:EINVALID_METADATA_SIGNATURE` if `metadata_signature` does not verify on the message `metadata` | `payer` | `value` | `DOMAIN_SEPARATOR` using the `compliance_public_key` published in the `payee`'s `DualAttestation::Credential` resource.
//
// ## Other Aborts
// These aborts should only happen when `payer` or `payee` have account limit restrictions or
// have been frozen by Libra administrators.
// * Aborts with `LibraAccount::EWITHDRAWAL_EXCEEDS_LIMITS` if `payer` has exceeded their daily
// withdrawal limits.
// * Aborts with `LibraAccount::EDEPOSIT_EXCEEDS_LIMITS` if `payee` has exceeded their daily deposit limits.
// * Aborts with `LibraAccount::EACCOUNT_FROZEN` if `payer`'s account is frozen.
func EncodePeerToPeerWithMetadataScript(currency libratypes.TypeTag, payee libratypes.AccountAddress, amount uint64, metadata []byte, metadata_signature []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), peer_to_peer_with_metadata_code...),
		TyArgs: []libratypes.TypeTag{currency},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{payee}, &libratypes.TransactionArgument__U64{amount}, &libratypes.TransactionArgument__U8Vector{metadata}, &libratypes.TransactionArgument__U8Vector{metadata_signature}},
	}
}

// Preburn `amount` `Token`s from `account`.
// This will only succeed if `account` already has a published `Preburn<Token>` resource.
func EncodePreburnScript(token libratypes.TypeTag, amount uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), preburn_code...),
		TyArgs: []libratypes.TypeTag{token},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{amount}},
	}
}

// (1) Rotate the authentication key of the sender to `public_key`
// (2) Publish a resource containing a 32-byte ed25519 public key and the rotation capability
//     of the sender under the sender's address.
// Aborts if the sender already has a `SharedEd25519PublicKey` resource.
// Aborts if the length of `new_public_key` is not 32.
func EncodePublishSharedEd25519PublicKeyScript(public_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), publish_shared_ed25519_public_key_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U8Vector{public_key}},
	}
}

// Set validator's config locally.
// Does not emit NewEpochEvent, the config is NOT changed in the validator set.
func EncodeRegisterValidatorConfigScript(validator_account libratypes.AccountAddress, consensus_pubkey []byte, validator_network_identity_pubkey []byte, validator_network_address []byte, fullnodes_network_identity_pubkey []byte, fullnodes_network_address []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), register_validator_config_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{validator_account}, &libratypes.TransactionArgument__U8Vector{consensus_pubkey}, &libratypes.TransactionArgument__U8Vector{validator_network_identity_pubkey}, &libratypes.TransactionArgument__U8Vector{validator_network_address}, &libratypes.TransactionArgument__U8Vector{fullnodes_network_identity_pubkey}, &libratypes.TransactionArgument__U8Vector{fullnodes_network_address}},
	}
}

// Removes a validator from the validator set.
// Fails if the validator_address is not in the validator set.
// Emits a NewEpochEvent.
func EncodeRemoveValidatorAndReconfigureScript(sliding_nonce uint64, validator_name []byte, validator_address libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), remove_validator_and_reconfigure_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__U8Vector{validator_name}, &libratypes.TransactionArgument__Address{validator_address}},
	}
}

// Rotate the sender's authentication key to `new_key`.
// `new_key` should be a 256 bit sha3 hash of an ed25519 public key.
// * Aborts with `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` if the `KeyRotationCapability` for `account` has already been extracted.
// * Aborts with `0` if the key rotation capability held by the account doesn't match the sender's address.
// * Aborts with `LibraAccount::EMALFORMED_AUTHENTICATION_KEY` if the length of `new_key` != 32.
func EncodeRotateAuthenticationKeyScript(new_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), rotate_authentication_key_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U8Vector{new_key}},
	}
}

// Rotate `account`'s authentication key to `new_key`.
// `new_key` should be a 256 bit sha3 hash of an ed25519 public key. This script also takes
// `sliding_nonce`, as a unique nonce for this operation. See sliding_nonce.move for details.
func EncodeRotateAuthenticationKeyWithNonceScript(sliding_nonce uint64, new_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), rotate_authentication_key_with_nonce_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__U8Vector{new_key}},
	}
}

// Rotate `account`'s authentication key to `new_key`.
// `new_key` should be a 256 bit sha3 hash of an ed25519 public key. This script also takes
// `sliding_nonce`, as a unique nonce for this operation. See sliding_nonce.move for details.
func EncodeRotateAuthenticationKeyWithNonceAdminScript(sliding_nonce uint64, new_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), rotate_authentication_key_with_nonce_admin_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__U8Vector{new_key}},
	}
}

// Rotate the authentication key of `account` to `new_key` using the `KeyRotationCapability`
// stored under `recovery_address`.
//
// ## Aborts
// * Aborts with `RecoveryAddress::ENOT_A_RECOVERY_ADDRESS` if `recovery_address` does not have a `RecoveryAddress` resource
// * Aborts with `RecoveryAddress::ECANNOT_ROTATE_KEY` if `account` is not `recovery_address` or `to_recover`.
// * Aborts with `LibraAccount::EMALFORMED_AUTHENTICATION_KEY` if `new_key` is not 32 bytes.
// * Aborts with `RecoveryAddress::ECANNOT_ROTATE_KEY` if `account` has not delegated its `KeyRotationCapability` to `recovery_address`.
func EncodeRotateAuthenticationKeyWithRecoveryAddressScript(recovery_address libratypes.AccountAddress, to_recover libratypes.AccountAddress, new_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), rotate_authentication_key_with_recovery_address_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{recovery_address}, &libratypes.TransactionArgument__Address{to_recover}, &libratypes.TransactionArgument__U8Vector{new_key}},
	}
}

// Rotate `account`'s base URL to `new_url` and its compliance public key to `new_key`.
// Aborts if `account` is not a ParentVASP or DesignatedDealer
// Aborts if `new_key` is not a well-formed public key
func EncodeRotateDualAttestationInfoScript(new_url []byte, new_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), rotate_dual_attestation_info_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U8Vector{new_url}, &libratypes.TransactionArgument__U8Vector{new_key}},
	}
}

// (1) Rotate the public key stored in `account`'s `SharedEd25519PublicKey` resource to
// `new_public_key`
// (2) Rotate the authentication key using the capability stored in `account`'s
// `SharedEd25519PublicKey` to a new value derived from `new_public_key`
// Aborts if `account` does not have a `SharedEd25519PublicKey` resource.
// Aborts if the length of `new_public_key` is not 32.
func EncodeRotateSharedEd25519PublicKeyScript(public_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), rotate_shared_ed25519_public_key_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U8Vector{public_key}},
	}
}

// Set validator's config and updates the config in the validator set.
// NewEpochEvent is emitted.
func EncodeSetValidatorConfigAndReconfigureScript(validator_account libratypes.AccountAddress, consensus_pubkey []byte, validator_network_identity_pubkey []byte, validator_network_address []byte, fullnodes_network_identity_pubkey []byte, fullnodes_network_address []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), set_validator_config_and_reconfigure_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{validator_account}, &libratypes.TransactionArgument__U8Vector{consensus_pubkey}, &libratypes.TransactionArgument__U8Vector{validator_network_identity_pubkey}, &libratypes.TransactionArgument__U8Vector{validator_network_address}, &libratypes.TransactionArgument__U8Vector{fullnodes_network_identity_pubkey}, &libratypes.TransactionArgument__U8Vector{fullnodes_network_address}},
	}
}

// Set validator's operator
func EncodeSetValidatorOperatorScript(operator_name []byte, operator_account libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), set_validator_operator_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U8Vector{operator_name}, &libratypes.TransactionArgument__Address{operator_account}},
	}
}

// Set validator operator as 'operator_account' of validator owner 'account' (via Admin Script).
// `operator_name` should match expected from operator account. This script also
// takes `sliding_nonce`, as a unique nonce for this operation. See `Sliding_nonce.move` for details.
func EncodeSetValidatorOperatorWithNonceAdminScript(sliding_nonce uint64, operator_name []byte, operator_account libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), set_validator_operator_with_nonce_admin_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__U8Vector{operator_name}, &libratypes.TransactionArgument__Address{operator_account}},
	}
}

// Mint 'mint_amount' to 'designated_dealer_address' for 'tier_index' tier.
// Max valid tier index is 3 since there are max 4 tiers per DD.
// Sender should be treasury compliance account and receiver authorized DD.
// `sliding_nonce` is a unique nonce for operation, see sliding_nonce.move for details.
func EncodeTieredMintScript(coin_type libratypes.TypeTag, sliding_nonce uint64, designated_dealer_address libratypes.AccountAddress, mint_amount uint64, tier_index uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), tiered_mint_code...),
		TyArgs: []libratypes.TypeTag{coin_type},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__Address{designated_dealer_address}, &libratypes.TransactionArgument__U64{mint_amount}, &libratypes.TransactionArgument__U64{tier_index}},
	}
}

// Unfreeze account `address`. Initiator must be authorized.
// `sliding_nonce` is a unique nonce for operation, see sliding_nonce.move for details.
func EncodeUnfreezeAccountScript(sliding_nonce uint64, to_unfreeze_account libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), unfreeze_account_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__Address{to_unfreeze_account}},
	}
}

// Unmints `amount_lbr` LBR from the sending account into the constituent coins and deposits
// the resulting coins into the sending account.
func EncodeUnmintLbrScript(amount_lbr uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), unmint_lbr_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{amount_lbr}},
	}
}

// Update the dual attesation limit to `new_micro_lbr_limit`.
func EncodeUpdateDualAttestationLimitScript(sliding_nonce uint64, new_micro_lbr_limit uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), update_dual_attestation_limit_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__U64{new_micro_lbr_limit}},
	}
}

// Update the on-chain exchange rate to LBR for the given `currency` to be given by
// `new_exchange_rate_numerator/new_exchange_rate_denominator`.
func EncodeUpdateExchangeRateScript(currency libratypes.TypeTag, sliding_nonce uint64, new_exchange_rate_numerator uint64, new_exchange_rate_denominator uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), update_exchange_rate_code...),
		TyArgs: []libratypes.TypeTag{currency},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__U64{new_exchange_rate_numerator}, &libratypes.TransactionArgument__U64{new_exchange_rate_denominator}},
	}
}

// Update Libra version.
// `sliding_nonce` is a unique nonce for operation, see sliding_nonce.move for details.
func EncodeUpdateLibraVersionScript(sliding_nonce uint64, major uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), update_libra_version_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__U64{sliding_nonce}, &libratypes.TransactionArgument__U64{major}},
	}
}

// Allows--true--or disallows--false--minting of `currency` based upon `allow_minting`.
func EncodeUpdateMintingAbilityScript(currency libratypes.TypeTag, allow_minting bool) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), update_minting_ability_code...),
		TyArgs: []libratypes.TypeTag{currency},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Bool{allow_minting}},
	}
}

func decode_add_currency_to_account_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 1 { return nil, fmt.Errorf("Was expecting 1 type arguments") }
	if len(script.Args) < 0 { return nil, fmt.Errorf("Was expecting 0 regular arguments") }
	var call ScriptCall__AddCurrencyToAccount
	call.Currency = script.TyArgs[0]
	return &call, nil
}

func decode_add_recovery_rotation_capability_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 1 { return nil, fmt.Errorf("Was expecting 1 regular arguments") }
	var call ScriptCall__AddRecoveryRotationCapability
	if val, err := decode_address_argument(script.Args[0]); err == nil {
		call.RecoveryAddress = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_add_to_script_allow_list_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 2 { return nil, fmt.Errorf("Was expecting 2 regular arguments") }
	var call ScriptCall__AddToScriptAllowList
	if val, err := decode_u8vector_argument(script.Args[0]); err == nil {
		call.Hash = val
	} else {
		return nil, err
	}

	if val, err := decode_u64_argument(script.Args[1]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_add_validator_and_reconfigure_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 3 { return nil, fmt.Errorf("Was expecting 3 regular arguments") }
	var call ScriptCall__AddValidatorAndReconfigure
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[1]); err == nil {
		call.ValidatorName = val
	} else {
		return nil, err
	}

	if val, err := decode_address_argument(script.Args[2]); err == nil {
		call.ValidatorAddress = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_burn_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 1 { return nil, fmt.Errorf("Was expecting 1 type arguments") }
	if len(script.Args) < 2 { return nil, fmt.Errorf("Was expecting 2 regular arguments") }
	var call ScriptCall__Burn
	call.Token = script.TyArgs[0]
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_address_argument(script.Args[1]); err == nil {
		call.PreburnAddress = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_burn_txn_fees_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 1 { return nil, fmt.Errorf("Was expecting 1 type arguments") }
	if len(script.Args) < 0 { return nil, fmt.Errorf("Was expecting 0 regular arguments") }
	var call ScriptCall__BurnTxnFees
	call.CoinType = script.TyArgs[0]
	return &call, nil
}

func decode_cancel_burn_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 1 { return nil, fmt.Errorf("Was expecting 1 type arguments") }
	if len(script.Args) < 1 { return nil, fmt.Errorf("Was expecting 1 regular arguments") }
	var call ScriptCall__CancelBurn
	call.Token = script.TyArgs[0]
	if val, err := decode_address_argument(script.Args[0]); err == nil {
		call.PreburnAddress = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_create_child_vasp_account_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 1 { return nil, fmt.Errorf("Was expecting 1 type arguments") }
	if len(script.Args) < 4 { return nil, fmt.Errorf("Was expecting 4 regular arguments") }
	var call ScriptCall__CreateChildVaspAccount
	call.CoinType = script.TyArgs[0]
	if val, err := decode_address_argument(script.Args[0]); err == nil {
		call.ChildAddress = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[1]); err == nil {
		call.AuthKeyPrefix = val
	} else {
		return nil, err
	}

	if val, err := decode_bool_argument(script.Args[2]); err == nil {
		call.AddAllCurrencies = val
	} else {
		return nil, err
	}

	if val, err := decode_u64_argument(script.Args[3]); err == nil {
		call.ChildInitialBalance = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_create_designated_dealer_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 1 { return nil, fmt.Errorf("Was expecting 1 type arguments") }
	if len(script.Args) < 5 { return nil, fmt.Errorf("Was expecting 5 regular arguments") }
	var call ScriptCall__CreateDesignatedDealer
	call.Currency = script.TyArgs[0]
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_address_argument(script.Args[1]); err == nil {
		call.Addr = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[2]); err == nil {
		call.AuthKeyPrefix = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[3]); err == nil {
		call.HumanName = val
	} else {
		return nil, err
	}

	if val, err := decode_bool_argument(script.Args[4]); err == nil {
		call.AddAllCurrencies = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_create_parent_vasp_account_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 1 { return nil, fmt.Errorf("Was expecting 1 type arguments") }
	if len(script.Args) < 5 { return nil, fmt.Errorf("Was expecting 5 regular arguments") }
	var call ScriptCall__CreateParentVaspAccount
	call.CoinType = script.TyArgs[0]
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_address_argument(script.Args[1]); err == nil {
		call.NewAccountAddress = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[2]); err == nil {
		call.AuthKeyPrefix = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[3]); err == nil {
		call.HumanName = val
	} else {
		return nil, err
	}

	if val, err := decode_bool_argument(script.Args[4]); err == nil {
		call.AddAllCurrencies = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_create_recovery_address_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 0 { return nil, fmt.Errorf("Was expecting 0 regular arguments") }
	var call ScriptCall__CreateRecoveryAddress
	return &call, nil
}

func decode_create_validator_account_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 4 { return nil, fmt.Errorf("Was expecting 4 regular arguments") }
	var call ScriptCall__CreateValidatorAccount
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_address_argument(script.Args[1]); err == nil {
		call.NewAccountAddress = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[2]); err == nil {
		call.AuthKeyPrefix = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[3]); err == nil {
		call.HumanName = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_create_validator_operator_account_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 4 { return nil, fmt.Errorf("Was expecting 4 regular arguments") }
	var call ScriptCall__CreateValidatorOperatorAccount
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_address_argument(script.Args[1]); err == nil {
		call.NewAccountAddress = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[2]); err == nil {
		call.AuthKeyPrefix = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[3]); err == nil {
		call.HumanName = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_freeze_account_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 2 { return nil, fmt.Errorf("Was expecting 2 regular arguments") }
	var call ScriptCall__FreezeAccount
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_address_argument(script.Args[1]); err == nil {
		call.ToFreezeAccount = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_mint_lbr_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 1 { return nil, fmt.Errorf("Was expecting 1 regular arguments") }
	var call ScriptCall__MintLbr
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.AmountLbr = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_peer_to_peer_with_metadata_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 1 { return nil, fmt.Errorf("Was expecting 1 type arguments") }
	if len(script.Args) < 4 { return nil, fmt.Errorf("Was expecting 4 regular arguments") }
	var call ScriptCall__PeerToPeerWithMetadata
	call.Currency = script.TyArgs[0]
	if val, err := decode_address_argument(script.Args[0]); err == nil {
		call.Payee = val
	} else {
		return nil, err
	}

	if val, err := decode_u64_argument(script.Args[1]); err == nil {
		call.Amount = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[2]); err == nil {
		call.Metadata = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[3]); err == nil {
		call.MetadataSignature = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_preburn_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 1 { return nil, fmt.Errorf("Was expecting 1 type arguments") }
	if len(script.Args) < 1 { return nil, fmt.Errorf("Was expecting 1 regular arguments") }
	var call ScriptCall__Preburn
	call.Token = script.TyArgs[0]
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.Amount = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_publish_shared_ed25519_public_key_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 1 { return nil, fmt.Errorf("Was expecting 1 regular arguments") }
	var call ScriptCall__PublishSharedEd25519PublicKey
	if val, err := decode_u8vector_argument(script.Args[0]); err == nil {
		call.PublicKey = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_register_validator_config_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 6 { return nil, fmt.Errorf("Was expecting 6 regular arguments") }
	var call ScriptCall__RegisterValidatorConfig
	if val, err := decode_address_argument(script.Args[0]); err == nil {
		call.ValidatorAccount = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[1]); err == nil {
		call.ConsensusPubkey = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[2]); err == nil {
		call.ValidatorNetworkIdentityPubkey = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[3]); err == nil {
		call.ValidatorNetworkAddress = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[4]); err == nil {
		call.FullnodesNetworkIdentityPubkey = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[5]); err == nil {
		call.FullnodesNetworkAddress = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_remove_validator_and_reconfigure_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 3 { return nil, fmt.Errorf("Was expecting 3 regular arguments") }
	var call ScriptCall__RemoveValidatorAndReconfigure
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[1]); err == nil {
		call.ValidatorName = val
	} else {
		return nil, err
	}

	if val, err := decode_address_argument(script.Args[2]); err == nil {
		call.ValidatorAddress = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_rotate_authentication_key_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 1 { return nil, fmt.Errorf("Was expecting 1 regular arguments") }
	var call ScriptCall__RotateAuthenticationKey
	if val, err := decode_u8vector_argument(script.Args[0]); err == nil {
		call.NewKey = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_rotate_authentication_key_with_nonce_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 2 { return nil, fmt.Errorf("Was expecting 2 regular arguments") }
	var call ScriptCall__RotateAuthenticationKeyWithNonce
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[1]); err == nil {
		call.NewKey = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_rotate_authentication_key_with_nonce_admin_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 2 { return nil, fmt.Errorf("Was expecting 2 regular arguments") }
	var call ScriptCall__RotateAuthenticationKeyWithNonceAdmin
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[1]); err == nil {
		call.NewKey = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_rotate_authentication_key_with_recovery_address_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 3 { return nil, fmt.Errorf("Was expecting 3 regular arguments") }
	var call ScriptCall__RotateAuthenticationKeyWithRecoveryAddress
	if val, err := decode_address_argument(script.Args[0]); err == nil {
		call.RecoveryAddress = val
	} else {
		return nil, err
	}

	if val, err := decode_address_argument(script.Args[1]); err == nil {
		call.ToRecover = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[2]); err == nil {
		call.NewKey = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_rotate_dual_attestation_info_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 2 { return nil, fmt.Errorf("Was expecting 2 regular arguments") }
	var call ScriptCall__RotateDualAttestationInfo
	if val, err := decode_u8vector_argument(script.Args[0]); err == nil {
		call.NewUrl = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[1]); err == nil {
		call.NewKey = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_rotate_shared_ed25519_public_key_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 1 { return nil, fmt.Errorf("Was expecting 1 regular arguments") }
	var call ScriptCall__RotateSharedEd25519PublicKey
	if val, err := decode_u8vector_argument(script.Args[0]); err == nil {
		call.PublicKey = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_set_validator_config_and_reconfigure_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 6 { return nil, fmt.Errorf("Was expecting 6 regular arguments") }
	var call ScriptCall__SetValidatorConfigAndReconfigure
	if val, err := decode_address_argument(script.Args[0]); err == nil {
		call.ValidatorAccount = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[1]); err == nil {
		call.ConsensusPubkey = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[2]); err == nil {
		call.ValidatorNetworkIdentityPubkey = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[3]); err == nil {
		call.ValidatorNetworkAddress = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[4]); err == nil {
		call.FullnodesNetworkIdentityPubkey = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[5]); err == nil {
		call.FullnodesNetworkAddress = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_set_validator_operator_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 2 { return nil, fmt.Errorf("Was expecting 2 regular arguments") }
	var call ScriptCall__SetValidatorOperator
	if val, err := decode_u8vector_argument(script.Args[0]); err == nil {
		call.OperatorName = val
	} else {
		return nil, err
	}

	if val, err := decode_address_argument(script.Args[1]); err == nil {
		call.OperatorAccount = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_set_validator_operator_with_nonce_admin_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 3 { return nil, fmt.Errorf("Was expecting 3 regular arguments") }
	var call ScriptCall__SetValidatorOperatorWithNonceAdmin
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[1]); err == nil {
		call.OperatorName = val
	} else {
		return nil, err
	}

	if val, err := decode_address_argument(script.Args[2]); err == nil {
		call.OperatorAccount = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_tiered_mint_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 1 { return nil, fmt.Errorf("Was expecting 1 type arguments") }
	if len(script.Args) < 4 { return nil, fmt.Errorf("Was expecting 4 regular arguments") }
	var call ScriptCall__TieredMint
	call.CoinType = script.TyArgs[0]
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_address_argument(script.Args[1]); err == nil {
		call.DesignatedDealerAddress = val
	} else {
		return nil, err
	}

	if val, err := decode_u64_argument(script.Args[2]); err == nil {
		call.MintAmount = val
	} else {
		return nil, err
	}

	if val, err := decode_u64_argument(script.Args[3]); err == nil {
		call.TierIndex = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_unfreeze_account_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 2 { return nil, fmt.Errorf("Was expecting 2 regular arguments") }
	var call ScriptCall__UnfreezeAccount
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_address_argument(script.Args[1]); err == nil {
		call.ToUnfreezeAccount = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_unmint_lbr_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 1 { return nil, fmt.Errorf("Was expecting 1 regular arguments") }
	var call ScriptCall__UnmintLbr
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.AmountLbr = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_update_dual_attestation_limit_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 2 { return nil, fmt.Errorf("Was expecting 2 regular arguments") }
	var call ScriptCall__UpdateDualAttestationLimit
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_u64_argument(script.Args[1]); err == nil {
		call.NewMicroLbrLimit = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_update_exchange_rate_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 1 { return nil, fmt.Errorf("Was expecting 1 type arguments") }
	if len(script.Args) < 3 { return nil, fmt.Errorf("Was expecting 3 regular arguments") }
	var call ScriptCall__UpdateExchangeRate
	call.Currency = script.TyArgs[0]
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_u64_argument(script.Args[1]); err == nil {
		call.NewExchangeRateNumerator = val
	} else {
		return nil, err
	}

	if val, err := decode_u64_argument(script.Args[2]); err == nil {
		call.NewExchangeRateDenominator = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_update_libra_version_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 0 { return nil, fmt.Errorf("Was expecting 0 type arguments") }
	if len(script.Args) < 2 { return nil, fmt.Errorf("Was expecting 2 regular arguments") }
	var call ScriptCall__UpdateLibraVersion
	if val, err := decode_u64_argument(script.Args[0]); err == nil {
		call.SlidingNonce = val
	} else {
		return nil, err
	}

	if val, err := decode_u64_argument(script.Args[1]); err == nil {
		call.Major = val
	} else {
		return nil, err
	}

	return &call, nil
}

func decode_update_minting_ability_script(script *libratypes.Script) (ScriptCall, error) {
	if len(script.TyArgs) < 1 { return nil, fmt.Errorf("Was expecting 1 type arguments") }
	if len(script.Args) < 1 { return nil, fmt.Errorf("Was expecting 1 regular arguments") }
	var call ScriptCall__UpdateMintingAbility
	call.Currency = script.TyArgs[0]
	if val, err := decode_bool_argument(script.Args[0]); err == nil {
		call.AllowMinting = val
	} else {
		return nil, err
	}

	return &call, nil
}

var add_currency_to_account_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 2, 3, 2, 6, 4, 8, 2, 5, 10, 7, 7, 17, 26, 8, 43, 16, 0, 0, 0, 1, 0, 1, 1, 1, 0, 2, 1, 6, 12, 0, 1, 9, 0, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 12, 97, 100, 100, 95, 99, 117, 114, 114, 101, 110, 99, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 3, 11, 0, 56, 0, 2};

var add_recovery_rotation_capability_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 4, 2, 4, 4, 3, 8, 10, 5, 18, 15, 7, 33, 107, 8, 140, 1, 16, 0, 0, 0, 1, 0, 2, 1, 0, 0, 3, 0, 1, 0, 1, 4, 2, 3, 0, 1, 6, 12, 1, 8, 0, 2, 8, 0, 5, 0, 2, 6, 12, 5, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 15, 82, 101, 99, 111, 118, 101, 114, 121, 65, 100, 100, 114, 101, 115, 115, 21, 75, 101, 121, 82, 111, 116, 97, 116, 105, 111, 110, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 101, 120, 116, 114, 97, 99, 116, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 23, 97, 100, 100, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 4, 3, 5, 11, 0, 17, 0, 10, 1, 17, 1, 2};

var add_to_script_allow_list_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 4, 3, 4, 10, 5, 14, 16, 7, 30, 93, 8, 123, 16, 0, 0, 0, 1, 0, 2, 0, 1, 0, 1, 3, 2, 1, 0, 2, 6, 12, 10, 2, 0, 2, 6, 12, 3, 3, 6, 12, 10, 2, 3, 32, 76, 105, 98, 114, 97, 84, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 80, 117, 98, 108, 105, 115, 104, 105, 110, 103, 79, 112, 116, 105, 111, 110, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 24, 97, 100, 100, 95, 116, 111, 95, 115, 99, 114, 105, 112, 116, 95, 97, 108, 108, 111, 119, 95, 108, 105, 115, 116, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 3, 1, 7, 10, 0, 10, 2, 17, 1, 11, 0, 11, 1, 17, 0, 2};

var add_validator_and_reconfigure_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 6, 3, 6, 15, 5, 21, 24, 7, 45, 92, 8, 137, 1, 16, 0, 0, 0, 1, 0, 2, 1, 3, 0, 1, 0, 2, 4, 2, 3, 0, 0, 5, 4, 1, 0, 2, 6, 12, 3, 0, 1, 5, 1, 10, 2, 2, 6, 12, 5, 4, 6, 12, 3, 10, 2, 5, 2, 1, 3, 11, 76, 105, 98, 114, 97, 83, 121, 115, 116, 101, 109, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 15, 86, 97, 108, 105, 100, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 14, 103, 101, 116, 95, 104, 117, 109, 97, 110, 95, 110, 97, 109, 101, 13, 97, 100, 100, 95, 118, 97, 108, 105, 100, 97, 116, 111, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 5, 6, 18, 10, 0, 10, 1, 17, 0, 10, 3, 17, 1, 11, 2, 33, 12, 4, 11, 4, 3, 14, 11, 0, 1, 6, 0, 0, 0, 0, 0, 0, 0, 0, 39, 11, 0, 10, 3, 17, 2, 2};

var burn_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 4, 3, 4, 11, 4, 15, 2, 5, 17, 17, 7, 34, 46, 8, 80, 16, 0, 0, 0, 1, 1, 2, 0, 1, 0, 0, 3, 2, 1, 1, 1, 1, 4, 2, 6, 12, 3, 0, 2, 6, 12, 5, 3, 6, 12, 3, 5, 1, 9, 0, 5, 76, 105, 98, 114, 97, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 4, 98, 117, 114, 110, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 3, 1, 7, 10, 0, 10, 1, 17, 0, 11, 0, 10, 2, 56, 0, 2};

var burn_txn_fees_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 2, 3, 2, 6, 4, 8, 2, 5, 10, 7, 7, 17, 25, 8, 42, 16, 0, 0, 0, 1, 0, 1, 1, 1, 0, 2, 1, 6, 12, 0, 1, 9, 0, 14, 84, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 70, 101, 101, 9, 98, 117, 114, 110, 95, 102, 101, 101, 115, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 3, 11, 0, 56, 0, 2};

var cancel_burn_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 2, 3, 2, 6, 4, 8, 2, 5, 10, 8, 7, 18, 25, 8, 43, 16, 0, 0, 0, 1, 0, 1, 1, 1, 0, 2, 2, 6, 12, 5, 0, 1, 9, 0, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 11, 99, 97, 110, 99, 101, 108, 95, 98, 117, 114, 110, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 4, 11, 0, 10, 1, 56, 0, 2};

var create_child_vasp_account_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 8, 1, 0, 2, 2, 2, 4, 3, 6, 22, 4, 28, 4, 5, 32, 35, 7, 67, 123, 8, 190, 1, 16, 6, 206, 1, 4, 0, 0, 0, 1, 1, 0, 0, 2, 0, 1, 1, 1, 0, 3, 2, 3, 0, 0, 4, 4, 1, 1, 1, 0, 5, 3, 1, 0, 0, 6, 2, 6, 4, 6, 12, 5, 10, 2, 1, 0, 1, 6, 12, 1, 8, 0, 5, 6, 8, 0, 5, 3, 10, 2, 10, 2, 5, 6, 12, 5, 10, 2, 1, 3, 1, 9, 0, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 18, 87, 105, 116, 104, 100, 114, 97, 119, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 25, 99, 114, 101, 97, 116, 101, 95, 99, 104, 105, 108, 100, 95, 118, 97, 115, 112, 95, 97, 99, 99, 111, 117, 110, 116, 27, 101, 120, 116, 114, 97, 99, 116, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 8, 112, 97, 121, 95, 102, 114, 111, 109, 27, 114, 101, 115, 116, 111, 114, 101, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 10, 2, 1, 0, 1, 1, 5, 3, 25, 10, 0, 10, 1, 11, 2, 10, 3, 56, 0, 10, 4, 6, 0, 0, 0, 0, 0, 0, 0, 0, 36, 3, 10, 5, 22, 11, 0, 17, 1, 12, 5, 14, 5, 10, 1, 10, 4, 7, 0, 7, 0, 56, 1, 11, 5, 17, 3, 5, 24, 11, 0, 1, 2};

var create_designated_dealer_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 4, 3, 4, 11, 4, 15, 2, 5, 17, 27, 7, 44, 73, 8, 117, 16, 0, 0, 0, 1, 1, 2, 0, 1, 0, 0, 3, 2, 1, 1, 1, 1, 4, 2, 6, 12, 3, 0, 5, 6, 12, 5, 10, 2, 10, 2, 1, 6, 6, 12, 3, 5, 10, 2, 10, 2, 1, 1, 9, 0, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 24, 99, 114, 101, 97, 116, 101, 95, 100, 101, 115, 105, 103, 110, 97, 116, 101, 100, 95, 100, 101, 97, 108, 101, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 3, 1, 10, 10, 0, 10, 1, 17, 0, 11, 0, 10, 2, 11, 3, 11, 4, 10, 5, 56, 0, 2};

var create_parent_vasp_account_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 4, 3, 4, 11, 4, 15, 2, 5, 17, 27, 7, 44, 75, 8, 119, 16, 0, 0, 0, 1, 1, 2, 0, 1, 0, 0, 3, 2, 1, 1, 1, 1, 4, 2, 6, 12, 3, 0, 5, 6, 12, 5, 10, 2, 10, 2, 1, 6, 6, 12, 3, 5, 10, 2, 10, 2, 1, 1, 9, 0, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 26, 99, 114, 101, 97, 116, 101, 95, 112, 97, 114, 101, 110, 116, 95, 118, 97, 115, 112, 95, 97, 99, 99, 111, 117, 110, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 3, 1, 10, 10, 0, 10, 1, 17, 0, 11, 0, 10, 2, 11, 3, 11, 4, 10, 5, 56, 0, 2};

var create_recovery_address_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 4, 2, 4, 4, 3, 8, 10, 5, 18, 12, 7, 30, 91, 8, 121, 16, 0, 0, 0, 1, 0, 2, 1, 0, 0, 3, 0, 1, 0, 1, 4, 2, 3, 0, 1, 6, 12, 1, 8, 0, 2, 6, 12, 8, 0, 0, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 15, 82, 101, 99, 111, 118, 101, 114, 121, 65, 100, 100, 114, 101, 115, 115, 21, 75, 101, 121, 82, 111, 116, 97, 116, 105, 111, 110, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 101, 120, 116, 114, 97, 99, 116, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 7, 112, 117, 98, 108, 105, 115, 104, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 3, 5, 10, 0, 11, 0, 17, 0, 17, 1, 2};

var create_validator_account_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 4, 3, 4, 10, 5, 14, 22, 7, 36, 73, 8, 109, 16, 0, 0, 0, 1, 1, 2, 0, 1, 0, 0, 3, 2, 1, 0, 2, 6, 12, 3, 0, 4, 6, 12, 5, 10, 2, 10, 2, 5, 6, 12, 3, 5, 10, 2, 10, 2, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 24, 99, 114, 101, 97, 116, 101, 95, 118, 97, 108, 105, 100, 97, 116, 111, 114, 95, 97, 99, 99, 111, 117, 110, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 3, 1, 9, 10, 0, 10, 1, 17, 0, 11, 0, 10, 2, 11, 3, 11, 4, 17, 1, 2};

var create_validator_operator_account_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 4, 3, 4, 10, 5, 14, 22, 7, 36, 82, 8, 118, 16, 0, 0, 0, 1, 1, 2, 0, 1, 0, 0, 3, 2, 1, 0, 2, 6, 12, 3, 0, 4, 6, 12, 5, 10, 2, 10, 2, 5, 6, 12, 3, 5, 10, 2, 10, 2, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 33, 99, 114, 101, 97, 116, 101, 95, 118, 97, 108, 105, 100, 97, 116, 111, 114, 95, 111, 112, 101, 114, 97, 116, 111, 114, 95, 97, 99, 99, 111, 117, 110, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 3, 1, 9, 10, 0, 10, 1, 17, 0, 11, 0, 10, 2, 11, 3, 11, 4, 17, 1, 2};

var freeze_account_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 4, 3, 4, 10, 5, 14, 14, 7, 28, 66, 8, 94, 16, 0, 0, 0, 1, 0, 2, 0, 1, 0, 1, 3, 2, 1, 0, 2, 6, 12, 5, 0, 2, 6, 12, 3, 3, 6, 12, 3, 5, 15, 65, 99, 99, 111, 117, 110, 116, 70, 114, 101, 101, 122, 105, 110, 103, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 14, 102, 114, 101, 101, 122, 101, 95, 97, 99, 99, 111, 117, 110, 116, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 3, 1, 7, 10, 0, 10, 1, 17, 1, 11, 0, 10, 2, 17, 0, 2};

var mint_lbr_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 2, 2, 2, 4, 3, 6, 15, 5, 21, 16, 7, 37, 99, 8, 136, 1, 16, 0, 0, 0, 1, 1, 0, 0, 2, 0, 1, 0, 0, 3, 1, 2, 0, 0, 4, 3, 2, 0, 1, 6, 12, 1, 8, 0, 0, 2, 6, 8, 0, 3, 2, 6, 12, 3, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 18, 87, 105, 116, 104, 100, 114, 97, 119, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 27, 101, 120, 116, 114, 97, 99, 116, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 27, 114, 101, 115, 116, 111, 114, 101, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 10, 115, 116, 97, 112, 108, 101, 95, 108, 98, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 4, 1, 9, 11, 0, 17, 0, 12, 2, 14, 2, 10, 1, 17, 2, 11, 2, 17, 1, 2};

var peer_to_peer_with_metadata_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 7, 1, 0, 2, 2, 2, 4, 3, 6, 16, 4, 22, 2, 5, 24, 29, 7, 53, 97, 8, 150, 1, 16, 0, 0, 0, 1, 1, 0, 0, 2, 0, 1, 0, 0, 3, 2, 3, 1, 1, 0, 4, 1, 3, 0, 1, 5, 1, 6, 12, 1, 8, 0, 5, 6, 8, 0, 5, 3, 10, 2, 10, 2, 0, 5, 6, 12, 5, 3, 10, 2, 10, 2, 1, 9, 0, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 18, 87, 105, 116, 104, 100, 114, 97, 119, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 27, 101, 120, 116, 114, 97, 99, 116, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 8, 112, 97, 121, 95, 102, 114, 111, 109, 27, 114, 101, 115, 116, 111, 114, 101, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 4, 1, 12, 11, 0, 17, 0, 12, 5, 14, 5, 10, 1, 10, 2, 11, 3, 11, 4, 56, 0, 11, 5, 17, 2, 2};

var preburn_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 7, 1, 0, 2, 2, 2, 4, 3, 6, 16, 4, 22, 2, 5, 24, 21, 7, 45, 96, 8, 141, 1, 16, 0, 0, 0, 1, 1, 0, 0, 2, 0, 1, 0, 0, 3, 2, 3, 1, 1, 0, 4, 1, 3, 0, 1, 5, 1, 6, 12, 1, 8, 0, 3, 6, 12, 6, 8, 0, 3, 0, 2, 6, 12, 3, 1, 9, 0, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 18, 87, 105, 116, 104, 100, 114, 97, 119, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 27, 101, 120, 116, 114, 97, 99, 116, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 7, 112, 114, 101, 98, 117, 114, 110, 27, 114, 101, 115, 116, 111, 114, 101, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 4, 1, 10, 10, 0, 17, 0, 12, 2, 11, 0, 14, 2, 10, 1, 56, 0, 11, 2, 17, 2, 2};

var publish_shared_ed25519_public_key_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 2, 3, 2, 5, 5, 7, 6, 7, 13, 31, 8, 44, 16, 0, 0, 0, 1, 0, 1, 0, 2, 6, 12, 10, 2, 0, 22, 83, 104, 97, 114, 101, 100, 69, 100, 50, 53, 53, 49, 57, 80, 117, 98, 108, 105, 99, 75, 101, 121, 7, 112, 117, 98, 108, 105, 115, 104, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 4, 11, 0, 11, 1, 17, 0, 2};

var register_validator_config_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 2, 3, 2, 5, 5, 7, 15, 7, 22, 27, 8, 49, 16, 0, 0, 0, 1, 0, 1, 0, 7, 6, 12, 5, 10, 2, 10, 2, 10, 2, 10, 2, 10, 2, 0, 15, 86, 97, 108, 105, 100, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 10, 115, 101, 116, 95, 99, 111, 110, 102, 105, 103, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 9, 11, 0, 10, 1, 11, 2, 11, 3, 11, 4, 11, 5, 11, 6, 17, 0, 2};

var remove_validator_and_reconfigure_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 6, 3, 6, 15, 5, 21, 24, 7, 45, 95, 8, 140, 1, 16, 0, 0, 0, 1, 0, 2, 1, 3, 0, 1, 0, 2, 4, 2, 3, 0, 0, 5, 4, 1, 0, 2, 6, 12, 3, 0, 1, 5, 1, 10, 2, 2, 6, 12, 5, 4, 6, 12, 3, 10, 2, 5, 2, 1, 3, 11, 76, 105, 98, 114, 97, 83, 121, 115, 116, 101, 109, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 15, 86, 97, 108, 105, 100, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 14, 103, 101, 116, 95, 104, 117, 109, 97, 110, 95, 110, 97, 109, 101, 16, 114, 101, 109, 111, 118, 101, 95, 118, 97, 108, 105, 100, 97, 116, 111, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 5, 6, 18, 10, 0, 10, 1, 17, 0, 10, 3, 17, 1, 11, 2, 33, 12, 4, 11, 4, 3, 14, 11, 0, 1, 6, 0, 0, 0, 0, 0, 0, 0, 0, 39, 11, 0, 10, 3, 17, 2, 2};

var rotate_authentication_key_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 4, 2, 4, 4, 3, 8, 25, 5, 33, 32, 7, 65, 175, 1, 8, 240, 1, 16, 0, 0, 0, 1, 0, 3, 1, 0, 1, 2, 0, 1, 0, 0, 4, 0, 2, 0, 0, 5, 3, 4, 0, 0, 6, 2, 5, 0, 0, 7, 6, 5, 0, 1, 6, 12, 1, 5, 1, 8, 0, 1, 6, 8, 0, 1, 6, 5, 0, 2, 6, 8, 0, 10, 2, 2, 6, 12, 10, 2, 3, 8, 0, 1, 3, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 6, 83, 105, 103, 110, 101, 114, 10, 97, 100, 100, 114, 101, 115, 115, 95, 111, 102, 21, 75, 101, 121, 82, 111, 116, 97, 116, 105, 111, 110, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 101, 120, 116, 114, 97, 99, 116, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 95, 97, 100, 100, 114, 101, 115, 115, 31, 114, 101, 115, 116, 111, 114, 101, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 25, 114, 111, 116, 97, 116, 101, 95, 97, 117, 116, 104, 101, 110, 116, 105, 99, 97, 116, 105, 111, 110, 95, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 7, 8, 20, 10, 0, 17, 1, 12, 2, 14, 2, 17, 2, 20, 11, 0, 17, 0, 33, 12, 3, 11, 3, 3, 14, 6, 0, 0, 0, 0, 0, 0, 0, 0, 39, 14, 2, 11, 1, 17, 4, 11, 2, 17, 3, 2};

var rotate_authentication_key_with_nonce_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 4, 2, 4, 4, 3, 8, 20, 5, 28, 23, 7, 51, 160, 1, 8, 211, 1, 16, 0, 0, 0, 1, 0, 3, 1, 0, 1, 2, 0, 1, 0, 0, 4, 2, 3, 0, 0, 5, 3, 1, 0, 0, 6, 4, 1, 0, 2, 6, 12, 3, 0, 1, 6, 12, 1, 8, 0, 2, 6, 8, 0, 10, 2, 3, 6, 12, 3, 10, 2, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 21, 75, 101, 121, 82, 111, 116, 97, 116, 105, 111, 110, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 101, 120, 116, 114, 97, 99, 116, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 114, 101, 115, 116, 111, 114, 101, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 25, 114, 111, 116, 97, 116, 101, 95, 97, 117, 116, 104, 101, 110, 116, 105, 99, 97, 116, 105, 111, 110, 95, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 5, 3, 12, 10, 0, 10, 1, 17, 0, 11, 0, 17, 1, 12, 3, 14, 3, 11, 2, 17, 3, 11, 3, 17, 2, 2};

var rotate_authentication_key_with_nonce_admin_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 4, 2, 4, 4, 3, 8, 20, 5, 28, 25, 7, 53, 160, 1, 8, 213, 1, 16, 0, 0, 0, 1, 0, 3, 1, 0, 1, 2, 0, 1, 0, 0, 4, 2, 3, 0, 0, 5, 3, 1, 0, 0, 6, 4, 1, 0, 2, 6, 12, 3, 0, 1, 6, 12, 1, 8, 0, 2, 6, 8, 0, 10, 2, 4, 6, 12, 6, 12, 3, 10, 2, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 21, 75, 101, 121, 82, 111, 116, 97, 116, 105, 111, 110, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 101, 120, 116, 114, 97, 99, 116, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 114, 101, 115, 116, 111, 114, 101, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 25, 114, 111, 116, 97, 116, 101, 95, 97, 117, 116, 104, 101, 110, 116, 105, 99, 97, 116, 105, 111, 110, 95, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 5, 3, 12, 11, 0, 10, 2, 17, 0, 11, 1, 17, 1, 12, 4, 14, 4, 11, 3, 17, 3, 11, 4, 17, 2, 2};

var rotate_authentication_key_with_recovery_address_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 2, 3, 2, 5, 5, 7, 8, 7, 15, 42, 8, 57, 16, 0, 0, 0, 1, 0, 1, 0, 4, 6, 12, 5, 5, 10, 2, 0, 15, 82, 101, 99, 111, 118, 101, 114, 121, 65, 100, 100, 114, 101, 115, 115, 25, 114, 111, 116, 97, 116, 101, 95, 97, 117, 116, 104, 101, 110, 116, 105, 99, 97, 116, 105, 111, 110, 95, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 6, 11, 0, 10, 1, 10, 2, 11, 3, 17, 0, 2};

var rotate_dual_attestation_info_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 2, 3, 2, 10, 5, 12, 13, 7, 25, 61, 8, 86, 16, 0, 0, 0, 1, 0, 1, 0, 0, 2, 0, 1, 0, 2, 6, 12, 10, 2, 0, 3, 6, 12, 10, 2, 10, 2, 15, 68, 117, 97, 108, 65, 116, 116, 101, 115, 116, 97, 116, 105, 111, 110, 15, 114, 111, 116, 97, 116, 101, 95, 98, 97, 115, 101, 95, 117, 114, 108, 28, 114, 111, 116, 97, 116, 101, 95, 99, 111, 109, 112, 108, 105, 97, 110, 99, 101, 95, 112, 117, 98, 108, 105, 99, 95, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 1, 7, 10, 0, 11, 1, 17, 0, 11, 0, 11, 2, 17, 1, 2};

var rotate_shared_ed25519_public_key_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 2, 3, 2, 5, 5, 7, 6, 7, 13, 34, 8, 47, 16, 0, 0, 0, 1, 0, 1, 0, 2, 6, 12, 10, 2, 0, 22, 83, 104, 97, 114, 101, 100, 69, 100, 50, 53, 53, 49, 57, 80, 117, 98, 108, 105, 99, 75, 101, 121, 10, 114, 111, 116, 97, 116, 101, 95, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 4, 11, 0, 11, 1, 17, 0, 2};

var set_validator_config_and_reconfigure_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 4, 3, 4, 10, 5, 14, 19, 7, 33, 69, 8, 102, 16, 0, 0, 0, 1, 1, 2, 0, 1, 0, 0, 3, 2, 1, 0, 7, 6, 12, 5, 10, 2, 10, 2, 10, 2, 10, 2, 10, 2, 0, 2, 6, 12, 5, 11, 76, 105, 98, 114, 97, 83, 121, 115, 116, 101, 109, 15, 86, 97, 108, 105, 100, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 10, 115, 101, 116, 95, 99, 111, 110, 102, 105, 103, 29, 117, 112, 100, 97, 116, 101, 95, 99, 111, 110, 102, 105, 103, 95, 97, 110, 100, 95, 114, 101, 99, 111, 110, 102, 105, 103, 117, 114, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 12, 10, 0, 10, 1, 11, 2, 11, 3, 11, 4, 11, 5, 11, 6, 17, 0, 11, 0, 10, 1, 17, 1, 2};

var set_validator_operator_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 4, 3, 4, 10, 5, 14, 19, 7, 33, 68, 8, 101, 16, 0, 0, 0, 1, 1, 2, 0, 1, 0, 0, 3, 2, 3, 0, 1, 5, 1, 10, 2, 2, 6, 12, 5, 0, 3, 6, 12, 10, 2, 5, 2, 1, 3, 15, 86, 97, 108, 105, 100, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 23, 86, 97, 108, 105, 100, 97, 116, 111, 114, 79, 112, 101, 114, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 14, 103, 101, 116, 95, 104, 117, 109, 97, 110, 95, 110, 97, 109, 101, 12, 115, 101, 116, 95, 111, 112, 101, 114, 97, 116, 111, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 4, 5, 15, 10, 2, 17, 0, 11, 1, 33, 12, 3, 11, 3, 3, 11, 11, 0, 1, 6, 0, 0, 0, 0, 0, 0, 0, 0, 39, 11, 0, 10, 2, 17, 1, 2};

var set_validator_operator_with_nonce_admin_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 6, 3, 6, 15, 5, 21, 26, 7, 47, 103, 8, 150, 1, 16, 0, 0, 0, 1, 0, 2, 0, 3, 0, 1, 0, 2, 4, 2, 3, 0, 1, 5, 4, 1, 0, 2, 6, 12, 3, 0, 1, 5, 1, 10, 2, 2, 6, 12, 5, 5, 6, 12, 6, 12, 3, 10, 2, 5, 2, 1, 3, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 15, 86, 97, 108, 105, 100, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 23, 86, 97, 108, 105, 100, 97, 116, 111, 114, 79, 112, 101, 114, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 14, 103, 101, 116, 95, 104, 117, 109, 97, 110, 95, 110, 97, 109, 101, 12, 115, 101, 116, 95, 111, 112, 101, 114, 97, 116, 111, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 5, 6, 18, 11, 0, 10, 2, 17, 0, 10, 4, 17, 1, 11, 3, 33, 12, 5, 11, 5, 3, 14, 11, 1, 1, 6, 0, 0, 0, 0, 0, 0, 0, 0, 39, 11, 1, 10, 4, 17, 2, 2};

var tiered_mint_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 4, 3, 4, 11, 4, 15, 2, 5, 17, 21, 7, 38, 60, 8, 98, 16, 0, 0, 0, 1, 1, 2, 0, 1, 0, 0, 3, 2, 1, 1, 1, 1, 4, 2, 6, 12, 3, 0, 4, 6, 12, 5, 3, 3, 5, 6, 12, 3, 5, 3, 3, 1, 9, 0, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 11, 116, 105, 101, 114, 101, 100, 95, 109, 105, 110, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 3, 1, 9, 10, 0, 10, 1, 17, 0, 11, 0, 10, 2, 10, 3, 10, 4, 56, 0, 2};

var unfreeze_account_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 4, 3, 4, 10, 5, 14, 14, 7, 28, 68, 8, 96, 16, 0, 0, 0, 1, 0, 2, 0, 1, 0, 1, 3, 2, 1, 0, 2, 6, 12, 5, 0, 2, 6, 12, 3, 3, 6, 12, 3, 5, 15, 65, 99, 99, 111, 117, 110, 116, 70, 114, 101, 101, 122, 105, 110, 103, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 16, 117, 110, 102, 114, 101, 101, 122, 101, 95, 97, 99, 99, 111, 117, 110, 116, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 3, 1, 7, 10, 0, 10, 1, 17, 1, 11, 0, 10, 2, 17, 0, 2};

var unmint_lbr_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 2, 2, 2, 4, 3, 6, 15, 5, 21, 16, 7, 37, 101, 8, 138, 1, 16, 0, 0, 0, 1, 1, 0, 0, 2, 0, 1, 0, 0, 3, 1, 2, 0, 0, 4, 3, 2, 0, 1, 6, 12, 1, 8, 0, 0, 2, 6, 8, 0, 3, 2, 6, 12, 3, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 18, 87, 105, 116, 104, 100, 114, 97, 119, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 27, 101, 120, 116, 114, 97, 99, 116, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 27, 114, 101, 115, 116, 111, 114, 101, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 12, 117, 110, 115, 116, 97, 112, 108, 101, 95, 108, 98, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 4, 1, 9, 11, 0, 17, 0, 12, 2, 14, 2, 10, 1, 17, 2, 11, 2, 17, 1, 2};

var update_dual_attestation_limit_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 4, 3, 4, 10, 5, 14, 10, 7, 24, 72, 8, 96, 16, 0, 0, 0, 1, 0, 2, 0, 1, 0, 1, 3, 0, 1, 0, 2, 6, 12, 3, 0, 3, 6, 12, 3, 3, 15, 68, 117, 97, 108, 65, 116, 116, 101, 115, 116, 97, 116, 105, 111, 110, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 20, 115, 101, 116, 95, 109, 105, 99, 114, 111, 108, 105, 98, 114, 97, 95, 108, 105, 109, 105, 116, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 1, 7, 10, 0, 10, 1, 17, 1, 11, 0, 10, 2, 17, 0, 2};

var update_exchange_rate_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 7, 1, 0, 6, 2, 6, 4, 3, 10, 16, 4, 26, 2, 5, 28, 25, 7, 53, 100, 8, 153, 1, 16, 0, 0, 0, 1, 0, 2, 0, 0, 2, 0, 0, 3, 0, 1, 0, 2, 4, 2, 3, 0, 1, 5, 4, 3, 1, 1, 2, 6, 2, 3, 3, 1, 8, 0, 2, 6, 12, 3, 0, 2, 6, 12, 8, 0, 4, 6, 12, 3, 3, 3, 1, 9, 0, 12, 70, 105, 120, 101, 100, 80, 111, 105, 110, 116, 51, 50, 5, 76, 105, 98, 114, 97, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 20, 99, 114, 101, 97, 116, 101, 95, 102, 114, 111, 109, 95, 114, 97, 116, 105, 111, 110, 97, 108, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 24, 117, 112, 100, 97, 116, 101, 95, 108, 98, 114, 95, 101, 120, 99, 104, 97, 110, 103, 101, 95, 114, 97, 116, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 5, 1, 11, 10, 0, 10, 1, 17, 1, 10, 2, 10, 3, 17, 0, 12, 4, 11, 0, 11, 4, 56, 0, 2};

var update_libra_version_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 4, 3, 4, 10, 5, 14, 10, 7, 24, 52, 8, 76, 16, 0, 0, 0, 1, 0, 2, 0, 1, 0, 1, 3, 0, 1, 0, 2, 6, 12, 3, 0, 3, 6, 12, 3, 3, 12, 76, 105, 98, 114, 97, 86, 101, 114, 115, 105, 111, 110, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 3, 115, 101, 116, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 1, 7, 10, 0, 10, 1, 17, 1, 11, 0, 10, 2, 17, 0, 2};

var update_minting_ability_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 2, 3, 2, 6, 4, 8, 2, 5, 10, 8, 7, 18, 29, 8, 47, 16, 0, 0, 0, 1, 0, 1, 1, 1, 0, 2, 2, 6, 12, 1, 0, 1, 9, 0, 5, 76, 105, 98, 114, 97, 22, 117, 112, 100, 97, 116, 101, 95, 109, 105, 110, 116, 105, 110, 103, 95, 97, 98, 105, 108, 105, 116, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 4, 11, 0, 10, 1, 56, 0, 2};

var script_decoder_map = map[string]func(*libratypes.Script) (ScriptCall, error) {
	string(add_currency_to_account_code): decode_add_currency_to_account_script,
	string(add_recovery_rotation_capability_code): decode_add_recovery_rotation_capability_script,
	string(add_to_script_allow_list_code): decode_add_to_script_allow_list_script,
	string(add_validator_and_reconfigure_code): decode_add_validator_and_reconfigure_script,
	string(burn_code): decode_burn_script,
	string(burn_txn_fees_code): decode_burn_txn_fees_script,
	string(cancel_burn_code): decode_cancel_burn_script,
	string(create_child_vasp_account_code): decode_create_child_vasp_account_script,
	string(create_designated_dealer_code): decode_create_designated_dealer_script,
	string(create_parent_vasp_account_code): decode_create_parent_vasp_account_script,
	string(create_recovery_address_code): decode_create_recovery_address_script,
	string(create_validator_account_code): decode_create_validator_account_script,
	string(create_validator_operator_account_code): decode_create_validator_operator_account_script,
	string(freeze_account_code): decode_freeze_account_script,
	string(mint_lbr_code): decode_mint_lbr_script,
	string(peer_to_peer_with_metadata_code): decode_peer_to_peer_with_metadata_script,
	string(preburn_code): decode_preburn_script,
	string(publish_shared_ed25519_public_key_code): decode_publish_shared_ed25519_public_key_script,
	string(register_validator_config_code): decode_register_validator_config_script,
	string(remove_validator_and_reconfigure_code): decode_remove_validator_and_reconfigure_script,
	string(rotate_authentication_key_code): decode_rotate_authentication_key_script,
	string(rotate_authentication_key_with_nonce_code): decode_rotate_authentication_key_with_nonce_script,
	string(rotate_authentication_key_with_nonce_admin_code): decode_rotate_authentication_key_with_nonce_admin_script,
	string(rotate_authentication_key_with_recovery_address_code): decode_rotate_authentication_key_with_recovery_address_script,
	string(rotate_dual_attestation_info_code): decode_rotate_dual_attestation_info_script,
	string(rotate_shared_ed25519_public_key_code): decode_rotate_shared_ed25519_public_key_script,
	string(set_validator_config_and_reconfigure_code): decode_set_validator_config_and_reconfigure_script,
	string(set_validator_operator_code): decode_set_validator_operator_script,
	string(set_validator_operator_with_nonce_admin_code): decode_set_validator_operator_with_nonce_admin_script,
	string(tiered_mint_code): decode_tiered_mint_script,
	string(unfreeze_account_code): decode_unfreeze_account_script,
	string(unmint_lbr_code): decode_unmint_lbr_script,
	string(update_dual_attestation_limit_code): decode_update_dual_attestation_limit_script,
	string(update_exchange_rate_code): decode_update_exchange_rate_script,
	string(update_libra_version_code): decode_update_libra_version_script,
	string(update_minting_ability_code): decode_update_minting_ability_script,
}

func decode_bool_argument(arg libratypes.TransactionArgument) (value bool, err error) {
	if arg, ok := arg.(*libratypes.TransactionArgument__Bool); ok {
		value = arg.Value
	} else {
		err = fmt.Errorf("Was expecting a Bool argument")
	}
	return
}


func decode_u64_argument(arg libratypes.TransactionArgument) (value uint64, err error) {
	if arg, ok := arg.(*libratypes.TransactionArgument__U64); ok {
		value = arg.Value
	} else {
		err = fmt.Errorf("Was expecting a U64 argument")
	}
	return
}


func decode_address_argument(arg libratypes.TransactionArgument) (value libratypes.AccountAddress, err error) {
	if arg, ok := arg.(*libratypes.TransactionArgument__Address); ok {
		value = arg.Value
	} else {
		err = fmt.Errorf("Was expecting a Address argument")
	}
	return
}


func decode_u8vector_argument(arg libratypes.TransactionArgument) (value []byte, err error) {
	if arg, ok := arg.(*libratypes.TransactionArgument__U8Vector); ok {
		value = arg.Value
	} else {
		err = fmt.Errorf("Was expecting a U8Vector argument")
	}
	return
}

