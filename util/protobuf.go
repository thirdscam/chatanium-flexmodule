package util

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func PbTimestamp2AsTimePtr(t *timestamppb.Timestamp) *time.Time {
	if t == nil {
		return nil
	}

	tm := t.AsTime()
	return &tm
}

func AsTimePtrToPbTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}

	return timestamppb.New(*t)
}

func StringPtrToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func IntPtrToInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}
