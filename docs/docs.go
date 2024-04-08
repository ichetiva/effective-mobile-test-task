package docs

import (
	_ "github.com/ichetiva/effective-mobile-test-task/cmd/api"
	_ "github.com/ichetiva/effective-mobile-test-task/pkg/responder"
)

//go:generate swagger generate spec -o ../public/swagger.json --scan-models
