package drivers

func GetListDriversSQL() []Driver {

	var list []Driver

	list = append(list, SQLITE)

	return list

}
