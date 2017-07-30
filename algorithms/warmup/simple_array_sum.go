package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	elements, _ := reader.ReadString('\n')
	elements = strings.TrimSpace(elements)

	sum := 0
	for _, e := range strings.Split(elements, " ") {
		element, _ := strconv.Atoi(e)
		sum += element
	}

	fmt.Println(sum)
}
