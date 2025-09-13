package main

func main() {}

// Интерфейс для компилируемых ЯП
type compilation interface {
	compile()
}

// Интерфейс для интерпретируемых ЯП
type interpretation interface {
	interpret()
}

type CPlusPlus struct{}

func (s CPlusPlus) compile() {
	println("Compile C++ code")
}

type CSharp struct{}

func (p CSharp) interpret() {
	println("Interpret C# code")
}

type CSharpToCompileAdapter struct {
	C_sharp CSharp
}

func CSharpToCompileAdapterConstructor(c CSharp) CSharpToCompileAdapter {
	return CSharpToCompileAdapter{
		C_sharp: c,
	}
}

func (c CSharpToCompileAdapter) compile() {
	c.C_sharp.interpret()
}
