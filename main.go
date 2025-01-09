package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/eseferr/blog-aggregator/internal/config"
	"github.com/eseferr/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main(){
cfg, err := config.Read()
if err != nil {
	fmt.Fprintf(os.Stderr, "error reading config: %v\n", err)
	os.Exit(1)
}

db, err := sql.Open("postgres",cfg.DBURL)
if err != nil {
	fmt.Println(err)
	os.Exit(1)
}
defer db.Close()
dbQueries := database.New(db)
currentState := State{
	cfg: &cfg,
	db: dbQueries,
}
commands := Commands{
	 commandHandlers: make(map[string]func(*State, Command) error),
}
commands.commandHandlers["login"] = handlerLogin
commands.commandHandlers["register"] = handlerRegister
commands.commandHandlers["reset"] = handlerReset
commands.commandHandlers["users"] = handlerGetUsers
commands.commandHandlers["agg"] = handlerAggregator
commands.commandHandlers["addfeed"] = handlerAddFeed
commands.commandHandlers["feeds"] = handlerFeeds
userCommand := os.Args
if len(userCommand) < 2{
	fmt.Println("Invalid Command")
	os.Exit(1)
}

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