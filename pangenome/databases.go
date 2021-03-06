package pangenome

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

var Schema = `
	Sequence: string @index(term) @upsert .
`

type KmerNode struct {
	UID          uint64     `json:"uid,omitempty"`
	Sequence     string     `json:"sequence,omitempty"`
	ForwardNodes []KmerNode `json:"forward,omitempty"`
	ReverseNodes []KmerNode `json:"reverse,omitempty"`
}

func setupDgraph(address string, port string) (*dgo.Dgraph, error) {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	db := fmt.Sprintf("%s:%s", address, port)
	d, err := grpc.Dial(db, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	dc := dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)

	setupSchema(dc)

	return dc, err
}

func setupSchema(c *dgo.Dgraph) {
	err := c.Alter(context.Background(), &api.Operation{
		Schema: Schema,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func setupBadger() *badger.DB {
	// Open the Badger database located in the user's home directory.
	// It will be created if it doesn't exist.
	opts := badger.DefaultOptions

	// Get currenty working directory.
	wdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir := path.Join(wdir, "badger")

	opts.Dir = dir
	opts.ValueDir = dir
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
