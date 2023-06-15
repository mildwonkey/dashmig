package legacydash

import (
	"encoding/json"

	"github.com/mildwonkey/dashmig/internal/kinds/dashboard"
	legdash "github.com/mildwonkey/dashmig/internal/legacydash/dashboard" //trolololol
)

func ReadDashv35(src []byte) (*dashboard.Dashboard, error) {
	orig := &legdash.Dashboard{}
	json.Unmarshal(src, orig)
	return migrateDashv36(orig)
}
