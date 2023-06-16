package q4

//Dado um array não vazio de números inteiros "nums", cada elemento aparece duas vezes, exceto um. Encontre esse único
//elemento.
//
//Você deve implementar uma solução com complexidade de tempo linear e sem memória extra.

func SingleNumber(nums []int) int {
	elementosDoSlice := make(map[int]int)

	for _, elemento := range nums {
		elementosDoSlice[elemento]++
	}
	
	numeroUnico := []int{}

	for chave, valor := range elementosDoSlice {
		if valor == 1 {
			numeroUnico = append(numeroUnico, chave)
		}
	}

	for g := 0; g < len(nums); g++ {
		if nums[g] == numeroUnico[0] {
			return g
		}
	}

	return 0
}
