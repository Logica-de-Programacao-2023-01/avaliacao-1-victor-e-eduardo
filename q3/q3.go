package q3

//Você recebe um tabuleiro retangular de M x N quadrados. Além disso, você tem um número ilimitado de peças de dominó
//padrão de 2 x 1 quadrados. Você pode girar as peças. Você deve colocar o máximo de peças de dominó possível no
//tabuleiro, seguindo as seguintes condições:
//
//1. Cada peça de dominó cobre completamente dois quadrados.
//2. Nenhuma duas peças de dominó se sobrepõem.
//3. Cada peça de dominó está completamente dentro do tabuleiro. É permitido que a peça toque as bordas do tabuleiro.
//
//Encontre o número máximo de peças de dominó que podem ser colocadas sob essas restrições.
//
//Se M ou N forem iguais ou menores que 0, a função deve retornar um erro.

func DominoPieces(m, n int) (int, error) {
	var dominós int
	if m, n <= 0 || (m*n) < 2 {
		return 0, fmt.Errorf("O tamanho do tabuleiro é inválido")
	} 
	if m%2 == 0 {
		return (m*n), nill 
	} else if n%2 == 0 {
		return (m*n), nill
	} else if m%2 != 0 && n%2 != 0 {
		return (
	return 0, nil
}
