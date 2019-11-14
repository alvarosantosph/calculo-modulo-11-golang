/*
 * Cálculo de dígito verificador no módulo 11 padrão febraban (Federação Brasileira de Bancos)
 * Linguagem: GoLang
 * Autor: Alvaro Santos
 * Data: 13/11/2019
 * GitHub: https://github.com/alvarosantosph
 * Email: alvaro_webmaster@hotmail.com
 */

func modulo11(num string) int {
	// Variaveis de controle
	tamanhoString := len(num) + 1
	soma := 0
	resto := 0
	dv := 0
	numeros := make([]string, tamanhoString)
	multiplicador := 2
	runes := []rune(num)
	for i := len(num); i > 0; i-- {
		// Multiplica da direita pra esquerda, incrementando o multiplicador de 2 a 9
		// Caso o multiplicador seja maior que 9 o mesmo recomeça em 2
		if multiplicador > 9 {
			// Pega cada numero isoladamente
			multiplicador = 2
			conteudo := string(runes[i-1 : i])
			conteudoInt, _ := strconv.Atoi(conteudo)
			calculo := conteudoInt * multiplicador
			numeros[i] = strconv.Itoa(calculo)
			multiplicador++
		} else {
			conteudo := string(runes[i-1 : i])
			conteudoInt, _ := strconv.Atoi(conteudo)
			calculo := conteudoInt * multiplicador
			numeros[i] = strconv.Itoa(calculo)
			multiplicador++
		}
	}
	// Realiza a soma de todos os elementos do array e calcula o digito verificador
	// na base 11 de acordo com a regra.
	for i := len(numeros); i > 0; i-- {
		if len(numeros[i-1]) > 0 {
			conteudoSoma, _ := strconv.Atoi(numeros[i-1])
			soma = soma + conteudoSoma
		}
	}
	resto = soma % 11
	dv = 11 - resto
	if dv > 9 || dv == 0 {
		dv = 1
	}