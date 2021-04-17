package result

import "github.com/Gumihoy/gumiho-sql-go/sql/db"

type SQLTransformResult struct {
	SourceSQL string
	TargetSQL string
	results   []*SQLObjectResult

	changes   []*SQLTransformChange
	warnnings []*SQLTransformWarnning
	errors    []*SQLTransformError

	createTableCount int

	insertCount            int
	deleteCount            int
	selectCount            int
	updateCount            int
	createTypeCount        int
	createPackageCount     int
	createPackageBodyCount int
	createFunctionCount    int
	createProcedureCount   int
}

func NewSQLTransformResult(sourceSQL string) *SQLTransformResult {
	var x SQLTransformResult
	x.SourceSQL = sourceSQL
	return &x
}

func (x *SQLTransformResult) AddObjectResult(result *SQLObjectResult) {
	if result == nil {
		return
	}
	x.results = append(x.results, result)
}
func (x *SQLTransformResult) AddChanges(changes []*SQLTransformChange) {
	if changes == nil || len(changes) == 0 {
		return
	}
	for _, change := range changes {
		x.changes = append(x.changes, change)
	}
}
func (x *SQLTransformResult) AddWarnnings(warnnings []*SQLTransformWarnning) {
	if warnnings == nil || len(warnnings) == 0 {
		return
	}
	for _, warnning := range warnnings {
		x.warnnings = append(x.warnnings, warnning)
	}
}
func (x *SQLTransformResult) AddErrors(errors []*SQLTransformError) {
	if errors == nil || len(errors) == 0 {
		return
	}
	for _, error := range errors {
		x.errors = append(x.errors, error)
	}
}

type SQLObjectResult struct {
	objectType db.SQLObjectType
	targetSQL  string
}

func NewSQLObjectResult(objectType db.SQLObjectType, targetSQL string) *SQLObjectResult {
	x := new(SQLObjectResult)
	x.objectType = objectType
	x.targetSQL = targetSQL
	return x
}

type SQLTransformChange struct {
}
type SQLTransformWarnning struct {
}
type SQLTransformError struct {
}
