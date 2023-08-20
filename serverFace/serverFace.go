package serverFace

type Serverf interface {
	Start()
	Close()
	Serve()
}
