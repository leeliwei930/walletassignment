package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/lock --feature sql/versioned-migration --feature sql/upsert ./schema
