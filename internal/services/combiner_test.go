package services

import "testing"

func Test_CombineFor(t *testing.T) {
	testCases := []struct {
		description string
		names       []string
		expected    []string
	}{
		{
			"2 names",
			[]string{
				"corey",
				"bob",
			},
			[]string{
				"corey ",
				"bob corey",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			combiner := Combiner{
				Names: testCase.names,
			}

			actual, _ := combiner.CombineFor()
			for index, expect := range testCase.expected {
				if actual[index] != expect {
					t.Errorf("results did not match expected %s, but got %s", expect, actual[index])
				}
			}
		})
	}
}

func Test_CombineRoutine(t *testing.T) {
	testCases := []struct {
		description string
		names       []string
		expected    []string
	}{
		{
			"2 names",
			[]string{
				"corey",
				"bob",
			},
			[]string{
				"corey ",
				"bob corey",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			combiner := Combiner{
				Names: testCase.names,
			}

			actual, _ := combiner.CombineRoutine()
			for index, expect := range testCase.expected {
				if actual[index] != expect {
					t.Errorf("results did not match expected %s, but got %s", expect, actual[index])
				}
			}
		})
	}
}
