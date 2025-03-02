// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql2

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/target/goalert/alert"
	"github.com/target/goalert/alert/alertlog"
	"github.com/target/goalert/assignment"
	"github.com/target/goalert/escalation"
	"github.com/target/goalert/integrationkey"
	"github.com/target/goalert/label"
	"github.com/target/goalert/limit"
	"github.com/target/goalert/notification"
	"github.com/target/goalert/notification/slack"
	"github.com/target/goalert/override"
	"github.com/target/goalert/schedule"
	"github.com/target/goalert/schedule/rotation"
	"github.com/target/goalert/schedule/rule"
	"github.com/target/goalert/service"
	"github.com/target/goalert/user"
	"github.com/target/goalert/user/contactmethod"
	"github.com/target/goalert/util/timeutil"
)

type AlertConnection struct {
	Nodes    []alert.Alert `json:"nodes"`
	PageInfo *PageInfo     `json:"pageInfo"`
}

type AlertDataPoint struct {
	Timestamp  time.Time `json:"timestamp"`
	AlertCount int       `json:"alertCount"`
}

type AlertLogEntryConnection struct {
	Nodes    []alertlog.Entry `json:"nodes"`
	PageInfo *PageInfo        `json:"pageInfo"`
}

type AlertMetricsOptions struct {
	RInterval         timeutil.ISORInterval `json:"rInterval"`
	FilterByServiceID []string              `json:"filterByServiceID,omitempty"`
}

type AlertPendingNotification struct {
	Destination string `json:"destination"`
}

type AlertRecentEventsOptions struct {
	Limit *int    `json:"limit,omitempty"`
	After *string `json:"after,omitempty"`
}

type AlertSearchOptions struct {
	FilterByStatus    []AlertStatus    `json:"filterByStatus,omitempty"`
	FilterByServiceID []string         `json:"filterByServiceID,omitempty"`
	Search            *string          `json:"search,omitempty"`
	First             *int             `json:"first,omitempty"`
	After             *string          `json:"after,omitempty"`
	FavoritesOnly     *bool            `json:"favoritesOnly,omitempty"`
	IncludeNotified   *bool            `json:"includeNotified,omitempty"`
	Omit              []int            `json:"omit,omitempty"`
	Sort              *AlertSearchSort `json:"sort,omitempty"`
	CreatedBefore     *time.Time       `json:"createdBefore,omitempty"`
	NotCreatedBefore  *time.Time       `json:"notCreatedBefore,omitempty"`
	ClosedBefore      *time.Time       `json:"closedBefore,omitempty"`
	NotClosedBefore   *time.Time       `json:"notClosedBefore,omitempty"`
}

type AuthSubjectConnection struct {
	Nodes    []user.AuthSubject `json:"nodes"`
	PageInfo *PageInfo          `json:"pageInfo"`
}

type CalcRotationHandoffTimesInput struct {
	Handoff          time.Time             `json:"handoff"`
	From             *time.Time            `json:"from,omitempty"`
	TimeZone         string                `json:"timeZone"`
	ShiftLengthHours *int                  `json:"shiftLengthHours,omitempty"`
	ShiftLength      *timeutil.ISODuration `json:"shiftLength,omitempty"`
	Count            int                   `json:"count"`
}

