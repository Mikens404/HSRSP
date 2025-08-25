package domain

import "testing"

func TestCalculateFare(t *testing.T) {
	type args struct {
		distance int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "適切な料金が返却される,500円",
			args: args{
				distance: 10,
			},
			want: 500,
		},
		{
			name: "適切な料金が返却される,600円",
			args: args{
				distance: 20,
			},
			want: 600,
		},
		{
			name: "適切な料金が返却される,500円",
			args: args{
				distance: 30,
			},
			want: 700,
		},
		{
			name: "マイナス値が渡されたら0円が返ってくる",
			args: args{
				distance: -10,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateFare(tt.args.distance); got != tt.want {
				t.Errorf("CalculateFare() = %v, want %v", got, tt.want)
			}
		})
	}
}
