package app

import "bufio"

// BatchIterator provides a convenient interface for reading lines in batches
type BatchIterator struct {
	max     int
	scanner *bufio.Scanner
	elems   []string
}

// NewBatchIterator returns a new batch iterator  reading from s.
func NewBatchIterator(batchSize int, s *bufio.Scanner) BatchIterator {
	return BatchIterator{
		max:     batchSize,
		scanner: s,
	}
}

// Next advances to next batch
func (it *BatchIterator) Next() bool {
	elems := make([]string, 0)
	scanner := it.scanner
	for i := 0; i < it.max; i++ {
		if scanner.Scan() {
			elems = append(elems, scanner.Text())
		}
		if scanner.Err() != nil {
			break
		}
	}
	it.elems = elems
	return len(elems) > 0
}

// Value returns current batch
func (it *BatchIterator) Value() []string {
	return it.elems
}

// Err returns error
func (it *BatchIterator) Err() error {
	return it.scanner.Err()
}
