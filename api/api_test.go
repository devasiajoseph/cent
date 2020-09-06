package api

import (
	"fmt"
	"testing"
)

func TestAPI(t *testing.T) {
	ans := APITest()
	if ans == "success" {
		fmt.Println("test success")

	} else {
		t.Errorf("Test Failed")
	}

}
