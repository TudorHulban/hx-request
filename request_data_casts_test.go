package hxrequest

import (
	"fmt"
	"testing"
)

func TestExtractOptionalTimestamp(t *testing.T) {
	data := RequestData{
		Content: map[string]string{
			"time": "2024-09-16T18:16",
		},
	}

	fmt.Println(
		data.ExtractOptionalTimestampInLocation(
			"time",
			LayoutDateTime,
			LocationRomania,
		),

		data.ExtractOptionalTimestampInLocation(
			"time",
			LayoutDateTime,
			LocationRomania,
		).Time.UnixNano(),
	)
}
