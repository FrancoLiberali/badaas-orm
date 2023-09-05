package query

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/ditrit/badaas/orm/model"
)

type GormQuery struct {
	GormDB          *gorm.DB
	ConcernedModels map[reflect.Type][]Table
	initialTable    Table
}

// Order specify order when retrieving models from database.
//
// if descending is true, the ordering is in descending direction.
//
// joinNumber can be used to select the join in case the field is joined more than once.
func (query *GormQuery) Order(field IFieldIdentifier, descending bool, joinNumber int) error {
	table, err := query.GetModelTable(field, joinNumber)
	if err != nil {
		return err
	}

	switch query.Dialector() {
	case Postgres:
		// postgres supports only order by selected fields
		query.AddSelect(table, field)
		query.GormDB = query.GormDB.Order(
			clause.OrderByColumn{
				Column: clause.Column{
					Name: query.getSelectAlias(table, field),
				},
				Desc: descending,
			},
		)

		return nil
	case SQLServer, SQLite, MySQL:
		query.GormDB = query.GormDB.Order(
			clause.OrderByColumn{
				Column: clause.Column{
					Name: field.ColumnSQL(
						query,
						table,
					),
				},
				Desc: descending,
			},
		)

		return nil
	}

	return nil
}

// Offset specify the number of records to skip before starting to return the records
//
// Offset conditions can be cancelled by using `Offset(-1)`.
func (query *GormQuery) Offset(offset int) {
	query.GormDB = query.GormDB.Offset(offset)
}

// Limit specify the number of records to be retrieved
//
// Limit conditions can be cancelled by using `Limit(-1)`
func (query *GormQuery) Limit(limit int) {
	query.GormDB = query.GormDB.Limit(limit)
}

// First finds the first record ordered by primary key, matching given conditions
func (query *GormQuery) First(dest any) error {
	return query.GormDB.First(dest).Error
}

// Take finds the first record returned by the database in no specified order, matching given conditions
func (query *GormQuery) Take(dest any) error {
	return query.GormDB.Take(dest).Error
}

// Last finds the last record ordered by primary key, matching given conditions
func (query *GormQuery) Last(dest any) error {
	return query.GormDB.Last(dest).Error
}

// Find finds all models matching given conditions
func (query *GormQuery) Find(dest any) error {
	return query.GormDB.Find(dest).Error
}

func (query *GormQuery) AddSelect(table Table, fieldID IFieldIdentifier) {
	query.GormDB.Statement.Selects = append(
		query.GormDB.Statement.Selects,
		fmt.Sprintf(
			"%s.%s AS %s",
			table.Alias,
			fieldID.ColumnName(query, table),
			query.getSelectAlias(table, fieldID),
		),
	)
}

func (query *GormQuery) getSelectAlias(table Table, fieldID IFieldIdentifier) string {
	return fmt.Sprintf(
		"\"%[1]s__%[2]s\"", // name used by gorm to load the fields inside the models
		table.Alias,
		fieldID.ColumnName(query, table),
	)
}

func (query *GormQuery) Preload(preloadQuery string, args ...interface{}) {
	query.GormDB = query.GormDB.Preload(preloadQuery, args...)
}

func (query *GormQuery) Unscoped() {
	query.GormDB = query.GormDB.Unscoped()
}

func (query *GormQuery) Where(whereQuery interface{}, args ...interface{}) {
	query.GormDB = query.GormDB.Where(whereQuery, args...)
}

func (query *GormQuery) Joins(joinQuery string, args ...interface{}) {
	query.GormDB = query.GormDB.Joins(joinQuery, args...)
}

func (query *GormQuery) AddConcernedModel(model model.Model, table Table) {
	tableList, isPresent := query.ConcernedModels[reflect.TypeOf(model)]
	if !isPresent {
		query.ConcernedModels[reflect.TypeOf(model)] = []Table{table}
	} else {
		tableList = append(tableList, table)
		query.ConcernedModels[reflect.TypeOf(model)] = tableList
	}
}

func (query *GormQuery) GetTables(modelType reflect.Type) []Table {
	tableList, isPresent := query.ConcernedModels[modelType]
	if !isPresent {
		return nil
	}

	return tableList
}

const UndefinedJoinNumber = -1

func (query *GormQuery) GetModelTable(field IFieldIdentifier, joinNumber int) (Table, error) {
	modelTables := query.GetTables(field.GetModelType())
	if modelTables == nil {
		return Table{}, fieldModelNotConcernedError(field)
	}

	if len(modelTables) == 1 {
		return modelTables[0], nil
	}

	if joinNumber == UndefinedJoinNumber {
		return Table{}, joinMustBeSelectedError(field)
	}

	return modelTables[joinNumber], nil
}

func (query GormQuery) ColumnName(table Table, fieldName string) string {
	return query.GormDB.NamingStrategy.ColumnName(table.Name, fieldName)
}

type Dialector string

