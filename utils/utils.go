package utils

import "strconv"

type ParamsCtxKey struct{}
type DBCtxKey struct{}
type LoggerCtxKey struct{}

func GetIdFromPath(path string) (int64, error) {
	idTest := path[1:]

	id, err := strconv.ParseInt(idTest, 10, 64)
	if err != nil {
		return -1, err
	}

	return id, nil
}
