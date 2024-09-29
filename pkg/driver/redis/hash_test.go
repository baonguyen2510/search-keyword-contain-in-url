package redis

import "testing"

func TestKey(t *testing.T) {
	type args struct {
		key string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "full {}",
			args: args{key: "{keyRedis}"},
			want: "keyRedis",
		},
		{
			name: "half {}",
			args: args{key: "{keyRedis"},
			want: "{keyRedis",
		},
		{
			name: "none {}",
			args: args{key: "keyRedis"},
			want: "keyRedis",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Key(tt.args.key); got != tt.want {
				t.Errorf("Key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlot(t *testing.T) {
	type args struct {
		key string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "with key",
			args: args{key: "{keyRedis}"},
			want: 4127,
		},
		{
			name: "with key without curly braces",
			args: args{key: "keyRedis"},
			want: 4127,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slot(tt.args.key); got != tt.want {
				t.Errorf("Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
