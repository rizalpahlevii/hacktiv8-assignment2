package helpers

import (
	"fmt"
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		tx.Rollback()
		fmt.Println("Transaction rolled back")
	} else {
		tx.Commit()
		fmt.Println("Transaction commited")
	}
}
