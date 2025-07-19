// cmd/ultrasmartcontractsplatform/main.go
package main

import (
"flag"
"log"
"os"

"ultrasmartcontractsplatform/internal/ultrasmartcontractsplatform"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := ultrasmartcontractsplatform.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
