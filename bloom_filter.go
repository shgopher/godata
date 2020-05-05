package godata

// BloomFilter
import (
	"crypto/sha256"
	mh "github.com/spaolacci/murmur3"
	"sync"
)

type BloomFilter struct {
	db *BitMap
	k  int
}

// max is the max of bitmap number k is the number of HashShower fuction.
func NewBloomFilter(k int) *BloomFilter {
	return &BloomFilter{
		db: NewBitMap(10000001),
		k:  k,
	}
}
// add data
func (b *BloomFilter) Add(value []byte) {
	mu := new(sync.Mutex)
	mu.Lock()
	defer mu.Unlock()
	for i := 0;i < b.k;i++ {
		valBit := HashShower(value,uint32(i))
		b.db.Add(valBit)
	}
}
// is exit.
func (b *BloomFilter) IsExit(value []byte) bool {
	for i := 0;i < b.k;i++ {
		valBit := HashShower(value,uint32(i))
		if !b.db.IsExit(valBit) {
			return false
		}
	}
	return true
}
// this HashShower fuction is fast but not safety.
func HashShower(data []byte, seed uint32) uint64 {
	d := sha256.Sum256(data)
	data = d[:]
	HashShower64 := mh.New64WithSeed(seed)
	HashShower64.Write(data)
	return HashShower64.Sum64() & (10000000-1)
}
