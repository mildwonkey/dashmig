package legacydash

import (
	"fmt"

	sjson "github.com/bitly/go-simplejson"
	"github.com/mildwonkey/dashmig/internal/kinds/dashboard"
)

func ReadDashv38(src []byte) (*dashboard.Dashboard, error) {
	orig, err := sjson.NewJson(src)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling to simplejson: %s", err)
	}

	// v38 is the current latest, so no migration needed
	return dashboard.ValidateDashJSON(orig)
}

func migrateDashv38(orig *sjson.Json) (*dashboard.Dashboard, error) {
	// v38 is the current latest, so no migration needed
	return dashboard.ValidateDashJSON(orig)
}
