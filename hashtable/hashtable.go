package hashtable

const ArrSize = 7

type HashTable struct {
	array [ArrSize]*bucket
}

func hash(key string) int {
	var hash int
	for _, v := range key {
		hash += int(v)
	}
	return hash % ArrSize
}

func (ht *HashTable) Insert(key string) {
	index := hash(key)
	ht.array[index].insert(key)
}

func (ht *HashTable) Delete(key string) {
	index := hash(key)
	ht.array[index].delete(key)
}

func (ht *HashTable) Search(key string) bool {
	index := hash(key)
	return ht.array[index].search(key)

}

func Init() *HashTable {
	ht := &HashTable{}
	for i := range ht.array {
		ht.array[i] = &bucket{}
	}
	return ht
}

// ht := hashtable.Init()
// list := []string{
// 	"ERIC",
// 	"KENNY",
// 	"KYLE",
// 	"STAN",
// 	"RANDY",
// 	"BUTTERS",
// }
// for _, name := range list {
// 	ht.Insert(name)
// }
// println(ht.Search("BUTTERS"))
