// Copyright 2019 Authors of Hubble
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

package logger

import (
	"io/ioutil"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	log  *logrus.Entry
	once sync.Once
)

// NoopLogger ...
func NoopLogger() *logrus.Entry {
	l := logrus.New()
	l.Out = ioutil.Discard
	e := logrus.NewEntry(l)
	return e
}

// GetLogger returns the logger properly set up accordingly with the debug flag.
func GetLogger() *logrus.Entry {
	once.Do(func() {
		logger := logrus.New()
		if viper.GetBool("debug") {
			logger.SetLevel(logrus.DebugLevel)
		}
		log = logrus.NewEntry(logger)
	})
	return log
}
