package mysql

import (
	"github.com/Gumihoy/gumiho-sql-go/translation/config"
	"github.com/Gumihoy/gumiho-sql-go/translation/visitor/oracle"
)

type Oracle2MySQLV5_6ASTTransformVisitor struct {
	*oracle.OracleASTTransformVisitor
}

func NewOracle2MySQLV5_6ASTTransformVisitor(config *config.SQLTransformConfig) *Oracle2MySQLV5_6ASTTransformVisitor {
	x := new(Oracle2MySQLV5_6ASTTransformVisitor)
	x.OracleASTTransformVisitor = oracle.NewOracleASTTransformVisitor(config)
	return x
}

type Oracle2MySQLV5_7ASTTransformVisitor struct {
	*Oracle2MySQLV5_6ASTTransformVisitor
}

func NewOracle2MySQLV5_7ASTTransformVisitor(config *config.SQLTransformConfig) *Oracle2MySQLV5_7ASTTransformVisitor {
	x := new(Oracle2MySQLV5_7ASTTransformVisitor)
	x.Oracle2MySQLV5_6ASTTransformVisitor = NewOracle2MySQLV5_6ASTTransformVisitor(config)
	return x
}

type Oracle2MySQLV8_0ASTTransformVisitor struct {
	*Oracle2MySQLV5_7ASTTransformVisitor
}


func NewOracle2MySQLV8_0ASTTransformVisitor(config *config.SQLTransformConfig) *Oracle2MySQLV8_0ASTTransformVisitor {
	x := new(Oracle2MySQLV8_0ASTTransformVisitor)
	x.Oracle2MySQLV5_7ASTTransformVisitor = NewOracle2MySQLV5_7ASTTransformVisitor(config)
	return x
}
