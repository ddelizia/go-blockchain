package module

import (
	"time"
	"crypto/sha256"
	"fmt"
	"strings"
)

type Block struct {
	Data         []Transaction
	Index        int
	Timestamp    time.Time
	Hash         string
	PreviousHash string
	Proof		 int
	PrevProof	 int
}

func (block *Block) GetTransactions () []Transaction{
	return block.Data
}

func sha256Algo (value string) string{
	hasher := sha256.New()
	hasher.Write([]byte(value))
	return fmt.Sprintf("%x",hasher.Sum(nil))
}

func (block *Block) hashBlock(){
	str := string(block.Proof)
	str += string(block.PrevProof)
	str += string(block.Index)
	str += string(block.Timestamp.String())
	str += string(fmt.Sprintf("",block.Data))
	str += string(block.PreviousHash)

	block.Hash = sha256Algo(str)
}

func verifyProof(prevProof int, proof int) bool{
	str := string(proof)
	str += string(prevProof)
	hash := sha256Algo(str)
	return strings.HasPrefix(hash, "0000")
}

