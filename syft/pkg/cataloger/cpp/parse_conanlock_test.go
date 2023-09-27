package cpp

import (
	"testing"

	"github.com/anchore/syft/syft/artifact"
	"github.com/anchore/syft/syft/file"
	"github.com/anchore/syft/syft/pkg"
	"github.com/anchore/syft/syft/pkg/cataloger/internal/pkgtest"
)

func TestParseConanlock(t *testing.T) {
	fixture := "test-fixtures/conan.lock"
	expected := []pkg.Package{
		{
			Name:         "mfast",
			Version:      "1.2.2",
			PURL:         "pkg:conan/my_user/mfast@1.2.2?channel=my_channel",
			Locations:    file.NewLocationSet(file.NewLocation(fixture)),
			Language:     pkg.CPP,
			Type:         pkg.ConanPkg,
			MetadataType: pkg.ConanLockMetadataType,
			Metadata: pkg.ConanLockMetadata{
				Ref: "mfast/1.2.2@my_user/my_channel#c6f6387c9b99780f0ee05e25f99d0f39",
				Options: map[string]string{
					"fPIC":                            "True",
					"shared":                          "False",
					"with_sqlite3":                    "False",
					"boost:addr2line_location":        "/usr/bin/addr2line",
					"boost:asio_no_deprecated":        "False",
					"boost:buildid":                   "None",
					"boost:bzip2":                     "True",
					"boost:debug_level":               "0",
					"boost:diagnostic_definitions":    "False",
					"boost:error_code_header_only":    "False",
					"boost:extra_b2_flags":            "None",
					"boost:fPIC":                      "True",
					"boost:filesystem_no_deprecated":  "False",
					"boost:header_only":               "False",
					"boost:i18n_backend":              "deprecated",
					"boost:i18n_backend_iconv":        "libc",
					"boost:i18n_backend_icu":          "False",
					"boost:layout":                    "system",
					"boost:lzma":                      "False",
					"boost:magic_autolink":            "False",
					"boost:multithreading":            "True",
					"boost:namespace":                 "boost",
					"boost:namespace_alias":           "False",
					"boost:numa":                      "True",
					"boost:pch":                       "True",
					"boost:python_executable":         "None",
					"boost:python_version":            "None",
					"boost:segmented_stacks":          "False",
					"boost:shared":                    "False",
					"boost:system_no_deprecated":      "False",
					"boost:system_use_utf8":           "False",
					"boost:visibility":                "hidden",
					"boost:with_stacktrace_backtrace": "True",
					"boost:without_atomic":            "False",
					"boost:without_chrono":            "False",
					"boost:without_container":         "False",
					"boost:without_context":           "False",
					"boost:without_contract":          "False",
					"boost:without_coroutine":         "False",
					"boost:without_date_time":         "False",
					"boost:without_exception":         "False",
					"boost:without_fiber":             "False",
					"boost:without_filesystem":        "False",
					"boost:without_graph":             "False",
					"boost:without_graph_parallel":    "True",
					"boost:without_iostreams":         "False",
					"boost:without_json":              "False",
					"boost:without_locale":            "False",
					"boost:without_log":               "False",
					"boost:without_math":              "False",
					"boost:without_mpi":               "True",
					"boost:without_nowide":            "False",
					"boost:without_program_options":   "False",
					"boost:without_python":            "True",
					"boost:without_random":            "False",
					"boost:without_regex":             "False",
					"boost:without_serialization":     "False",
					"boost:without_stacktrace":        "False",
					"boost:without_system":            "False",
					"boost:without_test":              "False",
					"boost:without_thread":            "False",
					"boost:without_timer":             "False",
					"boost:without_type_erasure":      "False",
					"boost:without_wave":              "False",
					"boost:zlib":                      "True",
					"boost:zstd":                      "False",
					"bzip2:build_executable":          "True",
					"bzip2:fPIC":                      "True",
					"bzip2:shared":                    "False",
					"libbacktrace:fPIC":               "True",
					"libbacktrace:shared":             "False",
					"tinyxml2:fPIC":                   "True",
					"tinyxml2:shared":                 "False",
					"zlib:fPIC":                       "True",
					"zlib:shared":                     "False",
				},
				Context:   "host",
				PackageID: "9d1f076b471417647c2022a78d5e2c1f834289ac",
				Prev:      "0ca9799450422cc55a92ccc6ffd57fba",
			},
		},
		{
			Name:         "boost",
			Version:      "1.75.0",
			PURL:         "pkg:conan/boost@1.75.0",
			Locations:    file.NewLocationSet(file.NewLocation(fixture)),
			Language:     pkg.CPP,
			Type:         pkg.ConanPkg,
			MetadataType: pkg.ConanLockMetadataType,
			Metadata: pkg.ConanLockMetadata{
				Ref: "boost/1.75.0#a9c318f067216f900900e044e7af4ab1",
				Options: map[string]string{
					"addr2line_location":        "/usr/bin/addr2line",
					"asio_no_deprecated":        "False",
					"buildid":                   "None",
					"bzip2":                     "True",
					"debug_level":               "0",
					"diagnostic_definitions":    "False",
					"error_code_header_only":    "False",
					"extra_b2_flags":            "None",
					"fPIC":                      "True",
					"filesystem_no_deprecated":  "False",
					"header_only":               "False",
					"i18n_backend":              "deprecated",
					"i18n_backend_iconv":        "libc",
					"i18n_backend_icu":          "False",
					"layout":                    "system",
					"lzma":                      "False",
					"magic_autolink":            "False",
					"multithreading":            "True",
					"namespace":                 "boost",
					"namespace_alias":           "False",
					"numa":                      "True",
					"pch":                       "True",
					"python_executable":         "None",
					"python_version":            "None",
					"segmented_stacks":          "False",
					"shared":                    "False",
					"system_no_deprecated":      "False",
					"system_use_utf8":           "False",
					"visibility":                "hidden",
					"with_stacktrace_backtrace": "True",
					"without_atomic":            "False",
					"without_chrono":            "False",
					"without_container":         "False",
					"without_context":           "False",
					"without_contract":          "False",
					"without_coroutine":         "False",
					"without_date_time":         "False",
					"without_exception":         "False",
					"without_fiber":             "False",
					"without_filesystem":        "False",
					"without_graph":             "False",
					"without_graph_parallel":    "True",
					"without_iostreams":         "False",
					"without_json":              "False",
					"without_locale":            "False",
					"without_log":               "False",
					"without_math":              "False",
					"without_mpi":               "True",
					"without_nowide":            "False",
					"without_program_options":   "False",
					"without_python":            "True",
					"without_random":            "False",
					"without_regex":             "False",
					"without_serialization":     "False",
					"without_stacktrace":        "False",
					"without_system":            "False",
					"without_test":              "False",
					"without_thread":            "False",
					"without_timer":             "False",
					"without_type_erasure":      "False",
					"without_wave":              "False",
					"zlib":                      "True",
					"zstd":                      "False",
					"bzip2:build_executable":    "True",
					"bzip2:fPIC":                "True",
					"bzip2:shared":              "False",
					"libbacktrace:fPIC":         "True",
					"libbacktrace:shared":       "False",
					"zlib:fPIC":                 "True",
					"zlib:shared":               "False",
				},
				Context:   "host",
				PackageID: "dc8aedd23a0f0a773a5fcdcfe1ae3e89c4205978",
				Prev:      "b9d7912e6131dfa453c725593b36c808",
			},
		},
		{
			Name:         "zlib",
			Version:      "1.2.12",
			PURL:         "pkg:conan/zlib@1.2.12",
			Locations:    file.NewLocationSet(file.NewLocation(fixture)),
			Language:     pkg.CPP,
			Type:         pkg.ConanPkg,
			MetadataType: pkg.ConanLockMetadataType,
			Metadata: pkg.ConanLockMetadata{
				Ref: "zlib/1.2.12#c67ce17f2e96b972d42393ce50a76a1a",
				Options: map[string]string{
					"fPIC":   "True",
					"shared": "False",
				},
				Context:   "host",
				PackageID: "dfbe50feef7f3c6223a476cd5aeadb687084a646",
				Prev:      "7cd359d44f89ab08e33b5db75605002c",
			},
		},
		{
			Name:         "bzip2",
			Version:      "1.0.8",
			PURL:         "pkg:conan/bzip2@1.0.8",
			Locations:    file.NewLocationSet(file.NewLocation(fixture)),
			Language:     pkg.CPP,
			Type:         pkg.ConanPkg,
			MetadataType: pkg.ConanLockMetadataType,
			Metadata: pkg.ConanLockMetadata{
				Ref: "bzip2/1.0.8#62a8031289639043797cf53fa876d0ef",
				Options: map[string]string{
					"build_executable": "True",
					"fPIC":             "True",
					"shared":           "False",
				},
				Context:   "host",
				PackageID: "c32092bf4d4bb47cf962af898e02823f499b017e",
				Prev:      "b746948bc999d6f17f52a1f76e729e80",
			},
		},
		{
			Name:         "libbacktrace",
			Version:      "cci.20210118",
			PURL:         "pkg:conan/libbacktrace@cci.20210118",
			Locations:    file.NewLocationSet(file.NewLocation(fixture)),
			Language:     pkg.CPP,
			Type:         pkg.ConanPkg,
			MetadataType: pkg.ConanLockMetadataType,
			Metadata: pkg.ConanLockMetadata{
				Ref: "libbacktrace/cci.20210118#76e40b760e0bcd602d46db56b22820ab",
				Options: map[string]string{
					"fPIC":   "True",
					"shared": "False",
				},
				Context:   "host",
				PackageID: "dfbe50feef7f3c6223a476cd5aeadb687084a646",
				Prev:      "98a976f017e894c27e9a158b807ec0c7",
			},
		},
		{
			Name:         "tinyxml2",
			Version:      "9.0.0",
			PURL:         "pkg:conan/tinyxml2@9.0.0",
			Locations:    file.NewLocationSet(file.NewLocation(fixture)),
			Language:     pkg.CPP,
			Type:         pkg.ConanPkg,
			MetadataType: pkg.ConanLockMetadataType,
			Metadata: pkg.ConanLockMetadata{
				Ref: "tinyxml2/9.0.0#9f13a36ebfc222cd55fe531a0a8d94d1",
				Options: map[string]string{
					"fPIC":   "True",
					"shared": "False",
				},
				Context: "host",
				// intentionally remove to test missing PackageID and Prev
				// PackageID: "6557f18ca99c0b6a233f43db00e30efaa525e27e",
				// Prev:      "548bb273d2980991baa519453d68e5cd",
			},
		},
	}

	// relationships require IDs to be set to be sorted similarly
	for i := range expected {
		expected[i].SetID()
	}

	var expectedRelationships = []artifact.Relationship{
		{
			From: expected[1], // boost
			To:   expected[0], // mfast
			Type: artifact.DependencyOfRelationship,
			Data: nil,
		},
		{
			From: expected[5], // tinyxml2
			To:   expected[0], // mfast
			Type: artifact.DependencyOfRelationship,
			Data: nil,
		},
		{
			From: expected[2], // zlib
			To:   expected[1], // boost
			Type: artifact.DependencyOfRelationship,
			Data: nil,
		},
		{
			From: expected[3], // bzip2
			To:   expected[1], // boost
			Type: artifact.DependencyOfRelationship,
			Data: nil,
		},
		{
			From: expected[4], // libbacktrace
			To:   expected[1], // boost
			Type: artifact.DependencyOfRelationship,
			Data: nil,
		},
	}

	pkgtest.TestFileParser(t, fixture, parseConanlock, expected, expectedRelationships)
}