package stdlib


import (
	"fmt"
	"github.com/libra/libra-client-sdk-go/libratypes"
)


// Structured representation of a call into a known Move script.
type ScriptCall interface {
	isScriptCall()
}

// # Summary
// Adds a zero `Currency` balance to the sending `account`. This will enable `account` to
// send, receive, and hold `Libra::Libra<Currency>` coins. This transaction can be
// successfully sent by any account that is allowed to hold balances
// (e.g., VASP, Designated Dealer).
//
// # Technical Description
// After the successful execution of this transaction the sending account will have a
// `LibraAccount::Balance<Currency>` resource with zero balance published under it. Only
// accounts that can hold balances can send this transaction, the sending account cannot
// already have a `LibraAccount::Balance<Currency>` published under it.
//
// # Parameters
// | Name       | Type      | Description                                                                                                                                         |
// | ------     | ------    | -------------                                                                                                                                       |
// | `Currency` | Type      | The Move type for the `Currency` being added to the sending account of the transaction. `Currency` must be an already-registered currency on-chain. |
// | `account`  | `&signer` | The signer of the sending account of the transaction.                                                                                               |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                             | Description                                                                |
// | ----------------            | --------------                           | -------------                                                              |
// | `Errors::NOT_PUBLISHED`     | `Libra::ECURRENCY_INFO`                  | The `Currency` is not a registered currency on-chain.                      |
// | `Errors::INVALID_ARGUMENT`  | `LibraAccount::EROLE_CANT_STORE_BALANCE` | The sending `account`'s role does not permit balances.                     |
// | `Errors::ALREADY_PUBLISHED` | `LibraAccount::EADD_EXISTING_CURRENCY`   | A balance for `Currency` is already published under the sending `account`. |
//
// # Related Scripts
// * `Script::create_child_vasp_account`
// * `Script::create_parent_vasp_account`
// * `Script::peer_to_peer_with_metadata`
type ScriptCall__AddCurrencyToAccount struct {
	Currency libratypes.TypeTag
}

func (*ScriptCall__AddCurrencyToAccount) isScriptCall() {}

// # Summary
// Stores the sending accounts ability to rotate its authentication key with a designated recovery
// account. Both the sending and recovery accounts need to belong to the same VASP and
// both be VASP accounts. After this transaction both the sending account and the
// specified recovery account can rotate the sender account's authentication key.
//
// # Technical Description
// Adds the `LibraAccount::KeyRotationCapability` for the sending account
// (`to_recover_account`) to the `RecoveryAddress::RecoveryAddress` resource under
// `recovery_address`. After this transaction has been executed successfully the account at
// `recovery_address` and the `to_recover_account` may rotate the authentication key of
// `to_recover_account` (the sender of this transaction).
//
// The sending account of this transaction (`to_recover_account`) must not have previously given away its unique key
// rotation capability, and must be a VASP account. The account at `recovery_address`
// must also be a VASP account belonging to the same VASP as the `to_recover_account`.
// Additionally the account at `recovery_address` must have already initialized itself as
// a recovery account address using the `Script::create_recovery_address` transaction script.
//
// The sending account's (`to_recover_account`) key rotation capability is
// removed in this transaction and stored in the `RecoveryAddress::RecoveryAddress`
// resource stored under the account at `recovery_address`.
//
// # Parameters
// | Name                 | Type      | Description                                                                                                |
// | ------               | ------    | -------------                                                                                              |
// | `to_recover_account` | `&signer` | The signer reference of the sending account of this transaction.                                           |
// | `recovery_address`   | `address` | The account address where the `to_recover_account`'s `LibraAccount::KeyRotationCapability` will be stored. |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                               | Description                                                                                     |
// | ----------------           | --------------                                             | -------------                                                                                   |
// | `Errors::INVALID_STATE`    | `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` | `to_recover_account` has already delegated/extracted its `LibraAccount::KeyRotationCapability`. |
// | `Errors::NOT_PUBLISHED`    | `RecoveryAddress::ERECOVERY_ADDRESS`                       | `recovery_address` does not have a `RecoveryAddress` resource published under it.               |
// | `Errors::INVALID_ARGUMENT` | `RecoveryAddress::EINVALID_KEY_ROTATION_DELEGATION`        | `to_recover_account` and `recovery_address` do not belong to the same VASP.                     |
//
// # Related Scripts
// * `Script::create_recovery_address`
// * `Script::rotate_authentication_key_with_recovery_address`
type ScriptCall__AddRecoveryRotationCapability struct {
	RecoveryAddress libratypes.AccountAddress
}

func (*ScriptCall__AddRecoveryRotationCapability) isScriptCall() {}

// # Summary
// Adds a script hash to the transaction allowlist. This transaction
// can only be sent by the Libra Root account. Scripts with this hash can be
// sent afterward the successful execution of this script.
//
// # Technical Description
//
// The sending account (`lr_account`) must be the Libra Root account. The script allow
// list must not already hold the script `hash` being added. The `sliding_nonce` must be
// a valid nonce for the Libra Root account. After this transaction has executed
// successfully a reconfiguration will be initiated, and the on-chain config
// `LibraTransactionPublishingOption::LibraTransactionPublishingOption`'s
// `script_allow_list` field will contain the new script `hash` and transactions
// with this `hash` can be successfully sent to the network.
//
// # Parameters
// | Name            | Type         | Description                                                                                     |
// | ------          | ------       | -------------                                                                                   |
// | `lr_account`    | `&signer`    | The signer reference of the sending account of this transaction. Must be the Libra Root signer. |
// | `hash`          | `vector<u8>` | The hash of the script to be added to the script allowlist.                                     |
// | `sliding_nonce` | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                      |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                                           | Description                                                                                |
// | ----------------           | --------------                                                         | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                                         | A `SlidingNonce` resource is not published under `lr_account`.                             |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`                                         | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`                                         | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`                                | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ELIBRA_ROOT`                                           | The sending account is not the Libra Root account.                                         |
// | `Errors::REQUIRES_ROLE`    | `Roles::ELIBRA_ROOT`                                                   | The sending account is not the Libra Root account.                                         |
// | `Errors::INVALID_ARGUMENT` | `LibraTransactionPublishingOption::EINVALID_SCRIPT_HASH`               | The script `hash` is an invalid length.                                                    |
// | `Errors::INVALID_ARGUMENT` | `LibraTransactionPublishingOption::EALLOWLIST_ALREADY_CONTAINS_SCRIPT` | The on-chain allowlist already contains the script `hash`.                                 |
type ScriptCall__AddToScriptAllowList struct {
	Hash []byte
	SlidingNonce uint64
}

func (*ScriptCall__AddToScriptAllowList) isScriptCall() {}

// # Summary
// Adds a validator account to the validator set, and triggers a
// reconfiguration of the system to admit the account to the validator set for the system. This
// transaction can only be successfully called by the Libra Root account.
//
// # Technical Description
// This script adds the account at `validator_address` to the validator set.
// This transaction emits a `LibraConfig::NewEpochEvent` event and triggers a
// reconfiguration. Once the reconfiguration triggered by this script's
// execution has been performed, the account at the `validator_address` is
// considered to be a validator in the network.
//
// This transaction script will fail if the `validator_address` address is already in the validator set
// or does not have a `ValidatorConfig::ValidatorConfig` resource already published under it.
//
// # Parameters
// | Name                | Type         | Description                                                                                                                        |
// | ------              | ------       | -------------                                                                                                                      |
// | `lr_account`        | `&signer`    | The signer reference of the sending account of this transaction. Must be the Libra Root signer.                                    |
// | `sliding_nonce`     | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                                         |
// | `validator_name`    | `vector<u8>` | ASCII-encoded human name for the validator. Must match the human name in the `ValidatorConfig::ValidatorConfig` for the validator. |
// | `validator_address` | `address`    | The validator account address to be added to the validator set.                                                                    |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                  | Description                                                                                                                               |
// | ----------------           | --------------                                | -------------                                                                                                                             |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                | A `SlidingNonce` resource is not published under `lr_account`.                                                                            |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`                | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not.                                                |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`                | The `sliding_nonce` is too far in the future.                                                                                             |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`       | The `sliding_nonce` has been previously recorded.                                                                                         |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ELIBRA_ROOT`                  | The sending account is not the Libra Root account.                                                                                        |
// | `Errors::REQUIRES_ROLE`    | `Roles::ELIBRA_ROOT`                          | The sending account is not the Libra Root account.                                                                                        |
// | 0                          | 0                                             | The provided `validator_name` does not match the already-recorded human name for the validator.                                           |
// | `Errors::INVALID_ARGUMENT` | `LibraSystem::EINVALID_PROSPECTIVE_VALIDATOR` | The validator to be added does not have a `ValidatorConfig::ValidatorConfig` resource published under it, or its `config` field is empty. |
// | `Errors::INVALID_ARGUMENT` | `LibraSystem::EALREADY_A_VALIDATOR`           | The `validator_address` account is already a registered validator.                                                                        |
// | `Errors::INVALID_STATE`    | `LibraConfig::EINVALID_BLOCK_TIME`            | An invalid time value was encountered in reconfiguration. Unlikely to occur.                                                              |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::create_validator_operator_account`
// * `Script::register_validator_config`
// * `Script::remove_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::set_validator_config_and_reconfigure`
type ScriptCall__AddValidatorAndReconfigure struct {
	SlidingNonce uint64
	ValidatorName []byte
	ValidatorAddress libratypes.AccountAddress
}

func (*ScriptCall__AddValidatorAndReconfigure) isScriptCall() {}

// # Summary
// Burns all coins held in the preburn resource at the specified
// preburn address and removes them from the system. The sending account must
// be the Treasury Compliance account.
// The account that holds the preburn resource will normally be a Designated
// Dealer, but there are no enforced requirements that it be one.
//
// # Technical Description
// This transaction permanently destroys all the coins of `Token` type
// stored in the `Libra::Preburn<Token>` resource published under the
// `preburn_address` account address.
//
// This transaction will only succeed if the sending `account` has a
// `Libra::BurnCapability<Token>`, and a `Libra::Preburn<Token>` resource
// exists under `preburn_address`, with a non-zero `to_burn` field. After the successful execution
// of this transaction the `total_value` field in the
// `Libra::CurrencyInfo<Token>` resource published under `0xA550C18` will be
// decremented by the value of the `to_burn` field of the preburn resource
// under `preburn_address` immediately before this transaction, and the
// `to_burn` field of the preburn resource will have a zero value.
//
// ## Events
// The successful execution of this transaction will emit a `Libra::BurnEvent` on the event handle
// held in the `Libra::CurrencyInfo<Token>` resource's `burn_events` published under
// `0xA550C18`.
//
// # Parameters
// | Name              | Type      | Description                                                                                                                  |
// | ------            | ------    | -------------                                                                                                                |
// | `Token`           | Type      | The Move type for the `Token` currency being burned. `Token` must be an already-registered currency on-chain.                |
// | `tc_account`      | `&signer` | The signer reference of the sending account of this transaction, must have a burn capability for `Token` published under it. |
// | `sliding_nonce`   | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                                   |
// | `preburn_address` | `address` | The address where the coins to-be-burned are currently held.                                                                 |
//
// # Common Abort Conditions
// | Error Category                | Error Reason                            | Description                                                                                           |
// | ----------------              | --------------                          | -------------                                                                                         |
// | `Errors::NOT_PUBLISHED`       | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `account`.                                           |
// | `Errors::INVALID_ARGUMENT`    | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not.            |
// | `Errors::INVALID_ARGUMENT`    | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                                         |
// | `Errors::INVALID_ARGUMENT`    | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                                     |
// | `Errors::REQUIRES_CAPABILITY` | `Libra::EBURN_CAPABILITY`               | The sending `account` does not have a `Libra::BurnCapability<Token>` published under it.              |
// | `Errors::NOT_PUBLISHED`       | `Libra::EPREBURN`                       | The account at `preburn_address` does not have a `Libra::Preburn<Token>` resource published under it. |
// | `Errors::INVALID_STATE`       | `Libra::EPREBURN_EMPTY`                 | The `Libra::Preburn<Token>` resource is empty (has a value of 0).                                     |
// | `Errors::NOT_PUBLISHED`       | `Libra::ECURRENCY_INFO`                 | The specified `Token` is not a registered currency on-chain.                                          |
//
// # Related Scripts
// * `Script::burn_txn_fees`
// * `Script::cancel_burn`
// * `Script::preburn`
type ScriptCall__Burn struct {
	Token libratypes.TypeTag
	SlidingNonce uint64
	PreburnAddress libratypes.AccountAddress
}

func (*ScriptCall__Burn) isScriptCall() {}

// # Summary
// Burns the transaction fees collected in the `CoinType` currency so that the
// Libra association may reclaim the backing coins off-chain. May only be sent
// by the Treasury Compliance account.
//
// # Technical Description
// Burns the transaction fees collected in `CoinType` so that the
// association may reclaim the backing coins. Once this transaction has executed
// successfully all transaction fees that will have been collected in
// `CoinType` since the last time this script was called with that specific
// currency. Both `balance` and `preburn` fields in the
// `TransactionFee::TransactionFee<CoinType>` resource published under the `0xB1E55ED`
// account address will have a value of 0 after the successful execution of this script.
//
// ## Events
// The successful execution of this transaction will emit a `Libra::BurnEvent` on the event handle
// held in the `Libra::CurrencyInfo<CoinType>` resource's `burn_events` published under
// `0xA550C18`.
//
// # Parameters
// | Name         | Type      | Description                                                                                                                                         |
// | ------       | ------    | -------------                                                                                                                                       |
// | `CoinType`   | Type      | The Move type for the `CoinType` being added to the sending account of the transaction. `CoinType` must be an already-registered currency on-chain. |
// | `tc_account` | `&signer` | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account.                                           |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                          | Description                                                 |
// | ----------------           | --------------                        | -------------                                               |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ETREASURY_COMPLIANCE` | The sending account is not the Treasury Compliance account. |
// | `Errors::NOT_PUBLISHED`    | `TransactionFee::ETRANSACTION_FEE`    | `CoinType` is not an accepted transaction fee currency.     |
// | `Errors::INVALID_ARGUMENT` | `Libra::ECOIN`                        | The collected fees in `CoinType` are zero.                  |
//
// # Related Scripts
// * `Script::burn`
// * `Script::cancel_burn`
type ScriptCall__BurnTxnFees struct {
	CoinType libratypes.TypeTag
}

func (*ScriptCall__BurnTxnFees) isScriptCall() {}

// # Summary
// Cancels and returns all coins held in the preburn area under
// `preburn_address` and returns the funds to the `preburn_address`'s balance.
// Can only be successfully sent by an account with Treasury Compliance role.
//
// # Technical Description
// Cancels and returns all coins held in the `Libra::Preburn<Token>` resource under the `preburn_address` and
// return the funds to the `preburn_address` account's `LibraAccount::Balance<Token>`.
// The transaction must be sent by an `account` with a `Libra::BurnCapability<Token>`
// resource published under it. The account at `preburn_address` must have a
// `Libra::Preburn<Token>` resource published under it, and its value must be nonzero. The transaction removes
// the entire balance held in the `Libra::Preburn<Token>` resource, and returns it back to the account's
// `LibraAccount::Balance<Token>` under `preburn_address`. Due to this, the account at
// `preburn_address` must already have a balance in the `Token` currency published
// before this script is called otherwise the transaction will fail.
//
// ## Events
// The successful execution of this transaction will emit:
// * A `Libra::CancelBurnEvent` on the event handle held in the `Libra::CurrencyInfo<Token>`
// resource's `burn_events` published under `0xA550C18`.
// * A `LibraAccount::ReceivedPaymentEvent` on the `preburn_address`'s
// `LibraAccount::LibraAccount` `received_events` event handle with both the `payer` and `payee`
// being `preburn_address`.
//
// # Parameters
// | Name              | Type      | Description                                                                                                                          |
// | ------            | ------    | -------------                                                                                                                        |
// | `Token`           | Type      | The Move type for the `Token` currenty that burning is being cancelled for. `Token` must be an already-registered currency on-chain. |
// | `account`         | `&signer` | The signer reference of the sending account of this transaction, must have a burn capability for `Token` published under it.         |
// | `preburn_address` | `address` | The address where the coins to-be-burned are currently held.                                                                         |
//
// # Common Abort Conditions
// | Error Category                | Error Reason                                     | Description                                                                                           |
// | ----------------              | --------------                                   | -------------                                                                                         |
// | `Errors::REQUIRES_CAPABILITY` | `Libra::EBURN_CAPABILITY`                        | The sending `account` does not have a `Libra::BurnCapability<Token>` published under it.              |
// | `Errors::NOT_PUBLISHED`       | `Libra::EPREBURN`                                | The account at `preburn_address` does not have a `Libra::Preburn<Token>` resource published under it. |
// | `Errors::NOT_PUBLISHED`       | `Libra::ECURRENCY_INFO`                          | The specified `Token` is not a registered currency on-chain.                                          |
// | `Errors::INVALID_ARGUMENT`    | `LibraAccount::ECOIN_DEPOSIT_IS_ZERO`            | The value held in the preburn resource was zero.                                                      |
// | `Errors::INVALID_ARGUMENT`    | `LibraAccount::EPAYEE_CANT_ACCEPT_CURRENCY_TYPE` | The account at `preburn_address` doesn't have a balance resource for `Token`.                         |
// | `Errors::LIMIT_EXCEEDED`      | `LibraAccount::EDEPOSIT_EXCEEDS_LIMITS`          | The depositing of the funds held in the prebun area would exceed the `account`'s account limits.      |
// | `Errors::INVALID_STATE`       | `DualAttestation::EPAYEE_COMPLIANCE_KEY_NOT_SET` | The `account` does not have a compliance key set on it but dual attestion checking was performed.     |
//
// # Related Scripts
// * `Script::burn_txn_fees`
// * `Script::burn`
// * `Script::preburn`
type ScriptCall__CancelBurn struct {
	Token libratypes.TypeTag
	PreburnAddress libratypes.AccountAddress
}

func (*ScriptCall__CancelBurn) isScriptCall() {}

// # Summary
// Creates a Child VASP account with its parent being the sending account of the transaction.
// The sender of the transaction must be a Parent VASP account.
//
// # Technical Description
// Creates a `ChildVASP` account for the sender `parent_vasp` at `child_address` with a balance of
// `child_initial_balance` in `CoinType` and an initial authentication key of
// `auth_key_prefix | child_address`.
//
// If `add_all_currencies` is true, the child address will have a zero balance in all available
// currencies in the system.
//
// The new account will be a child account of the transaction sender, which must be a
// Parent VASP account. The child account will be recorded against the limit of
// child accounts of the creating Parent VASP account.
//
// ## Events
// Successful execution with a `child_initial_balance` greater than zero will emit:
// * A `LibraAccount::SentPaymentEvent` with the `payer` field being the Parent VASP's address,
// and payee field being `child_address`. This is emitted on the Parent VASP's
// `LibraAccount::LibraAccount` `sent_events` handle.
// * A `LibraAccount::ReceivedPaymentEvent` with the  `payer` field being the Parent VASP's address,
// and payee field being `child_address`. This is emitted on the new Child VASPS's
// `LibraAccount::LibraAccount` `received_events` handle.
//
// # Parameters
// | Name                    | Type         | Description                                                                                                                                 |
// | ------                  | ------       | -------------                                                                                                                               |
// | `CoinType`              | Type         | The Move type for the `CoinType` that the child account should be created with. `CoinType` must be an already-registered currency on-chain. |
// | `parent_vasp`           | `&signer`    | The signer reference of the sending account. Must be a Parent VASP account.                                                                 |
// | `child_address`         | `address`    | Address of the to-be-created Child VASP account.                                                                                            |
// | `auth_key_prefix`       | `vector<u8>` | The authentication key prefix that will be used initially for the newly created account.                                                    |
// | `add_all_currencies`    | `bool`       | Whether to publish balance resources for all known currencies when the account is created.                                                  |
// | `child_initial_balance` | `u64`        | The initial balance in `CoinType` to give the child account when it's created.                                                              |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                                             | Description                                                                              |
// | ----------------            | --------------                                           | -------------                                                                            |
// | `Errors::INVALID_ARGUMENT`  | `LibraAccount::EMALFORMED_AUTHENTICATION_KEY`            | The `auth_key_prefix` was not of length 32.                                              |
// | `Errors::REQUIRES_ROLE`     | `Roles::EPARENT_VASP`                                    | The sending account wasn't a Parent VASP account.                                        |
// | `Errors::ALREADY_PUBLISHED` | `Roles::EROLE_ID`                                        | The `child_address` address is already taken.                                            |
// | `Errors::LIMIT_EXCEEDED`    | `VASP::ETOO_MANY_CHILDREN`                               | The sending account has reached the maximum number of allowed child accounts.            |
// | `Errors::NOT_PUBLISHED`     | `Libra::ECURRENCY_INFO`                                  | The `CoinType` is not a registered currency on-chain.                                    |
// | `Errors::INVALID_STATE`     | `LibraAccount::EWITHDRAWAL_CAPABILITY_ALREADY_EXTRACTED` | The withdrawal capability for the sending account has already been extracted.            |
// | `Errors::NOT_PUBLISHED`     | `LibraAccount::EPAYER_DOESNT_HOLD_CURRENCY`              | The sending account doesn't have a balance in `CoinType`.                                |
// | `Errors::LIMIT_EXCEEDED`    | `LibraAccount::EINSUFFICIENT_BALANCE`                    | The sending account doesn't have at least `child_initial_balance` of `CoinType` balance. |
// | `Errors::INVALID_ARGUMENT`  | `LibraAccount::ECANNOT_CREATE_AT_VM_RESERVED`            | The `child_address` is the reserved address 0x0.                                         |
//
// # Related Scripts
// * `Script::create_parent_vasp_account`
// * `Script::add_currency_to_account`
// * `Script::rotate_authentication_key`
// * `Script::add_recovery_rotation_capability`
// * `Script::create_recovery_address`
type ScriptCall__CreateChildVaspAccount struct {
	CoinType libratypes.TypeTag
	ChildAddress libratypes.AccountAddress
	AuthKeyPrefix []byte
	AddAllCurrencies bool
	ChildInitialBalance uint64
}

func (*ScriptCall__CreateChildVaspAccount) isScriptCall() {}

// # Summary
// Creates a Designated Dealer account with the provided information, and initializes it with
// default mint tiers. The transaction can only be sent by the Treasury Compliance account.
//
// # Technical Description
// Creates an account with the Designated Dealer role at `addr` with authentication key
// `auth_key_prefix` | `addr` and a 0 balance of type `Currency`. If `add_all_currencies` is true,
// 0 balances for all available currencies in the system will also be added. This can only be
// invoked by an account with the TreasuryCompliance role.
//
// At the time of creation the account is also initialized with default mint tiers of (500_000,
// 5000_000, 50_000_000, 500_000_000), and preburn areas for each currency that is added to the
// account.
//
// # Parameters
// | Name                 | Type         | Description                                                                                                                                         |
// | ------               | ------       | -------------                                                                                                                                       |
// | `Currency`           | Type         | The Move type for the `Currency` that the Designated Dealer should be initialized with. `Currency` must be an already-registered currency on-chain. |
// | `tc_account`         | `&signer`    | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account.                                           |
// | `sliding_nonce`      | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                                                          |
// | `addr`               | `address`    | Address of the to-be-created Designated Dealer account.                                                                                             |
// | `auth_key_prefix`    | `vector<u8>` | The authentication key prefix that will be used initially for the newly created account.                                                            |
// | `human_name`         | `vector<u8>` | ASCII-encoded human name for the Designated Dealer.                                                                                                 |
// | `add_all_currencies` | `bool`       | Whether to publish preburn, balance, and tier info resources for all known (SCS) currencies or just `Currency` when the account is created.         |
//