type ClearTemporarySchedulesInput struct {
	ScheduleID string    `json:"scheduleID"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
}

type ConfigHint struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type ConfigValue struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Value       string     `json:"value"`
	Type        ConfigType `json:"type"`
	Password    bool       `json:"password"`
	Deprecated  string     `json:"deprecated"`
}

type ConfigValueInput struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type CreateAlertInput struct {
	Summary   string  `json:"summary"`
	Details   *string `json:"details,omitempty"`
	ServiceID string  `json:"serviceID"`
	Sanitize  *bool   `json:"sanitize,omitempty"`
}

type CreateBasicAuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserID   string `json:"userID"`
}

type CreateEscalationPolicyInput struct {
	Name        string                            `json:"name"`
	Description *string                           `json:"description,omitempty"`
	Repeat      *int                              `json:"repeat,omitempty"`
	Favorite    *bool                             `json:"favorite,omitempty"`
	Steps       []CreateEscalationPolicyStepInput `json:"steps,omitempty"`
}

type CreateEscalationPolicyStepInput struct {
	EscalationPolicyID *string                `json:"escalationPolicyID,omitempty"`
	DelayMinutes       int                    `json:"delayMinutes"`
	Targets            []assignment.RawTarget `json:"targets,omitempty"`
	NewRotation        *CreateRotationInput   `json:"newRotation,omitempty"`
	NewSchedule        *CreateScheduleInput   `json:"newSchedule,omitempty"`
}

type CreateGQLAPIKeyInput struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	AllowedFields []string  `json:"allowedFields"`
	ExpiresAt     time.Time `json:"expiresAt"`
	Role          UserRole  `json:"role"`
}

type CreateHeartbeatMonitorInput struct {
	ServiceID      *string `json:"serviceID,omitempty"`
	Name           string  `json:"name"`
	TimeoutMinutes int     `json:"timeoutMinutes"`
}

type CreateIntegrationKeyInput struct {
	ServiceID *string            `json:"serviceID,omitempty"`
	Type      IntegrationKeyType `json:"type"`
	Name      string             `json:"name"`
}

type CreateRotationInput struct {
	Name        string        `json:"name"`
	Description *string       `json:"description,omitempty"`
	TimeZone    string        `json:"timeZone"`
	Start       time.Time     `json:"start"`
	Favorite    *bool         `json:"favorite,omitempty"`
	Type        rotation.Type `json:"type"`
	ShiftLength *int          `json:"shiftLength,omitempty"`
	UserIDs     []string      `json:"userIDs,omitempty"`
}

type CreateScheduleInput struct {
	Name             string                    `json:"name"`
	Description      *string                   `json:"description,omitempty"`
	TimeZone         string                    `json:"timeZone"`
	Favorite         *bool                     `json:"favorite,omitempty"`
	Targets          []ScheduleTargetInput     `json:"targets,omitempty"`
	NewUserOverrides []CreateUserOverrideInput `json:"newUserOverrides,omitempty"`
}

type CreateServiceInput struct {
	Name                 string                        `json:"name"`
	Description          *string                       `json:"description,omitempty"`
	Favorite             *bool                         `json:"favorite,omitempty"`
	EscalationPolicyID   *string                       `json:"escalationPolicyID,omitempty"`
	NewEscalationPolicy  *CreateEscalationPolicyInput  `json:"newEscalationPolicy,omitempty"`
	NewIntegrationKeys   []CreateIntegrationKeyInput   `json:"newIntegrationKeys,omitempty"`
	Labels               []SetLabelInput               `json:"labels,omitempty"`
	NewHeartbeatMonitors []CreateHeartbeatMonitorInput `json:"newHeartbeatMonitors,omitempty"`
}

type CreateUserCalendarSubscriptionInput struct {
	Name            string `json:"name"`
	ReminderMinutes []int  `json:"reminderMinutes,omitempty"`
	ScheduleID      string `json:"scheduleID"`
	Disabled        *bool  `json:"disabled,omitempty"`
}

type CreateUserContactMethodInput struct {
	UserID                  string                           `json:"userID"`
	Type                    contactmethod.Type               `json:"type"`
	Name                    string                           `json:"name"`
	Value                   string                           `json:"value"`
	NewUserNotificationRule *CreateUserNotificationRuleInput `json:"newUserNotificationRule,omitempty"`
}

type CreateUserInput struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Name     *string   `json:"name,omitempty"`
	Email    *string   `json:"email,omitempty"`
	Role     *UserRole `json:"role,omitempty"`
	Favorite *bool     `json:"favorite,omitempty"`
}

type CreateUserNotificationRuleInput struct {
	UserID          *string `json:"userID,omitempty"`
	ContactMethodID *string `json:"contactMethodID,omitempty"`
	DelayMinutes    int     `json:"delayMinutes"`
}

type CreateUserOverrideInput struct {
	ScheduleID   *string   `json:"scheduleID,omitempty"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	AddUserID    *string   `json:"addUserID,omitempty"`
	RemoveUserID *string   `json:"removeUserID,omitempty"`
}

type CreatedGQLAPIKey struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

type DebugCarrierInfoInput struct {
	Number string `json:"number"`
}

