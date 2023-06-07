package config

import (
	"main/scylla"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

func Config() {
	logger := scylla.CreateLogger("info")
	cluster := scylla.CreateCluster(gocql.Quorum, "catalog", "scylla-node1")
	session, err := gocql.NewSession(*cluster)
	if err != nil {
		logger.Fatal("Unable to connect to Scylla", zap.Error(err))
	}
	defer session.Close()
	scylla.SelectQuery(session, logger)
	scylla.InsertQuery(session, logger)
	scylla.SelectQuery(session, logger)
	scylla.DeleteQuery(session, logger)
	scylla.SelectQuery(session, logger)
}
