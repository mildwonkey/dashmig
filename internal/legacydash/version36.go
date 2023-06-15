package legacydash

import (
	"fmt"

	"github.com/bitly/go-simplejson"
	"github.com/mildwonkey/dashmig/internal/kinds/dashboard"
)

func ReadDashv36(src []byte) (*dashboard.Dashboard, error) {
	orig, err := simplejson.NewJson(src)
	if err != nil {
		return nil, err
	}

	return migrateDashv37(orig)
}

func migrateDashv36(orig *simplejson.Json) (*dashboard.Dashboard, error) {
	// implement v35 -> v36 migration

	// Migrate datasource to refs in annotations
	annos := orig.Get("annotations").Get("list").MustArray()

	type ads struct {
		dsType string
		uid    string
	}

	a := make(map[string]interface{}, len(annos))

	for _, anno := range annos {
		fmt.Printf("anno:\n%v\n\n", anno)
		anno := anno.(map[string]interface{})
		if ds, ok := anno["datasource"]; ok {
			fmt.Printf("ds: %v\n", ds)
		}

	}

	// I just deleted the list. so this is a bit wrong.
	orig.SetPath([]string{"annotations", "list"}, a)

	//fmt.Printf("original annotations: %#v\n", annos)
	fmt.Printf("new annotations: %#v\n", orig.Get("annotations").Get("list"))

	return migrateDashv37(orig)
}
