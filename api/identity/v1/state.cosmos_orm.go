// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package identityv1

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type AccountTable interface {
	Insert(ctx context.Context, account *Account) error
	InsertReturningSequence(ctx context.Context, account *Account) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, account *Account) error
	Save(ctx context.Context, account *Account) error
	Delete(ctx context.Context, account *Account) error
	Has(ctx context.Context, sequence uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, sequence uint64) (*Account, error)
	HasByAddress(ctx context.Context, address string) (found bool, err error)
	// GetByAddress returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByAddress(ctx context.Context, address string) (*Account, error)
	HasByPublicKey(ctx context.Context, public_key []byte) (found bool, err error)
	// GetByPublicKey returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByPublicKey(ctx context.Context, public_key []byte) (*Account, error)
	List(ctx context.Context, prefixKey AccountIndexKey, opts ...ormlist.Option) (AccountIterator, error)
	ListRange(ctx context.Context, from, to AccountIndexKey, opts ...ormlist.Option) (AccountIterator, error)
	DeleteBy(ctx context.Context, prefixKey AccountIndexKey) error
	DeleteRange(ctx context.Context, from, to AccountIndexKey) error

	doNotImplement()
}

type AccountIterator struct {
	ormtable.Iterator
}

func (i AccountIterator) Value() (*Account, error) {
	var account Account
	err := i.UnmarshalMessage(&account)
	return &account, err
}

type AccountIndexKey interface {
	id() uint32
	values() []interface{}
	accountIndexKey()
}

// primary key starting index..
type AccountPrimaryKey = AccountSequenceIndexKey

type AccountSequenceIndexKey struct {
	vs []interface{}
}

func (x AccountSequenceIndexKey) id() uint32            { return 0 }
func (x AccountSequenceIndexKey) values() []interface{} { return x.vs }
func (x AccountSequenceIndexKey) accountIndexKey()      {}

func (this AccountSequenceIndexKey) WithSequence(sequence uint64) AccountSequenceIndexKey {
	this.vs = []interface{}{sequence}
	return this
}

type AccountAddressIndexKey struct {
	vs []interface{}
}

func (x AccountAddressIndexKey) id() uint32            { return 1 }
func (x AccountAddressIndexKey) values() []interface{} { return x.vs }
func (x AccountAddressIndexKey) accountIndexKey()      {}

func (this AccountAddressIndexKey) WithAddress(address string) AccountAddressIndexKey {
	this.vs = []interface{}{address}
	return this
}

type AccountPublicKeyIndexKey struct {
	vs []interface{}
}

func (x AccountPublicKeyIndexKey) id() uint32            { return 2 }
func (x AccountPublicKeyIndexKey) values() []interface{} { return x.vs }
func (x AccountPublicKeyIndexKey) accountIndexKey()      {}

func (this AccountPublicKeyIndexKey) WithPublicKey(public_key []byte) AccountPublicKeyIndexKey {
	this.vs = []interface{}{public_key}
	return this
}

type accountTable struct {
	table ormtable.AutoIncrementTable
}

func (this accountTable) Insert(ctx context.Context, account *Account) error {
	return this.table.Insert(ctx, account)
}

func (this accountTable) Update(ctx context.Context, account *Account) error {
	return this.table.Update(ctx, account)
}

func (this accountTable) Save(ctx context.Context, account *Account) error {
	return this.table.Save(ctx, account)
}

func (this accountTable) Delete(ctx context.Context, account *Account) error {
	return this.table.Delete(ctx, account)
}

func (this accountTable) InsertReturningSequence(ctx context.Context, account *Account) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, account)
}

func (this accountTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this accountTable) Has(ctx context.Context, sequence uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, sequence)
}

func (this accountTable) Get(ctx context.Context, sequence uint64) (*Account, error) {
	var account Account
	found, err := this.table.PrimaryKey().Get(ctx, &account, sequence)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &account, nil
}

func (this accountTable) HasByAddress(ctx context.Context, address string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		address,
	)
}

func (this accountTable) GetByAddress(ctx context.Context, address string) (*Account, error) {
	var account Account
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &account,
		address,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &account, nil
}

