package main

import (
	"github.com/spf13/cobra"
	"github.com/urfave/negroni"
	"github.com/yaien/clothes-store-api/api/routes"
	"github.com/yaien/clothes-store-api/core"
	"log"
)

func server() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		Short: "Start cloth store api server",
		Run: func(cmd *cobra.Command, args []string) {
			app, err := core.NewApp()
			if err != nil {
				log.Fatal(err)
			}
			server := negroni.Classic()
			router := routes.Register(app)
			server.UseHandler(router)
			server.Run(app.Config.Address)
		},
	}
}
