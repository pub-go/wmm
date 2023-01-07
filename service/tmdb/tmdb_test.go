package tmdb

import (
	"context"
	"testing"
)

var ctx = context.Background()

func TestSearchTV(t *testing.T) {
	err := SearchTV(ctx, "动物王国大冒险")
	if err != nil {
		t.Errorf("err=%+v", err)
	}
}
