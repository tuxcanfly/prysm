// Package bls implements a go-wrapper around a library implementing the
// the BLS12-381 curve and signature scheme. This package exposes a public API for
// verifying and aggregating BLS signatures used by Ethereum 2.0.
package bls

import (
	"github.com/prysmaticlabs/prysm/shared/bls/iface"
	sa "github.com/prysmaticlabs/prysm/shared/bls/supranational"
)

// SecretKeyFromBytes creates a BLS private key from a BigEndian byte slice.
func SecretKeyFromBytes(privKey []byte) (SecretKey, error) {
	return sa.SecretKeyFromBytes(privKey)
}

// PublicKeyFromBytes creates a BLS public key from a  BigEndian byte slice.
func PublicKeyFromBytes(pubKey []byte) (PublicKey, error) {
	return sa.PublicKeyFromBytes(pubKey)
}

// SignatureFromBytes creates a BLS signature from a LittleEndian byte slice.
func SignatureFromBytes(sig []byte) (Signature, error) {
	return sa.SignatureFromBytes(sig)
}

// AggregateSignatures converts a list of signatures into a single, aggregated sig.
func AggregateSignatures(sigs []iface.Signature) iface.Signature {
	return sa.AggregateSignatures(sigs)
}

// VerifyMultipleSignatures verifies multiple signatures for distinct messages securely.
func VerifyMultipleSignatures(sigs []iface.Signature, msgs [][32]byte, pubKeys []iface.PublicKey) (bool, error) {
	ag := AggregateSignatures(sigs)
	return ag.AggregateVerify(pubKeys, msgs), nil
}

// NewAggregateSignature creates a blank aggregate signature.
func NewAggregateSignature() iface.Signature {
	return sa.NewAggregateSignature()
}

// RandKey creates a new private key using a random input.
func RandKey() iface.SecretKey {
	return sa.RandKey()
}
