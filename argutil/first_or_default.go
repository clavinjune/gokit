package argutil

func FirstOrDefault[T comparable](defaultArg T, args ...T) T {
	var empty T
	if len(args) >= 1 && args[0] != empty {
		return args[0]
	} else {
		return defaultArg
	}
}
