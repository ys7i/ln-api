package daocore

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// User represents a row from 'db.users'.
type User struct {
	ID           int    `json:"id"`            // id
	Name         string `json:"name"`          // name
	Email        string `json:"email"`         // email
	PasswordHash string `json:"password_hash"` // password_hash
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted returns true when the User has been marked for deletion from
// the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(ctx context.Context, db DB) error {
	switch {
	case u._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case u._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO db.users (` +
		`name, email, password_hash` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`
	// run
	logf(sqlstr, u.Name, u.Email, u.PasswordHash)
	res, err := db.ExecContext(ctx, sqlstr, u.Name, u.Email, u.PasswordHash)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	u.ID = int(id)
	// set exists
	u._exists = true
	return nil
}

// Update updates a User in the database.
func (u *User) Update(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case u._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE db.users SET ` +
		`name = ?, email = ?, password_hash = ? ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, u.Name, u.Email, u.PasswordHash, u.ID)
	if _, err := db.ExecContext(ctx, sqlstr, u.Name, u.Email, u.PasswordHash, u.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the User to the database.
func (u *User) Save(ctx context.Context, db DB) error {
	if u.Exists() {
		return u.Update(ctx, db)
	}
	return u.Insert(ctx, db)
}

// Upsert performs an upsert for User.
func (u *User) Upsert(ctx context.Context, db DB) error {
	switch {
	case u._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO db.users (` +
		`id, name, email, password_hash` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`name = VALUES(name), email = VALUES(email), password_hash = VALUES(password_hash)`
	// run
	logf(sqlstr, u.ID, u.Name, u.Email, u.PasswordHash)
	if _, err := db.ExecContext(ctx, sqlstr, u.ID, u.Name, u.Email, u.PasswordHash); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Delete deletes the User from the database.
func (u *User) Delete(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return nil
	case u._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM db.users ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, u.ID)
	if _, err := db.ExecContext(ctx, sqlstr, u.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	u._deleted = true
	return nil
}

// UserByEmail retrieves a row from 'db.users' as a User.
//
// Generated from index 'email'.
func UserByEmail(ctx context.Context, db DB, email string) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, email, password_hash ` +
		`FROM db.users ` +
		`WHERE email = ?`
	// run
	logf(sqlstr, email)
	u := User{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, email).Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}

// UserByID retrieves a row from 'db.users' as a User.
//
// Generated from index 'users_id_pkey'.
func UserByID(ctx context.Context, db DB, id int) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, email, password_hash ` +
		`FROM db.users ` +
		`WHERE id = ?`
	// run
	logf(sqlstr, id)
	u := User{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}
