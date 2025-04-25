package dao

import (
	"database/sql"
	"sensor-monitoring/internal/domain"
)

type SensorDataDAO struct {
	ID                      sql.NullInt64  `json:"id"`
	OrginatingNumber        sql.NullString `json:"originating_number,omitempty"`
	SensorType              sql.NullString `json:"sensor_type,omitempty"`
	Transcript              sql.NullString `json:"transcript,omitempty"`
	RecordingFile           sql.NullString `json:"recording_file,omitempty"`
	StirShakenIdentityToken sql.NullString `json:"stir_shaken_identity_token,omitempty"`
	Attestation             sql.NullString `json:"attestation,omitempty"`
	CertificateUrl          sql.NullString `json:"certificate_url,omitempty"`
	SpC                     sql.NullString `json:"spc,omitempty"`
	InferredViolation       sql.NullString `json:"inferred_violation,omitempty"`
	ShakenFailed            sql.NullString `json:"shaken_failed,omitempty"`
	UserId                  sql.NullInt64  `json:"user_id,omitempty"`
	NoCert                  sql.NullBool   `json:"no_cert,omitempty"`
	FlaggedBy               sql.NullString `json:"flagged_by,omitempty"`
	Flagged                 sql.NullBool   `json:"flagged,omitempty"`
	CallbackTn              sql.NullString `json:"callback_tn,omitempty"`
	RecordingInbound        sql.NullString `json:"recording_inbound,omitempty"`
	RecordingOutbound       sql.NullString `json:"recording_outbound,omitempty"`
	InferredBrand           sql.NullString `json:"inferred_brand,omitempty"`
	Context                 sql.NullString `json:"context,omitempty"`
	RiskRating              sql.NullInt64  `json:"risk_rating,omitempty"`
	CreatedAt               sql.NullTime   `json:"created_at,omitempty"`
	UpdatedAt               sql.NullTime   `json:"updated_at,omitempty"`
	DeletedAt               sql.NullTime   `json:"deleted_at,omitempty"`
	Timestamp               sql.NullString `json:"timestamp,omitempty"`
}

func (s *SensorDataDAO) FromDAO() *domain.SensorData {
	return &domain.SensorData{
		ID: func() *int {
			id := int(s.ID.Int64)
			return &id
		}(),
		OrginatingNumber:        &s.OrginatingNumber.String,
		SensorType:              &s.SensorType.String,
		Transcript:              &s.Transcript.String,
		RecordingFile:           &s.RecordingFile.String,
		StirShakenIdentityToken: &s.StirShakenIdentityToken.String,
		Attestation:             &s.Attestation.String,
		CertificateUrl:          &s.CertificateUrl.String,
		SpC:                     &s.SpC.String,
		InferredViolation:       &s.InferredViolation.String,
		ShakenFailed:            &s.ShakenFailed.String,
		UserId: func() *int {
			userId := int(s.UserId.Int64)
			return &userId
		}(),
		NoCert:            &s.NoCert.Bool,
		FlaggedBy:         &s.FlaggedBy.String,
		Flagged:           &s.Flagged.Bool,
		CallbackTn:        &s.CallbackTn.String,
		RecordingInbound:  &s.RecordingInbound.String,
		RecordingOutbound: &s.RecordingOutbound.String,
		InferredBrand:     &s.InferredBrand.String,
		Context:           &s.Context.String,
		RiskRating: func() *int {
			riskRating := int(s.RiskRating.Int64)
			return &riskRating
		}(),
		CreatedAt: &s.CreatedAt.Time,
		UpdatedAt: &s.UpdatedAt.Time,
		DeletedAt: &s.DeletedAt.Time,
		Timestamp: &s.Timestamp.String,
	}
}
