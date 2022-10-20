package stringchains

import (
	"fmt"
	"unicode/utf8"
)

func StringChains() {
	stringToByte()
}

func runesCount() {
	c := "Hola Mundo こんにちは"
	fmt.Printf("%s tiene %v runas\n", c, utf8.RuneCountInString(c))
}

func quoting() {
	strSpn := "Hola mundo"
	strJpn := "おはよう！"

	fmt.Println("Spanish salute: ", strSpn)
	fmt.Println("Japanese salute: ", strJpn)
	strQuote := " \"entrecomillado\" \nnueva linea"

	fmt.Println("Quoting between quotes: ", strQuote)

	strAccent := `El acento grave permite escribir 
	el contenido en lineas diferentes sin
	importar los caracteres como \  o en otros idiomas
	eludiendo a su vez saltos de linea はい！　así es ñ `

	fmt.Println(strAccent)
}

func stringToByte() {
	cad := "una cadena"

	bytes := []byte(cad)
	runas := []rune(cad)

	//Modificando el contenido de las porciones
	bytes[0] = 'U'
	runas = append(runas, '!')

	str1 := string(bytes)
	str2 := string(runas)

	fmt.Println(str1)
	fmt.Println(str2)
}
