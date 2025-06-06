// SPDX-License-Identifier: Apache-2.0

package version

import (
	"testing"

	"github.com/go-vela/server/version"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	const (
		goVersion   = "go1.20.7"
		osVersion   = "darwin"
		archVersion = "arm64"
	)

	t.Run("Valid Version Output", func(t *testing.T) {
		// to avoid flaky tests, set the values
		Go = goVersion
		OS = osVersion
		Compiler = "gc"
		Tag = "v1.1.1"
		Commit = "000"
		Date = "111"
		Arch = archVersion

		expected := version.Version{
			Canonical: "v1.1.1",
			Major:     1,
			Minor:     1,
			Patch:     1,
			Metadata: version.Metadata{
				Architecture:    "arm64",
				BuildDate:       "111",
				Compiler:        "gc",
				GitCommit:       "000",
				GoVersion:       "go1.20.7",
				OperatingSystem: "darwin",
			},
		}
		assert.Equal(t, expected, *New())
	})
	t.Run("no tag", func(t *testing.T) {
		// to avoid flaky tests, set the values
		Go = goVersion
		OS = osVersion
		Compiler = "gc"
		Tag = ""
		Commit = "000"
		Date = "111"
		Arch = archVersion

		expected := version.Version{
			Canonical: "v0.0.0",
			Metadata: version.Metadata{
				Architecture:    "arm64",
				BuildDate:       "111",
				Compiler:        "gc",
				GitCommit:       "000",
				GoVersion:       "go1.20.7",
				OperatingSystem: "darwin",
			},
		}
		assert.Equal(t, expected, *New())
	})
	t.Run("invalid tag", func(t *testing.T) {
		// to avoid flaky tests, set the values
		Go = goVersion
		OS = osVersion
		Compiler = "gc"
		Tag = "something"
		Commit = "000"
		Date = "111"
		Arch = archVersion

		expected := version.Version{
			Canonical: "something",
			Metadata: version.Metadata{
				Architecture:    "arm64",
				BuildDate:       "111",
				Compiler:        "gc",
				GitCommit:       "000",
				GoVersion:       "go1.20.7",
				OperatingSystem: "darwin",
			},
		}
		assert.Equal(t, expected, *New())
	})
}
