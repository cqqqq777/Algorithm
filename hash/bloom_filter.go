package hash

type HFunc func(str string) int

type BloomFilter struct {
	bitmap   []byte
	hashFunc []HFunc
}

func NewBloomFilter(m int, hashFunctions ...HFunc) *BloomFilter {
	return &BloomFilter{
		bitmap:   make([]byte, m/8+1),
		hashFunc: hashFunctions,
	}
}

func (b *BloomFilter) Add(str string) {
	bits := b.hash(str)
	for _, v := range bits {
		b.setBit(v)
	}
}

func (b *BloomFilter) Query(str string) bool {
	bits := b.hash(str)
	for _, v := range bits {
		if b.getBit(v) == 0 {
			return false
		}
	}
	return true
}

func (b *BloomFilter) hash(str string) []int {
	result := make([]int, 0, len(b.hashFunc))
	for _, v := range b.hashFunc {
		result = append(result, v(str))
	}
	return result
}

func (b *BloomFilter) getBit(offset int) int {
	blockIndex := offset / 8
	bitIndex := offset % 8
	return int(b.bitmap[blockIndex] >> bitIndex & 1)
}

func (b *BloomFilter) setBit(offset int) {
	blockIndex := offset / 8
	bitIndex := offset % 8
	b.bitmap[blockIndex] = (1 << bitIndex) | b.bitmap[blockIndex]
}
