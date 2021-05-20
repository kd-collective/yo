// Code generated by yo. DO NOT EDIT.

// Package models contains the types.
package models

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
	"google.golang.org/grpc/codes"
)

// CustomPrimitiveType represents a row from 'CustomPrimitiveTypes'.
type CustomPrimitiveType struct {
	PKey              string  `spanner:"PKey" json:"PKey"`                           // PKey
	FTInt64           int64   `spanner:"FTInt64" json:"FTInt64"`                     // FTInt64
	FTInt64null       int64   `spanner:"FTInt64Null" json:"FTInt64Null"`             // FTInt64Null
	FTInt32           int32   `spanner:"FTInt32" json:"FTInt32"`                     // FTInt32
	FTInt32null       int32   `spanner:"FTInt32Null" json:"FTInt32Null"`             // FTInt32Null
	FTInt16           int16   `spanner:"FTInt16" json:"FTInt16"`                     // FTInt16
	FTInt16null       int16   `spanner:"FTInt16Null" json:"FTInt16Null"`             // FTInt16Null
	FTInt8            int8    `spanner:"FTInt8" json:"FTInt8"`                       // FTInt8
	FTInt8null        int8    `spanner:"FTInt8Null" json:"FTInt8Null"`               // FTInt8Null
	FTUInt64          uint64  `spanner:"FTUInt64" json:"FTUInt64"`                   // FTUInt64
	FTUInt64null      uint64  `spanner:"FTUInt64Null" json:"FTUInt64Null"`           // FTUInt64Null
	FTUInt32          uint32  `spanner:"FTUInt32" json:"FTUInt32"`                   // FTUInt32
	FTUInt32null      uint32  `spanner:"FTUInt32Null" json:"FTUInt32Null"`           // FTUInt32Null
	FTUInt16          uint16  `spanner:"FTUInt16" json:"FTUInt16"`                   // FTUInt16
	FTUInt16null      uint16  `spanner:"FTUInt16Null" json:"FTUInt16Null"`           // FTUInt16Null
	FTUInt8           uint8   `spanner:"FTUInt8" json:"FTUInt8"`                     // FTUInt8
	FTUInt8null       uint8   `spanner:"FTUInt8Null" json:"FTUInt8Null"`             // FTUInt8Null
	FTArrayInt64      []int64 `spanner:"FTArrayInt64" json:"FTArrayInt64"`           // FTArrayInt64
	FTArrayInt64null  []int64 `spanner:"FTArrayInt64Null" json:"FTArrayInt64Null"`   // FTArrayInt64Null
	FTArrayInt32      []int64 `spanner:"FTArrayInt32" json:"FTArrayInt32"`           // FTArrayInt32
	FTArrayInt32null  []int64 `spanner:"FTArrayInt32Null" json:"FTArrayInt32Null"`   // FTArrayInt32Null
	FTArrayInt16      []int64 `spanner:"FTArrayInt16" json:"FTArrayInt16"`           // FTArrayInt16
	FTArrayInt16null  []int64 `spanner:"FTArrayInt16Null" json:"FTArrayInt16Null"`   // FTArrayInt16Null
	FTArrayInt8       []int64 `spanner:"FTArrayInt8" json:"FTArrayInt8"`             // FTArrayInt8
	FTArrayInt8null   []int64 `spanner:"FTArrayInt8Null" json:"FTArrayInt8Null"`     // FTArrayInt8Null
	FTArrayUINt64     []int64 `spanner:"FTArrayUInt64" json:"FTArrayUInt64"`         // FTArrayUInt64
	FTArrayUINt64null []int64 `spanner:"FTArrayUInt64Null" json:"FTArrayUInt64Null"` // FTArrayUInt64Null
	FTArrayUINt32     []int64 `spanner:"FTArrayUInt32" json:"FTArrayUInt32"`         // FTArrayUInt32
	FTArrayUINt32null []int64 `spanner:"FTArrayUInt32Null" json:"FTArrayUInt32Null"` // FTArrayUInt32Null
	FTArrayUINt16     []int64 `spanner:"FTArrayUInt16" json:"FTArrayUInt16"`         // FTArrayUInt16
	FTArrayUINt16null []int64 `spanner:"FTArrayUInt16Null" json:"FTArrayUInt16Null"` // FTArrayUInt16Null
	FTArrayUINt8      []int64 `spanner:"FTArrayUInt8" json:"FTArrayUInt8"`           // FTArrayUInt8
	FTArrayUINt8null  []int64 `spanner:"FTArrayUInt8Null" json:"FTArrayUInt8Null"`   // FTArrayUInt8Null
}

func CustomPrimitiveTypePrimaryKeys() []string {
	return []string{
		"PKey",
	}
}

