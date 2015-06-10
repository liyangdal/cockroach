// Copyright 2015 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Peter Mattis (peter@cockroachlabs.com)

package client

// This file contains the experimental Cockroach table-based interface. The
// contents will eventually be dispersed into {batch,db,txn}.go, but are
// collected here during initial development.

type model struct {
	fields fieldMap
}

// BindModel ...
func (db *DB) BindModel(name string, obj interface{}) error {
	return nil
}

// GetStruct ...
func (db *DB) GetStruct(obj interface{}, columns ...string) error {
	b := db.NewBatch()
	b.GetStruct(obj, columns...)
	_, err := runOneResult(db, b)
	return err
}

// PutStruct ...
func (db *DB) PutStruct(obj interface{}, columns ...string) error {
	b := db.NewBatch()
	b.PutStruct(obj, columns...)
	_, err := runOneResult(db, b)
	return err
}

// IncStruct ...
func (db *DB) IncStruct(obj interface{}, value int64, column string) error {
	b := db.NewBatch()
	b.IncStruct(obj, value, column)
	_, err := runOneResult(db, b)
	return err
}

// ScanStruct ...
func (db *DB) ScanStruct(start, end interface{}, maxRows int64) error {
	b := db.NewBatch()
	b.ScanStruct(start, end, maxRows)
	_, err := runOneResult(db, b)
	return err
}

// DelStruct ...
func (db *DB) DelStruct(obj interface{}, columns ...string) error {
	b := db.NewBatch()
	b.DelStruct(obj, columns...)
	_, err := runOneResult(db, b)
	return err
}

// GetStruct ...
func (txn *Txn) GetStruct(obj interface{}, columns ...string) error {
	b := txn.NewBatch()
	b.GetStruct(obj, columns...)
	_, err := runOneResult(txn, b)
	return err
}

// PutStruct ...
func (txn *Txn) PutStruct(obj interface{}, columns ...string) error {
	b := txn.NewBatch()
	b.PutStruct(obj, columns...)
	_, err := runOneResult(txn, b)
	return err
}

// IncStruct ...
func (txn *Txn) IncStruct(obj interface{}, value int64, column string) error {
	b := txn.NewBatch()
	b.IncStruct(obj, value, column)
	_, err := runOneResult(txn, b)
	return err
}

// ScanStruct ...
func (txn *Txn) ScanStruct(start, end interface{}, maxRows int64) error {
	b := txn.NewBatch()
	b.ScanStruct(start, end, maxRows)
	_, err := runOneResult(txn, b)
	return err
}

// DelStruct ...
func (txn *Txn) DelStruct(obj interface{}, columns ...string) error {
	b := txn.NewBatch()
	b.DelStruct(obj, columns...)
	_, err := runOneResult(txn, b)
	return err
}

// GetStruct ...
func (b *Batch) GetStruct(obj interface{}, columns ...string) {
}

// PutStruct ...
func (b *Batch) PutStruct(obj interface{}, columns ...string) {
}

// IncStruct ...
func (b *Batch) IncStruct(obj interface{}, value int64, column string) {
}

// ScanStruct ...
func (b *Batch) ScanStruct(start, end interface{}, maxRows int64) {
}

// DelStruct ...
func (b *Batch) DelStruct(obj interface{}, columns ...string) {
}
