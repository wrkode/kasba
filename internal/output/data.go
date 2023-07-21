package output

import (
	"github.com/wrkode/kasba/internal/nodeinfo"
	"github.com/wrkode/kasba/internal/util"
)

type TemplateData struct {
	CreatedAt     string
	BOMFormat     string
	Version       string
	NodeInfo      nodeinfo.NodesInfo
	NetworkPlugin string
	Longhorn      bool
	AppsInfo      util.WorkloadInfo
	Errors        util.Errors
}
