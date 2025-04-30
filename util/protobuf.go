package util

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func BufTimestamp2AsTime(t *timestamppb.Timestamp) *time.Time {
	if t == nil {
		return nil
	}

	tm := t.AsTime()
	return &tm
}
