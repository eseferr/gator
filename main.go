package main

import (
	"fmt"
	"os"

	"github.com/eseferr/blog-aggregator/internal/config"
)

func main(){
cfg, err := config.Read()
if err != nil {
	fmt.Fprintf(os.Stderr, "error reading config: %v\n", err)
	os.Exit(1)
}
err = cfg.SetUser("ergin")
if err !=nil {
	fmt.Fprintf(os.Stderr, "error setting user: %v\n", err)
	os.Exit(1)
}
newCfg, err := config.Read()
    if err != nil {
		fmt.Fprintf(os.Stderr, "error reading config: %v\n", err)
        os.Exit(1)
    }
	fmt.Println(newCfg)
}