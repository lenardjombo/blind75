package arraysandhashing

func groupAnagrams(strs []string) [][]string {
    res := make(map[string][]string)

    for _, s := range strs {
        sortedS := sortString(s)
        res[sortedS] = append(res[sortedS], s)
    }

    var result [][]string
    for _, group := range res {
        result = append(result, group)
    }
    return result
}

func sortString(s string) string {
	characters := []rune(s)
	sort.Slice(characters, func(i, j int) bool {
		return characters[i] < characters[j]
	})
	return string(characters)
}
