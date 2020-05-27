package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// connect to the cluster
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	fmt.Println("Cassandra connection OK")
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
