package logg

import (
	"cmd/pkg/util"
	"log"
	"os"
)

const loggFail = "../../formkit.log"

var (
	// Info - Информация.
	Info *log.Logger
	// Error - Ошибка.
	Error *log.Logger
)

// InitLogg - Инициализация лога.
func InitLogg() {
	file, err := os.OpenFile(loggFail, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("\033[1;31m[E] Error while opening logger file: %s\033[0m", err.Error())
	}

	Info = log.New(file, "INFO:  ", log.Lmsgprefix|log.LstdFlags)
	Error = log.New(file, "ERROR: ", log.Lmsgprefix|log.LstdFlags)
}

// LogI - Лог информаций.
func LogI(message string) {
	log.Printf("\033[1;34m[I] %s\033[0m", message)
	Info.Printf("%s - %s", util.FileWithFuncAndLineNum(), message)
}

// LogE - Лог ошибок.
func LogE(message error, err string) {
	log.Printf("%s \033[1;31m[E] %s: %s\033[0m", util.FileWithLineNum(), message, err)
	Error.Fatalf("%s - %s: %s", util.FileWithFuncAndLineNum(), message, err)
}