// # Common Abort Conditions
// | Error Category              | Error Reason                            | Description                                                                                |
// | ----------------            | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`     | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `tc_account`.                             |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS`  | `CoreAddresses::ETREASURY_COMPLIANCE`   | The sending account is not the Treasury Compliance account.                                |
// | `Errors::REQUIRES_ROLE`     | `Roles::ETREASURY_COMPLIANCE`           | The sending account is not the Treasury Compliance account.                                |
// | `Errors::NOT_PUBLISHED`     | `Libra::ECURRENCY_INFO`                 | The `Currency` is not a registered currency on-chain.                                      |
// | `Errors::ALREADY_PUBLISHED` | `Roles::EROLE_ID`                       | The `addr` address is already taken.                                                       |
//
// # Related Scripts
// * `Script::tiered_mint`
// * `Script::peer_to_peer_with_metadata`
// * `Script::rotate_dual_attestation_info`
type ScriptCall__CreateDesignatedDealer struct {
	Currency libratypes.TypeTag
	SlidingNonce uint64
	Addr libratypes.AccountAddress
	AuthKeyPrefix []byte
	HumanName []byte
	AddAllCurrencies bool
}

func (*ScriptCall__CreateDesignatedDealer) isScriptCall() {}

// # Summary
// Creates a Parent VASP account with the specified human name. Must be called by the Treasury Compliance account.
//
// # Technical Description
// Creates an account with the Parent VASP role at `address` with authentication key
// `auth_key_prefix` | `new_account_address` and a 0 balance of type `CoinType`. If
// `add_all_currencies` is true, 0 balances for all available currencies in the system will
// also be added. This can only be invoked by an TreasuryCompliance account.
// `sliding_nonce` is a unique nonce for operation, see `SlidingNonce` for details.
//
// # Parameters
// | Name                  | Type         | Description                                                                                                                                                    |
// | ------                | ------       | -------------                                                                                                                                                  |
// | `CoinType`            | Type         | The Move type for the `CoinType` currency that the Parent VASP account should be initialized with. `CoinType` must be an already-registered currency on-chain. |
// | `tc_account`          | `&signer`    | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account.                                                      |
// | `sliding_nonce`       | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                                                                     |
// | `new_account_address` | `address`    | Address of the to-be-created Parent VASP account.                                                                                                              |
// | `auth_key_prefix`     | `vector<u8>` | The authentication key prefix that will be used initially for the newly created account.                                                                       |
// | `human_name`          | `vector<u8>` | ASCII-encoded human name for the Parent VASP.                                                                                                                  |
// | `add_all_currencies`  | `bool`       | Whether to publish balance resources for all known currencies when the account is created.                                                                     |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                            | Description                                                                                |
// | ----------------            | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`     | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `tc_account`.                             |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS`  | `CoreAddresses::ETREASURY_COMPLIANCE`   | The sending account is not the Treasury Compliance account.                                |
// | `Errors::REQUIRES_ROLE`     | `Roles::ETREASURY_COMPLIANCE`           | The sending account is not the Treasury Compliance account.                                |
// | `Errors::NOT_PUBLISHED`     | `Libra::ECURRENCY_INFO`                 | The `CoinType` is not a registered currency on-chain.                                      |
// | `Errors::ALREADY_PUBLISHED` | `Roles::EROLE_ID`                       | The `new_account_address` address is already taken.                                        |
//
// # Related Scripts
// * `Script::create_child_vasp_account`
// * `Script::add_currency_to_account`
// * `Script::rotate_authentication_key`
// * `Script::add_recovery_rotation_capability`
// * `Script::create_recovery_address`
// * `Script::rotate_dual_attestation_info`
type ScriptCall__CreateParentVaspAccount struct {
	CoinType libratypes.TypeTag
	SlidingNonce uint64
	NewAccountAddress libratypes.AccountAddress
	AuthKeyPrefix []byte
	HumanName []byte
	AddAllCurrencies bool
}

func (*ScriptCall__CreateParentVaspAccount) isScriptCall() {}

// # Summary
// Initializes the sending account as a recovery address that may be used by
// the VASP that it belongs to. The sending account must be a VASP account.
// Multiple recovery addresses can exist for a single VASP, but accounts in
// each must be disjoint.
//
// # Technical Description
// Publishes a `RecoveryAddress::RecoveryAddress` resource under `account`. It then
// extracts the `LibraAccount::KeyRotationCapability` for `account` and adds
// it to the resource. After the successful execution of this transaction
// other accounts may add their key rotation to this resource so that `account`
// may be used as a recovery account for those accounts.
//
// # Parameters
// | Name      | Type      | Description                                           |
// | ------    | ------    | -------------                                         |
// | `account` | `&signer` | The signer of the sending account of the transaction. |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                                               | Description                                                                                   |
// | ----------------            | --------------                                             | -------------                                                                                 |
// | `Errors::INVALID_STATE`     | `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` | `account` has already delegated/extracted its `LibraAccount::KeyRotationCapability`.          |
// | `Errors::INVALID_ARGUMENT`  | `RecoveryAddress::ENOT_A_VASP`                             | `account` is not a VASP account.                                                              |
// | `Errors::INVALID_ARGUMENT`  | `RecoveryAddress::EKEY_ROTATION_DEPENDENCY_CYCLE`          | A key rotation recovery cycle would be created by adding `account`'s key rotation capability. |
// | `Errors::ALREADY_PUBLISHED` | `RecoveryAddress::ERECOVERY_ADDRESS`                       | A `RecoveryAddress::RecoveryAddress` resource has already been published under `account`.     |
//
// # Related Scripts
// * `Script::add_recovery_rotation_capability`
// * `Script::rotate_authentication_key_with_recovery_address`
type ScriptCall__CreateRecoveryAddress struct {
}

func (*ScriptCall__CreateRecoveryAddress) isScriptCall() {}

// # Summary
// Creates a Validator account. This transaction can only be sent by the Libra
// Root account.
//
// # Technical Description
// Creates an account with a Validator role at `new_account_address`, with authentication key
// `auth_key_prefix` | `new_account_address`. It publishes a
// `ValidatorConfig::ValidatorConfig` resource with empty `config`, and
// `operator_account` fields. The `human_name` field of the
// `ValidatorConfig::ValidatorConfig` is set to the passed in `human_name`.
// This script does not add the validator to the validator set or the system,
// but only creates the account.
//
// # Parameters
// | Name                  | Type         | Description                                                                                     |
// | ------                | ------       | -------------                                                                                   |
// | `lr_account`          | `&signer`    | The signer reference of the sending account of this transaction. Must be the Libra Root signer. |
// | `sliding_nonce`       | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                      |
// | `new_account_address` | `address`    | Address of the to-be-created Validator account.                                                 |
// | `auth_key_prefix`     | `vector<u8>` | The authentication key prefix that will be used initially for the newly created account.        |
// | `human_name`          | `vector<u8>` | ASCII-encoded human name for the validator.                                                     |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                            | Description                                                                                |
// | ----------------            | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`     | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `lr_account`.                             |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS`  | `CoreAddresses::ELIBRA_ROOT`            | The sending account is not the Libra Root account.                                         |
// | `Errors::REQUIRES_ROLE`     | `Roles::ELIBRA_ROOT`                    | The sending account is not the Libra Root account.                                         |
// | `Errors::ALREADY_PUBLISHED` | `Roles::EROLE_ID`                       | The `new_account_address` address is already taken.                                        |
//
// # Related Scripts
// * `Script::add_validator_and_reconfigure`
// * `Script::create_validator_operator_account`
// * `Script::register_validator_config`
// * `Script::remove_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::set_validator_config_and_reconfigure`
type ScriptCall__CreateValidatorAccount struct {
	SlidingNonce uint64
	NewAccountAddress libratypes.AccountAddress
	AuthKeyPrefix []byte
	HumanName []byte
}

func (*ScriptCall__CreateValidatorAccount) isScriptCall() {}

// # Summary
// Creates a Validator Operator account. This transaction can only be sent by the Libra
// Root account.
//
// # Technical Description
// Creates an account with a Validator Operator role at `new_account_address`, with authentication key
// `auth_key_prefix` | `new_account_address`. It publishes a
// `ValidatorOperatorConfig::ValidatorOperatorConfig` resource with the specified `human_name`.
// This script does not assign the validator operator to any validator accounts but only creates the account.
//
// # Parameters
// | Name                  | Type         | Description                                                                                     |
// | ------                | ------       | -------------                                                                                   |
// | `lr_account`          | `&signer`    | The signer reference of the sending account of this transaction. Must be the Libra Root signer. |
// | `sliding_nonce`       | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                      |
// | `new_account_address` | `address`    | Address of the to-be-created Validator account.                                                 |
// | `auth_key_prefix`     | `vector<u8>` | The authentication key prefix that will be used initially for the newly created account.        |
// | `human_name`          | `vector<u8>` | ASCII-encoded human name for the validator.                                                     |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                            | Description                                                                                |
// | ----------------            | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`     | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `lr_account`.                             |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS`  | `CoreAddresses::ELIBRA_ROOT`            | The sending account is not the Libra Root account.                                         |
// | `Errors::REQUIRES_ROLE`     | `Roles::ELIBRA_ROOT`                    | The sending account is not the Libra Root account.                                         |
// | `Errors::ALREADY_PUBLISHED` | `Roles::EROLE_ID`                       | The `new_account_address` address is already taken.                                        |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::add_validator_and_reconfigure`
// * `Script::register_validator_config`
// * `Script::remove_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::set_validator_config_and_reconfigure`
type ScriptCall__CreateValidatorOperatorAccount struct {
	SlidingNonce uint64
	NewAccountAddress libratypes.AccountAddress
	AuthKeyPrefix []byte
	HumanName []byte
}

func (*ScriptCall__CreateValidatorOperatorAccount) isScriptCall() {}

// # Summary
// Freezes the account at `address`. The sending account of this transaction
// must be the Treasury Compliance account. The account being frozen cannot be
// the Libra Root or Treasury Compliance account. After the successful
// execution of this transaction no transactions may be sent from the frozen
// account, and the frozen account may not send or receive coins.
//
// # Technical Description
// Sets the `AccountFreezing::FreezingBit` to `true` and emits a
// `AccountFreezing::FreezeAccountEvent`. The transaction sender must be the
// Treasury Compliance account, but the account at `to_freeze_account` must
// not be either `0xA550C18` (the Libra Root address), or `0xB1E55ED` (the
// Treasury Compliance address). Note that this is a per-account property
// e.g., freezing a Parent VASP will not effect the status any of its child
// accounts and vice versa.
//

// ## Events
// Successful execution of this transaction will emit a `AccountFreezing::FreezeAccountEvent` on
// the `freeze_event_handle` held in the `AccountFreezing::FreezeEventsHolder` resource published
// under `0xA550C18` with the `frozen_address` being the `to_freeze_account`.
//
// # Parameters
// | Name                | Type      | Description                                                                                               |
// | ------              | ------    | -------------                                                                                             |
// | `tc_account`        | `&signer` | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account. |
// | `sliding_nonce`     | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                |
// | `to_freeze_account` | `address` | The account address to be frozen.                                                                         |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                 | Description                                                                                |
// | ----------------           | --------------                               | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`               | A `SlidingNonce` resource is not published under `tc_account`.                             |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`               | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`               | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`      | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ETREASURY_COMPLIANCE`        | The sending account is not the Treasury Compliance account.                                |
// | `Errors::REQUIRES_ROLE`    | `Roles::ETREASURY_COMPLIANCE`                | The sending account is not the Treasury Compliance account.                                |
// | `Errors::INVALID_ARGUMENT` | `AccountFreezing::ECANNOT_FREEZE_TC`         | `to_freeze_account` was the Treasury Compliance account (`0xB1E55ED`).                     |
// | `Errors::INVALID_ARGUMENT` | `AccountFreezing::ECANNOT_FREEZE_LIBRA_ROOT` | `to_freeze_account` was the Libra Root account (`0xA550C18`).                              |
//
// # Related Scripts
// * `Scripts::unfreeze_account`
type ScriptCall__FreezeAccount struct {
	SlidingNonce uint64
	ToFreezeAccount libratypes.AccountAddress
}

func (*ScriptCall__FreezeAccount) isScriptCall() {}

// # Summary
// Transfers a given number of coins in a specified currency from one account to another.
// Transfers over a specified amount defined on-chain that are between two different VASPs, or
// other accounts that have opted-in will be subject to on-chain checks to ensure the receiver has
// agreed to receive the coins.  This transaction can be sent by any account that can hold a
// balance, and to any account that can hold a balance. Both accounts must hold balances in the
// currency being transacted.
//
// # Technical Description
//
// Transfers `amount` coins of type `Currency` from `payer` to `payee` with (optional) associated
// `metadata` and an (optional) `metadata_signature` on the message
// `metadata` | `Signer::address_of(payer)` | `amount` | `DualAttestation::DOMAIN_SEPARATOR`.
// The `metadata` and `metadata_signature` parameters are only required if `amount` >=
// `DualAttestation::get_cur_microlibra_limit` LBR and `payer` and `payee` are distinct VASPs.
// However, a transaction sender can opt in to dual attestation even when it is not required
// (e.g., a DesignatedDealer -> VASP payment) by providing a non-empty `metadata_signature`.
// Standardized `metadata` LCS format can be found in `libra_types::transaction::metadata::Metadata`.
//
// ## Events
// Successful execution of this script emits two events:
// * A `LibraAccount::SentPaymentEvent` on `payer`'s `LibraAccount::LibraAccount` `sent_events` handle; and
// * A `LibraAccount::ReceivedPaymentEvent` on `payee`'s `LibraAccount::LibraAccount` `received_events` handle.
//
// # Parameters
// | Name                 | Type         | Description                                                                                                                  |
// | ------               | ------       | -------------                                                                                                                |
// | `Currency`           | Type         | The Move type for the `Currency` being sent in this transaction. `Currency` must be an already-registered currency on-chain. |
// | `payer`              | `&signer`    | The signer reference of the sending account that coins are being transferred from.                                           |
// | `payee`              | `address`    | The address of the account the coins are being transferred to.                                                               |
// | `metadata`           | `vector<u8>` | Optional metadata about this payment.                                                                                        |
// | `metadata_signature` | `vector<u8>` | Optional signature over `metadata` and payment information. See                                                              |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                     | Description                                                                                                                         |
// | ----------------           | --------------                                   | -------------                                                                                                                       |
// | `Errors::NOT_PUBLISHED`    | `LibraAccount::EPAYER_DOESNT_HOLD_CURRENCY`      | `payer` doesn't hold a balance in `Currency`.                                                                                       |
// | `Errors::LIMIT_EXCEEDED`   | `LibraAccount::EINSUFFICIENT_BALANCE`            | `amount` is greater than `payer`'s balance in `Currency`.                                                                           |
// | `Errors::INVALID_ARGUMENT` | `LibraAccount::ECOIN_DEPOSIT_IS_ZERO`            | `amount` is zero.                                                                                                                   |
// | `Errors::NOT_PUBLISHED`    | `LibraAccount::EPAYEE_DOES_NOT_EXIST`            | No account exists at the `payee` address.                                                                                           |
// | `Errors::INVALID_ARGUMENT` | `LibraAccount::EPAYEE_CANT_ACCEPT_CURRENCY_TYPE` | An account exists at `payee`, but it does not accept payments in `Currency`.                                                        |
// | `Errors::INVALID_STATE`    | `AccountFreezing::EACCOUNT_FROZEN`               | The `payee` account is frozen.                                                                                                      |
// | `Errors::INVALID_ARGUMENT` | `DualAttestation::EMALFORMED_METADATA_SIGNATURE` | `metadata_signature` is not 64 bytes.                                                                                               |
// | `Errors::INVALID_ARGUMENT` | `DualAttestation::EINVALID_METADATA_SIGNATURE`   | `metadata_signature` does not verify on the against the `payee'`s `DualAttestation::Credential` `compliance_public_key` public key. |
// | `Errors::LIMIT_EXCEEDED`   | `LibraAccount::EWITHDRAWAL_EXCEEDS_LIMITS`       | `payer` has exceeded its daily withdrawal limits for the backing coins of LBR.                                                      |
// | `Errors::LIMIT_EXCEEDED`   | `LibraAccount::EDEPOSIT_EXCEEDS_LIMITS`          | `payee` has exceeded its daily deposit limits for LBR.                                                                              |
//
// # Related Scripts
// * `Script::create_child_vasp_account`
// * `Script::create_parent_vasp_account`
// * `Script::add_currency_to_account`
type ScriptCall__PeerToPeerWithMetadata struct {
	Currency libratypes.TypeTag
	Payee libratypes.AccountAddress
	Amount uint64
	Metadata []byte
	MetadataSignature []byte
}

func (*ScriptCall__PeerToPeerWithMetadata) isScriptCall() {}

// # Summary
// Moves a specified number of coins in a given currency from the account's
// balance to its preburn area after which the coins may be burned. This
// transaction may be sent by any account that holds a balance and preburn area
// in the specified currency.
//
// # Technical Description
// Moves the specified `amount` of coins in `Token` currency from the sending `account`'s
// `LibraAccount::Balance<Token>` to the `Libra::Preburn<Token>` published under the same
// `account`. `account` must have both of these resources published under it at the start of this
// transaction in order for it to execute successfully.
//
// ## Events
// Successful execution of this script emits two events:
// * `LibraAccount::SentPaymentEvent ` on `account`'s `LibraAccount::LibraAccount` `sent_events`
// handle with the `payee` and `payer` fields being `account`'s address; and
// * A `Libra::PreburnEvent` with `Token`'s currency code on the
// `Libra::CurrencyInfo<Token`'s `preburn_events` handle for `Token` and with
// `preburn_address` set to `account`'s address.
//
// # Parameters
// | Name      | Type      | Description                                                                                                                      |
// | ------    | ------    | -------------                                                                                                                    |
// | `Token`   | Type      | The Move type for the `Token` currency being moved to the preburn area. `Token` must be an already-registered currency on-chain. |
// | `account` | `&signer` | The signer reference of the sending account.                                                                                     |
// | `amount`  | `u64`     | The amount in `Token` to be moved to the preburn area.                                                                           |
//
// # Common Abort Conditions
// | Error Category           | Error Reason                                             | Description                                                                             |
// | ----------------         | --------------                                           | -------------                                                                           |
// | `Errors::NOT_PUBLISHED`  | `Libra::ECURRENCY_INFO`                                  | The `Token` is not a registered currency on-chain.                                      |
// | `Errors::INVALID_STATE`  | `LibraAccount::EWITHDRAWAL_CAPABILITY_ALREADY_EXTRACTED` | The withdrawal capability for `account` has already been extracted.                     |
// | `Errors::LIMIT_EXCEEDED` | `LibraAccount::EINSUFFICIENT_BALANCE`                    | `amount` is greater than `payer`'s balance in `Token`.                                  |
// | `Errors::NOT_PUBLISHED`  | `LibraAccount::EPAYER_DOESNT_HOLD_CURRENCY`              | `account` doesn't hold a balance in `Token`.                                            |
// | `Errors::NOT_PUBLISHED`  | `Libra::EPREBURN`                                        | `account` doesn't have a `Libra::Preburn<Token>` resource published under it.           |
// | `Errors::INVALID_STATE`  | `Libra::EPREBURN_OCCUPIED`                               | The `value` field in the `Libra::Preburn<Token>` resource under the sender is non-zero. |
//
// # Related Scripts
// * `Script::cancel_burn`
// * `Script::burn`
// * `Script::burn_txn_fees`
type ScriptCall__Preburn struct {
	Token libratypes.TypeTag
	Amount uint64
}

func (*ScriptCall__Preburn) isScriptCall() {}

// # Summary
// Rotates the authentication key of the sending account to the
// newly-specified public key and publishes a new shared authentication key
// under the sender's account. Any account can send this transaction.
//
// # Technical Description
// Rotates the authentication key of the sending account to `public_key`,
// and publishes a `SharedEd25519PublicKey::SharedEd25519PublicKey` resource
// containing the 32-byte ed25519 `public_key` and the `LibraAccount::KeyRotationCapability` for
// `account` under `account`.
//
// # Parameters
// | Name         | Type         | Description                                                                               |
// | ------       | ------       | -------------                                                                             |
// | `account`    | `&signer`    | The signer reference of the sending account of the transaction.                           |
// | `public_key` | `vector<u8>` | 32-byte Ed25519 public key for `account`' authentication key to be rotated to and stored. |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                                               | Description                                                                                         |
// | ----------------            | --------------                                             | -------------                                                                                       |
// | `Errors::INVALID_STATE`     | `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` | `account` has already delegated/extracted its `LibraAccount::KeyRotationCapability` resource.       |
// | `Errors::ALREADY_PUBLISHED` | `SharedEd25519PublicKey::ESHARED_KEY`                      | The `SharedEd25519PublicKey::SharedEd25519PublicKey` resource is already published under `account`. |
// | `Errors::INVALID_ARGUMENT`  | `SharedEd25519PublicKey::EMALFORMED_PUBLIC_KEY`            | `public_key` is an invalid ed25519 public key.                                                      |
//
// # Related Scripts
// * `Script::rotate_shared_ed25519_public_key`
type ScriptCall__PublishSharedEd25519PublicKey struct {
	PublicKey []byte
}

func (*ScriptCall__PublishSharedEd25519PublicKey) isScriptCall() {}

// # Summary
// Updates a validator's configuration. This does not reconfigure the system and will not update
// the configuration in the validator set that is seen by other validators in the network. Can
// only be successfully sent by a Validator Operator account that is already registered with a
// validator.
//
// # Technical Description
// This updates the fields with corresponding names held in the `ValidatorConfig::ValidatorConfig`
// config resource held under `validator_account`. It does not emit a `LibraConfig::NewEpochEvent`
// so the copy of this config held in the validator set will not be updated, and the changes are
// only "locally" under the `validator_account` account address.
//
// # Parameters
// | Name                          | Type         | Description                                                                                                                  |
// | ------                        | ------       | -------------                                                                                                                |
// | `validator_operator_account`  | `&signer`    | Signer reference of the sending account. Must be the registered validator operator for the validator at `validator_address`. |
// | `validator_account`           | `address`    | The address of the validator's `ValidatorConfig::ValidatorConfig` resource being updated.                                    |
// | `consensus_pubkey`            | `vector<u8>` | New Ed25519 public key to be used in the updated `ValidatorConfig::ValidatorConfig`.                                         |
// | `validator_network_addresses` | `vector<u8>` | New set of `validator_network_addresses` to be used in the updated `ValidatorConfig::ValidatorConfig`.                       |
// | `fullnode_network_addresses`  | `vector<u8>` | New set of `fullnode_network_addresses` to be used in the updated `ValidatorConfig::ValidatorConfig`.                        |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                   | Description                                                                                           |
// | ----------------           | --------------                                 | -------------                                                                                         |
// | `Errors::NOT_PUBLISHED`    | `ValidatorConfig::EVALIDATOR_CONFIG`           | `validator_address` does not have a `ValidatorConfig::ValidatorConfig` resource published under it.   |
// | `Errors::INVALID_ARGUMENT` | `ValidatorConfig::EINVALID_TRANSACTION_SENDER` | `validator_operator_account` is not the registered operator for the validator at `validator_address`. |
// | `Errors::INVALID_ARGUMENT` | `ValidatorConfig::EINVALID_CONSENSUS_KEY`      | `consensus_pubkey` is not a valid ed25519 public key.                                                 |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::create_validator_operator_account`
// * `Script::add_validator_and_reconfigure`
// * `Script::remove_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::set_validator_config_and_reconfigure`
type ScriptCall__RegisterValidatorConfig struct {
	ValidatorAccount libratypes.AccountAddress
	ConsensusPubkey []byte
	ValidatorNetworkAddresses []byte
	FullnodeNetworkAddresses []byte
}