func (this accountTable) HasByPublicKey(ctx context.Context, public_key []byte) (found bool, err error) {
	return this.table.GetIndexByID(2).(ormtable.UniqueIndex).Has(ctx,
		public_key,
	)
}

func (this accountTable) GetByPublicKey(ctx context.Context, public_key []byte) (*Account, error) {
	var account Account
	found, err := this.table.GetIndexByID(2).(ormtable.UniqueIndex).Get(ctx, &account,
		public_key,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &account, nil
}

func (this accountTable) List(ctx context.Context, prefixKey AccountIndexKey, opts ...ormlist.Option) (AccountIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return AccountIterator{it}, err
}

func (this accountTable) ListRange(ctx context.Context, from, to AccountIndexKey, opts ...ormlist.Option) (AccountIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return AccountIterator{it}, err
}

func (this accountTable) DeleteBy(ctx context.Context, prefixKey AccountIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this accountTable) DeleteRange(ctx context.Context, from, to AccountIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this accountTable) doNotImplement() {}

var _ AccountTable = accountTable{}

func NewAccountTable(db ormtable.Schema) (AccountTable, error) {
	table := db.GetTable(&Account{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Account{}).ProtoReflect().Descriptor().FullName()))
	}
	return accountTable{table.(ormtable.AutoIncrementTable)}, nil
}

type BlockchainTable interface {
	Insert(ctx context.Context, blockchain *Blockchain) error
	InsertReturningIndex(ctx context.Context, blockchain *Blockchain) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, blockchain *Blockchain) error
	Save(ctx context.Context, blockchain *Blockchain) error
	Delete(ctx context.Context, blockchain *Blockchain) error
	Has(ctx context.Context, index uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, index uint64) (*Blockchain, error)
	HasByChainId(ctx context.Context, chain_id string) (found bool, err error)
	// GetByChainId returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByChainId(ctx context.Context, chain_id string) (*Blockchain, error)
	HasByName(ctx context.Context, name string) (found bool, err error)
	// GetByName returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByName(ctx context.Context, name string) (*Blockchain, error)
	List(ctx context.Context, prefixKey BlockchainIndexKey, opts ...ormlist.Option) (BlockchainIterator, error)
	ListRange(ctx context.Context, from, to BlockchainIndexKey, opts ...ormlist.Option) (BlockchainIterator, error)
	DeleteBy(ctx context.Context, prefixKey BlockchainIndexKey) error
	DeleteRange(ctx context.Context, from, to BlockchainIndexKey) error

	doNotImplement()
}

type BlockchainIterator struct {
	ormtable.Iterator
}

func (i BlockchainIterator) Value() (*Blockchain, error) {
	var blockchain Blockchain
	err := i.UnmarshalMessage(&blockchain)
	return &blockchain, err
}

type BlockchainIndexKey interface {
	id() uint32
	values() []interface{}
	blockchainIndexKey()
}

// primary key starting index..
type BlockchainPrimaryKey = BlockchainIndexIndexKey

type BlockchainIndexIndexKey struct {
	vs []interface{}
}

func (x BlockchainIndexIndexKey) id() uint32            { return 0 }
func (x BlockchainIndexIndexKey) values() []interface{} { return x.vs }
func (x BlockchainIndexIndexKey) blockchainIndexKey()   {}

func (this BlockchainIndexIndexKey) WithIndex(index uint64) BlockchainIndexIndexKey {
	this.vs = []interface{}{index}
	return this
}

type BlockchainChainIdIndexKey struct {
	vs []interface{}
}

func (x BlockchainChainIdIndexKey) id() uint32            { return 1 }
func (x BlockchainChainIdIndexKey) values() []interface{} { return x.vs }
func (x BlockchainChainIdIndexKey) blockchainIndexKey()   {}

func (this BlockchainChainIdIndexKey) WithChainId(chain_id string) BlockchainChainIdIndexKey {
	this.vs = []interface{}{chain_id}
	return this
}

type BlockchainNameIndexKey struct {
	vs []interface{}
}

func (x BlockchainNameIndexKey) id() uint32            { return 2 }
func (x BlockchainNameIndexKey) values() []interface{} { return x.vs }
func (x BlockchainNameIndexKey) blockchainIndexKey()   {}

