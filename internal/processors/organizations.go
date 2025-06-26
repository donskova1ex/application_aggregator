package processors

import (
	"context"
	"github.com/donskova1ex/application_aggregator/internal/domain"
	"log/slog"
)

type OrganizationsRepository interface {
	CreateOrganization(ctx context.Context, organization *domain.Organization) (*domain.Organization, error)
	GetOrganizationByUUID(ctx context.Context, uuid string) (*domain.Organization, error)
	DeleteOrganizationByUUID(ctx context.Context, uuid string) error
	UpdateOrganization(ctx context.Context, uuid string, organization *domain.Organization) (*domain.Organization, error)
	GetOrganizations(ctx context.Context) ([]*domain.Organization, error)
}

type OrganizationLogger interface {
	Error(msg string, arg ...any)
	Info(msg string, arg ...any)
}

type Organization struct {
	repo   OrganizationsRepository
	logger OrganizationLogger
}

func NewOrganization(repository OrganizationsRepository, logger OrganizationLogger) *Organization {
	return &Organization{
		repo:   repository,
		logger: logger,
	}
}

func (o *Organization) CreateOrganization(ctx context.Context, organization *domain.Organization) (*domain.Organization, error) {
	result, err := o.repo.CreateOrganization(ctx, organization)
	if err != nil {
		o.logger.Error(
			"error creating organization",
			slog.String("error", err.Error()),
		)
		return nil, err
	}

	o.logger.Info(
		"creating organization successfully",
		slog.String("organization", organization.Name),
	)
	return result, nil
}

func (o *Organization) GetOrganizationByUUID(ctx context.Context, uuid string) (*domain.Organization, error) {
	result, err := o.repo.GetOrganizationByUUID(ctx, uuid)
	if err != nil {
		o.logger.Error(
			"error getting organization by uuid",
			slog.String("error", err.Error()),
			slog.String("uuid", uuid),
		)
		return nil, err
	}

	o.logger.Info(
		"getting organization by uuid successfully",
		slog.String("uuid", uuid),
		slog.String("organization", result.Name),
	)
	return result, nil
}

func (o *Organization) DeleteOrganizationByUUID(ctx context.Context, uuid string) error {
	err := o.repo.DeleteOrganizationByUUID(ctx, uuid)
	if err != nil {
		o.logger.Error(
			"error deleting organization by uuid",
			slog.String("error", err.Error()),
			slog.String("uuid", uuid),
		)
		return err
	}

	o.logger.Info(
		"deleting organization by uuid successfully",
		slog.String("uuid", uuid),
	)
	return nil
}

func (o *Organization) UpdateOrganization(ctx context.Context, uuid string, organization *domain.Organization) (*domain.Organization, error) {
	result, err := o.repo.UpdateOrganization(ctx, uuid, organization)
	if err != nil {
		o.logger.Error(
			"error updating organization",
			slog.String("error", err.Error()),
			slog.String("uuid", uuid),
		)
		return nil, err
	}
	o.logger.Info(
		"updating organization successfully",
		slog.String("uuid", uuid),
		slog.String("organization", result.Name),
	)
	return result, nil
}

func (o *Organization) GetOrganizations(ctx context.Context) ([]*domain.Organization, error) {
	result, err := o.repo.GetOrganizations(ctx)
	if err != nil {
		o.logger.Error(
			"error getting organizations list",
			slog.String("error", err.Error()))
	}

	o.logger.Info(
		"getting organizations successfully",
	)
	return result, err
}
