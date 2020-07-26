// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

type gmachineMemory struct {
	 P  uint64
	 Memory [DefaultMemSize]uint64
}

func New() *gmachineMemory {
	return &gmachineMemory{
		P: 0,
		Memory: [DefaultMemSize]uint64{1:DefaultMemSize-1},
	}
}