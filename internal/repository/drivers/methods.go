package drivers

func GetListSupportedDriversSQL() []Driver {

	var list []Driver

	list = append(list, SQLITE, SQLITE3)

	return list

}

func DriverSupported(driver Driver) bool {

	for _, d := range GetListSupportedDriversSQL() {

		if driver == d {
			return true
		}

	}

	return false

}
