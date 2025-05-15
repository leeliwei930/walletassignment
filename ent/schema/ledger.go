package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Ledger holds the schema definition for the Ledger entity.
type Ledger struct {
	ent.Schema
}

// Fields of the Ledger.
func (Ledger) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("wallet_id", uuid.UUID{}),
		field.Int("amount").Positive(),
		field.String("description"),
		field.String("transaction_type"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Ledger.
func (Ledger) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("wallet", Wallet.Type).Ref("ledgers").Field("wallet_id").Required().Unique(),
	}
}

func (Ledger) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("wallet_id", "transaction_type"),
		index.Fields("wallet_id", "created_at"),
		index.Fields("transaction_type"),
		index.Fields("created_at"),
	}
}
