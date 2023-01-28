package piscine

func ConcatParams(args []string) string {
	arg := args[0]
	for r := 1; r < len(args); r++ {
		arg += "\n" + args[r]
	}
	return arg
}