type DebugMessage struct {
	ID          string     `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	Type        string     `json:"type"`
	Status      string     `json:"status"`
	UserID      *string    `json:"userID,omitempty"`
	UserName    *string    `json:"userName,omitempty"`
	Source      *string    `json:"source,omitempty"`
	Destination string     `json:"destination"`
	ServiceID   *string    `json:"serviceID,omitempty"`
	ServiceName *string    `json:"serviceName,omitempty"`
	AlertID     *int       `json:"alertID,omitempty"`
	ProviderID  *string    `json:"providerID,omitempty"`
	SentAt      *time.Time `json:"sentAt,omitempty"`
	RetryCount  int        `json:"retryCount"`
}

type DebugMessageStatusInfo struct {
	State *NotificationState `json:"state"`
}

type DebugMessageStatusInput struct {
	ProviderMessageID string `json:"providerMessageID"`
}

type DebugMessagesInput struct {
	First         *int       `json:"first,omitempty"`
	CreatedBefore *time.Time `json:"createdBefore,omitempty"`
	CreatedAfter  *time.Time `json:"createdAfter,omitempty"`
}

type DebugSendSMSInfo struct {
	ID          string `json:"id"`
	ProviderURL string `json:"providerURL"`
	FromNumber  string `json:"fromNumber"`
}

type DebugSendSMSInput struct {
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"body"`
}

type EscalationPolicyConnection struct {
	Nodes    []escalation.Policy `json:"nodes"`
	PageInfo *PageInfo           `json:"pageInfo"`
}

type EscalationPolicySearchOptions struct {
	First          *int     `json:"first,omitempty"`
	After          *string  `json:"after,omitempty"`
	Search         *string  `json:"search,omitempty"`
	Omit           []string `json:"omit,omitempty"`
	FavoritesOnly  *bool    `json:"favoritesOnly,omitempty"`
	FavoritesFirst *bool    `json:"favoritesFirst,omitempty"`
}

type GQLAPIKey struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	CreatedAt     time.Time       `json:"createdAt"`
	CreatedBy     *user.User      `json:"createdBy,omitempty"`
	UpdatedAt     time.Time       `json:"updatedAt"`
	UpdatedBy     *user.User      `json:"updatedBy,omitempty"`
	LastUsed      *GQLAPIKeyUsage `json:"lastUsed,omitempty"`
	ExpiresAt     time.Time       `json:"expiresAt"`
	AllowedFields []string        `json:"allowedFields"`
	Role          UserRole        `json:"role"`
}

type GQLAPIKeyUsage struct {
	Time time.Time `json:"time"`
	Ua   string    `json:"ua"`
	IP   string    `json:"ip"`
}

type IntegrationKeyConnection struct {
	Nodes    []integrationkey.IntegrationKey `json:"nodes"`
	PageInfo *PageInfo                       `json:"pageInfo"`
}

type IntegrationKeySearchOptions struct {
	First  *int     `json:"first,omitempty"`
	After  *string  `json:"after,omitempty"`
	Search *string  `json:"search,omitempty"`
	Omit   []string `json:"omit,omitempty"`
}

type IntegrationKeyTypeInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Label   string `json:"label"`
	Enabled bool   `json:"enabled"`
}

type LabelConnection struct {
	Nodes    []label.Label `json:"nodes"`
	PageInfo *PageInfo     `json:"pageInfo"`
}

type LabelKeySearchOptions struct {
	First  *int     `json:"first,omitempty"`
	After  *string  `json:"after,omitempty"`
	Search *string  `json:"search,omitempty"`
	Omit   []string `json:"omit,omitempty"`
}

type LabelSearchOptions struct {
	First      *int     `json:"first,omitempty"`
	After      *string  `json:"after,omitempty"`
	Search     *string  `json:"search,omitempty"`
	UniqueKeys *bool    `json:"uniqueKeys,omitempty"`
	Omit       []string `json:"omit,omitempty"`
}

type LabelValueSearchOptions struct {
	Key    string   `json:"key"`
	First  *int     `json:"first,omitempty"`
	After  *string  `json:"after,omitempty"`
	Search *string  `json:"search,omitempty"`
	Omit   []string `json:"omit,omitempty"`
}

