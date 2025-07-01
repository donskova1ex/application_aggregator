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
	"github.com/lib/pq"
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
		return nil, fmt.Errorf("invalid loan application uuid: %w", internal.ErrUUIDValidation)
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

func (repo *PostgresRepository) CreateLoanApplication(ctx context.Context, loanApplication *domain.LoanApplication) (*domain.LoanApplication, error) {
	if !tools.ValidUUID(loanApplication.IncomingOrganizationUuid) {
		return nil, fmt.Errorf("invalid incoming organization uuid: %w", internal.ErrUUIDValidation)
	}
	if !tools.ValidPhone(loanApplication.Phone) {
		return nil, fmt.Errorf("invalid phone number: %w", internal.ErrPhoneValidation)
	}

	query := `INSERT INTO loan_applications (uuid, phone, value, incoming_organization_uuid)
					SELECT 
						$1, 
						$2, 
						$3, 
						$4
					WHERE NOT EXISTS (
						SELECT 1 
						FROM loan_applications 
						WHERE phone = $2
						AND incoming_organization_uuid = $4
						AND created_at::date = CURRENT_DATE
						)`
	newUUID := uuid.NewString()

	//result, err := repo.db.ExecContext(ctx, query, newUUID, loanApplication.Phone, loanApplication.Value, loanApplication.IncomingOrganizationUuid)
	row := repo.db.QueryRowContext(ctx, query, newUUID, loanApplication.Phone, loanApplication.Value, loanApplication.IncomingOrganizationUuid)
	var pqErr *pq.Error
	err := row.Err()
	if err != nil {
		if errors.As(err, &pqErr) {
			if pqErr.Constraint == "chk_positive_value" {
				return nil, fmt.Errorf(`error loan application value: %s`, loanApplication.Phone)
			}
			if pqErr.Constraint == "loan_applications_uuid_key" {
				return nil, fmt.Errorf(`error loan application uuid checking: %s`, loanApplication.IncomingOrganizationUuid)
			}
		}
	}
	//TODO: continue
	loanApplication.Uuid = newUUID

	return loanApplication, nil
}
