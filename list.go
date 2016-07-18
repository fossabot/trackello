// Copyright © 2016 Nick Klauer <klauer@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trackello

import (
	"fmt"
	"sort"

	"github.com/VojtechVitek/go-trello"
	"github.com/klauern/trackello/rest"
	"github.com/pkg/errors"
)

// List is both the Trello List + other stats on the actions in it.
type List struct {
	name  string
	cards map[cardID]Card
	stats *statistics
	list  *trello.List
}

// Print will print out a list and all of the cards to the command-line.
func (l *List) Print() {
	if len(l.name) > 0 {
		fmt.Printf("%s\n", l.name)
		cardSlice := make([]Card, 0, len(l.cards))
		for _, card := range l.cards {
			cardSlice = append(cardSlice, card)
		}
		sort.Sort(ByStatisticsCountRev(cardSlice))
		for _, card := range cardSlice {
			fmt.Printf("  * %s\n", card.String())
		}
	}
}

// NewList constructs a new *List based off of a go-trello *List.
func NewList(l *trello.List) *List {
	return &List{
		name:  l.Name,
		cards: make(map[cardID]Card),
		stats: &statistics{},
		list:  l,
	}
}

// MapActions will map all of the Actions that occurred on a List.
func (l *List) MapActions() (bool, error) {
	args := rest.CreateArgsForBoardActions()
	actions, err := l.list.Actions(args...)
	if err != nil {
		fmt.Println("error in MapActions")
		return false, errors.Wrapf(err, "Error getting List \"%s\" Actions: ", l.name)
	}
	for _, action := range actions {
		card, ok := l.cards[cardID(action.Data.Card.Id)]
		if !ok {
			fmt.Printf("Card info is %+v\n", card)
			switch action.Type {
			case "updateList", "createList":
				continue
			}

			fmt.Printf("Not Ok for id %s, action information is %+v\n", action.Data.Card.Id, action)
			if len(action.Data.Card.Id) == 0 {
				continue
			}
		}
		card, ok = l.cards[cardID(action.Data.Card.Id)]
		// maybe an action doesn't exist on this card anymore, but it did at some time?  Need to RE-investigate what
		// card this action belongs to in order to determine whether we're dealing with missing cards in the map
		if !ok {
			panic("Still not ok")
		}
		card.AddCalculation(action)
		l.cards[cardID(action.Data.Card.Id)] = card
	}
	return true, nil
}

// MapCards maps all of the cards for a list into the List.cards map[string]Card based on the cardID.
func (l *List) MapCards() error {
	cards, err := l.list.Cards()
	if err != nil {
		fmt.Printf("Error MapCards %s\n", err)
		return err
	}
	for _, card := range cards {
		l.cards[cardID(card.Id)] = NewCard(&card)
	}
	return nil
}

func makeList(listMap map[string]List) []List {
	list := make([]List, len(listMap))
	for _, v := range listMap {
		list = append(list, v)
	}
	return list
}

// ByListName is a super type of a List[], with functions that implement the sort interface.
type ByListName []List

// Len returns the length of the ByListName slice.
func (l ByListName) Len() int {
	return len(l)
}

// Swap will swap the positions of two trackello.List items in the ByListName slice.
func (l ByListName) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Less determines which of the two trackello.List items is before other based on the List Name.
func (l ByListName) Less(i, j int) bool {
	return l[i].name < l[j].name
}