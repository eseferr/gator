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
currentState := State{
	&cfg,
}
commands := Commands{
	 commandHandlers: make(map[string]func(*State, Command) error),
}
commands.commandHandlers["login"] = handlerLogin
userCommand := os.Args
if len(userCommand) < 2{
	fmt.Println("Invalid Command")
	os.Exit(1)
}
fmt.Println(len(userCommand))

 cmd := Command{
	Name:userCommand[1],
	Args:userCommand[2:],
 }
 
 err = commands.run(&currentState,cmd)
 if err !=nil{
	fmt.Println(err)
	os.Exit(1)
 }

}