func (*ScriptCall__RegisterValidatorConfig) isScriptCall() {}

// # Summary
// This script removes a validator account from the validator set, and triggers a reconfiguration
// of the system to remove the validator from the system. This transaction can only be
// successfully called by the Libra Root account.
//
// # Technical Description
// This script removes the account at `validator_address` from the validator set. This transaction
// emits a `LibraConfig::NewEpochEvent` event. Once the reconfiguration triggered by this event
// has been performed, the account at `validator_address` is no longer considered to be a
// validator in the network. This transaction will fail if the validator at `validator_address`
// is not in the validator set.
//
// # Parameters
// | Name                | Type         | Description                                                                                                                        |
// | ------              | ------       | -------------                                                                                                                      |
// | `lr_account`        | `&signer`    | The signer reference of the sending account of this transaction. Must be the Libra Root signer.                                    |
// | `sliding_nonce`     | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                                         |
// | `validator_name`    | `vector<u8>` | ASCII-encoded human name for the validator. Must match the human name in the `ValidatorConfig::ValidatorConfig` for the validator. |
// | `validator_address` | `address`    | The validator account address to be removed from the validator set.                                                                |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                            | Description                                                                                     |
// | ----------------           | --------------                          | -------------                                                                                   |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `lr_account`.                                  |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not.      |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                                   |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                               |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`          | The sending account is not the Libra Root account or Treasury Compliance account                |
// | 0                          | 0                                       | The provided `validator_name` does not match the already-recorded human name for the validator. |
// | `Errors::INVALID_ARGUMENT` | `LibraSystem::ENOT_AN_ACTIVE_VALIDATOR` | The validator to be removed is not in the validator set.                                        |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ELIBRA_ROOT`            | The sending account is not the Libra Root account.                                              |
// | `Errors::REQUIRES_ROLE`    | `Roles::ELIBRA_ROOT`                    | The sending account is not the Libra Root account.                                              |
// | `Errors::INVALID_STATE`    | `LibraConfig::EINVALID_BLOCK_TIME`      | An invalid time value was encountered in reconfiguration. Unlikely to occur.                    |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::create_validator_operator_account`
// * `Script::register_validator_config`
// * `Script::add_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::set_validator_config_and_reconfigure`
type ScriptCall__RemoveValidatorAndReconfigure struct {
	SlidingNonce uint64
	ValidatorName []byte
	ValidatorAddress libratypes.AccountAddress
}

func (*ScriptCall__RemoveValidatorAndReconfigure) isScriptCall() {}

// # Summary
// Rotates the transaction sender's authentication key to the supplied new authentication key. May
// be sent by any account.
//
// # Technical Description
// Rotate the `account`'s `LibraAccount::LibraAccount` `authentication_key` field to `new_key`.
// `new_key` must be a valid ed25519 public key, and `account` must not have previously delegated
// its `LibraAccount::KeyRotationCapability`.
//
// # Parameters
// | Name      | Type         | Description                                                 |
// | ------    | ------       | -------------                                               |
// | `account` | `&signer`    | Signer reference of the sending account of the transaction. |
// | `new_key` | `vector<u8>` | New ed25519 public key to be used for `account`.            |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                               | Description                                                                              |
// | ----------------           | --------------                                             | -------------                                                                            |
// | `Errors::INVALID_STATE`    | `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` | `account` has already delegated/extracted its `LibraAccount::KeyRotationCapability`.     |
// | `Errors::INVALID_ARGUMENT` | `LibraAccount::EMALFORMED_AUTHENTICATION_KEY`              | `new_key` was an invalid length.                                                         |
//
// # Related Scripts
// * `Script::rotate_authentication_key_with_nonce`
// * `Script::rotate_authentication_key_with_nonce_admin`
// * `Script::rotate_authentication_key_with_recovery_address`
type ScriptCall__RotateAuthenticationKey struct {
	NewKey []byte
}

func (*ScriptCall__RotateAuthenticationKey) isScriptCall() {}

// # Summary
// Rotates the sender's authentication key to the supplied new authentication key. May be sent by
// any account that has a sliding nonce resource published under it (usually this is Treasury
// Compliance or Libra Root accounts).
//
// # Technical Description
// Rotates the `account`'s `LibraAccount::LibraAccount` `authentication_key` field to `new_key`.
// `new_key` must be a valid ed25519 public key, and `account` must not have previously delegated
// its `LibraAccount::KeyRotationCapability`.
//
// # Parameters
// | Name            | Type         | Description                                                                |
// | ------          | ------       | -------------                                                              |
// | `account`       | `&signer`    | Signer reference of the sending account of the transaction.                |
// | `sliding_nonce` | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction. |
// | `new_key`       | `vector<u8>` | New ed25519 public key to be used for `account`.                           |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                               | Description                                                                                |
// | ----------------           | --------------                                             | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                             | A `SlidingNonce` resource is not published under `account`.                                |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`                             | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`                             | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`                    | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::INVALID_STATE`    | `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` | `account` has already delegated/extracted its `LibraAccount::KeyRotationCapability`.       |
// | `Errors::INVALID_ARGUMENT` | `LibraAccount::EMALFORMED_AUTHENTICATION_KEY`              | `new_key` was an invalid length.                                                           |
//
// # Related Scripts
// * `Script::rotate_authentication_key`
// * `Script::rotate_authentication_key_with_nonce_admin`
// * `Script::rotate_authentication_key_with_recovery_address`
type ScriptCall__RotateAuthenticationKeyWithNonce struct {
	SlidingNonce uint64
	NewKey []byte
}

func (*ScriptCall__RotateAuthenticationKeyWithNonce) isScriptCall() {}

// # Summary
// Rotates the specified account's authentication key to the supplied new authentication key. May
// only be sent by the Libra Root account as a write set transaction.
//
// # Technical Description
// Rotate the `account`'s `LibraAccount::LibraAccount` `authentication_key` field to `new_key`.
// `new_key` must be a valid ed25519 public key, and `account` must not have previously delegated
// its `LibraAccount::KeyRotationCapability`.
//
// # Parameters
// | Name            | Type         | Description                                                                                                  |
// | ------          | ------       | -------------                                                                                                |
// | `lr_account`    | `&signer`    | The signer reference of the sending account of the write set transaction. May only be the Libra Root signer. |
// | `account`       | `&signer`    | Signer reference of account specified in the `execute_as` field of the write set transaction.                |
// | `sliding_nonce` | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction for Libra Root.                    |
// | `new_key`       | `vector<u8>` | New ed25519 public key to be used for `account`.                                                             |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                               | Description                                                                                                |
// | ----------------           | --------------                                             | -------------                                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                             | A `SlidingNonce` resource is not published under `lr_account`.                                             |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`                             | The `sliding_nonce` in `lr_account` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`                             | The `sliding_nonce` in `lr_account` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`                    | The `sliding_nonce` in` lr_account` has been previously recorded.                                          |
// | `Errors::INVALID_STATE`    | `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` | `account` has already delegated/extracted its `LibraAccount::KeyRotationCapability`.                       |
// | `Errors::INVALID_ARGUMENT` | `LibraAccount::EMALFORMED_AUTHENTICATION_KEY`              | `new_key` was an invalid length.                                                                           |
//
// # Related Scripts
// * `Script::rotate_authentication_key`
// * `Script::rotate_authentication_key_with_nonce`
// * `Script::rotate_authentication_key_with_recovery_address`
type ScriptCall__RotateAuthenticationKeyWithNonceAdmin struct {
	SlidingNonce uint64
	NewKey []byte
}

func (*ScriptCall__RotateAuthenticationKeyWithNonceAdmin) isScriptCall() {}

// # Summary
// Rotates the authentication key of a specified account that is part of a recovery address to a
// new authentication key. Only used for accounts that are part of a recovery address (see
// `Script::add_recovery_rotation_capability` for account restrictions).
//
// # Technical Description
// Rotates the authentication key of the `to_recover` account to `new_key` using the
// `LibraAccount::KeyRotationCapability` stored in the `RecoveryAddress::RecoveryAddress` resource
// published under `recovery_address`. This transaction can be sent either by the `to_recover`
// account, or by the account where the `RecoveryAddress::RecoveryAddress` resource is published
// that contains `to_recover`'s `LibraAccount::KeyRotationCapability`.
//
// # Parameters
// | Name               | Type         | Description                                                                                                                    |
// | ------             | ------       | -------------                                                                                                                  |
// | `account`          | `&signer`    | Signer reference of the sending account of the transaction.                                                                    |
// | `recovery_address` | `address`    | Address where `RecoveryAddress::RecoveryAddress` that holds `to_recover`'s `LibraAccount::KeyRotationCapability` is published. |
// | `to_recover`       | `address`    | The address of the account whose authentication key will be updated.                                                           |
// | `new_key`          | `vector<u8>` | New ed25519 public key to be used for the account at the `to_recover` address.                                                 |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                  | Description                                                                                                                                          |
// | ----------------           | --------------                                | -------------                                                                                                                                        |
// | `Errors::NOT_PUBLISHED`    | `RecoveryAddress::ERECOVERY_ADDRESS`          | `recovery_address` does not have a `RecoveryAddress::RecoveryAddress` resource published under it.                                                   |
// | `Errors::INVALID_ARGUMENT` | `RecoveryAddress::ECANNOT_ROTATE_KEY`         | The address of `account` is not `recovery_address` or `to_recover`.                                                                                  |
// | `Errors::INVALID_ARGUMENT` | `RecoveryAddress::EACCOUNT_NOT_RECOVERABLE`   | `to_recover`'s `LibraAccount::KeyRotationCapability`  is not in the `RecoveryAddress::RecoveryAddress`  resource published under `recovery_address`. |
// | `Errors::INVALID_ARGUMENT` | `LibraAccount::EMALFORMED_AUTHENTICATION_KEY` | `new_key` was an invalid length.                                                                                                                     |
//
// # Related Scripts
// * `Script::rotate_authentication_key`
// * `Script::rotate_authentication_key_with_nonce`
// * `Script::rotate_authentication_key_with_nonce_admin`
type ScriptCall__RotateAuthenticationKeyWithRecoveryAddress struct {
	RecoveryAddress libratypes.AccountAddress
	ToRecover libratypes.AccountAddress
	NewKey []byte
}

func (*ScriptCall__RotateAuthenticationKeyWithRecoveryAddress) isScriptCall() {}

// # Summary
// Updates the url used for off-chain communication, and the public key used to verify dual
// attestation on-chain. Transaction can be sent by any account that has dual attestation
// information published under it. In practice the only such accounts are Designated Dealers and
// Parent VASPs.
//
// # Technical Description
// Updates the `base_url` and `compliance_public_key` fields of the `DualAttestation::Credential`
// resource published under `account`. The `new_key` must be a valid ed25519 public key.
//
// ## Events
// Successful execution of this transaction emits two events:
// * A `DualAttestation::ComplianceKeyRotationEvent` containing the new compliance public key, and
// the blockchain time at which the key was updated emitted on the `DualAttestation::Credential`
// `compliance_key_rotation_events` handle published under `account`; and
// * A `DualAttestation::BaseUrlRotationEvent` containing the new base url to be used for
// off-chain communication, and the blockchain time at which the url was updated emitted on the
// `DualAttestation::Credential` `base_url_rotation_events` handle published under `account`.
//
// # Parameters
// | Name      | Type         | Description                                                               |
// | ------    | ------       | -------------                                                             |
// | `account` | `&signer`    | Signer reference of the sending account of the transaction.               |
// | `new_url` | `vector<u8>` | ASCII-encoded url to be used for off-chain communication with `account`.  |
// | `new_key` | `vector<u8>` | New ed25519 public key to be used for on-chain dual attestation checking. |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                           | Description                                                                |
// | ----------------           | --------------                         | -------------                                                              |
// | `Errors::NOT_PUBLISHED`    | `DualAttestation::ECREDENTIAL`         | A `DualAttestation::Credential` resource is not published under `account`. |
// | `Errors::INVALID_ARGUMENT` | `DualAttestation::EINVALID_PUBLIC_KEY` | `new_key` is not a valid ed25519 public key.                               |
//
// # Related Scripts
// * `Script::create_parent_vasp_account`
// * `Script::create_designated_dealer`
// * `Script::rotate_dual_attestation_info`
type ScriptCall__RotateDualAttestationInfo struct {
	NewUrl []byte
	NewKey []byte
}

func (*ScriptCall__RotateDualAttestationInfo) isScriptCall() {}

// # Summary
// Rotates the authentication key in a `SharedEd25519PublicKey`. This transaction can be sent by
// any account that has previously published a shared ed25519 public key using
// `Script::publish_shared_ed25519_public_key`.
//
// # Technical Description
// This first rotates the public key stored in `account`'s
// `SharedEd25519PublicKey::SharedEd25519PublicKey` resource to `public_key`, after which it
// rotates the authentication key using the capability stored in `account`'s
// `SharedEd25519PublicKey::SharedEd25519PublicKey` to a new value derived from `public_key`
//
// # Parameters
// | Name         | Type         | Description                                                     |
// | ------       | ------       | -------------                                                   |
// | `account`    | `&signer`    | The signer reference of the sending account of the transaction. |
// | `public_key` | `vector<u8>` | 32-byte Ed25519 public key.                                     |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                    | Description                                                                                   |
// | ----------------           | --------------                                  | -------------                                                                                 |
// | `Errors::NOT_PUBLISHED`    | `SharedEd25519PublicKey::ESHARED_KEY`           | A `SharedEd25519PublicKey::SharedEd25519PublicKey` resource is not published under `account`. |
// | `Errors::INVALID_ARGUMENT` | `SharedEd25519PublicKey::EMALFORMED_PUBLIC_KEY` | `public_key` is an invalid ed25519 public key.                                                |
//
// # Related Scripts
// * `Script::publish_shared_ed25519_public_key`
type ScriptCall__RotateSharedEd25519PublicKey struct {
	PublicKey []byte
}

func (*ScriptCall__RotateSharedEd25519PublicKey) isScriptCall() {}

// # Summary
// Updates a validator's configuration, and triggers a reconfiguration of the system to update the
// validator set with this new validator configuration.  Can only be successfully sent by a
// Validator Operator account that is already registered with a validator.
//
// # Technical Description
// This updates the fields with corresponding names held in the `ValidatorConfig::ValidatorConfig`
// config resource held under `validator_account`. It then emits a `LibraConfig::NewEpochEvent` to
// trigger a reconfiguration of the system.  This reconfiguration will update the validator set
// on-chain with the updated `ValidatorConfig::ValidatorConfig`.
//
// # Parameters
// | Name                          | Type         | Description                                                                                                                  |
// | ------                        | ------       | -------------                                                                                                                |
// | `validator_operator_account`  | `&signer`    | Signer reference of the sending account. Must be the registered validator operator for the validator at `validator_address`. |
// | `validator_account`           | `address`    | The address of the validator's `ValidatorConfig::ValidatorConfig` resource being updated.                                    |
// | `consensus_pubkey`            | `vector<u8>` | New Ed25519 public key to be used in the updated `ValidatorConfig::ValidatorConfig`.                                         |
// | `validator_network_addresses` | `vector<u8>` | New set of `validator_network_addresses` to be used in the updated `ValidatorConfig::ValidatorConfig`.                       |
// | `fullnode_network_addresses`  | `vector<u8>` | New set of `fullnode_network_addresses` to be used in the updated `ValidatorConfig::ValidatorConfig`.                        |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                   | Description                                                                                           |
// | ----------------           | --------------                                 | -------------                                                                                         |
// | `Errors::NOT_PUBLISHED`    | `ValidatorConfig::EVALIDATOR_CONFIG`           | `validator_address` does not have a `ValidatorConfig::ValidatorConfig` resource published under it.   |
// | `Errors::REQUIRES_ROLE     | `Roles::EVALIDATOR_OPERATOR`                   | `validator_operator_account` does not have a Validator Operator role.                                 |
// | `Errors::INVALID_ARGUMENT` | `ValidatorConfig::EINVALID_TRANSACTION_SENDER` | `validator_operator_account` is not the registered operator for the validator at `validator_address`. |
// | `Errors::INVALID_ARGUMENT` | `ValidatorConfig::EINVALID_CONSENSUS_KEY`      | `consensus_pubkey` is not a valid ed25519 public key.                                                 |
// | `Errors::INVALID_STATE`    | `LibraConfig::EINVALID_BLOCK_TIME`             | An invalid time value was encountered in reconfiguration. Unlikely to occur.                          |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::create_validator_operator_account`
// * `Script::add_validator_and_reconfigure`
// * `Script::remove_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::register_validator_config`
type ScriptCall__SetValidatorConfigAndReconfigure struct {
	ValidatorAccount libratypes.AccountAddress
	ConsensusPubkey []byte
	ValidatorNetworkAddresses []byte
	FullnodeNetworkAddresses []byte
}

func (*ScriptCall__SetValidatorConfigAndReconfigure) isScriptCall() {}

// # Summary
// Sets the validator operator for a validator in the validator's configuration resource "locally"
// and does not reconfigure the system. Changes from this transaction will not picked up by the
// system until a reconfiguration of the system is triggered. May only be sent by an account with
// Validator role.
//
// # Technical Description
// Sets the account at `operator_account` address and with the specified `human_name` as an
// operator for the sending validator account. The account at `operator_account` address must have
// a Validator Operator role and have a `ValidatorOperatorConfig::ValidatorOperatorConfig`
// resource published under it. The sending `account` must be a Validator and have a
// `ValidatorConfig::ValidatorConfig` resource published under it. This script does not emit a
// `LibraConfig::NewEpochEvent` and no reconfiguration of the system is initiated by this script.
//
// # Parameters
// | Name               | Type         | Description                                                                                  |
// | ------             | ------       | -------------                                                                                |
// | `account`          | `&signer`    | The signer reference of the sending account of the transaction.                              |
// | `operator_name`    | `vector<u8>` | Validator operator's human name.                                                             |
// | `operator_account` | `address`    | Address of the validator operator account to be added as the `account` validator's operator. |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                          | Description                                                                                                                                                  |
// | ----------------           | --------------                                        | -------------                                                                                                                                                |
// | `Errors::NOT_PUBLISHED`    | `ValidatorOperatorConfig::EVALIDATOR_OPERATOR_CONFIG` | The `ValidatorOperatorConfig::ValidatorOperatorConfig` resource is not published under `operator_account`.                                                   |
// | 0                          | 0                                                     | The `human_name` field of the `ValidatorOperatorConfig::ValidatorOperatorConfig` resource under `operator_account` does not match the provided `human_name`. |
// | `Errors::REQUIRES_ROLE`    | `Roles::EVALIDATOR`                                   | `account` does not have a Validator account role.                                                                                                            |
// | `Errors::INVALID_ARGUMENT` | `ValidatorConfig::ENOT_A_VALIDATOR_OPERATOR`          | The account at `operator_account` does not have a `ValidatorOperatorConfig::ValidatorOperatorConfig` resource.                                               |
// | `Errors::NOT_PUBLISHED`    | `ValidatorConfig::EVALIDATOR_CONFIG`                  | A `ValidatorConfig::ValidatorConfig` is not published under `account`.                                                                                       |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::create_validator_operator_account`
// * `Script::register_validator_config`
// * `Script::remove_validator_and_reconfigure`
// * `Script::add_validator_and_reconfigure`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::set_validator_config_and_reconfigure`
type ScriptCall__SetValidatorOperator struct {
	OperatorName []byte
	OperatorAccount libratypes.AccountAddress
}

func (*ScriptCall__SetValidatorOperator) isScriptCall() {}

// # Summary
// Sets the validator operator for a validator in the validator's configuration resource "locally"
// and does not reconfigure the system. Changes from this transaction will not picked up by the
// system until a reconfiguration of the system is triggered. May only be sent by the Libra Root
// account as a write set transaction.
//
// # Technical Description
// Sets the account at `operator_account` address and with the specified `human_name` as an
// operator for the validator `account`. The account at `operator_account` address must have a
// Validator Operator role and have a `ValidatorOperatorConfig::ValidatorOperatorConfig` resource
// published under it. The account represented by the `account` signer must be a Validator and
// have a `ValidatorConfig::ValidatorConfig` resource published under it. No reconfiguration of
// the system is initiated by this script.
//
// # Parameters
// | Name               | Type         | Description                                                                                                  |
// | ------             | ------       | -------------                                                                                                |
// | `lr_account`       | `&signer`    | The signer reference of the sending account of the write set transaction. May only be the Libra Root signer. |
// | `account`          | `&signer`    | Signer reference of account specified in the `execute_as` field of the write set transaction.                |
// | `sliding_nonce`    | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction for Libra Root.                    |
// | `operator_name`    | `vector<u8>` | Validator operator's human name.                                                                             |
// | `operator_account` | `address`    | Address of the validator operator account to be added as the `account` validator's operator.                 |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                          | Description                                                                                                                                                  |
// | ----------------           | --------------                                        | -------------                                                                                                                                                |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                        | A `SlidingNonce` resource is not published under `lr_account`.                                                                                               |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`                        | The `sliding_nonce` in `lr_account` is too old and it's impossible to determine if it's duplicated or not.                                                   |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`                        | The `sliding_nonce` in `lr_account` is too far in the future.                                                                                                |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`               | The `sliding_nonce` in` lr_account` has been previously recorded.                                                                                            |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                        | The sending account is not the Libra Root account or Treasury Compliance account                                                                             |
// | `Errors::NOT_PUBLISHED`    | `ValidatorOperatorConfig::EVALIDATOR_OPERATOR_CONFIG` | The `ValidatorOperatorConfig::ValidatorOperatorConfig` resource is not published under `operator_account`.                                                   |
// | 0                          | 0                                                     | The `human_name` field of the `ValidatorOperatorConfig::ValidatorOperatorConfig` resource under `operator_account` does not match the provided `human_name`. |
// | `Errors::REQUIRES_ROLE`    | `Roles::EVALIDATOR`                                   | `account` does not have a Validator account role.                                                                                                            |
// | `Errors::INVALID_ARGUMENT` | `ValidatorConfig::ENOT_A_VALIDATOR_OPERATOR`          | The account at `operator_account` does not have a `ValidatorOperatorConfig::ValidatorOperatorConfig` resource.                                               |
// | `Errors::NOT_PUBLISHED`    | `ValidatorConfig::EVALIDATOR_CONFIG`                  | A `ValidatorConfig::ValidatorConfig` is not published under `account`.                                                                                       |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::create_validator_operator_account`
// * `Script::register_validator_config`
// * `Script::remove_validator_and_reconfigure`
// * `Script::add_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_config_and_reconfigure`
type ScriptCall__SetValidatorOperatorWithNonceAdmin struct {
	SlidingNonce uint64
	OperatorName []byte
	OperatorAccount libratypes.AccountAddress
}

func (*ScriptCall__SetValidatorOperatorWithNonceAdmin) isScriptCall() {}

