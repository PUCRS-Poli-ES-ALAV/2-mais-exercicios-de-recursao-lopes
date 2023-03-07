package main

import "fmt"

// 1. Modele e implemente um método recursivo que calcule o fatorial de um número n passado como parâmetro.
func Fatorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("n must be greater than zero")
	}
	if n == 0 {
		return 1, nil
	}

	fat, err := Fatorial(n - 1)
	if err != nil {
		return -1, err
	}

	return n * fat, nil
}

// 2. Modele e implemente um método recursivo que calcule o somatório de um número n (passado como parâmetro) até 0.
func Somatorio(n int) int {
	if n == 0 {
		return 0
	}

	return n + Somatorio(n-1)
}

// 3. Modele e implemente um método recursivo que calcule o n-ésimo número da sequência de fibonacci.
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

// 4. Modele e implemente um método recursivo que calcule o somatório dos números inteiros entre os números k e j, passados como parâmetro.

func SomatorioIntervalo(k, j int) int {
	if k > j {
		return SomatorioIntervalo(j, k)
	}
	if k == j {
		return k
	}

	return k + SomatorioIntervalo(k+1, j)
}

// 5. Modele e implemente um método recursivo que recebe um String em retorna true se este String for um palíndrome, false caso contrário.
//
// 	boolean isPal(String s)

func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return IsPalindrome(s[1 : len(s)-1])
}

// 6.. Modele e implemente um método recursivo que recebe um inteiro zero ou positivo e retorna um String com o número em binário.
// 	String convBase2(int n)

func ConvIntToBin(n int) string {
	if n == 0 {
		return "0"
	}
	if n == 1 {
		return "1"
	}

	return ConvIntToBin(n/2) + fmt.Sprintf("%d", n%2)
}

// 7. Modele e implemente um método recursivo que calcule o somatório dos números contidos em um ArrayList de inteiros, passado como parâmetro.

func SomaLista(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	if len(numbers) == 1 {
		return numbers[0]
	}

	return numbers[0] + SomaLista(numbers[1:])
}
