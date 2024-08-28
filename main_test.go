package main

import "testing"

func TestAddSum(t *testing.T) {
	var tests = map[string]struct {
		a            int
		b            int
		expectResult int
	}{
		"test simple addition": {
			a:            1,
			b:            1,
			expectResult: 2,
		},
		"test another addition": {
			a:            10,
			b:            10,
			expectResult: 20,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actualResult := AddSum(test.a, test.b)

			if actualResult != test.expectResult {
				t.Errorf("Case %q: wanted %d, got %d", name, test.expectResult, actualResult)
			}
		})
	}
}
