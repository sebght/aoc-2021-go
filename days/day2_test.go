package days

import (
	"testing"
)

func TestFirstPartDay2OneForward(t *testing.T) {
	result := FirstPartDay2([]string{"forward 1"})
	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestFirstPartDay2OneForwardOneDown(t *testing.T) {
	result := FirstPartDay2([]string{"forward 1", "down 3"})
	if result != 3 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 3)
	}
}

func TestFirstPartDay2OneForwardOneUp(t *testing.T) {
	result := FirstPartDay2([]string{"forward 1", "up 3"})
	if result != -3 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, -3)
	}
}

func TestFirstPartDay2OneForwardOneIncorrect(t *testing.T) {
	result := FirstPartDay2([]string{"forward 1", "toto 3"})
	if result != 0 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 0)
	}
}

func TestFirstPartDay2OfficalExample(t *testing.T) {
	result := FirstPartDay2([]string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"})
	if result != 150 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 150)
	}
}
