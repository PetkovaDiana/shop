package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/PetkovaDiana/shop/internal/repository/entities"
	repoErrors "github.com/PetkovaDiana/shop/internal/repository/errors"
	domainModels "github.com/PetkovaDiana/shop/internal/service/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type auth struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) Authorization {
	return &auth{db: db}
}

func (a *auth) CreateClient(ctx context.Context, client domainModels.CreateClient) error {
	checkQuery := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(
			"true",
		).
		From(fmt.Sprintf("%s AS cl", entities.Table_Client)).
		Where(squirrel.Eq{fmt.Sprintf("cl.%s", entities.Field_Client_Email): client.Email})

	checkSql, checkArgs, err := checkQuery.ToSql()
	if err != nil {
		return fmt.Errorf("error building check sql query: %v", err)
	}

	var isExists bool
	err = a.db.QueryRow(ctx, checkSql, checkArgs...).Scan(&isExists)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("error checking for existing email: %v", err)
	}

	if isExists {
		return fmt.Errorf("email %s already in use", client.Email)
	}

	query := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert(entities.Table_Client).
		Columns(
			entities.Field_Client_Name,
			entities.Field_Client_Last_Name,
			entities.Field_Client_Number,
			entities.Field_Client_Password,
			entities.Field_Client_Email).
		Values(
			client.Name,
			client.LastName,
			client.Number,
			client.PasswordHashed,
			client.Email,
		)

	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("error building sql query: %v", err)
	}

	_, err = a.db.Exec(ctx, sql, args...)
	return err

}

func (a *auth) GetClient(ctx context.Context, email string) (*domainModels.Client, error) {
	var resultClient domainModels.Client
	query := squirrel.Select(
		fmt.Sprintf("cl.%s, cl.%s, cl.%s, cl.%s, cl.%s, cl.%s",
			entities.Field_Client_ID,
			entities.Field_Client_Name,
			entities.Field_Client_Last_Name,
			entities.Field_Client_Number,
			entities.Field_Client_Password,
			entities.Field_Client_Email,
		),
	).From(fmt.Sprintf("%s AS cl", entities.Table_Client)).
		Where(squirrel.Eq{fmt.Sprintf("cl.%s", entities.Field_Client_Email): email}).
		Limit(1)

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	if err = a.db.QueryRow(ctx, sql, args...).Scan(
		&resultClient.ID, &resultClient.Name, &resultClient.LastName,
		&resultClient.Number, &resultClient.PasswordHashed,
		&resultClient.Email); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repoErrors.ErrClientNotFound{ClientEmail: email}
		} else {
			return nil, err
		}
	}

	return &resultClient, nil

}
