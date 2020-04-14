// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for github.com/go-pg/pg/orm, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: github.com/go-pg/pg/orm (exports: Query; functions: Q)

// Package orm is a stub of github.com/go-pg/pg/orm, generated by depstubber.
package orm

import (
	context "context"
	io "io"
	reflect "reflect"
)

type ColumnScanner interface {
	ScanColumn(_ int, _ string, _ interface{}, _ int) interface {
		Error() string
	}
}

type CreateTableOptions struct {
	Temp          bool
	IfNotExists   bool
	Varchar       int
	FKConstraints bool
}

type DB interface {
	Context() context.Context
	CopyFrom(_ io.Reader, _ interface{}, _ ...interface{}) (Result, interface {
		Error() string
	})
	CopyTo(_ io.Writer, _ interface{}, _ ...interface{}) (Result, interface {
		Error() string
	})
	Delete(_ interface{}) interface {
		Error() string
	}
	Exec(_ interface{}, _ ...interface{}) (Result, interface {
		Error() string
	})
	ExecContext(_ context.Context, _ interface{}, _ ...interface{}) (Result, interface {
		Error() string
	})
	ExecOne(_ interface{}, _ ...interface{}) (Result, interface {
		Error() string
	})
	ExecOneContext(_ context.Context, _ interface{}, _ ...interface{}) (Result, interface {
		Error() string
	})
	ForceDelete(_ interface{}) interface {
		Error() string
	}
	FormatQuery(_ []uint8, _ string, _ ...interface{}) []uint8
	Insert(_ ...interface{}) interface {
		Error() string
	}
	Model(_ ...interface{}) *Query
	ModelContext(_ context.Context, _ ...interface{}) *Query
	Query(_ interface{}, _ interface{}, _ ...interface{}) (Result, interface {
		Error() string
	})
	QueryContext(_ context.Context, _ interface{}, _ interface{}, _ ...interface{}) (Result, interface {
		Error() string
	})
	QueryOne(_ interface{}, _ interface{}, _ ...interface{}) (Result, interface {
		Error() string
	})
	QueryOneContext(_ context.Context, _ interface{}, _ interface{}, _ ...interface{}) (Result, interface {
		Error() string
	})
	Select(_ interface{}) interface {
		Error() string
	}
	Update(_ interface{}) interface {
		Error() string
	}
}

type DropTableOptions struct {
	IfExists bool
	Cascade  bool
}

type Field struct {
	Field    reflect.StructField
	Type     reflect.Type
	GoName   string
	SQLName  string
	Column   interface{}
	SQLType  string
	Index    []int
	Default  interface{}
	OnDelete string
	OnUpdate string
}

func (_ *Field) AppendValue(_ []uint8, _ reflect.Value, _ int) []uint8 {
	return nil
}

func (_ *Field) Copy() *Field {
	return nil
}

func (_ *Field) HasFlag(_ uint8) bool {
	return false
}

func (_ *Field) IsZeroValue(_ reflect.Value) bool {
	return false
}

func (_ *Field) OmitZero() bool {
	return false
}

func (_ *Field) ScanValue(_ reflect.Value, _ interface{}, _ int) interface {
	Error() string
} {
	return nil
}

func (_ *Field) SetFlag(_ uint8) {}

func (_ *Field) Value(_ reflect.Value) reflect.Value {
	return reflect.Value{}
}

type Method struct {
	Index int
}

func (_ *Method) AppendValue(_ []uint8, _ reflect.Value, _ int) []uint8 {
	return nil
}

func (_ *Method) Has(_ int8) bool {
	return false
}

func (_ *Method) Value(_ reflect.Value) reflect.Value {
	return reflect.Value{}
}

type Model interface {
	AddModel(_ ColumnScanner) interface {
		Error() string
	}
	AfterDelete(_ context.Context, _ DB) interface {
		Error() string
	}
	AfterInsert(_ context.Context, _ DB) interface {
		Error() string
	}
	AfterQuery(_ context.Context, _ DB) interface {
		Error() string
	}
	AfterSelect(_ context.Context, _ DB) interface {
		Error() string
	}
	AfterUpdate(_ context.Context, _ DB) interface {
		Error() string
	}
	BeforeDelete(_ context.Context, _ DB) interface {
		Error() string
	}
	BeforeInsert(_ context.Context, _ DB) interface {
		Error() string
	}
	BeforeSelectQuery(_ context.Context, _ DB, _ *Query) (*Query, interface {
		Error() string
	})
	BeforeUpdate(_ context.Context, _ DB) interface {
		Error() string
	}
	Init() interface {
		Error() string
	}
	NewModel() ColumnScanner
}

