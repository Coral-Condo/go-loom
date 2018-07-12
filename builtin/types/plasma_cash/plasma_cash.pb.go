// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto

/*
Package plasma_cash is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto

It has these top-level messages:
	PlasmaCashCoin
	PlasmaCashAccount
	PlasmaBlock
	PlasmaTx
	PlasmaBookKeeping
	GetCurrentBlockRequest
	GetCurrentBlockResponse
	GetBlockRequest
	GetBlockResponse
	SubmitBlockToMainnetRequest
	SubmitBlockToMainnetResponse
	PlasmaTxRequest
	PlasmaTxResponse
	DepositRequest
	DepositResponse
	PlasmaCashExitCoinRequest
	PlasmaCashWithdrawCoinRequest
	Pending
	PlasmaCashParams
	PlasmaCashInitRequest
	PlasmaCashBalanceOfRequest
	PlasmaCashBalanceOfResponse
*/
package plasma_cash

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PlasmaCashCoinState int32

const (
	PlasmaCashCoinState_DEPOSITED  PlasmaCashCoinState = 0
	PlasmaCashCoinState_EXITING    PlasmaCashCoinState = 1
	PlasmaCashCoinState_CHALLENGED PlasmaCashCoinState = 2
	PlasmaCashCoinState_EXITED     PlasmaCashCoinState = 3
)

var PlasmaCashCoinState_name = map[int32]string{
	0: "DEPOSITED",
	1: "EXITING",
	2: "CHALLENGED",
	3: "EXITED",
}
var PlasmaCashCoinState_value = map[string]int32{
	"DEPOSITED":  0,
	"EXITING":    1,
	"CHALLENGED": 2,
	"EXITED":     3,
}

func (x PlasmaCashCoinState) String() string {
	return proto.EnumName(PlasmaCashCoinState_name, int32(x))
}
func (PlasmaCashCoinState) EnumDescriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{0} }

// Plasma Cash coin holds a single ERC721 token
type PlasmaCashCoin struct {
	// Unique ID
	Slot  uint64              `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	State PlasmaCashCoinState `protobuf:"varint,2,opt,name=state,proto3,enum=PlasmaCashCoinState" json:"state,omitempty"`
	// ERC721 token ID
	Token *types.BigUInt `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	// ERC721 token contract address
	Contract *types.Address `protobuf:"bytes,4,opt,name=contract" json:"contract,omitempty"`
}

func (m *PlasmaCashCoin) Reset()                    { *m = PlasmaCashCoin{} }
func (m *PlasmaCashCoin) String() string            { return proto.CompactTextString(m) }
func (*PlasmaCashCoin) ProtoMessage()               {}
func (*PlasmaCashCoin) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{0} }

func (m *PlasmaCashCoin) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *PlasmaCashCoin) GetState() PlasmaCashCoinState {
	if m != nil {
		return m.State
	}
	return PlasmaCashCoinState_DEPOSITED
}

func (m *PlasmaCashCoin) GetToken() *types.BigUInt {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *PlasmaCashCoin) GetContract() *types.Address {
	if m != nil {
		return m.Contract
	}
	return nil
}

type PlasmaCashAccount struct {
	Owner *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	// Address of ERC721 contract the tokens in the Plasma coins originated from
	Contract *types.Address `protobuf:"bytes,2,opt,name=contract" json:"contract,omitempty"`
	// Plasma coins in this account, identified by their slot number.
	Slots []uint64 `protobuf:"varint,3,rep,packed,name=slots" json:"slots,omitempty"`
}

func (m *PlasmaCashAccount) Reset()                    { *m = PlasmaCashAccount{} }
func (m *PlasmaCashAccount) String() string            { return proto.CompactTextString(m) }
func (*PlasmaCashAccount) ProtoMessage()               {}
func (*PlasmaCashAccount) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{1} }

func (m *PlasmaCashAccount) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PlasmaCashAccount) GetContract() *types.Address {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *PlasmaCashAccount) GetSlots() []uint64 {
	if m != nil {
		return m.Slots
	}
	return nil
}