func (this BlockchainNameIndexKey) WithName(name string) BlockchainNameIndexKey {
	this.vs = []interface{}{name}
	return this
}

type blockchainTable struct {
	table ormtable.AutoIncrementTable
}

func (this blockchainTable) Insert(ctx context.Context, blockchain *Blockchain) error {
	return this.table.Insert(ctx, blockchain)
}

func (this blockchainTable) Update(ctx context.Context, blockchain *Blockchain) error {
	return this.table.Update(ctx, blockchain)
}

func (this blockchainTable) Save(ctx context.Context, blockchain *Blockchain) error {
	return this.table.Save(ctx, blockchain)
}

func (this blockchainTable) Delete(ctx context.Context, blockchain *Blockchain) error {
	return this.table.Delete(ctx, blockchain)
}

func (this blockchainTable) InsertReturningIndex(ctx context.Context, blockchain *Blockchain) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, blockchain)
}

func (this blockchainTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this blockchainTable) Has(ctx context.Context, index uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, index)
}

func (this blockchainTable) Get(ctx context.Context, index uint64) (*Blockchain, error) {
	var blockchain Blockchain
	found, err := this.table.PrimaryKey().Get(ctx, &blockchain, index)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &blockchain, nil
}

func (this blockchainTable) HasByChainId(ctx context.Context, chain_id string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		chain_id,
	)
}

func (this blockchainTable) GetByChainId(ctx context.Context, chain_id string) (*Blockchain, error) {
	var blockchain Blockchain
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &blockchain,
		chain_id,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &blockchain, nil
}

func (this blockchainTable) HasByName(ctx context.Context, name string) (found bool, err error) {
	return this.table.GetIndexByID(2).(ormtable.UniqueIndex).Has(ctx,
		name,
	)
}

func (this blockchainTable) GetByName(ctx context.Context, name string) (*Blockchain, error) {
	var blockchain Blockchain
	found, err := this.table.GetIndexByID(2).(ormtable.UniqueIndex).Get(ctx, &blockchain,
		name,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &blockchain, nil
}

func (this blockchainTable) List(ctx context.Context, prefixKey BlockchainIndexKey, opts ...ormlist.Option) (BlockchainIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return BlockchainIterator{it}, err
}

func (this blockchainTable) ListRange(ctx context.Context, from, to BlockchainIndexKey, opts ...ormlist.Option) (BlockchainIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return BlockchainIterator{it}, err
}

func (this blockchainTable) DeleteBy(ctx context.Context, prefixKey BlockchainIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this blockchainTable) DeleteRange(ctx context.Context, from, to BlockchainIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this blockchainTable) doNotImplement() {}

var _ BlockchainTable = blockchainTable{}

func NewBlockchainTable(db ormtable.Schema) (BlockchainTable, error) {
	table := db.GetTable(&Blockchain{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Blockchain{}).ProtoReflect().Descriptor().FullName()))
	}
	return blockchainTable{table.(ormtable.AutoIncrementTable)}, nil
}

type AccumulatorTable interface {
	Insert(ctx context.Context, accumulator *Accumulator) error
	InsertReturningIndex(ctx context.Context, accumulator *Accumulator) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, accumulator *Accumulator) error
	Save(ctx context.Context, accumulator *Accumulator) error
	Delete(ctx context.Context, accumulator *Accumulator) error
	Has(ctx context.Context, index uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, index uint64) (*Accumulator, error)
	HasByControllerKey(ctx context.Context, controller string, key string) (found bool, err error)
	// GetByControllerKey returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByControllerKey(ctx context.Context, controller string, key string) (*Accumulator, error)
	List(ctx context.Context, prefixKey AccumulatorIndexKey, opts ...ormlist.Option) (AccumulatorIterator, error)
	ListRange(ctx context.Context, from, to AccumulatorIndexKey, opts ...ormlist.Option) (AccumulatorIterator, error)
	DeleteBy(ctx context.Context, prefixKey AccumulatorIndexKey) error
	DeleteRange(ctx context.Context, from, to AccumulatorIndexKey) error

	doNotImplement()
}

type AccumulatorIterator struct {
	ormtable.Iterator
}

func (i AccumulatorIterator) Value() (*Accumulator, error) {
	var accumulator Accumulator
	err := i.UnmarshalMessage(&accumulator)
	return &accumulator, err
}

