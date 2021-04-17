package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/comment"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/database"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/function"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/index"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/package"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/packagebody"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/procedure"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/role"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/schema"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/sequence"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/server"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/set"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/show"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/synonym"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/table"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/trigger"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/type"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/typebody"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/user"
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/statement/view"
)

// ---- COMMENT
func AcceptCommentOnAuditPolicyStatement(visitor ISQLVisitor, x *comment.SQLCommentOnAuditPolicyStatement) {
	if visitor.VisitCommentOnAuditPolicyStatement(visitor, x) {
		Accept(visitor, x.Name())
		Accept(visitor, x.Comment())
	}
}
func AcceptCommentOnColumnStatement(visitor ISQLVisitor, x *comment.SQLCommentOnColumnStatement) {
	if visitor.VisitCommentOnColumnStatement(visitor, x) {
		Accept(visitor, x.Name())
		Accept(visitor, x.Comment())
	}
}
func AcceptCommentOnEditionStatement(visitor ISQLVisitor, x *comment.SQLCommentOnEditionStatement) {
	if visitor.VisitCommentOnEditionStatement(visitor, x) {
		Accept(visitor, x.Name())
		Accept(visitor, x.Comment())
	}
}
func AcceptCommentOnIndextypeStatement(visitor ISQLVisitor, x *comment.SQLCommentOnIndextypeStatement) {
	if visitor.VisitCommentOnIndextypeStatement(visitor, x) {
		Accept(visitor, x.Name())
		Accept(visitor, x.Comment())
	}
}
func AcceptCommentOnMaterializedViewStatement(visitor ISQLVisitor, x *comment.SQLCommentOnMaterializedViewStatement) {
	if visitor.VisitCommentOnMaterializedViewStatement(visitor, x) {
		Accept(visitor, x.Name())
		Accept(visitor, x.Comment())
	}
}
func AcceptCommentOnMiningModelStatement(visitor ISQLVisitor, x *comment.SQLCommentOnMiningModelStatement) {
	if visitor.VisitCommentOnMiningModelStatement(visitor, x) {
		Accept(visitor, x.Name())
		Accept(visitor, x.Comment())
	}
}
func AcceptCommentOnOperatorStatement(visitor ISQLVisitor, x *comment.SQLCommentOnOperatorStatement) {
	if visitor.VisitCommentOnOperatorStatement(visitor, x) {
		Accept(visitor, x.Name())
		Accept(visitor, x.Comment())
	}
}
func AcceptCommentOnTableStatement(visitor ISQLVisitor, x *comment.SQLCommentOnTableStatement) {
	if visitor.VisitCommentOnTableStatement(visitor, x) {
		Accept(visitor, x.Name())
		Accept(visitor, x.Comment())
	}
}


// ---------------------------------------- DDL --------------------------------------------------------------------
// ---- Database
func AcceptAlterDatabaseStatement(visitor ISQLVisitor, x *database.SQLAlterDatabaseStatement) {
	if visitor.VisitAlterDatabaseStatement(visitor, x) {

	}
}
func AcceptCreateDatabaseStatement(visitor ISQLVisitor, x *database.SQLCreateDatabaseStatement) {
	if visitor.VisitCreateDatabaseStatement(visitor, x) {

		Accept(visitor, x.Name())

		for _, child := range x.Options() {
			Accept(visitor, child)
		}
	}
}
func AcceptDropDatabaseStatement(visitor ISQLVisitor, x *database.SQLDropDatabaseStatement) {
	if visitor.VisitDropDatabaseStatement(visitor, x) {

		Accept(visitor, x.Name())

	}
}


// ---- Function
func AcceptAlterFunctionStatement(visitor ISQLVisitor, x *function.SQLAlterFunctionStatement) {
	if visitor.VisitAlterFunctionStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}
func AcceptCreateFunctionStatement(visitor ISQLVisitor, x *function.SQLCreateFunctionStatement) {
	if visitor.VisitCreateFunctionStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}
