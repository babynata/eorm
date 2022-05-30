// Copyright 2021 gotomicro
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package test 是用于辅助测试的包。仅限于内部使用
package test

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// SimpleStruct 包含所有 eorm 支持的类型
type SimpleStruct struct {
	Bool bool
	BoolPtr *bool

	Int int
	IntPtr *int

	Int8 int8
	Int8Ptr *int8

	Int16 int16
	Int16Ptr *int16

	Int32 int32
	Int32Ptr *int32

	Int64 int64
	Int64Ptr *int64

	Uint uint
	UintPtr *uint

	Uint8 uint8
	Uint8Ptr *uint8

	Uint16 uint16
	Uint16Ptr *uint16

	Uint32 uint32
	Uint32Ptr *uint32

	Uint64 uint64
	Uint64Ptr *uint64

	Float32 float32
	Float32Ptr *float32

	Float64 float64
	Float64Ptr *float64

	Byte byte
	BytePtr *byte
	ByteArray []byte

	String string

	// 特殊类型
	NullStringPtr *sql.NullString
	NullInt16Ptr *sql.NullInt16
	NullInt32Ptr *sql.NullInt32
	NullInt64Ptr *sql.NullInt64
	NullBoolPtr *sql.NullBool
	NullBytePtr *sql.NullByte
	NullTimePtr *sql.NullTime
	NullFloat64 *sql.NullFloat64
	JsonColumn *JsonColumn
}

// JsonColumn 是自定义的 JSON 类型字段
// Val 字段必须是结构体指针
type JsonColumn struct {
	Val any
	Valid bool
}

func (j *JsonColumn) Scan(src any) error {
	var bs []byte
	switch val := src.(type) {
	case string:
		bs = []byte(val)
	case []byte:
		bs = val
	case *[]byte:
		bs = *val
	default:
		return fmt.Errorf("不合法类型 %+v", src)
	}
	err := json.Unmarshal(bs, j.Val)
	if err != nil {
		return err
	}
	j.Valid = true
	return nil
}

// Value 参考 sql.NullXXX 类型定义的
func (j JsonColumn) Value() (driver.Value, error) {
	if !j.Valid {
		return nil, nil
	}
	return j.Val, nil
}