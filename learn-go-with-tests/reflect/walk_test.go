package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Profile struct {
	Age int
	City string
}

type Person struct {
	Name string
	Profile Profile
}


func TestWalk(t *testing.T) {
	t.Run("not maps", func(t *testing.T) {
		cases := []struct {
			Name string
			Input interface{}
			ExpectedCalls []string
		}{
			{
				"Struce with one string field",
				struct {
					Name string
				}{"Chris"},
				[]string{"Chris"},
			},
			{
				"Struct with two string fields",
				struct {
					Name string
					City string
				}{
					"Chris",
					"London",
				},
				[]string{"Chris", "London"},
			},
			{
				"Struct with non string field",
				struct {
					Name string
					Age int
				}{
					"Chris",
					3,
				},
				[]string{"Chris"},
			},
			{
				"Nested field",
				Person{
					Name:    "Chris",
					Profile: Profile{
						33,
						"London",
					},
				},
				[]string{"Chris", "London"},
			},
			{
				"Pointers to things",
				&Person{
					Name:    "Chris",
					Profile: Profile{
						33,
						"London",
					},
				},
				[]string{"Chris", "London"},
			},
			{
				"Slices",
				[]Profile{
					{33, "London"},
					{22, "China"},
				},
				[]string{"London", "China"},
			},
			{
				"Arrays",
				[2]Profile{
					{33, "London"},
					{22, "Japan"},
				},
				[]string{"London", "Japan"},
			},
			{
				"Maps",
				map[string]string{
					"Foo":"Bar",
					"Baz":"Baz",
				},
				[]string{"Bar", "Baz"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T){
				var got []string
				Walk(test.Input, func(input string){
					got = append(got, input)
				})
				fmt.Printf("got %v, want %v", got, test.ExpectedCalls)
				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, test.ExpectedCalls)
				}
			})
		}
	})

	//t.Run("with maps", func(t *testing.T) {
	//	aMap := map[string]string{
	//		"Foo":"Bar",
	//		"Baz": "Boz",
	//	}
	//
	//	var got []string
	//	Walk(aMap, func(input string) {
	//		got = append(got, input)
	//	})
	//
	//	assertContains(t, got, "Bar")
	//	assertContains(t, got, "Boz")
	//})
}

func assertContains(t *testing.T, haystack []string, needle string)  {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain '%s' but it didnt", haystack, needle)
	}
}
