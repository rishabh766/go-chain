package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Wallets struct {
	Wallets map[string]*Wallet
}

func NewWallets() *Wallets {
	ws := Wallets{}
	ws.Wallets = make(map[string]*Wallet)
	return &ws
}

func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := fmt.Sprintf("%x", wallet.Address())
	ws.Wallets[address] = wallet
	return address
}

func (ws *Wallets) SaveToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(ws)
	if err != nil {
		panic(err)
	}
}

func (ws *Wallets) LoadFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(ws)
	if err != nil {
		panic(err)
	}
}
