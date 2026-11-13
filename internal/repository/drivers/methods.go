package drivers

func (d Driver) GetListDriversSQL() []Driver {

	var list []Driver

	list = append(list, SQLITE, SQLITE3)

	return list

}
