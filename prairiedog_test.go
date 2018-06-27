package prairiedog

import (
	"fmt"
	"testing"

	"github.com/superphy/prairiedog/kmers"
)

func BenchmarkNew(*testing.B) {
	NewGraph()
}
func ExampleKmers() {
	km := kmers.New("testdata/172.fa")
	header, _ := km.Next()
	fmt.Println(header)
	// Output: >gi|1062504329|gb|CP014670.1| Escherichia coli strain CFSAN004177, complete genome
}

func ExampleKmersNext() {
	km := kmers.New("testdata/172.fa")
	header, kmer := km.Next()
	fmt.Println(header)
	fmt.Println(kmer)
	header, kmer = km.Next()
	fmt.Println(header)
	fmt.Println(kmer)
	// Output:
	// >gi|1062504329|gb|CP014670.1| Escherichia coli strain CFSAN004177, complete genome
	// TCGCTTTCGTT
	// >gi|1062504329|gb|CP014670.1| Escherichia coli strain CFSAN004177, complete genome
	// CGCTTTCGTTC
}

func ExampleKmersIndex() {
	km := kmers.New("testdata/ECI-2866_lcl.fasta")
	fmt.Println(len(km.Headers))
	fmt.Println(len(km.Sequences))
	fmt.Println(len(km.Sequences) == len(km.Headers))
	// Note: index starts at 0.
	fmt.Println(km.Headers[0])
	fmt.Println(km.Headers[1])
	fmt.Println(km.Sequences[0])
	fmt.Println(km.Sequences[1])
	fmt.Println(len(km.Sequences[0]))
	// Output:
	// 297
	// 297
	// true
	// >lcl|ECI-2866|NODE_177_length_532_cov_12.8938_ID_353
	// >lcl|ECI-2866|NODE_222_length_438_cov_0.710611_ID_443
	// AACGCGCACTGACGTGAGGCCAGCGTACAGGCCGGATTATCGACATATTTCTGACAGGTGCCGTTATCTGCGGACTGTGTGACATATTTATCCCGGTATGCCCAGCACGCCTGTGTGATGCTCCAGGGTTTACCTTCCATCACACCTGTTTTCGTCCCCCCCGGCTCTGAACACTCAGTACCTTTCAGCACGCCATCCGCTTTATTAAACGGACAACTCTCCACCCACTCCACCCGTGGAACCCATTCCTTATCACGGACCTTCATCCTGAGTTTCAGCGTAAAGGTGGAAGCACCACTGACAAGCGATTCATAGACCATCCTGTCACCATTCCCGTGCGGGAGGCAATTACCGTTTGCAGTACAGCCACTACCGATCAGAACCTGCCCCTGTGTGACAGAAAACCCGGAGGGCACAGGTATGGTGAAAGTCCCGCTCTGTCCTGTAACCAGTTGCACATTAAAGGCTGTGTTCATAAAGTCGTAACGGGAGTTAAGGAAATATAGCCCTGCATGAGCCGACAGCGAGGCAC
	// ACATCGTGCCGCATTGTTGGCAGAGGGAATTCCTTTTCATTGCTTTTATTATCCCTGTGTTAGTGAAATACTACGTTAGGGTTTGGAACACGTAAGAAAAATGGCGTTGTCAATGGGATTGTTTTTTTTTTATGCCGGTCAGATCTCAAAAACTAGGCCAGAGATCAATTCTACTTGACCTCATGACAGTTTACTGCCGCTGCTGCCGGAATCCAAATCTCGTGGTATCCTAACTCAAGGAGTCGGCATGAAGTCCATCGAAGCATATTTTCTGTTCATCAGGTATTGACTAGTGACTCTGCAAGGACAAATCACCTTACTACATCCTGGTCCATGGTGAAGTCTAGCTTTGATACCTTGAGTTGTCCATTCCCGGAAATGCACCTCCGGGCCAGGGGTGCTCGCTCTGACCTTCGTGTCCCATGGAACTTCAGCCAG
	// 540
}

// ExampleKmersIndexEnd checks end case.
func ExampleKmersIndexEnd() {
	km := kmers.New("testdata/GCA_900015695.1_ED647_contigs_genomic.fna")
	header, kmer := km.Next()
	for ; header != ""; header, kmer = km.Next() {
	}
	fmt.Println(header)
	fmt.Println(kmer)
	// Output:
	// >FAVS01000195.1 Escherichia coli strain ED647 genome assembly, contig: out_195, whole genome shotgun sequence
	// TGTGGCTCAGC
}
