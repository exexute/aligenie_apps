package news

import (
	"fmt"
	"testing"
)

func TestGetNews(t *testing.T) {
	_, res := get815()
	fmt.Println(res)
}
