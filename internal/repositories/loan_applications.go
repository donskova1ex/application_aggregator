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
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return loanApplications, nil
	case err != nil:
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

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, fmt.Errorf(`loan application not found: %w`, internal.ErrRecordNotFound)
	case err != nil:
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
						) RETURNING uuid`
	newUUID := uuid.NewString()

	row := repo.db.QueryRowContext(ctx, query, newUUID, loanApplication.Phone, loanApplication.Value, loanApplication.IncomingOrganizationUuid)
	err := row.Err()
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Constraint {
			case "chk_positive_value":
				return nil, fmt.Errorf(`error loan application value: %w`, internal.ErrPositiveValue)
			case "loan_applications_uuid_key":
				return nil, fmt.Errorf(`error loan application uuid checking: %w`, internal.ErrEntityUUIDDuplicate)
			case "loan_applications_incoming_organization_uuid_fkey":
				return nil, fmt.Errorf("error loan application incoming organization uuid: %w", internal.ErrIncomingOrganizationUUID)
			}
		}
		return nil, fmt.Errorf("error creating loan application: %w", err)
	}

	var entUUID string
	err = row.Scan(&entUUID)
	if err != nil {
		return nil, fmt.Errorf("error new loan application uuid scan: %w", err)
	}
	loanApplication.Uuid = entUUID

	return loanApplication, nil
}

func (repo *PostgresRepository) UpdateLoanApplication(ctx context.Context, uuid string, loanApplication *domain.LoanApplication) (*domain.LoanApplication, error) {
	if !tools.ValidUUID(uuid) {
		return nil, fmt.Errorf("invalid loan application uuid: %w", internal.ErrUUIDValidation)
	}
	if !tools.ValidUUID(loanApplication.IncomingOrganizationUuid) {
		return nil, fmt.Errorf("invalid incoming organization uuid: %w", internal.ErrIssueOrganizationUUID)
	}
	_, err := repo.GetLoanApplicationsByUUID(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf(`loan application not found: %w`, internal.ErrRecordNotFound)
	}

	query := `UPDATE loan_applications
					SET 
						issue_organization_uuid = $1,
						updated_at = NOW()
					FROM 
						organizations
					WHERE 
						loan_applications.uuid = $2
					RETURNING *`

	var changingLoanApplication *domain.LoanApplication
	err = repo.db.GetContext(ctx, &changingLoanApplication, query, loanApplication.IssueOrganizationUuid, uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf(`loan application not found: %w`, internal.ErrRecordNotFound)
		}

		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Constraint {
			case "loan_applications_issue_organization_uuid_fkey":
				return nil, fmt.Errorf("invalid issue organization uuid: %w", internal.ErrIssueOrganizationUUID)
			}
		}
		return nil, fmt.Errorf("error updating loan application: %w", err)
	}

	return changingLoanApplication, nil
}
