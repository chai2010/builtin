// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builtin

import (
	"encoding/json"
	"reflect"
	"testing"
)

type tUser struct {
	Name  string
	Role  string
	Age   int32
	Notes []string
}

func (user *tUser) DoubleAge() int32 {
	return 2 * user.Age
}

type tEmployee struct {
	Name      string
	Age       int32
	EmployeId int64
	DoubleAge int32
	SuperRule string
	Notes     []string
}

func (employee *tEmployee) Role(role string) {
	employee.SuperRule = "Super " + role
}

func TestCopy_struct(t *testing.T) {
	user := tUser{Name: "Gopher", Age: 18, Role: "Admin", Notes: []string{"hello world"}}
	employee := tEmployee{}

	Copy(&employee, &user)

	if employee.Name != "Gopher" {
		t.Errorf("Name haven't been copied correctly.")
	}
	if employee.Age != 18 {
		t.Errorf("Age haven't been copied correctly.")
	}
	if employee.DoubleAge != 36 {
		t.Errorf("Copy copy from method doesn't work")
	}
	if employee.SuperRule != "Super Admin" {
		t.Errorf("Copy Attributes should support copy to method")
	}

	if !reflect.DeepEqual(employee.Notes, []string{"hello world"}) {
		t.Errorf("Copy a map")
	}

	user.Notes = append(user.Notes, "welcome")
	if !reflect.DeepEqual(user.Notes, []string{"hello world", "welcome"}) {
		t.Errorf("tUser's Note should be changed")
	}

	if !reflect.DeepEqual(employee.Notes, []string{"hello world"}) {
		t.Errorf("tEmployee's Note should not be changed")
	}

	employee.Notes = append(employee.Notes, "golang")
	if !reflect.DeepEqual(employee.Notes, []string{"hello world", "golang"}) {
		t.Errorf("tEmployee's Note should be changed")
	}

	if !reflect.DeepEqual(user.Notes, []string{"hello world", "welcome"}) {
		t.Errorf("tEmployee's Note should not be changed")
	}
}

func TestCopy_slice(t *testing.T) {
	user := tUser{Name: "Gopher", Age: 18, Role: "Admin", Notes: []string{"hello world"}}
	users := []tUser{{Name: "jinzhu 2", Age: 30, Role: "Dev"}}
	employees := []tEmployee{}

	Copy(&employees, &user)
	if len(employees) != 1 {
		t.Errorf("Should only have one elem when copy struct to slice")
	}

	Copy(&employees, &users)
	if len(employees) != 2 {
		t.Errorf("Should have two elems when copy additional slice to slice")
	}

	if employees[0].Name != "Gopher" {
		t.Errorf("Name haven't been copied correctly.")
	}
	if employees[0].Age != 18 {
		t.Errorf("Age haven't been copied correctly.")
	}
	if employees[0].DoubleAge != 36 {
		t.Errorf("Copy copy from method doesn't work")
	}
	if employees[0].SuperRule != "Super Admin" {
		t.Errorf("Copy Attributes should support copy to method")
	}

	if employees[1].Name != "jinzhu 2" {
		t.Errorf("Name haven't been copied correctly.")
	}
	if employees[1].Age != 30 {
		t.Errorf("Age haven't been copied correctly.")
	}
	if employees[1].DoubleAge != 60 {
		t.Errorf("Copy copy from method doesn't work")
	}
	if employees[1].SuperRule != "Super Dev" {
		t.Errorf("Copy Attributes should support copy to method")
	}

	employee := employees[0]
	user.Notes = append(user.Notes, "welcome")
	if !reflect.DeepEqual(user.Notes, []string{"hello world", "welcome"}) {
		t.Errorf("tUser's Note should be changed")
	}

	if !reflect.DeepEqual(employee.Notes, []string{"hello world"}) {
		t.Errorf("tEmployee's Note should not be changed")
	}

	employee.Notes = append(employee.Notes, "golang")
	if !reflect.DeepEqual(employee.Notes, []string{"hello world", "golang"}) {
		t.Errorf("tEmployee's Note should be changed")
	}

	if !reflect.DeepEqual(user.Notes, []string{"hello world", "welcome"}) {
		t.Errorf("tEmployee's Note should not be changed")
	}
}

func BenchmarkCopy_struct(b *testing.B) {
	user := tUser{Name: "Gopher", Age: 18, Role: "Admin", Notes: []string{"hello world"}}
	for x := 0; x < b.N; x++ {
		Copy(&tEmployee{}, &user)
	}
}

func BenchmarkNamaCopy(b *testing.B) {
	user := tUser{Name: "Gopher", Age: 18, Role: "Admin", Notes: []string{"hello world"}}
	for x := 0; x < b.N; x++ {
		employee := &tEmployee{
			Name:      user.Name,
			Age:       user.Age,
			DoubleAge: user.DoubleAge(),
			Notes:     user.Notes,
		}
		employee.Role(user.Role)
	}
}

func BenchmarkJsonMarshalCopy(b *testing.B) {
	user := tUser{Name: "Gopher", Age: 18, Role: "Admin", Notes: []string{"hello world"}}
	for x := 0; x < b.N; x++ {
		data, _ := json.Marshal(user)
		var employee tEmployee
		json.Unmarshal(data, &employee)
		employee.DoubleAge = user.DoubleAge()
		employee.Role(user.Role)
	}
}
