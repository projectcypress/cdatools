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
			"Proper date with timezone, and without hour, minute, or second",
			"20160808+1000",
			false,
			1470578400,
		},
		{
			"Proper date with hour",
			"2016080804",
			false,
			1470628800,
		},
		{
			"Proper date with hour and timezone",
			"2016080804+0312",
			false,
			1470617280,
		},
		{
			"Proper date with hour and minute",
			"201608080434",
			false,
			1470630840,
		},
		{
			"Proper date with hour, minute, and second",
			"20160808043434",
			false,
			1470630874,
		},
		{
			"Proper date with hour, minute, second, and positive timezone",
			"20160808043434+0300",
			false,
			1470620074,
		},
		{
			"Proper date with hour, minute, second, and negative timezone",
			"20160808043434-0300",
			false,
			1470641674,
		},
		{
			"Proper date with letters at seconds position",
			"201608080434as",
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
			"2016080804as34",
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
			"20160808043",
			true,
			0,
		},
		{
			"Proper string, but length 13 instead of 12",
			"2016080804343",
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
		if actual != nil {
			//t.Logf("test case: %s: actual was: %d, expected was: %d", tt.name, *actual, tt.expected)
		}
	}
}
