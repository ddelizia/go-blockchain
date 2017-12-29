package module

import (
	"time"
	"fmt"
	log "github.com/Sirupsen/logrus"
	//"encoding/base64"
	"errors"
)

type Transaction struct {
	Sender    int     `json:"sender"`
	Recipient int     `json:"recipient"`
	Amount    float64 `json:"amount"`
}

type BlockChain struct {
	Blocks       []Block
	CreationTime time.Time
	state        map[string]float64
	Nodes        []int
}

func (b *BlockChain) Init(node int) {

	b.Nodes = append(b.Nodes, node)
	// Create genesis block
	if (len(b.Blocks) == 0) {
		b.CreationTime = time.Now()
		log.Info("Creating the first block")
		var block = Block{
			Index:        0,
			Timestamp:    time.Now(),
			PreviousHash: "0000000000000000000000000000000000000000000000000000000000000000",
			PrevProof:    0,
		}
		b.Blocks = append(b.Blocks, block)
		b.AddTransactionToCurrentBlock(0, 1, 1000)
		b.Mine(0)
		log.Info(fmt.Sprintf("Genesis block created: ", b))
	}
}

func (b *BlockChain) AddTransactionToCurrentBlock(sender int, recipient int, amount float64) (*Transaction, error) {
	transaction := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
	t, err := b.updateState(transaction)

	if err == nil {
		lastBlock := b.LastBlock()
		lastBlock.Data = append(lastBlock.Data, transaction)

		log.Info("Transaction added to block: ", lastBlock.Index)
	}
	return t, err

}

func (b *BlockChain) LastBlock() *Block {
	return &b.Blocks[len(b.Blocks)-1]
}

func (b *BlockChain) AddBlock() *Block {
	block := Block{
		Index:     (len(b.Blocks)),
		Timestamp: time.Now(),
	}
	b.Blocks = append(b.Blocks, block)

	log.Info(fmt.Sprintf("New block created %d - Block information:%s", block.Index, block))

	return &block
}

func (b *BlockChain) updateState(t Transaction) (*Transaction, error) {

	if t.Amount == 0 {
		return &t, errors.New("Amount = 0 is not valid")
	}

	if t.Recipient == t.Sender {
		return &t, errors.New("Sender = Recipient: transaction not valid")
	}

	valueSender, ok := b.state[string(t.Sender)]
	if ok {
		if valueSender > t.Amount {
			return &t, errors.New("Sender has not amount available to send")
		}
		b.state[string(t.Sender)] = valueSender - t.Amount
	}
	valueRecipient, ok := b.state[string(t.Recipient)]
	if ok {
		b.state[string(t.Recipient)] = valueRecipient + t.Amount
	}

	log.Info(fmt.Sprintf("Current state: ", len(b.state)))
	return &t, nil
}

func (b *BlockChain) Mine(node int) {
	lastBlock := b.LastBlock()
	i := 0
	for ; ; i++ { // while - loop
		if verifyProof(lastBlock.PrevProof, i) == true {
			break
		}
	}
	lastBlock.Proof = i
	b.AddTransactionToCurrentBlock(0, node, 1)
	lastBlock.hashBlock()
	b.AddBlock()
	newBlock := b.LastBlock()
	newBlock.PrevProof = lastBlock.Proof
	newBlock.PreviousHash = lastBlock.Hash
}
