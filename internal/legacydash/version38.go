package legacydash

import (
	"github.com/mildwonkey/dashmig/internal/kinds/dashboard"
)

func ReadDashv38(src []byte) (*dashboard.Dashboard, error) {
	return dashboard.ValidateBytesDash(src)
}

func migrateDashv38(orig *dashboard.Dashboard) (*dashboard.Dashboard, error) {
	return dashboard.ValidateDash(orig)
}
