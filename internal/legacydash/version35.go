package legacydash

import (
	"github.com/bitly/go-simplejson"
	"github.com/mildwonkey/dashmig/internal/kinds/dashboard"
)

func ReadDashv35(src []byte) (*dashboard.Dashboard, error) {
	orig, err := simplejson.NewJson(src)
	if err != nil {
		return nil, err
	}
	return migrateDashv36(orig)
}