func Q(_ string, _ ...interface{}) *interface{} {
	return nil
}

type Query struct{}

func (_ *Query) AppendFormat(_ []uint8, _ QueryFormatter) []uint8 {
	return nil
}

func (_ *Query) Apply(_ func(*Query) (*Query, interface {
	Error() string
})) *Query {
	return nil
}

func (_ *Query) Column(_ ...string) *Query {
	return nil
}

func (_ *Query) ColumnExpr(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) Context(_ context.Context) *Query {
	return nil
}

func (_ *Query) Copy() *Query {
	return nil
}

func (_ *Query) CopyFrom(_ io.Reader, _ interface{}, _ ...interface{}) (Result, interface {
	Error() string
}) {
	return nil, nil
}

func (_ *Query) CopyTo(_ io.Writer, _ interface{}, _ ...interface{}) (Result, interface {
	Error() string
}) {
	return nil, nil
}

func (_ *Query) Count() (int, interface {
	Error() string
}) {
	return 0, nil
}

func (_ *Query) CountEstimate(_ int) (int, interface {
	Error() string
}) {
	return 0, nil
}

func (_ *Query) CreateTable(_ *CreateTableOptions) interface {
	Error() string
} {
	return nil
}

func (_ *Query) DB(_ DB) *Query {
	return nil
}

func (_ *Query) Delete(_ ...interface{}) (Result, interface {
	Error() string
}) {
	return nil, nil
}

func (_ *Query) Deleted() *Query {
	return nil
}

func (_ *Query) DropTable(_ *DropTableOptions) interface {
	Error() string
} {
	return nil
}

func (_ *Query) ExcludeColumn(_ ...string) *Query {
	return nil
}

func (_ *Query) Exec(_ interface{}, _ ...interface{}) (Result, interface {
	Error() string
}) {
	return nil, nil
}

func (_ *Query) ExecOne(_ interface{}, _ ...interface{}) (Result, interface {
	Error() string
}) {
	return nil, nil
}

func (_ *Query) Exists() (bool, interface {
	Error() string
}) {
	return false, nil
}

func (_ *Query) First() interface {
	Error() string
} {
	return nil
}

func (_ *Query) For(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) ForEach(_ interface{}) interface {
	Error() string
} {
	return nil
}

func (_ *Query) ForceDelete(_ ...interface{}) (Result, interface {
	Error() string
}) {
	return nil, nil
}

func (_ *Query) FormatQuery(_ []uint8, _ string, _ ...interface{}) []uint8 {
	return nil
}

func (_ *Query) Formatter(_ QueryFormatter) *Query {
	return nil
}

func (_ *Query) GetModel() TableModel {
	return nil
}

func (_ *Query) Group(_ ...string) *Query {
	return nil
}

func (_ *Query) GroupExpr(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) Having(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) Insert(_ ...interface{}) (Result, interface {
	Error() string
}) {
	return nil, nil
}

func (_ *Query) Join(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) JoinOn(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) JoinOnOr(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) Last() interface {
	Error() string
} {
	return nil
}

func (_ *Query) Limit(_ int) *Query {
	return nil
}

func (_ *Query) Model(_ ...interface{}) *Query {
	return nil
}

func (_ *Query) New() *Query {
	return nil
}

func (_ *Query) Offset(_ int) *Query {
	return nil
}

func (_ *Query) OnConflict(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) Order(_ ...string) *Query {
	return nil
}

func (_ *Query) OrderExpr(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) Query(_ interface{}, _ interface{}, _ ...interface{}) (Result, interface {
	Error() string
}) {
	return nil, nil
}

func (_ *Query) QueryOne(_ interface{}, _ interface{}, _ ...interface{}) (Result, interface {
	Error() string
}) {
	return nil, nil
}

func (_ *Query) Relation(_ string, _ ...func(*Query) (*Query, interface {
	Error() string
})) *Query {
	return nil
}

func (_ *Query) Returning(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) Select(_ ...interface{}) interface {
	Error() string
} {
	return nil
}

func (_ *Query) SelectAndCount(_ ...interface{}) (int, interface {
	Error() string
}) {
	return 0, nil
}

func (_ *Query) SelectAndCountEstimate(_ int, _ ...interface{}) (int, interface {
	Error() string
}) {
	return 0, nil
}

