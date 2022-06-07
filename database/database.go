package database

var connection string
var statusConnection string
var iphost string

//var connectionDatabase string

func init() {
	connection = "MySQL"
	statusConnection = "TRUE"
	iphost = "localhost"
	//connectionDatabase = "localhost:192.168.100.24"
}

func GetDatabase() string {
	return connection
}

func StatusConnection() string {
	return statusConnection
}

func IpConnection() string {
	return iphost
}
