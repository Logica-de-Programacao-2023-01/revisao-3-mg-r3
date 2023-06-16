package q1

//Dado um array de números inteiros "nums" e um número inteiro "target", retorne os índices dos dois números cuja soma
//seja igual a "target".
//
//Você pode assumir que cada entrada terá exatamente uma solução e não poderá usar o mesmo elemento duas vezes.
//
//Você pode retornar a resposta em qualquer ordem.

func TwoSum(nums []int, target int) []int {
	indicesArray := []int{}

	for i, valor := range nums {
		for k, outroValor := range nums {
			if i != k && valor+outroValor == target {
				indicesArray = append(indicesArray, i, k)
			}
		}
	}
	return indicesArray
}
