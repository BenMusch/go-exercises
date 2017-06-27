package accumulate

const testVersion = 1

func Accumulate(list []string, f func(string) string) []string {
	for i, val := range list {
		list[i] = f(val)
	}
	return list
}
