package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type LoginPermission struct {
	ent.Schema
}

func (LoginPermission) Fields() []ent.Field {
	return []ent.Field{
		// Field id autoincrements automatically, due to the name "id", as its a magic identifier in ent.
		field.
			Int("id").
			StorageKey("Id"),
		field.String("name"),
	}
}

func (LoginPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("logins", Login.Type).
			Through("login_permission_allocations", LoginPermissionAllocation.Type),
	}
}
