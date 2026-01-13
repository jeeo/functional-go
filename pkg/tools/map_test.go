package tools_test

import (
	"testing"

	"github.com/jeeo/functional-go/pkg/iterable"
	"github.com/jeeo/functional-go/pkg/model"
	"github.com/jeeo/functional-go/pkg/tools"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
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
	expected := []*struct{ Name string }{{"John Doe-foo!"}, {"foo-foo!"}, {"bar-foo!"}}
	peopleIterator := iterable.NewIterable(people)
	foer := func(p model.Person) *model.Person {
		newPerson := p
		newPerson.Name += "-foo!"
		return &newPerson
	}
	cutAges := func(p model.Person) *struct{ Name string } {
		return &struct{ Name string }{
			p.Name,
		}
	}
	foerIterable := tools.Map(peopleIterator, foer)
	cutterIterable := tools.Map(foerIterable, cutAges)

	var results []*struct{ Name string }
	for cutterIterable.HasNext() {
		results = append(results, cutterIterable.Next())
	}

	assert.EqualValues(t, expected, results)
}
