package vmnet

import "testing"

func TestParseSubnet(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    Subnet
		wantErr bool
	}{
		{
			name: "default",
			want: Subnet{
				Gateway: NetGateway,
				DHCPEnd: NetDHCPEnd,
				Netmask: NetMask,
			},
		},
		{
			name:  "custom",
			value: "192.168.107.0/24",
			want: Subnet{
				Gateway: "192.168.107.1",
				DHCPEnd: "192.168.107.254",
				Netmask: "255.255.255.0",
			},
		},
		{
			name:  "small custom subnet",
			value: "10.20.30.0/28",
			want: Subnet{
				Gateway: "10.20.30.1",
				DHCPEnd: "10.20.30.14",
				Netmask: "255.255.255.240",
			},
		},
		{name: "invalid", value: "invalid", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSubnet(tt.value)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseSubnet() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ParseSubnet() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