type PlasmaBlock struct {
	Uid          *types.BigUInt `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Transactions []*PlasmaTx    `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
	Signature    []byte         `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	MerkleHash   []byte         `protobuf:"bytes,4,opt,name=merkle_hash,json=merkleHash,proto3" json:"merkle_hash,omitempty"`
	Hash         []byte         `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	Proof        []byte         `protobuf:"bytes,6,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *PlasmaBlock) Reset()                    { *m = PlasmaBlock{} }
func (m *PlasmaBlock) String() string            { return proto.CompactTextString(m) }
func (*PlasmaBlock) ProtoMessage()               {}
func (*PlasmaBlock) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{2} }

func (m *PlasmaBlock) GetUid() *types.BigUInt {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *PlasmaBlock) GetTransactions() []*PlasmaTx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *PlasmaBlock) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PlasmaBlock) GetMerkleHash() []byte {
	if m != nil {
		return m.MerkleHash
	}
	return nil
}

func (m *PlasmaBlock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PlasmaBlock) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

type PlasmaTx struct {
	Slot          uint64         `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	PreviousBlock *types.BigUInt `protobuf:"bytes,2,opt,name=previous_block,json=previousBlock" json:"previous_block,omitempty"`
	Denomination  *types.BigUInt `protobuf:"bytes,3,opt,name=denomination" json:"denomination,omitempty"`
	NewOwner      *types.Address `protobuf:"bytes,4,opt,name=new_owner,json=newOwner" json:"new_owner,omitempty"`
	Signature     []byte         `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Hash          []byte         `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	MerkleHash    []byte         `protobuf:"bytes,7,opt,name=merkle_hash,json=merkleHash,proto3" json:"merkle_hash,omitempty"`
	Sender        *types.Address `protobuf:"bytes,8,opt,name=sender" json:"sender,omitempty"`
	Proof         []byte         `protobuf:"bytes,9,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *PlasmaTx) Reset()                    { *m = PlasmaTx{} }
func (m *PlasmaTx) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTx) ProtoMessage()               {}
func (*PlasmaTx) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{3} }

func (m *PlasmaTx) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *PlasmaTx) GetPreviousBlock() *types.BigUInt {
	if m != nil {
		return m.PreviousBlock
	}
	return nil
}

func (m *PlasmaTx) GetDenomination() *types.BigUInt {
	if m != nil {
		return m.Denomination
	}
	return nil
}

func (m *PlasmaTx) GetNewOwner() *types.Address {
	if m != nil {
		return m.NewOwner
	}
	return nil
}

func (m *PlasmaTx) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PlasmaTx) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PlasmaTx) GetMerkleHash() []byte {
	if m != nil {
		return m.MerkleHash
	}
	return nil
}

func (m *PlasmaTx) GetSender() *types.Address {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *PlasmaTx) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

type PlasmaBookKeeping struct {
	CurrentHeight *types.BigUInt `protobuf:"bytes,1,opt,name=current_height,json=currentHeight" json:"current_height,omitempty"`
}

func (m *PlasmaBookKeeping) Reset()                    { *m = PlasmaBookKeeping{} }
func (m *PlasmaBookKeeping) String() string            { return proto.CompactTextString(m) }
func (*PlasmaBookKeeping) ProtoMessage()               {}
func (*PlasmaBookKeeping) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{4} }

func (m *PlasmaBookKeeping) GetCurrentHeight() *types.BigUInt {
	if m != nil {
		return m.CurrentHeight
	}
	return nil
}

type GetCurrentBlockRequest struct {
}

func (m *GetCurrentBlockRequest) Reset()                    { *m = GetCurrentBlockRequest{} }
func (m *GetCurrentBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCurrentBlockRequest) ProtoMessage()               {}
func (*GetCurrentBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{5} }

type GetCurrentBlockResponse struct {
	BlockHeight *types.BigUInt `protobuf:"bytes,1,opt,name=block_height,json=blockHeight" json:"block_height,omitempty"`
}

func (m *GetCurrentBlockResponse) Reset()         { *m = GetCurrentBlockResponse{} }
func (m *GetCurrentBlockResponse) String() string { return proto.CompactTextString(m) }
func (*GetCurrentBlockResponse) ProtoMessage()    {}
func (*GetCurrentBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{6}
}

func (m *GetCurrentBlockResponse) GetBlockHeight() *types.BigUInt {
	if m != nil {
		return m.BlockHeight
	}
	return nil
}

type GetBlockRequest struct {
	BlockHeight *types.BigUInt `protobuf:"bytes,1,opt,name=block_height,json=blockHeight" json:"block_height,omitempty"`
}

func (m *GetBlockRequest) Reset()                    { *m = GetBlockRequest{} }
func (m *GetBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBlockRequest) ProtoMessage()               {}
func (*GetBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{7} }

func (m *GetBlockRequest) GetBlockHeight() *types.BigUInt {
	if m != nil {
		return m.BlockHeight
	}
	return nil
}

type GetBlockResponse struct {
	Block *PlasmaBlock `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
}

