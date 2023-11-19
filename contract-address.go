package chaindexing

type ContractAddress struct {
	ID                    uint16 `gorm:"column:id"`
	Address               string `gorm:"column:address;not null"`
	ContractName          string `gorm:"column:contract_name;not null"`
	ChainID               int32  `gorm:"column:chain_id;not null"`
	StartBlockNumber      int64  `gorm:"column:start_block_number;not null"`
	NextBlockToIngestFrom int64  `gorm:"column:next_block_to_ingest_from;not null"`
	NextBlockToHandleFrom int64  `gorm:"column:next_block_to_handle_from;not null"`
}

func (c *ContractAddress) TableName() string {
	return "chaindexing_contract_address"
}
