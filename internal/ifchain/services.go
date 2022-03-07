package ifchain

import (
	"log"
	"os"

	"github.com/ruckuus/ifchain/pkg/dummydb"
)

type ServiceIface interface {
	Save(*Object) error
}

type serviceImpl struct {
	ServiceIface

	log *log.Logger
}

type serviceValidator struct {
	ServiceIface
	log *log.Logger
}

type serviceCache struct {
	ServiceIface

	cache *dummydb.DB
	log   *log.Logger
}

type serviceOrm struct {
	db  *dummydb.DB
	log *log.Logger
}

func NewService(db *dummydb.DB, cache *dummydb.DB) ServiceIface {
	log := log.New(os.Stdout, "svc: ", log.Lshortfile|log.Lmicroseconds)
	svc := &serviceImpl{
		ServiceIface: &serviceValidator{
			ServiceIface: &serviceCache{
				ServiceIface: &serviceOrm{
					db:  db,
					log: log,
				},
				cache: cache,
				log:   log,
			},
			log: log,
		},
		log: log,
	}

	return svc
}
