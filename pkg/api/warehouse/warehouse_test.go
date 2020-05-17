package warehouse

import (
	"context"
	"github.com/breathbath/api-go-wrapper/internal/common"
	"testing"
)

//works
func TestErplyClient_GetWarehouses(t *testing.T) {
	const (
		//fill your data here
		sk = ""
		cc = ""
	)

	cli := NewClient(common.NewClient(sk, cc, "", nil))
	resp, err := cli.GetWarehouses(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