func (m *GetBlockResponse) Reset()                    { *m = GetBlockResponse{} }
func (m *GetBlockResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBlockResponse) ProtoMessage()               {}
func (*GetBlockResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{8} }

func (m *GetBlockResponse) GetBlock() *PlasmaBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

// This only originates from the validator
type SubmitBlockToMainnetRequest struct {
}

func (m *SubmitBlockToMainnetRequest) Reset()         { *m = SubmitBlockToMainnetRequest{} }
func (m *SubmitBlockToMainnetRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitBlockToMainnetRequest) ProtoMessage()    {}
func (*SubmitBlockToMainnetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{9}
}

type SubmitBlockToMainnetResponse struct {
	MerkleHash []byte `protobuf:"bytes,1,opt,name=merkle_hash,json=merkleHash,proto3" json:"merkle_hash,omitempty"`
}

func (m *SubmitBlockToMainnetResponse) Reset()         { *m = SubmitBlockToMainnetResponse{} }
func (m *SubmitBlockToMainnetResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitBlockToMainnetResponse) ProtoMessage()    {}
func (*SubmitBlockToMainnetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{10}
}

func (m *SubmitBlockToMainnetResponse) GetMerkleHash() []byte {
	if m != nil {
		return m.MerkleHash
	}
	return nil
}

type PlasmaTxRequest struct {
	Plasmatx *PlasmaTx `protobuf:"bytes,1,opt,name=plasmatx" json:"plasmatx,omitempty"`
}

func (m *PlasmaTxRequest) Reset()                    { *m = PlasmaTxRequest{} }
func (m *PlasmaTxRequest) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTxRequest) ProtoMessage()               {}
func (*PlasmaTxRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{11} }

func (m *PlasmaTxRequest) GetPlasmatx() *PlasmaTx {
	if m != nil {
		return m.Plasmatx
	}
	return nil
}

type PlasmaTxResponse struct {
}

func (m *PlasmaTxResponse) Reset()                    { *m = PlasmaTxResponse{} }
func (m *PlasmaTxResponse) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTxResponse) ProtoMessage()               {}
func (*PlasmaTxResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{12} }

