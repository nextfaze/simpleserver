package main

import "os"
import "log"
import "strconv"
import "net/http"
import "github.com/codegangsta/cli"

func main() {
	app := cli.NewApp()
	app.Name = "SimpleServer"
	app.Usage = "Simple static server to quickly run in any directory"

	app.Flags = []cli.Flag{
		cli.StringFlag{"directory, d", ".", "directory"},
		cli.IntFlag{"port, p", 8080, "port number"},
	}

	app.Action = func(c *cli.Context) {
		if c.IsSet("directory") == false || c.IsSet("port") == false {
			cli.ShowAppHelp(c)
		} else {
			staticHandler := http.FileServer(http.Dir(c.GlobalString("directory")))
			http.Handle("/", staticHandler)

			log.Fatal(http.ListenAndServe(":"+strconv.Itoa(c.GlobalInt("port")), nil))
		}
	}

	app.Run(os.Args)
}
