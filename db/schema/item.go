package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("文件描述"),
		field.Bytes("file").Comment("文件内容"),
		field.Time("created_At").
			Default(time.Now).
			Comment("创建时间"),
		field.Time("updated_At").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间"),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return nil
}
