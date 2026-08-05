package mysql

// DB 表示一个数据库连接
type DB struct {
	driverName     string
	dataSourceName string
}

// Open 打开一个数据库连接
func Open(driverName, dataSourceName string) (*DB, error) {
	db := &DB{
		driverName:     driverName,
		dataSourceName: dataSourceName,
	}
	return db, nil
}