func AcceptDropFunctionStatement(visitor ISQLVisitor, x *function.SQLDropFunctionStatement) {
	if visitor.VisitDropFunctionStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}

// ---- Index
func AcceptAlterIndexStatement(visitor ISQLVisitor, x *index.SQLAlterIndexStatement) {
	if visitor.VisitAlterIndexStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}
func AcceptCreateIndexStatement(visitor ISQLVisitor, x *index.SQLCreateIndexStatement) {
	if visitor.VisitCreateIndexStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}
func AcceptDropIndexStatement(visitor ISQLVisitor, x *index.SQLDropIndexStatement) {
	if visitor.VisitDropIndexStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}

// ---- Package
func AcceptAlterPackageStatement(visitor ISQLVisitor, x *package_.SQLAlterPackageStatement) {
	if visitor.VisitAlterPackageStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}
func AcceptCreatePackageStatement(visitor ISQLVisitor, x *package_.SQLCreatePackageStatement) {
	if visitor.VisitCreatePackageStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}
func AcceptDropPackageStatement(visitor ISQLVisitor, x *package_.SQLDropPackageStatement) {
	if visitor.VisitDropPackageStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}

// ---- Package Body
func AcceptAlterPackageBodyStatement(visitor ISQLVisitor, x *packagebody.SQLAlterPackageBodyStatement) {
	if visitor.VisitAlterPackageBodyStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}
func AcceptCreatePackageBoydStatement(visitor ISQLVisitor, x *packagebody.SQLCreatePackageBoydStatement) {
	if visitor.VisitCreatePackageBoydStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}
func AcceptDropPackageBodyStatement(visitor ISQLVisitor, x *packagebody.SQLDropPackageBodyStatement) {
	if visitor.VisitDropPackageBodyStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}

// ---- Procedure
func AcceptAlterProcedureStatement(visitor ISQLVisitor, x *procedure.SQLAlterProcedureStatement) {
	if visitor.VisitAlterProcedureStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}
func AcceptCreateProcedureStatement(visitor ISQLVisitor, x *procedure.SQLCreateProcedureStatement) {
	if visitor.VisitCreateProcedureStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}
func AcceptDropProcedureStatement(visitor ISQLVisitor, x *procedure.SQLDropProcedureStatement) {
	if visitor.VisitDropProcedureStatement(visitor, x) {

		// Accept(visitor, x.Name())

	}
}

// ---- Role
func AcceptAlterRoleStatement(visitor ISQLVisitor, x *role.SQLAlterRoleStatement) {
	if visitor.VisitAlterRoleStatement(visitor, x) {

	}
}
func AcceptCreateRoleStatement(visitor ISQLVisitor, x *role.SQLCreateRoleStatement) {
	if visitor.VisitCreateRoleStatement(visitor, x) {

	}
}
func AcceptDropRoleStatement(visitor ISQLVisitor, x *role.SQLDropRoleStatement) {
	if visitor.VisitDropRoleStatement(visitor, x) {

	}
}

// ---- Schema
func AcceptAlterSchemaStatement(visitor ISQLVisitor, x *schema.SQLAlterSchemaStatement) {
	if visitor.VisitAlterSchemaStatement(visitor, x) {

	}
}

func AcceptCreateSchemaStatement(visitor ISQLVisitor, x *schema.SQLCreateSchemaStatement) {
	if visitor.VisitCreateSchemaStatement(visitor, x) {

		Accept(visitor, x.Name())

		for _, child := range x.Options() {
			Accept(visitor, child)
		}
	}
}
func AcceptDropSchemaStatement(visitor ISQLVisitor, x *schema.SQLDropSchemaStatement) {
	if visitor.VisitDropSchemaStatement(visitor, x) {

		Accept(visitor, x.Name())

	}
}


