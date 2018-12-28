package model

import (
	"fmt"
	"reflect"
)

// User struct for model of user
type User struct {
	name string
	age  int
}

// RepoUser for list methods of User struct
type RepoUser interface {
	GetAll() ([]map[string]string, error)
}

// NewUserService for get instance of user
func NewUserService(name string, age int) *User {
	return &User{
		name: name,
		age:  age,
	}
}

// GetAll fot getting all user data
func (usr *User) GetAll() ([]map[string]string, error) {
	var result []map[string]string

	var er error
	er = fmt.Errorf("Udin gak ketemu")

	result = []map[string]string{
		map[string]string{"name": "udin", "age": "50"},
		map[string]string{"name": "bruce", "age": "30"},
	}
	er = nil

	return result, er
}

// GetFromDatabase for get data base on your struct
func GetFromDatabase(params interface{}) ([]map[string]string, error) {

	var result []map[string]string

	var elem reflect.Value
	validateElem := true
	var num int
	typeInterface := reflect.TypeOf(params).Kind()
	var er error

	er = fmt.Errorf("Udin gak ketemu")

	switch typeInterface {
	case reflect.Slice:
		elem = reflect.ValueOf(params)
		num = elem.Len()
	case reflect.Ptr:
		elem = reflect.ValueOf(params).Elem()
		num = elem.Len()
	case reflect.Struct:
		elem = reflect.ValueOf(params)
		num = elem.NumField()

	default:
		validateElem = false
	}

	if validateElem == false {
		return result, fmt.Errorf("type not found")
	}

	// fmt.Println("top", typeInterface, elem, elem.Interface(), num)

	for a := 0; a < num; a++ {
		if typeInterface == reflect.Struct {
			st := elem.Type().Field(a)
			newMap := map[string]string{
				string(st.Name): string(st.Type.String()),
			}

			result = append(result, newMap)
			er = nil
			continue
		}

		i := elem.Index(a)
		if typeInterface == reflect.Slice {
			var newMap = make(map[string]string, i.Len())
			for _, v := range i.MapKeys() {
				dol := i.MapIndex(v)
				newMap[v.String()] = dol.String()
			}
			result = append(result, newMap)
			er = nil
			continue
		}

		// if reflect.TypeOf(i).Kind() == reflect.Struct {
		// 	for j := 0; j < i.NumField(); j++ {
		// 		// st := i.Type().Field(j)
		// 		fmt.Println(i, i.Interface())
		// 		// newMap := map[string]string{
		// 		// 	string(st.Name): string("a"),
		// 		// }
		// 		// fmt.Println(newMap)
		// 		// result = append(result, newMap)
		// 	}
		// 	er = nil
		// }

	}

	return result, er
}
