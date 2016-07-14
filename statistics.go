package trackello

import (
	"fmt"

	"github.com/VojtechVitek/go-trello"
	"github.com/fatih/color"
)

// statistics provides a way to show statistical information about a list, card or whatnot by aggregating the updates,
// comments, checklists, and other actions under a specific item.
type statistics struct {
	comments, // represented by a horizontal ellepsis ⋯ 0x22EF
	updates, // represented by a keyboard 0x2328
	creates, // represented by plus +
	checkListItemUpdates int // represented by check mark ✓ 0x2713
}

// Statistics represents the statistics for all the actions generated for a list, card, etc.
type Statistics interface {
	AddCalculation(trello.Action)
}

func (c *Card) AddCalculation(a trello.Action) {
	if c.stats == nil {
		c.stats = &statistics{}
	}
	switch a.Type {
	case "commentCard":
		c.stats.comments++
	case "updateCard":
		c.stats.updates++
	case "createCard", "addChecklistToCard", "addAttachmentToCard":
		c.stats.creates++
	case "updateCheckItemStateOnCard":
		c.stats.checkListItemUpdates++
	default:
		fmt.Printf("Unmapped action type: %s.  Defaulting to update\n", a.Type)
		c.stats.updates++
	}
}

// PrintStatistics will print the statistics information out.
// Example format: [ 3 +  2 ≡  0 ✓  1 … ]
func (s *statistics) PrintStatistics() string {
	if s == nil {
		s = &statistics{}
	}
	stats := "[" + color.CyanString("%d +", s.updates)
	stats = stats + color.RedString(" %d ≡", s.comments)
	stats = stats + color.GreenString(" %d ✓", s.checkListItemUpdates)
	stats = stats + color.MagentaString(" %d …", s.creates)
	stats = stats + "]"
	return stats
}

func (l *List) AddCalculation(a trello.Action) {}
