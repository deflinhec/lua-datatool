// Copyright 2023 Deflinhec
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

package formatter

// An Option is used to define the behavior and rules of a Format.
type Option func(*options)

type KeySortType int8

const (
	KeySortNone KeySortType = iota
	KeySortHashFnv64
	KeySortHashFnv64Reversed
	KeySortAlphabetical
	KeySortAlphabeticalReversed
)

type style struct {
	// MaxLineLength is the maximum length of a line.
	maxLineLength int

	// KeySort is the type of key sorting.
	keySort KeySortType

	// KeySortPriority is a list of keys that will be sorted first.
	keySortPriority []string

	// StepIndentWidth is the width of the indentation for each step.
	indentWidhtStep int
}

type options struct {
	style

	// RootIndentWidth is the width of the indentation.
	rootIndentWidth int

	// RootTrait is the trait of the root element.
	rootTrait string
}

// Trait sets the trait of the root element.
func RootTrait(trait string) Option {
	return func(o *options) {
		o.rootTrait = trait
	}
}

// RootIndentWidth sets the width of the indentation.
func RootIndentWidth(width int) Option {
	return func(o *options) {
		o.rootIndentWidth = width
	}
}

// IndentWidthStep sets the width of the indentation for each step.
func IndentWidthStep(width int) Option {
	return func(o *options) {
		o.indentWidhtStep = width
	}
}

// MaxLineLength sets the maximum length of a line.
func MaxLineLength(length int) Option {
	return func(o *options) {
		o.maxLineLength = length
	}
}

// KeySort sets the key sort type.
func KeySort(sort KeySortType) Option {
	return func(o *options) {
		o.keySort = sort
	}
}

// KeySortPriority sets the key sort priority.
func KeySortPriority(priority ...string) Option {
	return func(o *options) {
		o.keySortPriority = priority
	}
}