type AccumulatorIndexKey interface {
	id() uint32
	values() []interface{}
	accumulatorIndexKey()
}

// primary key starting index..
type AccumulatorPrimaryKey = AccumulatorIndexIndexKey

type AccumulatorIndexIndexKey struct {
	vs []interface{}
}

func (x AccumulatorIndexIndexKey) id() uint32            { return 0 }
func (x AccumulatorIndexIndexKey) values() []interface{} { return x.vs }
func (x AccumulatorIndexIndexKey) accumulatorIndexKey()  {}

func (this AccumulatorIndexIndexKey) WithIndex(index uint64) AccumulatorIndexIndexKey {
	this.vs = []interface{}{index}
	return this
}

type AccumulatorControllerKeyIndexKey struct {
	vs []interface{}
}

func (x AccumulatorControllerKeyIndexKey) id() uint32            { return 1 }
func (x AccumulatorControllerKeyIndexKey) values() []interface{} { return x.vs }
func (x AccumulatorControllerKeyIndexKey) accumulatorIndexKey()  {}

func (this AccumulatorControllerKeyIndexKey) WithController(controller string) AccumulatorControllerKeyIndexKey {
	this.vs = []interface{}{controller}
	return this
}

func (this AccumulatorControllerKeyIndexKey) WithControllerKey(controller string, key string) AccumulatorControllerKeyIndexKey {
	this.vs = []interface{}{controller, key}
	return this
}

type accumulatorTable struct {
	table ormtable.AutoIncrementTable
}

func (this accumulatorTable) Insert(ctx context.Context, accumulator *Accumulator) error {
	return this.table.Insert(ctx, accumulator)
}

func (this accumulatorTable) Update(ctx context.Context, accumulator *Accumulator) error {
	return this.table.Update(ctx, accumulator)
}

func (this accumulatorTable) Save(ctx context.Context, accumulator *Accumulator) error {
	return this.table.Save(ctx, accumulator)
}

func (this accumulatorTable) Delete(ctx context.Context, accumulator *Accumulator) error {
	return this.table.Delete(ctx, accumulator)
}

func (this accumulatorTable) InsertReturningIndex(ctx context.Context, accumulator *Accumulator) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, accumulator)
}

func (this accumulatorTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this accumulatorTable) Has(ctx context.Context, index uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, index)
}

func (this accumulatorTable) Get(ctx context.Context, index uint64) (*Accumulator, error) {
	var accumulator Accumulator
	found, err := this.table.PrimaryKey().Get(ctx, &accumulator, index)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &accumulator, nil
}

func (this accumulatorTable) HasByControllerKey(ctx context.Context, controller string, key string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		controller,
		key,
	)
}

func (this accumulatorTable) GetByControllerKey(ctx context.Context, controller string, key string) (*Accumulator, error) {
	var accumulator Accumulator
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &accumulator,
		controller,
		key,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &accumulator, nil
}

func (this accumulatorTable) List(ctx context.Context, prefixKey AccumulatorIndexKey, opts ...ormlist.Option) (AccumulatorIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return AccumulatorIterator{it}, err
}

func (this accumulatorTable) ListRange(ctx context.Context, from, to AccumulatorIndexKey, opts ...ormlist.Option) (AccumulatorIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return AccumulatorIterator{it}, err
}

func (this accumulatorTable) DeleteBy(ctx context.Context, prefixKey AccumulatorIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this accumulatorTable) DeleteRange(ctx context.Context, from, to AccumulatorIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this accumulatorTable) doNotImplement() {}

var _ AccumulatorTable = accumulatorTable{}

func NewAccumulatorTable(db ormtable.Schema) (AccumulatorTable, error) {
	table := db.GetTable(&Accumulator{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Accumulator{}).ProtoReflect().Descriptor().FullName()))
	}
	return accumulatorTable{table.(ormtable.AutoIncrementTable)}, nil
}

