package buyandsell_test

import (
	"strconv"
	"testing"

	"github.com/IamNator/gpi_interview/buyandsell"
)

type args struct {
	buy  []int
	sell []int
}

var tests = []struct {
	name string
	args args
	want string
}{
	{
		name: "testing with zero",
		args: args{
			buy:  []int{0, 0, 0},
			sell: []int{1, 1},
		},
		want: "-2",
	},
	{
		name: "testing negative numbers",
		args: args{
			buy:  []int{-1, -3, -8},
			sell: []int{3, 1, 7},
		},
		want: "-23",
	}, {
		name: "testing with small numbers",
		args: args{
			buy:  []int{3, 5, 7},
			sell: []int{1, 1},
		},
		want: "13",
	}, {
		name: "testing with number greater than a thousand but less than ten thousand",
		args: args{
			buy:  []int{4000, 8000, 9000},
			sell: []int{1000, 2000},
		},
		want: "18000",
	},
	{
		name: "testing with numbers greater than ten thousand but less than a hundred thousand",
		args: args{
			buy:  []int{23000, 56000, 91000},
			sell: []int{9000, 45000},
		},
		want: "116000",
	},
}

func TestBuyCement(t *testing.T) {

	for _, tt := range tests {
		var tcement buyandsell.Cement
		t.Run(tt.name, func(t *testing.T) {
			total := 0
			for _, v := range tt.args.buy {
				tcement.BuyCement(v)
				total += v
			}

			if got := tcement.String(); got != strconv.Itoa(total) {
				t.Errorf("BuyCement() = %v, want %v", got, total)
			}

		})
	}
}

func TestSellCement(t *testing.T) {

	for _, tt := range tests {

		var tcement buyandsell.Cement
		t.Run(tt.name, func(t *testing.T) {
			total := 0
			for _, v := range tt.args.sell {
				tcement.SellCement(v)
				total += v
			}

			if got := tcement.String(); got != strconv.Itoa(total) {
				t.Errorf("SellCement() = %v, want %v", got, total)
			}

		})
	}
}

func TestBuyAndSellCement(t *testing.T) {

	for _, tt := range tests {
		var tcement buyandsell.Cement

		t.Run(tt.name, func(t *testing.T) {

			for _, v := range tt.args.buy {
				tcement.BuyCement(v)
			}

			for _, v := range tt.args.buy {
				tcement.SellCement(v)
			}

			if got := tcement.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBuyCement(b *testing.B) {

	for _, tt := range tests {
		var tcement buyandsell.Cement

		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, v := range tt.args.buy {
					tcement.BuyCement(v)
				}
			}
		})
	}
}

func BenchmarkSellCement(b *testing.B) {

	for _, tt := range tests {
		var tcement buyandsell.Cement

		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, v := range tt.args.sell {
					tcement.SellCement(v)

				}
			}
		})
	}
}
