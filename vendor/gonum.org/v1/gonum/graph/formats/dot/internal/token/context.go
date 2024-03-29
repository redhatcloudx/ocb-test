// Code generated by gocc; DO NOT EDIT.

// This file is dual licensed under CC0 and The Gonum License.
//
// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Copyright ©2017 Robin Eklind.
// This file is made available under a Creative Commons CC0 1.0
// Universal Public Domain Dedication.

package token

// Context allows user-defined data to be associated with the
// lexer/scanner to be associated with each token that lexer
// produces.
type Context interface{}

// Sourcer is a Context interface which presents a Source() method
// identifying e.g the filename for the current code.
type Sourcer interface {
	Source() string
}