// ---- Sequence
func AcceptAlterSequenceStatement(visitor ISQLVisitor, x *sequence.SQLAlterSequenceStatement) {
	if visitor.VisitAlterSequenceStatement(visitor, x) {
		// Accept(visitor, x.Name())
	}
}
func AcceptCreateSequenceStatement(visitor ISQLVisitor, x *sequence.SQLCreateSequenceStatement) {
	if visitor.VisitCreateSequenceStatement(visitor, x) {
		Accept(visitor, x.Name())
		// for _, child := range x.o() {
		// 	Accept(visitor, child)
		// }
	}
}
func AcceptDropSequenceStatement(visitor ISQLVisitor, x *sequence.SQLDropSequenceStatement) {
	if visitor.VisitDropSequenceStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}

// ---- server
func AcceptAlterServerStatement(visitor ISQLVisitor, x *server.SQLAlterServerStatement) {
	if visitor.VisitAlterServerStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}
func AcceptCreateServerStatement(visitor ISQLVisitor, x *server.SQLCreateServerStatement) {
	if visitor.VisitCreateServerStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}
func AcceptDropServerStatement(visitor ISQLVisitor, x *server.SQLDropServerStatement) {
	if visitor.VisitDropServerStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}

// ---- Synonym
func AcceptAlterSynonymStatement(visitor ISQLVisitor, x *synonym.SQLAlterSynonymStatement) {
	if visitor.VisitAlterSynonymStatement(visitor, x) {
		// Accept(visitor, x.Name())
	}
}
func AcceptCreateSynonymStatement(visitor ISQLVisitor, x *synonym.SQLCreateSynonymStatement) {
	if visitor.VisitCreateSynonymStatement(visitor, x) {
		// Accept(visitor, x.Name())
	}
}
func AcceptDropSynonymStatement(visitor ISQLVisitor, x *synonym.SQLDropSynonymStatement) {
	if visitor.VisitDropSynonymStatement(visitor, x) {
		// Accept(visitor, x.Name())
	}
}

// ---- OnTable
func AcceptAlterTableStatement(visitor ISQLVisitor, x *table.SQLAlterTableStatement) {
	if visitor.VisitAlterTableStatement(visitor, x) {

	}
}
func AcceptCreateTableStatement(visitor ISQLVisitor, x *table.SQLCreateTableStatement) {
	if visitor.VisitCreateTableStatement(visitor, x) {

		Accept(visitor, x.Name())

		for _, child := range x.Elements() {
			Accept(visitor, child)
		}

		Accept(visitor, x.PartitionBy())

		Accept(visitor, x.SubQuery())

	}
}
func AcceptDropTableStatement(visitor ISQLVisitor, x *table.SQLDropTableStatement) {
	if visitor.VisitDropTableStatement(visitor, x) {

		Accept(visitor, x.Name())

	}
}


// ---- Trigger
func AcceptAlterTriggerStatement(visitor ISQLVisitor, x *trigger.SQLAlterTriggerStatement) {
	if visitor.VisitAlterTriggerStatement(visitor, x) {
		// Accept(visitor, x.Name())
	}
}
func AcceptCreateTriggerStatement(visitor ISQLVisitor, x *trigger.SQLCreateTriggerStatement) {
	if visitor.VisitCreateTriggerStatement(visitor, x) {
		// Accept(visitor, x.Name())
	}
}
func AcceptDropTriggerStatement(visitor ISQLVisitor, x *trigger.SQLDropTriggerStatement) {
	if visitor.VisitDropTriggerStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}

// ---- Type
func AcceptAlterTypeStatement(visitor ISQLVisitor, x *type_.SQLAlterTypeStatement) {
	if visitor.VisitAlterTypeStatement(visitor, x) {
		// Accept(visitor, x.Name())
	}
}
func AcceptCreateTypeStatement(visitor ISQLVisitor, x *type_.SQLCreateTypeStatement) {
	if visitor.VisitCreateTypeStatement(visitor, x) {
		// Accept(visitor, x.Name())
	}
}
func AcceptDropTypeStatement(visitor ISQLVisitor, x *type_.SQLDropTypeStatement) {
	if visitor.VisitDropTypeStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}

