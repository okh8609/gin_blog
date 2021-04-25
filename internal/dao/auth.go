package dao

import "github.com/okh8609/gin_blog/internal/model"

func (d *Dao) GetAuth(uuid string) (model.Auth, error) {
	auth := model.Auth{UUID: uuid}
	return auth.Get(d.gormDB)
}

func (d *Dao) CreateAuth(uuid string, password string) error {
	auth := model.Auth{
		UUID:     uuid,
		Password: password,
	}
	return auth.Create(d.gormDB)
}

func (d *Dao) UpdateAuth(uuid string, password string) error {
	auth := model.Auth{
		UUID: uuid,
	}
	values := map[string]interface{}{
		"password": password,
	}
	return auth.Update(d.gormDB, values)
}

func (d *Dao) DeleteAuth(uuid string) error {
	auth := model.Auth{UUID: uuid}
	return auth.Delete(d.gormDB)
}
