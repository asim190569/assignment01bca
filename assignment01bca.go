// You can edit this code!
// Click here and start typing.

// 19I 0569
// Asim Kamran

package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Block struct {
	nonce       int
	transaction string
	prevHash    string
	currHash    string
}

func Newblock(nonce int, transaction string) *Block {
	sk := new(Block)
	sk.nonce = nonce
	sk.transaction = transaction
	return sk
}

type Blockchain struct {
	list []*Block
}

func (ls *Blockchain) Addblock(nonce int, transaction string) *Block {
	st := Newblock(nonce, transaction)

	if VerifyChain(ls) {
		ls.list = append(ls.list, st)
		Blockhash(ls)

		return st
	} else {
		return nil
	}
}

func ListBlocks(bk *Blockchain) {

	for i := 0; i < len(bk.list); i++ {
		fmt.Printf("%s Block No %d %s\n", strings.Repeat("*", 25), i, strings.Repeat("*", 25))
		fmt.Println("Block Nonce    ", bk.list[i].nonce)
		fmt.Println("Block Transaction   ", bk.list[i].transaction)
		fmt.Println("Block Previous hash   ", bk.list[i].prevHash)
		fmt.Println("Block Current hash   ", bk.list[i].currHash)

	}

}

func CalculateHash(hashh string) string {

	return fmt.Sprintf("%x", sha256.Sum256([]byte(hashh)))
}

func Blockhash(bk *Blockchain) {

	for i := 0; i < len(bk.list); i++ {

		bk.list[i].currHash = CalculateHash(strconv.Itoa(bk.list[i].nonce) + bk.list[i].transaction + bk.list[i].prevHash)
		if i < len(bk.list)-1 {
			bk.list[i+1].prevHash = CalculateHash(strconv.Itoa(bk.list[i].nonce) + bk.list[i].transaction + bk.list[i].prevHash)
		}
	}

}
func VerifyChain(bk *Blockchain) bool {
	var st = ""
	for i := 0; i < len(bk.list); i++ {

		st = CalculateHash(strconv.Itoa(bk.list[i].nonce) + bk.list[i].transaction + bk.list[i].prevHash)

		if st != bk.list[i].currHash {
			fmt.Printf("Tempered Block no %d\n", i)
			return false

		}

	}

	return true

}
func ChangeBlock(bk *Blockchain, nonce int, transaction string) {

	for i := 0; i < len(bk.list); i++ {
		if nonce == bk.list[i].nonce {

			bk.list[i].transaction = transaction
			fmt.Println("change done ")
			return
		}
	}

	fmt.Println("block not found!")

}
