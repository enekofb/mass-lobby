package aws

import (
	"fmt"
	"time"
)

func Randomize(toRandomize string) string {
	return fmt.Sprintf("%s-%d", toRandomize, uint64(time.Now().UnixNano()))
}
