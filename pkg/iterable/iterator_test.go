package iterable_test

import (
	"fmt"
	"testing"

	"github.com/jeeo/functional-go/pkg/iterable"
	"github.com/jeeo/functional-go/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestHasNext(t *testing.T) {
	people := []model.Person{
		{
			Name: "John Doe",
			Age:  20,
		},
		{
			Name: "foo",
			Age:  21,
		},
		{
			Name: "bar",
			Age:  20,
		},
	}
	expectedResult := []bool{true, true, false}
	iterable := iterable.NewIterable(people)
	result := make([]bool, 0)

	for range people {
		fmt.Println(iterable.Next())
		result = append(result, iterable.HasNext())
	}

	assert.EqualValues(t, expectedResult, result)
}

func TestNext(t *testing.T) {
	people := []model.Person{
		{
			Name: "John Doe",
			Age:  20,
		},
		{
			Name: "foo",
			Age:  21,
		},
		{
			Name: "bar",
			Age:  20,
		},
	}
	expectedResult := []*model.Person{&people[0], &people[1], &people[2]}
	result := make([]*model.Person, 0)
	iterable := iterable.NewIterable(people)
	for iterable.HasNext() {
		result = append(result, iterable.Next())
	}

	assert.EqualValues(t, expectedResult, result)
}
