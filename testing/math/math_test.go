package math

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestAddPositive(t *testing.T) {
	result := Add(1, 2)

	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}

}

func TestAddNegative(t *testing.T) {
	result := Add(-1, -2)

	if result != -3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}

}

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test positive",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "Test 1+4",
			args: args{
				a: 1,
				b: 4,
			},
			want: 5,
		},
		{
			name: "Test 10+10",
			args: args{
				a: 10,
				b: 10,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.args.a, tt.args.b)
			assert.Equal(t, got, tt.want)
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	fmt.Println(b.N, " loops")
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
