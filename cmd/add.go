package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

var (
	collName string
	newID    string
)

func init() {
	addCmd.Flags().StringVarP(&collName, "col", "c", "", "collection name")
	addCmd.Flags().StringVar(&newID, "nid", "", "new document id")
	addCmd.MarkFlagRequired("col")
	databaseCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [json data]",
	Short: "add a document to the specified collection",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 || len(args) > 1 {
			return errors.New("invalid number of argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			app      *firebase.App
			fDB      *firestore.Client
			err      error
			jsonData interface{}
			resp     *firestore.WriteResult
		)

		// create firebase
		ctx := context.Background()
		sa := option.WithCredentialsFile(saFile)
		if app, err = firebase.NewApp(ctx, nil, sa); err != nil {
			log.Fatal(err)
		}

		if fDB, err = app.Firestore(ctx); err != nil {
			log.Fatal(err)
		}
		defer fDB.Close()

		if err = json.Unmarshal([]byte(args[0]), &jsonData); err != nil {
			log.Println(args[0])
			log.Fatal(err)
		}

		cRef := fDB.Collection(collName)
		if newID != "" {
			resp, err = cRef.Doc(newID).Set(ctx, jsonData)
		} else {
			_, resp, err = cRef.Add(ctx, jsonData)
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Document successfully created at : ", resp.UpdateTime.Format(time.RFC850))
	},
}
