package pages

import (
	"encoding/json"
	"os"
	"testing"
)

type pageCountCase struct {
	Description string `json:"description"`
	Items       int    `json:"items"`
	PerPage     int    `json:"perPage"`
	Expected    int    `json:"expected"`
}

func TestPageCount(t *testing.T) {
	raw, err := os.ReadFile("tests/unit_test/page_count.json")
	if err != nil {
		t.Fatal(err)
	}
	var cases []pageCountCase
	if err := json.Unmarshal(raw, &cases); err != nil {
		t.Fatal(err)
	}
	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			if got := PageCount(c.Items, c.PerPage); got != c.Expected {
				t.Errorf("PageCount(%d, %d) = %d, want %d", c.Items, c.PerPage, got, c.Expected)
			}
		})
	}
}