const (
	Postgres  Dialector = "postgres"
	MySQL     Dialector = "mysql"
	SQLite    Dialector = "sqlite"
	SQLServer Dialector = "sqlserver"
)

func (query GormQuery) Dialector() Dialector {
	return Dialector(query.GormDB.Dialector.Name())
}

func NewGormQuery(db *gorm.DB, initialModel model.Model, initialTable Table) *GormQuery {
	query := &GormQuery{
		GormDB:          db.Model(initialModel).Select(initialTable.Name + ".*"),
		ConcernedModels: map[reflect.Type][]Table{},
		initialTable:    initialTable,
	}

	query.AddConcernedModel(initialModel, initialTable)

	return query
}

// Get the name of the table in "db" in which the data for "entity" is saved
// returns error is table name can not be found by gorm,
// probably because the type of "entity" is not registered using AddModel
func getTableName(db *gorm.DB, entity any) (string, error) {
	schemaName, err := schema.Parse(entity, &sync.Map{}, db.NamingStrategy)
	if err != nil {
		return "", err
	}

	return schemaName.Table, nil
}

// Find finds all models matching given conditions
func (query *GormQuery) Update(values map[IFieldIdentifier]any) (int64, error) {
	updateMap := map[string]any{}

	for field, value := range values {
		// TODO ver este 0
		table, err := query.GetModelTable(field, 0)
		if err != nil {
			// TODO aca falta agregar el metodo usado
			return 0, err
		}

		updateMap[field.ColumnName(query, table)] = value
	}

	// TODO tambien sacar el preload en caso de que hagan un preload collection
	// Tambien ver el tema de los order y eso
	// y si pongo el returning tambien ver que eso no rompa los find
	query.GormDB.Statement.Selects = []string{}

	// postgre y sqlite son con from, que es lo mismo que hacer un exists en el where
	// mysql y sqlserver permiten update join, lo cual es lo mismo mas que permiten hacer el update de mas de una tabla a la vez
	// pero sqlserver necesita la repeticion de la tabla inicial, al menos segun la doc, se podria probar

	joinClauses := []clause.Join{}
	joins := query.GormDB.Statement.Joins
	lastI := len(joins) - 1
	for i := range query.GormDB.Statement.Joins {
		// TODO quizas para evitarme estos split podria usar bien los joins en la creacion directamente, con el on y eso
		join := joins[lastI-i]
		joinName := strings.ReplaceAll(join.Name, "INNER JOIN ", "")
		joinName = strings.ReplaceAll(joinName, "LEFT JOIN ", "")
		log.Println(joinName)
		joinNameSplit := strings.Split(joinName, " ON ")
		tableNameAndAlias := joinNameSplit[0]
		onStatement := joinNameSplit[1]
		log.Println(tableNameAndAlias)
		log.Println(onStatement)
		tableNameAndAliasSplit := strings.Split(tableNameAndAlias, " ")
		tableName := tableNameAndAliasSplit[0]
		tableAlias := tableNameAndAliasSplit[1]
		log.Println(tableName)
		log.Println(tableAlias)

		if i == lastI || true {
			query.GormDB.Statement.AddClause(
				clause.From{
					Tables: []clause.Table{
						{
							Name:  tableName,
							Alias: tableAlias,
							Raw:   true, // prevent gorm from putting the alias in quotes
						},
					},
					// TODO ver que pone LEFT JOIN siempre
					// Joins: joinClauses,
				},
			)

			query.GormDB = query.GormDB.Where(onStatement, join.Conds...)
		} else {
			onExpressions := []clause.Expression{}
			if strings.Contains(onStatement, query.initialTable.Name) {
				// onStatementSplit := strings.Split(onStatement, " AND ")
				// joinConditionIndex := pie.FindFirstUsing(onStatementSplit, func(condition string) bool {
				// 	return strings.Contains(onStatement, query.initialTable.Name)
				// })
				// log.Println("index")
				// log.Println(joinConditionIndex)
				// // TODO ver si este da que no esta
				// joinCondition := onStatementSplit[joinConditionIndex]
				// onStatement = strings.Join(
				// 	append(onStatementSplit[:joinConditionIndex], onStatementSplit[joinConditionIndex+1:]...),
				// 	" AND ",
				// )
				// join.Conds = append(join.Conds[:joinConditionIndex], join.Conds[joinConditionIndex+1:]...)
				query.GormDB = query.GormDB.Where(onStatement, join.Conds...)
			} else {
				onExpressions = []clause.Expression{
					clause.Expr{SQL: onStatement, Vars: join.Conds},
				}
			}

			joinClauses = append(joinClauses, clause.Join{
				Type: join.JoinType,
				Table: clause.Table{
					Name:  tableName,
					Alias: tableAlias,
					Raw:   true, // prevent gorm from putting the alias in quotes
				},
				ON: clause.Where{Exprs: onExpressions},
			})
		}
	}

	update := query.GormDB.Updates(updateMap)

	return update.RowsAffected, update.Error
}