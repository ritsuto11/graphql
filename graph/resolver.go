package graph

import "my_gql_server/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos   []*model.Todo
	devices []*model.Device
}
