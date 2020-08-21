package libratypes


import (
	"fmt"
	"github.com/facebookincubator/serde-reflection/serde-generate/runtime/golang/serde"
)


type AccessPath struct {
	Address AccountAddress
	Path []byte
}

func (obj *AccessPath) Serialize(serializer serde.Serializer) error {
	if err := obj.Address.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeBytes(obj.Path); err != nil { return err }
	return nil
}

func DeserializeAccessPath(deserializer serde.Deserializer) (AccessPath, error) {
	var obj AccessPath
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.Address = val } else { return obj, err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Path = val } else { return obj, err }
	return obj, nil
}


type AccountAddress struct {
	Value [16]uint8
}

func (obj *AccountAddress) Serialize(serializer serde.Serializer) error {
	if err := serialize_array16_u8_array(obj.Value, serializer); err != nil { return err }
	return nil
}

func DeserializeAccountAddress(deserializer serde.Deserializer) (AccountAddress, error) {
	var obj AccountAddress
	if val, err := deserialize_array16_u8_array(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type BlockMetadata struct {
	Id HashValue
	Round uint64
	TimestampUsecs uint64
	PreviousBlockVotes []AccountAddress
	Proposer AccountAddress
}

func (obj *BlockMetadata) Serialize(serializer serde.Serializer) error {
	if err := obj.Id.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(obj.Round); err != nil { return err }
	if err := serializer.SerializeU64(obj.TimestampUsecs); err != nil { return err }
	if err := serialize_vector_AccountAddress(obj.PreviousBlockVotes, serializer); err != nil { return err }
	if err := obj.Proposer.Serialize(serializer); err != nil { return err }
	return nil
}

func DeserializeBlockMetadata(deserializer serde.Deserializer) (BlockMetadata, error) {
	var obj BlockMetadata
	if val, err := DeserializeHashValue(deserializer); err == nil { obj.Id = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.Round = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.TimestampUsecs = val } else { return obj, err }
	if val, err := deserialize_vector_AccountAddress(deserializer); err == nil { obj.PreviousBlockVotes = val } else { return obj, err }
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.Proposer = val } else { return obj, err }
	return obj, nil
}


type ChainId struct {
	Value uint8
}

func (obj *ChainId) Serialize(serializer serde.Serializer) error {
	if err := serializer.SerializeU8(obj.Value); err != nil { return err }
	return nil
}

func DeserializeChainId(deserializer serde.Deserializer) (ChainId, error) {
	var obj ChainId
	if val, err := deserializer.DeserializeU8(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type ChangeSet struct {
	WriteSet WriteSet
	Events []ContractEvent
}

func (obj *ChangeSet) Serialize(serializer serde.Serializer) error {
	if err := obj.WriteSet.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_ContractEvent(obj.Events, serializer); err != nil { return err }
	return nil
}

func DeserializeChangeSet(deserializer serde.Deserializer) (ChangeSet, error) {
	var obj ChangeSet
	if val, err := DeserializeWriteSet(deserializer); err == nil { obj.WriteSet = val } else { return obj, err }
	if val, err := deserialize_vector_ContractEvent(deserializer); err == nil { obj.Events = val } else { return obj, err }
	return obj, nil
}


type ContractEvent interface {
	isContractEvent()
	Serialize(serializer serde.Serializer) error
}

func DeserializeContractEvent(deserializer serde.Deserializer) (ContractEvent, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_ContractEvent__V0(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for ContractEvent: %d", index)
	}
}

type ContractEvent__V0 struct {
	Value ContractEventV0
}

func (*ContractEvent__V0) isContractEvent() {}

func (obj *ContractEvent__V0) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_ContractEvent__V0(deserializer serde.Deserializer) (ContractEvent__V0, error) {
	var obj ContractEvent__V0
	if val, err := DeserializeContractEventV0(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type ContractEventV0 struct {
	Key EventKey
	SequenceNumber uint64
	TypeTag TypeTag
	EventData []byte
}

func (obj *ContractEventV0) Serialize(serializer serde.Serializer) error {
	if err := obj.Key.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(obj.SequenceNumber); err != nil { return err }
	if err := obj.TypeTag.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeBytes(obj.EventData); err != nil { return err }
	return nil
}

func DeserializeContractEventV0(deserializer serde.Deserializer) (ContractEventV0, error) {
	var obj ContractEventV0
	if val, err := DeserializeEventKey(deserializer); err == nil { obj.Key = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.SequenceNumber = val } else { return obj, err }
	if val, err := DeserializeTypeTag(deserializer); err == nil { obj.TypeTag = val } else { return obj, err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.EventData = val } else { return obj, err }
	return obj, nil
}


type Ed25519PublicKey struct {
	Value []byte
}

func (obj *Ed25519PublicKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.SerializeBytes(obj.Value); err != nil { return err }
	return nil
}

func DeserializeEd25519PublicKey(deserializer serde.Deserializer) (Ed25519PublicKey, error) {
	var obj Ed25519PublicKey
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type Ed25519Signature struct {
	Value []byte
}

func (obj *Ed25519Signature) Serialize(serializer serde.Serializer) error {
	if err := serializer.SerializeBytes(obj.Value); err != nil { return err }
	return nil
}

func DeserializeEd25519Signature(deserializer serde.Deserializer) (Ed25519Signature, error) {
	var obj Ed25519Signature
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type EventKey struct {
	Value []byte
}

func (obj *EventKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.SerializeBytes(obj.Value); err != nil { return err }
	return nil
}

func DeserializeEventKey(deserializer serde.Deserializer) (EventKey, error) {
	var obj EventKey
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type GeneralMetadata interface {
	isGeneralMetadata()
	Serialize(serializer serde.Serializer) error
}

func DeserializeGeneralMetadata(deserializer serde.Deserializer) (GeneralMetadata, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_GeneralMetadata__GeneralMetadataVersion0(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for GeneralMetadata: %d", index)
	}
}

type GeneralMetadata__GeneralMetadataVersion0 struct {
	Value GeneralMetadataV0
}

func (*GeneralMetadata__GeneralMetadataVersion0) isGeneralMetadata() {}

func (obj *GeneralMetadata__GeneralMetadataVersion0) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_GeneralMetadata__GeneralMetadataVersion0(deserializer serde.Deserializer) (GeneralMetadata__GeneralMetadataVersion0, error) {
	var obj GeneralMetadata__GeneralMetadataVersion0
	if val, err := DeserializeGeneralMetadataV0(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type GeneralMetadataV0 struct {
	ToSubaddress *[]byte
	FromSubaddress *[]byte
	ReferencedEvent *uint64
}

func (obj *GeneralMetadataV0) Serialize(serializer serde.Serializer) error {
	if err := serialize_option_bytes(obj.ToSubaddress, serializer); err != nil { return err }
	if err := serialize_option_bytes(obj.FromSubaddress, serializer); err != nil { return err }
	if err := serialize_option_u64(obj.ReferencedEvent, serializer); err != nil { return err }
	return nil
}

func DeserializeGeneralMetadataV0(deserializer serde.Deserializer) (GeneralMetadataV0, error) {
	var obj GeneralMetadataV0
	if val, err := deserialize_option_bytes(deserializer); err == nil { obj.ToSubaddress = val } else { return obj, err }
	if val, err := deserialize_option_bytes(deserializer); err == nil { obj.FromSubaddress = val } else { return obj, err }
	if val, err := deserialize_option_u64(deserializer); err == nil { obj.ReferencedEvent = val } else { return obj, err }
	return obj, nil
}


type HashValue struct {
	Value []byte
}

func (obj *HashValue) Serialize(serializer serde.Serializer) error {
	if err := serializer.SerializeBytes(obj.Value); err != nil { return err }
	return nil
}

func DeserializeHashValue(deserializer serde.Deserializer) (HashValue, error) {
	var obj HashValue
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type Identifier struct {
	Value string
}

func (obj *Identifier) Serialize(serializer serde.Serializer) error {
	if err := serializer.SerializeStr(obj.Value); err != nil { return err }
	return nil
}

func DeserializeIdentifier(deserializer serde.Deserializer) (Identifier, error) {
	var obj Identifier
	if val, err := deserializer.DeserializeStr(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type Metadata interface {
	isMetadata()
	Serialize(serializer serde.Serializer) error
}

func DeserializeMetadata(deserializer serde.Deserializer) (Metadata, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_Metadata__Undefined(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_Metadata__GeneralMetadata(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_Metadata__TravelRuleMetadata(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_Metadata__UnstructuredBytesMetadata(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for Metadata: %d", index)
	}
}

type Metadata__Undefined struct {
}

func (*Metadata__Undefined) isMetadata() {}

func (obj *Metadata__Undefined) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(0)
	return nil
}

func load_Metadata__Undefined(deserializer serde.Deserializer) (Metadata__Undefined, error) {
	var obj Metadata__Undefined
	return obj, nil
}


type Metadata__GeneralMetadata struct {
	Value GeneralMetadata
}

func (*Metadata__GeneralMetadata) isMetadata() {}

func (obj *Metadata__GeneralMetadata) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_Metadata__GeneralMetadata(deserializer serde.Deserializer) (Metadata__GeneralMetadata, error) {
	var obj Metadata__GeneralMetadata
	if val, err := DeserializeGeneralMetadata(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type Metadata__TravelRuleMetadata struct {
	Value TravelRuleMetadata
}

func (*Metadata__TravelRuleMetadata) isMetadata() {}

func (obj *Metadata__TravelRuleMetadata) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(2)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_Metadata__TravelRuleMetadata(deserializer serde.Deserializer) (Metadata__TravelRuleMetadata, error) {
	var obj Metadata__TravelRuleMetadata
	if val, err := DeserializeTravelRuleMetadata(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type Metadata__UnstructuredBytesMetadata struct {
	Value UnstructuredBytesMetadata
}

func (*Metadata__UnstructuredBytesMetadata) isMetadata() {}

func (obj *Metadata__UnstructuredBytesMetadata) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(3)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_Metadata__UnstructuredBytesMetadata(deserializer serde.Deserializer) (Metadata__UnstructuredBytesMetadata, error) {
	var obj Metadata__UnstructuredBytesMetadata
	if val, err := DeserializeUnstructuredBytesMetadata(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type Module struct {
	Code []byte
}

func (obj *Module) Serialize(serializer serde.Serializer) error {
	if err := serializer.SerializeBytes(obj.Code); err != nil { return err }
	return nil
}

func DeserializeModule(deserializer serde.Deserializer) (Module, error) {
	var obj Module
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Code = val } else { return obj, err }
	return obj, nil
}


type MultiEd25519PublicKey struct {
	Value []byte
}

func (obj *MultiEd25519PublicKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.SerializeBytes(obj.Value); err != nil { return err }
	return nil
}

func DeserializeMultiEd25519PublicKey(deserializer serde.Deserializer) (MultiEd25519PublicKey, error) {
	var obj MultiEd25519PublicKey
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type MultiEd25519Signature struct {
	Value []byte
}

func (obj *MultiEd25519Signature) Serialize(serializer serde.Serializer) error {
	if err := serializer.SerializeBytes(obj.Value); err != nil { return err }
	return nil
}

func DeserializeMultiEd25519Signature(deserializer serde.Deserializer) (MultiEd25519Signature, error) {
	var obj MultiEd25519Signature
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type RawTransaction struct {
	Sender AccountAddress
	SequenceNumber uint64
	Payload TransactionPayload
	MaxGasAmount uint64
	GasUnitPrice uint64
	GasCurrencyCode string
	ExpirationTimestampSecs uint64
	ChainId ChainId
}

func (obj *RawTransaction) Serialize(serializer serde.Serializer) error {
	if err := obj.Sender.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(obj.SequenceNumber); err != nil { return err }
	if err := obj.Payload.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(obj.MaxGasAmount); err != nil { return err }
	if err := serializer.SerializeU64(obj.GasUnitPrice); err != nil { return err }
	if err := serializer.SerializeStr(obj.GasCurrencyCode); err != nil { return err }
	if err := serializer.SerializeU64(obj.ExpirationTimestampSecs); err != nil { return err }
	if err := obj.ChainId.Serialize(serializer); err != nil { return err }
	return nil
}

func DeserializeRawTransaction(deserializer serde.Deserializer) (RawTransaction, error) {
	var obj RawTransaction
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.Sender = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.SequenceNumber = val } else { return obj, err }
	if val, err := DeserializeTransactionPayload(deserializer); err == nil { obj.Payload = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.MaxGasAmount = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.GasUnitPrice = val } else { return obj, err }
	if val, err := deserializer.DeserializeStr(); err == nil { obj.GasCurrencyCode = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.ExpirationTimestampSecs = val } else { return obj, err }
	if val, err := DeserializeChainId(deserializer); err == nil { obj.ChainId = val } else { return obj, err }
	return obj, nil
}


type Script struct {
	Code []byte
	TyArgs []TypeTag
	Args []TransactionArgument
}

func (obj *Script) Serialize(serializer serde.Serializer) error {
	if err := serializer.SerializeBytes(obj.Code); err != nil { return err }
	if err := serialize_vector_TypeTag(obj.TyArgs, serializer); err != nil { return err }
	if err := serialize_vector_TransactionArgument(obj.Args, serializer); err != nil { return err }
	return nil
}

func DeserializeScript(deserializer serde.Deserializer) (Script, error) {
	var obj Script
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Code = val } else { return obj, err }
	if val, err := deserialize_vector_TypeTag(deserializer); err == nil { obj.TyArgs = val } else { return obj, err }
	if val, err := deserialize_vector_TransactionArgument(deserializer); err == nil { obj.Args = val } else { return obj, err }
	return obj, nil
}


type SignedTransaction struct {
	RawTxn RawTransaction
	Authenticator TransactionAuthenticator
}

func (obj *SignedTransaction) Serialize(serializer serde.Serializer) error {
	if err := obj.RawTxn.Serialize(serializer); err != nil { return err }
	if err := obj.Authenticator.Serialize(serializer); err != nil { return err }
	return nil
}

func DeserializeSignedTransaction(deserializer serde.Deserializer) (SignedTransaction, error) {
	var obj SignedTransaction
	if val, err := DeserializeRawTransaction(deserializer); err == nil { obj.RawTxn = val } else { return obj, err }
	if val, err := DeserializeTransactionAuthenticator(deserializer); err == nil { obj.Authenticator = val } else { return obj, err }
	return obj, nil
}


type StructTag struct {
	Address AccountAddress
	Module Identifier
	Name Identifier
	TypeParams []TypeTag
}

func (obj *StructTag) Serialize(serializer serde.Serializer) error {
	if err := obj.Address.Serialize(serializer); err != nil { return err }
	if err := obj.Module.Serialize(serializer); err != nil { return err }
	if err := obj.Name.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_TypeTag(obj.TypeParams, serializer); err != nil { return err }
	return nil
}

func DeserializeStructTag(deserializer serde.Deserializer) (StructTag, error) {
	var obj StructTag
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.Address = val } else { return obj, err }
	if val, err := DeserializeIdentifier(deserializer); err == nil { obj.Module = val } else { return obj, err }
	if val, err := DeserializeIdentifier(deserializer); err == nil { obj.Name = val } else { return obj, err }
	if val, err := deserialize_vector_TypeTag(deserializer); err == nil { obj.TypeParams = val } else { return obj, err }
	return obj, nil
}


type Transaction interface {
	isTransaction()
	Serialize(serializer serde.Serializer) error
}

func DeserializeTransaction(deserializer serde.Deserializer) (Transaction, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_Transaction__UserTransaction(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_Transaction__GenesisTransaction(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_Transaction__BlockMetadata(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for Transaction: %d", index)
	}
}

type Transaction__UserTransaction struct {
	Value SignedTransaction
}

func (*Transaction__UserTransaction) isTransaction() {}

func (obj *Transaction__UserTransaction) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_Transaction__UserTransaction(deserializer serde.Deserializer) (Transaction__UserTransaction, error) {
	var obj Transaction__UserTransaction
	if val, err := DeserializeSignedTransaction(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type Transaction__GenesisTransaction struct {
	Value WriteSetPayload
}

func (*Transaction__GenesisTransaction) isTransaction() {}

func (obj *Transaction__GenesisTransaction) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_Transaction__GenesisTransaction(deserializer serde.Deserializer) (Transaction__GenesisTransaction, error) {
	var obj Transaction__GenesisTransaction
	if val, err := DeserializeWriteSetPayload(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type Transaction__BlockMetadata struct {
	Value BlockMetadata
}

func (*Transaction__BlockMetadata) isTransaction() {}

func (obj *Transaction__BlockMetadata) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(2)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_Transaction__BlockMetadata(deserializer serde.Deserializer) (Transaction__BlockMetadata, error) {
	var obj Transaction__BlockMetadata
	if val, err := DeserializeBlockMetadata(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type TransactionArgument interface {
	isTransactionArgument()
	Serialize(serializer serde.Serializer) error
}

func DeserializeTransactionArgument(deserializer serde.Deserializer) (TransactionArgument, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_TransactionArgument__U8(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TransactionArgument__U64(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_TransactionArgument__U128(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_TransactionArgument__Address(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 4:
		if val, err := load_TransactionArgument__U8Vector(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 5:
		if val, err := load_TransactionArgument__Bool(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TransactionArgument: %d", index)
	}
}

type TransactionArgument__U8 struct {
	Value uint8
}

func (*TransactionArgument__U8) isTransactionArgument() {}

func (obj *TransactionArgument__U8) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(0)
	if err := serializer.SerializeU8(obj.Value); err != nil { return err }
	return nil
}

func load_TransactionArgument__U8(deserializer serde.Deserializer) (TransactionArgument__U8, error) {
	var obj TransactionArgument__U8
	if val, err := deserializer.DeserializeU8(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type TransactionArgument__U64 struct {
	Value uint64
}

func (*TransactionArgument__U64) isTransactionArgument() {}

func (obj *TransactionArgument__U64) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(1)
	if err := serializer.SerializeU64(obj.Value); err != nil { return err }
	return nil
}

func load_TransactionArgument__U64(deserializer serde.Deserializer) (TransactionArgument__U64, error) {
	var obj TransactionArgument__U64
	if val, err := deserializer.DeserializeU64(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type TransactionArgument__U128 struct {
	Value serde.Uint128
}

func (*TransactionArgument__U128) isTransactionArgument() {}

func (obj *TransactionArgument__U128) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(2)
	if err := serializer.SerializeU128(obj.Value); err != nil { return err }
	return nil
}

func load_TransactionArgument__U128(deserializer serde.Deserializer) (TransactionArgument__U128, error) {
	var obj TransactionArgument__U128
	if val, err := deserializer.DeserializeU128(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type TransactionArgument__Address struct {
	Value AccountAddress
}

func (*TransactionArgument__Address) isTransactionArgument() {}

func (obj *TransactionArgument__Address) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(3)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_TransactionArgument__Address(deserializer serde.Deserializer) (TransactionArgument__Address, error) {
	var obj TransactionArgument__Address
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type TransactionArgument__U8Vector struct {
	Value []byte
}

func (*TransactionArgument__U8Vector) isTransactionArgument() {}

func (obj *TransactionArgument__U8Vector) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(4)
	if err := serializer.SerializeBytes(obj.Value); err != nil { return err }
	return nil
}

func load_TransactionArgument__U8Vector(deserializer serde.Deserializer) (TransactionArgument__U8Vector, error) {
	var obj TransactionArgument__U8Vector
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type TransactionArgument__Bool struct {
	Value bool
}

func (*TransactionArgument__Bool) isTransactionArgument() {}

func (obj *TransactionArgument__Bool) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(5)
	if err := serializer.SerializeBool(obj.Value); err != nil { return err }
	return nil
}

func load_TransactionArgument__Bool(deserializer serde.Deserializer) (TransactionArgument__Bool, error) {
	var obj TransactionArgument__Bool
	if val, err := deserializer.DeserializeBool(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type TransactionAuthenticator interface {
	isTransactionAuthenticator()
	Serialize(serializer serde.Serializer) error
}

func DeserializeTransactionAuthenticator(deserializer serde.Deserializer) (TransactionAuthenticator, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_TransactionAuthenticator__Ed25519(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TransactionAuthenticator__MultiEd25519(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TransactionAuthenticator: %d", index)
	}
}

type TransactionAuthenticator__Ed25519 struct {
	PublicKey Ed25519PublicKey
	Signature Ed25519Signature
}

func (*TransactionAuthenticator__Ed25519) isTransactionAuthenticator() {}

func (obj *TransactionAuthenticator__Ed25519) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(0)
	if err := obj.PublicKey.Serialize(serializer); err != nil { return err }
	if err := obj.Signature.Serialize(serializer); err != nil { return err }
	return nil
}

func load_TransactionAuthenticator__Ed25519(deserializer serde.Deserializer) (TransactionAuthenticator__Ed25519, error) {
	var obj TransactionAuthenticator__Ed25519
	if val, err := DeserializeEd25519PublicKey(deserializer); err == nil { obj.PublicKey = val } else { return obj, err }
	if val, err := DeserializeEd25519Signature(deserializer); err == nil { obj.Signature = val } else { return obj, err }
	return obj, nil
}


type TransactionAuthenticator__MultiEd25519 struct {
	PublicKey MultiEd25519PublicKey
	Signature MultiEd25519Signature
}

func (*TransactionAuthenticator__MultiEd25519) isTransactionAuthenticator() {}

func (obj *TransactionAuthenticator__MultiEd25519) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(1)
	if err := obj.PublicKey.Serialize(serializer); err != nil { return err }
	if err := obj.Signature.Serialize(serializer); err != nil { return err }
	return nil
}

func load_TransactionAuthenticator__MultiEd25519(deserializer serde.Deserializer) (TransactionAuthenticator__MultiEd25519, error) {
	var obj TransactionAuthenticator__MultiEd25519
	if val, err := DeserializeMultiEd25519PublicKey(deserializer); err == nil { obj.PublicKey = val } else { return obj, err }
	if val, err := DeserializeMultiEd25519Signature(deserializer); err == nil { obj.Signature = val } else { return obj, err }
	return obj, nil
}


type TransactionPayload interface {
	isTransactionPayload()
	Serialize(serializer serde.Serializer) error
}

func DeserializeTransactionPayload(deserializer serde.Deserializer) (TransactionPayload, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_TransactionPayload__WriteSet(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TransactionPayload__Script(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_TransactionPayload__Module(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TransactionPayload: %d", index)
	}
}

type TransactionPayload__WriteSet struct {
	Value WriteSetPayload
}

func (*TransactionPayload__WriteSet) isTransactionPayload() {}

func (obj *TransactionPayload__WriteSet) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_TransactionPayload__WriteSet(deserializer serde.Deserializer) (TransactionPayload__WriteSet, error) {
	var obj TransactionPayload__WriteSet
	if val, err := DeserializeWriteSetPayload(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type TransactionPayload__Script struct {
	Value Script
}

func (*TransactionPayload__Script) isTransactionPayload() {}

func (obj *TransactionPayload__Script) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_TransactionPayload__Script(deserializer serde.Deserializer) (TransactionPayload__Script, error) {
	var obj TransactionPayload__Script
	if val, err := DeserializeScript(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type TransactionPayload__Module struct {
	Value Module
}

func (*TransactionPayload__Module) isTransactionPayload() {}

func (obj *TransactionPayload__Module) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(2)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_TransactionPayload__Module(deserializer serde.Deserializer) (TransactionPayload__Module, error) {
	var obj TransactionPayload__Module
	if val, err := DeserializeModule(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type TravelRuleMetadata interface {
	isTravelRuleMetadata()
	Serialize(serializer serde.Serializer) error
}

func DeserializeTravelRuleMetadata(deserializer serde.Deserializer) (TravelRuleMetadata, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_TravelRuleMetadata__TravelRuleMetadataVersion0(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TravelRuleMetadata: %d", index)
	}
}

type TravelRuleMetadata__TravelRuleMetadataVersion0 struct {
	Value TravelRuleMetadataV0
}

func (*TravelRuleMetadata__TravelRuleMetadataVersion0) isTravelRuleMetadata() {}

func (obj *TravelRuleMetadata__TravelRuleMetadataVersion0) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_TravelRuleMetadata__TravelRuleMetadataVersion0(deserializer serde.Deserializer) (TravelRuleMetadata__TravelRuleMetadataVersion0, error) {
	var obj TravelRuleMetadata__TravelRuleMetadataVersion0
	if val, err := DeserializeTravelRuleMetadataV0(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type TravelRuleMetadataV0 struct {
	OffChainReferenceId *string
}

func (obj *TravelRuleMetadataV0) Serialize(serializer serde.Serializer) error {
	if err := serialize_option_str(obj.OffChainReferenceId, serializer); err != nil { return err }
	return nil
}

func DeserializeTravelRuleMetadataV0(deserializer serde.Deserializer) (TravelRuleMetadataV0, error) {
	var obj TravelRuleMetadataV0
	if val, err := deserialize_option_str(deserializer); err == nil { obj.OffChainReferenceId = val } else { return obj, err }
	return obj, nil
}


type TypeTag interface {
	isTypeTag()
	Serialize(serializer serde.Serializer) error
}

func DeserializeTypeTag(deserializer serde.Deserializer) (TypeTag, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_TypeTag__Bool(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_TypeTag__U8(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_TypeTag__U64(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_TypeTag__U128(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 4:
		if val, err := load_TypeTag__Address(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 5:
		if val, err := load_TypeTag__Signer(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 6:
		if val, err := load_TypeTag__Vector(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 7:
		if val, err := load_TypeTag__Struct(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for TypeTag: %d", index)
	}
}

type TypeTag__Bool struct {
}

func (*TypeTag__Bool) isTypeTag() {}

func (obj *TypeTag__Bool) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(0)
	return nil
}

func load_TypeTag__Bool(deserializer serde.Deserializer) (TypeTag__Bool, error) {
	var obj TypeTag__Bool
	return obj, nil
}


type TypeTag__U8 struct {
}

func (*TypeTag__U8) isTypeTag() {}

func (obj *TypeTag__U8) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(1)
	return nil
}

func load_TypeTag__U8(deserializer serde.Deserializer) (TypeTag__U8, error) {
	var obj TypeTag__U8
	return obj, nil
}


type TypeTag__U64 struct {
}

func (*TypeTag__U64) isTypeTag() {}

func (obj *TypeTag__U64) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(2)
	return nil
}

func load_TypeTag__U64(deserializer serde.Deserializer) (TypeTag__U64, error) {
	var obj TypeTag__U64
	return obj, nil
}


type TypeTag__U128 struct {
}

func (*TypeTag__U128) isTypeTag() {}

func (obj *TypeTag__U128) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(3)
	return nil
}

func load_TypeTag__U128(deserializer serde.Deserializer) (TypeTag__U128, error) {
	var obj TypeTag__U128
	return obj, nil
}


type TypeTag__Address struct {
}

func (*TypeTag__Address) isTypeTag() {}

func (obj *TypeTag__Address) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(4)
	return nil
}

func load_TypeTag__Address(deserializer serde.Deserializer) (TypeTag__Address, error) {
	var obj TypeTag__Address
	return obj, nil
}


type TypeTag__Signer struct {
}

func (*TypeTag__Signer) isTypeTag() {}

func (obj *TypeTag__Signer) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(5)
	return nil
}

func load_TypeTag__Signer(deserializer serde.Deserializer) (TypeTag__Signer, error) {
	var obj TypeTag__Signer
	return obj, nil
}


type TypeTag__Vector struct {
	Value TypeTag
}

func (*TypeTag__Vector) isTypeTag() {}

func (obj *TypeTag__Vector) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(6)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_TypeTag__Vector(deserializer serde.Deserializer) (TypeTag__Vector, error) {
	var obj TypeTag__Vector
	if val, err := DeserializeTypeTag(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type TypeTag__Struct struct {
	Value StructTag
}

func (*TypeTag__Struct) isTypeTag() {}

func (obj *TypeTag__Struct) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(7)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_TypeTag__Struct(deserializer serde.Deserializer) (TypeTag__Struct, error) {
	var obj TypeTag__Struct
	if val, err := DeserializeStructTag(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type UnstructuredBytesMetadata struct {
	Metadata *[]byte
}

func (obj *UnstructuredBytesMetadata) Serialize(serializer serde.Serializer) error {
	if err := serialize_option_bytes(obj.Metadata, serializer); err != nil { return err }
	return nil
}

func DeserializeUnstructuredBytesMetadata(deserializer serde.Deserializer) (UnstructuredBytesMetadata, error) {
	var obj UnstructuredBytesMetadata
	if val, err := deserialize_option_bytes(deserializer); err == nil { obj.Metadata = val } else { return obj, err }
	return obj, nil
}


type WriteOp interface {
	isWriteOp()
	Serialize(serializer serde.Serializer) error
}

func DeserializeWriteOp(deserializer serde.Deserializer) (WriteOp, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_WriteOp__Deletion(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_WriteOp__Value(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for WriteOp: %d", index)
	}
}

type WriteOp__Deletion struct {
}

func (*WriteOp__Deletion) isWriteOp() {}

func (obj *WriteOp__Deletion) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(0)
	return nil
}

func load_WriteOp__Deletion(deserializer serde.Deserializer) (WriteOp__Deletion, error) {
	var obj WriteOp__Deletion
	return obj, nil
}


type WriteOp__Value struct {
	Value []byte
}

func (*WriteOp__Value) isWriteOp() {}

func (obj *WriteOp__Value) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(1)
	if err := serializer.SerializeBytes(obj.Value); err != nil { return err }
	return nil
}

func load_WriteOp__Value(deserializer serde.Deserializer) (WriteOp__Value, error) {
	var obj WriteOp__Value
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type WriteSet struct {
	Value WriteSetMut
}

func (obj *WriteSet) Serialize(serializer serde.Serializer) error {
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func DeserializeWriteSet(deserializer serde.Deserializer) (WriteSet, error) {
	var obj WriteSet
	if val, err := DeserializeWriteSetMut(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type WriteSetMut struct {
	WriteSet []struct {Field0 AccessPath; Field1 WriteOp}
}

func (obj *WriteSetMut) Serialize(serializer serde.Serializer) error {
	if err := serialize_vector_tuple2_AccessPath_WriteOp(obj.WriteSet, serializer); err != nil { return err }
	return nil
}

func DeserializeWriteSetMut(deserializer serde.Deserializer) (WriteSetMut, error) {
	var obj WriteSetMut
	if val, err := deserialize_vector_tuple2_AccessPath_WriteOp(deserializer); err == nil { obj.WriteSet = val } else { return obj, err }
	return obj, nil
}


type WriteSetPayload interface {
	isWriteSetPayload()
	Serialize(serializer serde.Serializer) error
}

func DeserializeWriteSetPayload(deserializer serde.Deserializer) (WriteSetPayload, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_WriteSetPayload__Direct(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_WriteSetPayload__Script(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for WriteSetPayload: %d", index)
	}
}

type WriteSetPayload__Direct struct {
	Value ChangeSet
}

func (*WriteSetPayload__Direct) isWriteSetPayload() {}

func (obj *WriteSetPayload__Direct) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	return nil
}

func load_WriteSetPayload__Direct(deserializer serde.Deserializer) (WriteSetPayload__Direct, error) {
	var obj WriteSetPayload__Direct
	if val, err := DeserializeChangeSet(deserializer); err == nil { obj.Value = val } else { return obj, err }
	return obj, nil
}


type WriteSetPayload__Script struct {
	ExecuteAs AccountAddress
	Script Script
}

func (*WriteSetPayload__Script) isWriteSetPayload() {}

func (obj *WriteSetPayload__Script) Serialize(serializer serde.Serializer) error {
	serializer.SerializeVariantIndex(1)
	if err := obj.ExecuteAs.Serialize(serializer); err != nil { return err }
	if err := obj.Script.Serialize(serializer); err != nil { return err }
	return nil
}

func load_WriteSetPayload__Script(deserializer serde.Deserializer) (WriteSetPayload__Script, error) {
	var obj WriteSetPayload__Script
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.ExecuteAs = val } else { return obj, err }
	if val, err := DeserializeScript(deserializer); err == nil { obj.Script = val } else { return obj, err }
	return obj, nil
}

func serialize_array16_u8_array(value [16]uint8, serializer serde.Serializer) error {
	for _, item := range(value) {
		if err := serializer.SerializeU8(item); err != nil { return err }
	}
	return nil
}

func deserialize_array16_u8_array(deserializer serde.Deserializer) ([16]uint8, error) {
	var obj [16]uint8
	for i := range(obj) {
		if val, err := deserializer.DeserializeU8(); err == nil { obj[i] = val } else { return obj, err }
	}
	return obj, nil
}

func serialize_option_bytes(value *[]byte, serializer serde.Serializer) error {
	if (value != nil) {
		if err := serializer.SerializeOptionTag(true); err != nil { return err }
		if err := serializer.SerializeBytes((*value)); err != nil { return err }
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil { return err }
	}
	return nil
}

func deserialize_option_bytes(deserializer serde.Deserializer) (*[]byte, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil { return nil, err }
	var value *[]byte
	if (tag) {
		if val, err := deserializer.DeserializeBytes(); err == nil { *value = val } else { return nil, err }
	}
	return value, nil
}

func serialize_option_str(value *string, serializer serde.Serializer) error {
	if (value != nil) {
		if err := serializer.SerializeOptionTag(true); err != nil { return err }
		if err := serializer.SerializeStr((*value)); err != nil { return err }
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil { return err }
	}
	return nil
}

func deserialize_option_str(deserializer serde.Deserializer) (*string, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil { return nil, err }
	var value *string
	if (tag) {
		if val, err := deserializer.DeserializeStr(); err == nil { *value = val } else { return nil, err }
	}
	return value, nil
}

func serialize_option_u64(value *uint64, serializer serde.Serializer) error {
	if (value != nil) {
		if err := serializer.SerializeOptionTag(true); err != nil { return err }
		if err := serializer.SerializeU64((*value)); err != nil { return err }
	} else {
		if err := serializer.SerializeOptionTag(false); err != nil { return err }
	}
	return nil
}

func deserialize_option_u64(deserializer serde.Deserializer) (*uint64, error) {
	tag, err := deserializer.DeserializeOptionTag()
	if err != nil { return nil, err }
	var value *uint64
	if (tag) {
		if val, err := deserializer.DeserializeU64(); err == nil { *value = val } else { return nil, err }
	}
	return value, nil
}

func serialize_tuple2_AccessPath_WriteOp(value struct {Field0 AccessPath; Field1 WriteOp}, serializer serde.Serializer) error {
	if err := value.Field0.Serialize(serializer); err != nil { return err }
	if err := value.Field1.Serialize(serializer); err != nil { return err }
	return nil
}

func deserialize_tuple2_AccessPath_WriteOp(deserializer serde.Deserializer) (struct {Field0 AccessPath; Field1 WriteOp}, error) {
	var obj struct {Field0 AccessPath; Field1 WriteOp}
	if val, err := DeserializeAccessPath(deserializer); err == nil { obj.Field0 = val } else { return obj, err }
	if val, err := DeserializeWriteOp(deserializer); err == nil { obj.Field1 = val } else { return obj, err }
	return obj, nil
}

func serialize_vector_AccountAddress(value []AccountAddress, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_AccountAddress(deserializer serde.Deserializer) ([]AccountAddress, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]AccountAddress, length)
	for i := range(obj) {
		if val, err := DeserializeAccountAddress(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_ContractEvent(value []ContractEvent, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_ContractEvent(deserializer serde.Deserializer) ([]ContractEvent, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]ContractEvent, length)
	for i := range(obj) {
		if val, err := DeserializeContractEvent(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_TransactionArgument(value []TransactionArgument, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_TransactionArgument(deserializer serde.Deserializer) ([]TransactionArgument, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]TransactionArgument, length)
	for i := range(obj) {
		if val, err := DeserializeTransactionArgument(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_TypeTag(value []TypeTag, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := item.Serialize(serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_TypeTag(deserializer serde.Deserializer) ([]TypeTag, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]TypeTag, length)
	for i := range(obj) {
		if val, err := DeserializeTypeTag(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

func serialize_vector_tuple2_AccessPath_WriteOp(value []struct {Field0 AccessPath; Field1 WriteOp}, serializer serde.Serializer) error {
	if err := serializer.SerializeLen(uint64(len(value))); err != nil { return err }
	for _, item := range(value) {
		if err := serialize_tuple2_AccessPath_WriteOp(item, serializer); err != nil { return err }
	}
	return nil
}

func deserialize_vector_tuple2_AccessPath_WriteOp(deserializer serde.Deserializer) ([]struct {Field0 AccessPath; Field1 WriteOp}, error) {
	length, err := deserializer.DeserializeLen()
	if err != nil { return nil, err }
	obj := make([]struct {Field0 AccessPath; Field1 WriteOp}, length)
	for i := range(obj) {
		if val, err := deserialize_tuple2_AccessPath_WriteOp(deserializer); err == nil { obj[i] = val } else { return nil, err }
	}
	return obj, nil
}

