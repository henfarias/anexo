package main

import(
	"encoding/hex"
	//"strings"
	"fmt"
	"bufio"
	"os"
)

func solve(chex, original, modificada string) string {

	cifra, _ := hex.DecodeString(chex)
	
	cifra2 := make([]byte, len(cifra))
	copy(cifra2, cifra)

	for i := 0; i < len(original); i++ {
		cifra2[i] ^= original[i] ^ modificada[i]
	}

	return hex.EncodeToString(cifra2)
}

func main() {

	cin := bufio.NewReader(os.Stdin)

	var chex, original, modificada string
	chex, _ = cin.ReadString('\n')
	original, _ = cin.ReadString('\n')
	modificada, _ = cin.ReadString('\n')

	//arrumando bug de string com tamanho maior por causa do quebra linha
	chex = chex[:len(chex)-1]
	original = original[:len(original)-1]
	modificada = modificada[:len(modificada)-1]

	fmt.Println(solve(chex, original, modificada))
}