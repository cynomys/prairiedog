package kmers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Kmers contains all vars required to generate Kmers.
type Kmers struct {
	src       string
	Headers   []string // location of all the headers in lines.
	Sequences []string // location of all the sequences in lines.
	li        int      // line index in Headers and Sequences.
	pi        int      // position index in a slice of Sequences.
	K         int
}

func (km *Kmers) load() {
	file, err := os.Open(km.src)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	seq := make([]byte, 0)
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(s, ">") {
			if len(seq) != 0 {
				km.Sequences = append(km.Sequences, string(seq))
				seq = nil
			}
			km.Headers = append(km.Headers, s)
		} else {
			seq = append(seq, []byte(s)...)
		}
	}
	km.Sequences = append(km.Sequences, string(seq))
	if len(km.Sequences) == 0 {
		log.Fatal("Couldn't load any sequences from file.")
	}
}

// New creates a new Kmere struct.
func New(s string) *Kmers {
	km := &Kmers{
		src: s,
		li:  0,
		pi:  0,
		K:   11,
	}
	km.load()
	return km
}

// HasNext returns true if the source file still has kmers.
func (km *Kmers) HasNext() bool {
	lastOfSequences := km.li == len(km.Sequences)-1
	endOfSeq := km.pi+km.K > len(km.Sequences[km.li])
	return !(lastOfSequences && endOfSeq)
}

// ContigHasNext returns true if the current contig in a source file still has kmers.
func (km *Kmers) ContigHasNext() bool {
	endOfSeq := km.pi+km.K > len(km.Sequences[km.li])
	return !endOfSeq
}

// Next emits the next kmer.
func (km *Kmers) Next() (string, string) {
	// Done.
	if !km.HasNext() {
		return "", ""
	}

	// Move to next sequence.
	if !km.ContigHasNext() {
		km.li++
		km.pi = 0
	}

	// K is greater than the size of the contig.
	if km.K > len(km.Sequences[km.li])-1 {
		log.Printf("WARNING: contig %s is shorter than the chosen k-value of %v. Skipping contig.", km.Headers[km.li], km.K)
		km.li++
		km.pi = 0
	}

	// Fasta header.
	header := km.Headers[km.li]
	// Slice of the sequence.
	// log.Printf("%v, %v, %v, %v", km.li, km.pi, lastOfSequences, endOfSeq)
	sl := km.Sequences[km.li][km.pi : km.pi+km.K]

	// Increment.
	km.pi++

	return header, sl
}
