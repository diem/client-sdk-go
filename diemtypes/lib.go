package diemtypes


import (
	"fmt"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/bcs"
)


type AccessPath struct {
	Address AccountAddress
	Path []byte
}

func (obj *AccessPath) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Address.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeBytes(obj.Path); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *AccessPath) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeAccessPath(deserializer serde.Deserializer) (AccessPath, error) {
	var obj AccessPath
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.Address = val } else { return obj, err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Path = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeAccessPath(input []byte) (AccessPath, error) {
	if input == nil {
		var obj AccessPath
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeAccessPath(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type AccountAddress [16]uint8

func (obj *AccountAddress) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_array16_u8_array((([16]uint8)(*obj)), serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *AccountAddress) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeAccountAddress(deserializer serde.Deserializer) (AccountAddress, error) {
	var obj [16]uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (AccountAddress)(obj), err }
	if val, err := deserialize_array16_u8_array(deserializer); err == nil { obj = val } else { return ((AccountAddress)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (AccountAddress)(obj), nil
}

func BcsDeserializeAccountAddress(input []byte) (AccountAddress, error) {
	if input == nil {
		var obj AccountAddress
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeAccountAddress(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type BlockMetadata struct {
	Id HashValue
	Round uint64
	TimestampUsecs uint64
	PreviousBlockVotes []AccountAddress
	Proposer AccountAddress
}

func (obj *BlockMetadata) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Id.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(obj.Round); err != nil { return err }
	if err := serializer.SerializeU64(obj.TimestampUsecs); err != nil { return err }
	if err := serialize_vector_AccountAddress(obj.PreviousBlockVotes, serializer); err != nil { return err }
	if err := obj.Proposer.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *BlockMetadata) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeBlockMetadata(deserializer serde.Deserializer) (BlockMetadata, error) {
	var obj BlockMetadata
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeHashValue(deserializer); err == nil { obj.Id = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.Round = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.TimestampUsecs = val } else { return obj, err }
	if val, err := deserialize_vector_AccountAddress(deserializer); err == nil { obj.PreviousBlockVotes = val } else { return obj, err }
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.Proposer = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeBlockMetadata(input []byte) (BlockMetadata, error) {
	if input == nil {
		var obj BlockMetadata
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeBlockMetadata(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ChainId uint8

func (obj *ChainId) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeU8(((uint8)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ChainId) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeChainId(deserializer serde.Deserializer) (ChainId, error) {
	var obj uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (ChainId)(obj), err }
	if val, err := deserializer.DeserializeU8(); err == nil { obj = val } else { return ((ChainId)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (ChainId)(obj), nil
}

func BcsDeserializeChainId(input []byte) (ChainId, error) {
	if input == nil {
		var obj ChainId
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeChainId(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ChangeSet struct {
	WriteSet WriteSet
	Events []ContractEvent
}

func (obj *ChangeSet) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.WriteSet.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_ContractEvent(obj.Events, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ChangeSet) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeChangeSet(deserializer serde.Deserializer) (ChangeSet, error) {
	var obj ChangeSet
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeWriteSet(deserializer); err == nil { obj.WriteSet = val } else { return obj, err }
	if val, err := deserialize_vector_ContractEvent(deserializer); err == nil { obj.Events = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeChangeSet(input []byte) (ChangeSet, error) {
	if input == nil {
		var obj ChangeSet
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeChangeSet(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ContractEvent interface {
	isContractEvent()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
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

func BcsDeserializeContractEvent(input []byte) (ContractEvent, error) {
	if input == nil {
		var obj ContractEvent
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeContractEvent(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type ContractEvent__V0 struct {
	Value ContractEventV0
}

func (*ContractEvent__V0) isContractEvent() {}

func (obj *ContractEvent__V0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ContractEvent__V0) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_ContractEvent__V0(deserializer serde.Deserializer) (ContractEvent__V0, error) {
	var obj ContractEvent__V0
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeContractEventV0(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type ContractEventV0 struct {
	Key EventKey
	SequenceNumber uint64
	TypeTag TypeTag
	EventData []byte
}

func (obj *ContractEventV0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Key.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(obj.SequenceNumber); err != nil { return err }
	if err := obj.TypeTag.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeBytes(obj.EventData); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *ContractEventV0) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeContractEventV0(deserializer serde.Deserializer) (ContractEventV0, error) {
	var obj ContractEventV0
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeEventKey(deserializer); err == nil { obj.Key = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.SequenceNumber = val } else { return obj, err }
	if val, err := DeserializeTypeTag(deserializer); err == nil { obj.TypeTag = val } else { return obj, err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.EventData = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeContractEventV0(input []byte) (ContractEventV0, error) {
	if input == nil {
		var obj ContractEventV0
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeContractEventV0(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Ed25519PublicKey []byte

func (obj *Ed25519PublicKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Ed25519PublicKey) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeEd25519PublicKey(deserializer serde.Deserializer) (Ed25519PublicKey, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (Ed25519PublicKey)(obj), err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj = val } else { return ((Ed25519PublicKey)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (Ed25519PublicKey)(obj), nil
}

func BcsDeserializeEd25519PublicKey(input []byte) (Ed25519PublicKey, error) {
	if input == nil {
		var obj Ed25519PublicKey
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeEd25519PublicKey(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Ed25519Signature []byte

func (obj *Ed25519Signature) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Ed25519Signature) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeEd25519Signature(deserializer serde.Deserializer) (Ed25519Signature, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (Ed25519Signature)(obj), err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj = val } else { return ((Ed25519Signature)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (Ed25519Signature)(obj), nil
}

func BcsDeserializeEd25519Signature(input []byte) (Ed25519Signature, error) {
	if input == nil {
		var obj Ed25519Signature
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeEd25519Signature(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type EventKey []byte

func (obj *EventKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *EventKey) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeEventKey(deserializer serde.Deserializer) (EventKey, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (EventKey)(obj), err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj = val } else { return ((EventKey)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (EventKey)(obj), nil
}

func BcsDeserializeEventKey(input []byte) (EventKey, error) {
	if input == nil {
		var obj EventKey
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeEventKey(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type GeneralMetadata interface {
	isGeneralMetadata()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
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

func BcsDeserializeGeneralMetadata(input []byte) (GeneralMetadata, error) {
	if input == nil {
		var obj GeneralMetadata
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeGeneralMetadata(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type GeneralMetadata__GeneralMetadataVersion0 struct {
	Value GeneralMetadataV0
}

func (*GeneralMetadata__GeneralMetadataVersion0) isGeneralMetadata() {}

func (obj *GeneralMetadata__GeneralMetadataVersion0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *GeneralMetadata__GeneralMetadataVersion0) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_GeneralMetadata__GeneralMetadataVersion0(deserializer serde.Deserializer) (GeneralMetadata__GeneralMetadataVersion0, error) {
	var obj GeneralMetadata__GeneralMetadataVersion0
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeGeneralMetadataV0(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type GeneralMetadataV0 struct {
	ToSubaddress *[]byte
	FromSubaddress *[]byte
	ReferencedEvent *uint64
}

func (obj *GeneralMetadataV0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_option_bytes(obj.ToSubaddress, serializer); err != nil { return err }
	if err := serialize_option_bytes(obj.FromSubaddress, serializer); err != nil { return err }
	if err := serialize_option_u64(obj.ReferencedEvent, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *GeneralMetadataV0) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeGeneralMetadataV0(deserializer serde.Deserializer) (GeneralMetadataV0, error) {
	var obj GeneralMetadataV0
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_option_bytes(deserializer); err == nil { obj.ToSubaddress = val } else { return obj, err }
	if val, err := deserialize_option_bytes(deserializer); err == nil { obj.FromSubaddress = val } else { return obj, err }
	if val, err := deserialize_option_u64(deserializer); err == nil { obj.ReferencedEvent = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeGeneralMetadataV0(input []byte) (GeneralMetadataV0, error) {
	if input == nil {
		var obj GeneralMetadataV0
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeGeneralMetadataV0(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type HashValue []byte

func (obj *HashValue) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *HashValue) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeHashValue(deserializer serde.Deserializer) (HashValue, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (HashValue)(obj), err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj = val } else { return ((HashValue)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (HashValue)(obj), nil
}

func BcsDeserializeHashValue(input []byte) (HashValue, error) {
	if input == nil {
		var obj HashValue
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeHashValue(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Identifier string

func (obj *Identifier) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeStr(((string)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Identifier) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeIdentifier(deserializer serde.Deserializer) (Identifier, error) {
	var obj string
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (Identifier)(obj), err }
	if val, err := deserializer.DeserializeStr(); err == nil { obj = val } else { return ((Identifier)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (Identifier)(obj), nil
}

func BcsDeserializeIdentifier(input []byte) (Identifier, error) {
	if input == nil {
		var obj Identifier
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeIdentifier(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Metadata interface {
	isMetadata()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
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

	case 4:
		if val, err := load_Metadata__RefundMetadata(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for Metadata: %d", index)
	}
}

func BcsDeserializeMetadata(input []byte) (Metadata, error) {
	if input == nil {
		var obj Metadata
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMetadata(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Metadata__Undefined struct {
}

func (*Metadata__Undefined) isMetadata() {}

func (obj *Metadata__Undefined) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Metadata__Undefined) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Metadata__Undefined(deserializer serde.Deserializer) (Metadata__Undefined, error) {
	var obj Metadata__Undefined
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Metadata__GeneralMetadata struct {
	Value GeneralMetadata
}

func (*Metadata__GeneralMetadata) isMetadata() {}

func (obj *Metadata__GeneralMetadata) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Metadata__GeneralMetadata) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Metadata__GeneralMetadata(deserializer serde.Deserializer) (Metadata__GeneralMetadata, error) {
	var obj Metadata__GeneralMetadata
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeGeneralMetadata(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Metadata__TravelRuleMetadata struct {
	Value TravelRuleMetadata
}

func (*Metadata__TravelRuleMetadata) isMetadata() {}

func (obj *Metadata__TravelRuleMetadata) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Metadata__TravelRuleMetadata) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Metadata__TravelRuleMetadata(deserializer serde.Deserializer) (Metadata__TravelRuleMetadata, error) {
	var obj Metadata__TravelRuleMetadata
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeTravelRuleMetadata(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Metadata__UnstructuredBytesMetadata struct {
	Value UnstructuredBytesMetadata
}

func (*Metadata__UnstructuredBytesMetadata) isMetadata() {}

func (obj *Metadata__UnstructuredBytesMetadata) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Metadata__UnstructuredBytesMetadata) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Metadata__UnstructuredBytesMetadata(deserializer serde.Deserializer) (Metadata__UnstructuredBytesMetadata, error) {
	var obj Metadata__UnstructuredBytesMetadata
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeUnstructuredBytesMetadata(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Metadata__RefundMetadata struct {
	Value RefundMetadata
}

func (*Metadata__RefundMetadata) isMetadata() {}

func (obj *Metadata__RefundMetadata) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(4)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Metadata__RefundMetadata) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Metadata__RefundMetadata(deserializer serde.Deserializer) (Metadata__RefundMetadata, error) {
	var obj Metadata__RefundMetadata
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeRefundMetadata(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Module struct {
	Code []byte
}

func (obj *Module) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeBytes(obj.Code); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Module) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeModule(deserializer serde.Deserializer) (Module, error) {
	var obj Module
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Code = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeModule(input []byte) (Module, error) {
	if input == nil {
		var obj Module
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeModule(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MultiEd25519PublicKey []byte

func (obj *MultiEd25519PublicKey) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MultiEd25519PublicKey) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeMultiEd25519PublicKey(deserializer serde.Deserializer) (MultiEd25519PublicKey, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (MultiEd25519PublicKey)(obj), err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj = val } else { return ((MultiEd25519PublicKey)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (MultiEd25519PublicKey)(obj), nil
}

func BcsDeserializeMultiEd25519PublicKey(input []byte) (MultiEd25519PublicKey, error) {
	if input == nil {
		var obj MultiEd25519PublicKey
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMultiEd25519PublicKey(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type MultiEd25519Signature []byte

func (obj *MultiEd25519Signature) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *MultiEd25519Signature) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeMultiEd25519Signature(deserializer serde.Deserializer) (MultiEd25519Signature, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (MultiEd25519Signature)(obj), err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj = val } else { return ((MultiEd25519Signature)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (MultiEd25519Signature)(obj), nil
}

func BcsDeserializeMultiEd25519Signature(input []byte) (MultiEd25519Signature, error) {
	if input == nil {
		var obj MultiEd25519Signature
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeMultiEd25519Signature(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
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
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Sender.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(obj.SequenceNumber); err != nil { return err }
	if err := obj.Payload.Serialize(serializer); err != nil { return err }
	if err := serializer.SerializeU64(obj.MaxGasAmount); err != nil { return err }
	if err := serializer.SerializeU64(obj.GasUnitPrice); err != nil { return err }
	if err := serializer.SerializeStr(obj.GasCurrencyCode); err != nil { return err }
	if err := serializer.SerializeU64(obj.ExpirationTimestampSecs); err != nil { return err }
	if err := obj.ChainId.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *RawTransaction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeRawTransaction(deserializer serde.Deserializer) (RawTransaction, error) {
	var obj RawTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.Sender = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.SequenceNumber = val } else { return obj, err }
	if val, err := DeserializeTransactionPayload(deserializer); err == nil { obj.Payload = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.MaxGasAmount = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.GasUnitPrice = val } else { return obj, err }
	if val, err := deserializer.DeserializeStr(); err == nil { obj.GasCurrencyCode = val } else { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.ExpirationTimestampSecs = val } else { return obj, err }
	if val, err := DeserializeChainId(deserializer); err == nil { obj.ChainId = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeRawTransaction(input []byte) (RawTransaction, error) {
	if input == nil {
		var obj RawTransaction
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeRawTransaction(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type RefundMetadata interface {
	isRefundMetadata()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeRefundMetadata(deserializer serde.Deserializer) (RefundMetadata, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_RefundMetadata__RefundMetadataV0(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for RefundMetadata: %d", index)
	}
}

func BcsDeserializeRefundMetadata(input []byte) (RefundMetadata, error) {
	if input == nil {
		var obj RefundMetadata
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeRefundMetadata(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type RefundMetadata__RefundMetadataV0 struct {
	Value RefundMetadataV0
}

func (*RefundMetadata__RefundMetadataV0) isRefundMetadata() {}

func (obj *RefundMetadata__RefundMetadataV0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *RefundMetadata__RefundMetadataV0) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_RefundMetadata__RefundMetadataV0(deserializer serde.Deserializer) (RefundMetadata__RefundMetadataV0, error) {
	var obj RefundMetadata__RefundMetadataV0
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeRefundMetadataV0(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type RefundMetadataV0 struct {
	TransactionVersion uint64
	Reason RefundReason
}

func (obj *RefundMetadataV0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeU64(obj.TransactionVersion); err != nil { return err }
	if err := obj.Reason.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *RefundMetadataV0) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeRefundMetadataV0(deserializer serde.Deserializer) (RefundMetadataV0, error) {
	var obj RefundMetadataV0
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj.TransactionVersion = val } else { return obj, err }
	if val, err := DeserializeRefundReason(deserializer); err == nil { obj.Reason = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeRefundMetadataV0(input []byte) (RefundMetadataV0, error) {
	if input == nil {
		var obj RefundMetadataV0
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeRefundMetadataV0(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type RefundReason interface {
	isRefundReason()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
}

func DeserializeRefundReason(deserializer serde.Deserializer) (RefundReason, error) {
	index, err := deserializer.DeserializeVariantIndex()
	if err != nil { return nil, err }

	switch index {
	case 0:
		if val, err := load_RefundReason__OtherReason(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 1:
		if val, err := load_RefundReason__InvalidSubaddress(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 2:
		if val, err := load_RefundReason__UserInitiatedPartialRefund(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	case 3:
		if val, err := load_RefundReason__UserInitiatedFullRefund(deserializer); err == nil {
			return &val, nil
		} else {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("Unknown variant index for RefundReason: %d", index)
	}
}

func BcsDeserializeRefundReason(input []byte) (RefundReason, error) {
	if input == nil {
		var obj RefundReason
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeRefundReason(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type RefundReason__OtherReason struct {
}

func (*RefundReason__OtherReason) isRefundReason() {}

func (obj *RefundReason__OtherReason) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *RefundReason__OtherReason) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_RefundReason__OtherReason(deserializer serde.Deserializer) (RefundReason__OtherReason, error) {
	var obj RefundReason__OtherReason
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type RefundReason__InvalidSubaddress struct {
}

func (*RefundReason__InvalidSubaddress) isRefundReason() {}

func (obj *RefundReason__InvalidSubaddress) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *RefundReason__InvalidSubaddress) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_RefundReason__InvalidSubaddress(deserializer serde.Deserializer) (RefundReason__InvalidSubaddress, error) {
	var obj RefundReason__InvalidSubaddress
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type RefundReason__UserInitiatedPartialRefund struct {
}

func (*RefundReason__UserInitiatedPartialRefund) isRefundReason() {}

func (obj *RefundReason__UserInitiatedPartialRefund) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *RefundReason__UserInitiatedPartialRefund) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_RefundReason__UserInitiatedPartialRefund(deserializer serde.Deserializer) (RefundReason__UserInitiatedPartialRefund, error) {
	var obj RefundReason__UserInitiatedPartialRefund
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type RefundReason__UserInitiatedFullRefund struct {
}

func (*RefundReason__UserInitiatedFullRefund) isRefundReason() {}

func (obj *RefundReason__UserInitiatedFullRefund) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *RefundReason__UserInitiatedFullRefund) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_RefundReason__UserInitiatedFullRefund(deserializer serde.Deserializer) (RefundReason__UserInitiatedFullRefund, error) {
	var obj RefundReason__UserInitiatedFullRefund
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Script struct {
	Code []byte
	TyArgs []TypeTag
	Args []TransactionArgument
}

func (obj *Script) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serializer.SerializeBytes(obj.Code); err != nil { return err }
	if err := serialize_vector_TypeTag(obj.TyArgs, serializer); err != nil { return err }
	if err := serialize_vector_TransactionArgument(obj.Args, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Script) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeScript(deserializer serde.Deserializer) (Script, error) {
	var obj Script
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj.Code = val } else { return obj, err }
	if val, err := deserialize_vector_TypeTag(deserializer); err == nil { obj.TyArgs = val } else { return obj, err }
	if val, err := deserialize_vector_TransactionArgument(deserializer); err == nil { obj.Args = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeScript(input []byte) (Script, error) {
	if input == nil {
		var obj Script
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeScript(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type SignedTransaction struct {
	RawTxn RawTransaction
	Authenticator TransactionAuthenticator
}

func (obj *SignedTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.RawTxn.Serialize(serializer); err != nil { return err }
	if err := obj.Authenticator.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *SignedTransaction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeSignedTransaction(deserializer serde.Deserializer) (SignedTransaction, error) {
	var obj SignedTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeRawTransaction(deserializer); err == nil { obj.RawTxn = val } else { return obj, err }
	if val, err := DeserializeTransactionAuthenticator(deserializer); err == nil { obj.Authenticator = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeSignedTransaction(input []byte) (SignedTransaction, error) {
	if input == nil {
		var obj SignedTransaction
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeSignedTransaction(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type StructTag struct {
	Address AccountAddress
	Module Identifier
	Name Identifier
	TypeParams []TypeTag
}

func (obj *StructTag) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Address.Serialize(serializer); err != nil { return err }
	if err := obj.Module.Serialize(serializer); err != nil { return err }
	if err := obj.Name.Serialize(serializer); err != nil { return err }
	if err := serialize_vector_TypeTag(obj.TypeParams, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *StructTag) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeStructTag(deserializer serde.Deserializer) (StructTag, error) {
	var obj StructTag
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.Address = val } else { return obj, err }
	if val, err := DeserializeIdentifier(deserializer); err == nil { obj.Module = val } else { return obj, err }
	if val, err := DeserializeIdentifier(deserializer); err == nil { obj.Name = val } else { return obj, err }
	if val, err := deserialize_vector_TypeTag(deserializer); err == nil { obj.TypeParams = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeStructTag(input []byte) (StructTag, error) {
	if input == nil {
		var obj StructTag
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeStructTag(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Transaction interface {
	isTransaction()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
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

func BcsDeserializeTransaction(input []byte) (Transaction, error) {
	if input == nil {
		var obj Transaction
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTransaction(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type Transaction__UserTransaction struct {
	Value SignedTransaction
}

func (*Transaction__UserTransaction) isTransaction() {}

func (obj *Transaction__UserTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Transaction__UserTransaction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Transaction__UserTransaction(deserializer serde.Deserializer) (Transaction__UserTransaction, error) {
	var obj Transaction__UserTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeSignedTransaction(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Transaction__GenesisTransaction struct {
	Value WriteSetPayload
}

func (*Transaction__GenesisTransaction) isTransaction() {}

func (obj *Transaction__GenesisTransaction) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Transaction__GenesisTransaction) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Transaction__GenesisTransaction(deserializer serde.Deserializer) (Transaction__GenesisTransaction, error) {
	var obj Transaction__GenesisTransaction
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeWriteSetPayload(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type Transaction__BlockMetadata struct {
	Value BlockMetadata
}

func (*Transaction__BlockMetadata) isTransaction() {}

func (obj *Transaction__BlockMetadata) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *Transaction__BlockMetadata) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_Transaction__BlockMetadata(deserializer serde.Deserializer) (Transaction__BlockMetadata, error) {
	var obj Transaction__BlockMetadata
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeBlockMetadata(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionArgument interface {
	isTransactionArgument()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
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

func BcsDeserializeTransactionArgument(input []byte) (TransactionArgument, error) {
	if input == nil {
		var obj TransactionArgument
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTransactionArgument(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionArgument__U8 uint8

func (*TransactionArgument__U8) isTransactionArgument() {}

func (obj *TransactionArgument__U8) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := serializer.SerializeU8(((uint8)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionArgument__U8) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionArgument__U8(deserializer serde.Deserializer) (TransactionArgument__U8, error) {
	var obj uint8
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (TransactionArgument__U8)(obj), err }
	if val, err := deserializer.DeserializeU8(); err == nil { obj = val } else { return ((TransactionArgument__U8)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__U8)(obj), nil
}

type TransactionArgument__U64 uint64

func (*TransactionArgument__U64) isTransactionArgument() {}

func (obj *TransactionArgument__U64) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := serializer.SerializeU64(((uint64)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionArgument__U64) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionArgument__U64(deserializer serde.Deserializer) (TransactionArgument__U64, error) {
	var obj uint64
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (TransactionArgument__U64)(obj), err }
	if val, err := deserializer.DeserializeU64(); err == nil { obj = val } else { return ((TransactionArgument__U64)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__U64)(obj), nil
}

type TransactionArgument__U128 serde.Uint128

func (*TransactionArgument__U128) isTransactionArgument() {}

func (obj *TransactionArgument__U128) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	if err := serializer.SerializeU128(((serde.Uint128)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionArgument__U128) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionArgument__U128(deserializer serde.Deserializer) (TransactionArgument__U128, error) {
	var obj serde.Uint128
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (TransactionArgument__U128)(obj), err }
	if val, err := deserializer.DeserializeU128(); err == nil { obj = val } else { return ((TransactionArgument__U128)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__U128)(obj), nil
}

type TransactionArgument__Address struct {
	Value AccountAddress
}

func (*TransactionArgument__Address) isTransactionArgument() {}

func (obj *TransactionArgument__Address) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionArgument__Address) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionArgument__Address(deserializer serde.Deserializer) (TransactionArgument__Address, error) {
	var obj TransactionArgument__Address
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionArgument__U8Vector []byte

func (*TransactionArgument__U8Vector) isTransactionArgument() {}

func (obj *TransactionArgument__U8Vector) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(4)
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionArgument__U8Vector) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionArgument__U8Vector(deserializer serde.Deserializer) (TransactionArgument__U8Vector, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (TransactionArgument__U8Vector)(obj), err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj = val } else { return ((TransactionArgument__U8Vector)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__U8Vector)(obj), nil
}

type TransactionArgument__Bool bool

func (*TransactionArgument__Bool) isTransactionArgument() {}

func (obj *TransactionArgument__Bool) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(5)
	if err := serializer.SerializeBool(((bool)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionArgument__Bool) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionArgument__Bool(deserializer serde.Deserializer) (TransactionArgument__Bool, error) {
	var obj bool
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (TransactionArgument__Bool)(obj), err }
	if val, err := deserializer.DeserializeBool(); err == nil { obj = val } else { return ((TransactionArgument__Bool)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (TransactionArgument__Bool)(obj), nil
}

type TransactionAuthenticator interface {
	isTransactionAuthenticator()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
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

func BcsDeserializeTransactionAuthenticator(input []byte) (TransactionAuthenticator, error) {
	if input == nil {
		var obj TransactionAuthenticator
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTransactionAuthenticator(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionAuthenticator__Ed25519 struct {
	PublicKey Ed25519PublicKey
	Signature Ed25519Signature
}

func (*TransactionAuthenticator__Ed25519) isTransactionAuthenticator() {}

func (obj *TransactionAuthenticator__Ed25519) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.PublicKey.Serialize(serializer); err != nil { return err }
	if err := obj.Signature.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionAuthenticator__Ed25519) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionAuthenticator__Ed25519(deserializer serde.Deserializer) (TransactionAuthenticator__Ed25519, error) {
	var obj TransactionAuthenticator__Ed25519
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeEd25519PublicKey(deserializer); err == nil { obj.PublicKey = val } else { return obj, err }
	if val, err := DeserializeEd25519Signature(deserializer); err == nil { obj.Signature = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionAuthenticator__MultiEd25519 struct {
	PublicKey MultiEd25519PublicKey
	Signature MultiEd25519Signature
}

func (*TransactionAuthenticator__MultiEd25519) isTransactionAuthenticator() {}

func (obj *TransactionAuthenticator__MultiEd25519) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := obj.PublicKey.Serialize(serializer); err != nil { return err }
	if err := obj.Signature.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionAuthenticator__MultiEd25519) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionAuthenticator__MultiEd25519(deserializer serde.Deserializer) (TransactionAuthenticator__MultiEd25519, error) {
	var obj TransactionAuthenticator__MultiEd25519
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeMultiEd25519PublicKey(deserializer); err == nil { obj.PublicKey = val } else { return obj, err }
	if val, err := DeserializeMultiEd25519Signature(deserializer); err == nil { obj.Signature = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionPayload interface {
	isTransactionPayload()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
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

func BcsDeserializeTransactionPayload(input []byte) (TransactionPayload, error) {
	if input == nil {
		var obj TransactionPayload
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTransactionPayload(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TransactionPayload__WriteSet struct {
	Value WriteSetPayload
}

func (*TransactionPayload__WriteSet) isTransactionPayload() {}

func (obj *TransactionPayload__WriteSet) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionPayload__WriteSet) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionPayload__WriteSet(deserializer serde.Deserializer) (TransactionPayload__WriteSet, error) {
	var obj TransactionPayload__WriteSet
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeWriteSetPayload(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionPayload__Script struct {
	Value Script
}

func (*TransactionPayload__Script) isTransactionPayload() {}

func (obj *TransactionPayload__Script) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionPayload__Script) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionPayload__Script(deserializer serde.Deserializer) (TransactionPayload__Script, error) {
	var obj TransactionPayload__Script
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeScript(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TransactionPayload__Module struct {
	Value Module
}

func (*TransactionPayload__Module) isTransactionPayload() {}

func (obj *TransactionPayload__Module) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TransactionPayload__Module) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TransactionPayload__Module(deserializer serde.Deserializer) (TransactionPayload__Module, error) {
	var obj TransactionPayload__Module
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeModule(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TravelRuleMetadata interface {
	isTravelRuleMetadata()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
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

func BcsDeserializeTravelRuleMetadata(input []byte) (TravelRuleMetadata, error) {
	if input == nil {
		var obj TravelRuleMetadata
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTravelRuleMetadata(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TravelRuleMetadata__TravelRuleMetadataVersion0 struct {
	Value TravelRuleMetadataV0
}

func (*TravelRuleMetadata__TravelRuleMetadataVersion0) isTravelRuleMetadata() {}

func (obj *TravelRuleMetadata__TravelRuleMetadataVersion0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TravelRuleMetadata__TravelRuleMetadataVersion0) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TravelRuleMetadata__TravelRuleMetadataVersion0(deserializer serde.Deserializer) (TravelRuleMetadata__TravelRuleMetadataVersion0, error) {
	var obj TravelRuleMetadata__TravelRuleMetadataVersion0
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeTravelRuleMetadataV0(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TravelRuleMetadataV0 struct {
	OffChainReferenceId *string
}

func (obj *TravelRuleMetadataV0) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_option_str(obj.OffChainReferenceId, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TravelRuleMetadataV0) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeTravelRuleMetadataV0(deserializer serde.Deserializer) (TravelRuleMetadataV0, error) {
	var obj TravelRuleMetadataV0
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_option_str(deserializer); err == nil { obj.OffChainReferenceId = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeTravelRuleMetadataV0(input []byte) (TravelRuleMetadataV0, error) {
	if input == nil {
		var obj TravelRuleMetadataV0
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTravelRuleMetadataV0(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TypeTag interface {
	isTypeTag()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
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

func BcsDeserializeTypeTag(input []byte) (TypeTag, error) {
	if input == nil {
		var obj TypeTag
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeTypeTag(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type TypeTag__Bool struct {
}

func (*TypeTag__Bool) isTypeTag() {}

func (obj *TypeTag__Bool) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Bool) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__Bool(deserializer serde.Deserializer) (TypeTag__Bool, error) {
	var obj TypeTag__Bool
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U8 struct {
}

func (*TypeTag__U8) isTypeTag() {}

func (obj *TypeTag__U8) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__U8) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__U8(deserializer serde.Deserializer) (TypeTag__U8, error) {
	var obj TypeTag__U8
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U64 struct {
}

func (*TypeTag__U64) isTypeTag() {}

func (obj *TypeTag__U64) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(2)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__U64) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__U64(deserializer serde.Deserializer) (TypeTag__U64, error) {
	var obj TypeTag__U64
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__U128 struct {
}

func (*TypeTag__U128) isTypeTag() {}

func (obj *TypeTag__U128) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(3)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__U128) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__U128(deserializer serde.Deserializer) (TypeTag__U128, error) {
	var obj TypeTag__U128
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Address struct {
}

func (*TypeTag__Address) isTypeTag() {}

func (obj *TypeTag__Address) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(4)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Address) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__Address(deserializer serde.Deserializer) (TypeTag__Address, error) {
	var obj TypeTag__Address
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Signer struct {
}

func (*TypeTag__Signer) isTypeTag() {}

func (obj *TypeTag__Signer) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(5)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Signer) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__Signer(deserializer serde.Deserializer) (TypeTag__Signer, error) {
	var obj TypeTag__Signer
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Vector struct {
	Value TypeTag
}

func (*TypeTag__Vector) isTypeTag() {}

func (obj *TypeTag__Vector) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(6)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Vector) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__Vector(deserializer serde.Deserializer) (TypeTag__Vector, error) {
	var obj TypeTag__Vector
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeTypeTag(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type TypeTag__Struct struct {
	Value StructTag
}

func (*TypeTag__Struct) isTypeTag() {}

func (obj *TypeTag__Struct) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(7)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *TypeTag__Struct) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_TypeTag__Struct(deserializer serde.Deserializer) (TypeTag__Struct, error) {
	var obj TypeTag__Struct
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeStructTag(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type UnstructuredBytesMetadata struct {
	Metadata *[]byte
}

func (obj *UnstructuredBytesMetadata) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_option_bytes(obj.Metadata, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *UnstructuredBytesMetadata) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeUnstructuredBytesMetadata(deserializer serde.Deserializer) (UnstructuredBytesMetadata, error) {
	var obj UnstructuredBytesMetadata
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_option_bytes(deserializer); err == nil { obj.Metadata = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeUnstructuredBytesMetadata(input []byte) (UnstructuredBytesMetadata, error) {
	if input == nil {
		var obj UnstructuredBytesMetadata
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeUnstructuredBytesMetadata(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type WriteOp interface {
	isWriteOp()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
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

func BcsDeserializeWriteOp(input []byte) (WriteOp, error) {
	if input == nil {
		var obj WriteOp
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeWriteOp(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type WriteOp__Deletion struct {
}

func (*WriteOp__Deletion) isWriteOp() {}

func (obj *WriteOp__Deletion) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *WriteOp__Deletion) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_WriteOp__Deletion(deserializer serde.Deserializer) (WriteOp__Deletion, error) {
	var obj WriteOp__Deletion
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type WriteOp__Value []byte

func (*WriteOp__Value) isWriteOp() {}

func (obj *WriteOp__Value) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := serializer.SerializeBytes((([]byte)(*obj))); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *WriteOp__Value) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_WriteOp__Value(deserializer serde.Deserializer) (WriteOp__Value, error) {
	var obj []byte
	if err := deserializer.IncreaseContainerDepth(); err != nil { return (WriteOp__Value)(obj), err }
	if val, err := deserializer.DeserializeBytes(); err == nil { obj = val } else { return ((WriteOp__Value)(obj)), err }
	deserializer.DecreaseContainerDepth()
	return (WriteOp__Value)(obj), nil
}

type WriteSet struct {
	Value WriteSetMut
}

func (obj *WriteSet) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *WriteSet) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeWriteSet(deserializer serde.Deserializer) (WriteSet, error) {
	var obj WriteSet
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeWriteSetMut(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeWriteSet(input []byte) (WriteSet, error) {
	if input == nil {
		var obj WriteSet
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeWriteSet(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type WriteSetMut struct {
	WriteSet []struct {Field0 AccessPath; Field1 WriteOp}
}

func (obj *WriteSetMut) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	if err := serialize_vector_tuple2_AccessPath_WriteOp(obj.WriteSet, serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *WriteSetMut) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func DeserializeWriteSetMut(deserializer serde.Deserializer) (WriteSetMut, error) {
	var obj WriteSetMut
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := deserialize_vector_tuple2_AccessPath_WriteOp(deserializer); err == nil { obj.WriteSet = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

func BcsDeserializeWriteSetMut(input []byte) (WriteSetMut, error) {
	if input == nil {
		var obj WriteSetMut
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeWriteSetMut(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type WriteSetPayload interface {
	isWriteSetPayload()
	Serialize(serializer serde.Serializer) error
	BcsSerialize() ([]byte, error)
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

func BcsDeserializeWriteSetPayload(input []byte) (WriteSetPayload, error) {
	if input == nil {
		var obj WriteSetPayload
		return obj, fmt.Errorf("Cannot deserialize null array")
	}
	deserializer := bcs.NewDeserializer(input);
	obj, err := DeserializeWriteSetPayload(deserializer)
	if err == nil && deserializer.GetBufferOffset() < uint64(len(input)) {
		return obj, fmt.Errorf("Some input bytes were not read")
	}
	return obj, err
}

type WriteSetPayload__Direct struct {
	Value ChangeSet
}

func (*WriteSetPayload__Direct) isWriteSetPayload() {}

func (obj *WriteSetPayload__Direct) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(0)
	if err := obj.Value.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *WriteSetPayload__Direct) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_WriteSetPayload__Direct(deserializer serde.Deserializer) (WriteSetPayload__Direct, error) {
	var obj WriteSetPayload__Direct
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeChangeSet(deserializer); err == nil { obj.Value = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
	return obj, nil
}

type WriteSetPayload__Script struct {
	ExecuteAs AccountAddress
	Script Script
}

func (*WriteSetPayload__Script) isWriteSetPayload() {}

func (obj *WriteSetPayload__Script) Serialize(serializer serde.Serializer) error {
	if err := serializer.IncreaseContainerDepth(); err != nil { return err }
	serializer.SerializeVariantIndex(1)
	if err := obj.ExecuteAs.Serialize(serializer); err != nil { return err }
	if err := obj.Script.Serialize(serializer); err != nil { return err }
	serializer.DecreaseContainerDepth()
	return nil
}

func (obj *WriteSetPayload__Script) BcsSerialize() ([]byte, error) {
	if obj == nil {
		return nil, fmt.Errorf("Cannot serialize null object")
	}
	serializer := bcs.NewSerializer();
	if err := obj.Serialize(serializer); err != nil { return nil, err }
	return serializer.GetBytes(), nil
}

func load_WriteSetPayload__Script(deserializer serde.Deserializer) (WriteSetPayload__Script, error) {
	var obj WriteSetPayload__Script
	if err := deserializer.IncreaseContainerDepth(); err != nil { return obj, err }
	if val, err := DeserializeAccountAddress(deserializer); err == nil { obj.ExecuteAs = val } else { return obj, err }
	if val, err := DeserializeScript(deserializer); err == nil { obj.Script = val } else { return obj, err }
	deserializer.DecreaseContainerDepth()
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
	if value != nil {
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
	if tag {
		value := new([]byte)
		if val, err := deserializer.DeserializeBytes(); err == nil { *value = val } else { return nil, err }
	        return value, nil
	} else {
		return nil, nil
	}
}

func serialize_option_str(value *string, serializer serde.Serializer) error {
	if value != nil {
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
	if tag {
		value := new(string)
		if val, err := deserializer.DeserializeStr(); err == nil { *value = val } else { return nil, err }
	        return value, nil
	} else {
		return nil, nil
	}
}

func serialize_option_u64(value *uint64, serializer serde.Serializer) error {
	if value != nil {
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
	if tag {
		value := new(uint64)
		if val, err := deserializer.DeserializeU64(); err == nil { *value = val } else { return nil, err }
	        return value, nil
	} else {
		return nil, nil
	}
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

