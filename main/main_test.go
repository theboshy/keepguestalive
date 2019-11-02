package main

import (
	"fmt"
	"testing"
)
const IP_TEST= "92.168.16.1"
func TestGuestPing(t *testing.T) {
	got, result := keepAlive(IP_TEST)
	if got != nil {
		t.Errorf("Error = %d; want result", got)
	}else {
		fmt.Print(result)
	}
}