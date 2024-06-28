package network

type Response struct {
	BindUuid	string		`json:"bind_uuid"`		//TODO create alone type
	Payload		any			`json:"payload"`		//TODO create on reflect
	Message		string		`json:"message"`
}

//func NewInstance(config *Config) (store *Store) {
//	return &Store{
//		config : config,
//	}
//}