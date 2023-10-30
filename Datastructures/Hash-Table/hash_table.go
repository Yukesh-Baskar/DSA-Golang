package hashtable

const size = 10

type HashTable struct {
	Array [size]*Bucket
}

type Bucket struct {
	Head *BucketNode
}

type BucketNode struct {
	Key  string
	Next *BucketNode
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum
}

func (ht *HashTable) Insert(key string) {
	index := hash(key)
	ht.Array[index].Insert(key)
}

func (b *Bucket) Insert(key string) {
	newNode := &BucketNode{Key: key}
	newNode.Next = b.Head
	b.Head = newNode
}

func NewHashTable() *HashTable {
	res := &HashTable{}
	// for i := range res.Array {
	// 	res.Array[i] = &Bucket{}
	// }
	return res
}
