package assignment01IBC
import (
"crypto/sha256"
"fmt"
)

type BlockData struct {
Transactions []string
}
type Block struct {
Data        BlockData
PrevPointer *Block
PrevHash    string
CurrentHash string
}

func CalculateHash(inputBlock *Block) string {
   var block_hash string = fmt.Sprintf("%v", *inputBlock)
   hash := sha256.Sum256([]byte(block_hash))
   return fmt.Sprintf("%x", hash)
}
 func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {
     if chainHead == nil {
          var temp_block Block
          chainHead = &temp_block
          chainHead.PrevPointer = nil
          chainHead.Data = dataToInsert
          cur_hash := CalculateHash(chainHead)
          chainHead.CurrentHash = cur_hash
          chainHead.PrevHash = ""
     }  else {
         var temp_Block Block
         temp_Block.PrevPointer = chainHead
         temp_Block.Data = dataToInsert
         temp_Block.PrevHash = chainHead.CurrentHash
         cur_hash := CalculateHash(chainHead)
         temp_Block.CurrentHash = cur_hash
         chainHead = &temp_Block
     }
   return chainHead
 }

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
    ptr := chainHead
    for ptr != nil {
      trans := ptr.Data.Transactions
      for i := range trans{
        if oldTrans == trans[i]  {
          trans[i] = newTrans
          fmt.Println("Block Chain Compromised")
          break
        }
        i++
      }
      ptr = ptr.PrevPointer
    }
}

func ListBlocks(chainHead *Block) {
    var temp *Block
    temp = chainHead
    for temp != nil {
      fmt.Println("Transaction = ", temp.Data.Transactions)
      temp = temp.PrevPointer
    }
}

func VerifyChain(chainHead *Block) {
    ptr := chainHead.PrevPointer
    for ptr.PrevPointer != nil {
      if ptr.CurrentHash != chainHead.PrevHash {
        fmt.Println("Block Chain not verified")
        break
      }
    }
    fmt.Println("Block Chain Verified")
}
