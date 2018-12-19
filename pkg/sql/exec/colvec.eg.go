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
)

func (m *memColumn) Append(vec ColVec, colType types.T, toLength uint64, fromLength uint16) {
	switch colType {
	case types.Bool:
		m.col = append(m.Bool()[:toLength], vec.Bool()[:fromLength]...)
	case types.Bytes:
		m.col = append(m.Bytes()[:toLength], vec.Bytes()[:fromLength]...)
	case types.Decimal:
		m.col = append(m.Decimal()[:toLength], vec.Decimal()[:fromLength]...)
	case types.Int8:
		m.col = append(m.Int8()[:toLength], vec.Int8()[:fromLength]...)
	case types.Int16:
		m.col = append(m.Int16()[:toLength], vec.Int16()[:fromLength]...)
	case types.Int32:
		m.col = append(m.Int32()[:toLength], vec.Int32()[:fromLength]...)
	case types.Int64:
		m.col = append(m.Int64()[:toLength], vec.Int64()[:fromLength]...)
	case types.Float32:
		m.col = append(m.Float32()[:toLength], vec.Float32()[:fromLength]...)
	case types.Float64:
		m.col = append(m.Float64()[:toLength], vec.Float64()[:fromLength]...)
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}

	if fromLength > 0 {
		m.nulls = append(m.nulls, make([]int64, (fromLength-1)>>6+1)...)

		if vec.HasNulls() {
			for i := uint16(0); i < fromLength; i++ {
				if vec.NullAt(i) {
					m.SetNull64(toLength + uint64(i))
				}
			}
		}
	}
}

func (m *memColumn) AppendWithSel(
	vec ColVec, sel []uint16, batchSize uint16, colType types.T, toLength uint64,
) {
	switch colType {
	case types.Bool:
		toCol := append(m.Bool()[:toLength], make([]bool, batchSize)...)
		fromCol := vec.Bool()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Bytes:
		toCol := append(m.Bytes()[:toLength], make([][]byte, batchSize)...)
		fromCol := vec.Bytes()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Decimal:
		toCol := append(m.Decimal()[:toLength], make([]apd.Decimal, batchSize)...)
		fromCol := vec.Decimal()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Int8:
		toCol := append(m.Int8()[:toLength], make([]int8, batchSize)...)
		fromCol := vec.Int8()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Int16:
		toCol := append(m.Int16()[:toLength], make([]int16, batchSize)...)
		fromCol := vec.Int16()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Int32:
		toCol := append(m.Int32()[:toLength], make([]int32, batchSize)...)
		fromCol := vec.Int32()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Int64:
		toCol := append(m.Int64()[:toLength], make([]int64, batchSize)...)
		fromCol := vec.Int64()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Float32:
		toCol := append(m.Float32()[:toLength], make([]float32, batchSize)...)
		fromCol := vec.Float32()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	case types.Float64:
		toCol := append(m.Float64()[:toLength], make([]float64, batchSize)...)
		fromCol := vec.Float64()

		for i := uint16(0); i < batchSize; i++ {
			toCol[uint64(i)+toLength] = fromCol[sel[i]]
		}

		m.col = toCol
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}

	if batchSize > 0 {
		m.nulls = append(m.nulls, make([]int64, (batchSize-1)>>6+1)...)
		for i := uint16(0); i < batchSize; i++ {
			if vec.NullAt(sel[i]) {
				m.SetNull64(toLength + uint64(i))
			}
		}
	}
}

