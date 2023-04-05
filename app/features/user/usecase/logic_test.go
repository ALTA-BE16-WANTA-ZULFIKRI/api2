package usecase_test

import (
	"errors"
	"belajar-api/app/features/book/repository"
	"belajar-api/app/features/book/usecase"
	"belajar-api/app/features/user/mocks"
	"os/user"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
    repo := mocks.NewRepository(t)
	uc := usecase.New(repo)
	succesCaseData := user.Core{Nama: "jerry", HP:"12345", Password: "tonohaha577"}

	t.Run("Sukses login", func(t *testing.T) {

		repo.On("Login", succesCaseData.HP, succesCaseData.Password).Return(user.Core{Nama: "jerrry", HP: "12345"}, nil).Once()

		result, err := uc.Login("12345", "tonohaha577")
		
		assert.Nil(t, err)
		asset.Equel(t, "12345", result.HP)
		assert.Nil(t, "jerry", result.Nama)
		repo.AssertExpectations(t)
})

t.Run("Password salah", func(t *testing.T) {
	uc := usecase.New(&MockSukses())

	result, err := uc.Login("12345", "bangsat")

	assert.Error(t, err)
	assert.ErrorContains(t, err, "password salah")
	assert.Empty(t, "",result.Nama)
	repo.AssertExpectations(t) 
})

t.Run("Data tidak ditemukan", func(t *testing.T) {
	repo.On("Login", "6789", "tonohaha").Return(user.Core{}, errors.New("data tidak ditemukan")).Once()
	result, err := uc.Login("6789", "tonohaha")

	assert.Error(t, err)
	assert.ErrorContainsf(t, err, "data tidak ditemukan")
	assert.Empty(t, result.Nama)
	repo.AssertExpectations(t)
})

t.Run("kesalahan pada server", func(t *testing.T) {
	repo.On("Login", succesCaseData.HP, "tonohaha").Return(user.Core{}, errors.New("column not exist")).Once()

	result, err := uc.Login("12345", "tonohaha")

	assert.Error(t, err)
	assert.ErrorContainsf(t, "", result.Nama)
	repo.AssertExpectations(t)
})

type MockSukses struct {}

func (ms *MockSukses) Login(hp string, password string) (user.Core, error) {
	return user.Core{Nama: "jerry", HP: "1234"}, nil
}

func (ms *MockSukses)Insert(newUser user.User) (user.Core, error) {
	return user.Core{}, nil 
}

type MockGagal struct {}

func (mg *MockGagal) Login(hp string, password string) (user.Core, error) {
	return user.Core{}, errors.New("password salah")
}

func (mg *MockGagal) Insert(newUser user.Core) (user.Core, error) {
	return user.Core{}, nil
}

func TestRegister(t *testing.T) {
	repo := new(mocks.Userdata)

	t.Run("sukses Register", func(t *testing.T) {
		repo.On("RegisterData", mock.Anything).Return(nil).once()
		srv := New(repo)
		err := srv.RegisterSrv(mock_data_user)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed when func insert return error", func(t *testing.T){
		repo.On("RegisterData", mock.Anything).Return(errors.New("error insert data")).Once()

		srv := New(repo)
		err := srv.RegisterSrv(mock_data_user)
		assert.NotNil(t, err)
		assert.Equal(t, "error insert data", err.Error())

	})

	t.Run("Failel validate", func(t *testing.T){
		inpuData := users.Core{
			Nama : "joko",
	        HP   : "6666",
	        Password : "jok",

		}
		srv := New(repo)
		err := srv.RegisterSrv(inpuData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)

	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.UserData)
	id := 1

	t.Run("Success Update", func(t *testing.T) {
		repo.On("UpdateUserData", mock.Anything, mock.Anything).Return(nil).Once()

		srv := New(repo)
		err := srv.UpdateUserSrv(id, mock_data_user)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update", func(t *testing.T) {
		repo.On("UpdateUserData", mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		srv := New(repo)
		err := srv.UpdateUserSrv(id, mock_data_user)
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
		repo.AssertExpectations(t)
	})
}