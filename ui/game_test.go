package ui

import (
	"testing"

	"github.com/matryer/is"
)

func TestGridPos(t *testing.T) {
	is := is.New(t)
	t.Parallel()
	type testCase struct {
		Idx, Lines, Cols, ExpectedL, ExpectedC int
	}
	testcases := []testCase{
		{Idx: 0, Lines: 0, Cols: 0, ExpectedL: -1, ExpectedC: -1},
		{Idx: 0, Lines: 1, Cols: 1, ExpectedL: 0, ExpectedC: 0},
		{Idx: 0, Lines: 2, Cols: 2, ExpectedL: 0, ExpectedC: 0},
		{Idx: 1, Lines: 2, Cols: 2, ExpectedL: 0, ExpectedC: 1},
		{Idx: 2, Lines: 2, Cols: 2, ExpectedL: 1, ExpectedC: 0},
		{Idx: 3, Lines: 2, Cols: 2, ExpectedL: 1, ExpectedC: 1},
		{Idx: 4, Lines: 2, Cols: 2, ExpectedL: -1, ExpectedC: -1},
	}
	for i, tc := range testcases {
		ActualL, ACtualC := gridPos(tc.Idx, tc.Lines, tc.Cols)
		t.Logf("case #%d: %+v", i, tc)
		is.Equal(ACtualC, tc.ExpectedC) // column should match
		is.Equal(ActualL, tc.ExpectedL) // line should match
	}
}
