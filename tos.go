// Copyright (c) 2015 Uber Technologies, Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package tchannel

import "errors"

var tosFieldMap map[string]int
var errNotRecognizedTosName = errors.New("Unknown Tos Name")

func init() {
	tosFieldMap = map[string]int{
		// DSCP Values
		"DF":   0x00,
		"CS0":  0x00,
		"CS1":  0x08,
		"CS2":  0x10,
		"CS3":  0x18,
		"CS4":  0x20,
		"CS5":  0x28,
		"CS6":  0x30,
		"CS7":  0x38,
		"BE":   0x00,
		"AF11": 0x0a,
		"AF12": 0x0c,
		"AF13": 0x0e,
		"AF21": 0x12,
		"AF22": 0x14,
		"AF23": 0x16,
		"AF31": 0x1a,
		"AF32": 0x1c,
		"AF33": 0x1e,
		"AF41": 0x22,
		"AF42": 0x24,
		"AF43": 0x26,
		"EF":   0x2e,
		// IP Precedence (Linux Socket Compat)
		"LOWDELAY":    0x10,
		"THROUGHPUT":  0x08,
		"RELIABILITY": 0x04,
		"LOWCOST":     0x02,
	}
}

// tosName can be any Definitions for DiffServ Codepoints
// as per RFC2474 or described by IP Precedence descriptions:
// "LOWDELAY", "THROUGHPUT", "RELIABILTIY", "LOWCOST"
func GetTosField(tosName string) (int, error) {
	if fieldValue, ok := tosFieldMap[tosName]; ok {
		return fieldValue, nil
	}
	return -1, errNotRecognizedTosName
}
