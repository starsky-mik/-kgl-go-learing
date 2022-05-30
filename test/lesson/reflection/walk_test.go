package reflection

import (
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/reflection"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name   string
		Obj    interface{}
		Result []string
	}{
		{
			Name: "Plain struct",
			Obj: struct {
				Name string
			}{"Mike"},
			Result: []string{"Mike"},
		}, {
			Name: "Plain struct with int field",
			Obj: struct {
				Name string
				Age  int
			}{"Mike", 30},
			Result: []string{"Mike"},
		}, {
			Name: "Struct with nested field",
			Obj: struct {
				Name    string
				Profile Profile
			}{"Mike", Profile{
				Age:  30,
				City: "Irkutsk",
			}},
			Result: []string{"Mike", "Irkutsk"},
		}, {
			Name: "Pointer",
			Obj: &struct {
				Name    string
				Profile Profile
			}{"Mike", Profile{
				Age:  30,
				City: "Irkutsk",
			}},
			Result: []string{"Mike", "Irkutsk"},
		}, {
			Name: "Slice",
			Obj: []Profile{
				{0, "Irkutsk"},
				{28, "Kaliningrad"},
			},
			Result: []string{"Irkutsk", "Kaliningrad"},
		}, {
			Name: "Array",
			Obj: [2]Profile{
				{0, "Irkutsk"},
				{28, "Kaliningrad"},
			},
			Result: []string{"Irkutsk", "Kaliningrad"},
		}, {
			Name: "Map",
			Obj: map[string]Profile{
				"first":  {City: "Irkutsk"},
				"second": {Age: 28, City: "Kaliningrad"},
			},
			Result: []string{"Irkutsk", "Kaliningrad"},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.Name, func(t *testing.T) {
			var got []string
			Walk(testCase.Obj, func(input string) {
				got = append(got, input)
			})

			AssertDeepEquals(t, got, testCase.Result)
		})
	}

	t.Run("Channel", func(t *testing.T) {
		ch := make(chan Profile)

		go func() {
			ch <- Profile{City: "Irkutsk"}
			ch <- Profile{Age: 28, City: "Kalinigrad"}
			close(ch)
		}()

		var got []string
		var want = []string{"Irkutsk", "Kalinigrad"}

		Walk(ch, func(input string) {
			got = append(got, input)
		})

		AssertDeepEquals(t, got, want)
	})

	t.Run("Function", func(t *testing.T) {
		fn := func() (Profile, Profile) {
			return Profile{City: "Irkutsk"}, Profile{Age: 28, City: "Kalinigrad"}
		}

		var got []string
		var want = []string{"Irkutsk", "Kalinigrad"}

		Walk(fn, func(input string) {
			got = append(got, input)
		})

		AssertDeepEquals(t, got, want)
	})
}
