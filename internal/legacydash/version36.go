package legacydash

import (
	"encoding/json"

	"github.com/mildwonkey/dashmig/internal/kinds/dashboard"
	legdash "github.com/mildwonkey/dashmig/internal/legacydash/dashboard"
)

var (
	datasourceStr string = "datasource"
)

func ReadDashv36(src []byte) (*dashboard.Dashboard, error) {
	orig := &dashboard.Dashboard{}
	json.Unmarshal(src, orig)
	return migrateDashv37(orig)
}

func migrateDashv36(orig *legdash.Dashboard) (*dashboard.Dashboard, error) {
	// implement v35 -> v36 migration
	ret := &dashboard.Dashboard{}

	// In this (very made up) example, there's only one change, so we'll cast
	// the original dash to the new model, then upgrade the annotations.
	jsonDash, err := json.Marshal(orig)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(jsonDash, ret)

	// Migrate datasource to refs in annotations
	annos := orig.Annotations.List
	updatedAnnos := make([]*dashboard.AnnotationContainer_AnnotationQuery, len(annos))
	for i, anno := range annos {
		aq := &dashboard.AnnotationContainer_AnnotationQuery{
			Name: anno.Name,
			Datasource: &dashboard.DataSourceRef{
				Type: &datasourceStr,
				Uid:  &anno.Datasource,
			},
			Enable:    anno.Enable,
			Hide:      anno.Hide,
			IconColor: anno.IconColor,
			Type:      anno.Type,
		}
		if anno.Filter != nil {
			aq.Filter = &dashboard.AnnotationPanelFilter{
				Exclude: anno.Filter.Exclude,
				Ids:     anno.Filter.Ids,
			}
		}
		if anno.Target != nil {
			aq.Target = &dashboard.AnnotationTarget{
				Limit:    anno.Target.Limit,
				MatchAny: anno.Target.MatchAny,
				Tags:     anno.Target.Tags,
				Type:     anno.Target.Type,
			}
		}
		updatedAnnos[i] = aq
	}

	ret.Annotations.List = updatedAnnos
	return migrateDashv37(ret)
}
