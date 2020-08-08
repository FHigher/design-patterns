package facade

import "fmt"

// FeedCat cat
type FeedCat struct{}

func (fc *FeedCat) play() {
	fmt.Println("feed cat play")
}

// FeedDog dog
type FeedDog struct{}

func (fd *FeedDog) play() {
	fmt.Println("feed dog play")
}

// FeedAnimal animal
type FeedAnimal struct{
	*FeedCat
	*FeedDog
}

// NewFeedAnimal animal
func NewFeedAnimal() *FeedAnimal {
	return &FeedAnimal{
		&FeedCat{},
		&FeedDog{},
	}
}

func (fa *FeedAnimal) play() {
	fa.FeedCat.play()
	fa.FeedDog.play()
}