func CustomPrimitiveTypeColumns() []string {
	return []string{
		"PKey",
		"FTInt64",
		"FTInt64Null",
		"FTInt32",
		"FTInt32Null",
		"FTInt16",
		"FTInt16Null",
		"FTInt8",
		"FTInt8Null",
		"FTUInt64",
		"FTUInt64Null",
		"FTUInt32",
		"FTUInt32Null",
		"FTUInt16",
		"FTUInt16Null",
		"FTUInt8",
		"FTUInt8Null",
		"FTArrayInt64",
		"FTArrayInt64Null",
		"FTArrayInt32",
		"FTArrayInt32Null",
		"FTArrayInt16",
		"FTArrayInt16Null",
		"FTArrayInt8",
		"FTArrayInt8Null",
		"FTArrayUInt64",
		"FTArrayUInt64Null",
		"FTArrayUInt32",
		"FTArrayUInt32Null",
		"FTArrayUInt16",
		"FTArrayUInt16Null",
		"FTArrayUInt8",
		"FTArrayUInt8Null",
	}
}

func CustomPrimitiveTypeWritableColumns() []string {
	return []string{
		"PKey",
		"FTInt64",
		"FTInt64Null",
		"FTInt32",
		"FTInt32Null",
		"FTInt16",
		"FTInt16Null",
		"FTInt8",
		"FTInt8Null",
		"FTUInt64",
		"FTUInt64Null",
		"FTUInt32",
		"FTUInt32Null",
		"FTUInt16",
		"FTUInt16Null",
		"FTUInt8",
		"FTUInt8Null",
		"FTArrayInt64",
		"FTArrayInt64Null",
		"FTArrayInt32",
		"FTArrayInt32Null",
		"FTArrayInt16",
		"FTArrayInt16Null",
		"FTArrayInt8",
		"FTArrayInt8Null",
		"FTArrayUInt64",
		"FTArrayUInt64Null",
		"FTArrayUInt32",
		"FTArrayUInt32Null",
		"FTArrayUInt16",
		"FTArrayUInt16Null",
		"FTArrayUInt8",
		"FTArrayUInt8Null",
	}
}

func (cpt *CustomPrimitiveType) columnsToPtrs(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "PKey":
			ret = append(ret, yoDecode(&cpt.PKey))
		case "FTInt64":
			ret = append(ret, yoDecode(&cpt.FTInt64))
		case "FTInt64Null":
			ret = append(ret, yoDecode(&cpt.FTInt64null))
		case "FTInt32":
			ret = append(ret, yoDecode(&cpt.FTInt32))
		case "FTInt32Null":
			ret = append(ret, yoDecode(&cpt.FTInt32null))
		case "FTInt16":
			ret = append(ret, yoDecode(&cpt.FTInt16))
		case "FTInt16Null":
			ret = append(ret, yoDecode(&cpt.FTInt16null))
		case "FTInt8":
			ret = append(ret, yoDecode(&cpt.FTInt8))
		case "FTInt8Null":
			ret = append(ret, yoDecode(&cpt.FTInt8null))
		case "FTUInt64":
			ret = append(ret, yoDecode(&cpt.FTUInt64))
		case "FTUInt64Null":
			ret = append(ret, yoDecode(&cpt.FTUInt64null))
		case "FTUInt32":
			ret = append(ret, yoDecode(&cpt.FTUInt32))
		case "FTUInt32Null":
			ret = append(ret, yoDecode(&cpt.FTUInt32null))
		case "FTUInt16":
			ret = append(ret, yoDecode(&cpt.FTUInt16))
		case "FTUInt16Null":
			ret = append(ret, yoDecode(&cpt.FTUInt16null))
		case "FTUInt8":
			ret = append(ret, yoDecode(&cpt.FTUInt8))
		case "FTUInt8Null":
			ret = append(ret, yoDecode(&cpt.FTUInt8null))
		case "FTArrayInt64":
			ret = append(ret, yoDecode(&cpt.FTArrayInt64))
		case "FTArrayInt64Null":
			ret = append(ret, yoDecode(&cpt.FTArrayInt64null))
		case "FTArrayInt32":
			ret = append(ret, yoDecode(&cpt.FTArrayInt32))
		case "FTArrayInt32Null":
			ret = append(ret, yoDecode(&cpt.FTArrayInt32null))
		case "FTArrayInt16":
			ret = append(ret, yoDecode(&cpt.FTArrayInt16))
		case "FTArrayInt16Null":
			ret = append(ret, yoDecode(&cpt.FTArrayInt16null))
		case "FTArrayInt8":
			ret = append(ret, yoDecode(&cpt.FTArrayInt8))
		case "FTArrayInt8Null":
			ret = append(ret, yoDecode(&cpt.FTArrayInt8null))
		case "FTArrayUInt64":
			ret = append(ret, yoDecode(&cpt.FTArrayUINt64))
		case "FTArrayUInt64Null":
			ret = append(ret, yoDecode(&cpt.FTArrayUINt64null))
		case "FTArrayUInt32":
			ret = append(ret, yoDecode(&cpt.FTArrayUINt32))
		case "FTArrayUInt32Null":
			ret = append(ret, yoDecode(&cpt.FTArrayUINt32null))
		case "FTArrayUInt16":
			ret = append(ret, yoDecode(&cpt.FTArrayUINt16))
		case "FTArrayUInt16Null":
			ret = append(ret, yoDecode(&cpt.FTArrayUINt16null))
		case "FTArrayUInt8":
			ret = append(ret, yoDecode(&cpt.FTArrayUINt8))
		case "FTArrayUInt8Null":
			ret = append(ret, yoDecode(&cpt.FTArrayUINt8null))
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}
	return ret, nil
}

