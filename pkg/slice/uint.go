package slice

func DoesExist(list []uint, value uint) bool {
	for _, item := range list {
		if item == value {

			return true
		}
	}

	return false
}

func Uint64toUint(numbers []uint64) []uint {
	uintList := make([]uint, 0)

	for _, num := range numbers {
		uintList = append(uintList, uint(num))
	}

	return uintList
}

func UintToUint64(numbers []uint) []uint64 {
	uint64List := make([]uint64, 0)

	for _, num := range numbers {
		uint64List = append(uint64List, uint64(num))
	}

	return uint64List
}
