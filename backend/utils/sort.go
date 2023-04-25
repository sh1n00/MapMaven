package utils

import "sort"

type Pair struct {
	Key   string
	Value float64
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func InstructSortByCosin(instructToCosin map[string]float64) PairList {
	pl := make(PairList, len(instructToCosin))
	i := 0
	for k, v := range instructToCosin {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}
