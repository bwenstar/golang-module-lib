package bye

import "fmt"

func Cya(name string) string {
    message := fmt.Sprintf("Cya %v!", name)
    return message
}
