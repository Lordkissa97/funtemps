package conv

func TerstFahrenheitToCelcius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}
	{
		tests := []tests{
			{134, 56.67},
		}

		for _, tc := range tests {
			got := FahrenheitToCelcius(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("expected: (%v), got:  %v", tc.want, got)
			}
		}
	}

}
