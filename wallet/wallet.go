package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
}

func NewWallet() *Wallet {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	return &Wallet{PrivateKey: privateKey, PublicKey: publicKey}
}

func (w *Wallet) Address() []byte {
	// Simplest representation just using public key bytes
	return w.PublicKey
}

func main() {
	wallet := NewWallet()
	fmt.Printf("Private Key: %x\n", wallet.PrivateKey.D.Bytes())
	fmt.Printf("Public Key: %x\n", wallet.PublicKey)
	fmt.Printf("Address: %x\n", wallet.Address())
}
