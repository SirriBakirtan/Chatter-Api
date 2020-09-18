package config

var Username = "dbAdmin"
var Password = "jtb7w3Ax5NMFRmze"
var Url = "personalprojectcluster.udqwt.mongodb.net"
var DbName = "Chatter"

func GetDatabaseConnectionString() string {
	return "mongodb+srv://" + Username + ":" + Password + "@" + Url + "/" + DbName + "?retryWrites=true&w=majority"
}
