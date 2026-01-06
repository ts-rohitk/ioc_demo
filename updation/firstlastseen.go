package updation

func FirstSeenUpdation(firstSeen *int64, existingFirstSeen *int64) *int64 {
	if *existingFirstSeen < *firstSeen {
		return existingFirstSeen
	}
	return firstSeen
}

func LastSeenUpdation(lastSeen *int64, existingLastSeen *int64) *int64 {
	if lastSeen != nil && existingLastSeen != nil {
		return existingLastSeen
	}

	if lastSeen == nil && existingLastSeen != nil {
		return existingLastSeen
	}

	if lastSeen != nil && existingLastSeen == nil {
		return lastSeen
	}

	return lastSeen
}
