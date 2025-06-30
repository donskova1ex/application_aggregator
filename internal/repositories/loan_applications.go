package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/donskova1ex/application_aggregator/internal"
	"github.com/donskova1ex/application_aggregator/internal/domain"
	"github.com/donskova1ex/application_aggregator/tools"
)

func (repo *PostgresRepository) LoanApplications(ctx context.Context) ([]*domain.LoanApplication, error) {
	query := `SELECT uuid, value, phone, incoming_organization_uuid FROM loan_applications`

	var loanApplications []*domain.LoanApplication
	err := repo.db.SelectContext(ctx, &loanApplications, query)
	if errors.Is(err, sql.ErrNoRows) {
		return loanApplications, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error getting loan applications: %w", err)
	}
	return loanApplications, nil
}

func (repo *PostgresRepository) GetLoanApplicationsByUUID(ctx context.Context, uuid string) (*domain.LoanApplication, error) {
	if !tools.ValidUUID(uuid) {
		return nil, internal.UUIDValidationFailed
	}

	query := `SELECT uuid, value, phone, incoming_organization_uuid FROM loan_applications WHERE uuid=$1`
	loanApplication := &domain.LoanApplication{}
	err := repo.db.GetContext(ctx, loanApplication, query, uuid)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf(`loan application not found: %w`, internal.ErrRecordNotFound)
	}
	if err != nil {
		return nil, fmt.Errorf("error getting loan application by uuid: %w", err)
	}

	return loanApplication, nil
}
