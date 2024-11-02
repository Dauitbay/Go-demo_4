package files

import (
	"fmt"
	"os"
)

func ReadFile() {

}

func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}

	_, err = file.WriteString(content)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Sucsessful Write")
}
