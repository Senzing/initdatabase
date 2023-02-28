package senzingschema

import (
	"context"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type SenzingSchema interface {
	Initialize(ctx context.Context) error
	RegisterObserver(ctx context.Context, observer observer.Observer) error
	SetLogLevel(ctx context.Context, logLevel logger.Level) error
	UnregisterObserver(ctx context.Context, observer observer.Observer) error
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "senzing-6212xxxx".
const ProductId = 6212

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message templates for g2config implementations.
var IdMessages = map[int]string{
	1:    "Enter Initialize().",
	2:    "Exit  Initialize() returned (%v).",
	3:    "Enter RegisterObserver(%s).",
	4:    "Exit  RegisterObserver(%s) returned (%v).",
	5:    "Enter SetLogLevel(%v).",
	6:    "Exit  SetLogLevel(%v) returned (%v).",
	7:    "Enter UnregisterObserver(%s).",
	8:    "Exit  UnregisterObserver(%s) returned (%v).",
	901:  "Exit  InitializeSenzingConfiguration() returned (%v).",
	2000: "Entry: %+v",
	2001: "SENZING_ENGINE_CONFIGURATION_JSON: %v",
	8001: "Initialize",
	8002: "RegisterObserver",
	8003: "SetLogLevel",
	8004: "UnregisterObserver",
}

// Status strings for specific messages.
var IdStatuses = map[int]string{}
