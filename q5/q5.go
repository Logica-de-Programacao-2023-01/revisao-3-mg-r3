package q5

//Uma frase é um palíndromo se, após converter todas as letras maiúsculas em letras minúsculas e remover todos os
//caracteres não alfanuméricos, ela for lida da mesma forma da esquerda para a direita e vice-versa. Caracteres
//alfanuméricos incluem letras e números.
//
//Dada uma string "s", retorne verdadeiro se for um palíndromo e falso caso contrário.
func reverse(str string) string {
	bytes := []byte(str)
	length := len(bytes)
	for i := 0; i < length/2; i++ {
		bytes[i], bytes[length-i-1] = bytes[length-i-1], bytes[i]
	}
	return string(bytes)
}

func IsPalindrome(s string) bool {
	var str string
	rs := reverse(s)
	if str == rs {
		return true
	} else {
		return false
	}
}

