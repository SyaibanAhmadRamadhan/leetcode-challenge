package problem

func RomanToInt(s string) int {
	romanVal := [128]int16{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	// total := int(0)
	// for i := 0; i < len(s); i++ {
	// 	value := romanVal[s[i]]
	// 	next := value
	// 	if len(s) > i+1 {
	// 		next = romanVal[s[i+1]]
	// 	}

	// 	if value < next {
	// 		total += (int(next) - int(value))
	// 		i++
	// 	} else {
	// 		total += int(value)
	// 	}
	// }

	total, prevValue := int(0), int16(0)
	for i := len(s) - 1; i >= 0; i-- {
		value := romanVal[s[i]]
		if value < prevValue {
			total -= int(value)
		} else {
			total += int(value)
			prevValue = value
		}
	}

	// total := int(0)
	// for i := 0; i < len(s); i++ {
	// 	if len(s) > i+1 {
	// 		pair := romanVal[s[i]]<<8 | romanVal[s[i+1]]
	// 		switch pair {
	// 		case romanVal['I']<<8 | romanVal['V']:
	// 			total += 4
	// 		case romanVal['I']<<8 | romanVal['X']:
	// 			total += 9
	// 		case romanVal['X']<<8 | romanVal['L']:
	// 			total += 40
	// 		case romanVal['X']<<8 | romanVal['C']:
	// 			total += 90
	// 		case romanVal['C']<<8 | romanVal['D']:
	// 			total += 400
	// 		case romanVal['C']<<8 | romanVal['M']:
	// 			total += 900
	// 		default:
	// 			total += int(romanVal[s[i]])
	// 			continue
	// 		}
	// 		i++
	// 	} else {
	// 		total += int(romanVal[s[i]])
	// 	}
	// }

	return int(total)
}
