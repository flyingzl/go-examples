package condition_test

import (
	"runtime"
	"testing"
)

func TestSwitchCondition(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("Mac OS")
	default:
		t.Log("other OS")
	}

	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("Unknown")
		}
	}

}
