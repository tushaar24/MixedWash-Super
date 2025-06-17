package utils

import (
  "encoding/json"
  "time"
)

type DateOnly struct {
  time.Time
}

// UnmarshalJSON parses a JSON string like "2025-06-05"
func (d *DateOnly) UnmarshalJSON(b []byte) error {
  // strip the quotes
  var s string
  if err := json.Unmarshal(b, &s); err != nil {
    return err
  }
  if s == "" {
    // treat empty as zero time
    d.Time = time.Time{}
    return nil
  }
  // parse only the date portion
  t, err := time.Parse("2006-01-02", s)
  if err != nil {
    return err
  }
  d.Time = t
  return nil
}

// MarshalJSON (optional) will write back out as "YYYY-MM-DD"
func (d DateOnly) MarshalJSON() ([]byte, error) {
  return json.Marshal(d.Format("2006-01-02"))
}

