/* 
Collecting command line arguments in CLI

c.Args keeps all of the arguments we entered. Since we know the order of the arguments,
we deduced that the first argument is the name and the remaining values are the marks. We
are checking a flag called save to save those details in a database or not (we don't have
database logic here, for simplicity). app.Version sets the version of the tool. All other
things remain the same as the last program.


*/

package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	// define flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "save",
			Value: "no",
			Usage: "Should save to database (yes/no)",
		},
	}

	app.Version = "1.0"
	// define action
	app.Action = func(c *cli.Context) error {
		var args []string
		if c.NArg() > 0 {
			// Fetch arguments in a array
			args = c.Args()
			personName := args[0]
			marks := args[1:len(args)]
			log.Println("Person: ", personName)
			log.Println("marks", marks)
		}
		// check the flag value
		if c.String("save") == "no" {
			log.Println("Skipping saving to the database")
		} else {
			// Add database logic here
			log.Println("Saving to the database", args)
		}
		return nil
	}

	app.Run(os.Args)
}


// run the program. If we don't give any flag, the default is save=no:
// ./storeMarks --save=yes Albert 89 85 97



