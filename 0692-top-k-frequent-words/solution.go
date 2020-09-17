func topKFrequent(words []string, k int) []string {
	count := make(map[string]int, len(words))
	for _, w := range words {
		count[w]++
	}
	fw := make(frewords, 0, len(count))
	for w, c := range count {
		fw = append(fw, &entry{
			word: w,
			freq: c,
		})
	}
	sort.Sort(fw)

	res := make([]string, k)
	for i := 0; i < k; i++ {
		res[i] = fw[i].word
	}
	return res
}

type entry struct {
	word string
	freq int
}

type frewords []*entry

func (fre frewords) Len() int { return len(fre) }

func (fre frewords) Less(i, j int) bool {
	if fre[i].freq == fre[j].freq {
		return fre[i].word < fre[j].word
	}
	return fre[i].freq > fre[j].freq
}

func (fre frewords) Swap(i, j int) {
	fre[i], fre[j] = fre[j], fre[i]
}