package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Wallet holds the schema definition for the Wallet entity.
type Wallet struct {
	ent.Schema
}

// Fields of the Wallet.
func (Wallet) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("user_id", uuid.UUID{}),
		field.Int("balance").Default(0),
		field.Int("currency_code").Default(10),
		field.Int("decimal_places").Default(2),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Wallet.
func (Wallet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("wallets").Field("user_id").Required().Unique(),
		edge.To("ledgers", Ledger.Type),
	}
}
