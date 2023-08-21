package lep

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBetween(t *testing.T) {
	type testParseEquals struct {
		left   interface{}
		start  interface{}
		end    interface{}
		result string
		err    error
	}
	var tests = []testParseEquals{
		{
			left:   Param("a"),
			start:  Integer(1),
			end:    Integer(2),
			result: `a between 1 and 2`,
		},
	}

	for _, tt := range tests {
		e, err := parseBetween(tt.left, tt.start, tt.end)
		if tt.err == nil && assert.NoError(t, err) {
			assert.IsType(t, (*Between)(nil), e)
			assert.Equal(t, tt.left, e.Param)
			assert.Equal(t, tt.start, e.Start)
			assert.Equal(t, tt.end, e.End)
			assert.Equal(t, tt.result, e.String())
		} else {
			assert.EqualError(t, err, tt.err.Error())
		}
	}
}
