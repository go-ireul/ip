package ip

import (
	"net"
	"testing"
)

func TestIsReservedIP(t *testing.T) {
	type args struct {
		ip net.IP
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "127.0.0.1",
			args: args{
				ip: net.ParseIP("127.0.0.1"),
			},
			want: true,
		},
		{
			name: "126.0.0.1",
			args: args{
				ip: net.ParseIP("126.0.0.1"),
			},
			want: false,
		},
		{
			name: "106.15.138.139",
			args: args{
				ip: net.ParseIP("106.15.138.139"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsReservedIP(tt.args.ip); got != tt.want {
				t.Errorf("IsReservedIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