type LinkAccountInfo struct {
	UserDetails    string       `json:"userDetails"`
	AlertID        *int         `json:"alertID,omitempty"`
	AlertNewStatus *AlertStatus `json:"alertNewStatus,omitempty"`
}

type MessageLogConnection struct {
	Nodes    []DebugMessage              `json:"nodes"`
	PageInfo *PageInfo                   `json:"pageInfo"`
	Stats    *notification.SearchOptions `json:"stats"`
}

type MessageLogSearchOptions struct {
	First         *int       `json:"first,omitempty"`
	After         *string    `json:"after,omitempty"`
	CreatedBefore *time.Time `json:"createdBefore,omitempty"`
	CreatedAfter  *time.Time `json:"createdAfter,omitempty"`
	Search        *string    `json:"search,omitempty"`
	Omit          []string   `json:"omit,omitempty"`
}

type NotificationState struct {
	Details           string              `json:"details"`
	Status            *NotificationStatus `json:"status,omitempty"`
	FormattedSrcValue string              `json:"formattedSrcValue"`
}

type PageInfo struct {
	EndCursor   *string `json:"endCursor,omitempty"`
	HasNextPage bool    `json:"hasNextPage"`
}

type PhoneNumberInfo struct {
	ID          string `json:"id"`
	CountryCode string `json:"countryCode"`
	RegionCode  string `json:"regionCode"`
	Formatted   string `json:"formatted"`
	Valid       bool   `json:"valid"`
	Error       string `json:"error"`
}

type RotationConnection struct {
	Nodes    []rotation.Rotation `json:"nodes"`
	PageInfo *PageInfo           `json:"pageInfo"`
}

type RotationSearchOptions struct {
	First          *int     `json:"first,omitempty"`
	After          *string  `json:"after,omitempty"`
	Search         *string  `json:"search,omitempty"`
	Omit           []string `json:"omit,omitempty"`
	FavoritesOnly  *bool    `json:"favoritesOnly,omitempty"`
	FavoritesFirst *bool    `json:"favoritesFirst,omitempty"`
}

type SWOConnection struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
	IsNext  bool   `json:"isNext"`
	Count   int    `json:"count"`
}

type SWONode struct {
	ID       string `json:"id"`
	CanExec  bool   `json:"canExec"`
	IsLeader bool   `json:"isLeader"`
	// The uptime of the node in seconds. Empty if the node/connection is *not* a GoAlert instance in SWO mode.
	Uptime      string          `json:"uptime"`
	ConfigError string          `json:"configError"`
	Connections []SWOConnection `json:"connections,omitempty"`
}

type SWOStatus struct {
	State         SWOState  `json:"state"`
	LastStatus    string    `json:"lastStatus"`
	LastError     string    `json:"lastError"`
	Nodes         []SWONode `json:"nodes"`
	MainDBVersion string    `json:"mainDBVersion"`
	NextDBVersion string    `json:"nextDBVersion"`
}

type ScheduleConnection struct {
	Nodes    []schedule.Schedule `json:"nodes"`
	PageInfo *PageInfo           `json:"pageInfo"`
}

type ScheduleRuleInput struct {
	ID            *string                 `json:"id,omitempty"`
	Start         *timeutil.Clock         `json:"start,omitempty"`
	End           *timeutil.Clock         `json:"end,omitempty"`
	WeekdayFilter *timeutil.WeekdayFilter `json:"weekdayFilter,omitempty"`
}

type ScheduleSearchOptions struct {
	First          *int     `json:"first,omitempty"`
	After          *string  `json:"after,omitempty"`
	Search         *string  `json:"search,omitempty"`
	Omit           []string `json:"omit,omitempty"`
	FavoritesOnly  *bool    `json:"favoritesOnly,omitempty"`
	FavoritesFirst *bool    `json:"favoritesFirst,omitempty"`
}

type ScheduleTarget struct {
	ScheduleID string                `json:"scheduleID"`
	Target     *assignment.RawTarget `json:"target"`
	Rules      []rule.Rule           `json:"rules"`
}

type ScheduleTargetInput struct {
	ScheduleID  *string               `json:"scheduleID,omitempty"`
	Target      *assignment.RawTarget `json:"target,omitempty"`
	NewRotation *CreateRotationInput  `json:"newRotation,omitempty"`
	Rules       []ScheduleRuleInput   `json:"rules"`
}

