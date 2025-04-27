package dao

import (
	"database/sql"
	"sensor-monitoring/internal/domain"
)

type SensorDataDAO struct {
	ID                      sql.NullInt64
	OrginatingNumber        sql.NullString
	SensorType              sql.NullString
	Transcript              sql.NullString
	RecordingFile           sql.NullString
	StirShakenIdentityToken sql.NullString
	Attestation             sql.NullString
	CertificateUrl          sql.NullString
	SpC                     sql.NullString
	InferredViolation       sql.NullString
	ShakenFailed            sql.NullString
	UserId                  sql.NullInt64
	NoCert                  sql.NullBool
	FlaggedBy               sql.NullString
	Flagged                 sql.NullBool
	CallbackTn              sql.NullString
	RecordingInbound        sql.NullString
	RecordingOutbound       sql.NullString
	InferredBrand           sql.NullString
	Context                 sql.NullString
	RiskRating              sql.NullInt64
	CreatedAt               sql.NullTime
	UpdatedAt               sql.NullTime
	DeletedAt               sql.NullTime
	Timestamp               sql.NullString
}

type SensorDataFilterDAO struct {
	InferredBrand       sql.NullString
	Attestation         sql.NullString
	HasRecording        sql.NullBool
	MinLengthTranscript sql.NullInt64
	PaginationInput     *PaginationDAO
	SortInput           *SortDAO
}

func (s *SensorDataDAO) FromDAO() *domain.SensorData {
	return &domain.SensorData{
		ID:                      fromNullInt64(s.ID),
		OrginatingNumber:        fromNullString(s.OrginatingNumber),
		SensorType:              fromNullString(s.SensorType),
		Transcript:              fromNullString(s.Transcript),
		RecordingFile:           fromNullString(s.RecordingFile),
		StirShakenIdentityToken: fromNullString(s.StirShakenIdentityToken),
		Attestation:             fromNullString(s.Attestation),
		CertificateUrl:          fromNullString(s.CertificateUrl),
		SpC:                     fromNullString(s.SpC),
		InferredViolation:       fromNullString(s.InferredViolation),
		ShakenFailed:            fromNullString(s.ShakenFailed),
		UserId:                  fromNullInt64(s.UserId),
		NoCert:                  fromNullBool(s.NoCert),
		FlaggedBy:               fromNullString(s.FlaggedBy),
		Flagged:                 fromNullBool(s.Flagged),
		CallbackTn:              fromNullString(s.CallbackTn),
		RecordingInbound:        fromNullString(s.RecordingInbound),
		RecordingOutbound:       fromNullString(s.RecordingOutbound),
		InferredBrand:           fromNullString(s.InferredBrand),
		Context:                 fromNullString(s.Context),
		RiskRating:              fromNullInt64(s.RiskRating),
		CreatedAt:               fromNullTime(s.CreatedAt),
		UpdatedAt:               fromNullTime(s.UpdatedAt),
		DeletedAt:               fromNullTime(s.DeletedAt),
		Timestamp:               fromNullString(s.Timestamp),
	}
}

func (s *SensorDataFilterDAO) FromSensorDataFilterInput(domain *domain.SensorDataFilterInput) *SensorDataFilterDAO {
	s.InferredBrand = toNullString(domain.InferredBrand)
	s.Attestation = toNullString(func() *string {
		if domain.Attestation == nil {
			return nil
		}
		return domain.Attestation.StringPtr()
	}())
	s.HasRecording = toNullBool(domain.HasRecording)
	s.MinLengthTranscript = toNullInt64(domain.MinLengthTranscript)
	s.PaginationInput = new(PaginationDAO).FromPaginationInput(&domain.PaginationInput)
	s.SortInput = new(SortDAO).FromSortInput(&domain.SortInput)
	return s
}