// # Summary
// Mints a specified number of coins in a currency to a Designated Dealer. The sending account
// must be the Treasury Compliance account, and coins can only be minted to a Designated Dealer
// account.
//
// # Technical Description
// Mints `mint_amount` of coins in the `CoinType` currency to Designated Dealer account at
// `designated_dealer_address`. The `tier_index` parameter specifies which tier should be used to
// check verify the off-chain approval policy, and is based in part on the on-chain tier values
// for the specific Designated Dealer, and the number of `CoinType` coins that have been minted to
// the dealer over the past 24 hours. Every Designated Dealer has 4 tiers for each currency that
// they support. The sending `tc_account` must be the Treasury Compliance account, and the
// receiver an authorized Designated Dealer account.
//
// ## Events
// Successful execution of the transaction will emit two events:
// * A `Libra::MintEvent` with the amount and currency code minted is emitted on the
// `mint_event_handle` in the stored `Libra::CurrencyInfo<CoinType>` resource stored under
// `0xA550C18`; and
// * A `DesignatedDealer::ReceivedMintEvent` with the amount, currency code, and Designated
// Dealer's address is emitted on the `mint_event_handle` in the stored `DesignatedDealer::Dealer`
// resource published under the `designated_dealer_address`.
//
// # Parameters
// | Name                        | Type      | Description                                                                                                |
// | ------                      | ------    | -------------                                                                                              |
// | `CoinType`                  | Type      | The Move type for the `CoinType` being minted. `CoinType` must be an already-registered currency on-chain. |
// | `tc_account`                | `&signer` | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account.  |
// | `sliding_nonce`             | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                 |
// | `designated_dealer_address` | `address` | The address of the Designated Dealer account being minted to.                                              |
// | `mint_amount`               | `u64`     | The number of coins to be minted.                                                                          |
// | `tier_index`                | `u64`     | The mint tier index to use for the Designated Dealer account.                                              |
//
// # Common Abort Conditions
// | Error Category                | Error Reason                                 | Description                                                                                                                  |
// | ----------------              | --------------                               | -------------                                                                                                                |
// | `Errors::NOT_PUBLISHED`       | `SlidingNonce::ESLIDING_NONCE`               | A `SlidingNonce` resource is not published under `tc_account`.                                                               |
// | `Errors::INVALID_ARGUMENT`    | `SlidingNonce::ENONCE_TOO_OLD`               | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not.                                   |
// | `Errors::INVALID_ARGUMENT`    | `SlidingNonce::ENONCE_TOO_NEW`               | The `sliding_nonce` is too far in the future.                                                                                |
// | `Errors::INVALID_ARGUMENT`    | `SlidingNonce::ENONCE_ALREADY_RECORDED`      | The `sliding_nonce` has been previously recorded.                                                                            |
// | `Errors::REQUIRES_ADDRESS`    | `CoreAddresses::ETREASURY_COMPLIANCE`        | `tc_account` is not the Treasury Compliance account.                                                                         |
// | `Errors::REQUIRES_ROLE`       | `Roles::ETREASURY_COMPLIANCE`                | `tc_account` is not the Treasury Compliance account.                                                                         |
// | `Errors::INVALID_ARGUMENT`    | `DesignatedDealer::EINVALID_MINT_AMOUNT`     | `mint_amount` is zero.                                                                                                       |
// | `Errors::NOT_PUBLISHED`       | `DesignatedDealer::EDEALER`                  | `DesignatedDealer::Dealer` or `DesignatedDealer::TierInfo<CoinType>` resource does not exist at `designated_dealer_address`. |
// | `Errors::INVALID_ARGUMENT`    | `DesignatedDealer::EINVALID_TIER_INDEX`      | The `tier_index` is out of bounds.                                                                                           |
// | `Errors::INVALID_ARGUMENT`    | `DesignatedDealer::EINVALID_AMOUNT_FOR_TIER` | `mint_amount` exceeds the maximum allowed amount for `tier_index`.                                                           |
// | `Errors::REQUIRES_CAPABILITY` | `Libra::EMINT_CAPABILITY`                    | `tc_account` does not have a `Libra::MintCapability<CoinType>` resource published under it.                                  |
// | `Errors::INVALID_STATE`       | `Libra::EMINTING_NOT_ALLOWED`                | Minting is not currently allowed for `CoinType` coins.                                                                       |
// | `Errors::LIMIT_EXCEEDED`      | `LibraAccount::EDEPOSIT_EXCEEDS_LIMITS`      | The depositing of the funds would exceed the `account`'s account limits.                                                     |
//
// # Related Scripts
// * `Script::create_designated_dealer`
// * `Script::peer_to_peer_with_metadata`
// * `Script::rotate_dual_attestation_info`
type ScriptCall__TieredMint struct {
	CoinType libratypes.TypeTag
	SlidingNonce uint64
	DesignatedDealerAddress libratypes.AccountAddress
	MintAmount uint64
	TierIndex uint64
}

func (*ScriptCall__TieredMint) isScriptCall() {}

// # Summary
// Unfreezes the account at `address`. The sending account of this transaction must be the
// Treasury Compliance account. After the successful execution of this transaction transactions
// may be sent from the previously frozen account, and coins may be sent and received.
//
// # Technical Description
// Sets the `AccountFreezing::FreezingBit` to `false` and emits a
// `AccountFreezing::UnFreezeAccountEvent`. The transaction sender must be the Treasury Compliance
// account. Note that this is a per-account property so unfreezing a Parent VASP will not effect
// the status any of its child accounts and vice versa.
//
// ## Events
// Successful execution of this script will emit a `AccountFreezing::UnFreezeAccountEvent` with
// the `unfrozen_address` set the `to_unfreeze_account`'s address.
//
// # Parameters
// | Name                  | Type      | Description                                                                                               |
// | ------                | ------    | -------------                                                                                             |
// | `tc_account`          | `&signer` | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account. |
// | `sliding_nonce`       | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                |
// | `to_unfreeze_account` | `address` | The account address to be frozen.                                                                         |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                            | Description                                                                                |
// | ----------------           | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `account`.                                |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ETREASURY_COMPLIANCE`   | The sending account is not the Treasury Compliance account.                                |
//
// # Related Scripts
// * `Scripts::freeze_account`
type ScriptCall__UnfreezeAccount struct {
	SlidingNonce uint64
	ToUnfreezeAccount libratypes.AccountAddress
}

func (*ScriptCall__UnfreezeAccount) isScriptCall() {}

// # Summary
// Update the dual attestation limit on-chain. Defined in terms of micro-LBR.  The transaction can
// only be sent by the Treasury Compliance account.  After this transaction all inter-VASP
// payments over this limit must be checked for dual attestation.
//
// # Technical Description
// Updates the `micro_lbr_limit` field of the `DualAttestation::Limit` resource published under
// `0xA550C18`. The amount is set in micro-LBR.
//
// # Parameters
// | Name                  | Type      | Description                                                                                               |
// | ------                | ------    | -------------                                                                                             |
// | `tc_account`          | `&signer` | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account. |
// | `sliding_nonce`       | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                |
// | `new_micro_lbr_limit` | `u64`     | The new dual attestation limit to be used on-chain.                                                       |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                            | Description                                                                                |
// | ----------------           | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `tc_account`.                             |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ETREASURY_COMPLIANCE`   | `tc_account` is not the Treasury Compliance account.                                       |
//
// # Related Scripts
// * `Scripts::update_exchange_rate`
// * `Scripts::update_minting_ability`
type ScriptCall__UpdateDualAttestationLimit struct {
	SlidingNonce uint64
	NewMicroLbrLimit uint64
}

func (*ScriptCall__UpdateDualAttestationLimit) isScriptCall() {}

// # Summary
// Update the rough on-chain exchange rate between a specified currency and LBR (as a conversion
// to micro-LBR). The transaction can only be sent by the Treasury Compliance account. After this
// transaction the updated exchange rate will be used for normalization of gas prices, and for
// dual attestation checking.
//
// # Technical Description
// Updates the on-chain exchange rate from the given `Currency` to micro-LBR.  The exchange rate
// is given by `new_exchange_rate_numerator/new_exchange_rate_denominator`.
//
// # Parameters
// | Name                            | Type      | Description                                                                                                                        |
// | ------                          | ------    | -------------                                                                                                                      |
// | `Currency`                      | Type      | The Move type for the `Currency` whose exchange rate is being updated. `Currency` must be an already-registered currency on-chain. |
// | `tc_account`                    | `&signer` | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account.                          |
// | `sliding_nonce`                 | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for the transaction.                                                          |
// | `new_exchange_rate_numerator`   | `u64`     | The numerator for the new to micro-LBR exchange rate for `Currency`.                                                               |
// | `new_exchange_rate_denominator` | `u64`     | The denominator for the new to micro-LBR exchange rate for `Currency`.                                                             |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                            | Description                                                                                |
// | ----------------           | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `tc_account`.                             |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ETREASURY_COMPLIANCE`   | `tc_account` is not the Treasury Compliance account.                                       |
// | `Errors::REQUIRES_ROLE`    | `Roles::ETREASURY_COMPLIANCE`           | `tc_account` is not the Treasury Compliance account.                                       |
// | `Errors::INVALID_ARGUMENT` | `FixedPoint32::EDENOMINATOR`            | `new_exchange_rate_denominator` is zero.                                                   |
// | `Errors::INVALID_ARGUMENT` | `FixedPoint32::ERATIO_OUT_OF_RANGE`     | The quotient is unrepresentable as a `FixedPoint32`.                                       |
// | `Errors::LIMIT_EXCEEDED`   | `FixedPoint32::ERATIO_OUT_OF_RANGE`     | The quotient is unrepresentable as a `FixedPoint32`.                                       |
//
// # Related Scripts
// * `Scripts::update_dual_attestation_limit`
// * `Scripts::update_minting_ability`
type ScriptCall__UpdateExchangeRate struct {
	Currency libratypes.TypeTag
	SlidingNonce uint64
	NewExchangeRateNumerator uint64
	NewExchangeRateDenominator uint64
}

func (*ScriptCall__UpdateExchangeRate) isScriptCall() {}

// # Summary
// Updates the Libra major version that is stored on-chain and is used by the VM.  This
// transaction can only be sent from the Libra Root account.
//
// # Technical Description
// Updates the `LibraVersion` on-chain config and emits a `LibraConfig::NewEpochEvent` to trigger
// a reconfiguration of the system. The `major` version that is passed in must be strictly greater
// than the current major version held on-chain. The VM reads this information and can use it to
// preserve backwards compatibility with previous major versions of the VM.
//
// # Parameters
// | Name            | Type      | Description                                                                |
// | ------          | ------    | -------------                                                              |
// | `account`       | `&signer` | Signer reference of the sending account. Must be the Libra Root account.   |
// | `sliding_nonce` | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction. |
// | `major`         | `u64`     | The `major` version of the VM to be used from this transaction on.         |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                  | Description                                                                                |
// | ----------------           | --------------                                | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                | A `SlidingNonce` resource is not published under `account`.                                |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`                | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`                | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`       | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ELIBRA_ROOT`                  | `account` is not the Libra Root account.                                                   |
// | `Errors::INVALID_ARGUMENT` | `LibraVersion::EINVALID_MAJOR_VERSION_NUMBER` | `major` is less-than or equal to the current major version stored on-chain.                |
type ScriptCall__UpdateLibraVersion struct {
	SlidingNonce uint64
	Major uint64
}

func (*ScriptCall__UpdateLibraVersion) isScriptCall() {}

// # Summary
// Script to allow or disallow minting of new coins in a specified currency.  This transaction can
// only be sent by the Treasury Compliance account.  Turning minting off for a currency will have
// no effect on coins already in circulation, and coins may still be removed from the system.
//
// # Technical Description
// This transaction sets the `can_mint` field of the `Libra::CurrencyInfo<Currency>` resource
// published under `0xA550C18` to the value of `allow_minting`. Minting of coins if allowed if
// this field is set to `true` and minting of new coins in `Currency` is disallowed otherwise.
// This transaction needs to be sent by the Treasury Compliance account.
//
// # Parameters
// | Name            | Type      | Description                                                                                                                          |
// | ------          | ------    | -------------                                                                                                                        |
// | `Currency`      | Type      | The Move type for the `Currency` whose minting ability is being updated. `Currency` must be an already-registered currency on-chain. |
// | `account`       | `&signer` | Signer reference of the sending account. Must be the Libra Root account.                                                             |
// | `allow_minting` | `bool`    | Whether to allow minting of new coins in `Currency`.                                                                                 |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                          | Description                                          |
// | ----------------           | --------------                        | -------------                                        |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ETREASURY_COMPLIANCE` | `tc_account` is not the Treasury Compliance account. |
// | `Errors::NOT_PUBLISHED`    | `Libra::ECURRENCY_INFO`               | `Currency` is not a registered currency on-chain.    |
//
// # Related Scripts
// * `Scripts::update_dual_attestation_limit`
// * `Scripts::update_exchange_rate`
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
	case *ScriptCall__PeerToPeerWithMetadata:
		return EncodePeerToPeerWithMetadataScript(call.Currency, call.Payee, call.Amount, call.Metadata, call.MetadataSignature)
	case *ScriptCall__Preburn:
		return EncodePreburnScript(call.Token, call.Amount)
	case *ScriptCall__PublishSharedEd25519PublicKey:
		return EncodePublishSharedEd25519PublicKeyScript(call.PublicKey)
	case *ScriptCall__RegisterValidatorConfig:
		return EncodeRegisterValidatorConfigScript(call.ValidatorAccount, call.ConsensusPubkey, call.ValidatorNetworkAddresses, call.FullnodeNetworkAddresses)
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
		return EncodeSetValidatorConfigAndReconfigureScript(call.ValidatorAccount, call.ConsensusPubkey, call.ValidatorNetworkAddresses, call.FullnodeNetworkAddresses)
	case *ScriptCall__SetValidatorOperator:
		return EncodeSetValidatorOperatorScript(call.OperatorName, call.OperatorAccount)
	case *ScriptCall__SetValidatorOperatorWithNonceAdmin:
		return EncodeSetValidatorOperatorWithNonceAdminScript(call.SlidingNonce, call.OperatorName, call.OperatorAccount)
	case *ScriptCall__TieredMint:
		return EncodeTieredMintScript(call.CoinType, call.SlidingNonce, call.DesignatedDealerAddress, call.MintAmount, call.TierIndex)
	case *ScriptCall__UnfreezeAccount:
		return EncodeUnfreezeAccountScript(call.SlidingNonce, call.ToUnfreezeAccount)
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

// # Summary
// Adds a zero `Currency` balance to the sending `account`. This will enable `account` to
// send, receive, and hold `Libra::Libra<Currency>` coins. This transaction can be
// successfully sent by any account that is allowed to hold balances
// (e.g., VASP, Designated Dealer).
//
// # Technical Description
// After the successful execution of this transaction the sending account will have a
// `LibraAccount::Balance<Currency>` resource with zero balance published under it. Only
// accounts that can hold balances can send this transaction, the sending account cannot
// already have a `LibraAccount::Balance<Currency>` published under it.
//
// # Parameters
// | Name       | Type      | Description                                                                                                                                         |
// | ------     | ------    | -------------                                                                                                                                       |
// | `Currency` | Type      | The Move type for the `Currency` being added to the sending account of the transaction. `Currency` must be an already-registered currency on-chain. |
// | `account`  | `&signer` | The signer of the sending account of the transaction.                                                                                               |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                             | Description                                                                |
// | ----------------            | --------------                           | -------------                                                              |
// | `Errors::NOT_PUBLISHED`     | `Libra::ECURRENCY_INFO`                  | The `Currency` is not a registered currency on-chain.                      |
// | `Errors::INVALID_ARGUMENT`  | `LibraAccount::EROLE_CANT_STORE_BALANCE` | The sending `account`'s role does not permit balances.                     |
// | `Errors::ALREADY_PUBLISHED` | `LibraAccount::EADD_EXISTING_CURRENCY`   | A balance for `Currency` is already published under the sending `account`. |
//
// # Related Scripts
// * `Script::create_child_vasp_account`
// * `Script::create_parent_vasp_account`
// * `Script::peer_to_peer_with_metadata`
func EncodeAddCurrencyToAccountScript(currency libratypes.TypeTag) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), add_currency_to_account_code...),
		TyArgs: []libratypes.TypeTag{currency},
		Args: []libratypes.TransactionArgument{},
	}
}

// # Summary
// Stores the sending accounts ability to rotate its authentication key with a designated recovery
// account. Both the sending and recovery accounts need to belong to the same VASP and
// both be VASP accounts. After this transaction both the sending account and the
// specified recovery account can rotate the sender account's authentication key.
//
// # Technical Description
// Adds the `LibraAccount::KeyRotationCapability` for the sending account
// (`to_recover_account`) to the `RecoveryAddress::RecoveryAddress` resource under
// `recovery_address`. After this transaction has been executed successfully the account at
// `recovery_address` and the `to_recover_account` may rotate the authentication key of
// `to_recover_account` (the sender of this transaction).
//
// The sending account of this transaction (`to_recover_account`) must not have previously given away its unique key
// rotation capability, and must be a VASP account. The account at `recovery_address`
// must also be a VASP account belonging to the same VASP as the `to_recover_account`.
// Additionally the account at `recovery_address` must have already initialized itself as
// a recovery account address using the `Script::create_recovery_address` transaction script.
//
// The sending account's (`to_recover_account`) key rotation capability is
// removed in this transaction and stored in the `RecoveryAddress::RecoveryAddress`
// resource stored under the account at `recovery_address`.
//
// # Parameters
// | Name                 | Type      | Description                                                                                                |
// | ------               | ------    | -------------                                                                                              |
// | `to_recover_account` | `&signer` | The signer reference of the sending account of this transaction.                                           |
// | `recovery_address`   | `address` | The account address where the `to_recover_account`'s `LibraAccount::KeyRotationCapability` will be stored. |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                               | Description                                                                                     |
// | ----------------           | --------------                                             | -------------                                                                                   |
// | `Errors::INVALID_STATE`    | `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` | `to_recover_account` has already delegated/extracted its `LibraAccount::KeyRotationCapability`. |
// | `Errors::NOT_PUBLISHED`    | `RecoveryAddress::ERECOVERY_ADDRESS`                       | `recovery_address` does not have a `RecoveryAddress` resource published under it.               |
// | `Errors::INVALID_ARGUMENT` | `RecoveryAddress::EINVALID_KEY_ROTATION_DELEGATION`        | `to_recover_account` and `recovery_address` do not belong to the same VASP.                     |
//
// # Related Scripts
// * `Script::create_recovery_address`
// * `Script::rotate_authentication_key_with_recovery_address`
func EncodeAddRecoveryRotationCapabilityScript(recovery_address libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), add_recovery_rotation_capability_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{recovery_address}},
	}
}

// # Summary
// Adds a script hash to the transaction allowlist. This transaction
// can only be sent by the Libra Root account. Scripts with this hash can be
// sent afterward the successful execution of this script.
//
// # Technical Description
//
// The sending account (`lr_account`) must be the Libra Root account. The script allow
// list must not already hold the script `hash` being added. The `sliding_nonce` must be
// a valid nonce for the Libra Root account. After this transaction has executed
// successfully a reconfiguration will be initiated, and the on-chain config
// `LibraTransactionPublishingOption::LibraTransactionPublishingOption`'s
// `script_allow_list` field will contain the new script `hash` and transactions
// with this `hash` can be successfully sent to the network.
//
// # Parameters
// | Name            | Type         | Description                                                                                     |
// | ------          | ------       | -------------                                                                                   |
// | `lr_account`    | `&signer`    | The signer reference of the sending account of this transaction. Must be the Libra Root signer. |
// | `hash`          | `vector<u8>` | The hash of the script to be added to the script allowlist.                                     |
// | `sliding_nonce` | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                      |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                                           | Description                                                                                |
// | ----------------           | --------------                                                         | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                                         | A `SlidingNonce` resource is not published under `lr_account`.                             |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`                                         | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`                                         | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`                                | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ELIBRA_ROOT`                                           | The sending account is not the Libra Root account.                                         |
// | `Errors::REQUIRES_ROLE`    | `Roles::ELIBRA_ROOT`                                                   | The sending account is not the Libra Root account.                                         |
// | `Errors::INVALID_ARGUMENT` | `LibraTransactionPublishingOption::EINVALID_SCRIPT_HASH`               | The script `hash` is an invalid length.                                                    |
// | `Errors::INVALID_ARGUMENT` | `LibraTransactionPublishingOption::EALLOWLIST_ALREADY_CONTAINS_SCRIPT` | The on-chain allowlist already contains the script `hash`.                                 |
func EncodeAddToScriptAllowListScript(hash []byte, sliding_nonce uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), add_to_script_allow_list_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U8Vector)(&hash), (*libratypes.TransactionArgument__U64)(&sliding_nonce)},
	}
}

// # Summary
// Adds a validator account to the validator set, and triggers a
// reconfiguration of the system to admit the account to the validator set for the system. This
// transaction can only be successfully called by the Libra Root account.
//
// # Technical Description
// This script adds the account at `validator_address` to the validator set.
// This transaction emits a `LibraConfig::NewEpochEvent` event and triggers a
// reconfiguration. Once the reconfiguration triggered by this script's
// execution has been performed, the account at the `validator_address` is
// considered to be a validator in the network.
//
// This transaction script will fail if the `validator_address` address is already in the validator set
// or does not have a `ValidatorConfig::ValidatorConfig` resource already published under it.
//
// # Parameters
// | Name                | Type         | Description                                                                                                                        |
// | ------              | ------       | -------------                                                                                                                      |
// | `lr_account`        | `&signer`    | The signer reference of the sending account of this transaction. Must be the Libra Root signer.                                    |
// | `sliding_nonce`     | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                                         |
// | `validator_name`    | `vector<u8>` | ASCII-encoded human name for the validator. Must match the human name in the `ValidatorConfig::ValidatorConfig` for the validator. |
// | `validator_address` | `address`    | The validator account address to be added to the validator set.                                                                    |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                  | Description                                                                                                                               |
// | ----------------           | --------------                                | -------------                                                                                                                             |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                | A `SlidingNonce` resource is not published under `lr_account`.                                                                            |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`                | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not.                                                |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`                | The `sliding_nonce` is too far in the future.                                                                                             |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`       | The `sliding_nonce` has been previously recorded.                                                                                         |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ELIBRA_ROOT`                  | The sending account is not the Libra Root account.                                                                                        |
// | `Errors::REQUIRES_ROLE`    | `Roles::ELIBRA_ROOT`                          | The sending account is not the Libra Root account.                                                                                        |
// | 0                          | 0                                             | The provided `validator_name` does not match the already-recorded human name for the validator.                                           |
// | `Errors::INVALID_ARGUMENT` | `LibraSystem::EINVALID_PROSPECTIVE_VALIDATOR` | The validator to be added does not have a `ValidatorConfig::ValidatorConfig` resource published under it, or its `config` field is empty. |
// | `Errors::INVALID_ARGUMENT` | `LibraSystem::EALREADY_A_VALIDATOR`           | The `validator_address` account is already a registered validator.                                                                        |
// | `Errors::INVALID_STATE`    | `LibraConfig::EINVALID_BLOCK_TIME`            | An invalid time value was encountered in reconfiguration. Unlikely to occur.                                                              |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::create_validator_operator_account`
// * `Script::register_validator_config`
// * `Script::remove_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::set_validator_config_and_reconfigure`
func EncodeAddValidatorAndReconfigureScript(sliding_nonce uint64, validator_name []byte, validator_address libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), add_validator_and_reconfigure_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), (*libratypes.TransactionArgument__U8Vector)(&validator_name), &libratypes.TransactionArgument__Address{validator_address}},
	}
}

