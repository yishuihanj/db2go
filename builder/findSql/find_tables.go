package findSql

//查找该数据库的数据库表
func FindTables(model SqlInterface) ([]string, error) {
	db := model.GetDB()
	rows, err := db.Query(model.FindTableString())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tablesName := make([]string, 0, 0)
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		tablesName = append(tablesName, tableName)
	}
	return tablesName, nil
}
