package test

import (
	"fmt"
	"reflect"
)

type GameUser struct {
	Name string
	Id   int32
}

func (user GameUser) getId() {
	fmt.Println("Name = ", user.Id)
}

type User interface {
	getName()

	getId()
}

func (user GameUser) getName() {
	fmt.Println("Name = ", user.Name)
}

func ShowTest() {

	//typ_User := reflect.TypeOf((*User)(nil))
	typ_User := reflect.TypeOf((*User)(nil))

	fmt.Println(typ_User)

}
