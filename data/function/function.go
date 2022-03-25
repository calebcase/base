package function

func Id[A any](a A) A {
	return a
}

func Const[A, B any](a A, _ B) A {
	return a
}

func Compose[A, B, C any](bc func(B) C, ab func(A) B, a A) C {
	return bc(ab(a))
}

func Flip[A, B, C any](abc func(A, B) C, b B, a A) C {
	return abc(a, b)
}

func Apply[A, B any](ab func(A) B, a A) B {
	return ab(a)
}

func On[A, B, C any](b func(B, B) C, u func(A) B, x A, y A) C {
	return b(u(x), u(y))
}
