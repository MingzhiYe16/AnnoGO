package db

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

// Session holds the Cassandra session
var Session *gocql.Session

func init() {
    cluster := gocql.NewCluster(os.Getenv("CASSANDRA_HOST")) // Your Cassandra IP
    cluster.Keyspace = "snp_keyspace"
    cluster.Consistency = gocql.One
    var err error
    Session, err = cluster.CreateSession()
    if err != nil {
        log.Fatalf("Could not connect to Cassandra: %v", err)
    }
}
