package main

import (
	"testing"
)

func Test_testValidity(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "valid string", args: args{str: "1-hello-2-world"}, want: true},
		{name: "invalid string", args: args{str: "hello world"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testValidity(tt.args.str); got != tt.want {
				t.Errorf("testValidity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_averageNumber(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
		{name: "test avarage", args: args{str: "test av"}, want: 3.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageNumber(tt.args.str); got != tt.want {
				t.Errorf("averageNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wholeStory(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test hello world", args: args{str: "1-hello-2-world"}, want: "hello world"},
		{name: "test error", args: args{str: " "}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := wholeStory(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("wholeStory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("wholeStory() = %v, want %v", got, tt.want)
			}
		})
	}
}
