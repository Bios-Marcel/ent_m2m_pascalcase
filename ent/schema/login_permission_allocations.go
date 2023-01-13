package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type LoginPermissionAllocation struct {
	ent.Schema
}

func (LoginPermissionAllocation) Fields() []ent.Field {
	return []ent.Field{
		// Field id autoincrements automatically, due to the name "id", as its a magic identifier in ent.
		field.
			Int("id").
			StorageKey("Id"),
		field.Int("login_id"),            /*.StorageKey("LoginId") will cause breakages*/
		field.Int("login_permission_id"), /*.StorageKey("LoginPermissionId") will cause breakages*/
	}
}

func (LoginPermissionAllocation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("logins", Login.Type).
			Required().
			Unique().
			Field("login_id"),
		edge.
			To("login_permissions", Login.Type).
			Required().
			Unique().
			Field("login_permission_id"),
	}
}
