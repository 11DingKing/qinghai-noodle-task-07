package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask07(t *testing.T) {
	now := time.Now()
	s := NewService(NewRegistry(), func() time.Time { return now })
	l := activeLicense(now)
	require.NoError(t, s.CheckTransfer(context.Background(), l, l.OwnerID, l.OwnerID))
}
