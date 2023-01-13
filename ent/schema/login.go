package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Login struct {
	ent.Schema
}

func (Login) Fields() []ent.Field {
	return []ent.Field{
		// Field id autoincrements automatically, due to the name "id", as its a magic identifier in ent.
		field.
			Int("id").
			StorageKey("Id"),
		field.String("name"),
	}
}

func (Login) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("login_permissions", LoginPermission.Type).
			Through("login_permission_allocations", LoginPermissionAllocation.Type).
			Ref("logins"),
	}
}
