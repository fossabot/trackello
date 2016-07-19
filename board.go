package trackello

import (
	"sync"

	"sort"

	"github.com/VojtechVitek/go-trello"
	"github.com/pkg/errors"
)

// Board is a super-type for a Trello board.  Board also contains a mutex and map of a List ID to a List.
type Board struct {
	id      string
	board   *trello.Board
	listMux *sync.RWMutex
	lists   map[string]List
}

// NewBoard will create a new Board type, using a trello.Board as a starting point.
func NewBoard(b *trello.Board) *Board {
	return &Board{
		id:      b.Id,
		board:   b,
		listMux: &sync.RWMutex{},
		lists:   make(map[string]List),
	}
}

// PopulateLists will Populate the board's lists with cards and missing data.
func (b *Board) PopulateLists() error {
	lists, err := b.board.Lists()
	if err != nil {
		return errors.Wrapf(err, "Unable to get Lists for Board %s", b.board.Name)
	}
	wg := sync.WaitGroup{}
	for _, list := range lists {
		list := list
		wg.Add(1)
		go func(list trello.List) {
			defer wg.Done()
			// 1. calculate the actions for the list
			trackList := NewList(&list)
			if err = trackList.MapCards(); err != nil {
				return
			}
			// 2. return the list actions to return to the board
			b.listMux.Lock()
			b.lists[trackList.list.Id] = *trackList
			b.listMux.Unlock()
		}(list)
	}
	wg.Wait()

	return nil
}

// MapActions queries Trello's API for all of the recent actions performed on a Board, and maps that to the
// board itself, into a list and card.
func (b *Board) MapActions() error {
	wg := sync.WaitGroup{}
	for _, list := range b.lists {
		wg.Add(1)
		go func(l List) {
			defer wg.Done()
			b.listMux.Lock()
			if _, err := l.MapActions(); err != nil {
				panic(err)
			}
			b.listMux.Unlock()
		}(list)
	}
	wg.Wait()
	return nil
}

// PrintActions will print the board actions out
func (b *Board) PrintActions() {
	lists := make([]List, len(b.lists))
	for _, list := range b.lists {
		lists = append(lists, list)
	}
	b.listMux.Lock()
	sort.Sort(ByListName(lists))
	b.listMux.Unlock()
	for _, list := range lists {
		b.listMux.RLock()
		list.PrintNonZeroActions()
		b.listMux.RUnlock()
	}
}
