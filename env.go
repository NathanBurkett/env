package env

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const sliceKey = "key"
const sliceValue = "value"

// Reader Reads environmental config file via bufio.NewScanner. While interpreting each row, will split the row by "="
type Reader struct {
	reader io.Reader
}

type envCollection []map[string]string

// Must Get an env key's value. Log panic if doesn't exist
func Must(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Panic(fmt.Sprintf("env variable %s does not exist", key))
	}

	return value
}

// NewReader Returns a new Reader instance
func NewReader(ioReader io.Reader) Reader {
	return Reader{
		reader: ioReader,
	}
}

// Read and set ENV
func (r Reader) Read() {
	values := r.readEnvFromIo()
	r.setEnvFromSlice(values)
}

// Read env from bufio.
func (r Reader) readEnvFromIo() []map[string]string {
	var vars envCollection

	scanner := bufio.NewScanner(r.reader)

	for scanner.Scan() {
		vars = r.readEnvRow(scanner, vars)
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	return vars
}

func (r Reader) readEnvRow(scanner *bufio.Scanner, env envCollection) envCollection {
	row := strings.SplitN(scanner.Text(), "=", 2)

	if _, exists := os.LookupEnv(row[0]); exists {
		return env
	}

	env = append(env, map[string]string{
		sliceKey:   row[0],
		sliceValue: row[1],
	})

	return env
}

func (r Reader) setEnvFromSlice(envSlice []map[string]string) {
	for i := 0; i < len(envSlice); i++ {
		row := envSlice[i]
		os.Setenv(row[sliceKey], row[sliceValue])
	}
}
