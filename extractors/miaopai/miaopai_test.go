package miaopai

import (
	"testing"

	"github.com/iawia002/annie/config"
	"github.com/iawia002/annie/test"
)

func TestDownload(t *testing.T) {
	config.InfoOnly = true
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "normal test",
			args: test.Args{
				URL:   "https://www.miaopai.com/show/nPWJvdR4z2Bg1Sz3PJpNYffjpDgEiuv4msALgw__.htm",
				Title: "情人节特辑：一个来自绝地求生的爱情故事，送给已经离开的你-绝地求生大逃杀的秒拍",
				Size:  12135847,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Download(tt.args.URL)
			test.CheckError(t, err)
			test.Check(t, tt.args, data[0])
		})
	}
}
