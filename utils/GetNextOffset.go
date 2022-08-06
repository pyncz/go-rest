package utils

func GetNextOffset(offset int64, total int64, length int64) int64 {
	// Send the next offset if there are still records left
	var nextOffset int64 = 0
	recordsIterated := offset + length
	if total > recordsIterated {
		nextOffset = recordsIterated + 1
	}
	return nextOffset
}
