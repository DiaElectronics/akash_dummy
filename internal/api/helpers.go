// Code generated by mtgroup-generator.
package api

import (
		"time"
	"strings"

	"demo/internal/api/restapi/models"
	"demo/internal/app"

	extauthapi "github.com/mtgroupit/mt-mock-extauthapi"
		"github.com/go-openapi/strfmt"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!


    func appListParams(apiLP *models.ListParams) *app.ListParams {
        var offset int64 = 0
        if apiLP.Offset != nil {
            offset = *apiLP.Offset
        }
        return &app.ListParams{
            Offset:       offset,
            Limit:        apiLP.Limit,
            FilterGroups: appFilterGroups(apiLP.FilterGroups),
            SortBy:       apiLP.SortBy,
            OrderBy:      apiLP.OrderBy,
        }
    }

	func appFilterGroups(apiFG []*models.FilterGroup) []*app.FilterGroup {
		appFG := []*app.FilterGroup{}
		for _, fg := range apiFG {
			appFG = append(appFG, &app.FilterGroup{
				Key:         fg.Key,
				LogicFilter: fg.LogicFilter,
				Filters:     appFilters(fg.Filters),
			})
		}
		return appFG
	}

	func appFilters(apiFP []*models.Filter) []*app.Filter {
		appF := []*app.Filter{}
		for _, fp := range apiFP {
			appF = append(appF, &app.Filter{
				Value:      fp.Value,
				Operator:   fp.Operator,
				IgnoreCase: fp.IgnoreCase,
			})
		}
		return appF
	}



	func fromDateTimesArray(dts []*strfmt.DateTime) (dates []*time.Time) {
		for _, date := range dts {
			dates = append(dates, (*time.Time)(date))
		}
		return
	}

	func toDateTimesArray(dates []*time.Time) (dts []*strfmt.DateTime) {
		for _, date := range dates {
			dts = append(dts, (*strfmt.DateTime)(date))
		}
		return
	}

	func fromDatesArray(ds []*strfmt.Date) (dates []*time.Time) {
		for _, date := range ds {
			dates = append(dates, (*time.Time)(date))
		}
		return
	}

	func toDatesArray(dates []*time.Time) (ds []*strfmt.Date) {
		for _, date := range dates {
			ds = append(ds, (*strfmt.Date)(date))
		}
		return
	}






func toAppProfile(prof *extauthapi.Profile) app.Profile {
	return app.Profile{
		ID:    prof.ID.String(),
		Authn: prof.Authn,
		Authz: app.Authz{
			User:    prof.Authz.User,
			Manager: prof.Authz.Manager,
			Admin:   prof.Authz.Admin,
		},
		IsolatedEntityID: prof.IsolatedEntityID.String(),
	}
}

func splitCommaSeparatedStr(commaSeparated string) (result []string) {
	for _, item := range strings.Split(commaSeparated, ",") {
		item = strings.TrimSpace(item)
		if item != "" {
			result = append(result, item)
		}
	}
	return
}
