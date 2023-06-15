package legacydash

import (
	"encoding/json"

	"github.com/mildwonkey/dashmig/internal/kinds/dashboard"
)

func ReadDashv37(src []byte) (*dashboard.Dashboard, error) {
	orig := &dashboard.Dashboard{}
	json.Unmarshal(src, orig)
	return migrateDashv38(orig)
}

func migrateDashv37(orig *dashboard.Dashboard) (*dashboard.Dashboard, error) {
	// implement v36 -> v37 migration
	return migrateDashv38(orig)
}
