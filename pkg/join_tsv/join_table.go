package join_tsv

import (
	"fmt"
	"github.com/mattak/stocks/cmd/basics"
	"strings"
)

func readData(file string) [][]string {
	result := [][]string{}
	lines := basics.ReadLines(file)
	for _, line := range lines {
		values := strings.Split(line, "\t")
		result = append(result, values)
	}
	return result
}

func readDataAsHash(file string, index int) map[string][]string {
	data := readData(file)
	hash := map[string][]string{}

	for _, values := range data {
		key := values[index]
		if index == 0 {
			hash[key] = values[1:]
		} else if index == len(values)-1 {
			hash[key] = values[:(index + 1)]
		} else {
			hash[key] = append(values[:index], values[(index+1):]...)
		}
	}

	return hash
}

func lengthFirstFields(data map[string][]string) int {
	for _, fields := range data {
		return len(fields)
	}
	return 0
}

func PrintTableTsv(data [][]string) {
	for _, fields := range data {
		fmt.Println(strings.Join(fields, "\t"))
	}
}

func fillData(data []string, v string) []string {
	for i := 0; i < len(data); i++ {
		data[i] = v
	}
	return data
}

func JoinTable(data1 [][]string, data2 map[string][]string, index1 int) [][]string {
	results := [][]string{}
	data2Length := lengthFirstFields(data2)

	for i := 0; i < len(data1); i++ {
		fields1 := data1[i]
		key := fields1[index1]

		var r []string
		fields2, ok := data2[key]

		if ok {
			r = append(fields1, fields2...)
		} else {
			r = make([]string, data2Length)
			r = fillData(r, "")
			r = append(fields1, r...)
		}

		results = append(results, r)
	}

	return results
}

func JoinTablesByFile(keys []int, files []string) [][]string {
	if len(files) < 2 {
		panic("args should have more than 2 tsv files")
	}

	result := readData(files[0])

	for i := 1; i < len(keys); i++ {
		maps := readDataAsHash(files[i], keys[i])
		result = JoinTable(result, maps, keys[0])
	}

	return result
}
