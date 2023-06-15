package legacydash

import (
	sjson "github.com/bitly/go-simplejson"
	"github.com/mildwonkey/dashmig/internal/kinds/dashboard"
)

func ReadDashv37(src []byte) (*dashboard.Dashboard, error) {
	orig, err := sjson.NewJson(src)
	if err != nil {
		return nil, err
	}
	return migrateDashv38(orig)
}

func migrateDashv37(orig *sjson.Json) (*dashboard.Dashboard, error) {
	// implement v36 -> v37 migration
	return migrateDashv38(orig)
}