// # Summary
// Burns all coins held in the preburn resource at the specified
// preburn address and removes them from the system. The sending account must
// be the Treasury Compliance account.
// The account that holds the preburn resource will normally be a Designated
// Dealer, but there are no enforced requirements that it be one.
//
// # Technical Description
// This transaction permanently destroys all the coins of `Token` type
// stored in the `Libra::Preburn<Token>` resource published under the
// `preburn_address` account address.
//
// This transaction will only succeed if the sending `account` has a
// `Libra::BurnCapability<Token>`, and a `Libra::Preburn<Token>` resource
// exists under `preburn_address`, with a non-zero `to_burn` field. After the successful execution
// of this transaction the `total_value` field in the
// `Libra::CurrencyInfo<Token>` resource published under `0xA550C18` will be
// decremented by the value of the `to_burn` field of the preburn resource
// under `preburn_address` immediately before this transaction, and the
// `to_burn` field of the preburn resource will have a zero value.
//
// ## Events
// The successful execution of this transaction will emit a `Libra::BurnEvent` on the event handle
// held in the `Libra::CurrencyInfo<Token>` resource's `burn_events` published under
// `0xA550C18`.
//
// # Parameters
// | Name              | Type      | Description                                                                                                                  |
// | ------            | ------    | -------------                                                                                                                |
// | `Token`           | Type      | The Move type for the `Token` currency being burned. `Token` must be an already-registered currency on-chain.                |
// | `tc_account`      | `&signer` | The signer reference of the sending account of this transaction, must have a burn capability for `Token` published under it. |
// | `sliding_nonce`   | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                                   |
// | `preburn_address` | `address` | The address where the coins to-be-burned are currently held.                                                                 |
//
// # Common Abort Conditions
// | Error Category                | Error Reason                            | Description                                                                                           |
// | ----------------              | --------------                          | -------------                                                                                         |
// | `Errors::NOT_PUBLISHED`       | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `account`.                                           |
// | `Errors::INVALID_ARGUMENT`    | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not.            |
// | `Errors::INVALID_ARGUMENT`    | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                                         |
// | `Errors::INVALID_ARGUMENT`    | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                                     |
// | `Errors::REQUIRES_CAPABILITY` | `Libra::EBURN_CAPABILITY`               | The sending `account` does not have a `Libra::BurnCapability<Token>` published under it.              |
// | `Errors::NOT_PUBLISHED`       | `Libra::EPREBURN`                       | The account at `preburn_address` does not have a `Libra::Preburn<Token>` resource published under it. |
// | `Errors::INVALID_STATE`       | `Libra::EPREBURN_EMPTY`                 | The `Libra::Preburn<Token>` resource is empty (has a value of 0).                                     |
// | `Errors::NOT_PUBLISHED`       | `Libra::ECURRENCY_INFO`                 | The specified `Token` is not a registered currency on-chain.                                          |
//
// # Related Scripts
// * `Script::burn_txn_fees`
// * `Script::cancel_burn`
// * `Script::preburn`
func EncodeBurnScript(token libratypes.TypeTag, sliding_nonce uint64, preburn_address libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), burn_code...),
		TyArgs: []libratypes.TypeTag{token},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), &libratypes.TransactionArgument__Address{preburn_address}},
	}
}

// # Summary
// Burns the transaction fees collected in the `CoinType` currency so that the
// Libra association may reclaim the backing coins off-chain. May only be sent
// by the Treasury Compliance account.
//
// # Technical Description
// Burns the transaction fees collected in `CoinType` so that the
// association may reclaim the backing coins. Once this transaction has executed
// successfully all transaction fees that will have been collected in
// `CoinType` since the last time this script was called with that specific
// currency. Both `balance` and `preburn` fields in the
// `TransactionFee::TransactionFee<CoinType>` resource published under the `0xB1E55ED`
// account address will have a value of 0 after the successful execution of this script.
//
// ## Events
// The successful execution of this transaction will emit a `Libra::BurnEvent` on the event handle
// held in the `Libra::CurrencyInfo<CoinType>` resource's `burn_events` published under
// `0xA550C18`.
//
// # Parameters
// | Name         | Type      | Description                                                                                                                                         |
// | ------       | ------    | -------------                                                                                                                                       |
// | `CoinType`   | Type      | The Move type for the `CoinType` being added to the sending account of the transaction. `CoinType` must be an already-registered currency on-chain. |
// | `tc_account` | `&signer` | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account.                                           |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                          | Description                                                 |
// | ----------------           | --------------                        | -------------                                               |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ETREASURY_COMPLIANCE` | The sending account is not the Treasury Compliance account. |
// | `Errors::NOT_PUBLISHED`    | `TransactionFee::ETRANSACTION_FEE`    | `CoinType` is not an accepted transaction fee currency.     |
// | `Errors::INVALID_ARGUMENT` | `Libra::ECOIN`                        | The collected fees in `CoinType` are zero.                  |
//
// # Related Scripts
// * `Script::burn`
// * `Script::cancel_burn`
func EncodeBurnTxnFeesScript(coin_type libratypes.TypeTag) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), burn_txn_fees_code...),
		TyArgs: []libratypes.TypeTag{coin_type},
		Args: []libratypes.TransactionArgument{},
	}
}

// # Summary
// Cancels and returns all coins held in the preburn area under
// `preburn_address` and returns the funds to the `preburn_address`'s balance.
// Can only be successfully sent by an account with Treasury Compliance role.
//
// # Technical Description
// Cancels and returns all coins held in the `Libra::Preburn<Token>` resource under the `preburn_address` and
// return the funds to the `preburn_address` account's `LibraAccount::Balance<Token>`.
// The transaction must be sent by an `account` with a `Libra::BurnCapability<Token>`
// resource published under it. The account at `preburn_address` must have a
// `Libra::Preburn<Token>` resource published under it, and its value must be nonzero. The transaction removes
// the entire balance held in the `Libra::Preburn<Token>` resource, and returns it back to the account's
// `LibraAccount::Balance<Token>` under `preburn_address`. Due to this, the account at
// `preburn_address` must already have a balance in the `Token` currency published
// before this script is called otherwise the transaction will fail.
//
// ## Events
// The successful execution of this transaction will emit:
// * A `Libra::CancelBurnEvent` on the event handle held in the `Libra::CurrencyInfo<Token>`
// resource's `burn_events` published under `0xA550C18`.
// * A `LibraAccount::ReceivedPaymentEvent` on the `preburn_address`'s
// `LibraAccount::LibraAccount` `received_events` event handle with both the `payer` and `payee`
// being `preburn_address`.
//
// # Parameters
// | Name              | Type      | Description                                                                                                                          |
// | ------            | ------    | -------------                                                                                                                        |
// | `Token`           | Type      | The Move type for the `Token` currenty that burning is being cancelled for. `Token` must be an already-registered currency on-chain. |
// | `account`         | `&signer` | The signer reference of the sending account of this transaction, must have a burn capability for `Token` published under it.         |
// | `preburn_address` | `address` | The address where the coins to-be-burned are currently held.                                                                         |
//
// # Common Abort Conditions
// | Error Category                | Error Reason                                     | Description                                                                                           |
// | ----------------              | --------------                                   | -------------                                                                                         |
// | `Errors::REQUIRES_CAPABILITY` | `Libra::EBURN_CAPABILITY`                        | The sending `account` does not have a `Libra::BurnCapability<Token>` published under it.              |
// | `Errors::NOT_PUBLISHED`       | `Libra::EPREBURN`                                | The account at `preburn_address` does not have a `Libra::Preburn<Token>` resource published under it. |
// | `Errors::NOT_PUBLISHED`       | `Libra::ECURRENCY_INFO`                          | The specified `Token` is not a registered currency on-chain.                                          |
// | `Errors::INVALID_ARGUMENT`    | `LibraAccount::ECOIN_DEPOSIT_IS_ZERO`            | The value held in the preburn resource was zero.                                                      |
// | `Errors::INVALID_ARGUMENT`    | `LibraAccount::EPAYEE_CANT_ACCEPT_CURRENCY_TYPE` | The account at `preburn_address` doesn't have a balance resource for `Token`.                         |
// | `Errors::LIMIT_EXCEEDED`      | `LibraAccount::EDEPOSIT_EXCEEDS_LIMITS`          | The depositing of the funds held in the prebun area would exceed the `account`'s account limits.      |
// | `Errors::INVALID_STATE`       | `DualAttestation::EPAYEE_COMPLIANCE_KEY_NOT_SET` | The `account` does not have a compliance key set on it but dual attestion checking was performed.     |
//
// # Related Scripts
// * `Script::burn_txn_fees`
// * `Script::burn`
// * `Script::preburn`
func EncodeCancelBurnScript(token libratypes.TypeTag, preburn_address libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), cancel_burn_code...),
		TyArgs: []libratypes.TypeTag{token},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{preburn_address}},
	}
}

// # Summary
// Creates a Child VASP account with its parent being the sending account of the transaction.
// The sender of the transaction must be a Parent VASP account.
//
// # Technical Description
// Creates a `ChildVASP` account for the sender `parent_vasp` at `child_address` with a balance of
// `child_initial_balance` in `CoinType` and an initial authentication key of
// `auth_key_prefix | child_address`.
//
// If `add_all_currencies` is true, the child address will have a zero balance in all available
// currencies in the system.
//
// The new account will be a child account of the transaction sender, which must be a
// Parent VASP account. The child account will be recorded against the limit of
// child accounts of the creating Parent VASP account.
//
// ## Events
// Successful execution with a `child_initial_balance` greater than zero will emit:
// * A `LibraAccount::SentPaymentEvent` with the `payer` field being the Parent VASP's address,
// and payee field being `child_address`. This is emitted on the Parent VASP's
// `LibraAccount::LibraAccount` `sent_events` handle.
// * A `LibraAccount::ReceivedPaymentEvent` with the  `payer` field being the Parent VASP's address,
// and payee field being `child_address`. This is emitted on the new Child VASPS's
// `LibraAccount::LibraAccount` `received_events` handle.
//
// # Parameters
// | Name                    | Type         | Description                                                                                                                                 |
// | ------                  | ------       | -------------                                                                                                                               |
// | `CoinType`              | Type         | The Move type for the `CoinType` that the child account should be created with. `CoinType` must be an already-registered currency on-chain. |
// | `parent_vasp`           | `&signer`    | The signer reference of the sending account. Must be a Parent VASP account.                                                                 |
// | `child_address`         | `address`    | Address of the to-be-created Child VASP account.                                                                                            |
// | `auth_key_prefix`       | `vector<u8>` | The authentication key prefix that will be used initially for the newly created account.                                                    |
// | `add_all_currencies`    | `bool`       | Whether to publish balance resources for all known currencies when the account is created.                                                  |
// | `child_initial_balance` | `u64`        | The initial balance in `CoinType` to give the child account when it's created.                                                              |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                                             | Description                                                                              |
// | ----------------            | --------------                                           | -------------                                                                            |
// | `Errors::INVALID_ARGUMENT`  | `LibraAccount::EMALFORMED_AUTHENTICATION_KEY`            | The `auth_key_prefix` was not of length 32.                                              |
// | `Errors::REQUIRES_ROLE`     | `Roles::EPARENT_VASP`                                    | The sending account wasn't a Parent VASP account.                                        |
// | `Errors::ALREADY_PUBLISHED` | `Roles::EROLE_ID`                                        | The `child_address` address is already taken.                                            |
// | `Errors::LIMIT_EXCEEDED`    | `VASP::ETOO_MANY_CHILDREN`                               | The sending account has reached the maximum number of allowed child accounts.            |
// | `Errors::NOT_PUBLISHED`     | `Libra::ECURRENCY_INFO`                                  | The `CoinType` is not a registered currency on-chain.                                    |
// | `Errors::INVALID_STATE`     | `LibraAccount::EWITHDRAWAL_CAPABILITY_ALREADY_EXTRACTED` | The withdrawal capability for the sending account has already been extracted.            |
// | `Errors::NOT_PUBLISHED`     | `LibraAccount::EPAYER_DOESNT_HOLD_CURRENCY`              | The sending account doesn't have a balance in `CoinType`.                                |
// | `Errors::LIMIT_EXCEEDED`    | `LibraAccount::EINSUFFICIENT_BALANCE`                    | The sending account doesn't have at least `child_initial_balance` of `CoinType` balance. |
// | `Errors::INVALID_ARGUMENT`  | `LibraAccount::ECANNOT_CREATE_AT_VM_RESERVED`            | The `child_address` is the reserved address 0x0.                                         |
//
// # Related Scripts
// * `Script::create_parent_vasp_account`
// * `Script::add_currency_to_account`
// * `Script::rotate_authentication_key`
// * `Script::add_recovery_rotation_capability`
// * `Script::create_recovery_address`
func EncodeCreateChildVaspAccountScript(coin_type libratypes.TypeTag, child_address libratypes.AccountAddress, auth_key_prefix []byte, add_all_currencies bool, child_initial_balance uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), create_child_vasp_account_code...),
		TyArgs: []libratypes.TypeTag{coin_type},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{child_address}, (*libratypes.TransactionArgument__U8Vector)(&auth_key_prefix), (*libratypes.TransactionArgument__Bool)(&add_all_currencies), (*libratypes.TransactionArgument__U64)(&child_initial_balance)},
	}
}

// # Summary
// Creates a Designated Dealer account with the provided information, and initializes it with
// default mint tiers. The transaction can only be sent by the Treasury Compliance account.
//
// # Technical Description
// Creates an account with the Designated Dealer role at `addr` with authentication key
// `auth_key_prefix` | `addr` and a 0 balance of type `Currency`. If `add_all_currencies` is true,
// 0 balances for all available currencies in the system will also be added. This can only be
// invoked by an account with the TreasuryCompliance role.
//
// At the time of creation the account is also initialized with default mint tiers of (500_000,
// 5000_000, 50_000_000, 500_000_000), and preburn areas for each currency that is added to the
// account.
//
// # Parameters
// | Name                 | Type         | Description                                                                                                                                         |
// | ------               | ------       | -------------                                                                                                                                       |
// | `Currency`           | Type         | The Move type for the `Currency` that the Designated Dealer should be initialized with. `Currency` must be an already-registered currency on-chain. |
// | `tc_account`         | `&signer`    | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account.                                           |
// | `sliding_nonce`      | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                                                          |
// | `addr`               | `address`    | Address of the to-be-created Designated Dealer account.                                                                                             |
// | `auth_key_prefix`    | `vector<u8>` | The authentication key prefix that will be used initially for the newly created account.                                                            |
// | `human_name`         | `vector<u8>` | ASCII-encoded human name for the Designated Dealer.                                                                                                 |
// | `add_all_currencies` | `bool`       | Whether to publish preburn, balance, and tier info resources for all known (SCS) currencies or just `Currency` when the account is created.         |
//

// # Common Abort Conditions
// | Error Category              | Error Reason                            | Description                                                                                |
// | ----------------            | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`     | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `tc_account`.                             |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS`  | `CoreAddresses::ETREASURY_COMPLIANCE`   | The sending account is not the Treasury Compliance account.                                |
// | `Errors::REQUIRES_ROLE`     | `Roles::ETREASURY_COMPLIANCE`           | The sending account is not the Treasury Compliance account.                                |
// | `Errors::NOT_PUBLISHED`     | `Libra::ECURRENCY_INFO`                 | The `Currency` is not a registered currency on-chain.                                      |
// | `Errors::ALREADY_PUBLISHED` | `Roles::EROLE_ID`                       | The `addr` address is already taken.                                                       |
//
// # Related Scripts
// * `Script::tiered_mint`
// * `Script::peer_to_peer_with_metadata`
// * `Script::rotate_dual_attestation_info`
func EncodeCreateDesignatedDealerScript(currency libratypes.TypeTag, sliding_nonce uint64, addr libratypes.AccountAddress, auth_key_prefix []byte, human_name []byte, add_all_currencies bool) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), create_designated_dealer_code...),
		TyArgs: []libratypes.TypeTag{currency},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), &libratypes.TransactionArgument__Address{addr}, (*libratypes.TransactionArgument__U8Vector)(&auth_key_prefix), (*libratypes.TransactionArgument__U8Vector)(&human_name), (*libratypes.TransactionArgument__Bool)(&add_all_currencies)},
	}
}

// # Summary
// Creates a Parent VASP account with the specified human name. Must be called by the Treasury Compliance account.
//
// # Technical Description
// Creates an account with the Parent VASP role at `address` with authentication key
// `auth_key_prefix` | `new_account_address` and a 0 balance of type `CoinType`. If
// `add_all_currencies` is true, 0 balances for all available currencies in the system will
// also be added. This can only be invoked by an TreasuryCompliance account.
// `sliding_nonce` is a unique nonce for operation, see `SlidingNonce` for details.
//
// # Parameters
// | Name                  | Type         | Description                                                                                                                                                    |
// | ------                | ------       | -------------                                                                                                                                                  |
// | `CoinType`            | Type         | The Move type for the `CoinType` currency that the Parent VASP account should be initialized with. `CoinType` must be an already-registered currency on-chain. |
// | `tc_account`          | `&signer`    | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account.                                                      |
// | `sliding_nonce`       | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                                                                     |
// | `new_account_address` | `address`    | Address of the to-be-created Parent VASP account.                                                                                                              |
// | `auth_key_prefix`     | `vector<u8>` | The authentication key prefix that will be used initially for the newly created account.                                                                       |
// | `human_name`          | `vector<u8>` | ASCII-encoded human name for the Parent VASP.                                                                                                                  |
// | `add_all_currencies`  | `bool`       | Whether to publish balance resources for all known currencies when the account is created.                                                                     |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                            | Description                                                                                |
// | ----------------            | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`     | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `tc_account`.                             |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS`  | `CoreAddresses::ETREASURY_COMPLIANCE`   | The sending account is not the Treasury Compliance account.                                |
// | `Errors::REQUIRES_ROLE`     | `Roles::ETREASURY_COMPLIANCE`           | The sending account is not the Treasury Compliance account.                                |
// | `Errors::NOT_PUBLISHED`     | `Libra::ECURRENCY_INFO`                 | The `CoinType` is not a registered currency on-chain.                                      |
// | `Errors::ALREADY_PUBLISHED` | `Roles::EROLE_ID`                       | The `new_account_address` address is already taken.                                        |
//
// # Related Scripts
// * `Script::create_child_vasp_account`
// * `Script::add_currency_to_account`
// * `Script::rotate_authentication_key`
// * `Script::add_recovery_rotation_capability`
// * `Script::create_recovery_address`
// * `Script::rotate_dual_attestation_info`
func EncodeCreateParentVaspAccountScript(coin_type libratypes.TypeTag, sliding_nonce uint64, new_account_address libratypes.AccountAddress, auth_key_prefix []byte, human_name []byte, add_all_currencies bool) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), create_parent_vasp_account_code...),
		TyArgs: []libratypes.TypeTag{coin_type},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), &libratypes.TransactionArgument__Address{new_account_address}, (*libratypes.TransactionArgument__U8Vector)(&auth_key_prefix), (*libratypes.TransactionArgument__U8Vector)(&human_name), (*libratypes.TransactionArgument__Bool)(&add_all_currencies)},
	}
}

// # Summary
// Initializes the sending account as a recovery address that may be used by
// the VASP that it belongs to. The sending account must be a VASP account.
// Multiple recovery addresses can exist for a single VASP, but accounts in
// each must be disjoint.
//
// # Technical Description
// Publishes a `RecoveryAddress::RecoveryAddress` resource under `account`. It then
// extracts the `LibraAccount::KeyRotationCapability` for `account` and adds
// it to the resource. After the successful execution of this transaction
// other accounts may add their key rotation to this resource so that `account`
// may be used as a recovery account for those accounts.
//
// # Parameters
// | Name      | Type      | Description                                           |
// | ------    | ------    | -------------                                         |
// | `account` | `&signer` | The signer of the sending account of the transaction. |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                                               | Description                                                                                   |
// | ----------------            | --------------                                             | -------------                                                                                 |
// | `Errors::INVALID_STATE`     | `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` | `account` has already delegated/extracted its `LibraAccount::KeyRotationCapability`.          |
// | `Errors::INVALID_ARGUMENT`  | `RecoveryAddress::ENOT_A_VASP`                             | `account` is not a VASP account.                                                              |
// | `Errors::INVALID_ARGUMENT`  | `RecoveryAddress::EKEY_ROTATION_DEPENDENCY_CYCLE`          | A key rotation recovery cycle would be created by adding `account`'s key rotation capability. |
// | `Errors::ALREADY_PUBLISHED` | `RecoveryAddress::ERECOVERY_ADDRESS`                       | A `RecoveryAddress::RecoveryAddress` resource has already been published under `account`.     |
//
// # Related Scripts
// * `Script::add_recovery_rotation_capability`
// * `Script::rotate_authentication_key_with_recovery_address`
func EncodeCreateRecoveryAddressScript() libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), create_recovery_address_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{},
	}
}

// # Summary
// Creates a Validator account. This transaction can only be sent by the Libra
// Root account.
//
// # Technical Description
// Creates an account with a Validator role at `new_account_address`, with authentication key
// `auth_key_prefix` | `new_account_address`. It publishes a
// `ValidatorConfig::ValidatorConfig` resource with empty `config`, and
// `operator_account` fields. The `human_name` field of the
// `ValidatorConfig::ValidatorConfig` is set to the passed in `human_name`.
// This script does not add the validator to the validator set or the system,
// but only creates the account.
//
// # Parameters
// | Name                  | Type         | Description                                                                                     |
// | ------                | ------       | -------------                                                                                   |
// | `lr_account`          | `&signer`    | The signer reference of the sending account of this transaction. Must be the Libra Root signer. |
// | `sliding_nonce`       | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                      |
// | `new_account_address` | `address`    | Address of the to-be-created Validator account.                                                 |
// | `auth_key_prefix`     | `vector<u8>` | The authentication key prefix that will be used initially for the newly created account.        |
// | `human_name`          | `vector<u8>` | ASCII-encoded human name for the validator.                                                     |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                            | Description                                                                                |
// | ----------------            | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`     | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `lr_account`.                             |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS`  | `CoreAddresses::ELIBRA_ROOT`            | The sending account is not the Libra Root account.                                         |
// | `Errors::REQUIRES_ROLE`     | `Roles::ELIBRA_ROOT`                    | The sending account is not the Libra Root account.                                         |
// | `Errors::ALREADY_PUBLISHED` | `Roles::EROLE_ID`                       | The `new_account_address` address is already taken.                                        |
//
// # Related Scripts
// * `Script::add_validator_and_reconfigure`
// * `Script::create_validator_operator_account`
// * `Script::register_validator_config`
// * `Script::remove_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::set_validator_config_and_reconfigure`
func EncodeCreateValidatorAccountScript(sliding_nonce uint64, new_account_address libratypes.AccountAddress, auth_key_prefix []byte, human_name []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), create_validator_account_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), &libratypes.TransactionArgument__Address{new_account_address}, (*libratypes.TransactionArgument__U8Vector)(&auth_key_prefix), (*libratypes.TransactionArgument__U8Vector)(&human_name)},
	}
}

