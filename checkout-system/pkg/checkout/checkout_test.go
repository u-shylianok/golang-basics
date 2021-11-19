package checkout

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/model"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/pricing/rules"
)

func TestCheckout_Total(t *testing.T) {
	type args struct {
		SKUs []string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "basic case from task: 01",
			args: args{
				SKUs: []string{"atv", "atv", "atv", "vga"},
			},
			want: 24900,
		},
		{
			name: "basic case from task: 02",
			args: args{
				SKUs: []string{"atv", "ipd", "ipd", "atv", "ipd", "ipd", "ipd"},
			},
			want: 271895,
		},
		{
			name: "basic case from task: 03",
			args: args{
				SKUs: []string{"mbp", "vga", "ipd"},
			},
			want: 194998,
		},
		{
			name: "additional case - (3*vga) + (2*mbp) + (ipd)",
			args: args{
				SKUs: []string{"vga", "vga", "vga", "mbp", "mbp", "ipd"},
			},
			want: 3000 + 139999 + 139999 + 54999,
		},
		{
			name: "additional case - (3*vga) + (2*mbp) + (ipd) changed order",
			args: args{
				SKUs: []string{"ipd", "vga", "mbp", "vga", "vga", "mbp"},
			},
			want: 1*3000 + 2*139999 + 54999,
		},
		{
			name: "additional case - all rules (6*ipd) + (4*atv) + (3*vga) + (2*mbp)",
			args: args{
				SKUs: []string{"ipd", "atv", "vga", "ipd", "mbp", "ipd", "vga", "ipd", "vga", "ipd", "mbp", "atv", "atv", "atv", "ipd"},
			},
			want: 6*49999 + 3*10950 + 1*3000 + 2*139999,
		},
		{
			name: "additional case - (10*atv)",
			args: args{
				SKUs: []string{"atv", "atv", "atv", "atv", "atv", "atv", "atv", "atv", "atv", "atv"},
			},
			want: 10*10950 - 3*10950,
		},
		{
			name: "additional case - (4*ipd)",
			args: args{
				SKUs: []string{"ipd", "ipd", "ipd", "ipd"},
			},
			want: 4 * 54999,
		},
		{
			name: "additional case - (5*ipd)",
			args: args{
				SKUs: []string{"ipd", "ipd", "ipd", "ipd", "ipd"},
			},
			want: 5 * 49999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			catalog := model.GetDefaultCatalog()
			rules := rules.GetDefaultRules()

			co := NewCheckout(catalog, rules)

			for _, sku := range tt.args.SKUs {
				co.Scan(sku)
			}

			got, err := co.Total()
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
