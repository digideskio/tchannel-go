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

var fieldMap map[string]int

const DF = 0
const CS0 = 0
const EF = 46
const CS6 = 48

func init() {
	dscpName := []string{
		"AF11",
		"AF12",
		"AF13",
		"CS2",
		"AF21",
		"AF22",
		"AF23",
		"CS3",
		"AF31",
		"AF32",
		"AF33",
		"CS4",
		"AF41",
		"AF42",
		"AF43",
		"CS5",
		// CS7 Reserved for future use
	}
	value := 10
	fieldMap = make(map[string]int)

	for _, name := range dscpName {
		fieldMap[name] = value
		value += 2
	}
	fieldMap["CS6"] = CS6
	fieldMap["CS0"] = CS0
	fieldMap["DF"] = DF
	fieldMap["EF"] = EF
	// CS7 reserved for future use
}

func GetDSCPFieldInt(dscpName string) int {
	return fieldMap[dscpName]
}
