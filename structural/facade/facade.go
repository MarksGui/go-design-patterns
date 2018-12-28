package facade

import (
	"fmt"
)

type LetterProcesser interface {
	WriteContent(string)
	writeEnvelope(string)
	letterIntoEnvelope()
	sendLetter()
}

type LetterProcesserImpl struct{}

func (LetterProcesserImpl) WriteContent(content string) {
	fmt.Printf("开始写信，内容:%s\n", content)
}

func (LetterProcesserImpl) writeEnvelope(address string) {
	fmt.Printf("开始写信封，地址:%s\n", address)
}

func (LetterProcesserImpl) letterIntoEnvelope() {
	fmt.Println("正在把信放入信封里面")
}

func (LetterProcesserImpl) sendLetter() {
	fmt.Println("把信送到邮局，投递")
}

type ModernPostOffice struct {
	lp LetterProcesser
}

func (m *ModernPostOffice) SendLetter(content, address string) {
	m.lp.WriteContent(content)
	m.lp.writeEnvelope(address)
	m.lp.letterIntoEnvelope()
	m.lp.sendLetter()
}

func NewModernPostOffice() *ModernPostOffice {
	return &ModernPostOffice{
		lp: LetterProcesserImpl{},
	}
}
