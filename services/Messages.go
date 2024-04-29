package services

import (
	"github.com/labstack/gommon/color"
	"log"
	"runtime"
)

type MessageService struct {
	Messages map[int][]string
}

const (
	MessageInfo                         = 100
	MessageWarning                      = 200
	MessageWarningSingleWordBlacklist   = 201
	MessageWarningSimpleWordTranslation = 202
	MessageError                        = 500
)

func (ms *MessageService) Warning(text string) {
	ms.Messages[MessageWarning] = append(ms.Messages[MessageWarning], text)
}

func (ms *MessageService) HandleError(err error) (b bool) {
	if err != nil {
		// notice that we're using 1, so it will actually log where
		// the error happened, 0 = this function, we don't want that.
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", fn, line, err)
		b = true
	}
	return
}

//this logs the function name as well.
func (ms *MessageService) FancyHandleError(err error) {
	if err != nil {
		// notice that we're using 1, so it will actually log the where
		// the error happened, 0 = this function, we don't want that.
		pc, fn, line, _ := runtime.Caller(1)
		log.Printf("["+color.Red("error")+"] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err)
	}
	return
}

func (ms *MessageService) Error(text string) {
	log.Print(text)
	ms.Messages[MessageError] = append(ms.Messages[MessageError], text)
}

func (ms *MessageService) Info(text string) {
	ms.Messages[MessageInfo] = append(ms.Messages[MessageInfo], text)
}

func (ms *MessageService) Set(text string, messageCode int) {
	ms.Messages[messageCode] = append(ms.Messages[messageCode], text)
}

func (ms *MessageService) Get() map[int][]string {
	return ms.Messages
}

func (ms *MessageService) GetWarnings() []string {
	return ms.Messages[MessageWarning]
}

func (ms *MessageService) GetInfos() []string {
	return ms.Messages[MessageInfo]
}

func (ms *MessageService) GetErrors() []string {
	return ms.Messages[MessageError]
}

func (ms *MessageService) Panic(e error) {
	log.Print(e.Error())
}

func (ms *MessageService) Init() {
}

var sentryInit = false
var Messages = MessageService{Messages: map[int][]string{}}
