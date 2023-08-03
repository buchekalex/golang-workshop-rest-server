package usecase_test

import (
	"context"
	"errors"
	"go-mongo-crud-rest-api/internal/entitiy"
	"go-mongo-crud-rest-api/internal/usecase"
	"testing"

	"github.com/golang/mock/gomock"
)

type mocks struct {
	repository  *usecase.MockRepository
	searchIndex *usecase.MockSearchIndex
}

func TestUserUseCase_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks{
		repository:  usecase.NewMockRepository(ctrl),
		searchIndex: usecase.NewMockSearchIndex(ctrl),
	}

	testCases := []struct {
		name    string
		email   string
		prepare func()
		want    *entitiy.User
		wantErr error
	}{
		{
			name:  "successful retrieval of user",
			email: "test@example.com",
			prepare: func() {
				m.repository.EXPECT().GetUser(gomock.Any(), "test@example.com").Return(&entitiy.User{Email: "test@example.com"}, nil)
			},
			want: &entitiy.User{
				Email: "test@example.com",
			},
			wantErr: nil,
		},
		{
			name:  "failed retrieval of user",
			email: "test@example.com",
			prepare: func() {
				m.repository.EXPECT().GetUser(gomock.Any(), "test@example.com").Return(nil, errors.New("user not found"))
			},
			want:    nil,
			wantErr: errors.New("user not found"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()

			u := usecase.NewUserUseCase(m.repository, m.searchIndex)
			got, err := u.GetUser(context.Background(), tt.email)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
