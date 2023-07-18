package greedy

import "sort"

type Meeting struct {
	Num        int
	Start, End int
}

// ArrangeMeeting 安排宣讲会
func ArrangeMeeting(meetings []*Meeting, start int) []*Meeting {
	if meetings == nil {
		return []*Meeting{}
	}
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].End < meetings[j].End
	})
	result := make([]*Meeting, 0, len(meetings))
	for _, v := range meetings {
		if v.Start >= start {
			result = append(result, v)
			start = v.End
		}
	}
	return result
}
