package main

import (
	"context"
	"fmt"
	"github.com/keanutaufan/anitrackr-server/platform/firebase_app"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
)

func main() {
	bgCtx := context.Background()
	fc, err := firebase_app.NewFirebaseClient(bgCtx, os.Getenv("FIREBASE_CREDENTIALS_FILE"))
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name: "firebase_admin",
		Commands: []*cli.Command{
			newFcCommand(fc),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newFcCommand(fc *firebase_app.FirebaseClient) *cli.Command {
	return &cli.Command{
		Name:  "auth",
		Usage: "authentication tools",
		Subcommands: []*cli.Command{
			{
				Name:  "set-id",
				Usage: "set the firebase id claims",
				Action: func(c *cli.Context) error {
					uid := c.Args().First()
					dbId := c.Args().Get(1)

					if uid == "" || dbId == "" {
						return fmt.Errorf("usage: firebase_admin auth set-id db-id")
					}

					dbIdInt, err := strconv.ParseInt(dbId, 10, 64)
					if err != nil {
						return err
					}

					claims := map[string]interface{}{
						"app_user_id": dbIdInt,
					}
					err = fc.Auth.SetCustomUserClaims(c.Context, uid, claims)
					if err != nil {
						return err
					}

					fmt.Printf("Successfully set user id %d for user %s\n", dbIdInt, uid)
					return nil
				},
			},
		},
	}
}
