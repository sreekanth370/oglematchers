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

package oglematchers_test

import (
	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type ErrorTest struct {
	matcherCalled bool
	suppliedCandidate interface{}
	wrappedResult MatchResult
	wrappedError error

	matcher Matcher
}

func init() { RegisterTestSuite(&ErrorTest{}) }

func (t *ErrorTest) SetUp() {
	wrapped := &fakeMatcher{
		func(c interface{}) (MatchResult, error) {
			t.matcherCalled = true
			t.suppliedCandidate = c
			return t.wrappedResult, t.wrappedError
		},
		"is foo",
	}

	t.matcher = Error(wrapped)
}

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func (t *ErrorTest) Description() {
	ExpectThat(t.matcher.Description(), Equals("error that is foo"))
}

func (t *ErrorTest) CandidateIsNil() {
	res, err := t.matcher.Matches(nil)

	ExpectThat(t.matcherCalled, Equals(false))
	ExpectThat(res, Equals(MATCH_UNDEFINED))
	ExpectThat(err, Equals(nil))
}

func (t *ErrorTest) CandidateIsString() {
}

func (t *ErrorTest) CallsWrappedMatcher() {
}

func (t *ErrorTest) ReturnsWrappedMatcherResult() {
}