func (_ *Query) SelectOrInsert(_ ...interface{}) (bool, interface {
	Error() string
}) {
	return false, nil
}

func (_ *Query) Set(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) Table(_ ...string) *Query {
	return nil
}

func (_ *Query) TableExpr(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) Update(_ ...interface{}) (Result, interface {
	Error() string
}) {
	return nil, nil
}

func (_ *Query) UpdateNotNull(_ ...interface{}) (Result, interface {
	Error() string
}) {
	return nil, nil
}

func (_ *Query) Value(_ string, _ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) Where(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) WhereGroup(_ func(*Query) (*Query, interface {
	Error() string
})) *Query {
	return nil
}

func (_ *Query) WhereIn(_ string, _ interface{}) *Query {
	return nil
}

func (_ *Query) WhereInMulti(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) WhereOr(_ string, _ ...interface{}) *Query {
	return nil
}

func (_ *Query) WhereOrGroup(_ func(*Query) (*Query, interface {
	Error() string
})) *Query {
	return nil
}

func (_ *Query) WherePK() *Query {
	return nil
}

func (_ *Query) WhereStruct(_ interface{}) *Query {
	return nil
}

func (_ *Query) With(_ string, _ *Query) *Query {
	return nil
}

func (_ *Query) WrapWith(_ string) *Query {
	return nil
}

type QueryFormatter interface {
	FormatQuery(_ []uint8, _ string, _ ...interface{}) []uint8
}

type Relation struct {
	Type          int
	Field         *Field
	JoinTable     *Table
	FKs           []*Field
	Polymorphic   *Field
	FKValues      []*Field
	M2MTableName  interface{}
	M2MTableAlias interface{}
	BaseFKs       []string
	JoinFKs       []string
}

func (_ *Relation) String() string {
	return ""
}

type Result interface {
	Model() Model
	RowsAffected() int
	RowsReturned() int
}

type Table struct {
	Type               reflect.Type
	TypeName           string
	Alias              interface{}
	ModelName          string
	Name               string
	FullName           interface{}
	FullNameForSelects interface{}
	Tablespace         interface{}
	Fields             []*Field
	PKs                []*Field
	DataFields         []*Field
	FieldsMap          map[string]*Field
	Methods            map[string]*Method
	Relations          map[string]*Relation
	Unique             map[string][]*Field
	SoftDeleteField    *Field
}

func (_ *Table) AddField(_ *Field) {}

func (_ *Table) AppendParam(_ []uint8, _ reflect.Value, _ string) ([]uint8, bool) {
	return nil, false
}

func (_ *Table) GetField(_ string) (*Field, interface {
	Error() string
}) {
	return nil, nil
}

func (_ *Table) HasField(_ string) bool {
	return false
}

func (_ *Table) HasFlag(_ uint16) bool {
	return false
}

func (_ *Table) RemoveField(_ *Field) {}

func (_ *Table) SetFlag(_ uint16) {}

func (_ *Table) String() string {
	return ""
}

type TableModel interface {
	AddJoin(_ interface{}) *interface{}
	AddModel(_ ColumnScanner) interface {
		Error() string
	}
	AfterDelete(_ context.Context, _ DB) interface {
		Error() string
	}
	AfterInsert(_ context.Context, _ DB) interface {
		Error() string
	}
	AfterQuery(_ context.Context, _ DB) interface {
		Error() string
	}
	AfterSelect(_ context.Context, _ DB) interface {
		Error() string
	}
	AfterUpdate(_ context.Context, _ DB) interface {
		Error() string
	}
	AppendParam(_ []uint8, _ QueryFormatter, _ string) ([]uint8, bool)
	BeforeDelete(_ context.Context, _ DB) interface {
		Error() string
	}
	BeforeInsert(_ context.Context, _ DB) interface {
		Error() string
	}
	BeforeSelectQuery(_ context.Context, _ DB, _ *Query) (*Query, interface {
		Error() string
	})
	BeforeUpdate(_ context.Context, _ DB) interface {
		Error() string
	}
	GetJoin(_ string) *interface{}
	GetJoins() []interface{}
	Index() []int
	Init() interface {
		Error() string
	}
	IsNil() bool
	Join(_ string, _ func(*Query) (*Query, interface {
		Error() string
	})) *interface{}
	Kind() reflect.Kind
	Mount(_ reflect.Value)
	NewModel() ColumnScanner
	ParentIndex() []int
	Relation() *Relation
	Root() reflect.Value
	Table() *Table
	Value() reflect.Value
}