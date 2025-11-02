package struct2buf

import (
	"time"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimestampPtr converts a *time.Time to *timestamppb.Timestamp
func TimestampPtr(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

// Timestamp converts a time.Time to *timestamppb.Timestamp
func Timestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
