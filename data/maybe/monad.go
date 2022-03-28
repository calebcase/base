package maybe

func MArrow[A, B any](v Maybe[A], f func(A) Maybe[B]) Maybe[B] {
	if j, ok := v.(Just[A]); ok {
		return f(j.Value)
	}

	return Nothing[B]{}
}

func MSeq[A, B any](_ Maybe[A], v Maybe[B]) Maybe[B] {
	return v
}

func MReturn[A any](v A) Maybe[A] {
	return Just[A]{v}
}
