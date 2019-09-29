package main

type Bag struct {
	items []int
}

func Insert(b *Bag, itemId int)  {
	b.items = append(b.items, itemId)
}

func (b *Bag) Insert(itemid int)  {
	b.items = append(b.items, itemid)
}

func main() {
	bag := new(Bag)
	//Insert(bag, 1001)

	bag.Insert(1001)
}