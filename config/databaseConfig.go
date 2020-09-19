package config

var Username = ""
var Password = ""
var Url = "personalprojectcluster.udqwt.mongodb.net"
var DbName = "Chatter"

func GetDatabaseConnectionString() string {
	return "mongodb+srv://" + Username + ":" + Password + "@" + Url + "/" + DbName + "?retryWrites=true&w=majority"
}
