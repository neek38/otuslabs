package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

type KeyValue struct {
	Key   string
	Value int
}

func Top10(inputtedString string) []string {
	regexToSplit := regexp.MustCompile(`[0-9\n\t.,!?;: «»()—\"']+`)
	onlySeparators := regexp.MustCompile(`^[0-9\n\t.,!?;:\- «»()—\"']+$`)
	wordsMap := make(map[string]int)
	var result []string

	//var keyValueSlice []KeyValue
	keyValueSlice := []KeyValue{}
	if inputtedString == "" {
		return result
	}

	split := regexToSplit.Split(inputtedString, -1)

	for i := range split {
		if !onlySeparators.MatchString(split[i]) {
			loweredString := strings.ToLower(split[i])
			_, ok := wordsMap[loweredString]
			if ok {
				wordsMap[loweredString]++
			} else {
				wordsMap[loweredString] = 1
			}
		}
	}

	for k, v := range wordsMap {
		keyValueSlice = append(keyValueSlice, KeyValue{k, v})
	}

	sort.Slice(keyValueSlice, func(i, j int) bool {
		return keyValueSlice[i].Value > keyValueSlice[j].Value
	})

	for i := 0; i < 10; i++ {
		result = append(result, keyValueSlice[i].Key)
	}

	return result
}
