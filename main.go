// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Bios-Marcel/ent_m2m_pascalcase/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:sqlite.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
}

func Do(ctx context.Context, client *ent.Client) error {
	for i := 0; i < 2; i++ {
		newPermission, err := client.LoginPermission.Create().SetName("my permission").Save(ctx)
		if err != nil {
			return err
		}

		_, err = client.Login.Create().SetName("my login").AddLoginPermissions(newPermission).Save(ctx)
		if err != nil {
			return err
		}
	}

	logins, err := client.Login.Query().WithLoginPermissions().All(ctx)
	if err != nil {
		return err
	}

	for _, login := range logins {
		loginBytes, err := json.Marshal(login)
		if err != nil {
			return err
		}

		fmt.Println(string(loginBytes))
	}

	return nil
}
