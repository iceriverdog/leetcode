package datastructure
// 字符串不能交换，转成byte， byte(s), s(byte)
func reverseVowels(s string) string {
	new := []byte(s)
	i,j := 0,len(new)-1
	for i < j {
		for !aeiou(new[i]) && i <j {
			i++
		}
		for !aeiou(new[j]) && i <j {
			j--
		}
		if i <j {
			new[i], new[j] = new[j],new[i]
			i++
			j--
		}

	}
	return string(new)
}

func aeiou(b byte) bool {
	switch b {
	case 'A', 'E','I', 'O','U','a','e','i','o','u' :
		return true
	default :
		return false
	}
}