type SendContactMethodVerificationInput struct {
	ContactMethodID string `json:"contactMethodID"`
}

type ServiceConnection struct {
	Nodes    []service.Service `json:"nodes"`
	PageInfo *PageInfo         `json:"pageInfo"`
}

type ServiceSearchOptions struct {
	First          *int     `json:"first,omitempty"`
	After          *string  `json:"after,omitempty"`
	Search         *string  `json:"search,omitempty"`
	Omit           []string `json:"omit,omitempty"`
	FavoritesOnly  *bool    `json:"favoritesOnly,omitempty"`
	FavoritesFirst *bool    `json:"favoritesFirst,omitempty"`
}

type SetAlertNoiseReasonInput struct {
	AlertID     int    `json:"alertID"`
	NoiseReason string `json:"noiseReason"`
}

type SetFavoriteInput struct {
	Target   *assignment.RawTarget `json:"target"`
	Favorite bool                  `json:"favorite"`
}

type SetLabelInput struct {
	Target *assignment.RawTarget `json:"target,omitempty"`
	Key    string                `json:"key"`
	Value  string                `json:"value"`
}

type SetScheduleOnCallNotificationRulesInput struct {
	ScheduleID string                        `json:"scheduleID"`
	Rules      []OnCallNotificationRuleInput `json:"rules"`
}

type SetTemporaryScheduleInput struct {
	ScheduleID string                `json:"scheduleID"`
	ClearStart *time.Time            `json:"clearStart,omitempty"`
	ClearEnd   *time.Time            `json:"clearEnd,omitempty"`
	Start      time.Time             `json:"start"`
	End        time.Time             `json:"end"`
	Shifts     []schedule.FixedShift `json:"shifts"`
}

type SlackChannelConnection struct {
	Nodes    []slack.Channel `json:"nodes"`
	PageInfo *PageInfo       `json:"pageInfo"`
}

type SlackChannelSearchOptions struct {
	First  *int     `json:"first,omitempty"`
	After  *string  `json:"after,omitempty"`
	Search *string  `json:"search,omitempty"`
	Omit   []string `json:"omit,omitempty"`
}

type SlackUserGroupConnection struct {
	Nodes    []slack.UserGroup `json:"nodes"`
	PageInfo *PageInfo         `json:"pageInfo"`
}

type SlackUserGroupSearchOptions struct {
	First  *int     `json:"first,omitempty"`
	After  *string  `json:"after,omitempty"`
	Search *string  `json:"search,omitempty"`
	Omit   []string `json:"omit,omitempty"`
}

type StringConnection struct {
	Nodes    []string  `json:"nodes"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type SystemLimit struct {
	ID          limit.ID `json:"id"`
	Description string   `json:"description"`
	Value       int      `json:"value"`
}

type SystemLimitInput struct {
	ID    limit.ID `json:"id"`
	Value int      `json:"value"`
}

type TimeSeriesBucket struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Count int       `json:"count"`
}

type TimeSeriesOptions struct {
	BucketDuration timeutil.ISODuration `json:"bucketDuration"`
	BucketOrigin   *time.Time           `json:"bucketOrigin,omitempty"`
}

type TimeZone struct {
	ID string `json:"id"`
}

type TimeZoneConnection struct {
	Nodes    []TimeZone `json:"nodes"`
	PageInfo *PageInfo  `json:"pageInfo"`
}

type TimeZoneSearchOptions struct {
	First  *int     `json:"first,omitempty"`
	After  *string  `json:"after,omitempty"`
	Search *string  `json:"search,omitempty"`
	Omit   []string `json:"omit,omitempty"`
}

type UpdateAlertsByServiceInput struct {
	ServiceID string      `json:"serviceID"`
	NewStatus AlertStatus `json:"newStatus"`
}

type UpdateAlertsInput struct {
	AlertIDs    []int        `json:"alertIDs"`
	NewStatus   *AlertStatus `json:"newStatus,omitempty"`
	NoiseReason *string      `json:"noiseReason,omitempty"`
}

type UpdateBasicAuthInput struct {
	Password    string  `json:"password"`
	OldPassword *string `json:"oldPassword,omitempty"`
	UserID      string  `json:"userID"`
}

type UpdateEscalationPolicyInput struct {
	ID          string   `json:"id"`
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Repeat      *int     `json:"repeat,omitempty"`
	StepIDs     []string `json:"stepIDs,omitempty"`
}

type UpdateEscalationPolicyStepInput struct {
	ID           string                 `json:"id"`
	DelayMinutes *int                   `json:"delayMinutes,omitempty"`
	Targets      []assignment.RawTarget `json:"targets,omitempty"`
}

type UpdateGQLAPIKeyInput struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateHeartbeatMonitorInput struct {
	ID             string  `json:"id"`
	Name           *string `json:"name,omitempty"`
	TimeoutMinutes *int    `json:"timeoutMinutes,omitempty"`
}

type UpdateRotationInput struct {
	ID              string         `json:"id"`
	Name            *string        `json:"name,omitempty"`
	Description     *string        `json:"description,omitempty"`
	TimeZone        *string        `json:"timeZone,omitempty"`
	Start           *time.Time     `json:"start,omitempty"`
	Type            *rotation.Type `json:"type,omitempty"`
	ShiftLength     *int           `json:"shiftLength,omitempty"`
	ActiveUserIndex *int           `json:"activeUserIndex,omitempty"`
	UserIDs         []string       `json:"userIDs,omitempty"`
}

type UpdateScheduleInput struct {
	ID          string  `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	TimeZone    *string `json:"timeZone,omitempty"`
}

