package token

import (
	"fmt"
	"testing"
)

func TestKind_String(t *testing.T) {
	tables := []struct {
		k Kind
		s string
	}{
		{KindUnknown, "unknown"},
		{KindUser, "User"},
		{KindGroupMembership, "GroupMembership"},
		{KindChatbotService, "ChatbotService"},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Checking string for Kind %d", i, table.k)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			kindStr := table.k.String()

			if kindStr != table.s {
				t.Errorf("[%d] unexpected string, got: %s, want: %s,", i, kindStr, table.s)
			}
		})
	}
}
