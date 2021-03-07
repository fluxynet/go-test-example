package gote

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	rc := m.Run()
	if rc != 0 {
		os.Exit(rc)
	}

	cov := testing.Coverage()
	j := bytes.NewBufferString(fmt.Sprintf(`{"percentage": "%f"}`, cov))

	cli := &http.Client{
		Timeout: time.Second * 3,
	}

	var url string

	if v, ok := os.LookupEnv("webhook_url"); ok {
		url = v
	}

	ioutil.WriteFile("foo.txt", []byte("["+url+"]"), 0777)

	cli.Post("http://127.0.0.1:8080", "text/text", j)
	os.Exit(0)
}

func TestSum(t *testing.T) {
	type args struct {
		nums      []int
		positives bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty",
			args: args{
				nums: nil,
			},
			want: 0,
		},
		{
			name: "1 Positive",
			args: args{
				nums: []int{10},
			},
			want: 10,
		},
		{
			name: "2 Positive",
			args: args{
				nums: []int{1, 2},
			},
			want: 3,
		},
		{
			name: "1 Negative",
			args: args{
				nums: []int{-1},
			},
			want: -1,
		},
		{
			name: "2 Negative",
			args: args{
				nums: []int{-1, -2},
			},
			want: -3,
		},
		{
			name: "4 Negative / Positive",
			args: args{
				nums: []int{100, -100, 9, -1},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.positives, tt.args.nums...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
