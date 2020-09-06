package centp

import (
	"fmt"
	"testing"
)

func TestCent(t *testing.T) {
	ans := CentpFunc()
	if ans < 0 {
		t.Errorf("Test Failed")
	} else {
		fmt.Println("test success")
	}

}
