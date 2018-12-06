package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"log"
	"sort"
	"container/list"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true

	//tasks := []string{"cook", "clean", "laundry", "eat", "sleep", "code"}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "lang,l",
			Value:  "english",
			Usage:  "language for the greeting",
			EnvVar: "LEGACY_COMPAT_LANG,APP_LANG,LANG",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from FILE",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:     "complete",
			Aliases:  []string{"c"},
			Usage:    "complete a task on list",
			Category: "Complete",
			Action: func(c *cli.Context) error {
				fmt.Println("args:", c.NArg())
				fmt.Println("complete task ", c.Args().First())
				return nil
			},
		},
		{
			Name:     "add",
			Aliases:  []string{"a"},
			Usage:    "add a task on list",
			Category: "Complete",
			Action: func(c *cli.Context) error {
				fmt.Println("add task ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:     "add",
					Usage:    "add a new template",
					Category: "template options",
					Action: func(c *cli.Context) error {
						fmt.Println("add a new task template ", c.Args().First())
						return nil
					},
				},
				{
					Name:     "remove",
					Usage:    "remove a new template",
					Category: "template options1",
					Action: func(c *cli.Context) error {
						fmt.Println("remove a new task template ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	app.BashComplete = func(c *cli.Context) {
		for _, comd := range c.App.Commands {
			fmt.Println("task:", comd.Name)
		}
	}

	//app.Action = func(ctx *cli.Context) error {
	//	name := "someone"
	//	if ctx.NArg() > 0 {
	//		name = ctx.Args().Get(0)
	//	}
	//
	//	if ctx.String("lang") == "spanish" {
	//		fmt.Println("hola ", name)
	//	} else {
	//		fmt.Println("hello ", name)
	//	}
	//
	//	return nil
	//}

	app.Writer = os.Stdout
	app.ErrWriter = os.Stderr

	app.ExitErrHandler = func(c *cli.Context, err error) {

	}
	var c = make(chan int)


	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
