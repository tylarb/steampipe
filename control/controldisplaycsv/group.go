package controldisplaycsv

import (
	"log"

	"github.com/turbot/go-kit/helpers"
	typehelpers "github.com/turbot/go-kit/types"

	"github.com/turbot/steampipe/control/execute"
)

type GroupCsvRenderer struct {
	group   *execute.ResultGroup
	columns *execute.ResultColumns
}

func NewGroupCsvRenderer(group *execute.ResultGroup, flatResults *execute.ResultColumns) *GroupCsvRenderer {
	return &GroupCsvRenderer{
		group:   group,
		columns: flatResults,
	}
}

func (r GroupCsvRenderer) Render() []interface{} {
	log.Printf("[TRACE] begin group  csv render '%s'\n", r.group.GroupId)
	defer log.Printf("[TRACE] end table csv render'%s'\n", r.group.GroupId)

	var results []interface{}

	// render children

	for _, childGroup := range r.group.Groups {
		results = append(results, NewGroupCsvRenderer(childGroup, r.columns).Render()...)
	}
	for _, run := range r.group.ControlRuns {
		results = append(results, r.renderControl(run)...)
	}

	return results
}

func (r GroupCsvRenderer) renderControl(run *execute.ControlRun) []interface{} {

	var res []interface{}
	for _, row := range run.Rows {
		for _, prop := range r.columns.GroupColumns {
			val, _ := helpers.GetFieldValueFromInterface(r.group, prop)
			res = append(res, typehelpers.ToString(val))
		}

		for _, prop := range r.columns.ResultColumns {
			val, _ := helpers.GetFieldValueFromInterface(row, prop)
			res = append(res, typehelpers.ToString(val))
		}
		for _, prop := range r.columns.DimensionColumns {
			val, _ := helpers.GetFieldValueFromInterface(row, prop)
			res = append(res, typehelpers.ToString(val))
		}
		tags := make(map[string]string)
		if run.Control.Tags != nil {
			tags = *run.Control.Tags
		}
		for _, prop := range r.columns.TagColumns {
			val := tags[prop]
			res = append(res, typehelpers.ToString(val))
		}
	}
	return res
}
