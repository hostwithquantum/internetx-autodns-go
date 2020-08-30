// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hostwithquantum/internetx-autodns-go/client/accounting_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/backup_mx_bulk_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/backup_mx_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/certificate_bulk_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/certificate_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/contact_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/customer_account_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/document_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/domain_bulk_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/domain_prereg_bulk_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/domain_prereg_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/domain_safe_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/domain_studio"
	"github.com/hostwithquantum/internetx-autodns-go/client/domain_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/guest_account_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/id4me_agent_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/id4me_identity_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/invoice_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/job_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/mail_proxy_bulk_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/mail_proxy_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/mail_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/object_user_assignment_bulk_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/object_user_assignment_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/poll_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/redirect_bulk_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/redirect_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/s_s_l_contact_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/session_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/ssl_contact_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/subject_product_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/subscription_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/tmch_mark_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/transfer_request_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/trusted_application_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/user_2fa_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/user_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/zone_bulk_tasks"
	"github.com/hostwithquantum/internetx-autodns-go/client/zone_tasks"
)

// Default internetx autodns go HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.autodns.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new internetx autodns go HTTP client.
func NewHTTPClient(formats strfmt.Registry) *InternetxAutodnsGo {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new internetx autodns go HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *InternetxAutodnsGo {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new internetx autodns go client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *InternetxAutodnsGo {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(InternetxAutodnsGo)
	cli.Transport = transport
	cli.AccountingTasks = accounting_tasks.New(transport, formats)
	cli.BackupMxBulkTasks = backup_mx_bulk_tasks.New(transport, formats)
	cli.BackupMxTasks = backup_mx_tasks.New(transport, formats)
	cli.CertificateBulkTasks = certificate_bulk_tasks.New(transport, formats)
	cli.CertificateTasks = certificate_tasks.New(transport, formats)
	cli.ContactTasks = contact_tasks.New(transport, formats)
	cli.CustomerAccountTasks = customer_account_tasks.New(transport, formats)
	cli.DocumentTasks = document_tasks.New(transport, formats)
	cli.DomainBulkTasks = domain_bulk_tasks.New(transport, formats)
	cli.DomainPreregBulkTasks = domain_prereg_bulk_tasks.New(transport, formats)
	cli.DomainPreregTasks = domain_prereg_tasks.New(transport, formats)
	cli.DomainSafeTasks = domain_safe_tasks.New(transport, formats)
	cli.DomainStudio = domain_studio.New(transport, formats)
	cli.DomainTasks = domain_tasks.New(transport, formats)
	cli.GuestAccountTasks = guest_account_tasks.New(transport, formats)
	cli.Id4meAgentTasks = id4me_agent_tasks.New(transport, formats)
	cli.Id4meIdentityTasks = id4me_identity_tasks.New(transport, formats)
	cli.InvoiceTasks = invoice_tasks.New(transport, formats)
	cli.JobTasks = job_tasks.New(transport, formats)
	cli.MailProxyBulkTasks = mail_proxy_bulk_tasks.New(transport, formats)
	cli.MailProxyTasks = mail_proxy_tasks.New(transport, formats)
	cli.MailTasks = mail_tasks.New(transport, formats)
	cli.ObjectUserAssignmentBulkTasks = object_user_assignment_bulk_tasks.New(transport, formats)
	cli.ObjectUserAssignmentTasks = object_user_assignment_tasks.New(transport, formats)
	cli.PollTasks = poll_tasks.New(transport, formats)
	cli.RedirectBulkTasks = redirect_bulk_tasks.New(transport, formats)
	cli.RedirectTasks = redirect_tasks.New(transport, formats)
	cli.SslContactTasks = s_s_l_contact_tasks.New(transport, formats)
	cli.SessionTasks = session_tasks.New(transport, formats)
	cli.SslContactTasks = ssl_contact_tasks.New(transport, formats)
	cli.SubjectProductTasks = subject_product_tasks.New(transport, formats)
	cli.SubscriptionTasks = subscription_tasks.New(transport, formats)
	cli.TmchMarkTasks = tmch_mark_tasks.New(transport, formats)
	cli.TransferRequestTasks = transfer_request_tasks.New(transport, formats)
	cli.TrustedApplicationTasks = trusted_application_tasks.New(transport, formats)
	cli.User2faTasks = user_2fa_tasks.New(transport, formats)
	cli.UserTasks = user_tasks.New(transport, formats)
	cli.ZoneBulkTasks = zone_bulk_tasks.New(transport, formats)
	cli.ZoneTasks = zone_tasks.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// InternetxAutodnsGo is a client for internetx autodns go
type InternetxAutodnsGo struct {
	AccountingTasks accounting_tasks.ClientService

	BackupMxBulkTasks backup_mx_bulk_tasks.ClientService

	BackupMxTasks backup_mx_tasks.ClientService

	CertificateBulkTasks certificate_bulk_tasks.ClientService

	CertificateTasks certificate_tasks.ClientService

	ContactTasks contact_tasks.ClientService

	CustomerAccountTasks customer_account_tasks.ClientService

	DocumentTasks document_tasks.ClientService

	DomainBulkTasks domain_bulk_tasks.ClientService

	DomainPreregBulkTasks domain_prereg_bulk_tasks.ClientService

	DomainPreregTasks domain_prereg_tasks.ClientService

	DomainSafeTasks domain_safe_tasks.ClientService

	DomainStudio domain_studio.ClientService

	DomainTasks domain_tasks.ClientService

	GuestAccountTasks guest_account_tasks.ClientService

	Id4meAgentTasks id4me_agent_tasks.ClientService

	Id4meIdentityTasks id4me_identity_tasks.ClientService

	InvoiceTasks invoice_tasks.ClientService

	JobTasks job_tasks.ClientService

	MailProxyBulkTasks mail_proxy_bulk_tasks.ClientService

	MailProxyTasks mail_proxy_tasks.ClientService

	MailTasks mail_tasks.ClientService

	ObjectUserAssignmentBulkTasks object_user_assignment_bulk_tasks.ClientService

	ObjectUserAssignmentTasks object_user_assignment_tasks.ClientService

	PollTasks poll_tasks.ClientService

	RedirectBulkTasks redirect_bulk_tasks.ClientService

	RedirectTasks redirect_tasks.ClientService

	SslContactTasks s_s_l_contact_tasks.ClientService

	SessionTasks session_tasks.ClientService

	SslContactTasks ssl_contact_tasks.ClientService

	SubjectProductTasks subject_product_tasks.ClientService

	SubscriptionTasks subscription_tasks.ClientService

	TmchMarkTasks tmch_mark_tasks.ClientService

	TransferRequestTasks transfer_request_tasks.ClientService

	TrustedApplicationTasks trusted_application_tasks.ClientService

	User2faTasks user_2fa_tasks.ClientService

	UserTasks user_tasks.ClientService

	ZoneBulkTasks zone_bulk_tasks.ClientService

	ZoneTasks zone_tasks.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *InternetxAutodnsGo) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AccountingTasks.SetTransport(transport)
	c.BackupMxBulkTasks.SetTransport(transport)
	c.BackupMxTasks.SetTransport(transport)
	c.CertificateBulkTasks.SetTransport(transport)
	c.CertificateTasks.SetTransport(transport)
	c.ContactTasks.SetTransport(transport)
	c.CustomerAccountTasks.SetTransport(transport)
	c.DocumentTasks.SetTransport(transport)
	c.DomainBulkTasks.SetTransport(transport)
	c.DomainPreregBulkTasks.SetTransport(transport)
	c.DomainPreregTasks.SetTransport(transport)
	c.DomainSafeTasks.SetTransport(transport)
	c.DomainStudio.SetTransport(transport)
	c.DomainTasks.SetTransport(transport)
	c.GuestAccountTasks.SetTransport(transport)
	c.Id4meAgentTasks.SetTransport(transport)
	c.Id4meIdentityTasks.SetTransport(transport)
	c.InvoiceTasks.SetTransport(transport)
	c.JobTasks.SetTransport(transport)
	c.MailProxyBulkTasks.SetTransport(transport)
	c.MailProxyTasks.SetTransport(transport)
	c.MailTasks.SetTransport(transport)
	c.ObjectUserAssignmentBulkTasks.SetTransport(transport)
	c.ObjectUserAssignmentTasks.SetTransport(transport)
	c.PollTasks.SetTransport(transport)
	c.RedirectBulkTasks.SetTransport(transport)
	c.RedirectTasks.SetTransport(transport)
	c.SslContactTasks.SetTransport(transport)
	c.SessionTasks.SetTransport(transport)
	c.SslContactTasks.SetTransport(transport)
	c.SubjectProductTasks.SetTransport(transport)
	c.SubscriptionTasks.SetTransport(transport)
	c.TmchMarkTasks.SetTransport(transport)
	c.TransferRequestTasks.SetTransport(transport)
	c.TrustedApplicationTasks.SetTransport(transport)
	c.User2faTasks.SetTransport(transport)
	c.UserTasks.SetTransport(transport)
	c.ZoneBulkTasks.SetTransport(transport)
	c.ZoneTasks.SetTransport(transport)
}