func (m *memColumn) Copy(src ColVec, srcStartIdx, srcEndIdx uint64, typ types.T) {
	switch typ {
	case types.Bool:
		copy(m.Bool(), src.Bool()[srcStartIdx:srcEndIdx])
	case types.Bytes:
		copy(m.Bytes(), src.Bytes()[srcStartIdx:srcEndIdx])
	case types.Decimal:
		copy(m.Decimal(), src.Decimal()[srcStartIdx:srcEndIdx])
	case types.Int8:
		copy(m.Int8(), src.Int8()[srcStartIdx:srcEndIdx])
	case types.Int16:
		copy(m.Int16(), src.Int16()[srcStartIdx:srcEndIdx])
	case types.Int32:
		copy(m.Int32(), src.Int32()[srcStartIdx:srcEndIdx])
	case types.Int64:
		copy(m.Int64(), src.Int64()[srcStartIdx:srcEndIdx])
	case types.Float32:
		copy(m.Float32(), src.Float32()[srcStartIdx:srcEndIdx])
	case types.Float64:
		copy(m.Float64(), src.Float64()[srcStartIdx:srcEndIdx])
	default:
		panic(fmt.Sprintf("unhandled type %d", typ))
	}
}

func (m *memColumn) CopyWithSelInt64(vec ColVec, sel []uint64, nSel uint16, colType types.T) {
	m.UnsetNulls()

	// todo (changangela): handle the case when nSel > ColBatchSize
	switch colType {
	case types.Bool:
		toCol := m.Bool()
		fromCol := vec.Bool()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Bytes:
		toCol := m.Bytes()
		fromCol := vec.Bytes()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Decimal:
		toCol := m.Decimal()
		fromCol := vec.Decimal()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int8:
		toCol := m.Int8()
		fromCol := vec.Int8()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int16:
		toCol := m.Int16()
		fromCol := vec.Int16()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int32:
		toCol := m.Int32()
		fromCol := vec.Int32()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int64:
		toCol := m.Int64()
		fromCol := vec.Int64()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Float32:
		toCol := m.Float32()
		fromCol := vec.Float32()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Float64:
		toCol := m.Float64()
		fromCol := vec.Float64()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt64(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

func (m *memColumn) CopyWithSelInt16(vec ColVec, sel []uint16, nSel uint16, colType types.T) {
	m.UnsetNulls()

	switch colType {
	case types.Bool:
		toCol := m.Bool()
		fromCol := vec.Bool()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Bytes:
		toCol := m.Bytes()
		fromCol := vec.Bytes()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Decimal:
		toCol := m.Decimal()
		fromCol := vec.Decimal()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int8:
		toCol := m.Int8()
		fromCol := vec.Int8()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int16:
		toCol := m.Int16()
		fromCol := vec.Int16()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int32:
		toCol := m.Int32()
		fromCol := vec.Int32()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Int64:
		toCol := m.Int64()
		fromCol := vec.Int64()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Float32:
		toCol := m.Float32()
		fromCol := vec.Float32()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	case types.Float64:
		toCol := m.Float64()
		fromCol := vec.Float64()

		if vec.HasNulls() {
			for i := uint16(0); i < nSel; i++ {
				if vec.NullAt(sel[i]) {
					m.SetNull(i)
				} else {
					toCol[i] = fromCol[sel[i]]
				}
			}
		} else {
			for i := uint16(0); i < nSel; i++ {
				toCol[i] = fromCol[sel[i]]
			}
		}
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}

func (m *memColumn) PrettyValueAt(colIdx uint16, colType types.T) string {
	if m.NullAt(colIdx) {
		return "NULL"
	}
	switch colType {
	case types.Bool:
		return fmt.Sprintf("%v", m.Bool()[colIdx])
	case types.Bytes:
		return fmt.Sprintf("%v", m.Bytes()[colIdx])
	case types.Decimal:
		return fmt.Sprintf("%v", m.Decimal()[colIdx])
	case types.Int8:
		return fmt.Sprintf("%v", m.Int8()[colIdx])
	case types.Int16:
		return fmt.Sprintf("%v", m.Int16()[colIdx])
	case types.Int32:
		return fmt.Sprintf("%v", m.Int32()[colIdx])
	case types.Int64:
		return fmt.Sprintf("%v", m.Int64()[colIdx])
	case types.Float32:
		return fmt.Sprintf("%v", m.Float32()[colIdx])
	case types.Float64:
		return fmt.Sprintf("%v", m.Float64()[colIdx])
	default:
		panic(fmt.Sprintf("unhandled type %d", colType))
	}
}