// # Summary
// Creates a Validator Operator account. This transaction can only be sent by the Libra
// Root account.
//
// # Technical Description
// Creates an account with a Validator Operator role at `new_account_address`, with authentication key
// `auth_key_prefix` | `new_account_address`. It publishes a
// `ValidatorOperatorConfig::ValidatorOperatorConfig` resource with the specified `human_name`.
// This script does not assign the validator operator to any validator accounts but only creates the account.
//
// # Parameters
// | Name                  | Type         | Description                                                                                     |
// | ------                | ------       | -------------                                                                                   |
// | `lr_account`          | `&signer`    | The signer reference of the sending account of this transaction. Must be the Libra Root signer. |
// | `sliding_nonce`       | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                      |
// | `new_account_address` | `address`    | Address of the to-be-created Validator account.                                                 |
// | `auth_key_prefix`     | `vector<u8>` | The authentication key prefix that will be used initially for the newly created account.        |
// | `human_name`          | `vector<u8>` | ASCII-encoded human name for the validator.                                                     |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                            | Description                                                                                |
// | ----------------            | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`     | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `lr_account`.                             |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT`  | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS`  | `CoreAddresses::ELIBRA_ROOT`            | The sending account is not the Libra Root account.                                         |
// | `Errors::REQUIRES_ROLE`     | `Roles::ELIBRA_ROOT`                    | The sending account is not the Libra Root account.                                         |
// | `Errors::ALREADY_PUBLISHED` | `Roles::EROLE_ID`                       | The `new_account_address` address is already taken.                                        |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::add_validator_and_reconfigure`
// * `Script::register_validator_config`
// * `Script::remove_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::set_validator_config_and_reconfigure`
func EncodeCreateValidatorOperatorAccountScript(sliding_nonce uint64, new_account_address libratypes.AccountAddress, auth_key_prefix []byte, human_name []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), create_validator_operator_account_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), &libratypes.TransactionArgument__Address{new_account_address}, (*libratypes.TransactionArgument__U8Vector)(&auth_key_prefix), (*libratypes.TransactionArgument__U8Vector)(&human_name)},
	}
}

// # Summary
// Freezes the account at `address`. The sending account of this transaction
// must be the Treasury Compliance account. The account being frozen cannot be
// the Libra Root or Treasury Compliance account. After the successful
// execution of this transaction no transactions may be sent from the frozen
// account, and the frozen account may not send or receive coins.
//
// # Technical Description
// Sets the `AccountFreezing::FreezingBit` to `true` and emits a
// `AccountFreezing::FreezeAccountEvent`. The transaction sender must be the
// Treasury Compliance account, but the account at `to_freeze_account` must
// not be either `0xA550C18` (the Libra Root address), or `0xB1E55ED` (the
// Treasury Compliance address). Note that this is a per-account property
// e.g., freezing a Parent VASP will not effect the status any of its child
// accounts and vice versa.
//

// ## Events
// Successful execution of this transaction will emit a `AccountFreezing::FreezeAccountEvent` on
// the `freeze_event_handle` held in the `AccountFreezing::FreezeEventsHolder` resource published
// under `0xA550C18` with the `frozen_address` being the `to_freeze_account`.
//
// # Parameters
// | Name                | Type      | Description                                                                                               |
// | ------              | ------    | -------------                                                                                             |
// | `tc_account`        | `&signer` | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account. |
// | `sliding_nonce`     | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                |
// | `to_freeze_account` | `address` | The account address to be frozen.                                                                         |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                 | Description                                                                                |
// | ----------------           | --------------                               | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`               | A `SlidingNonce` resource is not published under `tc_account`.                             |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`               | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`               | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`      | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ETREASURY_COMPLIANCE`        | The sending account is not the Treasury Compliance account.                                |
// | `Errors::REQUIRES_ROLE`    | `Roles::ETREASURY_COMPLIANCE`                | The sending account is not the Treasury Compliance account.                                |
// | `Errors::INVALID_ARGUMENT` | `AccountFreezing::ECANNOT_FREEZE_TC`         | `to_freeze_account` was the Treasury Compliance account (`0xB1E55ED`).                     |
// | `Errors::INVALID_ARGUMENT` | `AccountFreezing::ECANNOT_FREEZE_LIBRA_ROOT` | `to_freeze_account` was the Libra Root account (`0xA550C18`).                              |
//
// # Related Scripts
// * `Scripts::unfreeze_account`
func EncodeFreezeAccountScript(sliding_nonce uint64, to_freeze_account libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), freeze_account_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), &libratypes.TransactionArgument__Address{to_freeze_account}},
	}
}

// # Summary
// Transfers a given number of coins in a specified currency from one account to another.
// Transfers over a specified amount defined on-chain that are between two different VASPs, or
// other accounts that have opted-in will be subject to on-chain checks to ensure the receiver has
// agreed to receive the coins.  This transaction can be sent by any account that can hold a
// balance, and to any account that can hold a balance. Both accounts must hold balances in the
// currency being transacted.
//
// # Technical Description
//
// Transfers `amount` coins of type `Currency` from `payer` to `payee` with (optional) associated
// `metadata` and an (optional) `metadata_signature` on the message
// `metadata` | `Signer::address_of(payer)` | `amount` | `DualAttestation::DOMAIN_SEPARATOR`.
// The `metadata` and `metadata_signature` parameters are only required if `amount` >=
// `DualAttestation::get_cur_microlibra_limit` LBR and `payer` and `payee` are distinct VASPs.
// However, a transaction sender can opt in to dual attestation even when it is not required
// (e.g., a DesignatedDealer -> VASP payment) by providing a non-empty `metadata_signature`.
// Standardized `metadata` LCS format can be found in `libra_types::transaction::metadata::Metadata`.
//
// ## Events
// Successful execution of this script emits two events:
// * A `LibraAccount::SentPaymentEvent` on `payer`'s `LibraAccount::LibraAccount` `sent_events` handle; and
// * A `LibraAccount::ReceivedPaymentEvent` on `payee`'s `LibraAccount::LibraAccount` `received_events` handle.
//
// # Parameters
// | Name                 | Type         | Description                                                                                                                  |
// | ------               | ------       | -------------                                                                                                                |
// | `Currency`           | Type         | The Move type for the `Currency` being sent in this transaction. `Currency` must be an already-registered currency on-chain. |
// | `payer`              | `&signer`    | The signer reference of the sending account that coins are being transferred from.                                           |
// | `payee`              | `address`    | The address of the account the coins are being transferred to.                                                               |
// | `metadata`           | `vector<u8>` | Optional metadata about this payment.                                                                                        |
// | `metadata_signature` | `vector<u8>` | Optional signature over `metadata` and payment information. See                                                              |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                     | Description                                                                                                                         |
// | ----------------           | --------------                                   | -------------                                                                                                                       |
// | `Errors::NOT_PUBLISHED`    | `LibraAccount::EPAYER_DOESNT_HOLD_CURRENCY`      | `payer` doesn't hold a balance in `Currency`.                                                                                       |
// | `Errors::LIMIT_EXCEEDED`   | `LibraAccount::EINSUFFICIENT_BALANCE`            | `amount` is greater than `payer`'s balance in `Currency`.                                                                           |
// | `Errors::INVALID_ARGUMENT` | `LibraAccount::ECOIN_DEPOSIT_IS_ZERO`            | `amount` is zero.                                                                                                                   |
// | `Errors::NOT_PUBLISHED`    | `LibraAccount::EPAYEE_DOES_NOT_EXIST`            | No account exists at the `payee` address.                                                                                           |
// | `Errors::INVALID_ARGUMENT` | `LibraAccount::EPAYEE_CANT_ACCEPT_CURRENCY_TYPE` | An account exists at `payee`, but it does not accept payments in `Currency`.                                                        |
// | `Errors::INVALID_STATE`    | `AccountFreezing::EACCOUNT_FROZEN`               | The `payee` account is frozen.                                                                                                      |
// | `Errors::INVALID_ARGUMENT` | `DualAttestation::EMALFORMED_METADATA_SIGNATURE` | `metadata_signature` is not 64 bytes.                                                                                               |
// | `Errors::INVALID_ARGUMENT` | `DualAttestation::EINVALID_METADATA_SIGNATURE`   | `metadata_signature` does not verify on the against the `payee'`s `DualAttestation::Credential` `compliance_public_key` public key. |
// | `Errors::LIMIT_EXCEEDED`   | `LibraAccount::EWITHDRAWAL_EXCEEDS_LIMITS`       | `payer` has exceeded its daily withdrawal limits for the backing coins of LBR.                                                      |
// | `Errors::LIMIT_EXCEEDED`   | `LibraAccount::EDEPOSIT_EXCEEDS_LIMITS`          | `payee` has exceeded its daily deposit limits for LBR.                                                                              |
//
// # Related Scripts
// * `Script::create_child_vasp_account`
// * `Script::create_parent_vasp_account`
// * `Script::add_currency_to_account`
func EncodePeerToPeerWithMetadataScript(currency libratypes.TypeTag, payee libratypes.AccountAddress, amount uint64, metadata []byte, metadata_signature []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), peer_to_peer_with_metadata_code...),
		TyArgs: []libratypes.TypeTag{currency},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{payee}, (*libratypes.TransactionArgument__U64)(&amount), (*libratypes.TransactionArgument__U8Vector)(&metadata), (*libratypes.TransactionArgument__U8Vector)(&metadata_signature)},
	}
}

// # Summary
// Moves a specified number of coins in a given currency from the account's
// balance to its preburn area after which the coins may be burned. This
// transaction may be sent by any account that holds a balance and preburn area
// in the specified currency.
//
// # Technical Description
// Moves the specified `amount` of coins in `Token` currency from the sending `account`'s
// `LibraAccount::Balance<Token>` to the `Libra::Preburn<Token>` published under the same
// `account`. `account` must have both of these resources published under it at the start of this
// transaction in order for it to execute successfully.
//
// ## Events
// Successful execution of this script emits two events:
// * `LibraAccount::SentPaymentEvent ` on `account`'s `LibraAccount::LibraAccount` `sent_events`
// handle with the `payee` and `payer` fields being `account`'s address; and
// * A `Libra::PreburnEvent` with `Token`'s currency code on the
// `Libra::CurrencyInfo<Token`'s `preburn_events` handle for `Token` and with
// `preburn_address` set to `account`'s address.
//
// # Parameters
// | Name      | Type      | Description                                                                                                                      |
// | ------    | ------    | -------------                                                                                                                    |
// | `Token`   | Type      | The Move type for the `Token` currency being moved to the preburn area. `Token` must be an already-registered currency on-chain. |
// | `account` | `&signer` | The signer reference of the sending account.                                                                                     |
// | `amount`  | `u64`     | The amount in `Token` to be moved to the preburn area.                                                                           |
//
// # Common Abort Conditions
// | Error Category           | Error Reason                                             | Description                                                                             |
// | ----------------         | --------------                                           | -------------                                                                           |
// | `Errors::NOT_PUBLISHED`  | `Libra::ECURRENCY_INFO`                                  | The `Token` is not a registered currency on-chain.                                      |
// | `Errors::INVALID_STATE`  | `LibraAccount::EWITHDRAWAL_CAPABILITY_ALREADY_EXTRACTED` | The withdrawal capability for `account` has already been extracted.                     |
// | `Errors::LIMIT_EXCEEDED` | `LibraAccount::EINSUFFICIENT_BALANCE`                    | `amount` is greater than `payer`'s balance in `Token`.                                  |
// | `Errors::NOT_PUBLISHED`  | `LibraAccount::EPAYER_DOESNT_HOLD_CURRENCY`              | `account` doesn't hold a balance in `Token`.                                            |
// | `Errors::NOT_PUBLISHED`  | `Libra::EPREBURN`                                        | `account` doesn't have a `Libra::Preburn<Token>` resource published under it.           |
// | `Errors::INVALID_STATE`  | `Libra::EPREBURN_OCCUPIED`                               | The `value` field in the `Libra::Preburn<Token>` resource under the sender is non-zero. |
//
// # Related Scripts
// * `Script::cancel_burn`
// * `Script::burn`
// * `Script::burn_txn_fees`
func EncodePreburnScript(token libratypes.TypeTag, amount uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), preburn_code...),
		TyArgs: []libratypes.TypeTag{token},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&amount)},
	}
}

// # Summary
// Rotates the authentication key of the sending account to the
// newly-specified public key and publishes a new shared authentication key
// under the sender's account. Any account can send this transaction.
//
// # Technical Description
// Rotates the authentication key of the sending account to `public_key`,
// and publishes a `SharedEd25519PublicKey::SharedEd25519PublicKey` resource
// containing the 32-byte ed25519 `public_key` and the `LibraAccount::KeyRotationCapability` for
// `account` under `account`.
//
// # Parameters
// | Name         | Type         | Description                                                                               |
// | ------       | ------       | -------------                                                                             |
// | `account`    | `&signer`    | The signer reference of the sending account of the transaction.                           |
// | `public_key` | `vector<u8>` | 32-byte Ed25519 public key for `account`' authentication key to be rotated to and stored. |
//
// # Common Abort Conditions
// | Error Category              | Error Reason                                               | Description                                                                                         |
// | ----------------            | --------------                                             | -------------                                                                                       |
// | `Errors::INVALID_STATE`     | `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` | `account` has already delegated/extracted its `LibraAccount::KeyRotationCapability` resource.       |
// | `Errors::ALREADY_PUBLISHED` | `SharedEd25519PublicKey::ESHARED_KEY`                      | The `SharedEd25519PublicKey::SharedEd25519PublicKey` resource is already published under `account`. |
// | `Errors::INVALID_ARGUMENT`  | `SharedEd25519PublicKey::EMALFORMED_PUBLIC_KEY`            | `public_key` is an invalid ed25519 public key.                                                      |
//
// # Related Scripts
// * `Script::rotate_shared_ed25519_public_key`
func EncodePublishSharedEd25519PublicKeyScript(public_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), publish_shared_ed25519_public_key_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U8Vector)(&public_key)},
	}
}

// # Summary
// Updates a validator's configuration. This does not reconfigure the system and will not update
// the configuration in the validator set that is seen by other validators in the network. Can
// only be successfully sent by a Validator Operator account that is already registered with a
// validator.
//
// # Technical Description
// This updates the fields with corresponding names held in the `ValidatorConfig::ValidatorConfig`
// config resource held under `validator_account`. It does not emit a `LibraConfig::NewEpochEvent`
// so the copy of this config held in the validator set will not be updated, and the changes are
// only "locally" under the `validator_account` account address.
//
// # Parameters
// | Name                          | Type         | Description                                                                                                                  |
// | ------                        | ------       | -------------                                                                                                                |
// | `validator_operator_account`  | `&signer`    | Signer reference of the sending account. Must be the registered validator operator for the validator at `validator_address`. |
// | `validator_account`           | `address`    | The address of the validator's `ValidatorConfig::ValidatorConfig` resource being updated.                                    |
// | `consensus_pubkey`            | `vector<u8>` | New Ed25519 public key to be used in the updated `ValidatorConfig::ValidatorConfig`.                                         |
// | `validator_network_addresses` | `vector<u8>` | New set of `validator_network_addresses` to be used in the updated `ValidatorConfig::ValidatorConfig`.                       |
// | `fullnode_network_addresses`  | `vector<u8>` | New set of `fullnode_network_addresses` to be used in the updated `ValidatorConfig::ValidatorConfig`.                        |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                   | Description                                                                                           |
// | ----------------           | --------------                                 | -------------                                                                                         |
// | `Errors::NOT_PUBLISHED`    | `ValidatorConfig::EVALIDATOR_CONFIG`           | `validator_address` does not have a `ValidatorConfig::ValidatorConfig` resource published under it.   |
// | `Errors::INVALID_ARGUMENT` | `ValidatorConfig::EINVALID_TRANSACTION_SENDER` | `validator_operator_account` is not the registered operator for the validator at `validator_address`. |
// | `Errors::INVALID_ARGUMENT` | `ValidatorConfig::EINVALID_CONSENSUS_KEY`      | `consensus_pubkey` is not a valid ed25519 public key.                                                 |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::create_validator_operator_account`
// * `Script::add_validator_and_reconfigure`
// * `Script::remove_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::set_validator_config_and_reconfigure`
func EncodeRegisterValidatorConfigScript(validator_account libratypes.AccountAddress, consensus_pubkey []byte, validator_network_addresses []byte, fullnode_network_addresses []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), register_validator_config_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{validator_account}, (*libratypes.TransactionArgument__U8Vector)(&consensus_pubkey), (*libratypes.TransactionArgument__U8Vector)(&validator_network_addresses), (*libratypes.TransactionArgument__U8Vector)(&fullnode_network_addresses)},
	}
}

// # Summary
// This script removes a validator account from the validator set, and triggers a reconfiguration
// of the system to remove the validator from the system. This transaction can only be
// successfully called by the Libra Root account.
//
// # Technical Description
// This script removes the account at `validator_address` from the validator set. This transaction
// emits a `LibraConfig::NewEpochEvent` event. Once the reconfiguration triggered by this event
// has been performed, the account at `validator_address` is no longer considered to be a
// validator in the network. This transaction will fail if the validator at `validator_address`
// is not in the validator set.
//
// # Parameters
// | Name                | Type         | Description                                                                                                                        |
// | ------              | ------       | -------------                                                                                                                      |
// | `lr_account`        | `&signer`    | The signer reference of the sending account of this transaction. Must be the Libra Root signer.                                    |
// | `sliding_nonce`     | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                                         |
// | `validator_name`    | `vector<u8>` | ASCII-encoded human name for the validator. Must match the human name in the `ValidatorConfig::ValidatorConfig` for the validator. |
// | `validator_address` | `address`    | The validator account address to be removed from the validator set.                                                                |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                            | Description                                                                                     |
// | ----------------           | --------------                          | -------------                                                                                   |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `lr_account`.                                  |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not.      |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                                   |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                               |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`          | The sending account is not the Libra Root account or Treasury Compliance account                |
// | 0                          | 0                                       | The provided `validator_name` does not match the already-recorded human name for the validator. |
// | `Errors::INVALID_ARGUMENT` | `LibraSystem::ENOT_AN_ACTIVE_VALIDATOR` | The validator to be removed is not in the validator set.                                        |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ELIBRA_ROOT`            | The sending account is not the Libra Root account.                                              |
// | `Errors::REQUIRES_ROLE`    | `Roles::ELIBRA_ROOT`                    | The sending account is not the Libra Root account.                                              |
// | `Errors::INVALID_STATE`    | `LibraConfig::EINVALID_BLOCK_TIME`      | An invalid time value was encountered in reconfiguration. Unlikely to occur.                    |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::create_validator_operator_account`
// * `Script::register_validator_config`
// * `Script::add_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::set_validator_config_and_reconfigure`
func EncodeRemoveValidatorAndReconfigureScript(sliding_nonce uint64, validator_name []byte, validator_address libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), remove_validator_and_reconfigure_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), (*libratypes.TransactionArgument__U8Vector)(&validator_name), &libratypes.TransactionArgument__Address{validator_address}},
	}
}

// # Summary
// Rotates the transaction sender's authentication key to the supplied new authentication key. May
// be sent by any account.
//
// # Technical Description
// Rotate the `account`'s `LibraAccount::LibraAccount` `authentication_key` field to `new_key`.
// `new_key` must be a valid ed25519 public key, and `account` must not have previously delegated
// its `LibraAccount::KeyRotationCapability`.
//
// # Parameters
// | Name      | Type         | Description                                                 |
// | ------    | ------       | -------------                                               |
// | `account` | `&signer`    | Signer reference of the sending account of the transaction. |
// | `new_key` | `vector<u8>` | New ed25519 public key to be used for `account`.            |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                               | Description                                                                              |
// | ----------------           | --------------                                             | -------------                                                                            |
// | `Errors::INVALID_STATE`    | `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` | `account` has already delegated/extracted its `LibraAccount::KeyRotationCapability`.     |
// | `Errors::INVALID_ARGUMENT` | `LibraAccount::EMALFORMED_AUTHENTICATION_KEY`              | `new_key` was an invalid length.                                                         |
//
// # Related Scripts
// * `Script::rotate_authentication_key_with_nonce`
// * `Script::rotate_authentication_key_with_nonce_admin`
// * `Script::rotate_authentication_key_with_recovery_address`
func EncodeRotateAuthenticationKeyScript(new_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), rotate_authentication_key_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U8Vector)(&new_key)},
	}
}

// # Summary
// Rotates the sender's authentication key to the supplied new authentication key. May be sent by
// any account that has a sliding nonce resource published under it (usually this is Treasury
// Compliance or Libra Root accounts).
//
// # Technical Description
// Rotates the `account`'s `LibraAccount::LibraAccount` `authentication_key` field to `new_key`.
// `new_key` must be a valid ed25519 public key, and `account` must not have previously delegated
// its `LibraAccount::KeyRotationCapability`.
//
// # Parameters
// | Name            | Type         | Description                                                                |
// | ------          | ------       | -------------                                                              |
// | `account`       | `&signer`    | Signer reference of the sending account of the transaction.                |
// | `sliding_nonce` | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction. |
// | `new_key`       | `vector<u8>` | New ed25519 public key to be used for `account`.                           |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                               | Description                                                                                |
// | ----------------           | --------------                                             | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                             | A `SlidingNonce` resource is not published under `account`.                                |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`                             | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`                             | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`                    | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::INVALID_STATE`    | `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` | `account` has already delegated/extracted its `LibraAccount::KeyRotationCapability`.       |
// | `Errors::INVALID_ARGUMENT` | `LibraAccount::EMALFORMED_AUTHENTICATION_KEY`              | `new_key` was an invalid length.                                                           |
//
// # Related Scripts
// * `Script::rotate_authentication_key`
// * `Script::rotate_authentication_key_with_nonce_admin`
// * `Script::rotate_authentication_key_with_recovery_address`
func EncodeRotateAuthenticationKeyWithNonceScript(sliding_nonce uint64, new_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), rotate_authentication_key_with_nonce_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), (*libratypes.TransactionArgument__U8Vector)(&new_key)},
	}
}

// # Summary
// Rotates the specified account's authentication key to the supplied new authentication key. May
// only be sent by the Libra Root account as a write set transaction.
//
// # Technical Description
// Rotate the `account`'s `LibraAccount::LibraAccount` `authentication_key` field to `new_key`.
// `new_key` must be a valid ed25519 public key, and `account` must not have previously delegated
// its `LibraAccount::KeyRotationCapability`.
//
// # Parameters
// | Name            | Type         | Description                                                                                                  |
// | ------          | ------       | -------------                                                                                                |
// | `lr_account`    | `&signer`    | The signer reference of the sending account of the write set transaction. May only be the Libra Root signer. |
// | `account`       | `&signer`    | Signer reference of account specified in the `execute_as` field of the write set transaction.                |
// | `sliding_nonce` | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction for Libra Root.                    |
// | `new_key`       | `vector<u8>` | New ed25519 public key to be used for `account`.                                                             |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                               | Description                                                                                                |
// | ----------------           | --------------                                             | -------------                                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                             | A `SlidingNonce` resource is not published under `lr_account`.                                             |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`                             | The `sliding_nonce` in `lr_account` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`                             | The `sliding_nonce` in `lr_account` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`                    | The `sliding_nonce` in` lr_account` has been previously recorded.                                          |
// | `Errors::INVALID_STATE`    | `LibraAccount::EKEY_ROTATION_CAPABILITY_ALREADY_EXTRACTED` | `account` has already delegated/extracted its `LibraAccount::KeyRotationCapability`.                       |
// | `Errors::INVALID_ARGUMENT` | `LibraAccount::EMALFORMED_AUTHENTICATION_KEY`              | `new_key` was an invalid length.                                                                           |
//
// # Related Scripts
// * `Script::rotate_authentication_key`
// * `Script::rotate_authentication_key_with_nonce`
// * `Script::rotate_authentication_key_with_recovery_address`
func EncodeRotateAuthenticationKeyWithNonceAdminScript(sliding_nonce uint64, new_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), rotate_authentication_key_with_nonce_admin_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), (*libratypes.TransactionArgument__U8Vector)(&new_key)},
	}
}

