package common

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicatesArrMap(t *testing.T) {
	tests := []struct {
		name    string
		results []map[string]interface{}
		want    []map[string]interface{}
	}{
		{
			name: "remove duplicates",
			results: []map[string]interface{}{
				{"url": "https://www.example.com"},
				{"url": "https://www.example.com"},
				{"url": "https://www.example2.com"},
			},
			want: []map[string]interface{}{
				{"url": "https://www.example.com"},
				{"url": "https://www.example2.com"},
			},
		},
		{
			name: "remove duplicates",
			results: []map[string]interface{}{
				{"url": "https://www.example.com"},
				{"url": "https://www.example2.com"},
				{"url": "https://www.example3.com"},
				{"url": "https://www.example3.com"},
			},
			want: []map[string]interface{}{
				{"url": "https://www.example.com"},
				{"url": "https://www.example2.com"},
				{"url": "https://www.example3.com"},
			},
		},
		{
			name: "no duplicates",
			results: []map[string]interface{}{
				{"url": "https://www.example.com"},
				{"url": "https://www.example2.com"},
				{"url": "https://www.example3.com"},
			},
			want: []map[string]interface{}{
				{"url": "https://www.example.com"},
				{"url": "https://www.example2.com"},
				{"url": "https://www.example3.com"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicatesArrMap(tt.results); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicatesArrMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