func (cpt *CustomPrimitiveType) columnsToValues(cols []string) ([]interface{}, error) {
	ret := make([]interface{}, 0, len(cols))
	for _, col := range cols {
		switch col {
		case "PKey":
			ret = append(ret, yoEncode(cpt.PKey))
		case "FTInt64":
			ret = append(ret, yoEncode(cpt.FTInt64))
		case "FTInt64Null":
			ret = append(ret, yoEncode(cpt.FTInt64null))
		case "FTInt32":
			ret = append(ret, yoEncode(cpt.FTInt32))
		case "FTInt32Null":
			ret = append(ret, yoEncode(cpt.FTInt32null))
		case "FTInt16":
			ret = append(ret, yoEncode(cpt.FTInt16))
		case "FTInt16Null":
			ret = append(ret, yoEncode(cpt.FTInt16null))
		case "FTInt8":
			ret = append(ret, yoEncode(cpt.FTInt8))
		case "FTInt8Null":
			ret = append(ret, yoEncode(cpt.FTInt8null))
		case "FTUInt64":
			ret = append(ret, yoEncode(cpt.FTUInt64))
		case "FTUInt64Null":
			ret = append(ret, yoEncode(cpt.FTUInt64null))
		case "FTUInt32":
			ret = append(ret, yoEncode(cpt.FTUInt32))
		case "FTUInt32Null":
			ret = append(ret, yoEncode(cpt.FTUInt32null))
		case "FTUInt16":
			ret = append(ret, yoEncode(cpt.FTUInt16))
		case "FTUInt16Null":
			ret = append(ret, yoEncode(cpt.FTUInt16null))
		case "FTUInt8":
			ret = append(ret, yoEncode(cpt.FTUInt8))
		case "FTUInt8Null":
			ret = append(ret, yoEncode(cpt.FTUInt8null))
		case "FTArrayInt64":
			ret = append(ret, yoEncode(cpt.FTArrayInt64))
		case "FTArrayInt64Null":
			ret = append(ret, yoEncode(cpt.FTArrayInt64null))
		case "FTArrayInt32":
			ret = append(ret, yoEncode(cpt.FTArrayInt32))
		case "FTArrayInt32Null":
			ret = append(ret, yoEncode(cpt.FTArrayInt32null))
		case "FTArrayInt16":
			ret = append(ret, yoEncode(cpt.FTArrayInt16))
		case "FTArrayInt16Null":
			ret = append(ret, yoEncode(cpt.FTArrayInt16null))
		case "FTArrayInt8":
			ret = append(ret, yoEncode(cpt.FTArrayInt8))
		case "FTArrayInt8Null":
			ret = append(ret, yoEncode(cpt.FTArrayInt8null))
		case "FTArrayUInt64":
			ret = append(ret, yoEncode(cpt.FTArrayUINt64))
		case "FTArrayUInt64Null":
			ret = append(ret, yoEncode(cpt.FTArrayUINt64null))
		case "FTArrayUInt32":
			ret = append(ret, yoEncode(cpt.FTArrayUINt32))
		case "FTArrayUInt32Null":
			ret = append(ret, yoEncode(cpt.FTArrayUINt32null))
		case "FTArrayUInt16":
			ret = append(ret, yoEncode(cpt.FTArrayUINt16))
		case "FTArrayUInt16Null":
			ret = append(ret, yoEncode(cpt.FTArrayUINt16null))
		case "FTArrayUInt8":
			ret = append(ret, yoEncode(cpt.FTArrayUINt8))
		case "FTArrayUInt8Null":
			ret = append(ret, yoEncode(cpt.FTArrayUINt8null))
		default:
			return nil, fmt.Errorf("unknown column: %s", col)
		}
	}

	return ret, nil
}

