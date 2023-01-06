package main

import "fmt"

type Mediator interface {
	CreateColleagues()
	ColleagueChanged()
}

type Colleague interface {
	SetMediator(mediator Mediator)
	SetColleagueEnabled(enabled bool)
}

type button struct {
	Enabled  bool
	mediator Mediator
}

func (b *button) SetMediator(mediator Mediator) {
	b.mediator = mediator
}

func (b *button) SetColleagueEnabled(enabled bool) {
	b.Enabled = enabled
}

type radioButton struct {
	enabled  bool
	checked  bool
	mediator Mediator
}

func (rb *radioButton) SetMediator(mediator Mediator) {
	rb.mediator = mediator
}

func (rb *radioButton) SetColleagueEnabled(enabled bool) {
	rb.enabled = enabled
}

func (rb *radioButton) Check(checked bool) {
	rb.checked = checked
	rb.mediator.ColleagueChanged()
}

type loginForm struct {
	RadioButton radioButton
	Button      button
}

func NewLoginForm() *loginForm {
	loginForm := &loginForm{}
	loginForm.CreateColleagues()
	return loginForm
}

func (lf *loginForm) CreateColleagues() {
	lf.RadioButton = radioButton{}
	lf.Button = button{}
	lf.RadioButton.SetMediator(lf)
	lf.Button.SetMediator(lf)
}

func (lf *loginForm) ColleagueChanged() {
	if !lf.RadioButton.checked {
		lf.Button.SetColleagueEnabled(false)
	} else {
		lf.Button.SetColleagueEnabled(true)
	}
}

func main() {
	loginForm := NewLoginForm()

	fmt.Println(loginForm.Button.Enabled)

	loginForm.RadioButton.Check(true)
	fmt.Println(loginForm.Button.Enabled)
}
