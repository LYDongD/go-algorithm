package strings

import (
	"testing"
)

func Test_number_of_segments_in_a_string(t *testing.T) {
	t.Log(CountSegments("Hello, my name is John"))
	t.Log(CountSegments("Hello, name John"))
	t.Log(CountSegments("  Hello, my name is John  "))
}
