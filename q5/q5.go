package q5

import "strings"

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {

	nova_letra := ""
	nova_letra = strings.ReplaceAll(s, "A", "")
	nova_letra = strings.ReplaceAll(nova_letra, "E", "")
	nova_letra = strings.ReplaceAll(nova_letra, "I", "")
	nova_letra = strings.ReplaceAll(nova_letra, "O", "")
	nova_letra = strings.ReplaceAll(nova_letra, "U", "")
	nova_letra = strings.ReplaceAll(nova_letra, "a", "")
	nova_letra = strings.ReplaceAll(nova_letra, "e", "")
	nova_letra = strings.ReplaceAll(nova_letra, "i", "")
	nova_letra = strings.ReplaceAll(nova_letra, "o", "")
	nova_letra = strings.ReplaceAll(nova_letra, "u", "")
	var nova_palavra string
	for i := 0; i < len(nova_letra); i++ {
		nova_palavra += "."
		nova_palavra += string(nova_letra[i])
	}
	novo_texto := strings.ToLower(nova_palavra)

	return novo_texto
}
