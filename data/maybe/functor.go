package maybe

func FMap[A, B any](f func(A) B, v Type[A]) Type[B] {
	if j, ok := v.(Just[A]); ok {
		return Just[B]{f(j.Value)}
	}

	return Nothing[B]{}
}

func FReplace[A, B any](a A, v Type[B]) Type[A] {
	if _, ok := v.(Just[B]); ok {
		return Just[A]{a}
	}

	return Nothing[A]{}
}
