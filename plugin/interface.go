package plugin

import (
	"github.com/Yuzuki616/Ratte-Interface/core"
	"github.com/Yuzuki616/Ratte-Interface/panel"
)

type Implemented interface {
	core.Core
	panel.Panel
}
