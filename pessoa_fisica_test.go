package brasil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCPFValidation(t *testing.T) {
	for index, value := range []string{
		// Estes números foram gerados com um site qualquer na internet
		// These numbers were generated from some site on the internet
		"390.533.447-05",
		"259.387.703-00",
		"841.695.027-01",
	} {
		cpf, err := ParseCPF(value)
		assert.NoError(t, err, "CPF #%d (%s) não validou!", index, value)
		assert.Equal(t, cpf[cpflen-2], cpf.firstVerifier(), "Primeiro digito verificador do CPF %s não validou!", value)
		assert.Equal(t, cpf[cpflen-1], cpf.secondVerifier(), "Segundo digito verificador do CPF %s não validou!", value)
	}
}
