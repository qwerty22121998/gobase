package tlf

func Map[I any, O any](arr []I, fn func(elem I) O) []O {
	res := make([]O, 0, len(arr))
	for _, v := range arr {
		res = append(res, fn(v))
	}
	return res
}

func Reduce[I any, O any](arr []I, fn func(acc O, elem I, idx int) O, init O) O {
	var res = init
	for idx, v := range arr {
		res = fn(res, v, idx)
	}
	return res
}
