package task

import "testing"

func TestTaskWorkExample(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TaskWorkExample(tt.args.SKUs); got != tt.want {
				t.Errorf("TaskWorkExample() = %v, want %v", got, tt.want)
			}
		})
	}
}
