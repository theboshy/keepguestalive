package notifyicon

import (
	"fmt"
	"testing"
)

func TestGuestPing(t *testing.T) {
	 got := startNotification("KeepAlive!","Keeping alive guest","puto el que lo lea")
	 if (got != true){
		 t.Errorf("Got an error in execution")
	 }else{
	 	fmt.Print("All ok Arre")
	 }
}