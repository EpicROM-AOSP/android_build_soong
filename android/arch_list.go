// Copyright 2020 Google Inc. All rights reserved.
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

package android

var archVariants = map[ArchType][]string{
	Arm: {
		"armv7-a",
		"armv7-a-neon",
		"armv8-a",
		"armv8-2a",
	},
	Arm64: {
		"armv8-a",
		"armv8-a-branchprot",
		"armv8-2a",
		"armv8-2a-dotprod",
	},
	X86: {
		"amberlake",
		"atom",
		"broadwell",
		"haswell",
		"icelake",
		"ivybridge",
		"kabylake",
		"sandybridge",
		"silvermont",
		"skylake",
		"stoneyridge",
		"tigerlake",
		"whiskeylake",
		"x86_64",
	},
	X86_64: {
		"amberlake",
		"broadwell",
		"haswell",
		"icelake",
		"ivybridge",
		"kabylake",
		"sandybridge",
		"silvermont",
		"skylake",
		"stoneyridge",
		"tigerlake",
		"whiskeylake",
	},
}

var cpuVariants = map[ArchType][]string{
	Arm: {
		"cortex-a7",
		"cortex-a8",
		"cortex-a9",
		"cortex-a15",
		"cortex-a53",
		"cortex-a53.a57",
		"cortex-a55",
		"cortex-a72",
		"cortex-a73",
		"cortex-a75",
		"cortex-a76",
		"krait",
		"kryo",
		"kryo300",
		"kryo385",
		"exynos-m1",
		"exynos-m2",
	},
	Arm64: {
		"cortex-a53",
		"cortex-a55",
		"cortex-a72",
		"cortex-a73",
		"cortex-a75",
		"cortex-a76",
		"kryo",
		"kryo300",
		"kryo385",
		"exynos-m1",
		"exynos-m2",
	},
	X86:    {},
	X86_64: {},
}

var archFeatures = map[ArchType][]string{
	Arm: {
		"aarch32",
		"neon",
	},
	Arm64: {
		"dotprod",
	},
	X86: {
		"ssse3",
		"sse4",
		"sse4_1",
		"sse4_2",
		"aes_ni",
		"avx",
		"avx2",
		"avx512",
		"popcnt",
		"movbe",
	},
	X86_64: {
		"ssse3",
		"sse4",
		"sse4_1",
		"sse4_2",
		"aes_ni",
		"avx",
		"avx2",
		"avx512",
		"popcnt",
	},
}

var androidArchFeatureMap = map[ArchType]map[string][]string{
	Arm: {
		"armv7-a-neon": {
			"neon",
		},
		"armv8-a": {
			"aarch32",
			"neon",
		},
		"armv8-2a": {
			"aarch32",
			"neon",
		},
	},
	Arm64: {
		"armv8-2a-dotprod": {
			"dotprod",
		},
	},
	X86: {
		"amberlake": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"aes_ni",
			"popcnt",
		},
		"atom": {
			"ssse3",
			"movbe",
		},
		"broadwell": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"aes_ni",
			"popcnt",
		},
		"haswell": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"aes_ni",
			"avx",
			"popcnt",
			"movbe",
		},
		"icelake": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"avx512",
			"aes_ni",
			"popcnt",
		},
		"ivybridge": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"aes_ni",
			"avx",
			"popcnt",
		},
		"kabylake": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"aes_ni",
			"popcnt",
		},
		"sandybridge": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"popcnt",
		},
		"silvermont": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"aes_ni",
			"popcnt",
			"movbe",
		},
		"skylake": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"avx512",
			"aes_ni",
			"popcnt",
		},
		"stoneyridge": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"aes_ni",
			"avx",
			"avx2",
			"popcnt",
			"movbe",
		},
		"tigerlake": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"avx512",
			"aes_ni",
			"popcnt",
		},
		"whiskeylake": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"avx512",
			"aes_ni",
			"popcnt",
		},
		"x86_64": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"popcnt",
		},
	},
	X86_64: {
		"" /*default */ : {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"popcnt",
		},
		"amberlake": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"aes_ni",
			"popcnt",
		},
		"broadwell": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"aes_ni",
			"popcnt",
		},
		"haswell": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"aes_ni",
			"avx",
			"popcnt",
		},
		"icelake": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"avx512",
			"aes_ni",
			"popcnt",
		},
		"ivybridge": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"aes_ni",
			"avx",
			"popcnt",
		},
		"kabylake": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"aes_ni",
			"popcnt",
		},
		"sandybridge": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"popcnt",
		},
		"silvermont": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"aes_ni",
			"popcnt",
		},
		"skylake": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"avx512",
			"aes_ni",
			"popcnt",
		},
		"stoneyridge": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"aes_ni",
			"avx",
			"avx2",
			"popcnt",
		},
		"tigerlake": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"avx512",
			"aes_ni",
			"popcnt",
		},
		"whiskeylake": {
			"ssse3",
			"sse4",
			"sse4_1",
			"sse4_2",
			"avx",
			"avx2",
			"avx512",
			"aes_ni",
			"popcnt",
		},
	},
}
