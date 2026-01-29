package validators

import (
	"context"
	"main/db"
)

func AccountAtempt(uuid string) error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
		}
	}()
	ctx := context.Background()
	pr, err := client.Account.FindUnique(
		db.Account.UUID.Equals(uuid),
	).Update(
		db.Account.Atempts.Increment(1),
	).Exec(ctx)
	if err != nil {
	}
	if pr != nil {
		return err
	}
	return nil
}
