package tchannel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssuredFowardValues(t *testing.T) {
	// RFC2597 Use of DiffServ
	// Assured Fowarding Group is correct
	// The RECOMMENDED values of the AF codepoints are as follows:
	// AF11 = '001010', AF12 = '001100', AF13 = '001110', AF21 = '010010',
	// AF22 = '010100', AF23 = '010110', AF31 = '011010', AF32 = '011100',
	// AF33 = '011110', AF41 = '100010', AF42 = '100100', and AF43 = '100110'.
	assert.Equal(t, GetDSCPFieldInt("AF11"), 10)
	assert.Equal(t, GetDSCPFieldInt("AF12"), 12)
	assert.Equal(t, GetDSCPFieldInt("AF13"), 14)
	assert.Equal(t, GetDSCPFieldInt("AF21"), 18)
	assert.Equal(t, GetDSCPFieldInt("AF22"), 20)
	assert.Equal(t, GetDSCPFieldInt("AF23"), 22)
	assert.Equal(t, GetDSCPFieldInt("AF31"), 26)
	assert.Equal(t, GetDSCPFieldInt("AF32"), 28)
	assert.Equal(t, GetDSCPFieldInt("AF33"), 30)
	assert.Equal(t, GetDSCPFieldInt("AF41"), 34)
	assert.Equal(t, GetDSCPFieldInt("AF42"), 36)
	assert.Equal(t, GetDSCPFieldInt("AF43"), 38)
}

func TestDefaultFowarding(t *testing.T) {
	// RFC4594: Default Forwarding (DF) and Class Selector 0
	// (CS0) provide equivalent behavior and use the same DS codepoint,
	// '000000'.
	assert.Equal(t, GetDSCPFieldInt("CS0"), 0)
	assert.Equal(t, GetDSCPFieldInt("DF"), 0)
}

func TestCSConfigValues(t *testing.T) {
	// RFC4594
	assert.Equal(t, GetDSCPFieldInt("CS2"), 16)
	assert.Equal(t, GetDSCPFieldInt("CS3"), 24)
	assert.Equal(t, GetDSCPFieldInt("CS4"), 32)
	assert.Equal(t, GetDSCPFieldInt("CS5"), 40)
	assert.Equal(t, GetDSCPFieldInt("CS6"), 48)
	assert.Equal(t, GetDSCPFieldInt("EF"), 46)
}