type UpdateServiceInput struct {
	ID                   string     `json:"id"`
	Name                 *string    `json:"name,omitempty"`
	Description          *string    `json:"description,omitempty"`
	EscalationPolicyID   *string    `json:"escalationPolicyID,omitempty"`
	MaintenanceExpiresAt *time.Time `json:"maintenanceExpiresAt,omitempty"`
}

type UpdateUserCalendarSubscriptionInput struct {
	ID              string  `json:"id"`
	Name            *string `json:"name,omitempty"`
	ReminderMinutes []int   `json:"reminderMinutes,omitempty"`
	Disabled        *bool   `json:"disabled,omitempty"`
}

type UpdateUserContactMethodInput struct {
	ID                  string  `json:"id"`
	Name                *string `json:"name,omitempty"`
	Value               *string `json:"value,omitempty"`
	EnableStatusUpdates *bool   `json:"enableStatusUpdates,omitempty"`
}

type UpdateUserInput struct {
	ID                          string    `json:"id"`
	Name                        *string   `json:"name,omitempty"`
	Email                       *string   `json:"email,omitempty"`
	Role                        *UserRole `json:"role,omitempty"`
	StatusUpdateContactMethodID *string   `json:"statusUpdateContactMethodID,omitempty"`
}

type UpdateUserOverrideInput struct {
	ID           string     `json:"id"`
	Start        *time.Time `json:"start,omitempty"`
	End          *time.Time `json:"end,omitempty"`
	AddUserID    *string    `json:"addUserID,omitempty"`
	RemoveUserID *string    `json:"removeUserID,omitempty"`
}

