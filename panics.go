// Copyright 2011 Aaron Jacobs. All Rights Reserved.
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
	"errors"
	"reflect"
)

// Panics matches zero-arg functions which, when invoked, panic with an error
// that matches the supplied matcher.
func Panics(m Matcher) Matcher {
	return &panicsMatcher{m}
}

type panicsMatcher struct {
	wrappedMatcher Matcher
}

func (m *panicsMatcher) Description() string {
	return "panics with: " + m.wrappedMatcher.Description()
}

func (m *panicsMatcher) Matches(c interface{}) (res MatchResult, err error) {
	// Make sure c is a zero-arg function.
	v := reflect.ValueOf(c)
	if v.Kind() != reflect.Func || v.Type().NumIn() != 0 {
		res = MATCH_UNDEFINED
		err = errors.New("which is not a zero-arg function")
		return
	}

	// Call the function and check its panic error.
	defer func() {
		if e := recover(); e != nil {
			res, err = t.wrappedMatcher.Matches(e)
		}
	}()

	v.Call([]Value{})

	// If we got here, the function didn't panic.
	res = MATCH_FALSE
	err = nil
	return
}