// This only originates from the validator
type DepositRequest struct {
	Slot         uint64         `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	DepositBlock *types.BigUInt `protobuf:"bytes,2,opt,name=deposit_block,json=depositBlock" json:"deposit_block,omitempty"`
	// For ERC20 this is the number of coins deposited, for ERC721 this is a token ID.
	Denomination *types.BigUInt `protobuf:"bytes,3,opt,name=denomination" json:"denomination,omitempty"`
	// Entity that made the deposit
	From *types.Address `protobuf:"bytes,4,opt,name=from" json:"from,omitempty"`
	// Contract from which the coins originated (i.e. the currency of the coins)
	Contract *types.Address `protobuf:"bytes,5,opt,name=contract" json:"contract,omitempty"`
}

func (m *DepositRequest) Reset()                    { *m = DepositRequest{} }
func (m *DepositRequest) String() string            { return proto.CompactTextString(m) }
func (*DepositRequest) ProtoMessage()               {}
func (*DepositRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{13} }

func (m *DepositRequest) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *DepositRequest) GetDepositBlock() *types.BigUInt {
	if m != nil {
		return m.DepositBlock
	}
	return nil
}

func (m *DepositRequest) GetDenomination() *types.BigUInt {
	if m != nil {
		return m.Denomination
	}
	return nil
}

func (m *DepositRequest) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *DepositRequest) GetContract() *types.Address {
	if m != nil {
		return m.Contract
	}
	return nil
}

type DepositResponse struct {
}

func (m *DepositResponse) Reset()                    { *m = DepositResponse{} }
func (m *DepositResponse) String() string            { return proto.CompactTextString(m) }
func (*DepositResponse) ProtoMessage()               {}
func (*DepositResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{14} }

type PlasmaCashExitCoinRequest struct {
	Owner *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Slot  uint64         `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (m *PlasmaCashExitCoinRequest) Reset()         { *m = PlasmaCashExitCoinRequest{} }
func (m *PlasmaCashExitCoinRequest) String() string { return proto.CompactTextString(m) }
func (*PlasmaCashExitCoinRequest) ProtoMessage()    {}
func (*PlasmaCashExitCoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{15}
}

func (m *PlasmaCashExitCoinRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PlasmaCashExitCoinRequest) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

type PlasmaCashWithdrawCoinRequest struct {
	Owner *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Slot  uint64         `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (m *PlasmaCashWithdrawCoinRequest) Reset()         { *m = PlasmaCashWithdrawCoinRequest{} }
func (m *PlasmaCashWithdrawCoinRequest) String() string { return proto.CompactTextString(m) }
func (*PlasmaCashWithdrawCoinRequest) ProtoMessage()    {}
func (*PlasmaCashWithdrawCoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{16}
}

func (m *PlasmaCashWithdrawCoinRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PlasmaCashWithdrawCoinRequest) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

type Pending struct {
	Transactions []*PlasmaTx `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *Pending) Reset()                    { *m = Pending{} }
func (m *Pending) String() string            { return proto.CompactTextString(m) }
func (*Pending) ProtoMessage()               {}
func (*Pending) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{17} }

func (m *Pending) GetTransactions() []*PlasmaTx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

// Initialization of state from Genesis.json
type PlasmaCashParams struct {
	BlockInterval uint64 `protobuf:"varint,1,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval,omitempty"`
}

func (m *PlasmaCashParams) Reset()                    { *m = PlasmaCashParams{} }
func (m *PlasmaCashParams) String() string            { return proto.CompactTextString(m) }
func (*PlasmaCashParams) ProtoMessage()               {}
func (*PlasmaCashParams) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{18} }

func (m *PlasmaCashParams) GetBlockInterval() uint64 {
	if m != nil {
		return m.BlockInterval
	}
	return 0
}

