package v1

import (
	"bufio"
	"strconv"
	"strings"
)

func readRecordioMessage(reader *bufio.Reader) ([]byte, error) {
	// Get size as string.
	sizeString, err := reader.ReadString('\n')
	if err != nil {
		return []byte{}, err
	}

	// Convert string to int64.
	sizeInt, err := strconv.ParseInt(strings.TrimSpace(sizeString), 10, 64)
	if err != nil {
		return []byte{}, err
	}

	// Read data specified by the size.
	dataBytes := make([]byte, sizeInt)
	sizeRead := 0

	for int64(sizeRead) < sizeInt {
		n, err := reader.Read(dataBytes[sizeRead:])
		if err != nil {
			return []byte{}, err
		}

		sizeRead += n
	}

	return dataBytes, nil
}
