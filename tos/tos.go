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

package tos

type ToS uint8

const (
	DF   ToS = 0x00
	CS0  ToS = 0x00
	CS1  ToS = 0x08
	CS2  ToS = 0x10
	CS3  ToS = 0x18
	CS4  ToS = 0x20
	CS5  ToS = 0x28
	CS6  ToS = 0x30
	CS7  ToS = 0x38
	BE   ToS = 0x00
	AF11 ToS = 0x0a
	AF12 ToS = 0x0c
	AF13 ToS = 0x0e
	AF21 ToS = 0x12
	AF22 ToS = 0x14
	AF23 ToS = 0x16
	AF31 ToS = 0x1a
	AF32 ToS = 0x1c
	AF33 ToS = 0x1e
	AF41 ToS = 0x22
	AF42 ToS = 0x24
	AF43 ToS = 0x26
	EF   ToS = 0x2e
	// IP Precedence (Linux Socket Compat)
	LOWDELAY    ToS = 0x10
	THROUGHPUT  ToS = 0x08
	RELIABILITY ToS = 0x04
	LOWCOST     ToS = 0x02
)
