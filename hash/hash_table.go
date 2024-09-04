package hash

// HashTable represents a hash table with chaining for collision resolution
type HashTable struct {
	table map[int]*Bucket
	size  int
}

// Bucket represents a linked list in each slot of the hash table
type Bucket struct {
	head *Node
}

// Node represents a single element in the linked list
type Node struct {
	key   string
	value int
	next  *Node
}

// NewHashTable creates a new hash table of a given size
func NewHashTable(size int) *HashTable {
	return &HashTable{
		table: make(map[int]*Bucket, size),
		size:  size,
	}
}

// HashFunction computes the hash value for a given key by division
func (ht *HashTable) HashFunction(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % ht.size
}

// Insert adds a key-value pair to the hash table
func (ht *HashTable) Insert(key string, value int) {
	index := ht.HashFunction(key)
	if ht.table[index] == nil {
		ht.table[index] = &Bucket{}
	}
	ht.table[index].Insert(key, value)
}

// Search finds the value associated with a given key
func (ht *HashTable) Search(key string) (int, bool) {
	index := ht.HashFunction(key)
	if ht.table[index] != nil {
		return ht.table[index].Search(key)
	}
	return 0, false
}

// Delete removes a key-value pair from the hash table
func (ht *HashTable) Delete(key string) {
	index := ht.HashFunction(key)
	if ht.table[index] != nil {
		ht.table[index].Delete(key)
	}
}

// Insert adds a node to the bucket's linked list
func (b *Bucket) Insert(key string, value int) {
	newNode := &Node{
		key:   key,
		value: value,
		next:  b.head,
	}
	b.head = newNode
}

// Search finds the value associated with a given key in the bucket's linked list
func (b *Bucket) Search(key string) (int, bool) {
	current := b.head
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}

// Delete removes a node from the bucket's linked list
func (b *Bucket) Delete(key string) {
	if b.head == nil {
		return
	}

	if b.head.key == key {
		b.head = b.head.next
		return
	}

	prev := b.head
	for prev.next != nil {
		if prev.next.key == key {
			prev.next = prev.next.next
			return
		}
		prev = prev.next
	}
}
