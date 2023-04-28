package bonus

import "sort"

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.

func CalculateTowers(barLengths []int) (int, int) {
	sort.Ints(barLengths)

	verificacao := 0
	maxheights := 0
	igualdade := false

	for i := 1; i < len(barLengths); i++ {
		if verificacao[i]


			igualdade == true {

		}

	}

	return 0, 0
}