type ControllerTable interface {
	Insert(ctx context.Context, controller *Controller) error
	InsertReturningSequence(ctx context.Context, controller *Controller) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, controller *Controller) error
	Save(ctx context.Context, controller *Controller) error
	Delete(ctx context.Context, controller *Controller) error
	Has(ctx context.Context, sequence uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, sequence uint64) (*Controller, error)
	HasByAddress(ctx context.Context, address string) (found bool, err error)
	// GetByAddress returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByAddress(ctx context.Context, address string) (*Controller, error)
	HasByPublicKey(ctx context.Context, public_key []byte) (found bool, err error)
	// GetByPublicKey returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByPublicKey(ctx context.Context, public_key []byte) (*Controller, error)
	HasByPeerId(ctx context.Context, peer_id string) (found bool, err error)
	// GetByPeerId returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByPeerId(ctx context.Context, peer_id string) (*Controller, error)
	HasByIpns(ctx context.Context, ipns string) (found bool, err error)
	// GetByIpns returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByIpns(ctx context.Context, ipns string) (*Controller, error)
	List(ctx context.Context, prefixKey ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error)
	ListRange(ctx context.Context, from, to ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error)
	DeleteBy(ctx context.Context, prefixKey ControllerIndexKey) error
	DeleteRange(ctx context.Context, from, to ControllerIndexKey) error

	doNotImplement()
}

type ControllerIterator struct {
	ormtable.Iterator
}

func (i ControllerIterator) Value() (*Controller, error) {
	var controller Controller
	err := i.UnmarshalMessage(&controller)
	return &controller, err
}

type ControllerIndexKey interface {
	id() uint32
	values() []interface{}
	controllerIndexKey()
}

// primary key starting index..
type ControllerPrimaryKey = ControllerSequenceIndexKey

type ControllerSequenceIndexKey struct {
	vs []interface{}
}

func (x ControllerSequenceIndexKey) id() uint32            { return 0 }
func (x ControllerSequenceIndexKey) values() []interface{} { return x.vs }
func (x ControllerSequenceIndexKey) controllerIndexKey()   {}

func (this ControllerSequenceIndexKey) WithSequence(sequence uint64) ControllerSequenceIndexKey {
	this.vs = []interface{}{sequence}
	return this
}

type ControllerAddressIndexKey struct {
	vs []interface{}
}

func (x ControllerAddressIndexKey) id() uint32            { return 1 }
func (x ControllerAddressIndexKey) values() []interface{} { return x.vs }
func (x ControllerAddressIndexKey) controllerIndexKey()   {}

func (this ControllerAddressIndexKey) WithAddress(address string) ControllerAddressIndexKey {
	this.vs = []interface{}{address}
	return this
}

type ControllerPublicKeyIndexKey struct {
	vs []interface{}
}

func (x ControllerPublicKeyIndexKey) id() uint32            { return 2 }
func (x ControllerPublicKeyIndexKey) values() []interface{} { return x.vs }
func (x ControllerPublicKeyIndexKey) controllerIndexKey()   {}

func (this ControllerPublicKeyIndexKey) WithPublicKey(public_key []byte) ControllerPublicKeyIndexKey {
	this.vs = []interface{}{public_key}
	return this
}

type ControllerPeerIdIndexKey struct {
	vs []interface{}
}

func (x ControllerPeerIdIndexKey) id() uint32            { return 3 }
func (x ControllerPeerIdIndexKey) values() []interface{} { return x.vs }
func (x ControllerPeerIdIndexKey) controllerIndexKey()   {}

func (this ControllerPeerIdIndexKey) WithPeerId(peer_id string) ControllerPeerIdIndexKey {
	this.vs = []interface{}{peer_id}
	return this
}

type ControllerIpnsIndexKey struct {
	vs []interface{}
}

func (x ControllerIpnsIndexKey) id() uint32            { return 4 }
func (x ControllerIpnsIndexKey) values() []interface{} { return x.vs }
func (x ControllerIpnsIndexKey) controllerIndexKey()   {}

func (this ControllerIpnsIndexKey) WithIpns(ipns string) ControllerIpnsIndexKey {
	this.vs = []interface{}{ipns}
	return this
}

type controllerTable struct {
	table ormtable.AutoIncrementTable
}

func (this controllerTable) Insert(ctx context.Context, controller *Controller) error {
	return this.table.Insert(ctx, controller)
}

func (this controllerTable) Update(ctx context.Context, controller *Controller) error {
	return this.table.Update(ctx, controller)
}

