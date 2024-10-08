// +build integration

package migration

import (
	"demo/internal/def"

	"github.com/powerman/gotest/testinit"
	"github.com/powerman/pqx"
	"github.com/powerman/structlog"
)

var (
	dbCfg          pqx.Config
	connectTimeout = 3 * def.TestSecond
)

func init() { testinit.Setup(2, setupIntegration) }

func setupIntegration() {
	const dbSuffix = "migration"

	dbCfg = pqx.Config{
		DBName:                          def.DBName,
		User:                            def.DBUser,
		Pass:                            def.DBPass,
		Host:                            def.DBHost,
		Port:                            def.DBPort,
		ConnectTimeout:                  3 * def.TestSecond,
		StatementTimeout:                3 * def.TestSecond,
		LockTimeout:                     3 * def.TestSecond,
		IdleInTransactionSessionTimeout: 3 * def.TestSecond,
		SSLMode:                         pqx.SSLDisable,
	}
	_, cleanup, err := pqx.EnsureTempDB(structlog.New(), dbSuffix, dbCfg)
	if err != nil {
		testinit.Fatal(err)
	}
	testinit.Teardown(cleanup)
	dbCfg.DBName += "_" + dbSuffix
}