// ---- Type Body
func AcceptAlterTypeBodyStatement(visitor ISQLVisitor, x *typebody.SQLAlterTypeBodyStatement) {
	if visitor.VisitAlterTypeBodyStatement(visitor, x) {
		// Accept(visitor, x.Name())
	}
}
func AcceptCreateTypeBodyStatement(visitor ISQLVisitor, x *typebody.SQLCreateTypeBodyStatement) {
	if visitor.VisitCreateTypeBodyStatement(visitor, x) {
		// Accept(visitor, x.Name())
	}
}
func AcceptDropTypeBodyStatement(visitor ISQLVisitor, x *typebody.SQLDropTypeBodyStatement) {
	if visitor.VisitDropTypeBodyStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}



// ---- User
func AcceptAlterUserStatement(visitor ISQLVisitor, x *user.SQLAlterUserStatement) {
	if visitor.VisitAlterUserStatement(visitor, x) {

	}
}
func AcceptCreateUserStatement(visitor ISQLVisitor, x *user.SQLCreateUserStatement) {
	if visitor.VisitCreateUserStatement(visitor, x) {

	}
}
func AcceptDropUserStatement(visitor ISQLVisitor, x *user.SQLDropUserStatement) {
	if visitor.VisitDropUserStatement(visitor, x) {

	}
}


// ---- View
func AcceptAlterViewStatement(visitor ISQLVisitor, x *view.SQLAlterViewStatement) {
	if visitor.VisitAlterViewStatement(visitor, x) {

	}
}

func AcceptCreateViewStatement(visitor ISQLVisitor, x *view.SQLCreateViewStatement) {
	if visitor.VisitCreateViewStatement(visitor, x) {


	}
}

func AcceptDropViewStatement(visitor ISQLVisitor, x *view.SQLDropViewStatement) {
	if visitor.VisitDropViewStatement(visitor, x) {

		for _, child := range x.Names() {
			Accept(visitor, child)
		}

	}
}


// ---------------------------------------- DML --------------------------------------------------------------------
func AcceptDeleteStatement(visitor ISQLVisitor, x *statement.SQLDeleteStatement)  {
	if visitor.VisitDeleteStatement(visitor, x) {
		Accept(visitor, x.TableReference())
		Accept(visitor, x.WhereClause())
		Accept(visitor, x.OrderByClause())
	}
}

func AcceptInsertStatement(visitor ISQLVisitor, x *statement.SQLInsertStatement)  {
	if visitor.VisitInsertStatement(visitor, x) {
		Accept(visitor, x.TableReference())

		for _, child := range x.Columns() {
			Accept(visitor, child)
		}

		for _, child := range x.Values() {
			Accept(visitor, child)
		}

		Accept(visitor, x.SubQuery())
		// Accept(visitor, x.table())

		for _, child := range x.SetAssignments() {
			Accept(visitor, child)
		}
		for _, child := range x.UpdateAssignments() {
			Accept(visitor, child)
		}
	}
}

func AcceptSelectStatement(visitor ISQLVisitor, x *statement.SQLSelectStatement) {
	if visitor.VisitSelectStatement(visitor, x) {
		Accept(visitor, x.Query())
	}
}

func AcceptUpdateStatement(visitor ISQLVisitor, x *statement.SQLUpdateStatement)  {
	if visitor.VisitUpdateStatement(visitor, x) {
		Accept(visitor, x.TableReference())

		for _, child := range x.Assignments() {
			Accept(visitor, child)
		}

		Accept(visitor, x.WhereClause())
		Accept(visitor, x.OrderByClause())
		Accept(visitor, x.LimitClause())
	}
}


