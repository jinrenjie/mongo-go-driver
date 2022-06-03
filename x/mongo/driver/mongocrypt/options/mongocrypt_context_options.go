// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package options

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

// DataKeyOptions specifies options for creating a new data key.
type DataKeyOptions struct {
	KeyAltNames []string
	MasterKey   bsoncore.Document
	KeyMaterial *string
}

// DataKey creates a new DataKeyOptions instance.
func DataKey() *DataKeyOptions {
	return &DataKeyOptions{}
}

// SetKeyAltNames specifies alternate key names.
func (dko *DataKeyOptions) SetKeyAltNames(names []string) *DataKeyOptions {
	dko.KeyAltNames = names
	return dko
}

// SetMasterKey specifies the master key.
func (dko *DataKeyOptions) SetMasterKey(key bsoncore.Document) *DataKeyOptions {
	dko.MasterKey = key
	return dko
}

// SetKeyMaterial specifies the key material.
func (dko *DataKeyOptions) SetKeyMaterial(key string) *DataKeyOptions {
	dko.KeyMaterial = &key
	return dko
}

// QueryType describes the type of query the result of Encrypt is used for.
type QueryType int

// These constants specify valid values for QueryType
const (
	QueryTypeEquality QueryType = 1
)

// ExplicitEncryptionOptions specifies options for configuring an explicit encryption context.
type ExplicitEncryptionOptions struct {
	KeyID            *primitive.Binary
	KeyAltName       *string
	Algorithm        string
	QueryType        *QueryType
	ContentionFactor *int64
}

// ExplicitEncryption creates a new ExplicitEncryptionOptions instance.
func ExplicitEncryption() *ExplicitEncryptionOptions {
	return &ExplicitEncryptionOptions{}
}

// SetKeyID sets the key identifier.
func (eeo *ExplicitEncryptionOptions) SetKeyID(keyID primitive.Binary) *ExplicitEncryptionOptions {
	eeo.KeyID = &keyID
	return eeo
}

// SetKeyAltName sets the key alternative name.
func (eeo *ExplicitEncryptionOptions) SetKeyAltName(keyAltName string) *ExplicitEncryptionOptions {
	eeo.KeyAltName = &keyAltName
	return eeo
}

// SetAlgorithm specifies an encryption algorithm.
func (eeo *ExplicitEncryptionOptions) SetAlgorithm(algorithm string) *ExplicitEncryptionOptions {
	eeo.Algorithm = algorithm
	return eeo
}

// SetQueryType specifies the query type.
func (eeo *ExplicitEncryptionOptions) SetQueryType(queryType QueryType) *ExplicitEncryptionOptions {
	eeo.QueryType = &queryType
	return eeo
}

// SetContentionFactor specifies the contention factor.
func (eeo *ExplicitEncryptionOptions) SetContentionFactor(contentionFactor int64) *ExplicitEncryptionOptions {
	eeo.ContentionFactor = &contentionFactor
	return eeo
}
