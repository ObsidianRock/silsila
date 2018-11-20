package blockchain

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

type BlockChain struct {
	LashHash []byte
	Database *badger.DB
}

type BlockChainIterator struct {
	CurrentHash []byte 
	Database *badger.DB
}


func InitBlockChain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(txn *badger.Txn) err {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {

			fmt.Println("NO Preexisting Blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis Proved")

			err = txn.Set(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = txn.Set([]byte("lh"), genesis.Hash)
			lashHash = genesis.Hash

			return err
		} else {
			item, err := txn.Get([]byte("lh"))
			if err != nil {
				log.Panic(err)
			}
			lashHash, err := item.Value()
			return err
		}
	})

	if err != nil {
		log.Panic(err)
	}

	blockChain := Blockchain{lashHash, db}
	return &blockChain
}


func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte 

	err := chain.Database.View(func(txn *badget.Txn) err {
		item, err := txn.Get([]byte("lh"))
		if err != nil {
			log.Panic(err)
		}
		
		lastHash, err := item.Value()
		return err 
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = txn.Set([]byte("lh"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		chain.lastHash = newBlock.Hash 
		return err 

	})

	if err != nil {
		log.Panic(err)
	}
}

func (chain *BlockChain) Iterator() *BlockChainIterator{
	iter := &BlockChainIterator{chain.LastHash, chain.Database}

	return iter 
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block 

	err := 	iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		if err != nil {
			log.Panic(err)
		}
		

		encodedBlock, err := item.Value()
		block := Deserialize(encodedBlock)
		return err 
	})
	if err != nil {
		log.Panic(err)
	}
	
	item.CurrentHash = block.Prevhash

	return block 
}