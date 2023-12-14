package url_test

import (
	"os"
	"testing"

	"github.com/DarcoProgramador/url-shortener/internal/url"
	"github.com/DarcoProgramador/url-shortener/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoMock *mocks.Repository

func TestMain(m *testing.M) {
	repoMock = &mocks.Repository{}

	//SaveUrl
	repoMock.On("SaveUrl", mock.Anything, "ErrorUriDB").Return(nil, url.ErrSaveUrl)
	repoMock.On("SaveUrl", mock.Anything, mock.Anything).Return(&url.Url{Id: 1, ShortUrl: "ShortenerSucces", LongUrl: "www.uri.com"}, nil)

	//GetUrl
	repoMock.On("GetUrl", "FakeShortUrl").Return("UrlSucces", nil)
	repoMock.On("GetUrl", mock.Anything).Return("", url.ErrGetUrl)

	code := m.Run()
	os.Exit(code)
}

func Test_serv_SaveUrl(t *testing.T) {
	type args struct {
		req *url.UrlCreateRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *url.UrlCreateResponse
		wantErr error
	}{
		{
			name: "SaveUrlSuccess",
			args: args{
				req: &url.UrlCreateRequest{
					"www.google.com",
				},
			},
			want: &url.UrlCreateResponse{
				Id:       1,
				ShortUrl: "ShortenerSucces",
			},
			wantErr: nil,
		},
		{
			name: "SaveUrlErrorSaveUrl",
			args: args{
				req: &url.UrlCreateRequest{
					"ErrorUriDB",
				},
			},
			want:    nil,
			wantErr: url.ErrSaveUrl,
		},
		{
			name: "SaveUrlErrorEmptyUrl",
			args: args{
				req: &url.UrlCreateRequest{},
			},
			want:    nil,
			wantErr: url.ErrInvalidUrl,
		},
	}

	service := url.NewService(repoMock)

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repoMock.Test(t)
			got, err := service.SaveUrl(tt.args.req)

			if got != nil && tt.want != nil {
				assert.Equal(t, tt.want, got)
			}
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}

func Test_serv_GetUrl(t *testing.T) {
	type args struct {
		shortUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "GetUrlSuccess",
			args: args{
				shortUrl: "FakeShortUrl",
			},
			want:    "UrlSucces",
			wantErr: nil,
		},
		{
			name: "GetUrlErrorGetUrl",
			args: args{
				shortUrl: "ErrorUriDB",
			},
			want:    "",
			wantErr: url.ErrGetUrl,
		},
		{
			name: "GetUrlErrorEmptyUrl",
			args: args{
				shortUrl: "",
			},
			want:    "",
			wantErr: url.ErrGetUrl,
		},
	}

	service := url.NewService(repoMock)

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repoMock.Test(t)
			got, err := service.GetUrl(tt.args.shortUrl)

			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
