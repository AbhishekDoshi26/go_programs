package hello

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello World!"
	got := Say([]string{"World!"})
	if want != got {
		t.Errorf("Wanted %s but got %s", want, got)
	}

	subtests := []struct {
		items  []string
		result string
	}{{
		result: "Hello world",
	}, {
		items:  []string{"World!"},
		result: "Hello World!",
	}, {
		items:  []string{"World!", "World!"},
		result: "Hello World!, World!",
	},
	}

	for _, test := range subtests {
		if result := Say(test.items); test.result != result {
			t.Errorf("Passed %v wanted %s but got %s", test.items, test.result, result)
		}
	}
}
