package deck

import (
"crypto/rand"
"errors"
"fmt"
"math/big"
)
type Types struct{ Name, Symbol string }

type Val struct {
Name string
Value int
}

type Card struct {
Types Types
Value Val
}

func (card *Card) GreaterThan(b *Card) bool {
return card.Value.Value > b.Value.Value
}

func (card *Card) LessThan(b *Card) bool {
return card.Value.Value < b.Value.Value
}

func (card *Card) Equal(b *Card) bool {
return card.Value.Value == b.Value.Value
}

func (card *Card) Facecard() (ans bool) {
n := card.Value.Name
return n == "J" || n == "Q" || n == "K" || n == "A"
}

func (card *Card) ToStr() string {
if card.Facecard() {
return fmt.Sprintf(" %c%s", card.Value.Name[0], card.Types.Symbol)
}
return fmt.Sprintf("%2d%s", card.Value.Value, card.Types.Symbol)
}

type Deck struct {
Cards []Card
}

func (d *Deck) Initialize() error {

types := []Types{
Types{"Club", "♣"},
Types{"Diamond", "♦"},
Types{"Heart", "♥"},
Types{"Spade", "♠"},
}
vals := []Val{
Val{"2", 2},
Val{"3", 3},
Val{"4", 4},
Val{"5", 5},
Val{"6", 6},
Val{"7", 7},
Val{"8", 8},
Val{"9", 9},
Val{"10", 10},
Val{"J", 11},
Val{"Q", 12},
Val{"K", 13},
Val{"A", 14},
}
d.Cards = nil
for _, t := range types {
for _, val := range vals {
d.Cards = append(d.Cards, Card{Types: t, Value: val})
}
}
d.Shuffle()
return nil
}

func (d *Deck) Shuffle() (err error) {
var old []Card
old = d.Cards
var shuffled []Card

for i := len(old); i > 0; i-- {
nBig, e := rand.Int(rand.Reader, big.NewInt(int64(i)))
if e != nil {
panic(e)
}
j := nBig.Int64()
shuffled = append(shuffled, old[j])
old = remove(old, j)
}
d.Cards = shuffled
return nil
}

func remove(slice []Card, i int64) []Card {
copy(slice[i:], slice[i+1:])
return slice[:len(slice)-1]
}

func (d *Deck) Draw(count int) (cards []Card, err error) {
if count > len(d.Cards) {
return nil, errors.New("You can not draw more cards")
}
hand := d.Cards[0:count]
d.Cards = d.Cards[count:]
return hand, nil
}

func (d *Deck) CardsLeft() int {
return len(d.Cards)
}
