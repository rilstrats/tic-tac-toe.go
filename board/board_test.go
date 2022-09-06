package board

import "testing"

func TestWinCheck(t *testing.T) {
    cases := []struct {
        in [10]string 
        want bool
    } {
        // unmodified list
        {[10]string {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}, false},

        // horizontal wins
        {[10]string {"0", "X", "X", "X", "4", "5", "6", "7", "8", "9"}, true},
        {[10]string {"0", "1", "2", "3", "X", "X", "X", "7", "8", "9"}, true},
        {[10]string {"0", "1", "2", "3", "4", "5", "6", "X", "X", "X"}, true},

        // vertical wins
        {[10]string {"0", "X", "2", "3", "X", "5", "6", "X", "8", "9"}, true},
        {[10]string {"0", "1", "X", "3", "4", "X", "6", "7", "X", "9"}, true},
        {[10]string {"0", "1", "2", "X", "4", "5", "X", "7", "8", "X"}, true},

        // diagonal wins
        {[10]string {"0", "X", "2", "3", "4", "X", "6", "7", "8", "X"}, true},
        {[10]string {"0", "1", "2", "X", "4", "X", "6", "X", "8", "9"}, true},
    }

    for _, cas := range cases {
        Squares = cas.in
        got := WinCheck("X")
        if got != cas.want {
            t.Errorf("Squares is %q, WinCheck returned %v, wanted %v", cas.in, got, cas.want)
        }
    }

}

