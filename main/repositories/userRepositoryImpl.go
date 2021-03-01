package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"testMekarApp/main/models"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func (s UserRepositoryImpl) SelectUsers() ([]*models.User, error) {
	queryIn := "SELECT usr.id, usr.nama, usr.tgl_lahir, usr.no_ktp, prf.jenis_pekerjaan, edu.jenis_pendidikan from pengguna usr join pekerjaan prf on usr.id_pekerjaan = prf.id join pendidikan edu on usr.id_pendidikan = edu.id;"
	data, err := s.db.Query(queryIn)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var result = []*models.User{}
	for data.Next() {
		var user = models.User{}
		var err = data.Scan(&user.ID, &user.Name, &user.DateBirth, &user.IdCardNumber, &user.Profession, &user.Education)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, &user)
	}
	if err = data.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s UserRepositoryImpl) InsertUser(user *models.User) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := "insert into pengguna(nama, tgl_lahir,no_ktp,id_pekerjaan,id_pendidikan) values(?,?,?);"
	res, err := tx.Exec(query, user.Name, user.DateBirth, user.IdCardNumber, user.IdProfession, user.IdEducation)
	if err != nil {
		tx.Rollback()
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Sprint(lastId)
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
	// err = s.db.QueryRow("Select jenis_pekerjaan from pekerjaan where id = ?", user.IdProfession)

}
func (s UserRepositoryImpl) SelectUserById(id int) (*models.User, error) {
	var user = new(models.User)
	err := s.db.QueryRow("SELECT usr.id, usr.nama, usr.tgl_lahir, usr.no_ktp, prf.jenis_pekerjaan, edu.jenis_pendidikan from pengguna usr join pekerjaan prf on usr.id_pekerjaan = prf.id join pendidikan edu on usr.id_pendidikan = edu.id where usr.id=? ;", id).Scan(&user.ID, &user.Name, &user.DateBirth, &user.IdCardNumber, &user.Profession, &user.Education)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s UserRepositoryImpl) UpdateUser(user *models.User) error {
	tx, err := s.db.Begin()
	query := "update pengguna set nama = ?, tgl_lahir =?, no_ktp=?, id_pekerjaan=?, id_pendidikan=? where id=?"
	_, err = tx.Exec(query, user.Name, user.DateBirth, user.IdCardNumber, user.IdEducation, user.IdEducation, user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	// userUpdate, err := s.SelectUserById(user.ID)
	return nil
}

func (s UserRepositoryImpl) DeleteUser(id int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := "delete from pengguna where id = ?"
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}
func InitUserRepositoryImpl(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db}
}
