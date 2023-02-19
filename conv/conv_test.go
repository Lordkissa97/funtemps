package conv

import "reflect"
import "testing"


/*
*
	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/
func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		Fahernheit float64
		Celsius  float64
	}

	tests := []test{
		{Fahernheit: 134, Celsius: 56.67},
	}

	for _, v := range tests {
		got := FarhenheitToCelsius(v.Fahernheit)
		if !reflect.DeepEqual(got, v.Celsius) {
			t.Errorf("FarhenheitToCelsius(%v) = %v, want %v", v.Fahernheit, got, v.Celsius)
		}
	}

	}
}

// De andre testfunksjonene implementeres her
// ...
}
