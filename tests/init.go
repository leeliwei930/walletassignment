package tests

import (
	"io"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/joho/godotenv"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	_ = godotenv.Load(".env.testing")
}

func ResponseToString(res *http.Response) string {
	resBodyBytes, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	return string(resBodyBytes)
}
