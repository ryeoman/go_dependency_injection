package autocomplete

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_GetSuggestion(t *testing.T) {

	type args struct {
		ctx     context.Context
		keyword string
	}
	tests := []struct {
		name       string
		args       args
		newUsecase func(mockSearchRepo *MockSearchProvider) *Usecase
		want       Suggestions
		wantErr    bool
	}{
		{
			name: "Given empty input keyword, will return empty suggestion",
			args: args{
				ctx:     context.Background(),
				keyword: "",
			},
			newUsecase: func(mockSearchRepo *MockSearchProvider) *Usecase {
				return &Usecase{}
			},
			want:    Suggestions{},
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockSearchRepo := NewMockSearchProvider(ctrl)

			u := test.newUsecase(mockSearchRepo)
			got, err := u.GetSuggestion(test.args.ctx, test.args.keyword)
			if test.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
