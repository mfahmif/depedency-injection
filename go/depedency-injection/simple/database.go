package simple

type Database struct {
	Name string
}

type DatabaseSQL Database   //create alias untuk menghindari provider lebih dari 1 memiliki type data yg sama
type DatabaseMysql Database //crete alias

func NewDatabaseSQL() *DatabaseSQL {
	return (*DatabaseSQL)(&Database{Name: "SQL"})
}

func NewDatabaseMysql() *DatabaseMysql {
	return (*DatabaseMysql)(&Database{Name: "Mysql"})
}

type DatabaseRepository struct {
	DatabaseSQL   DatabaseSQL
	DatabaseMysql DatabaseMysql
}

func NewDatabaseRepository(sql *DatabaseSQL, mysql *DatabaseMysql) *DatabaseRepository {
	return &DatabaseRepository{DatabaseSQL: *sql, DatabaseMysql: *mysql}
}
