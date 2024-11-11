package claud

type ClaudDb struct {
	url string
}

func NewCloudDb(url string) *ClaudDb {
	return &ClaudDb{
		url: url,
	}
}

func (db *ClaudDb) Read() ([]byte, error) {
	return []byte{}, nil
}

func (db *ClaudDb) Write(content []byte) {

}
