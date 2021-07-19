package hey

import "fmt"

func Hey(name string) string {
    message := fmt.Sprintf("Hey, %v!", name)
    return message
}

func Cya(name string) string {
    message := fmt.Sprintf("Cya %v!", name)
    return message
}
