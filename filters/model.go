package filters

type PersonScore struct {
	Name  string
	Score float64
}

type ByScore []PersonScore

func (bs ByScore) Len() int           { return len(bs) }
func (bs ByScore) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
func (bs ByScore) Less(i, j int) bool { return bs[i].Score > bs[j].Score }
