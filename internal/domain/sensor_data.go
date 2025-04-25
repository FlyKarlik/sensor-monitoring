package domain

import "time"

type SensorData struct {
	ID                      *int       `json:"id"`
	OrginatingNumber        *string    `json:"originating_number,omitempty"`
	SensorType              *string    `json:"sensor_type,omitempty"`
	Transcript              *string    `json:"transcript,omitempty"`
	RecordingFile           *string    `json:"recording_file,omitempty"`
	StirShakenIdentityToken *string    `json:"stir_shaken_identity_token,omitempty"`
	Attestation             *string    `json:"attestation,omitempty"`
	CertificateUrl          *string    `json:"certificate_url,omitempty"`
	SpC                     *string    `json:"spc,omitempty"`
	InferredViolation       *string    `json:"inferred_violation,omitempty"`
	ShakenFailed            *string    `json:"shaken_failed,omitempty"`
	UserId                  *int       `json:"user_id,omitempty"`
	NoCert                  *bool      `json:"no_cert,omitempty"`
	FlaggedBy               *string    `json:"flagged_by,omitempty"`
	Flagged                 *bool      `json:"flagged,omitempty"`
	CallbackTn              *string    `json:"callback_tn,omitempty"`
	RecordingInbound        *string    `json:"recording_inbound,omitempty"`
	RecordingOutbound       *string    `json:"recording_outbound,omitempty"`
	InferredBrand           *string    `json:"inferred_brand,omitempty"`
	Context                 *string    `json:"context,omitempty"`
	RiskRating              *int       `json:"risk_rating,omitempty"`
	CreatedAt               *time.Time `json:"created_at,omitempty"`
	UpdatedAt               *time.Time `json:"updated_at,omitempty"`
	DeletedAt               *time.Time `json:"deleted_at,omitempty"`
	Timestamp               *string    `json:"timestamp,omitempty"`
}

type SearchSensorDataInput struct {
	SensorDataFilterInput `json:"sensor_data_filter,omitempty"`
}

type SensorDataFilterInput struct {
	InferredBrand   *string `json:"inferred_brand,omitempty"`
	PaginationInput `json:"pagination,omitempty"`
	SortInput       `json:"sort,omitempty"`
}
