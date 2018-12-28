package model

import (
	"testing"

	mock "github.com/hemat/model/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	usr := NewUserService("Udin", 20)

	// t.Run("Run failed scenario", func(t *testing.T) {
	// 	data, err := usr.GetAll()

	// 	assert.Error(t, err)
	// 	assert.Len(t, data, 0)

	// })

	t.Run("Run success scenario", func(t *testing.T) {
		mock := new(mock.RepoUser)

		expectReturn := []map[string]string{
			map[string]string{"name": "udin", "age": "50"},
			map[string]string{"name": "bruce", "age": "30"},
		}

		mock.On("GetAll").Return(expectReturn, nil)

		data, err := usr.GetAll()
		exData, _ := mock.GetAll()

		assert.NoError(t, err)
		assert.Len(t, data, 2)
		assert.Equal(t, exData, data)

		mock.AssertExpectations(t)
		mock.AssertNumberOfCalls(t, "GetAll", 1)

	})

	// t.Run("Run with model sucks !", func(t *testing.T) {
	// 	var usr User

	// 	var expectRes []map[string]string

	// 	res, err := usr.GetAll()

	// 	assert.IsType(t, reflect.TypeOf(expectRes), reflect.TypeOf(res))
	// 	assert.Error(t, err)

	// })
}

func TestGetFromDatabase(t *testing.T) {

	type fakeStruct struct {
		name string
	}

	var fake []fakeStruct

	_, err := GetFromDatabase(fake)

	assert.Error(t, err)

	// fake2 := []fakeStruct{
	// 	fakeStruct{
	// 		name: "Samsul 1",
	// 		age:  11,
	// 	},
	// 	fakeStruct{
	// 		name: "Samsul 2",
	// 		age:  12,
	// 	},
	// }

	// res, errSec2 := GetFromDatabase(&fake2)

	// assert.NoError(t, errSec2)
	// assert.Len(t, res, 1)

	newParams := []map[string]string{
		map[string]string{"name": "paijo", "age": "20"},
	}

	resC, errC := GetFromDatabase(newParams)
	assert.NoError(t, errC)
	assert.Len(t, resC, 1)

	res3, err3 := GetFromDatabase(fakeStruct{})
	assert.NoError(t, err3)
	assert.Len(t, res3, 1)
}