// newCustomPrimitiveType_Decoder returns a decoder which reads a row from *spanner.Row
// into CustomPrimitiveType. The decoder is not goroutine-safe. Don't use it concurrently.
func newCustomPrimitiveType_Decoder(cols []string) func(*spanner.Row) (*CustomPrimitiveType, error) {
	return func(row *spanner.Row) (*CustomPrimitiveType, error) {
		var cpt CustomPrimitiveType
		ptrs, err := cpt.columnsToPtrs(cols)
		if err != nil {
			return nil, err
		}

		if err := row.Columns(ptrs...); err != nil {
			return nil, err
		}

		return &cpt, nil
	}
}

// Insert returns a Mutation to insert a row into a table. If the row already
// exists, the write or transaction fails.
func (cpt *CustomPrimitiveType) Insert(ctx context.Context) *spanner.Mutation {
	values, _ := cpt.columnsToValues(CustomPrimitiveTypeWritableColumns())
	return spanner.Insert("CustomPrimitiveTypes", CustomPrimitiveTypeWritableColumns(), values)
}

// Update returns a Mutation to update a row in a table. If the row does not
// already exist, the write or transaction fails.
func (cpt *CustomPrimitiveType) Update(ctx context.Context) *spanner.Mutation {
	values, _ := cpt.columnsToValues(CustomPrimitiveTypeWritableColumns())
	return spanner.Update("CustomPrimitiveTypes", CustomPrimitiveTypeWritableColumns(), values)
}

// InsertOrUpdate returns a Mutation to insert a row into a table. If the row
// already exists, it updates it instead. Any column values not explicitly
// written are preserved.
func (cpt *CustomPrimitiveType) InsertOrUpdate(ctx context.Context) *spanner.Mutation {
	values, _ := cpt.columnsToValues(CustomPrimitiveTypeWritableColumns())
	return spanner.InsertOrUpdate("CustomPrimitiveTypes", CustomPrimitiveTypeWritableColumns(), values)
}

// Replace returns a Mutation to insert a row into a table, deleting any
// existing row. Unlike InsertOrUpdate, this means any values not explicitly
// written become NULL.
func (cpt *CustomPrimitiveType) Replace(ctx context.Context) *spanner.Mutation {
	values, _ := cpt.columnsToValues(CustomPrimitiveTypeWritableColumns())
	return spanner.Replace("CustomPrimitiveTypes", CustomPrimitiveTypeWritableColumns(), values)
}

// UpdateColumns returns a Mutation to update specified columns of a row in a table.
func (cpt *CustomPrimitiveType) UpdateColumns(ctx context.Context, cols ...string) (*spanner.Mutation, error) {
	// add primary keys to columns to update by primary keys
	colsWithPKeys := append(cols, CustomPrimitiveTypePrimaryKeys()...)

	values, err := cpt.columnsToValues(colsWithPKeys)
	if err != nil {
		return nil, newErrorWithCode(codes.InvalidArgument, "CustomPrimitiveType.UpdateColumns", "CustomPrimitiveTypes", err)
	}

	return spanner.Update("CustomPrimitiveTypes", colsWithPKeys, values), nil
}

// FindCustomPrimitiveType gets a CustomPrimitiveType by primary key
func FindCustomPrimitiveType(ctx context.Context, db YODB, pKey string) (*CustomPrimitiveType, error) {
	key := spanner.Key{yoEncode(pKey)}
	row, err := db.ReadRow(ctx, "CustomPrimitiveTypes", key, CustomPrimitiveTypeColumns())
	if err != nil {
		return nil, newError("FindCustomPrimitiveType", "CustomPrimitiveTypes", err)
	}

	decoder := newCustomPrimitiveType_Decoder(CustomPrimitiveTypeColumns())
	cpt, err := decoder(row)
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "FindCustomPrimitiveType", "CustomPrimitiveTypes", err)
	}

	return cpt, nil
}

// ReadCustomPrimitiveType retrieves multiples rows from CustomPrimitiveType by KeySet as a slice.
func ReadCustomPrimitiveType(ctx context.Context, db YODB, keys spanner.KeySet) ([]*CustomPrimitiveType, error) {
	var res []*CustomPrimitiveType

	decoder := newCustomPrimitiveType_Decoder(CustomPrimitiveTypeColumns())

	rows := db.Read(ctx, "CustomPrimitiveTypes", keys, CustomPrimitiveTypeColumns())
	err := rows.Do(func(row *spanner.Row) error {
		cpt, err := decoder(row)
		if err != nil {
			return err
		}
		res = append(res, cpt)

		return nil
	})
	if err != nil {
		return nil, newErrorWithCode(codes.Internal, "ReadCustomPrimitiveType", "CustomPrimitiveTypes", err)
	}

	return res, nil
}

// Delete deletes the CustomPrimitiveType from the database.
func (cpt *CustomPrimitiveType) Delete(ctx context.Context) *spanner.Mutation {
	values, _ := cpt.columnsToValues(CustomPrimitiveTypePrimaryKeys())
	return spanner.Delete("CustomPrimitiveTypes", spanner.Key(values))
}
