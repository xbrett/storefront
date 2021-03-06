package database

//SKUDataAccess is for defining different types of databases
type SKUDataAccess interface {
	Insert(string, string, string, int) error
	Get(string) string
	Print() string
	Update(string, int) error
	Delete(string) error
}

type sku struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

//NewSQL initializes a pointer to a sql database
func NewSQL() SKUDataAccess {
	sql := SQLdb{db: openDatabaseConnection()}
	return &sql
}

//NewInMemoryDB initializes a pointer to a local database
func NewInMemoryDB() SKUDataAccess {
	mem := MemDb{db: newConnection()}
	return &mem
}
