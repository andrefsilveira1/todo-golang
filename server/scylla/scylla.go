package scylla

import (
	"os"
	"time"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SelectQuery(session *gocql.Session, logger *zap.Logger) {
	logger.Info("Displaying results:")
	q := session.Query("SELECT title, description, completed FROM database")
	var title, description, completed string
	it := q.Iter()
	defer func() {
		if err := it.Close(); err != nil {
			logger.Warn("select database", zap.Error(err))
		}
	}()
	for it.Scan(&title, &description, &completed) {
		logger.Info("\t" + title + "|" + description + "|" + completed)
	}
}

func InsertQuery(session *gocql.Session, logger *zap.Logger) {
	logger.Info("Inserting new data")
	if err := session.Query("INSERT INTO database (title, description, completed) VALUES ('André', 'Teste de insert', 'false')").Exec(); err != nil {
		logger.Error("insert database", zap.Error(err))
	}
}

func DeleteQuery(session *gocql.Session, logger *zap.Logger) {
	logger.Info("Deleting First data")
	if err := session.Query("DELETE FROM database WHERE title='André' and description='Teste de insert'").Exec(); err != nil {
		logger.Error("delete database", zap.Error(err))
	}
}

func CreateCluster(consistency gocql.Consistency, keyspace string, hosts ...string) *gocql.ClusterConfig {
	retryPolicy := &gocql.ExponentialBackoffRetryPolicy{
		Min:        time.Second,
		Max:        10 * time.Second,
		NumRetries: 5,
	}
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Timeout = 5 * time.Second
	cluster.RetryPolicy = retryPolicy
	cluster.Consistency = consistency
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	return cluster
}

func CreateLogger(level string) *zap.Logger {
	lvl := zap.NewAtomicLevel()
	if err := lvl.UnmarshalText([]byte(level)); err != nil {
		lvl.SetLevel(zap.InfoLevel)
	}
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		lvl,
	))
	return logger
}
