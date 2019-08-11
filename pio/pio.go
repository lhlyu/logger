package pio

import (
	"io"
)

var NewLine = []byte("\n")

var Default = NewRegistry()

func RegisterPrinter(p *Printer) *Registry {
	return Default.RegisterPrinter(p)
}

func Register(printerName string, output io.Writer) *Printer {
	return Default.Register(printerName, output)
}

func Get(printerName string) *Printer {
	return Default.Get(printerName)
}

func Remove(printerName string) {
	Default.Remove(printerName)
}

func Print(v interface{}) (int, error) {
	return Default.Print(v)
}

func Println(v interface{}) (int, error) {
	return Default.Println(v)
}

func Scan(r io.Reader, addNewLine bool) (cancel func()) {
	return Default.Scan(r, addNewLine)
}
