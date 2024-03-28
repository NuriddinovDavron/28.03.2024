package kv

import (
	"database/sql"

	"github.com/spf13/cast"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) Set(key, value string, seconds int) error {
	query := `insert into tests(id, model) values($1, 42)`

	if _, err := p.db.Exec(query, key, value); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Get(key string) (string, error) {
	query := `select id, model from tests where id=$1`
	var resp string
	if err := p.db.QueryRow(query, key).Scan(&resp); err != nil {
		return "", err
	}

	return resp, nil
}

func (p *Postgres) Delete(key string) error {
	query := `delete from tests where id = $1`
	if _, err := p.db.Exec(query, key); err != nil{
		return err
	}
	return nil
}

func (p *Postgres) List() (map[string]string, error){
	pairs := make(map[string]string)

	query := `select id, model from tests`

	cursor, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	n := 1
	for cursor.Next(){
		var res string
		if err := cursor.Scan(&res); err != nil {
			return nil, err
		}

		pairs[cast.ToString(n)] = res
		n++
	}

	return pairs, nil
}
