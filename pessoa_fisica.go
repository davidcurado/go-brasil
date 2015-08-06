package brasil

import (
	"fmt"
	"regexp"
	"strings"
)

const cpflen = 11

// CPF é um tipo que define um número de CPF (Cadastro de pessoa física)
type CPF [cpflen]uint8

var (
	cpfre = regexp.MustCompile("^[0-9]{11}$")
)

// ParseCPF é uma função que cria CPFs
func ParseCPF(value string) (*CPF, error) {
	var cpf CPF
	parsed := strings.Replace(value, "-", "", -1)
	parsed = strings.Replace(parsed, ".", "", -1)
	if !cpfre.Match([]byte(parsed)) {
		return nil, fmt.Errorf("O valor %s não parece estar no formato de um CPF", value)
	}
	for index, char := range parsed {
		cpf[index] = uint8(char) - uint8('0')
	}
	return &cpf, nil
}

func (cpf CPF) String() string {
	var val [cpflen]uint8
	const zero = uint8('0')
	for index, char := range cpf {
		val[index] = zero + char
	}
	return string(val[:])
}

func (cpf CPF) firstVerifier() uint8 {
	sum := 0
	for i := 0; i < cpflen-2; i++ {
		sum += int(cpf[i]) * (10 - i)
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return uint8(remainder - 11)
}

func (cpf CPF) secondVerifier() uint8 {
	// TODO: implement this!
	return cpf[10]
}
