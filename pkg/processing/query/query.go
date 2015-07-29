package query

import "sync"

type Request struct {
	TimeRange TimeRange
	Queries   map[string]*Query
	Results   []*TimeSeries
}

func HandleRequest(req *Request) error {

	context := &QueryContext{
		TimeRange: req.TimeRange,
		Queries:   req.Queries,
	}

	_, batchGroups := GetBatchGroups(req)

	return nil
}

func GetBatchGroups(req *Request) (error, []*BatchGroup) {
	return nil, nil
}

type TimeRange struct {
}

type Query struct {
	Query         string
	DataSourcesId int64
	Results       []*TimeSeries
	Exclude       bool
}

type QueryContext struct {
	TimeRange TimeRange
	Queries   map[string]*Query
	Lock      sync.RWMutex
}

type BatchGroup struct {
	Context  *QueryContext
	Queries  []*Query
	Depends  []*sync.WaitGroup
	Children []*sync.WaitGroup
}

func (bg *BatchGroup) Process() error {
	for _, dep := range bg.Depends {
		dep.Wait()
	}
}

type QueryExecutor interface {
	Execute(query *Query, context *QueryContext) error
}

type TimeSeries struct {
}