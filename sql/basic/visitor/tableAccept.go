package visitor

import (
	"github.com/Gumihoy/gumiho-sql-go/sql/basic/ast/expr/table"
)

func AcceptTableColumn(visitor ISQLVisitor, x *table.SQLTableColumn) {
	if visitor.VisitTableColumn(visitor, x) {

		Accept(visitor, x.Name())
		Accept(visitor, x.DataType())

	}
}
func AcceptPrimaryKeyTableConstraint(visitor ISQLVisitor, x *table.SQLPrimaryKeyTableConstraint) {
	if visitor.VisitPrimaryKeyTableConstraint(visitor, x) {

		Accept(visitor, x.Name())

	}
}
func AcceptUniqueTableConstraint(visitor ISQLVisitor, x *table.SQLUniqueTableConstraint) {
	if visitor.VisitUniqueTableConstraint(visitor, x) {
		Accept(visitor, x.Name())
		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}
func AcceptUniqueIndexTableConstraint(visitor ISQLVisitor, x *table.SQLUniqueIndexTableConstraint) {
	if visitor.VisitUniqueIndexTableConstraint(visitor, x) {
		Accept(visitor, x.Name())
		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}
func AcceptUniqueKeyTableConstraint(visitor ISQLVisitor, x *table.SQLUniqueKeyTableConstraint) {
	if visitor.VisitUniqueKeyTableConstraint(visitor, x) {
		Accept(visitor, x.Name())
		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}
func AcceptForeignKeyTableConstraint(visitor ISQLVisitor, x *table.SQLForeignKeyTableConstraint) {
	if visitor.VisitForeignKeyTableConstraint(visitor, x) {

		Accept(visitor, x.Name())

	}
}
func AcceptCheckTableConstraint(visitor ISQLVisitor, x *table.SQLCheckTableConstraint) {
	if visitor.VisitCheckTableConstraint(visitor, x) {

		Accept(visitor, x.Name())

		Accept(visitor, x.Condition())
	}
}

func AcceptTableLikeClause(visitor ISQLVisitor, x *table.SQLTableLikeClause) {
	if visitor.VisitTableLikeClause(visitor, x) {

		Accept(visitor, x.Name())
	}
}

func AcceptNullColumnConstraint(visitor ISQLVisitor, x *table.SQLNullColumnConstraint)  {
	if visitor.VisitNullColumnConstraint(visitor, x) {

	}
}
func AcceptNotNullColumnConstraint(visitor ISQLVisitor, x *table.SQLNotNullColumnConstraint)  {
	if visitor.VisitNotNullColumnConstraint(visitor, x) {

	}
}
func AcceptPrimaryKeyColumnConstraint(visitor ISQLVisitor, x *table.SQLPrimaryKeyColumnConstraint)  {
	if visitor.VisitPrimaryKeyColumnConstraint(visitor, x) {
		Accept(visitor, x.Name())
	}
}
func AcceptKeyColumnConstraint(visitor ISQLVisitor, x *table.SQLKeyColumnConstraint)  {
	if visitor.VisitKeyColumnConstraint(visitor, x) {

	}
}

func AcceptUniqueColumnConstraint(visitor ISQLVisitor, x *table.SQLUniqueColumnConstraint)  {
	if visitor.VisitUniqueColumnConstraint(visitor, x) {
		Accept(visitor, x.Name())
	}
}
func AcceptCheckColumnConstraint(visitor ISQLVisitor, x *table.SQLCheckColumnConstraint)  {
	if visitor.VisitCheckColumnConstraint(visitor, x) {
		Accept(visitor, x.Name())
	}
}



func AcceptDefaultClause(visitor ISQLVisitor, x *table.SQLDefaultClause)  {
	if visitor.VisitDefaultClause(visitor, x) {
		Accept(visitor, x.Value())
	}
}
func AcceptAutoIncrementExpr(visitor ISQLVisitor, x *table.SQLAutoIncrementExpr)  {
	if visitor.VisitAutoIncrementExpr(visitor, x) {
	}
}
func AcceptVisibleExpr(visitor ISQLVisitor, x *table.SQLVisibleExpr)  {
	if visitor.VisitVisibleExpr(visitor, x) {
	}
}
func AcceptInvisibleExpr(visitor ISQLVisitor, x *table.SQLInvisibleExpr)  {
	if visitor.VisitInvisibleExpr(visitor, x) {
	}
}

func AcceptCommentExpr(visitor ISQLVisitor, x *table.SQLCommentExpr)  {
	if visitor.VisitCommentExpr(visitor, x) {
	}
}




func AcceptCharsetAssignExpr(visitor ISQLVisitor, x *table.SQLCharsetAssignExpr)  {
	if visitor.VisitCharsetAssignExpr(visitor, x) {
		Accept(visitor, x.Value())
	}
}
func AcceptCharacterSetAssignExpr(visitor ISQLVisitor, x *table.SQLCharacterSetAssignExpr)  {
	if visitor.VisitCharacterSetAssignExpr(visitor, x) {
		Accept(visitor, x.Value())
	}
}








func AcceptPartitionByHash(visitor ISQLVisitor, x *table.SQLPartitionByHash) {
	if visitor.VisitPartitionByHash(visitor, x) {

		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}
func AcceptPartitionByKey(visitor ISQLVisitor, x *table.SQLPartitionByKey) {
	if visitor.VisitPartitionByKey(visitor, x) {

		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}
func AcceptPartitionByRange(visitor ISQLVisitor, x *table.SQLPartitionByRange) {
	if visitor.VisitPartitionByRange(visitor, x) {

		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}
func AcceptPartitionByList(visitor ISQLVisitor, x *table.SQLPartitionByList) {
	if visitor.VisitPartitionByList(visitor, x) {

		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}



func AcceptSubPartitionByHash(visitor ISQLVisitor, x *table.SQLSubPartitionByHash) {
	if visitor.VisitSubPartitionByHash(visitor, x) {

		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}
func AcceptSubPartitionByKey(visitor ISQLVisitor, x *table.SQLSubPartitionByKey) {
	if visitor.VisitSubPartitionByKey(visitor, x) {

		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}


func AcceptPartitionDefinition(visitor ISQLVisitor, x *table.SQLPartitionDefinition) {
	if visitor.VisitPartitionDefinition(visitor, x) {

		Accept(visitor, x.Name())
		// for _, child := range x.Columns() {
		// 	Accept(visitor, child)
		// }
	}
}

func AcceptPartitionValuesLessThan(visitor ISQLVisitor, x *table.SQLPartitionValuesLessThan) {
	if visitor.VisitPartitionValuesLessThan(visitor, x) {

		for _, child := range x.Values() {
			Accept(visitor, child)
		}
	}
}
func AcceptPartitionValuesLessThanMaxValue(visitor ISQLVisitor, x *table.SQLPartitionValuesLessThanMaxValue) {
	if visitor.VisitPartitionValuesLessThanMaxValue(visitor, x) {
	}
}
func AcceptPartitionValuesIn(visitor ISQLVisitor, x *table.SQLPartitionValuesIn) {
	if visitor.VisitPartitionValuesIn(visitor, x) {

		for _, child := range x.Values() {
			Accept(visitor, child)
		}
	}
}



func AcceptSubPartitionDefinition(visitor ISQLVisitor, x *table.SQLSubPartitionDefinition) {
	if visitor.VisitSubPartitionDefinition(visitor, x) {

		Accept(visitor, x.Name())
		// for _, child := range x.Columns() {
		// 	Accept(visitor, child)
		// }
	}
}

// ------------------------ Alter Expr
func AcceptAddColumnAlterTableAction(visitor ISQLVisitor, x *table.SQLAddColumnAlterTableAction) {
	if visitor.VisitAddColumnAlterTableAction(visitor, x) {

		for _, child := range x.Columns() {
			Accept(visitor, child)
		}
	}
}
func AcceptAlterColumnAlterTableAction(visitor ISQLVisitor, x *table.SQLAlterColumnAlterTableAction) {
	if visitor.VisitAlterColumnAlterTableAction(visitor, x) {

	}
}
func AcceptDropColumnAlterTableAction(visitor ISQLVisitor, x *table.SQLDropColumnAlterTableAction) {
	if visitor.VisitDropColumnAlterTableAction(visitor, x) {

	}
}
func AcceptAddTableConstraintAlterTableAction(visitor ISQLVisitor, x *table.SQLAddTableConstraintAlterTableAction) {
	if visitor.VisitAddTableConstraintAlterTableAction(visitor, x) {

	}
}


func AcceptDropIndexAlterTableAction(visitor ISQLVisitor, x *table.SQLDropIndexAlterTableAction) {
	if visitor.VisitDropIndexAlterTableAction(visitor, x) {

	}
}
func AcceptDropKeyAlterTableAction(visitor ISQLVisitor, x *table.SQLDropKeyAlterTableAction) {
	if visitor.VisitDropKeyAlterTableAction(visitor, x) {

	}
}



func AcceptDropConstraintTableConstraintAlterTableAction(visitor ISQLVisitor, x *table.SQLDropTableConstraintAlterTableAction) {
	if visitor.VisitDropConstraintTableConstraintAlterTableAction(visitor, x) {

	}
}
func AcceptDropPrimaryKeyTableConstraintAlterTableAction(visitor ISQLVisitor, x *table.SQLDropPrimaryKeyTableConstraintAlterTableAction) {
	if visitor.VisitDropPrimaryKeyTableConstraintAlterTableAction(visitor, x) {

	}
}
func AcceptDropUniqueTableConstraintAlterTableAction(visitor ISQLVisitor, x *table.SQLDropUniqueTableConstraintAlterTableAction) {
	if visitor.VisitDropUniqueTableConstraintAlterTableAction(visitor, x) {

	}
}
func AcceptDropForeignKeyTableConstraintAlterTableAction(visitor ISQLVisitor, x *table.SQLDropForeignKeyTableConstraintAlterTableAction) {
	if visitor.VisitDropForeignKeyTableConstraintAlterTableAction(visitor, x) {

	}
}
func AcceptDropCheckTableConstraintAlterTableAction(visitor ISQLVisitor, x *table.SQLDropCheckTableConstraintAlterTableAction) {
	if visitor.VisitDropCheckTableConstraintAlterTableAction(visitor, x) {

	}
}



func AcceptAddPartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLAddPartitionAlterTableAction) {
	if visitor.VisitAddPartitionAlterTableAction(visitor, x) {

	}
}
func AcceptDropPartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLDropPartitionAlterTableAction) {
	if visitor.VisitDropPartitionAlterTableAction(visitor, x) {

	}
}
func AcceptDiscardPartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLDiscardPartitionAlterTableAction) {
	if visitor.VisitDiscardPartitionAlterTableAction(visitor, x) {

	}
}
func AcceptImportPartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLImportPartitionAlterTableAction) {
	if visitor.VisitImportPartitionAlterTableAction(visitor, x) {

	}
}
func AcceptTruncatePartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLTruncatePartitionAlterTableAction) {
	if visitor.VisitTruncatePartitionAlterTableAction(visitor, x) {

	}
}
func AcceptCoalescePartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLCoalescePartitionAlterTableAction) {
	if visitor.VisitCoalescePartitionAlterTableAction(visitor, x) {

	}
}
func AcceptReorganizePartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLReorganizePartitionAlterTableAction) {
	if visitor.VisitReorganizePartitionAlterTableAction(visitor, x) {

	}
}
func AcceptExchangePartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLExchangePartitionAlterTableAction) {
	if visitor.VisitExchangePartitionAlterTableAction(visitor, x) {

	}
}
func AcceptAnalyzePartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLAnalyzePartitionAlterTableAction) {
	if visitor.VisitAnalyzePartitionAlterTableAction(visitor, x) {

	}
}
func AcceptCheckPartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLCheckPartitionAlterTableAction) {
	if visitor.VisitCheckPartitionAlterTableAction(visitor, x) {

	}
}
func AcceptOptimizePartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLOptimizePartitionAlterTableAction) {
	if visitor.VisitOptimizePartitionAlterTableAction(visitor, x) {

	}
}
func AcceptRebuildPartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLRebuildPartitionAlterTableAction) {
	if visitor.VisitRebuildPartitionAlterTableAction(visitor, x) {

	}
}
func AcceptRepairPartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLRepairPartitionAlterTableAction) {
	if visitor.VisitRepairPartitionAlterTableAction(visitor, x) {

	}
}
func AcceptRemovePartitionAlterTableAction(visitor ISQLVisitor, x *table.SQLRemovePartitionAlterTableAction) {
	if visitor.VisitRemovePartitionAlterTableAction(visitor, x) {

	}
}
/**

  | ANALYZE PARTITION {partition_names | ALL}
  | CHECK PARTITION {partition_names | ALL}
  | OPTIMIZE PARTITION {partition_names | ALL}
  | REBUILD PARTITION {partition_names | ALL}
  | REPAIR PARTITION {partition_names | ALL}
  | REMOVE PARTITIONING
 */




// ------------------------ Drop Expr
func AcceptDropTableStatementRestrictOption(visitor ISQLVisitor, x *table.SQLDropTableStatementRestrictOption) {

}

func AcceptDropTableStatementCascadeOption(visitor ISQLVisitor, x *table.SQLDropTableStatementCascadeOption) {

}
