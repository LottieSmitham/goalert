// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: queries.sql

package gadb

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/sqlc-dev/pqtype"
)

const alertFeedback = `-- name: AlertFeedback :many
SELECT
    alert_id,
    noise_reason
FROM
    alert_feedback
WHERE
    alert_id = ANY($1::int[])
`

type AlertFeedbackRow struct {
	AlertID     int64
	NoiseReason string
}

func (q *Queries) AlertFeedback(ctx context.Context, dollar_1 []int32) ([]AlertFeedbackRow, error) {
	rows, err := q.db.QueryContext(ctx, alertFeedback, pq.Array(dollar_1))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AlertFeedbackRow
	for rows.Next() {
		var i AlertFeedbackRow
		if err := rows.Scan(&i.AlertID, &i.NoiseReason); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const alertHasEPState = `-- name: AlertHasEPState :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            escalation_policy_state
        WHERE
            alert_id = $1) AS has_ep_state
`

func (q *Queries) AlertHasEPState(ctx context.Context, alertID int64) (bool, error) {
	row := q.db.QueryRowContext(ctx, alertHasEPState, alertID)
	var has_ep_state bool
	err := row.Scan(&has_ep_state)
	return has_ep_state, err
}

const alertLogHBIntervalMinutes = `-- name: AlertLogHBIntervalMinutes :one
SELECT
    (EXTRACT(EPOCH FROM heartbeat_interval) / 60)::int
FROM
    heartbeat_monitors
WHERE
    id = $1
`

func (q *Queries) AlertLogHBIntervalMinutes(ctx context.Context, id uuid.UUID) (int32, error) {
	row := q.db.QueryRowContext(ctx, alertLogHBIntervalMinutes, id)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}

const alertLogInsertEP = `-- name: AlertLogInsertEP :exec
INSERT INTO alert_logs(alert_id, event, sub_type, sub_user_id, sub_integration_key_id, sub_hb_monitor_id, sub_channel_id, sub_classifier, meta, message)
SELECT
    a.id,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
FROM
    alerts a
    JOIN services svc ON svc.id = a.service_id
        AND svc.escalation_policy_id = $1
WHERE
    a.status != 'closed'
`

type AlertLogInsertEPParams struct {
	EscalationPolicyID  uuid.UUID
	Event               EnumAlertLogEvent
	SubType             NullEnumAlertLogSubjectType
	SubUserID           uuid.NullUUID
	SubIntegrationKeyID uuid.NullUUID
	SubHbMonitorID      uuid.NullUUID
	SubChannelID        uuid.NullUUID
	SubClassifier       string
	Meta                pqtype.NullRawMessage
	Message             string
}

func (q *Queries) AlertLogInsertEP(ctx context.Context, arg AlertLogInsertEPParams) error {
	_, err := q.db.ExecContext(ctx, alertLogInsertEP,
		arg.EscalationPolicyID,
		arg.Event,
		arg.SubType,
		arg.SubUserID,
		arg.SubIntegrationKeyID,
		arg.SubHbMonitorID,
		arg.SubChannelID,
		arg.SubClassifier,
		arg.Meta,
		arg.Message,
	)
	return err
}

const alertLogInsertMany = `-- name: AlertLogInsertMany :exec
INSERT INTO alert_logs(alert_id, event, sub_type, sub_user_id, sub_integration_key_id, sub_hb_monitor_id, sub_channel_id, sub_classifier, meta, message)
SELECT
    unnest,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
FROM
    unnest($1::bigint[])
`

type AlertLogInsertManyParams struct {
	Column1             []int64
	Event               EnumAlertLogEvent
	SubType             NullEnumAlertLogSubjectType
	SubUserID           uuid.NullUUID
	SubIntegrationKeyID uuid.NullUUID
	SubHbMonitorID      uuid.NullUUID
	SubChannelID        uuid.NullUUID
	SubClassifier       string
	Meta                pqtype.NullRawMessage
	Message             string
}

func (q *Queries) AlertLogInsertMany(ctx context.Context, arg AlertLogInsertManyParams) error {
	_, err := q.db.ExecContext(ctx, alertLogInsertMany,
		pq.Array(arg.Column1),
		arg.Event,
		arg.SubType,
		arg.SubUserID,
		arg.SubIntegrationKeyID,
		arg.SubHbMonitorID,
		arg.SubChannelID,
		arg.SubClassifier,
		arg.Meta,
		arg.Message,
	)
	return err
}

const alertLogInsertSvc = `-- name: AlertLogInsertSvc :exec
INSERT INTO alert_logs(alert_id, event, sub_type, sub_user_id, sub_integration_key_id, sub_hb_monitor_id, sub_channel_id, sub_classifier, meta, message)
SELECT
    a.id,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
FROM
    alerts a
WHERE
    a.service_id = $1
    AND (($2 = 'closed'::enum_alert_log_event
            AND a.status != 'closed')
        OR ($2::enum_alert_log_event IN ('acknowledged', 'notification_sent')
            AND a.status = 'triggered'))
`

type AlertLogInsertSvcParams struct {
	ServiceID           uuid.NullUUID
	Event               EnumAlertLogEvent
	SubType             NullEnumAlertLogSubjectType
	SubUserID           uuid.NullUUID
	SubIntegrationKeyID uuid.NullUUID
	SubHbMonitorID      uuid.NullUUID
	SubChannelID        uuid.NullUUID
	SubClassifier       string
	Meta                pqtype.NullRawMessage
	Message             string
}

func (q *Queries) AlertLogInsertSvc(ctx context.Context, arg AlertLogInsertSvcParams) error {
	_, err := q.db.ExecContext(ctx, alertLogInsertSvc,
		arg.ServiceID,
		arg.Event,
		arg.SubType,
		arg.SubUserID,
		arg.SubIntegrationKeyID,
		arg.SubHbMonitorID,
		arg.SubChannelID,
		arg.SubClassifier,
		arg.Meta,
		arg.Message,
	)
	return err
}

const alertLogLookupCMType = `-- name: AlertLogLookupCMType :one
SELECT
    "type" AS cm_type
FROM
    user_contact_methods
WHERE
    id = $1
`

func (q *Queries) AlertLogLookupCMType(ctx context.Context, id uuid.UUID) (EnumUserContactMethodType, error) {
	row := q.db.QueryRowContext(ctx, alertLogLookupCMType, id)
	var cm_type EnumUserContactMethodType
	err := row.Scan(&cm_type)
	return cm_type, err
}

const allPendingMsgDests = `-- name: AllPendingMsgDests :many
SELECT DISTINCT
    usr.name AS user_name,
    cm.type AS cm_type,
    nc.name AS nc_name,
    nc.type AS nc_type
FROM
    outgoing_messages om
    LEFT JOIN users usr ON usr.id = om.user_id
    LEFT JOIN notification_channels nc ON nc.id = om.channel_id
    LEFT JOIN user_contact_methods cm ON cm.id = om.contact_method_id
WHERE
    om.last_status = 'pending'
    AND (now() - om.created_at) > INTERVAL '15 seconds'
    AND (om.alert_id = $1::bigint
        OR (om.message_type = 'alert_notification_bundle'
            AND om.service_id = $2::uuid))
`

type AllPendingMsgDestsParams struct {
	AlertID   int64
	ServiceID uuid.UUID
}

type AllPendingMsgDestsRow struct {
	UserName sql.NullString
	CmType   NullEnumUserContactMethodType
	NcName   sql.NullString
	NcType   NullEnumNotifChannelType
}

func (q *Queries) AllPendingMsgDests(ctx context.Context, arg AllPendingMsgDestsParams) ([]AllPendingMsgDestsRow, error) {
	rows, err := q.db.QueryContext(ctx, allPendingMsgDests, arg.AlertID, arg.ServiceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AllPendingMsgDestsRow
	for rows.Next() {
		var i AllPendingMsgDestsRow
		if err := rows.Scan(
			&i.UserName,
			&i.CmType,
			&i.NcName,
			&i.NcType,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const authLinkAddAuthSubject = `-- name: AuthLinkAddAuthSubject :exec
INSERT INTO auth_subjects(provider_id, subject_id, user_id)
    VALUES ($1, $2, $3)
`

type AuthLinkAddAuthSubjectParams struct {
	ProviderID string
	SubjectID  string
	UserID     uuid.UUID
}

func (q *Queries) AuthLinkAddAuthSubject(ctx context.Context, arg AuthLinkAddAuthSubjectParams) error {
	_, err := q.db.ExecContext(ctx, authLinkAddAuthSubject, arg.ProviderID, arg.SubjectID, arg.UserID)
	return err
}

const authLinkAddReq = `-- name: AuthLinkAddReq :exec
INSERT INTO auth_link_requests(id, provider_id, subject_id, expires_at, metadata)
    VALUES ($1, $2, $3, $4, $5)
`

type AuthLinkAddReqParams struct {
	ID         uuid.UUID
	ProviderID string
	SubjectID  string
	ExpiresAt  time.Time
	Metadata   json.RawMessage
}

func (q *Queries) AuthLinkAddReq(ctx context.Context, arg AuthLinkAddReqParams) error {
	_, err := q.db.ExecContext(ctx, authLinkAddReq,
		arg.ID,
		arg.ProviderID,
		arg.SubjectID,
		arg.ExpiresAt,
		arg.Metadata,
	)
	return err
}

const authLinkMetadata = `-- name: AuthLinkMetadata :one
SELECT
    metadata
FROM
    auth_link_requests
WHERE
    id = $1
    AND expires_at > now()
`

func (q *Queries) AuthLinkMetadata(ctx context.Context, id uuid.UUID) (json.RawMessage, error) {
	row := q.db.QueryRowContext(ctx, authLinkMetadata, id)
	var metadata json.RawMessage
	err := row.Scan(&metadata)
	return metadata, err
}

const authLinkUseReq = `-- name: AuthLinkUseReq :one
DELETE FROM auth_link_requests
WHERE id = $1
    AND expires_at > now()
RETURNING
    provider_id,
    subject_id
`

type AuthLinkUseReqRow struct {
	ProviderID string
	SubjectID  string
}

func (q *Queries) AuthLinkUseReq(ctx context.Context, id uuid.UUID) (AuthLinkUseReqRow, error) {
	row := q.db.QueryRowContext(ctx, authLinkUseReq, id)
	var i AuthLinkUseReqRow
	err := row.Scan(&i.ProviderID, &i.SubjectID)
	return i, err
}

const calSubAuthUser = `-- name: CalSubAuthUser :one
UPDATE user_calendar_subscriptions
SET last_access = now()
WHERE NOT disabled
    AND id = $1
    AND date_trunc('second', created_at) = $2 RETURNING user_id
`

type CalSubAuthUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
}

func (q *Queries) CalSubAuthUser(ctx context.Context, arg CalSubAuthUserParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, calSubAuthUser, arg.ID, arg.CreatedAt)
	var user_id uuid.UUID
	err := row.Scan(&user_id)
	return user_id, err
}

const calSubRenderInfo = `-- name: CalSubRenderInfo :one
SELECT
    now()::timestamptz AS now,
    sub.schedule_id,
    sched.name AS schedule_name,
    sub.config,
    sub.user_id
FROM
    user_calendar_subscriptions sub
    JOIN schedules sched ON sched.id = schedule_id
WHERE
    sub.id = $1
`

type CalSubRenderInfoRow struct {
	Now          time.Time
	ScheduleID   uuid.UUID
	ScheduleName string
	Config       json.RawMessage
	UserID       uuid.UUID
}

func (q *Queries) CalSubRenderInfo(ctx context.Context, id uuid.UUID) (CalSubRenderInfoRow, error) {
	row := q.db.QueryRowContext(ctx, calSubRenderInfo, id)
	var i CalSubRenderInfoRow
	err := row.Scan(
		&i.Now,
		&i.ScheduleID,
		&i.ScheduleName,
		&i.Config,
		&i.UserID,
	)
	return i, err
}

const createCalSub = `-- name: CreateCalSub :one
INSERT INTO user_calendar_subscriptions (
        id,
        NAME,
        user_id,
        disabled,
        schedule_id,
        config
    )
VALUES ($1, $2, $3, $4, $5, $6) RETURNING created_at
`

type CreateCalSubParams struct {
	ID         uuid.UUID
	Name       string
	UserID     uuid.UUID
	Disabled   bool
	ScheduleID uuid.UUID
	Config     json.RawMessage
}

func (q *Queries) CreateCalSub(ctx context.Context, arg CreateCalSubParams) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, createCalSub,
		arg.ID,
		arg.Name,
		arg.UserID,
		arg.Disabled,
		arg.ScheduleID,
		arg.Config,
	)
	var created_at time.Time
	err := row.Scan(&created_at)
	return created_at, err
}

const deleteManyCalSub = `-- name: DeleteManyCalSub :exec
DELETE FROM user_calendar_subscriptions
WHERE id = ANY($1::uuid [ ])
    AND user_id = $2
`

type DeleteManyCalSubParams struct {
	Column1 []uuid.UUID
	UserID  uuid.UUID
}

func (q *Queries) DeleteManyCalSub(ctx context.Context, arg DeleteManyCalSubParams) error {
	_, err := q.db.ExecContext(ctx, deleteManyCalSub, pq.Array(arg.Column1), arg.UserID)
	return err
}

const findManyCalSubByUser = `-- name: FindManyCalSubByUser :many
SELECT id,
    NAME,
    user_id,
    disabled,
    schedule_id,
    config,
    last_access
FROM user_calendar_subscriptions
WHERE user_id = $1
`

type FindManyCalSubByUserRow struct {
	ID         uuid.UUID
	Name       string
	UserID     uuid.UUID
	Disabled   bool
	ScheduleID uuid.UUID
	Config     json.RawMessage
	LastAccess sql.NullTime
}

func (q *Queries) FindManyCalSubByUser(ctx context.Context, userID uuid.UUID) ([]FindManyCalSubByUserRow, error) {
	rows, err := q.db.QueryContext(ctx, findManyCalSubByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindManyCalSubByUserRow
	for rows.Next() {
		var i FindManyCalSubByUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.UserID,
			&i.Disabled,
			&i.ScheduleID,
			&i.Config,
			&i.LastAccess,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findOneCalSub = `-- name: FindOneCalSub :one
SELECT id,
    NAME,
    user_id,
    disabled,
    schedule_id,
    config,
    last_access
FROM user_calendar_subscriptions
WHERE id = $1
`

type FindOneCalSubRow struct {
	ID         uuid.UUID
	Name       string
	UserID     uuid.UUID
	Disabled   bool
	ScheduleID uuid.UUID
	Config     json.RawMessage
	LastAccess sql.NullTime
}

func (q *Queries) FindOneCalSub(ctx context.Context, id uuid.UUID) (FindOneCalSubRow, error) {
	row := q.db.QueryRowContext(ctx, findOneCalSub, id)
	var i FindOneCalSubRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.Disabled,
		&i.ScheduleID,
		&i.Config,
		&i.LastAccess,
	)
	return i, err
}

const findOneCalSubForUpdate = `-- name: FindOneCalSubForUpdate :one
SELECT id,
    NAME,
    user_id,
    disabled,
    schedule_id,
    config,
    last_access
FROM user_calendar_subscriptions
WHERE id = $1 FOR
UPDATE
`

type FindOneCalSubForUpdateRow struct {
	ID         uuid.UUID
	Name       string
	UserID     uuid.UUID
	Disabled   bool
	ScheduleID uuid.UUID
	Config     json.RawMessage
	LastAccess sql.NullTime
}

func (q *Queries) FindOneCalSubForUpdate(ctx context.Context, id uuid.UUID) (FindOneCalSubForUpdateRow, error) {
	row := q.db.QueryRowContext(ctx, findOneCalSubForUpdate, id)
	var i FindOneCalSubForUpdateRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.Disabled,
		&i.ScheduleID,
		&i.Config,
		&i.LastAccess,
	)
	return i, err
}

const getRulesForIntegrationKey = `-- name: GetRulesForIntegrationKey :many
SELECT
    r.id,
    r.name,
    r.service_id,
    r.filter,
    r.send_alert,
    r.actions
FROM
    service_rule_integration_keys AS sk
JOIN service_rules AS r 
	ON sk.service_rule_id = r.id
    AND r.service_id = $1
    AND sk.integration_key_id = $2
`

type GetRulesForIntegrationKeyParams struct {
	ServiceID        uuid.UUID
	IntegrationKeyID uuid.UUID
}

type GetRulesForIntegrationKeyRow struct {
	ID        uuid.UUID
	Name      string
	ServiceID uuid.UUID
	Filter    string
	SendAlert bool
	Actions   pqtype.NullRawMessage
}

func (q *Queries) GetRulesForIntegrationKey(ctx context.Context, arg GetRulesForIntegrationKeyParams) ([]GetRulesForIntegrationKeyRow, error) {
	rows, err := q.db.QueryContext(ctx, getRulesForIntegrationKey, arg.ServiceID, arg.IntegrationKeyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRulesForIntegrationKeyRow
	for rows.Next() {
		var i GetRulesForIntegrationKeyRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ServiceID,
			&i.Filter,
			&i.SendAlert,
			&i.Actions,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRulesForService = `-- name: GetRulesForService :many
SELECT
    service_rules.id,
    service_rules.name,
    service_rules.service_id,
    service_rules.filter,
    service_rules.send_alert,
    service_rules.actions,
    STRING_AGG (
      service_rule_integration_keys.integration_key_id::text,
      ','
    )::TEXT integration_keys
FROM
    service_rules
JOIN service_rule_integration_keys ON service_rule_integration_keys.service_rule_id = service_rules.id
WHERE 
    service_rules.service_id = $1
GROUP BY
		service_rules.id
`

type GetRulesForServiceRow struct {
	ID              uuid.UUID
	Name            string
	ServiceID       uuid.UUID
	Filter          string
	SendAlert       bool
	Actions         pqtype.NullRawMessage
	IntegrationKeys string
}

func (q *Queries) GetRulesForService(ctx context.Context, serviceID uuid.UUID) ([]GetRulesForServiceRow, error) {
	rows, err := q.db.QueryContext(ctx, getRulesForService, serviceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRulesForServiceRow
	for rows.Next() {
		var i GetRulesForServiceRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ServiceID,
			&i.Filter,
			&i.SendAlert,
			&i.Actions,
			&i.IntegrationKeys,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertServiceRule = `-- name: InsertServiceRule :exec
INSERT INTO service_rules(name, service_id, filter, send_alert, actions)
    VALUES ($1, $2, $3, $4, $5)
`

type InsertServiceRuleParams struct {
	Name      string
	ServiceID uuid.UUID
	Filter    string
	SendAlert bool
	Actions   pqtype.NullRawMessage
}

func (q *Queries) InsertServiceRule(ctx context.Context, arg InsertServiceRuleParams) error {
	_, err := q.db.ExecContext(ctx, insertServiceRule,
		arg.Name,
		arg.ServiceID,
		arg.Filter,
		arg.SendAlert,
		arg.Actions,
	)
	return err
}

const lockOneAlertService = `-- name: LockOneAlertService :one
SELECT
    maintenance_expires_at NOTNULL::bool AS is_maint_mode,
    alerts.status
FROM
    services svc
    JOIN alerts ON alerts.service_id = svc.id
WHERE
    alerts.id = $1
FOR UPDATE
`

type LockOneAlertServiceRow struct {
	IsMaintMode bool
	Status      EnumAlertStatus
}

func (q *Queries) LockOneAlertService(ctx context.Context, id int64) (LockOneAlertServiceRow, error) {
	row := q.db.QueryRowContext(ctx, lockOneAlertService, id)
	var i LockOneAlertServiceRow
	err := row.Scan(&i.IsMaintMode, &i.Status)
	return i, err
}

const noticeUnackedAlertsByService = `-- name: NoticeUnackedAlertsByService :one
SELECT
    count(*),
    (
        SELECT
            max
        FROM
            config_limits
        WHERE
            id = 'unacked_alerts_per_service'
    )
FROM
    alerts
WHERE
    service_id = $1::uuid
    AND status = 'triggered'
`

type NoticeUnackedAlertsByServiceRow struct {
	Count int64
	Max   int32
}

func (q *Queries) NoticeUnackedAlertsByService(ctx context.Context, dollar_1 uuid.UUID) (NoticeUnackedAlertsByServiceRow, error) {
	row := q.db.QueryRowContext(ctx, noticeUnackedAlertsByService, dollar_1)
	var i NoticeUnackedAlertsByServiceRow
	err := row.Scan(&i.Count, &i.Max)
	return i, err
}

const now = `-- name: Now :one
SELECT now()::timestamptz
`

func (q *Queries) Now(ctx context.Context) (time.Time, error) {
	row := q.db.QueryRowContext(ctx, now)
	var column_1 time.Time
	err := row.Scan(&column_1)
	return column_1, err
}

const requestAlertEscalationByTime = `-- name: RequestAlertEscalationByTime :one
UPDATE
    escalation_policy_state
SET
    force_escalation = TRUE
WHERE
    alert_id = $1
    AND (last_escalation <= $2::timestamptz
        OR last_escalation IS NULL)
RETURNING
    TRUE
`

type RequestAlertEscalationByTimeParams struct {
	AlertID int64
	Column2 time.Time
}

func (q *Queries) RequestAlertEscalationByTime(ctx context.Context, arg RequestAlertEscalationByTimeParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, requestAlertEscalationByTime, arg.AlertID, arg.Column2)
	var column_1 bool
	err := row.Scan(&column_1)
	return column_1, err
}

const setAlertFeedback = `-- name: SetAlertFeedback :exec
INSERT INTO alert_feedback(alert_id, noise_reason)
    VALUES ($1, $2)
ON CONFLICT (alert_id)
    DO UPDATE SET
        noise_reason = $2
    WHERE
        alert_feedback.alert_id = $1
`

type SetAlertFeedbackParams struct {
	AlertID     int64
	NoiseReason string
}

func (q *Queries) SetAlertFeedback(ctx context.Context, arg SetAlertFeedbackParams) error {
	_, err := q.db.ExecContext(ctx, setAlertFeedback, arg.AlertID, arg.NoiseReason)
	return err
}

const statusMgrCMInfo = `-- name: StatusMgrCMInfo :one
SELECT
    user_id,
    type
FROM
    user_contact_methods
WHERE
    id = $1
    AND NOT disabled
    AND enable_status_updates
`

type StatusMgrCMInfoRow struct {
	UserID uuid.UUID
	Type   EnumUserContactMethodType
}

func (q *Queries) StatusMgrCMInfo(ctx context.Context, id uuid.UUID) (StatusMgrCMInfoRow, error) {
	row := q.db.QueryRowContext(ctx, statusMgrCMInfo, id)
	var i StatusMgrCMInfoRow
	err := row.Scan(&i.UserID, &i.Type)
	return i, err
}

const statusMgrCleanupDisabledSubs = `-- name: StatusMgrCleanupDisabledSubs :exec
DELETE FROM alert_status_subscriptions sub USING user_contact_methods cm
WHERE sub.contact_method_id = cm.id
    AND (cm.disabled
        OR NOT cm.enable_status_updates)
`

func (q *Queries) StatusMgrCleanupDisabledSubs(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, statusMgrCleanupDisabledSubs)
	return err
}

const statusMgrDeleteSub = `-- name: StatusMgrDeleteSub :exec
DELETE FROM alert_status_subscriptions
WHERE id = $1
`

func (q *Queries) StatusMgrDeleteSub(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, statusMgrDeleteSub, id)
	return err
}

const statusMgrLogEntry = `-- name: StatusMgrLogEntry :one
SELECT
    id,
    sub_user_id AS user_id
FROM
    alert_logs
WHERE
    alert_id = $1::bigint
    AND event = $2::enum_alert_log_event
    AND timestamp > now() - '1 hour'::interval
ORDER BY
    id DESC
LIMIT 1
`

type StatusMgrLogEntryParams struct {
	AlertID   int64
	EventType EnumAlertLogEvent
}

type StatusMgrLogEntryRow struct {
	ID     int64
	UserID uuid.NullUUID
}

func (q *Queries) StatusMgrLogEntry(ctx context.Context, arg StatusMgrLogEntryParams) (StatusMgrLogEntryRow, error) {
	row := q.db.QueryRowContext(ctx, statusMgrLogEntry, arg.AlertID, arg.EventType)
	var i StatusMgrLogEntryRow
	err := row.Scan(&i.ID, &i.UserID)
	return i, err
}

const statusMgrNextUpdate = `-- name: StatusMgrNextUpdate :one
SELECT
    sub.id,
    channel_id,
    contact_method_id,
    alert_id,
    (
        SELECT
            status
        FROM
            alerts a
        WHERE
            a.id = sub.alert_id)
FROM
    alert_status_subscriptions sub
WHERE
    sub.last_alert_status != (
        SELECT
            status
        FROM
            alerts a
        WHERE
            a.id = sub.alert_id)
LIMIT 1
FOR UPDATE
    SKIP LOCKED
`

type StatusMgrNextUpdateRow struct {
	ID              int64
	ChannelID       uuid.NullUUID
	ContactMethodID uuid.NullUUID
	AlertID         int64
	Status          EnumAlertStatus
}

func (q *Queries) StatusMgrNextUpdate(ctx context.Context) (StatusMgrNextUpdateRow, error) {
	row := q.db.QueryRowContext(ctx, statusMgrNextUpdate)
	var i StatusMgrNextUpdateRow
	err := row.Scan(
		&i.ID,
		&i.ChannelID,
		&i.ContactMethodID,
		&i.AlertID,
		&i.Status,
	)
	return i, err
}

const statusMgrSendChannelMsg = `-- name: StatusMgrSendChannelMsg :exec
INSERT INTO outgoing_messages (id, message_type, channel_id, alert_id, alert_log_id)
    VALUES ($1::uuid, 'alert_status_update', $2::uuid, $3::bigint, $4)
`

type StatusMgrSendChannelMsgParams struct {
	ID        uuid.UUID
	ChannelID uuid.UUID
	AlertID   int64
	LogID     sql.NullInt64
}

func (q *Queries) StatusMgrSendChannelMsg(ctx context.Context, arg StatusMgrSendChannelMsgParams) error {
	_, err := q.db.ExecContext(ctx, statusMgrSendChannelMsg,
		arg.ID,
		arg.ChannelID,
		arg.AlertID,
		arg.LogID,
	)
	return err
}

const statusMgrSendUserMsg = `-- name: StatusMgrSendUserMsg :exec
INSERT INTO outgoing_messages (id, message_type, contact_method_id, user_id, alert_id, alert_log_id)
    VALUES ($1::uuid, 'alert_status_update', $2::uuid, $3::uuid, $4::bigint, $5)
`

type StatusMgrSendUserMsgParams struct {
	ID      uuid.UUID
	CmID    uuid.UUID
	UserID  uuid.UUID
	AlertID int64
	LogID   sql.NullInt64
}

func (q *Queries) StatusMgrSendUserMsg(ctx context.Context, arg StatusMgrSendUserMsgParams) error {
	_, err := q.db.ExecContext(ctx, statusMgrSendUserMsg,
		arg.ID,
		arg.CmID,
		arg.UserID,
		arg.AlertID,
		arg.LogID,
	)
	return err
}

const statusMgrUpdateCMForced = `-- name: StatusMgrUpdateCMForced :exec
UPDATE
    user_contact_methods
SET
    enable_status_updates = TRUE
WHERE
    TYPE = 'SLACK_DM'
    AND NOT enable_status_updates
`

func (q *Queries) StatusMgrUpdateCMForced(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, statusMgrUpdateCMForced)
	return err
}

const statusMgrUpdateSub = `-- name: StatusMgrUpdateSub :exec
UPDATE
    alert_status_subscriptions
SET
    last_alert_status = $2
WHERE
    id = $1
`

type StatusMgrUpdateSubParams struct {
	ID              int64
	LastAlertStatus EnumAlertStatus
}

func (q *Queries) StatusMgrUpdateSub(ctx context.Context, arg StatusMgrUpdateSubParams) error {
	_, err := q.db.ExecContext(ctx, statusMgrUpdateSub, arg.ID, arg.LastAlertStatus)
	return err
}

const updateCalSub = `-- name: UpdateCalSub :exec
UPDATE user_calendar_subscriptions
SET NAME = $1,
    disabled = $2,
    config = $3,
    last_update = now()
WHERE id = $4
    AND user_id = $5
`

type UpdateCalSubParams struct {
	Name     string
	Disabled bool
	Config   json.RawMessage
	ID       uuid.UUID
	UserID   uuid.UUID
}

func (q *Queries) UpdateCalSub(ctx context.Context, arg UpdateCalSubParams) error {
	_, err := q.db.ExecContext(ctx, updateCalSub,
		arg.Name,
		arg.Disabled,
		arg.Config,
		arg.ID,
		arg.UserID,
	)
	return err
}
