package hamming

import "errors"

// приведение типов - https://github.com/FanMakaron/myproject.go/blob/main/docs/types/casting.md
// тип rune - https://github.com/FanMakaron/myproject.go/blob/main/docs/types/rune.md
// ключевое слово range в for - https://github.com/FanMakaron/myproject.go/blob/main/docs/keywords/range.md


func Distance(a, b string) (int, error) {
	missCount := 0
	runesA := []rune(a)
	runesB := []rune(b)

	if len(runesA) != len(runesB) {
		return 0, errors.New("length does not match.")
	}

	for i, _ := range runesA {
		if runesA[i] != runesB[i] {
			missCount++
		}
	}

	return missCount, nil
}
