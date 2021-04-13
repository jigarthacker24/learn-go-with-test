package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Jigar"},
			ExpectedCalls: []string{"Jigar"},
		},
		{
			Name: "struct with two string value",
			Input: struct {
				Name string
				City string
			}{"Jigar", "Pune"},
			ExpectedCalls: []string{"Jigar", "Pune"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Jigar", 30},
			ExpectedCalls: []string{"Jigar"},
		},
		{
			Name: "Nested fileds",
			Input: Person{
				Name: "Jigar",
				Profile: Profile{
					Age:  32,
					City: "Pune",
				},
			},
			ExpectedCalls: []string{"Jigar", "Pune"},
		},
		{
			Name: "Nested fileds as pointer",
			Input: &Person{
				Name: "Jigar",
				Profile: Profile{
					Age:  32,
					City: "Pune",
				},
			},
			ExpectedCalls: []string{"Jigar", "Pune"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{30, "Pune"},
				{33, "Ahmedabad"},
			},
			ExpectedCalls: []string{"Pune", "Ahmedabad"},
		},
		{
			Name: "Array",
			Input: [2]Profile{
				{30, "Ahmedabad"},
				{33, "Pune"},
			},
			ExpectedCalls: []string{"Ahmedabad", "Pune"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			walk(tt.Input, func(input string) {
				got = append(got, input)
			})

			if len(got) != len(tt.ExpectedCalls) {
				t.Errorf("wrong number of function calls, got %d, want %d", len(got), len(tt.ExpectedCalls))
			}

			if !reflect.DeepEqual(got, tt.ExpectedCalls) {
				t.Errorf("want: %v, got: %v", tt.ExpectedCalls, got)
			}
		})

	}

	t.Run("with maps", func(t *testing.T) {
		in := map[string]string{
			"Vineet": "Ahmedabad",
			"Jigar":  "Pune",
		}

		got := []string{}
		walk(in, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Ahmedabad")
		assertContains(t, got, "Pune")
	})

	t.Run("with chan", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{30, "GJ"}
			aChannel <- Profile{33, "MH"}
			close(aChannel)
		}()

		var got []string
		want := []string{"GJ", "MH"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}

	})

	t.Run("func", func(t *testing.T) {
		aFunc := func() (Profile, Profile) {
			return Profile{30, "UP"}, Profile{33, "Bihar"}
		}

		var got []string
		want := []string{"UP", "Bihar"}

		walk(aFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})

}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false

	for _, val := range haystack {
		if val == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
