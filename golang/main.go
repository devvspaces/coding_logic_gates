package logic

func AND(a bool, b bool) bool {
	return a && b
}

func OR(a bool, b bool) bool {
	return a || b
}

func NOT(a bool) bool {
	return !a
}

func NAND(a bool, b bool) bool {
	return NOT(AND(a, b))
}

func XOR(a bool, b bool) bool {
	return a != b
}

func XNOR(a bool, b bool) bool {
	return NOT(XOR(a, b))
}