// # Summary
// Rotates the authentication key of a specified account that is part of a recovery address to a
// new authentication key. Only used for accounts that are part of a recovery address (see
// `Script::add_recovery_rotation_capability` for account restrictions).
//
// # Technical Description
// Rotates the authentication key of the `to_recover` account to `new_key` using the
// `LibraAccount::KeyRotationCapability` stored in the `RecoveryAddress::RecoveryAddress` resource
// published under `recovery_address`. This transaction can be sent either by the `to_recover`
// account, or by the account where the `RecoveryAddress::RecoveryAddress` resource is published
// that contains `to_recover`'s `LibraAccount::KeyRotationCapability`.
//
// # Parameters
// | Name               | Type         | Description                                                                                                                    |
// | ------             | ------       | -------------                                                                                                                  |
// | `account`          | `&signer`    | Signer reference of the sending account of the transaction.                                                                    |
// | `recovery_address` | `address`    | Address where `RecoveryAddress::RecoveryAddress` that holds `to_recover`'s `LibraAccount::KeyRotationCapability` is published. |
// | `to_recover`       | `address`    | The address of the account whose authentication key will be updated.                                                           |
// | `new_key`          | `vector<u8>` | New ed25519 public key to be used for the account at the `to_recover` address.                                                 |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                  | Description                                                                                                                                          |
// | ----------------           | --------------                                | -------------                                                                                                                                        |
// | `Errors::NOT_PUBLISHED`    | `RecoveryAddress::ERECOVERY_ADDRESS`          | `recovery_address` does not have a `RecoveryAddress::RecoveryAddress` resource published under it.                                                   |
// | `Errors::INVALID_ARGUMENT` | `RecoveryAddress::ECANNOT_ROTATE_KEY`         | The address of `account` is not `recovery_address` or `to_recover`.                                                                                  |
// | `Errors::INVALID_ARGUMENT` | `RecoveryAddress::EACCOUNT_NOT_RECOVERABLE`   | `to_recover`'s `LibraAccount::KeyRotationCapability`  is not in the `RecoveryAddress::RecoveryAddress`  resource published under `recovery_address`. |
// | `Errors::INVALID_ARGUMENT` | `LibraAccount::EMALFORMED_AUTHENTICATION_KEY` | `new_key` was an invalid length.                                                                                                                     |
//
// # Related Scripts
// * `Script::rotate_authentication_key`
// * `Script::rotate_authentication_key_with_nonce`
// * `Script::rotate_authentication_key_with_nonce_admin`
func EncodeRotateAuthenticationKeyWithRecoveryAddressScript(recovery_address libratypes.AccountAddress, to_recover libratypes.AccountAddress, new_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), rotate_authentication_key_with_recovery_address_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{recovery_address}, &libratypes.TransactionArgument__Address{to_recover}, (*libratypes.TransactionArgument__U8Vector)(&new_key)},
	}
}

// # Summary
// Updates the url used for off-chain communication, and the public key used to verify dual
// attestation on-chain. Transaction can be sent by any account that has dual attestation
// information published under it. In practice the only such accounts are Designated Dealers and
// Parent VASPs.
//
// # Technical Description
// Updates the `base_url` and `compliance_public_key` fields of the `DualAttestation::Credential`
// resource published under `account`. The `new_key` must be a valid ed25519 public key.
//
// ## Events
// Successful execution of this transaction emits two events:
// * A `DualAttestation::ComplianceKeyRotationEvent` containing the new compliance public key, and
// the blockchain time at which the key was updated emitted on the `DualAttestation::Credential`
// `compliance_key_rotation_events` handle published under `account`; and
// * A `DualAttestation::BaseUrlRotationEvent` containing the new base url to be used for
// off-chain communication, and the blockchain time at which the url was updated emitted on the
// `DualAttestation::Credential` `base_url_rotation_events` handle published under `account`.
//
// # Parameters
// | Name      | Type         | Description                                                               |
// | ------    | ------       | -------------                                                             |
// | `account` | `&signer`    | Signer reference of the sending account of the transaction.               |
// | `new_url` | `vector<u8>` | ASCII-encoded url to be used for off-chain communication with `account`.  |
// | `new_key` | `vector<u8>` | New ed25519 public key to be used for on-chain dual attestation checking. |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                           | Description                                                                |
// | ----------------           | --------------                         | -------------                                                              |
// | `Errors::NOT_PUBLISHED`    | `DualAttestation::ECREDENTIAL`         | A `DualAttestation::Credential` resource is not published under `account`. |
// | `Errors::INVALID_ARGUMENT` | `DualAttestation::EINVALID_PUBLIC_KEY` | `new_key` is not a valid ed25519 public key.                               |
//
// # Related Scripts
// * `Script::create_parent_vasp_account`
// * `Script::create_designated_dealer`
// * `Script::rotate_dual_attestation_info`
func EncodeRotateDualAttestationInfoScript(new_url []byte, new_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), rotate_dual_attestation_info_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U8Vector)(&new_url), (*libratypes.TransactionArgument__U8Vector)(&new_key)},
	}
}

// # Summary
// Rotates the authentication key in a `SharedEd25519PublicKey`. This transaction can be sent by
// any account that has previously published a shared ed25519 public key using
// `Script::publish_shared_ed25519_public_key`.
//
// # Technical Description
// This first rotates the public key stored in `account`'s
// `SharedEd25519PublicKey::SharedEd25519PublicKey` resource to `public_key`, after which it
// rotates the authentication key using the capability stored in `account`'s
// `SharedEd25519PublicKey::SharedEd25519PublicKey` to a new value derived from `public_key`
//
// # Parameters
// | Name         | Type         | Description                                                     |
// | ------       | ------       | -------------                                                   |
// | `account`    | `&signer`    | The signer reference of the sending account of the transaction. |
// | `public_key` | `vector<u8>` | 32-byte Ed25519 public key.                                     |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                    | Description                                                                                   |
// | ----------------           | --------------                                  | -------------                                                                                 |
// | `Errors::NOT_PUBLISHED`    | `SharedEd25519PublicKey::ESHARED_KEY`           | A `SharedEd25519PublicKey::SharedEd25519PublicKey` resource is not published under `account`. |
// | `Errors::INVALID_ARGUMENT` | `SharedEd25519PublicKey::EMALFORMED_PUBLIC_KEY` | `public_key` is an invalid ed25519 public key.                                                |
//
// # Related Scripts
// * `Script::publish_shared_ed25519_public_key`
func EncodeRotateSharedEd25519PublicKeyScript(public_key []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), rotate_shared_ed25519_public_key_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U8Vector)(&public_key)},
	}
}

// # Summary
// Updates a validator's configuration, and triggers a reconfiguration of the system to update the
// validator set with this new validator configuration.  Can only be successfully sent by a
// Validator Operator account that is already registered with a validator.
//
// # Technical Description
// This updates the fields with corresponding names held in the `ValidatorConfig::ValidatorConfig`
// config resource held under `validator_account`. It then emits a `LibraConfig::NewEpochEvent` to
// trigger a reconfiguration of the system.  This reconfiguration will update the validator set
// on-chain with the updated `ValidatorConfig::ValidatorConfig`.
//
// # Parameters
// | Name                          | Type         | Description                                                                                                                  |
// | ------                        | ------       | -------------                                                                                                                |
// | `validator_operator_account`  | `&signer`    | Signer reference of the sending account. Must be the registered validator operator for the validator at `validator_address`. |
// | `validator_account`           | `address`    | The address of the validator's `ValidatorConfig::ValidatorConfig` resource being updated.                                    |
// | `consensus_pubkey`            | `vector<u8>` | New Ed25519 public key to be used in the updated `ValidatorConfig::ValidatorConfig`.                                         |
// | `validator_network_addresses` | `vector<u8>` | New set of `validator_network_addresses` to be used in the updated `ValidatorConfig::ValidatorConfig`.                       |
// | `fullnode_network_addresses`  | `vector<u8>` | New set of `fullnode_network_addresses` to be used in the updated `ValidatorConfig::ValidatorConfig`.                        |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                   | Description                                                                                           |
// | ----------------           | --------------                                 | -------------                                                                                         |
// | `Errors::NOT_PUBLISHED`    | `ValidatorConfig::EVALIDATOR_CONFIG`           | `validator_address` does not have a `ValidatorConfig::ValidatorConfig` resource published under it.   |
// | `Errors::REQUIRES_ROLE     | `Roles::EVALIDATOR_OPERATOR`                   | `validator_operator_account` does not have a Validator Operator role.                                 |
// | `Errors::INVALID_ARGUMENT` | `ValidatorConfig::EINVALID_TRANSACTION_SENDER` | `validator_operator_account` is not the registered operator for the validator at `validator_address`. |
// | `Errors::INVALID_ARGUMENT` | `ValidatorConfig::EINVALID_CONSENSUS_KEY`      | `consensus_pubkey` is not a valid ed25519 public key.                                                 |
// | `Errors::INVALID_STATE`    | `LibraConfig::EINVALID_BLOCK_TIME`             | An invalid time value was encountered in reconfiguration. Unlikely to occur.                          |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::create_validator_operator_account`
// * `Script::add_validator_and_reconfigure`
// * `Script::remove_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::register_validator_config`
func EncodeSetValidatorConfigAndReconfigureScript(validator_account libratypes.AccountAddress, consensus_pubkey []byte, validator_network_addresses []byte, fullnode_network_addresses []byte) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), set_validator_config_and_reconfigure_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{&libratypes.TransactionArgument__Address{validator_account}, (*libratypes.TransactionArgument__U8Vector)(&consensus_pubkey), (*libratypes.TransactionArgument__U8Vector)(&validator_network_addresses), (*libratypes.TransactionArgument__U8Vector)(&fullnode_network_addresses)},
	}
}

// # Summary
// Sets the validator operator for a validator in the validator's configuration resource "locally"
// and does not reconfigure the system. Changes from this transaction will not picked up by the
// system until a reconfiguration of the system is triggered. May only be sent by an account with
// Validator role.
//
// # Technical Description
// Sets the account at `operator_account` address and with the specified `human_name` as an
// operator for the sending validator account. The account at `operator_account` address must have
// a Validator Operator role and have a `ValidatorOperatorConfig::ValidatorOperatorConfig`
// resource published under it. The sending `account` must be a Validator and have a
// `ValidatorConfig::ValidatorConfig` resource published under it. This script does not emit a
// `LibraConfig::NewEpochEvent` and no reconfiguration of the system is initiated by this script.
//
// # Parameters
// | Name               | Type         | Description                                                                                  |
// | ------             | ------       | -------------                                                                                |
// | `account`          | `&signer`    | The signer reference of the sending account of the transaction.                              |
// | `operator_name`    | `vector<u8>` | Validator operator's human name.                                                             |
// | `operator_account` | `address`    | Address of the validator operator account to be added as the `account` validator's operator. |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                          | Description                                                                                                                                                  |
// | ----------------           | --------------                                        | -------------                                                                                                                                                |
// | `Errors::NOT_PUBLISHED`    | `ValidatorOperatorConfig::EVALIDATOR_OPERATOR_CONFIG` | The `ValidatorOperatorConfig::ValidatorOperatorConfig` resource is not published under `operator_account`.                                                   |
// | 0                          | 0                                                     | The `human_name` field of the `ValidatorOperatorConfig::ValidatorOperatorConfig` resource under `operator_account` does not match the provided `human_name`. |
// | `Errors::REQUIRES_ROLE`    | `Roles::EVALIDATOR`                                   | `account` does not have a Validator account role.                                                                                                            |
// | `Errors::INVALID_ARGUMENT` | `ValidatorConfig::ENOT_A_VALIDATOR_OPERATOR`          | The account at `operator_account` does not have a `ValidatorOperatorConfig::ValidatorOperatorConfig` resource.                                               |
// | `Errors::NOT_PUBLISHED`    | `ValidatorConfig::EVALIDATOR_CONFIG`                  | A `ValidatorConfig::ValidatorConfig` is not published under `account`.                                                                                       |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::create_validator_operator_account`
// * `Script::register_validator_config`
// * `Script::remove_validator_and_reconfigure`
// * `Script::add_validator_and_reconfigure`
// * `Script::set_validator_operator_with_nonce_admin`
// * `Script::set_validator_config_and_reconfigure`
func EncodeSetValidatorOperatorScript(operator_name []byte, operator_account libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), set_validator_operator_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U8Vector)(&operator_name), &libratypes.TransactionArgument__Address{operator_account}},
	}
}

// # Summary
// Sets the validator operator for a validator in the validator's configuration resource "locally"
// and does not reconfigure the system. Changes from this transaction will not picked up by the
// system until a reconfiguration of the system is triggered. May only be sent by the Libra Root
// account as a write set transaction.
//
// # Technical Description
// Sets the account at `operator_account` address and with the specified `human_name` as an
// operator for the validator `account`. The account at `operator_account` address must have a
// Validator Operator role and have a `ValidatorOperatorConfig::ValidatorOperatorConfig` resource
// published under it. The account represented by the `account` signer must be a Validator and
// have a `ValidatorConfig::ValidatorConfig` resource published under it. No reconfiguration of
// the system is initiated by this script.
//
// # Parameters
// | Name               | Type         | Description                                                                                                  |
// | ------             | ------       | -------------                                                                                                |
// | `lr_account`       | `&signer`    | The signer reference of the sending account of the write set transaction. May only be the Libra Root signer. |
// | `account`          | `&signer`    | Signer reference of account specified in the `execute_as` field of the write set transaction.                |
// | `sliding_nonce`    | `u64`        | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction for Libra Root.                    |
// | `operator_name`    | `vector<u8>` | Validator operator's human name.                                                                             |
// | `operator_account` | `address`    | Address of the validator operator account to be added as the `account` validator's operator.                 |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                          | Description                                                                                                                                                  |
// | ----------------           | --------------                                        | -------------                                                                                                                                                |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                        | A `SlidingNonce` resource is not published under `lr_account`.                                                                                               |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`                        | The `sliding_nonce` in `lr_account` is too old and it's impossible to determine if it's duplicated or not.                                                   |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`                        | The `sliding_nonce` in `lr_account` is too far in the future.                                                                                                |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`               | The `sliding_nonce` in` lr_account` has been previously recorded.                                                                                            |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                        | The sending account is not the Libra Root account or Treasury Compliance account                                                                             |
// | `Errors::NOT_PUBLISHED`    | `ValidatorOperatorConfig::EVALIDATOR_OPERATOR_CONFIG` | The `ValidatorOperatorConfig::ValidatorOperatorConfig` resource is not published under `operator_account`.                                                   |
// | 0                          | 0                                                     | The `human_name` field of the `ValidatorOperatorConfig::ValidatorOperatorConfig` resource under `operator_account` does not match the provided `human_name`. |
// | `Errors::REQUIRES_ROLE`    | `Roles::EVALIDATOR`                                   | `account` does not have a Validator account role.                                                                                                            |
// | `Errors::INVALID_ARGUMENT` | `ValidatorConfig::ENOT_A_VALIDATOR_OPERATOR`          | The account at `operator_account` does not have a `ValidatorOperatorConfig::ValidatorOperatorConfig` resource.                                               |
// | `Errors::NOT_PUBLISHED`    | `ValidatorConfig::EVALIDATOR_CONFIG`                  | A `ValidatorConfig::ValidatorConfig` is not published under `account`.                                                                                       |
//
// # Related Scripts
// * `Script::create_validator_account`
// * `Script::create_validator_operator_account`
// * `Script::register_validator_config`
// * `Script::remove_validator_and_reconfigure`
// * `Script::add_validator_and_reconfigure`
// * `Script::set_validator_operator`
// * `Script::set_validator_config_and_reconfigure`
func EncodeSetValidatorOperatorWithNonceAdminScript(sliding_nonce uint64, operator_name []byte, operator_account libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), set_validator_operator_with_nonce_admin_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), (*libratypes.TransactionArgument__U8Vector)(&operator_name), &libratypes.TransactionArgument__Address{operator_account}},
	}
}

// # Summary
// Mints a specified number of coins in a currency to a Designated Dealer. The sending account
// must be the Treasury Compliance account, and coins can only be minted to a Designated Dealer
// account.
//
// # Technical Description
// Mints `mint_amount` of coins in the `CoinType` currency to Designated Dealer account at
// `designated_dealer_address`. The `tier_index` parameter specifies which tier should be used to
// check verify the off-chain approval policy, and is based in part on the on-chain tier values
// for the specific Designated Dealer, and the number of `CoinType` coins that have been minted to
// the dealer over the past 24 hours. Every Designated Dealer has 4 tiers for each currency that
// they support. The sending `tc_account` must be the Treasury Compliance account, and the
// receiver an authorized Designated Dealer account.
//
// ## Events
// Successful execution of the transaction will emit two events:
// * A `Libra::MintEvent` with the amount and currency code minted is emitted on the
// `mint_event_handle` in the stored `Libra::CurrencyInfo<CoinType>` resource stored under
// `0xA550C18`; and
// * A `DesignatedDealer::ReceivedMintEvent` with the amount, currency code, and Designated
// Dealer's address is emitted on the `mint_event_handle` in the stored `DesignatedDealer::Dealer`
// resource published under the `designated_dealer_address`.
//
// # Parameters
// | Name                        | Type      | Description                                                                                                |
// | ------                      | ------    | -------------                                                                                              |
// | `CoinType`                  | Type      | The Move type for the `CoinType` being minted. `CoinType` must be an already-registered currency on-chain. |
// | `tc_account`                | `&signer` | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account.  |
// | `sliding_nonce`             | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                 |
// | `designated_dealer_address` | `address` | The address of the Designated Dealer account being minted to.                                              |
// | `mint_amount`               | `u64`     | The number of coins to be minted.                                                                          |
// | `tier_index`                | `u64`     | The mint tier index to use for the Designated Dealer account.                                              |
//
// # Common Abort Conditions
// | Error Category                | Error Reason                                 | Description                                                                                                                  |
// | ----------------              | --------------                               | -------------                                                                                                                |
// | `Errors::NOT_PUBLISHED`       | `SlidingNonce::ESLIDING_NONCE`               | A `SlidingNonce` resource is not published under `tc_account`.                                                               |
// | `Errors::INVALID_ARGUMENT`    | `SlidingNonce::ENONCE_TOO_OLD`               | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not.                                   |
// | `Errors::INVALID_ARGUMENT`    | `SlidingNonce::ENONCE_TOO_NEW`               | The `sliding_nonce` is too far in the future.                                                                                |
// | `Errors::INVALID_ARGUMENT`    | `SlidingNonce::ENONCE_ALREADY_RECORDED`      | The `sliding_nonce` has been previously recorded.                                                                            |
// | `Errors::REQUIRES_ADDRESS`    | `CoreAddresses::ETREASURY_COMPLIANCE`        | `tc_account` is not the Treasury Compliance account.                                                                         |
// | `Errors::REQUIRES_ROLE`       | `Roles::ETREASURY_COMPLIANCE`                | `tc_account` is not the Treasury Compliance account.                                                                         |
// | `Errors::INVALID_ARGUMENT`    | `DesignatedDealer::EINVALID_MINT_AMOUNT`     | `mint_amount` is zero.                                                                                                       |
// | `Errors::NOT_PUBLISHED`       | `DesignatedDealer::EDEALER`                  | `DesignatedDealer::Dealer` or `DesignatedDealer::TierInfo<CoinType>` resource does not exist at `designated_dealer_address`. |
// | `Errors::INVALID_ARGUMENT`    | `DesignatedDealer::EINVALID_TIER_INDEX`      | The `tier_index` is out of bounds.                                                                                           |
// | `Errors::INVALID_ARGUMENT`    | `DesignatedDealer::EINVALID_AMOUNT_FOR_TIER` | `mint_amount` exceeds the maximum allowed amount for `tier_index`.                                                           |
// | `Errors::REQUIRES_CAPABILITY` | `Libra::EMINT_CAPABILITY`                    | `tc_account` does not have a `Libra::MintCapability<CoinType>` resource published under it.                                  |
// | `Errors::INVALID_STATE`       | `Libra::EMINTING_NOT_ALLOWED`                | Minting is not currently allowed for `CoinType` coins.                                                                       |
// | `Errors::LIMIT_EXCEEDED`      | `LibraAccount::EDEPOSIT_EXCEEDS_LIMITS`      | The depositing of the funds would exceed the `account`'s account limits.                                                     |
//
// # Related Scripts
// * `Script::create_designated_dealer`
// * `Script::peer_to_peer_with_metadata`
// * `Script::rotate_dual_attestation_info`
func EncodeTieredMintScript(coin_type libratypes.TypeTag, sliding_nonce uint64, designated_dealer_address libratypes.AccountAddress, mint_amount uint64, tier_index uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), tiered_mint_code...),
		TyArgs: []libratypes.TypeTag{coin_type},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), &libratypes.TransactionArgument__Address{designated_dealer_address}, (*libratypes.TransactionArgument__U64)(&mint_amount), (*libratypes.TransactionArgument__U64)(&tier_index)},
	}
}

// # Summary
// Unfreezes the account at `address`. The sending account of this transaction must be the
// Treasury Compliance account. After the successful execution of this transaction transactions
// may be sent from the previously frozen account, and coins may be sent and received.
//
// # Technical Description
// Sets the `AccountFreezing::FreezingBit` to `false` and emits a
// `AccountFreezing::UnFreezeAccountEvent`. The transaction sender must be the Treasury Compliance
// account. Note that this is a per-account property so unfreezing a Parent VASP will not effect
// the status any of its child accounts and vice versa.
//
// ## Events
// Successful execution of this script will emit a `AccountFreezing::UnFreezeAccountEvent` with
// the `unfrozen_address` set the `to_unfreeze_account`'s address.
//
// # Parameters
// | Name                  | Type      | Description                                                                                               |
// | ------                | ------    | -------------                                                                                             |
// | `tc_account`          | `&signer` | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account. |
// | `sliding_nonce`       | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                |
// | `to_unfreeze_account` | `address` | The account address to be frozen.                                                                         |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                            | Description                                                                                |
// | ----------------           | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `account`.                                |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ETREASURY_COMPLIANCE`   | The sending account is not the Treasury Compliance account.                                |
//
// # Related Scripts
// * `Scripts::freeze_account`
func EncodeUnfreezeAccountScript(sliding_nonce uint64, to_unfreeze_account libratypes.AccountAddress) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), unfreeze_account_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), &libratypes.TransactionArgument__Address{to_unfreeze_account}},
	}
}

