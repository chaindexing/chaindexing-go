package chaindexing

type ContractAddressRepository struct {
	db Database
}

// NewContractAddressRepository constructor
func NewContractAddressRepository(db *Database) ContractAddressRepository {
	database := GetDB()

	if db != nil {
		database = *db
	}

	return ContractAddressRepository{database}
}

func (c ContractAddressRepository) Create(contractAddress ContractAddress) error {
	return c.db.Create(&contractAddress)
}

func (c ContractAddressRepository) CreateMany(contractAddresses ...ContractAddress) error {
	return c.db.Create(contractAddresses)
}

func (c ContractAddressRepository) Count() (count int64) {
	c.db.Count(&ContractAddress{}, &count)

	return
}
