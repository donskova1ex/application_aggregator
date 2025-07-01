package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/donskova1ex/application_aggregator/internal"
	"github.com/donskova1ex/application_aggregator/internal/domain"
	"github.com/donskova1ex/application_aggregator/tools"
	"github.com/google/uuid"
	"time"
)

func (repo *PostgresRepository) CreateOrganization(ctx context.Context, organization *domain.Organization) (*domain.Organization, error) {
	query := `INSERT INTO organizations(uuid, name) VALUES ($1, $2) ON CONFLICT ON CONSTRAINT organizations_name_key DO NOTHING RETURNING id`

	newUUID := uuid.NewString()
	result, err := repo.db.ExecContext(ctx, query, newUUID, organization.Name)
	if err != nil {
		return nil, fmt.Errorf("error creating organization: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("organization with name [%s] already exists", organization.Name)
	}

	newOrganization := &domain.Organization{
		Uuid: newUUID,
		Name: organization.Name,
	}
	return newOrganization, nil
}

func (repo *PostgresRepository) GetOrganizationByUUID(ctx context.Context, uuid string) (*domain.Organization, error) {
	if !tools.ValidUUID(uuid) {
		return nil, fmt.Errorf("invalid organization uuid: %w", internal.ErrUUIDValidation)
	}

	query := `SELECT uuid, name FROM organizations WHERE uuid = $1`

	organization := &domain.Organization{}

	err := repo.db.GetContext(ctx, organization, query, uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("organization with uuid not found: %w", internal.ErrRecordNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("error getting organization by uuid: %w", err)
	}

	return organization, nil
}

func (repo *PostgresRepository) DeleteOrganizationByUUID(ctx context.Context, uuid string) error {
	if !tools.ValidUUID(uuid) {
		return fmt.Errorf("invalid organization uuid: %w", internal.ErrUUIDValidation)
	}

	querty := `DELETE FROM organizations WHERE uuid = $1`

	result, err := repo.db.ExecContext(ctx, querty, uuid)

	if err != nil {
		return fmt.Errorf("error deleting organization with this uuid: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("organization with uuid not found: %w", internal.ErrRecordNotFound)
	}

	return nil
}
func (repo *PostgresRepository) UpdateOrganization(ctx context.Context, uuid string, organization *domain.Organization) (*domain.Organization, error) {
	if !tools.ValidUUID(uuid) {
		return nil, fmt.Errorf("invalid organization uuid: %w", internal.ErrUUIDValidation)
	}

	query := `UPDATE organizations SET name = $1, updated_at = $2 WHERE uuid = $3`

	updateTime := time.Now()

	result, err := repo.db.ExecContext(ctx, query, organization.Name, updateTime, uuid)
	if err != nil {
		return nil, fmt.Errorf("error updating organization: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("organization with uuid not found: %w", internal.ErrRecordNotFound)
	}

	return organization, nil
}
func (repo *PostgresRepository) GetOrganizations(ctx context.Context) ([]*domain.Organization, error) {
	var organizations []*domain.Organization

	query := `SELECT uuid, name FROM organizations`
	err := repo.db.SelectContext(ctx, &organizations, query)
	if errors.Is(err, sql.ErrNoRows) {
		return organizations, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error getting organizations: %w", err)
	}

	return organizations, nil
}
