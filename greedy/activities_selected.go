package greedy

import (
	"sort"
)

// Activity represents an activity with a start and end time
type Activity struct {
	start int
	end   int
}

// ByEndTime implements sort.Interface for sorting activities by end time
type ByEndTime []Activity

func (a ByEndTime) Len() int           { return len(a) }
func (a ByEndTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByEndTime) Less(i, j int) bool { return a[i].end < a[j].end }

// ActivitySelection selects the maximum number of non-overlapping activities
func ActivitySelection(activities []Activity) []Activity {
	// Sort activities by end time
	sort.Sort(ByEndTime(activities))

	selected := []Activity{}
	lastEndTime := -1

	for _, activity := range activities {
		if activity.start > lastEndTime {
			selected = append(selected, activity)
			lastEndTime = activity.end
		}
	}

	return selected
}