type UserConnection struct {
	Nodes    []user.User `json:"nodes"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type UserOverrideConnection struct {
	Nodes    []override.UserOverride `json:"nodes"`
	PageInfo *PageInfo               `json:"pageInfo"`
}

type UserOverrideSearchOptions struct {
	First              *int       `json:"first,omitempty"`
	After              *string    `json:"after,omitempty"`
	Omit               []string   `json:"omit,omitempty"`
	ScheduleID         *string    `json:"scheduleID,omitempty"`
	FilterAddUserID    []string   `json:"filterAddUserID,omitempty"`
	FilterRemoveUserID []string   `json:"filterRemoveUserID,omitempty"`
	FilterAnyUserID    []string   `json:"filterAnyUserID,omitempty"`
	Start              *time.Time `json:"start,omitempty"`
	End                *time.Time `json:"end,omitempty"`
}

type UserSearchOptions struct {
	First          *int                `json:"first,omitempty"`
	After          *string             `json:"after,omitempty"`
	Search         *string             `json:"search,omitempty"`
	Omit           []string            `json:"omit,omitempty"`
	CMValue        *string             `json:"CMValue,omitempty"`
	CMType         *contactmethod.Type `json:"CMType,omitempty"`
	FavoritesOnly  *bool               `json:"favoritesOnly,omitempty"`
	FavoritesFirst *bool               `json:"favoritesFirst,omitempty"`
}

type UserSession struct {
	ID           string    `json:"id"`
	Current      bool      `json:"current"`
	UserAgent    string    `json:"userAgent"`
	CreatedAt    time.Time `json:"createdAt"`
	LastAccessAt time.Time `json:"lastAccessAt"`
}

type VerifyContactMethodInput struct {
	ContactMethodID string `json:"contactMethodID"`
	Code            int    `json:"code"`
}

type AlertSearchSort string

const (
	AlertSearchSortStatusID      AlertSearchSort = "statusID"
	AlertSearchSortDateID        AlertSearchSort = "dateID"
	AlertSearchSortDateIDReverse AlertSearchSort = "dateIDReverse"
)

var AllAlertSearchSort = []AlertSearchSort{
	AlertSearchSortStatusID,
	AlertSearchSortDateID,
	AlertSearchSortDateIDReverse,
}

func (e AlertSearchSort) IsValid() bool {
	switch e {
	case AlertSearchSortStatusID, AlertSearchSortDateID, AlertSearchSortDateIDReverse:
		return true
	}
	return false
}

func (e AlertSearchSort) String() string {
	return string(e)
}

func (e *AlertSearchSort) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AlertSearchSort(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AlertSearchSort", str)
	}
	return nil
}

func (e AlertSearchSort) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AlertStatus string

const (
	AlertStatusStatusAcknowledged   AlertStatus = "StatusAcknowledged"
	AlertStatusStatusClosed         AlertStatus = "StatusClosed"
	AlertStatusStatusUnacknowledged AlertStatus = "StatusUnacknowledged"
)

var AllAlertStatus = []AlertStatus{
	AlertStatusStatusAcknowledged,
	AlertStatusStatusClosed,
	AlertStatusStatusUnacknowledged,
}

func (e AlertStatus) IsValid() bool {
	switch e {
	case AlertStatusStatusAcknowledged, AlertStatusStatusClosed, AlertStatusStatusUnacknowledged:
		return true
	}
	return false
}

func (e AlertStatus) String() string {
	return string(e)
}

func (e *AlertStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AlertStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AlertStatus", str)
	}
	return nil
}

func (e AlertStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ConfigType string

const (
	ConfigTypeString     ConfigType = "string"
	ConfigTypeStringList ConfigType = "stringList"
	ConfigTypeInteger    ConfigType = "integer"
	ConfigTypeBoolean    ConfigType = "boolean"
)

var AllConfigType = []ConfigType{
	ConfigTypeString,
	ConfigTypeStringList,
	ConfigTypeInteger,
	ConfigTypeBoolean,
}

func (e ConfigType) IsValid() bool {
	switch e {
	case ConfigTypeString, ConfigTypeStringList, ConfigTypeInteger, ConfigTypeBoolean:
		return true
	}
	return false
}

func (e ConfigType) String() string {
	return string(e)
}

func (e *ConfigType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ConfigType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ConfigType", str)
	}
	return nil
}

func (e ConfigType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type IntegrationKeyType string

const (
	IntegrationKeyTypeGeneric                IntegrationKeyType = "generic"
	IntegrationKeyTypeGrafana                IntegrationKeyType = "grafana"
	IntegrationKeyTypeSite24x7               IntegrationKeyType = "site24x7"
	IntegrationKeyTypePrometheusAlertmanager IntegrationKeyType = "prometheusAlertmanager"
	IntegrationKeyTypeEmail                  IntegrationKeyType = "email"
)

var AllIntegrationKeyType = []IntegrationKeyType{
	IntegrationKeyTypeGeneric,
	IntegrationKeyTypeGrafana,
	IntegrationKeyTypeSite24x7,
	IntegrationKeyTypePrometheusAlertmanager,
	IntegrationKeyTypeEmail,
}

func (e IntegrationKeyType) IsValid() bool {
	switch e {
	case IntegrationKeyTypeGeneric, IntegrationKeyTypeGrafana, IntegrationKeyTypeSite24x7, IntegrationKeyTypePrometheusAlertmanager, IntegrationKeyTypeEmail:
		return true
	}
	return false
}

func (e IntegrationKeyType) String() string {
	return string(e)
}

func (e *IntegrationKeyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IntegrationKeyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid IntegrationKeyType", str)
	}
	return nil
}

func (e IntegrationKeyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NotificationStatus string

const (
	NotificationStatusOk    NotificationStatus = "OK"
	NotificationStatusWarn  NotificationStatus = "WARN"
	NotificationStatusError NotificationStatus = "ERROR"
)

var AllNotificationStatus = []NotificationStatus{
	NotificationStatusOk,
	NotificationStatusWarn,
	NotificationStatusError,
}

func (e NotificationStatus) IsValid() bool {
	switch e {
	case NotificationStatusOk, NotificationStatusWarn, NotificationStatusError:
		return true
	}
	return false
}

func (e NotificationStatus) String() string {
	return string(e)
}

func (e *NotificationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NotificationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NotificationStatus", str)
	}
	return nil
}

func (e NotificationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SWOAction string

const (
	SWOActionReset   SWOAction = "reset"
	SWOActionExecute SWOAction = "execute"
)

var AllSWOAction = []SWOAction{
	SWOActionReset,
	SWOActionExecute,
}

func (e SWOAction) IsValid() bool {
	switch e {
	case SWOActionReset, SWOActionExecute:
		return true
	}
	return false
}

func (e SWOAction) String() string {
	return string(e)
}

func (e *SWOAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SWOAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SWOAction", str)
	}
	return nil
}

func (e SWOAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SWOState string

const (
	SWOStateUnknown   SWOState = "unknown"
	SWOStateResetting SWOState = "resetting"
	SWOStateIdle      SWOState = "idle"
	SWOStateSyncing   SWOState = "syncing"
	SWOStatePausing   SWOState = "pausing"
	SWOStateExecuting SWOState = "executing"
	SWOStateDone      SWOState = "done"
)

var AllSWOState = []SWOState{
	SWOStateUnknown,
	SWOStateResetting,
	SWOStateIdle,
	SWOStateSyncing,
	SWOStatePausing,
	SWOStateExecuting,
	SWOStateDone,
}

func (e SWOState) IsValid() bool {
	switch e {
	case SWOStateUnknown, SWOStateResetting, SWOStateIdle, SWOStateSyncing, SWOStatePausing, SWOStateExecuting, SWOStateDone:
		return true
	}
	return false
}

func (e SWOState) String() string {
	return string(e)
}

func (e *SWOState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SWOState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SWOState", str)
	}
	return nil
}

func (e SWOState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StatusUpdateState string

const (
	StatusUpdateStateDisabled       StatusUpdateState = "DISABLED"
	StatusUpdateStateEnabled        StatusUpdateState = "ENABLED"
	StatusUpdateStateEnabledForced  StatusUpdateState = "ENABLED_FORCED"
	StatusUpdateStateDisabledForced StatusUpdateState = "DISABLED_FORCED"
)

var AllStatusUpdateState = []StatusUpdateState{
	StatusUpdateStateDisabled,
	StatusUpdateStateEnabled,
	StatusUpdateStateEnabledForced,
	StatusUpdateStateDisabledForced,
}

func (e StatusUpdateState) IsValid() bool {
	switch e {
	case StatusUpdateStateDisabled, StatusUpdateStateEnabled, StatusUpdateStateEnabledForced, StatusUpdateStateDisabledForced:
		return true
	}
	return false
}

func (e StatusUpdateState) String() string {
	return string(e)
}

func (e *StatusUpdateState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StatusUpdateState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StatusUpdateState", str)
	}
	return nil
}

func (e StatusUpdateState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserRole string

const (
	UserRoleUnknown UserRole = "unknown"
	UserRoleUser    UserRole = "user"
	UserRoleAdmin   UserRole = "admin"
)

var AllUserRole = []UserRole{
	UserRoleUnknown,
	UserRoleUser,
	UserRoleAdmin,
}

func (e UserRole) IsValid() bool {
	switch e {
	case UserRoleUnknown, UserRoleUser, UserRoleAdmin:
		return true
	}
	return false
}

func (e UserRole) String() string {
	return string(e)
}

func (e *UserRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRole", str)
	}
	return nil
}

func (e UserRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
