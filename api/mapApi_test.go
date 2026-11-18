package api

import (
	"testing"
)

func TestMap(t *testing.T) {
	mapApi := MapApi{}
	cases := []struct {
		isForward bool
		expected  []MapLocation
	}{
		{
			isForward: true,
			expected: []MapLocation{
				{
					Name: "canalave-city-area",
				},
				{
					Name: "eterna-city-area",
				},
				{
					Name: "pastoria-city-area",
				},
				{
					Name: "sunyshore-city-area",
				},
				{
					Name: "sinnoh-pokemon-league-area",
				},
				{
					Name: "oreburgh-mine-1f",
				},
				{
					Name: "oreburgh-mine-b1f",
				},
				{
					Name: "valley-windworks-area",
				},
				{
					Name: "eterna-forest-area",
				},
				{
					Name: "fuego-ironworks-area",
				},
				{
					Name: "mt-coronet-1f-route-207",
				},
				{
					Name: "mt-coronet-2f",
				},
				{
					Name: "mt-coronet-3f",
				},
				{
					Name: "mt-coronet-exterior-snowfall",
				},
				{
					Name: "mt-coronet-exterior-blizzard",
				},
				{
					Name: "mt-coronet-4f",
				},
				{
					Name: "mt-coronet-4f-small-room",
				},
				{
					Name: "mt-coronet-5f",
				},
				{
					Name: "mt-coronet-6f",
				},
				{
					Name: "mt-coronet-1f-from-exterior",
				},
			},
		},
		{
			isForward: true,
			expected: []MapLocation{
				{
					Name: "mt-coronet-1f-route-216",
				},
				{
					Name: "mt-coronet-1f-route-211",
				},
				{
					Name: "mt-coronet-b1f",
				},
				{
					Name: "great-marsh-area-1",
				},
				{
					Name: "great-marsh-area-2",
				},
				{
					Name: "great-marsh-area-3",
				},
				{
					Name: "great-marsh-area-4",
				},
				{
					Name: "great-marsh-area-5",
				},
				{
					Name: "great-marsh-area-6",
				},
				{
					Name: "solaceon-ruins-2f",
				},
				{
					Name: "solaceon-ruins-1f",
				},
				{
					Name: "solaceon-ruins-b1f-a",
				},
				{
					Name: "solaceon-ruins-b1f-b",
				},
				{
					Name: "solaceon-ruins-b1f-c",
				},
				{
					Name: "solaceon-ruins-b2f-a",
				},
				{
					Name: "solaceon-ruins-b2f-b",
				},
				{
					Name: "solaceon-ruins-b2f-c",
				},
				{
					Name: "solaceon-ruins-b3f-a",
				},
				{
					Name: "solaceon-ruins-b3f-b",
				},
				{
					Name: "solaceon-ruins-b3f-c",
				},
			},
		},
		{
			isForward: false,
			expected: []MapLocation{
				{
					Name: "canalave-city-area",
				},
				{
					Name: "eterna-city-area",
				},
				{
					Name: "pastoria-city-area",
				},
				{
					Name: "sunyshore-city-area",
				},
				{
					Name: "sinnoh-pokemon-league-area",
				},
				{
					Name: "oreburgh-mine-1f",
				},
				{
					Name: "oreburgh-mine-b1f",
				},
				{
					Name: "valley-windworks-area",
				},
				{
					Name: "eterna-forest-area",
				},
				{
					Name: "fuego-ironworks-area",
				},
				{
					Name: "mt-coronet-1f-route-207",
				},
				{
					Name: "mt-coronet-2f",
				},
				{
					Name: "mt-coronet-3f",
				},
				{
					Name: "mt-coronet-exterior-snowfall",
				},
				{
					Name: "mt-coronet-exterior-blizzard",
				},
				{
					Name: "mt-coronet-4f",
				},
				{
					Name: "mt-coronet-4f-small-room",
				},
				{
					Name: "mt-coronet-5f",
				},
				{
					Name: "mt-coronet-6f",
				},
				{
					Name: "mt-coronet-1f-from-exterior",
				},
			},
		},
	}

	for _, c := range cases {
		var actual []MapLocation
		if c.isForward {
			actual = mapApi.GetNext()
		} else {
			actual = mapApi.GetPrevious()
		}

		if len(actual) != len(c.expected) {
			t.Errorf("TestMap: Expected len(): %d but got %d", len(c.expected), len(actual))
		}

		for i := range actual {
			if actual[i].Name != c.expected[i].Name {
				t.Errorf("TestMap: expected location name: '%s' on index: %d \t got: '%s'", c.expected[i].Name, i, actual[i].Name)
			}
		}
	}
}
