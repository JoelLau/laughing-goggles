package testutil

import "log/slog"

var DiscardLogger = slog.New(slog.DiscardHandler)
