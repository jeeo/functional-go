package tools_test

import (
	"testing"

	"github.com/jeeo/functional-go/pkg/iterable"
	"github.com/jeeo/functional-go/pkg/model"
	"github.com/jeeo/functional-go/pkg/tools"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
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

	expected := []*struct {
		model.Person
		Like bool
	}{
		{
			Person: model.Person{
				Name: "John Doe-foo!",
				Age:  20,
			},
			Like: true,
		},
		{
			Person: model.Person{
				Name: "bar-foo!",
				Age:  20,
			},
			Like: true,
		},
	}

	peopleIterator := iterable.NewIterable(people)
	foer := func(p model.Person) *model.Person {
		newPerson := p
		newPerson.Name += "-foo!"
		return &newPerson
	}
	addProperty := func(p model.Person) *struct {
		model.Person
		Like bool
	} {
		return &struct {
			model.Person
			Like bool
		}{
			Person: p,
			Like:   true,
		}
	}
	foerIterable := tools.Map(peopleIterator, foer)
	ageFilterIterable := tools.Filter(foerIterable, func(p model.Person) bool { return p.Age <= 20 })
	coolIterable := tools.Map(ageFilterIterable, addProperty)
	var results []*struct {
		model.Person
		Like bool
	}
	for coolIterable.HasNext() {
		results = append(results, coolIterable.Next())
	}

	assert.EqualValues(t, expected, results)
}
