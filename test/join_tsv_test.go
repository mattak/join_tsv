package join_tsv_test

import (
	"github.com/mattak/join_tsv/pkg/join_tsv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type JoinTsvTestContext struct{}

func (JoinTsvTestContext) setup() {
	os.RemoveAll("/tmp/test")
}

func (JoinTsvTestContext) tearDown() {
	os.RemoveAll("/tmp/test")
}

func TestParseInt(t *testing.T) {
	context := JoinTsvTestContext{}
	context.setup()
	defer context.tearDown()

	result1 := join_tsv.JoinTablesByFile(
		[]int{0, 0},
		[]string{
			"./data/table1.tsv",
			"./data/table2.tsv",
		})
	assert.Equal(t, [][]string{
		{"1", "a", "a1", "a", "A1"},
		{"2", "b", "b1", "b", "B1"},
		{"3", "c", "c1", "c", "C1"},
	}, result1)

	result2 := join_tsv.JoinTablesByFile(
		[]int{0, 0},
		[]string{
			"./data/table1.tsv",
			"./data/table3.tsv",
		})
	assert.Equal(t, [][]string{
		{"1", "a", "a1", "α"},
		{"2", "b", "b1", ""},
		{"3", "c", "c1", "γ"},
	}, result2)
}
