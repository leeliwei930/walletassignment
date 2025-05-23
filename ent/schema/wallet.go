package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.String("currency_code").Default("USD"),
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

func (Wallet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}
