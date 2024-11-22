package calc

import "testing"

func TestAddition_Calculate(t1 *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1+2=3", args: args{a: 1, b: 2}, want: 3},
		{name: "0+1=1", args: args{a: 0, b: 1}, want: 1},
		{name: "0+0=0", args: args{a: 0, b: 0}, want: 0},
		{name: "0+-2=-2", args: args{a: 0, b: -2}, want: -2},
		{name: "-2+-2=-4", args: args{a: -2, b: -2}, want: -4},
		{name: "4+-2=-4", args: args{a: 4, b: -2}, want: 2},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Addition{}
			if got := t.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t1.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdditionAnotherWay_Calculate(t1 *testing.T) {
	calculator := &Addition{}
	result := calculator.Calculate(1, 2)
	if result != 3 {
		t1.Errorf("Calculate() = %v, want %v", result, 3)
	}
}
