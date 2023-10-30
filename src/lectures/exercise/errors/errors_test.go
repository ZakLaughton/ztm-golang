package timeparse

import (
	"fmt"
	"testing"
)

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:12", true},
		{"1:33:44", true},
		{"bad", false},
		{"1--3:44", false},
		{"0:59:59", true},
		{"", false},
		{"11:22", false},
		{"aa:bb:cc", false},
		{"5:23:", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		fmt.Println("data: ", data)
		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}
}

/***** CHATGPT SOLUTION *****/

// func TestTimeparse(t *testing.T) {
// 	timeComp, err := parseTimeString("14:07:33")
// 	if err != nil {
// 		t.Fatalf("Expected no error, got %v", err)
// 	}
// 	if timeComp.Hour != 14 {
// 		t.Fatalf("Expected hour to be 14, got %d", timeComp.Hour)
// 	}
// 	if timeComp.Minute != 7 {
// 		t.Fatalf("Expected minute to be 7, got %d", timeComp.Minute)
// 	}
// 	if timeComp.Second != 33 {
// 		t.Fatalf("Expected second to be 33, got %d", timeComp.Second)
// 	}

// 	// You can add more test cases for invalid time strings, boundary values, etc.
// }
