package main

import (
	"8_mockery/mocks"
	"8_mockery/model"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func fixtures() (s *UserService, m *mocks.UserDB) {
	m = &mocks.UserDB{}
	s = &UserService{
		userDB: m,
	}
	return s, m
}
func TestSave(t *testing.T) {
	// given
	s, m := fixtures()
	user1 := model.User{
		Email: "user1@gmail.com",
		Name:  "user1",
	}
	userMatcher := func(u *model.User) bool {
		return u.Email == user1.Email && u.Name == user1.Name
	}
	m.On("Save", mock.Anything, mock.MatchedBy(userMatcher)).Return(nil)

	// when
	err := s.Save(context.TODO(), &user1)

	// then
	assert.NoError(t, err)
	m.AssertCalled(t, "Save", mock.Anything, mock.MatchedBy(userMatcher))
	m.AssertNumberOfCalls(t, "Save", 1)
}

func TestSave_Fail(t *testing.T) {
	cases := []struct {
		Name       string
		User       model.User
		SetupMock  func(m *mocks.UserDB)
		Msg        string
		AssertMock func(t *testing.T, m *mocks.UserDB)
	}{
		{
			Name: "empty email",
			User: model.User{},
			Msg:  "invalid email",
			AssertMock: func(t *testing.T, m *mocks.UserDB) {
				m.AssertNotCalled(t, "Save")
			},
		}, {
			Name: "empty name",
			User: model.User{Email: "user1@email.com"},
			Msg:  "invalid name",
			AssertMock: func(t *testing.T, m *mocks.UserDB) {
				m.AssertNotCalled(t, "Save")
			},
		}, {
			Name: "duplicate email",
			User: model.User{Email: "user1@email.com", Name: "user1"},
			SetupMock: func(m *mocks.UserDB) {
				m.On("Save", mock.Anything, mock.Anything).Return(ErrKeyConflict)
			},
			Msg: "duplicate email",
			AssertMock: func(t *testing.T, m *mocks.UserDB) {
				m.AssertNumberOfCalls(t, "Save", 1)
			},
		}, {
			Name: "any error",
			User: model.User{Email: "user1@email.com", Name: "user1"},
			SetupMock: func(m *mocks.UserDB) {
				m.On("Save", mock.Anything, mock.Anything).Return(errors.New("any error"))
			},
			Msg: "any error",
			AssertMock: func(t *testing.T, m *mocks.UserDB) {
				m.AssertNumberOfCalls(t, "Save", 1)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			// t.Parallel() // 위의 TestCases를 병렬로 실행할 수 있습니다.
			// given
			s, m := fixtures()
			if tc.SetupMock != nil {
				tc.SetupMock(m)
			}

			// when
			err := s.Save(context.TODO(), &tc.User)

			// then
			assert.Error(t, err)
			if tc.AssertMock != nil {
				tc.AssertMock(t, m)
			}
		})
	}
}