// ---------------------------------------- SET --------------------------------------------------------------------
func AcceptSetVariableAssignmentStatement(visitor ISQLVisitor, x *set.SQLSetVariableAssignmentStatement)  {
	if visitor.VisitSetVariableAssignmentStatement(visitor, x) {
		// Accept(visitor, x.TableReference())
		// Accept(visitor, x.WhereClause())
		// Accept(visitor, x.OrderByClause())
	}
}
func AcceptSetCharacterSetStatement(visitor ISQLVisitor, x *set.SQLSetCharacterSetStatement)  {
	if visitor.VisitSetCharacterSetStatement(visitor, x) {
		// Accept(visitor, x.TableReference())
		// Accept(visitor, x.WhereClause())
		// Accept(visitor, x.OrderByClause())
	}
}
func AcceptSetCharsetStatement(visitor ISQLVisitor, x *set.SQLSetCharsetStatement)  {
	if visitor.VisitSetCharsetStatement(visitor, x) {
		// Accept(visitor, x.TableReference())
		// Accept(visitor, x.WhereClause())
		// Accept(visitor, x.OrderByClause())
	}
}
func AcceptSetNamesStatement(visitor ISQLVisitor, x *set.SQLSetNamesStatement)  {
	if visitor.VisitSetNamesStatement(visitor, x) {
		// Accept(visitor, x.TableReference())
		// Accept(visitor, x.WhereClause())
		// Accept(visitor, x.OrderByClause())
	}
}



// ---------------------------------------- SHOW --------------------------------------------------------------------
func AcceptShowCreateDatabaseStatement(visitor ISQLVisitor, x *show.SQLShowCreateDatabaseStatement)  {
	if visitor.VisitShowCreateDatabaseStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}
func AcceptShowCreateEventStatement(visitor ISQLVisitor, x *show.SQLShowCreateEventStatement)  {
	if visitor.VisitShowCreateEventStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}
func AcceptShowCreateFunctionStatement(visitor ISQLVisitor, x *show.SQLShowCreateFunctionStatement)  {
	if visitor.VisitShowCreateFunctionStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}
func AcceptShowCreateProcedureStatement(visitor ISQLVisitor, x *show.SQLShowCreateProcedureStatement)  {
	if visitor.VisitShowCreateProcedureStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}
func AcceptShowCreateTableStatement(visitor ISQLVisitor, x *show.SQLShowCreateTableStatement)  {
	if visitor.VisitShowCreateTableStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}
func AcceptShowCreateTriggerStatement(visitor ISQLVisitor, x *show.SQLShowCreateTriggerStatement)  {
	if visitor.VisitShowCreateTriggerStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}
func AcceptShowCreateViewStatement(visitor ISQLVisitor, x *show.SQLShowCreateViewStatement)  {
	if visitor.VisitShowCreateViewStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}

// ---------------------------------------- EXPLAIN --------------------------------------------------------------------
func AcceptDescStatement(visitor ISQLVisitor, x *statement.SQLDescStatement)  {
	if visitor.VisitDescStatement(visitor, x) {
		Accept(visitor, x.Table())
		Accept(visitor, x.Column())

	}
}
func AcceptDescribeStatement(visitor ISQLVisitor, x *statement.SQLDescribeStatement)  {
	if visitor.VisitDescribeStatement(visitor, x) {
		Accept(visitor, x.Table())
		Accept(visitor, x.Column())

	}
}
func AcceptExplainStatement(visitor ISQLVisitor, x *statement.SQLExplainStatement)  {
	if visitor.VisitExplainStatement(visitor, x) {
		Accept(visitor, x.Table())
		Accept(visitor, x.Column())

	}
}

func AcceptHelpStatement(visitor ISQLVisitor, x *statement.SQLHelpStatement)  {
	if visitor.VisitHelpStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}

func AcceptUseStatement(visitor ISQLVisitor, x *statement.SQLUseStatement)  {
	if visitor.VisitUseStatement(visitor, x) {
		Accept(visitor, x.Name())
	}
}