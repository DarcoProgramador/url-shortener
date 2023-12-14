package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateShortUrl(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "should return a short url 1",
			args: "https://www.google.com",
			want: "gmYncnVz",
		},
		{
			name: "should return a short url 2",
			args: "https://www.youtube.com/",
			want: "ANZkYUDg",
		},
		{
			name: "should return a empty string",
			args: "",
			want: "",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := GenerateShortUrl(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