type PlasmaCashInitRequest struct {
	Params *PlasmaCashParams `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
}

func (m *PlasmaCashInitRequest) Reset()                    { *m = PlasmaCashInitRequest{} }
func (m *PlasmaCashInitRequest) String() string            { return proto.CompactTextString(m) }
func (*PlasmaCashInitRequest) ProtoMessage()               {}
func (*PlasmaCashInitRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{19} }

func (m *PlasmaCashInitRequest) GetParams() *PlasmaCashParams {
	if m != nil {
		return m.Params
	}
	return nil
}

type PlasmaCashBalanceOfRequest struct {
	Owner    *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Contract *types.Address `protobuf:"bytes,2,opt,name=contract" json:"contract,omitempty"`
}

func (m *PlasmaCashBalanceOfRequest) Reset()         { *m = PlasmaCashBalanceOfRequest{} }
func (m *PlasmaCashBalanceOfRequest) String() string { return proto.CompactTextString(m) }
func (*PlasmaCashBalanceOfRequest) ProtoMessage()    {}
func (*PlasmaCashBalanceOfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{20}
}

func (m *PlasmaCashBalanceOfRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *PlasmaCashBalanceOfRequest) GetContract() *types.Address {
	if m != nil {
		return m.Contract
	}
	return nil
}

type PlasmaCashBalanceOfResponse struct {
	Coins []*PlasmaCashCoin `protobuf:"bytes,1,rep,name=coins" json:"coins,omitempty"`
}

func (m *PlasmaCashBalanceOfResponse) Reset()         { *m = PlasmaCashBalanceOfResponse{} }
func (m *PlasmaCashBalanceOfResponse) String() string { return proto.CompactTextString(m) }
func (*PlasmaCashBalanceOfResponse) ProtoMessage()    {}
func (*PlasmaCashBalanceOfResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{21}
}

func (m *PlasmaCashBalanceOfResponse) GetCoins() []*PlasmaCashCoin {
	if m != nil {
		return m.Coins
	}
	return nil
}

func init() {
	proto.RegisterType((*PlasmaCashCoin)(nil), "PlasmaCashCoin")
	proto.RegisterType((*PlasmaCashAccount)(nil), "PlasmaCashAccount")
	proto.RegisterType((*PlasmaBlock)(nil), "PlasmaBlock")
	proto.RegisterType((*PlasmaTx)(nil), "PlasmaTx")
	proto.RegisterType((*PlasmaBookKeeping)(nil), "PlasmaBookKeeping")
	proto.RegisterType((*GetCurrentBlockRequest)(nil), "GetCurrentBlockRequest")
	proto.RegisterType((*GetCurrentBlockResponse)(nil), "GetCurrentBlockResponse")
	proto.RegisterType((*GetBlockRequest)(nil), "GetBlockRequest")
	proto.RegisterType((*GetBlockResponse)(nil), "GetBlockResponse")
	proto.RegisterType((*SubmitBlockToMainnetRequest)(nil), "SubmitBlockToMainnetRequest")
	proto.RegisterType((*SubmitBlockToMainnetResponse)(nil), "SubmitBlockToMainnetResponse")
	proto.RegisterType((*PlasmaTxRequest)(nil), "PlasmaTxRequest")
	proto.RegisterType((*PlasmaTxResponse)(nil), "PlasmaTxResponse")
	proto.RegisterType((*DepositRequest)(nil), "DepositRequest")
	proto.RegisterType((*DepositResponse)(nil), "DepositResponse")
	proto.RegisterType((*PlasmaCashExitCoinRequest)(nil), "PlasmaCashExitCoinRequest")
	proto.RegisterType((*PlasmaCashWithdrawCoinRequest)(nil), "PlasmaCashWithdrawCoinRequest")
	proto.RegisterType((*Pending)(nil), "Pending")
	proto.RegisterType((*PlasmaCashParams)(nil), "PlasmaCashParams")
	proto.RegisterType((*PlasmaCashInitRequest)(nil), "PlasmaCashInitRequest")
	proto.RegisterType((*PlasmaCashBalanceOfRequest)(nil), "PlasmaCashBalanceOfRequest")
	proto.RegisterType((*PlasmaCashBalanceOfResponse)(nil), "PlasmaCashBalanceOfResponse")
	proto.RegisterEnum("PlasmaCashCoinState", PlasmaCashCoinState_name, PlasmaCashCoinState_value)
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto", fileDescriptorPlasmaCash)
}

var fileDescriptorPlasmaCash = []byte{
	// 926 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x5f, 0x6f, 0x22, 0x37,
	0x10, 0xef, 0x42, 0x20, 0x30, 0xfc, 0x8d, 0x7b, 0x6d, 0x69, 0x2e, 0xd7, 0xa0, 0x55, 0x23, 0xd1,
	0x53, 0x03, 0x55, 0x4e, 0xaa, 0xae, 0x52, 0xa5, 0x2a, 0x04, 0x94, 0xd0, 0xde, 0x05, 0xb4, 0x49,
	0x95, 0xbe, 0x21, 0xb3, 0x38, 0x60, 0x01, 0xf6, 0x76, 0xed, 0x3d, 0xd2, 0x2f, 0xd2, 0xef, 0xd3,
	0x87, 0xbe, 0xf7, 0x13, 0xe4, 0x21, 0x9f, 0xa4, 0x5a, 0xdb, 0xb0, 0xcb, 0x86, 0xf4, 0x4e, 0xbd,
	0x17, 0xe4, 0x99, 0xdf, 0xfc, 0xf1, 0xfc, 0xc6, 0x33, 0x0b, 0xfc, 0x3c, 0xa1, 0x72, 0x1a, 0x8c,
	0x9a, 0x2e, 0x5f, 0xb4, 0xe6, 0x9c, 0x2f, 0x18, 0x91, 0x4b, 0xee, 0xcf, 0x5a, 0x13, 0x7e, 0x1c,
	0x8a, 0xad, 0x51, 0x40, 0xe7, 0x92, 0xb2, 0x96, 0xfc, 0xc3, 0x23, 0xa2, 0xe5, 0xcd, 0xb1, 0x58,
	0xe0, 0xa1, 0x8b, 0xc5, 0x34, 0x7e, 0x6e, 0x7a, 0x3e, 0x97, 0x7c, 0xff, 0x38, 0x16, 0x6b, 0xc2,
	0x27, 0xbc, 0xa5, 0xd4, 0xa3, 0xe0, 0x56, 0x49, 0x4a, 0x50, 0x27, 0x63, 0xfe, 0xdd, 0x7b, 0x52,
	0xeb, 0x94, 0xea, 0x57, 0x7b, 0xd8, 0x7f, 0x5a, 0x50, 0x1e, 0xa8, 0xb4, 0x67, 0x58, 0x4c, 0xcf,
	0x38, 0x65, 0x08, 0xc1, 0x8e, 0x98, 0x73, 0x59, 0xb3, 0xea, 0x56, 0x63, 0xc7, 0x51, 0x67, 0xf4,
	0x12, 0x32, 0x42, 0x62, 0x49, 0x6a, 0xa9, 0xba, 0xd5, 0x28, 0x9f, 0x3c, 0x6b, 0x6e, 0xfa, 0x5c,
	0x85, 0x98, 0xa3, 0x4d, 0xd0, 0x57, 0x90, 0x91, 0x7c, 0x46, 0x58, 0x2d, 0x5d, 0xb7, 0x1a, 0x85,
	0x93, 0x5c, 0xb3, 0x4d, 0x27, 0xbf, 0xf6, 0x98, 0x74, 0xb4, 0x1a, 0x7d, 0x0d, 0x39, 0x97, 0x33,
	0xe9, 0x63, 0x57, 0xd6, 0x76, 0x8c, 0xc9, 0xe9, 0x78, 0xec, 0x13, 0x21, 0x9c, 0x35, 0x62, 0x73,
	0xd8, 0x8b, 0x72, 0x9c, 0xba, 0x2e, 0x0f, 0x98, 0x0c, 0x43, 0xf3, 0x25, 0x23, 0xbe, 0xba, 0x5b,
	0xdc, 0x4f, 0xab, 0x37, 0x42, 0xa7, 0x9e, 0x0a, 0x8d, 0x9e, 0x41, 0x26, 0x2c, 0x4a, 0xd4, 0xd2,
	0xf5, 0x74, 0x63, 0xc7, 0xd1, 0x82, 0xfd, 0x97, 0x05, 0x05, 0x9d, 0xb1, 0x3d, 0xe7, 0xee, 0x0c,
	0xed, 0x43, 0x3a, 0xa0, 0xe3, 0x75, 0xa6, 0x55, 0x11, 0xa1, 0x12, 0x1d, 0x43, 0x51, 0xfa, 0x98,
	0x09, 0xec, 0x4a, 0xca, 0x99, 0xa8, 0xa5, 0xea, 0xe9, 0x46, 0xe1, 0x24, 0x6f, 0x58, 0xb9, 0xbe,
	0x73, 0x36, 0x60, 0x74, 0x00, 0x79, 0x41, 0x27, 0x0c, 0xcb, 0xc0, 0x27, 0x8a, 0x95, 0xa2, 0x13,
	0x29, 0xd0, 0x21, 0x14, 0x16, 0xc4, 0x9f, 0xcd, 0xc9, 0x70, 0x8a, 0xc5, 0x54, 0x51, 0x52, 0x74,
	0x40, 0xab, 0x2e, 0xb0, 0x98, 0x86, 0x0d, 0x51, 0x48, 0x46, 0x21, 0xea, 0x1c, 0xd6, 0xe0, 0xf9,
	0x9c, 0xdf, 0xd6, 0xb2, 0x4a, 0xa9, 0x05, 0xfb, 0x9f, 0x14, 0xe4, 0x56, 0x77, 0xd8, 0xda, 0xc7,
	0x36, 0x94, 0x3d, 0x9f, 0xbc, 0xa3, 0x3c, 0x10, 0xc3, 0x51, 0x58, 0xe6, 0x9a, 0x26, 0x53, 0x5f,
	0x7b, 0xef, 0xe1, 0xfe, 0xb0, 0x34, 0x30, 0x36, 0x8a, 0x09, 0xa7, 0xe4, 0xc5, 0x45, 0xf4, 0x2d,
	0x14, 0xc7, 0x84, 0xf1, 0x05, 0x65, 0x38, 0x2c, 0xef, 0x51, 0x9b, 0x37, 0x50, 0xf4, 0x0a, 0xf2,
	0x8c, 0x2c, 0x87, 0xba, 0x6d, 0x89, 0x76, 0xb7, 0x8b, 0x0f, 0xf7, 0x87, 0xb9, 0x4b, 0xb2, 0xec,
	0x87, 0xa8, 0x93, 0x63, 0xe6, 0xb4, 0x49, 0x58, 0x26, 0x49, 0xd8, 0x8a, 0x8f, 0x6c, 0x8c, 0x8f,
	0x04, 0x89, 0xbb, 0x8f, 0x48, 0xac, 0x43, 0x56, 0x10, 0x36, 0x26, 0x7e, 0x2d, 0x97, 0x78, 0x18,
	0x46, 0x1f, 0x51, 0x9a, 0x8f, 0x53, 0x7a, 0xb3, 0x7a, 0x87, 0x6d, 0xce, 0x67, 0xbf, 0x10, 0xe2,
	0x51, 0x36, 0x09, 0x69, 0x74, 0x03, 0xdf, 0x27, 0x4c, 0x0e, 0xa7, 0x84, 0x4e, 0xa6, 0x32, 0xf9,
	0x4c, 0x34, 0x8d, 0x67, 0xda, 0xe6, 0x42, 0x99, 0x38, 0x25, 0x37, 0x2e, 0xda, 0x35, 0xf8, 0xfc,
	0x9c, 0x48, 0x63, 0xa2, 0x89, 0x26, 0xbf, 0x07, 0x44, 0x48, 0xfb, 0x06, 0xbe, 0x78, 0x84, 0x08,
	0x8f, 0x33, 0x41, 0xd0, 0x8f, 0x50, 0x54, 0x6d, 0x7b, 0x2a, 0x6d, 0xe5, 0xe1, 0xfe, 0xb0, 0xa0,
	0x5c, 0x4c, 0xd2, 0xc2, 0x28, 0x12, 0xec, 0x3e, 0x54, 0xce, 0xc9, 0x46, 0xae, 0x8f, 0x0c, 0xf8,
	0x3d, 0x54, 0xa3, 0x80, 0xe6, 0x8a, 0x36, 0x64, 0xf4, 0xcb, 0xd2, 0xa1, 0x8a, 0xcd, 0xd8, 0x50,
	0x39, 0x1a, 0xb2, 0x5f, 0xc0, 0xf3, 0xab, 0x60, 0xb4, 0xa0, 0xda, 0xf5, 0x9a, 0xbf, 0xc5, 0x94,
	0x31, 0x22, 0x57, 0x04, 0xfc, 0x04, 0x07, 0xdb, 0x61, 0x93, 0x22, 0xd1, 0x6c, 0x2b, 0xd9, 0x6c,
	0xfb, 0x35, 0x54, 0xd6, 0xa3, 0x68, 0x0a, 0x3d, 0x82, 0x9c, 0x5e, 0xaf, 0xf2, 0xce, 0xdc, 0x2c,
	0x36, 0xae, 0x6b, 0xc8, 0x46, 0x50, 0x8d, 0x3c, 0x75, 0x3a, 0xfb, 0x6f, 0x0b, 0xca, 0x1d, 0xe2,
	0x71, 0x41, 0x57, 0x37, 0xdc, 0x3a, 0x5b, 0xc7, 0x50, 0x1a, 0x6b, 0xab, 0xed, 0xa3, 0x15, 0x0e,
	0x86, 0x82, 0xff, 0xcf, 0x18, 0x1d, 0xc0, 0xce, 0xad, 0xcf, 0x17, 0x8f, 0x16, 0xa6, 0xd2, 0x6e,
	0xec, 0xbd, 0xcc, 0x93, 0x2b, 0x75, 0x0f, 0x2a, 0xeb, 0x32, 0x4c, 0x69, 0x7d, 0xf8, 0x32, 0xda,
	0xb2, 0xdd, 0x3b, 0x2a, 0xc3, 0x6d, 0xbe, 0x2a, 0xf2, 0x7d, 0xdb, 0x76, 0x45, 0x42, 0x2a, 0x22,
	0xc1, 0xbe, 0x82, 0x17, 0x51, 0xc0, 0x1b, 0x2a, 0xa7, 0x63, 0x1f, 0x2f, 0x3f, 0x36, 0xe8, 0x6b,
	0xd8, 0x1d, 0x10, 0x36, 0x0e, 0x27, 0x2f, 0xb9, 0x79, 0xad, 0xff, 0xdc, 0xbc, 0xf6, 0x0f, 0xab,
	0x76, 0x86, 0xd7, 0x19, 0x60, 0x1f, 0x2f, 0x04, 0x3a, 0x82, 0xb2, 0x7e, 0xf2, 0x94, 0x49, 0xe2,
	0xbf, 0xc3, 0x73, 0xd3, 0xc5, 0x92, 0xd2, 0xf6, 0x8c, 0xd2, 0x6e, 0xc3, 0x67, 0x91, 0x6b, 0x8f,
	0x45, 0xbd, 0xff, 0x06, 0xb2, 0x9e, 0x8a, 0x64, 0x4a, 0xd8, 0x6b, 0x26, 0x53, 0x38, 0xc6, 0xc0,
	0x1e, 0xc1, 0x7e, 0x84, 0xb5, 0xf1, 0x1c, 0x33, 0x97, 0xf4, 0x6f, 0x3f, 0x94, 0x8a, 0x0f, 0xfa,
	0x9a, 0xd9, 0x1d, 0x78, 0xbe, 0x35, 0x87, 0x99, 0x95, 0x23, 0xc8, 0xb8, 0x9c, 0xae, 0x99, 0xaa,
	0x24, 0xbe, 0xdc, 0x8e, 0x46, 0x5f, 0xbe, 0x85, 0x4f, 0xb7, 0x7c, 0xd2, 0x51, 0x09, 0xf2, 0x9d,
	0xee, 0xa0, 0x7f, 0xd5, 0xbb, 0xee, 0x76, 0xaa, 0x9f, 0xa0, 0x02, 0xec, 0x76, 0x7f, 0xeb, 0x5d,
	0xf7, 0x2e, 0xcf, 0xab, 0x16, 0x2a, 0x03, 0x9c, 0x5d, 0x9c, 0xbe, 0x79, 0xd3, 0xbd, 0x3c, 0xef,
	0x76, 0xaa, 0x29, 0x04, 0x90, 0x0d, 0xc1, 0x6e, 0xa7, 0x9a, 0x1e, 0x65, 0xd5, 0xbf, 0x8b, 0x57,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x62, 0x59, 0x00, 0x0c, 0x09, 0x00, 0x00,
}
