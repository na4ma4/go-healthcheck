package healthcheck

import protoreflect "google.golang.org/protobuf/reflect/protoreflect"

type genericStatusEnum[T any] interface {
	comparable

	Less(T) bool
	Number() protoreflect.EnumNumber
}

func statusLess[T comparable, P genericStatusEnum[T]](src []P, l, r P) bool {
	if l.Number() == r.Number() {
		return false
	}
	for _, v := range src {
		if v.Number() == l.Number() {
			return false
		} else if v.Number() == r.Number() {
			return true
		}
	}

	return true
}
