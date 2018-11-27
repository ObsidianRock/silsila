package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/ObsidianRock/silsila/wallet"
)

type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

type TxInput struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    []byte
}

type TxOutputs struct {
	Outputs []TxOutput
}

func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)
	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-1]
	out.PubKeyHash = pubKeyHash
}

func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

func (outs TxOutputs) Serialize() []byte {
	var buffer bytes.Buffer
	encode := gob.NewEncoder(&buffer)
	err := encode.Encode(outs)
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}

func (outs TxOutputs) DeserializeOutputs(data []byte) TxOutputs {

	var outputs TxOutputs
	decode := gob.NewDecoder(bytes.NewReader(data))
	err := decode.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}

	return outputs
}
