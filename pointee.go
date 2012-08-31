// Copyright 2012 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package oglematchers

import (
	"fmt"
	"reflect"
)

// Return a matcher that matches non-nil pointers whose pointee matches the
// wrapped matcher.
func Pointee(m Matcher) Matcher {
	return Not(m)
}

type pointeeMatcher struct {
	wrapped Matcher
}

func (m *pointeeMatcher) Matches(c interface{}) (err error) {
	// Make sure the candidate is of the appropriate type.
	cv := reflect.ValueOf(c)
	if !cv.IsValid() || cv.Kind() != reflect.Ptr {
		return NewFatalError("which is not a pointer")
	}

	// Make sure the candidate is non-nil.
	if cv.IsNil() {
		return fmt.Errorf("")
	}

	// Defer to the wrapped matcher.
	return m.wrapped.Matches(cv.Elem())
}

func (m *pointeeMatcher) Description() string {
	return fmt.Sprintf("pointee(%s)", m.wrapped.Description())
}
