//go:build tools

package tools

import (
	// Migrations
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	// Protobuf Generation
	_ "github.com/bufbuild/buf/cmd/buf"
)
