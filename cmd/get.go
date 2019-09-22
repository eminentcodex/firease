package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

var (
	documentPath string
)

func init() {
	getCmd.Flags().StringVarP(&documentPath, "path", "u", "", "document path to fetch by id")
	getCmd.MarkFlagRequired("path")
	databaseCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get [document url]",
	Short: "fetch a document from a collection by id",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			app     *firebase.App
			fDB     *firestore.Client
			dRef    *firestore.DocumentSnapshot
			docByte []byte
			err     error
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

		// parse document path to get the collection and id
		if strings.Contains(documentPath, "/") == false {
			log.Fatal("invalid document path")
		}

		path := strings.Split(strings.Trim(documentPath, "/"), "/")
		if len(path) < 2 {
			log.Fatal("invalid document path")
		}

		if dRef, err = fDB.Collection(path[0]).Doc(path[1]).Get(ctx); err != nil {
			log.Fatal(err)
		}

		d := dRef.Data()
		// convert to json
		if docByte, err = json.MarshalIndent(d, " ", "    "); err != nil {
			log.Fatal("unable to format document data")
		}

		fmt.Println(string(docByte))
	},
}
