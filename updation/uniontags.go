package updation

func UnionTag(incommingTags []string, existingTags []string) []string {
	uniqueFinder := make(map[string]bool)

	if len(incommingTags) <= 0 {
		return existingTags
	}

	if len(existingTags) <= 0 {
		return existingTags
	}

	for _, v := range existingTags {
		if _, ok := uniqueFinder[v]; !ok {
			uniqueFinder[v] = true
		}
	}

	for _, v := range incommingTags {
		if !uniqueFinder[v] {
			uniqueFinder[v] = true
		}
	}

	var newTags []string
	for tag := range uniqueFinder {
		newTags = append(newTags, tag)
	}

	return newTags
}
