// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/dsykes16/entpoc/ent/todo"
	"github.com/rs/xid"
)

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Text      string
	Status    todo.Status
	Priority  *int
	ParentID  *xid.ID
	ChildIDs  []xid.ID
}

// Mutate applies the CreateTodoInput on the TodoMutation builder.
func (i *CreateTodoInput) Mutate(m *TodoMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetText(i.Text)
	m.SetStatus(i.Status)
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if v := i.ChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
}

// SetInput applies the change-set in the CreateTodoInput on the TodoCreate builder.
func (c *TodoCreate) SetInput(i CreateTodoInput) *TodoCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateTodoInput represents a mutation input for updating todos.
type UpdateTodoInput struct {
	UpdatedAt      *time.Time
	Text           *string
	Status         *todo.Status
	Priority       *int
	ClearParent    bool
	ParentID       *xid.ID
	ClearChildren  bool
	AddChildIDs    []xid.ID
	RemoveChildIDs []xid.ID
}

// Mutate applies the UpdateTodoInput on the TodoMutation builder.
func (i *UpdateTodoInput) Mutate(m *TodoMutation) {
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.Text; v != nil {
		m.SetText(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if i.ClearParent {
		m.ClearParent()
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if i.ClearChildren {
		m.ClearChildren()
	}
	if v := i.AddChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.RemoveChildIDs; len(v) > 0 {
		m.RemoveChildIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdate builder.
func (c *TodoUpdate) SetInput(i UpdateTodoInput) *TodoUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateTodoInput on the TodoUpdateOne builder.
func (c *TodoUpdateOne) SetInput(i UpdateTodoInput) *TodoUpdateOne {
	i.Mutate(c.Mutation())
	return c
}