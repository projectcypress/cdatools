package importer_test

import (
	"testing"

	"github.com/projectcypress/cdatools/importer"
)

func TestTimestampToSeconds(t *testing.T) {
	tests := []struct {
		// test case name
		name string
		// input to TimestampToSeconds
		n string
		// true indicates the return should have been nil
		expectNil bool
		// expected value of successful timestamp
		expected int64
	}{
		{
			"Empty String",
			"",
			true,
			0,
		},
		{
			"Alphabet Soup",
			"asdfghjk",
			true,
			0,
		},
		{
			"Proper date without hour, minute, or second",
			"20160808",
			false,
			1470614400,
		},
		{
			"Proper date with hour",
			"2016080834",
			false,
			1470736800,
		},
		{
			"Proper date with hour and minute",
			"201608083434",
			false,
			1470738840,
		},
		{
			"Proper date with hour, minute, and second",
			"20160808343434",
			false,
			1470738874,
		},
		{
			"Proper date with letters at seconds position",
			"201608083434as",
			true,
			0,
		},
		{
			"Proper date with letters at hour position",
			"20160808as3434",
			true,
			0,
		},
		{
			"Proper date with letters at minute position",
			"2016080834as34",
			true,
			0,
		},
		{
			"Proper string, but length 9 instead of 8",
			"201608083",
			true,
			0,
		},
		{
			"Proper string, but length 11 instead of 10",
			"20160808343",
			true,
			0,
		},
		{
			"Proper string, but length 13 instead of 12",
			"2016080834343",
			true,
			0,
		},
	}

	for _, tt := range tests {
		actual := importer.TimestampToSeconds(tt.n)
		if actual == nil && !tt.expectNil {
			t.Errorf("test case: %s: actual was nil instead of producing a timestamp", tt.name)
		} else if actual != nil && tt.expectNil {
			t.Errorf("test case: %s: actual was: %d instead of nil\n", tt.name, *actual)
		}

		if actual != nil && (*actual != tt.expected) {
			t.Errorf("test case: %s: actual was: %d when expected was: %d", tt.name, *actual, tt.expected)
		}
	}
}
