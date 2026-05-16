package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryFactory func(tx *sql.Tx) any

type UowInterface interface {
	Register(name string, rc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (any, error)
	Do(ctx context.Context, fn func(uow *UoW) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string) error
}

type UoW struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUoW(ctx context.Context, db *sql.DB) *UoW {
	return &UoW{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *UoW) Register(name string, rf RepositoryFactory) {
	u.Repositories[name] = rf
}

func (u *UoW) UnRegister(name string) error {
	delete(u.Repositories, name)
	return nil
}

func (u *UoW) GetRepository(ctx context.Context, name string) (any, error) {
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}

		u.Tx = tx
	}

	repo := u.Repositories[name](u.Tx)
	return repo, nil
}

func (u *UoW) Do(ctx context.Context, fn func(uow *UoW) error) error {
	if u.Tx != nil {
		return fmt.Errorf("transaction already started")
	}

	tx, err := u.Db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	u.Tx = tx

	err = fn(u)

	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return errors.New(fmt.Sprintf("Original error: %s, rollback error: %s", err.Error(), errRb.Error()))
		}
		return err
	}

	return u.CommitOrRollback()
}

func (u *UoW) Rollback() error {
	if u.Tx == nil {
		return errors.New("No transaction to rollback")
	}

	err := u.Tx.Rollback()
	if err != nil {
		return err
	}

	u.Tx = nil

	return nil
}

func (u *UoW) CommitOrRollback() error {
	err := u.Tx.Commit()
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return errors.New(fmt.Sprintf("Original error: %s, rollback error: %s", err.Error(), errRb.Error()))
		}
		return err
	}

	return nil
}
