package template_method

import "testing"

func TestTemplateMethod(t *testing.T) {
	tea := NewTea()
	tea.TemplateMethod()

	coffee := NewCoffee()
	coffee.TemplateMethod()
}
