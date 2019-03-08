// Code generated by execgen; DO NOT EDIT.
// Copyright 2018 The Cockroach Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package exec

import (
	"fmt"

	"github.com/cockroachdb/apd"
	"github.com/cockroachdb/cockroach/pkg/sql/exec/types"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
)

// EncDatumRowsToColVec converts one column from EncDatumRows to a column
// vector. columnIdx is the 0-based index of the column in the EncDatumRows.
func EncDatumRowsToColVec(
	rows sqlbase.EncDatumRows,
	vec ColVec,
	columnIdx int,
	columnType *sqlbase.ColumnType,
	alloc *sqlbase.DatumAlloc,
) error {

	switch columnType.SemanticType {
	case sqlbase.ColumnType_STRING:

		nRows := uint16(len(rows))
		col := vec.Bytes()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.([]byte)
			}
		}
	case sqlbase.ColumnType_DATE:

		nRows := uint16(len(rows))
		col := vec.Int64()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(int64)
			}
		}
	case sqlbase.ColumnType_BOOL:

		nRows := uint16(len(rows))
		col := vec.Bool()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(bool)
			}
		}
	case sqlbase.ColumnType_INT:
		switch columnType.Width {
		case 0:

			nRows := uint16(len(rows))
			col := vec.Int64()
			datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
			for i := uint16(0); i < nRows; i++ {
				if rows[i][columnIdx].Datum == nil {
					if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := rows[i][columnIdx].Datum
				if datum == tree.DNull {
					vec.SetNull(i)
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}
					col[i] = v.(int64)
				}
			}
		case 8:

			nRows := uint16(len(rows))
			col := vec.Int8()
			datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
			for i := uint16(0); i < nRows; i++ {
				if rows[i][columnIdx].Datum == nil {
					if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := rows[i][columnIdx].Datum
				if datum == tree.DNull {
					vec.SetNull(i)
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}
					col[i] = v.(int8)
				}
			}
		case 16:

			nRows := uint16(len(rows))
			col := vec.Int16()
			datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
			for i := uint16(0); i < nRows; i++ {
				if rows[i][columnIdx].Datum == nil {
					if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := rows[i][columnIdx].Datum
				if datum == tree.DNull {
					vec.SetNull(i)
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}
					col[i] = v.(int16)
				}
			}
		case 32:

			nRows := uint16(len(rows))
			col := vec.Int32()
			datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
			for i := uint16(0); i < nRows; i++ {
				if rows[i][columnIdx].Datum == nil {
					if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := rows[i][columnIdx].Datum
				if datum == tree.DNull {
					vec.SetNull(i)
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}
					col[i] = v.(int32)
				}
			}
		case 64:

			nRows := uint16(len(rows))
			col := vec.Int64()
			datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
			for i := uint16(0); i < nRows; i++ {
				if rows[i][columnIdx].Datum == nil {
					if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
						return err
					}
				}
				datum := rows[i][columnIdx].Datum
				if datum == tree.DNull {
					vec.SetNull(i)
				} else {
					v, err := datumToPhysicalFn(datum)
					if err != nil {
						return err
					}
					col[i] = v.(int64)
				}
			}
		default:
			panic(fmt.Sprintf("unsupported width %d for column type %s", columnType.Width, columnType.SQLString()))
		}
	case sqlbase.ColumnType_FLOAT:

		nRows := uint16(len(rows))
		col := vec.Float64()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(float64)
			}
		}
	case sqlbase.ColumnType_DECIMAL:

		nRows := uint16(len(rows))
		col := vec.Decimal()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(apd.Decimal)
			}
		}
	case sqlbase.ColumnType_BYTES:

		nRows := uint16(len(rows))
		col := vec.Bytes()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.([]byte)
			}
		}
	case sqlbase.ColumnType_OID:

		nRows := uint16(len(rows))
		col := vec.Int64()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.(int64)
			}
		}
	case sqlbase.ColumnType_NAME:

		nRows := uint16(len(rows))
		col := vec.Bytes()
		datumToPhysicalFn := types.GetDatumToPhysicalFn(*columnType)
		for i := uint16(0); i < nRows; i++ {
			if rows[i][columnIdx].Datum == nil {
				if err := rows[i][columnIdx].EnsureDecoded(columnType, alloc); err != nil {
					return err
				}
			}
			datum := rows[i][columnIdx].Datum
			if datum == tree.DNull {
				vec.SetNull(i)
			} else {
				v, err := datumToPhysicalFn(datum)
				if err != nil {
					return err
				}
				col[i] = v.([]byte)
			}
		}
	default:
		panic(fmt.Sprintf("unsupported column type %s", columnType.SQLString()))
	}
	return nil
}
