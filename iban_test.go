package iban

import (
	"fmt"
	"testing"
)

// Source: https://en.wikipedia.org/wiki/International_Bank_Account_Number
var ibans = []string{
	"FI4250001510000023",                // Finland
	"BE71 0961 2345 6769",               // Belgium
	"FR76 3000 6000 0112 3456 7890 189", // France
	"DE91 1000 0000 0123 4567 89",       // Germany
	"GR96 0810 0010 0000 0123 4567 890", // Greece
	"RO09 BCYP 0000 0012 3456 7890",     // Romania
	"SA44 2000 0001 2345 6789 1234",     // Saudi Arabia
	"ES79 2100 0813 6101 2345 6789",     // Spain
	"CH56 0483 5012 3456 7800 9",        // Switzerland
	"GB98 MIDL 0700 9312 3456 78",       // United Kindom
}

var uglyStrings = []string{
	" fi42 5000 1510 0000 23 ",
	"G   B 9  8    M ID L  0 70 0  93 1 2   34 5 6  7 8",
}

func TestValidate(t *testing.T) {
	for _, address := range ibans {
		ret, msg := Validate(address)
		if !ret {
			t.Errorf("Error with address %s, error: %s", address, msg)
		}
	}

}

func BenchmarkValidate(t *testing.B) {
	for i := 0; i < t.N; i++ {
		for _, address := range ibans {
			ret, msg := Validate(address)
			if !ret {
				t.Errorf("Error with address %s, error: %s", address, msg)
			}
		}
	}

}

func TestUglyAddress(t *testing.T) {
	for _, address := range uglyStrings {
		ret, msg := Validate(address)
		if !ret {
			t.Errorf("Error with address %s, error: %s", address, msg)
		}
	}
}

func TestDebug(t *testing.T) {
	ret, msg := Validate(" fi42 5000 1510 0000 23 ")
	if !ret {
		t.Errorf("Error debug1: %s", msg)

	}
	fmt.Println("-----------")
	ret2, msg2 := Validate("FI4250001510500023")
	if ret2 {
		t.Errorf("Error (this should be invalid): %s", msg2)
	}
}
