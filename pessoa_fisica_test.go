package brasil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCPFValidation(t *testing.T) {
	for index, value := range []CPF{
		// Estes números foram gerados com um site qualquer na internet
		// These numbers were generated from some site on the internet
		"390.533.447-05",
		"259.387.703-00",
		"841.695.027-01",
	} {
		assert.Equal(t, value.String()[9], value.firstVerifier(), "Primeiro digito verificador do CPF %s não validou!", string(value))
		assert.Equal(t, value.String()[10], value.secondVerifier(), "Segundo digito verificador do CPF %s não validou!", string(value))
		assert.NoError(t, value.Validate(), "CPF #%d (%s) não validou!", index, string(value))
	}
}
