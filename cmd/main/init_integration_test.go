// +build integration

package main

import (
	"fmt"
	"net"
	"strconv"

	"demo/internal/api/restapi/client"

	"github.com/powerman/gotest/testinit"
	"github.com/powerman/pqx"
	"github.com/powerman/structlog"

	"demo/internal/def"
)

func init() { testinit.Setup(2, setupIntegration) }

var (
	extClient *client.Demo
)

func setupIntegration() {
	const dbSuffix = "cmd_test"
	const migrationDir = "../../migration"
	const host = "localhost"

	dbCfg := pqx.Config{
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

	cfg.db = dbCfg
	cfg.gooseDir = migrationDir
	cfg.api.Host = host
	cfg.api.Port = unusedTCPPort(host)

	go func() {
		errc := make(chan error)
		go runServe(errc)
		testinit.Fatal(<-errc)
	}()


	tsURL := fmt.Sprintf("https://%s:%d", cfg.api.Host, cfg.api.Port)
	extClient = client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
}

func unusedTCPPort(host string) (port int) {
	var portStr string
	ln, err := net.Listen("tcp", host+":0")
	if err == nil {
		err = ln.Close()
	}
	if err == nil {
		_, portStr, err = net.SplitHostPort(ln.Addr().String())
	}
	if err == nil {
		port, err = strconv.Atoi(portStr)
	}
	if err != nil {
		testinit.Fatal(err)
	}
	return port
}