func (this controllerTable) Save(ctx context.Context, controller *Controller) error {
	return this.table.Save(ctx, controller)
}

func (this controllerTable) Delete(ctx context.Context, controller *Controller) error {
	return this.table.Delete(ctx, controller)
}

func (this controllerTable) InsertReturningSequence(ctx context.Context, controller *Controller) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, controller)
}

func (this controllerTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this controllerTable) Has(ctx context.Context, sequence uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, sequence)
}

func (this controllerTable) Get(ctx context.Context, sequence uint64) (*Controller, error) {
	var controller Controller
	found, err := this.table.PrimaryKey().Get(ctx, &controller, sequence)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) HasByAddress(ctx context.Context, address string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		address,
	)
}

func (this controllerTable) GetByAddress(ctx context.Context, address string) (*Controller, error) {
	var controller Controller
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &controller,
		address,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) HasByPublicKey(ctx context.Context, public_key []byte) (found bool, err error) {
	return this.table.GetIndexByID(2).(ormtable.UniqueIndex).Has(ctx,
		public_key,
	)
}

func (this controllerTable) GetByPublicKey(ctx context.Context, public_key []byte) (*Controller, error) {
	var controller Controller
	found, err := this.table.GetIndexByID(2).(ormtable.UniqueIndex).Get(ctx, &controller,
		public_key,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) HasByPeerId(ctx context.Context, peer_id string) (found bool, err error) {
	return this.table.GetIndexByID(3).(ormtable.UniqueIndex).Has(ctx,
		peer_id,
	)
}

func (this controllerTable) GetByPeerId(ctx context.Context, peer_id string) (*Controller, error) {
	var controller Controller
	found, err := this.table.GetIndexByID(3).(ormtable.UniqueIndex).Get(ctx, &controller,
		peer_id,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) HasByIpns(ctx context.Context, ipns string) (found bool, err error) {
	return this.table.GetIndexByID(4).(ormtable.UniqueIndex).Has(ctx,
		ipns,
	)
}

func (this controllerTable) GetByIpns(ctx context.Context, ipns string) (*Controller, error) {
	var controller Controller
	found, err := this.table.GetIndexByID(4).(ormtable.UniqueIndex).Get(ctx, &controller,
		ipns,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &controller, nil
}

func (this controllerTable) List(ctx context.Context, prefixKey ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ControllerIterator{it}, err
}

func (this controllerTable) ListRange(ctx context.Context, from, to ControllerIndexKey, opts ...ormlist.Option) (ControllerIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ControllerIterator{it}, err
}

func (this controllerTable) DeleteBy(ctx context.Context, prefixKey ControllerIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this controllerTable) DeleteRange(ctx context.Context, from, to ControllerIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this controllerTable) doNotImplement() {}

var _ ControllerTable = controllerTable{}

func NewControllerTable(db ormtable.Schema) (ControllerTable, error) {
	table := db.GetTable(&Controller{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Controller{}).ProtoReflect().Descriptor().FullName()))
	}
	return controllerTable{table.(ormtable.AutoIncrementTable)}, nil
}

type StateStore interface {
	AccountTable() AccountTable
	BlockchainTable() BlockchainTable
	AccumulatorTable() AccumulatorTable
	ControllerTable() ControllerTable

	doNotImplement()
}

type stateStore struct {
	account     AccountTable
	blockchain  BlockchainTable
	accumulator AccumulatorTable
	controller  ControllerTable
}

func (x stateStore) AccountTable() AccountTable {
	return x.account
}

func (x stateStore) BlockchainTable() BlockchainTable {
	return x.blockchain
}

func (x stateStore) AccumulatorTable() AccumulatorTable {
	return x.accumulator
}

func (x stateStore) ControllerTable() ControllerTable {
	return x.controller
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	accountTable, err := NewAccountTable(db)
	if err != nil {
		return nil, err
	}

	blockchainTable, err := NewBlockchainTable(db)
	if err != nil {
		return nil, err
	}

	accumulatorTable, err := NewAccumulatorTable(db)
	if err != nil {
		return nil, err
	}

	controllerTable, err := NewControllerTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		accountTable,
		blockchainTable,
		accumulatorTable,
		controllerTable,
	}, nil
}
