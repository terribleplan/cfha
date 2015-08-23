package core

type Status int

const (
  Unknown Status = iota
  Up Status = iota
  Down Status = iota
)

func (t Status) String() string {
  if t == Unknown {
    return "Unknown"
  } else if t == Up {
    return "Up"
  } else if t == Down {
    return "Down"
  }
  return ""
}

type Transition struct {
  To Status
  From Status
  RecordValue string
}
