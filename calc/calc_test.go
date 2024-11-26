package calc

import (
	"fmt"
	"testing"
)

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

func TestSubtraction_Calculate(t1 *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "5-2=3", args: args{a: 5, b: 2}, want: 3},
		{name: "1-1=0", args: args{a: 1, b: 1}, want: 0},
		{name: "0-0=0", args: args{a: 0, b: 0}, want: 0},
		{name: "0--2=2", args: args{a: 0, b: -2}, want: 2},
		{name: "-2--2=0", args: args{a: -2, b: -2}, want: 0},
		{name: "-2-0=-2", args: args{a: -2, b: 0}, want: -2},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Subtraction{}
			if got := t.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t1.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplication_Calculate(t1 *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "2*2=4", args: args{a: 2, b: 2}, want: 4},
		{name: "0*1=0", args: args{a: 0, b: 1}, want: 0},
		{name: "1*5=5", args: args{a: 1, b: 5}, want: 5},
		{name: "0*0=0", args: args{a: 0, b: 0}, want: 0},
		{name: "-2*-2=-4", args: args{a: -2, b: -2}, want: 4},
		{name: "-2*2=-4", args: args{a: -2, b: 2}, want: -4},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Multiplication{}
			if got := t.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t1.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision_Calculate(t1 *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "2/2=1", args: args{a: 2, b: 2}, want: 1},
		{name: "0/1=0", args: args{a: 0, b: 1}, want: 0},
		{name: "5/2=2", args: args{a: 5, b: 2}, want: 2},
		{name: "5/-2=-2", args: args{a: 5, b: -2}, want: -2},
		{name: "-5/2=-2", args: args{a: -5, b: 2}, want: -2},
		{name: "-4+-2=2", args: args{a: -4, b: -2}, want: 2},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Division{}
			if got := t.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t1.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivideByZero_Calculate(t1 *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			errorMessage := fmt.Sprintf("%v", err)
			panicDivideByZeroMessage := "runtime error: integer divide by zero"
			assertEquals(t1, errorMessage, panicDivideByZeroMessage)
		}
	}()
	calculator := &Division{}
	calculator.Calculate(5, 0)

}

func assertEquals(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("actual = %s, expected = %s", actual, expected)
	}
}
