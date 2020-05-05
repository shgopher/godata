package godata

// bit map or bit array.
type BitMap struct {
	bits []uint8
	max  uint64
}

// return a new bit map. if max > 2^64-1 ,return nil
func NewBitMap(max uint64) *BitMap {
	return &BitMap{
		max:  max,
		bits: make([]uint8, max>>3+1),
	}
}

// is right return true ,or false
func (b *BitMap) Add(num uint64) {
	if num > b.max {
		return
	}
	byteIndex := num >> 3
	position := num & 0x7
	b.bits[byteIndex] |= 1<< position
}
func (b *BitMap) IsExit(num uint64) bool {
	if num > b.max {
		return false
	}
	byteIndex := num >> 3
	position := num & 0x7
	return (b.bits[byteIndex]) & (1<<position) != 0
 }

// init  the  bitMap.
func (b *BitMap) Init() {
	b.bits = make([]uint8, b.max>>3+1)
}

func(b *BitMap)Delete(num uint64){
	if num > b.max {
		return
	}
	byteIndex := num >> 3
	position := num & 0x7
	if (b.bits[byteIndex]) & (1<< position) != 0{
		b.bits[byteIndex] ^= (1<<position)
	}
}