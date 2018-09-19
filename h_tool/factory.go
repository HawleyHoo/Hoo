package h_tool

import "fmt"

type GirlFriend struct {
	nationality string
	eyesColor   string
	language    string
}

type AbstractFactory interface {
	CreateMyLove() GirlFriend
}

type IndianGirlFriendFactory struct {
}

type KoreanGirlFriendFactory struct {
}

func (a IndianGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Indian", "Black", "Hindi"}
}

func (a KoreanGirlFriendFactory) CreateMyLove() GirlFriend {
	return GirlFriend{"Korean", "Brown", "Korean"}
}

func getGirlFriend(typeGf string) GirlFriend {

	var gffact AbstractFactory
	switch typeGf {
	case "Indian":
		gffact = IndianGirlFriendFactory{}
		return gffact.CreateMyLove()
	case "Korean":
		gffact = KoreanGirlFriendFactory{}
		return gffact.CreateMyLove()
	}
	return GirlFriend{}
}

func maintest() {

	a := getGirlFriend("Indian")

	fmt.Println(a.eyesColor)
}

/*
var UserNotFoundError = errors.New("user not found")

type DataStore interface {
	Name() string
	FindUserNameById(id int64) (string, error)
}

//The first implementation.
type PostgreSQLDataStore struct {
	DSN string
	DB  sql.DB
}

func (pds *PostgreSQLDataStore) Name() string {
	return "PostgreSQLDataStore"
}

func (pds *PostgreSQLDataStore) FindUserNameById(id int64) (string, error) {
	var username string
	err := pds.DB.Query("SELECT username FROM users WHERE id=$1", id).Scan(&username)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", UserNotFoundError
		}
		return "", err
	}
	return username, nil
}

//The second implementation.
type MemoryDataStore struct {
	sync.RWMutex
	Users map[int64]string
}

func (mds *MemoryDataStore) Name() string {
	return "MemoryDataStore"
}

func (mds *MemoryDataStore) FindUserNameById(id int64) (string, error) {
	mds.RWMutex.RLock()
	defer mds.RWMutex.RUnlock()
	username, ok := mds.Users[id];
	if !ok {
		return "", UserNotFoundError
	}
	return username, nil
}

type DataStoreFactory func(conf map[string]string) (DataStore, error)

func NewPostgreSQLDataStore(conf map[string]string) (DataStore, error) {
	dsn, ok := conf.Get("DATASTORE_POSTGRES_DSN", "")
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s is required for the postgres datastore", "DATASTORE_POSTGRES_DSN"))
	}

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Panicf("Failed to connect to datastore: %s", err.Error())
		return nil, datastore.FailedToConnect
	}

	return &PostgresDataStore{
		DSN: dsn,
		DB:  db,
	}, nil
}

func NewMemoryDataStore(conf map[string]string) (DataStore, error) {
	return &MemoryDataStore{
		Users: &map[int64]string{
			1: "mnbbrown",
			0: "root",
		},
		RWMutex: &sync.RWMutex{},
	}, nil
}

var datastoreFactories = make(map[string]DataStoreFactory)

func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Panicf("Datastore factory %s does not exist.", name)
	}
	_, registered := datastoreFactories[name]
	if registered {
		log.Errorf("Datastore factory %s already registered. Ignoring.", name)
	}
	datastoreFactories[name] = factory
}

func init() {
	Register("postgres", NewPostgreSQLDataStore)
	Register("memory", NewMemoryDataStore)
}

func CreateDataStore(conf map[string]string) (DataStore, error) {

	// Query configuration for datastore defaulting to "memory".
	engineName := conf.Get("DATASTORE", "memory")

	engineFactory, ok := datastoreFactories[engineName]
	if !ok {
		// Factory has not been registered.
		// Make a list of all available datastore factories for logging.
		availableDatastores := make([]string, len(datastoreFactories))
		for k, _ := range datastoreFactories {
			availableDatastores = append(availableDatastores, k)
		}
		return nil, errors.New(fmt.Sprintf("Invalid Datastore name. Must be one of: %s", strings.Join(availableDatastores, ", ")))
	}

	// Run the factory with the configuration.
	return engineFactory(conf)
}

func test() {
	datastore, err := CreateDataStore(&map[string]string{
		"DATASTORE":              "postgres",
		"DATASTORE_POSTGRES_DSN": "dbname=factoriesareamazing",
	})

	datastore, err = CreateDataStore(&map[string]string{
		"DATASTORE": "memory",
	})
}
*/