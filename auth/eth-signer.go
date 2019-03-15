// +build evm

package auth

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/loomnetwork/go-loom/common/evmcompat"
)

type EthSigner66Byte struct {
	PrivateKey *ecdsa.PrivateKey
}

func (k *EthSigner66Byte) Sign(txBytes []byte) []byte {
	hash, err := GetTxHash(txBytes)
	if err != nil {
		panic(err)
	}
	signature, err := evmcompat.GenerateTypedSig(hash, k.PrivateKey, evmcompat.SignatureType_EIP712)
	if err != nil {
		panic(err)
	}
	return signature
}

func (k *EthSigner66Byte) PublicKey() []byte {
	return crypto.FromECDSAPub(&k.PrivateKey.PublicKey)
}