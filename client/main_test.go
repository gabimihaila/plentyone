package main

import "testing"

func TestSendGETRequest(t *testing.T) {
	type args struct {
		endpoint string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendGETRequest(tt.args.endpoint)
		})
	}
}
