package brasil

import (
	"errors"
	"regexp"
	"strings"
)

// CPF é um tipo que define um número de CPF (Cadastro de pessoa física)
type CPF string

var (
	cpfre = regexp.MustCompile("^[0-9]{11}$")
)

func (cpf CPF) String() string {
	rawval := string(cpf)
	rawval = strings.Replace(rawval, "-", "", -1)
	rawval = strings.Replace(rawval, ".", "", -1)
	return rawval
}

func (cpf CPF) firstVerifier() byte {
	values := []byte(cpf.String())
	zero := byte('0')
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(values[i]-zero) * (10 - i)
	}
	remainder := sum % 11
	if remainder < 2 {
		return '0'
	}
	return zero + byte(remainder-11)
}

func (cpf CPF) secondVerifier() byte {
	// TODO: implement this!
	return cpf.String()[10]
}

// Validate verifica se a string tem o formato adequado de um CPF
func (cpf CPF) Validate() (err error) {
	if !cpfre.MatchString(cpf.String()) {
		return errors.New("O valor passado não parece ser um CPF")
	}
	if cpf.String()[9] != cpf.firstVerifier() {
		return errors.New("Primeiro digito verificador não passou no teste")
	}
	return nil
}
