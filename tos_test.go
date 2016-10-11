package tchannel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Two return value shim
func makeAsOne(a, b interface{}) []interface{} {
	return []interface{}{a, b}
}

func TestAssuredFowardValues(t *testing.T) {
	// RFC2597 Use of DiffServ
	// Assured Fowarding Group is correct
	// The RECOMMENDED values of the AF codepoints are as follows:
	// AF11 = '001010', AF12 = '001100', AF13 = '001110', AF21 = '010010',
	// AF22 = '010100', AF23 = '010110', AF31 = '011010', AF32 = '011100',
	// AF33 = '011110', AF41 = '100010', AF42 = '100100', and AF43 = '100110'.
	assert.Equal(t, makeAsOne(GetTosField("AF11")), makeAsOne(10, nil))
	assert.Equal(t, makeAsOne(GetTosField("AF12")), makeAsOne(12, nil))
	assert.Equal(t, makeAsOne(GetTosField("AF13")), makeAsOne(14, nil))
	assert.Equal(t, makeAsOne(GetTosField("AF21")), makeAsOne(18, nil))
	assert.Equal(t, makeAsOne(GetTosField("AF22")), makeAsOne(20, nil))
	assert.Equal(t, makeAsOne(GetTosField("AF23")), makeAsOne(22, nil))
	assert.Equal(t, makeAsOne(GetTosField("AF31")), makeAsOne(26, nil))
	assert.Equal(t, makeAsOne(GetTosField("AF32")), makeAsOne(28, nil))
	assert.Equal(t, makeAsOne(GetTosField("AF33")), makeAsOne(30, nil))
	assert.Equal(t, makeAsOne(GetTosField("AF41")), makeAsOne(34, nil))
	assert.Equal(t, makeAsOne(GetTosField("AF42")), makeAsOne(36, nil))
	assert.Equal(t, makeAsOne(GetTosField("AF43")), makeAsOne(38, nil))
}

func TestDefaultFowarding(t *testing.T) {
	// RFC4594: Default Forwarding (DF) and Class Selector 0
	// (CS0) provide equivalent behavior and use the same DS codepoint,
	// '000000'.
	assert.Equal(t, makeAsOne(GetTosField("CS0")), makeAsOne(0, nil))
	assert.Equal(t, makeAsOne(GetTosField("DF")), makeAsOne(0, nil))
}

func TestCSConfigValues(t *testing.T) {
	// RFC4594
	assert.Equal(t, makeAsOne(GetTosField("CS2")), makeAsOne(16, nil))
	assert.Equal(t, makeAsOne(GetTosField("CS3")), makeAsOne(24, nil))
	assert.Equal(t, makeAsOne(GetTosField("CS4")), makeAsOne(32, nil))
	assert.Equal(t, makeAsOne(GetTosField("CS5")), makeAsOne(40, nil))
	assert.Equal(t, makeAsOne(GetTosField("CS6")), makeAsOne(48, nil))
	assert.Equal(t, makeAsOne(GetTosField("EF")), makeAsOne(46, nil))
}

func TestBadTosName(t *testing.T) {
	assert.Equal(t, makeAsOne(GetTosField("LOL")), makeAsOne(-1, notRecognizedTosName))
}
