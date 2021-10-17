package config

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func ExampleUsage() {
	var buf bytes.Buffer
	var w = bufio.NewWriter(&buf)
	PrintUsage(w)
	w.Flush()

	// why? -> fmt deletes spaces at end of "Output" comments
	for _, l := range strings.Split(buf.String(), "\n") {
		fmt.Printf("%s /\n", l)
	}

	// Output:
	// KEY          TYPE      DEFAULT    REQUIRED    DESCRIPTION /
	// HTTP_ADDR    String    :8077                   /
	// GRPC_ADDR    String    :8078                   /
	//  /
}
