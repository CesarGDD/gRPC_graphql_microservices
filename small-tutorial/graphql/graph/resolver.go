package graph

import (
	"context"

	"github.com/99designs/gqlgen/plugin/federation/test_data/model"
)

type queryResolver struct{}

// Resolver para la consulta person
func (r *queryResolver) Person(ctx context.Context, id int) (*model.Person, error) {
	// Aquí puedes agregar la lógica para buscar la persona en una base de datos
	// o un servicio externo. Por simplicidad, vamos a devolver un valor fijo.
	person := &model.Person{
		ID:    id,
		Name:  "Juan Perez",
		Email: "juanperez@example.com",
	}

	return person, nil
}
