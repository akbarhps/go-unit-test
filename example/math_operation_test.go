package example

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSum(b *testing.B) {
	for i := 1; i < b.N; i++ {
		Sum(10000000000, 100000000)
	}
}

func BenchmarkAll(b *testing.B) {
	b.Run("Benchmark Sum", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			Sum(10000000000, 100000000)
		}
	})

	b.Run("Benchmark Multiply", func(b *testing.B) {
		for i := 1; i < b.N; i++ {
			Multiply(10000000000, 100000000)
		}
	})
}

func BenchmarkMultiplyTable(b *testing.B) {
	table := []struct {
		TestName   string
		TestParam1 int
		TestParam2 int
		TestExpect int
	}{
		{
			TestName:   "TestMultiply(10*10)",
			TestParam1: 10,
			TestParam2: 10,
			TestExpect: 100,
		},
		{
			TestName:   "TestMultiply(20*10)",
			TestParam1: 20,
			TestParam2: 10,
			TestExpect: 200,
		},
	}

	for _, test := range table {
		b.Run(test.TestName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result := Multiply(test.TestParam1, test.TestParam2)
				assert.Equal(b, test.TestExpect, result, "result is not the same as expected")
			}
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Before test")

	m.Run() // run tests on this package

	fmt.Println("After test")
}

func TestSum(t *testing.T) {
	expect := 5 + 5
	result := Sum(5, 5)

	assert.Equal(t, expect, result, "result is not the same as expected")
	// if result != expect {
	// t.Fail()
	// t.Fatal("Expecting", expect, ", got:", result)
	// }
}

func TestMultiply(t *testing.T) {
	expect := 5 * 10
	result := Multiply(5, 10)

	require.Equal(t, expect, result, "result is not the same as expected")
	// if result != expect {
	// t.FailNow()
	// t.Fatal("Expecting", expect, ", got:", result)
	// }
}

func TestMultiplySkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("this test can't run on windows")
	}

	expect := 0 * 0
	result := Multiply(0, 0)
	assert.Equal(t, expect, result, "result is not the same as expected")
}

func TestAll(t *testing.T) {
	t.Run("Test Sum Function", func(t *testing.T) {
		result := Sum(10, 10)
		assert.Equal(t, 20, result, "result is not the same as expected")
	})
	t.Run("Test Multiply Function", func(t *testing.T) {
		result := Multiply(10, 10)
		assert.Equal(t, 100, result, "result is not the same as expected")
	})
}

func TestTableSum(t *testing.T) {
	table := []struct {
		TestName   string
		TestParam1 int
		TestParam2 int
		TestExpect int
	}{
		{
			TestName:   "TestSum(10+10)",
			TestParam1: 10,
			TestParam2: 10,
			TestExpect: 20,
		},
		{
			TestName:   "TestSum(20+10)",
			TestParam1: 20,
			TestParam2: 10,
			TestExpect: 30,
		},
	}

	for _, test := range table {
		t.Run(test.TestName, func(t *testing.T) {
			result := Sum(test.TestParam1, test.TestParam2)
			assert.Equal(t, test.TestExpect, result, "result is not the same as expected")
		})
	}
}
