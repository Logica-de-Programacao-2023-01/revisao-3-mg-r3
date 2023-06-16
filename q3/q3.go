package q3

//Você está subindo uma escada. Leva n passos para chegar ao topo.
//
//A cada vez, você pode subir 1 ou 2 degraus. De quantas maneiras distintas você pode subir até o topo?

func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	maneirasDeSubir := 0
	sobe1, sobe2 := 1, 2

	for i := 3; i <= n; i++ {
		maneirasDeSubir = sobe1 + sobe2
		sobe1, sobe2 = sobe2, maneirasDeSubir
	}

	return maneirasDeSubir
}
