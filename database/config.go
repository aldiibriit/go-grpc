package database

var PeduliATM []byte

func init() {
	PeduliATM = []byte(`
	{
		"DBType": "mysql",
		"Username": "root",
		"Password": "",
		"Protocol": "tcp",
		"Host": "127.0.0.1",
		"Port": "3306",
		"Database": "peduli-atm"
	}`)
}
