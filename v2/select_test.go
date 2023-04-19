//go:build v2

package orm

import (
	"database/sql"
	"reflect"
	"testing"
)

type TestModel struct {
	Id        int64
	FirstName string
	Age       int8
	LastName  *sql.NullString
}

func TestSelector_Build(t *testing.T) {
	type testCase[T any] struct {
		name    string
		q       QueryBuilder
		want    *Query
		wantErr bool
	}
	//goland:noinspection ALL
	tests := []testCase[TestModel]{
		// TODO: Add test cases.
		{
			// from 不调用
			name: "from not called",
			q:    NewSelector[TestModel](),
			want: &Query{
				SQL: "SELECT * FROM `TestModel`;",
			},
		},
		{
			// from 调用
			name: "from called",
			q:    NewSelector[TestModel]().From("`test_model_t`"),
			want: &Query{
				SQL: "SELECT * FROM `test_model_t`;",
			},
		},
		{
			// 调用 FROM，但是传入空字符串
			name: "empty from",
			q:    NewSelector[TestModel]().From(""),
			want: &Query{
				SQL: "SELECT * FROM `TestModel`;",
			},
		},
		{
			// 调用 FROM，同时出入看了 DB
			name: "with db",
			q:    NewSelector[TestModel]().From("`test_db`.`test_model`"),
			want: &Query{
				SQL: "SELECT * FROM `test_db`.`test_model`;",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Build()
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
		})
	}
}
