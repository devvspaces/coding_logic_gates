package logic

import (
	"testing"
)

func TestGates(t *testing.T) {

	t.Run("Testing the AND Gate", func(t *testing.T) {
		truthTable := []map[string][]bool{
			{
				"input":  {true, true},
				"output": {true},
			},
			{
				"input":  {true, false},
				"output": {false},
			},
			{
				"input":  {false, true},
				"output": {false},
			},
			{
				"input":  {false, false},
				"output": {false},
			},
		}

		assertTruthHelper(t, truthTable, AND)
	})

	t.Run("Testing the OR Gate", func(t *testing.T) {
		truthTable := []map[string][]bool{
			{
				"input":  {true, true},
				"output": {true},
			},
			{
				"input":  {true, false},
				"output": {true},
			},
			{
				"input":  {false, true},
				"output": {true},
			},
			{
				"input":  {false, false},
				"output": {false},
			},
		}

		assertTruthHelper(t, truthTable, OR)
	})

	t.Run("Testing the XOR Gate", func(t *testing.T) {
		truthTable := []map[string][]bool{
			{
				"input":  {true, true},
				"output": {false},
			},
			{
				"input":  {true, false},
				"output": {true},
			},
			{
				"input":  {false, true},
				"output": {true},
			},
			{
				"input":  {false, false},
				"output": {false},
			},
		}

		assertTruthHelper(t, truthTable, XOR)
	})

	t.Run("Testing the NAND Gate", func(t *testing.T) {
		truthTable := []map[string][]bool{
			{
				"input":  {true, true},
				"output": {false},
			},
			{
				"input":  {true, false},
				"output": {true},
			},
			{
				"input":  {false, true},
				"output": {true},
			},
			{
				"input":  {false, false},
				"output": {true},
			},
		}

		assertTruthHelper(t, truthTable, NAND)
	})

	t.Run("Testing the XNOR Gate", func(t *testing.T) {
		truthTable := []map[string][]bool{
			{
				"input":  {true, true},
				"output": {true},
			},
			{
				"input":  {true, false},
				"output": {false},
			},
			{
				"input":  {false, true},
				"output": {false},
			},
			{
				"input":  {false, false},
				"output": {true},
			},
		}

		assertTruthHelper(t, truthTable, XNOR)
	})

	t.Run("Testing the NOT Gate", func(t *testing.T) {
		truthTable := []map[string]bool{
			{
				"input":  false,
				"output": true,
			},
			{
				"input":  true,
				"output": false,
			},
		}

		for _, el := range truthTable {
			input := el["input"]
			output := el["output"]
			got := NOT(input)
			if got != output {
				t.Errorf("Expected %t, but got %t", output, got)
			}
		}
	})
}

func assertTruthHelper(t testing.TB, truthTable []map[string][]bool, fn func(a bool, b bool) bool) {

	t.Helper()

	for _, el := range truthTable {
		input := el["input"]
		output := el["output"][0]
		got := fn(input[0], input[1])
		if got != output {
			t.Errorf("Expected %t, but got %t", output, got)
		}
	}

}
