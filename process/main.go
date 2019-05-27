package process

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

// Process initializes and runs the process service
func Process(c chan os.Signal, dataBase *gorm.DB) {
	for {
		//fmt.Println("Serving process")
		time.Sleep(1 * time.Second)
	}
}