// # Summary
// Update the dual attestation limit on-chain. Defined in terms of micro-LBR.  The transaction can
// only be sent by the Treasury Compliance account.  After this transaction all inter-VASP
// payments over this limit must be checked for dual attestation.
//
// # Technical Description
// Updates the `micro_lbr_limit` field of the `DualAttestation::Limit` resource published under
// `0xA550C18`. The amount is set in micro-LBR.
//
// # Parameters
// | Name                  | Type      | Description                                                                                               |
// | ------                | ------    | -------------                                                                                             |
// | `tc_account`          | `&signer` | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account. |
// | `sliding_nonce`       | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction.                                |
// | `new_micro_lbr_limit` | `u64`     | The new dual attestation limit to be used on-chain.                                                       |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                            | Description                                                                                |
// | ----------------           | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `tc_account`.                             |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ETREASURY_COMPLIANCE`   | `tc_account` is not the Treasury Compliance account.                                       |
//
// # Related Scripts
// * `Scripts::update_exchange_rate`
// * `Scripts::update_minting_ability`
func EncodeUpdateDualAttestationLimitScript(sliding_nonce uint64, new_micro_lbr_limit uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), update_dual_attestation_limit_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), (*libratypes.TransactionArgument__U64)(&new_micro_lbr_limit)},
	}
}

// # Summary
// Update the rough on-chain exchange rate between a specified currency and LBR (as a conversion
// to micro-LBR). The transaction can only be sent by the Treasury Compliance account. After this
// transaction the updated exchange rate will be used for normalization of gas prices, and for
// dual attestation checking.
//
// # Technical Description
// Updates the on-chain exchange rate from the given `Currency` to micro-LBR.  The exchange rate
// is given by `new_exchange_rate_numerator/new_exchange_rate_denominator`.
//
// # Parameters
// | Name                            | Type      | Description                                                                                                                        |
// | ------                          | ------    | -------------                                                                                                                      |
// | `Currency`                      | Type      | The Move type for the `Currency` whose exchange rate is being updated. `Currency` must be an already-registered currency on-chain. |
// | `tc_account`                    | `&signer` | The signer reference of the sending account of this transaction. Must be the Treasury Compliance account.                          |
// | `sliding_nonce`                 | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for the transaction.                                                          |
// | `new_exchange_rate_numerator`   | `u64`     | The numerator for the new to micro-LBR exchange rate for `Currency`.                                                               |
// | `new_exchange_rate_denominator` | `u64`     | The denominator for the new to micro-LBR exchange rate for `Currency`.                                                             |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                            | Description                                                                                |
// | ----------------           | --------------                          | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`          | A `SlidingNonce` resource is not published under `tc_account`.                             |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`          | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`          | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED` | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ETREASURY_COMPLIANCE`   | `tc_account` is not the Treasury Compliance account.                                       |
// | `Errors::REQUIRES_ROLE`    | `Roles::ETREASURY_COMPLIANCE`           | `tc_account` is not the Treasury Compliance account.                                       |
// | `Errors::INVALID_ARGUMENT` | `FixedPoint32::EDENOMINATOR`            | `new_exchange_rate_denominator` is zero.                                                   |
// | `Errors::INVALID_ARGUMENT` | `FixedPoint32::ERATIO_OUT_OF_RANGE`     | The quotient is unrepresentable as a `FixedPoint32`.                                       |
// | `Errors::LIMIT_EXCEEDED`   | `FixedPoint32::ERATIO_OUT_OF_RANGE`     | The quotient is unrepresentable as a `FixedPoint32`.                                       |
//
// # Related Scripts
// * `Scripts::update_dual_attestation_limit`
// * `Scripts::update_minting_ability`
func EncodeUpdateExchangeRateScript(currency libratypes.TypeTag, sliding_nonce uint64, new_exchange_rate_numerator uint64, new_exchange_rate_denominator uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), update_exchange_rate_code...),
		TyArgs: []libratypes.TypeTag{currency},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), (*libratypes.TransactionArgument__U64)(&new_exchange_rate_numerator), (*libratypes.TransactionArgument__U64)(&new_exchange_rate_denominator)},
	}
}

// # Summary
// Updates the Libra major version that is stored on-chain and is used by the VM.  This
// transaction can only be sent from the Libra Root account.
//
// # Technical Description
// Updates the `LibraVersion` on-chain config and emits a `LibraConfig::NewEpochEvent` to trigger
// a reconfiguration of the system. The `major` version that is passed in must be strictly greater
// than the current major version held on-chain. The VM reads this information and can use it to
// preserve backwards compatibility with previous major versions of the VM.
//
// # Parameters
// | Name            | Type      | Description                                                                |
// | ------          | ------    | -------------                                                              |
// | `account`       | `&signer` | Signer reference of the sending account. Must be the Libra Root account.   |
// | `sliding_nonce` | `u64`     | The `sliding_nonce` (see: `SlidingNonce`) to be used for this transaction. |
// | `major`         | `u64`     | The `major` version of the VM to be used from this transaction on.         |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                                  | Description                                                                                |
// | ----------------           | --------------                                | -------------                                                                              |
// | `Errors::NOT_PUBLISHED`    | `SlidingNonce::ESLIDING_NONCE`                | A `SlidingNonce` resource is not published under `account`.                                |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_OLD`                | The `sliding_nonce` is too old and it's impossible to determine if it's duplicated or not. |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_TOO_NEW`                | The `sliding_nonce` is too far in the future.                                              |
// | `Errors::INVALID_ARGUMENT` | `SlidingNonce::ENONCE_ALREADY_RECORDED`       | The `sliding_nonce` has been previously recorded.                                          |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ELIBRA_ROOT`                  | `account` is not the Libra Root account.                                                   |
// | `Errors::INVALID_ARGUMENT` | `LibraVersion::EINVALID_MAJOR_VERSION_NUMBER` | `major` is less-than or equal to the current major version stored on-chain.                |
func EncodeUpdateLibraVersionScript(sliding_nonce uint64, major uint64) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), update_libra_version_code...),
		TyArgs: []libratypes.TypeTag{},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__U64)(&sliding_nonce), (*libratypes.TransactionArgument__U64)(&major)},
	}
}

// # Summary
// Script to allow or disallow minting of new coins in a specified currency.  This transaction can
// only be sent by the Treasury Compliance account.  Turning minting off for a currency will have
// no effect on coins already in circulation, and coins may still be removed from the system.
//
// # Technical Description
// This transaction sets the `can_mint` field of the `Libra::CurrencyInfo<Currency>` resource
// published under `0xA550C18` to the value of `allow_minting`. Minting of coins if allowed if
// this field is set to `true` and minting of new coins in `Currency` is disallowed otherwise.
// This transaction needs to be sent by the Treasury Compliance account.
//
// # Parameters
// | Name            | Type      | Description                                                                                                                          |
// | ------          | ------    | -------------                                                                                                                        |
// | `Currency`      | Type      | The Move type for the `Currency` whose minting ability is being updated. `Currency` must be an already-registered currency on-chain. |
// | `account`       | `&signer` | Signer reference of the sending account. Must be the Libra Root account.                                                             |
// | `allow_minting` | `bool`    | Whether to allow minting of new coins in `Currency`.                                                                                 |
//
// # Common Abort Conditions
// | Error Category             | Error Reason                          | Description                                          |
// | ----------------           | --------------                        | -------------                                        |
// | `Errors::REQUIRES_ADDRESS` | `CoreAddresses::ETREASURY_COMPLIANCE` | `tc_account` is not the Treasury Compliance account. |
// | `Errors::NOT_PUBLISHED`    | `Libra::ECURRENCY_INFO`               | `Currency` is not a registered currency on-chain.    |
//
// # Related Scripts
// * `Scripts::update_dual_attestation_limit`
// * `Scripts::update_exchange_rate`
func EncodeUpdateMintingAbilityScript(currency libratypes.TypeTag, allow_minting bool) libratypes.Script {
	return libratypes.Script {
		Code: append([]byte(nil), update_minting_ability_code...),
		TyArgs: []libratypes.TypeTag{currency},
		Args: []libratypes.TransactionArgument{(*libratypes.TransactionArgument__Bool)(&allow_minting)},
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
	if len(script.Args) < 4 { return nil, fmt.Errorf("Was expecting 4 regular arguments") }
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
		call.ValidatorNetworkAddresses = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[3]); err == nil {
		call.FullnodeNetworkAddresses = val
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
	if len(script.Args) < 4 { return nil, fmt.Errorf("Was expecting 4 regular arguments") }
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
		call.ValidatorNetworkAddresses = val
	} else {
		return nil, err
	}

	if val, err := decode_u8vector_argument(script.Args[3]); err == nil {
		call.FullnodeNetworkAddresses = val
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

var peer_to_peer_with_metadata_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 7, 1, 0, 2, 2, 2, 4, 3, 6, 16, 4, 22, 2, 5, 24, 29, 7, 53, 97, 8, 150, 1, 16, 0, 0, 0, 1, 1, 0, 0, 2, 0, 1, 0, 0, 3, 2, 3, 1, 1, 0, 4, 1, 3, 0, 1, 5, 1, 6, 12, 1, 8, 0, 5, 6, 8, 0, 5, 3, 10, 2, 10, 2, 0, 5, 6, 12, 5, 3, 10, 2, 10, 2, 1, 9, 0, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 18, 87, 105, 116, 104, 100, 114, 97, 119, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 27, 101, 120, 116, 114, 97, 99, 116, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 8, 112, 97, 121, 95, 102, 114, 111, 109, 27, 114, 101, 115, 116, 111, 114, 101, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 4, 1, 12, 11, 0, 17, 0, 12, 5, 14, 5, 10, 1, 10, 2, 11, 3, 11, 4, 56, 0, 11, 5, 17, 2, 2};

var preburn_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 7, 1, 0, 2, 2, 2, 4, 3, 6, 16, 4, 22, 2, 5, 24, 21, 7, 45, 96, 8, 141, 1, 16, 0, 0, 0, 1, 1, 0, 0, 2, 0, 1, 0, 0, 3, 2, 3, 1, 1, 0, 4, 1, 3, 0, 1, 5, 1, 6, 12, 1, 8, 0, 3, 6, 12, 6, 8, 0, 3, 0, 2, 6, 12, 3, 1, 9, 0, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 18, 87, 105, 116, 104, 100, 114, 97, 119, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 27, 101, 120, 116, 114, 97, 99, 116, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 7, 112, 114, 101, 98, 117, 114, 110, 27, 114, 101, 115, 116, 111, 114, 101, 95, 119, 105, 116, 104, 100, 114, 97, 119, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 4, 1, 10, 10, 0, 17, 0, 12, 2, 11, 0, 14, 2, 10, 1, 56, 0, 11, 2, 17, 2, 2};

var publish_shared_ed25519_public_key_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 2, 3, 2, 5, 5, 7, 6, 7, 13, 31, 8, 44, 16, 0, 0, 0, 1, 0, 1, 0, 2, 6, 12, 10, 2, 0, 22, 83, 104, 97, 114, 101, 100, 69, 100, 50, 53, 53, 49, 57, 80, 117, 98, 108, 105, 99, 75, 101, 121, 7, 112, 117, 98, 108, 105, 115, 104, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 4, 11, 0, 11, 1, 17, 0, 2};

var register_validator_config_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 2, 3, 2, 5, 5, 7, 11, 7, 18, 27, 8, 45, 16, 0, 0, 0, 1, 0, 1, 0, 5, 6, 12, 5, 10, 2, 10, 2, 10, 2, 0, 15, 86, 97, 108, 105, 100, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 10, 115, 101, 116, 95, 99, 111, 110, 102, 105, 103, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 7, 11, 0, 10, 1, 11, 2, 11, 3, 11, 4, 17, 0, 2};

var remove_validator_and_reconfigure_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 6, 3, 6, 15, 5, 21, 24, 7, 45, 95, 8, 140, 1, 16, 0, 0, 0, 1, 0, 2, 1, 3, 0, 1, 0, 2, 4, 2, 3, 0, 0, 5, 4, 1, 0, 2, 6, 12, 3, 0, 1, 5, 1, 10, 2, 2, 6, 12, 5, 4, 6, 12, 3, 10, 2, 5, 2, 1, 3, 11, 76, 105, 98, 114, 97, 83, 121, 115, 116, 101, 109, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 15, 86, 97, 108, 105, 100, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 14, 103, 101, 116, 95, 104, 117, 109, 97, 110, 95, 110, 97, 109, 101, 16, 114, 101, 109, 111, 118, 101, 95, 118, 97, 108, 105, 100, 97, 116, 111, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 5, 6, 18, 10, 0, 10, 1, 17, 0, 10, 3, 17, 1, 11, 2, 33, 12, 4, 11, 4, 3, 14, 11, 0, 1, 6, 0, 0, 0, 0, 0, 0, 0, 0, 39, 11, 0, 10, 3, 17, 2, 2};

var rotate_authentication_key_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 2, 2, 2, 4, 3, 6, 15, 5, 21, 18, 7, 39, 125, 8, 164, 1, 16, 0, 0, 0, 1, 1, 0, 0, 2, 0, 1, 0, 0, 3, 1, 2, 0, 0, 4, 3, 2, 0, 1, 6, 12, 1, 8, 0, 0, 2, 6, 8, 0, 10, 2, 2, 6, 12, 10, 2, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 21, 75, 101, 121, 82, 111, 116, 97, 116, 105, 111, 110, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 101, 120, 116, 114, 97, 99, 116, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 114, 101, 115, 116, 111, 114, 101, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 25, 114, 111, 116, 97, 116, 101, 95, 97, 117, 116, 104, 101, 110, 116, 105, 99, 97, 116, 105, 111, 110, 95, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 4, 1, 9, 11, 0, 17, 0, 12, 2, 14, 2, 11, 1, 17, 2, 11, 2, 17, 1, 2};

var rotate_authentication_key_with_nonce_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 4, 2, 4, 4, 3, 8, 20, 5, 28, 23, 7, 51, 160, 1, 8, 211, 1, 16, 0, 0, 0, 1, 0, 3, 1, 0, 1, 2, 0, 1, 0, 0, 4, 2, 3, 0, 0, 5, 3, 1, 0, 0, 6, 4, 1, 0, 2, 6, 12, 3, 0, 1, 6, 12, 1, 8, 0, 2, 6, 8, 0, 10, 2, 3, 6, 12, 3, 10, 2, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 21, 75, 101, 121, 82, 111, 116, 97, 116, 105, 111, 110, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 101, 120, 116, 114, 97, 99, 116, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 114, 101, 115, 116, 111, 114, 101, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 25, 114, 111, 116, 97, 116, 101, 95, 97, 117, 116, 104, 101, 110, 116, 105, 99, 97, 116, 105, 111, 110, 95, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 5, 3, 12, 10, 0, 10, 1, 17, 0, 11, 0, 17, 1, 12, 3, 14, 3, 11, 2, 17, 3, 11, 3, 17, 2, 2};

var rotate_authentication_key_with_nonce_admin_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 4, 2, 4, 4, 3, 8, 20, 5, 28, 25, 7, 53, 160, 1, 8, 213, 1, 16, 0, 0, 0, 1, 0, 3, 1, 0, 1, 2, 0, 1, 0, 0, 4, 2, 3, 0, 0, 5, 3, 1, 0, 0, 6, 4, 1, 0, 2, 6, 12, 3, 0, 1, 6, 12, 1, 8, 0, 2, 6, 8, 0, 10, 2, 4, 6, 12, 6, 12, 3, 10, 2, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 21, 75, 101, 121, 82, 111, 116, 97, 116, 105, 111, 110, 67, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 101, 120, 116, 114, 97, 99, 116, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 31, 114, 101, 115, 116, 111, 114, 101, 95, 107, 101, 121, 95, 114, 111, 116, 97, 116, 105, 111, 110, 95, 99, 97, 112, 97, 98, 105, 108, 105, 116, 121, 25, 114, 111, 116, 97, 116, 101, 95, 97, 117, 116, 104, 101, 110, 116, 105, 99, 97, 116, 105, 111, 110, 95, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 5, 3, 12, 11, 0, 10, 2, 17, 0, 11, 1, 17, 1, 12, 4, 14, 4, 11, 3, 17, 3, 11, 4, 17, 2, 2};

var rotate_authentication_key_with_recovery_address_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 2, 3, 2, 5, 5, 7, 8, 7, 15, 42, 8, 57, 16, 0, 0, 0, 1, 0, 1, 0, 4, 6, 12, 5, 5, 10, 2, 0, 15, 82, 101, 99, 111, 118, 101, 114, 121, 65, 100, 100, 114, 101, 115, 115, 25, 114, 111, 116, 97, 116, 101, 95, 97, 117, 116, 104, 101, 110, 116, 105, 99, 97, 116, 105, 111, 110, 95, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 6, 11, 0, 10, 1, 10, 2, 11, 3, 17, 0, 2};

var rotate_dual_attestation_info_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 2, 3, 2, 10, 5, 12, 13, 7, 25, 61, 8, 86, 16, 0, 0, 0, 1, 0, 1, 0, 0, 2, 0, 1, 0, 2, 6, 12, 10, 2, 0, 3, 6, 12, 10, 2, 10, 2, 15, 68, 117, 97, 108, 65, 116, 116, 101, 115, 116, 97, 116, 105, 111, 110, 15, 114, 111, 116, 97, 116, 101, 95, 98, 97, 115, 101, 95, 117, 114, 108, 28, 114, 111, 116, 97, 116, 101, 95, 99, 111, 109, 112, 108, 105, 97, 110, 99, 101, 95, 112, 117, 98, 108, 105, 99, 95, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 2, 1, 7, 10, 0, 11, 1, 17, 0, 11, 0, 11, 2, 17, 1, 2};

var rotate_shared_ed25519_public_key_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 2, 3, 2, 5, 5, 7, 6, 7, 13, 34, 8, 47, 16, 0, 0, 0, 1, 0, 1, 0, 2, 6, 12, 10, 2, 0, 22, 83, 104, 97, 114, 101, 100, 69, 100, 50, 53, 53, 49, 57, 80, 117, 98, 108, 105, 99, 75, 101, 121, 10, 114, 111, 116, 97, 116, 101, 95, 107, 101, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 4, 11, 0, 11, 1, 17, 0, 2};

var set_validator_config_and_reconfigure_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 4, 3, 4, 10, 5, 14, 15, 7, 29, 69, 8, 98, 16, 0, 0, 0, 1, 1, 2, 0, 1, 0, 0, 3, 2, 1, 0, 5, 6, 12, 5, 10, 2, 10, 2, 10, 2, 0, 2, 6, 12, 5, 11, 76, 105, 98, 114, 97, 83, 121, 115, 116, 101, 109, 15, 86, 97, 108, 105, 100, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 10, 115, 101, 116, 95, 99, 111, 110, 102, 105, 103, 29, 117, 112, 100, 97, 116, 101, 95, 99, 111, 110, 102, 105, 103, 95, 97, 110, 100, 95, 114, 101, 99, 111, 110, 102, 105, 103, 117, 114, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 10, 10, 0, 10, 1, 11, 2, 11, 3, 11, 4, 17, 0, 11, 0, 10, 1, 17, 1, 2};

var set_validator_operator_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 4, 3, 4, 10, 5, 14, 19, 7, 33, 68, 8, 101, 16, 0, 0, 0, 1, 1, 2, 0, 1, 0, 0, 3, 2, 3, 0, 1, 5, 1, 10, 2, 2, 6, 12, 5, 0, 3, 6, 12, 10, 2, 5, 2, 1, 3, 15, 86, 97, 108, 105, 100, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 23, 86, 97, 108, 105, 100, 97, 116, 111, 114, 79, 112, 101, 114, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 14, 103, 101, 116, 95, 104, 117, 109, 97, 110, 95, 110, 97, 109, 101, 12, 115, 101, 116, 95, 111, 112, 101, 114, 97, 116, 111, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 4, 5, 15, 10, 2, 17, 0, 11, 1, 33, 12, 3, 11, 3, 3, 11, 11, 0, 1, 6, 0, 0, 0, 0, 0, 0, 0, 0, 39, 11, 0, 10, 2, 17, 1, 2};

var set_validator_operator_with_nonce_admin_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 6, 3, 6, 15, 5, 21, 26, 7, 47, 103, 8, 150, 1, 16, 0, 0, 0, 1, 0, 2, 0, 3, 0, 1, 0, 2, 4, 2, 3, 0, 1, 5, 4, 1, 0, 2, 6, 12, 3, 0, 1, 5, 1, 10, 2, 2, 6, 12, 5, 5, 6, 12, 6, 12, 3, 10, 2, 5, 2, 1, 3, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 15, 86, 97, 108, 105, 100, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 23, 86, 97, 108, 105, 100, 97, 116, 111, 114, 79, 112, 101, 114, 97, 116, 111, 114, 67, 111, 110, 102, 105, 103, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 14, 103, 101, 116, 95, 104, 117, 109, 97, 110, 95, 110, 97, 109, 101, 12, 115, 101, 116, 95, 111, 112, 101, 114, 97, 116, 111, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 5, 6, 18, 11, 0, 10, 2, 17, 0, 10, 4, 17, 1, 11, 3, 33, 12, 5, 11, 5, 3, 14, 11, 1, 1, 6, 0, 0, 0, 0, 0, 0, 0, 0, 39, 11, 1, 10, 4, 17, 2, 2};

var tiered_mint_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 6, 1, 0, 4, 3, 4, 11, 4, 15, 2, 5, 17, 21, 7, 38, 60, 8, 98, 16, 0, 0, 0, 1, 1, 2, 0, 1, 0, 0, 3, 2, 1, 1, 1, 1, 4, 2, 6, 12, 3, 0, 4, 6, 12, 5, 3, 3, 5, 6, 12, 3, 5, 3, 3, 1, 9, 0, 12, 76, 105, 98, 114, 97, 65, 99, 99, 111, 117, 110, 116, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 11, 116, 105, 101, 114, 101, 100, 95, 109, 105, 110, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 3, 1, 9, 10, 0, 10, 1, 17, 0, 11, 0, 10, 2, 10, 3, 10, 4, 56, 0, 2};

var unfreeze_account_code = []byte {161, 28, 235, 11, 1, 0, 0, 0, 5, 1, 0, 4, 3, 4, 10, 5, 14, 14, 7, 28, 68, 8, 96, 16, 0, 0, 0, 1, 0, 2, 0, 1, 0, 1, 3, 2, 1, 0, 2, 6, 12, 5, 0, 2, 6, 12, 3, 3, 6, 12, 3, 5, 15, 65, 99, 99, 111, 117, 110, 116, 70, 114, 101, 101, 122, 105, 110, 103, 12, 83, 108, 105, 100, 105, 110, 103, 78, 111, 110, 99, 101, 16, 117, 110, 102, 114, 101, 101, 122, 101, 95, 97, 99, 99, 111, 117, 110, 116, 21, 114, 101, 99, 111, 114, 100, 95, 110, 111, 110, 99, 101, 95, 111, 114, 95, 97, 98, 111, 114, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 3, 1, 7, 10, 0, 10, 1, 17, 1, 11, 0, 10, 2, 17, 0, 2};

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
	string(update_dual_attestation_limit_code): decode_update_dual_attestation_limit_script,
	string(update_exchange_rate_code): decode_update_exchange_rate_script,
	string(update_libra_version_code): decode_update_libra_version_script,
	string(update_minting_ability_code): decode_update_minting_ability_script,
}

func decode_bool_argument(arg libratypes.TransactionArgument) (value bool, err error) {
	if arg, ok := arg.(*libratypes.TransactionArgument__Bool); ok {
		value = bool(*arg)
	} else {
		err = fmt.Errorf("Was expecting a Bool argument")
	}
	return
}


func decode_u64_argument(arg libratypes.TransactionArgument) (value uint64, err error) {
	if arg, ok := arg.(*libratypes.TransactionArgument__U64); ok {
		value = uint64(*arg)
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
		value = []byte(*arg)
	} else {
		err = fmt.Errorf("Was expecting a U8Vector argument")
	}
	return
}

