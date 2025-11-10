package pkg1

import (
	"iter"
)

type SliceSeq[E any] struct {
	seq iter.Seq2[int, E]
}
