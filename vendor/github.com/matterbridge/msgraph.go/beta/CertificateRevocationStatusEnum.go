// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CertificateRevocationStatus undocumented
type CertificateRevocationStatus int

const (
	// CertificateRevocationStatusVNone undocumented
	CertificateRevocationStatusVNone CertificateRevocationStatus = 0
	// CertificateRevocationStatusVPending undocumented
	CertificateRevocationStatusVPending CertificateRevocationStatus = 1
	// CertificateRevocationStatusVIssued undocumented
	CertificateRevocationStatusVIssued CertificateRevocationStatus = 2
	// CertificateRevocationStatusVFailed undocumented
	CertificateRevocationStatusVFailed CertificateRevocationStatus = 3
	// CertificateRevocationStatusVRevoked undocumented
	CertificateRevocationStatusVRevoked CertificateRevocationStatus = 4
)

// CertificateRevocationStatusPNone returns a pointer to CertificateRevocationStatusVNone
func CertificateRevocationStatusPNone() *CertificateRevocationStatus {
	v := CertificateRevocationStatusVNone
	return &v
}

// CertificateRevocationStatusPPending returns a pointer to CertificateRevocationStatusVPending
func CertificateRevocationStatusPPending() *CertificateRevocationStatus {
	v := CertificateRevocationStatusVPending
	return &v
}

// CertificateRevocationStatusPIssued returns a pointer to CertificateRevocationStatusVIssued
func CertificateRevocationStatusPIssued() *CertificateRevocationStatus {
	v := CertificateRevocationStatusVIssued
	return &v
}

// CertificateRevocationStatusPFailed returns a pointer to CertificateRevocationStatusVFailed
func CertificateRevocationStatusPFailed() *CertificateRevocationStatus {
	v := CertificateRevocationStatusVFailed
	return &v
}

// CertificateRevocationStatusPRevoked returns a pointer to CertificateRevocationStatusVRevoked
func CertificateRevocationStatusPRevoked() *CertificateRevocationStatus {
	v := CertificateRevocationStatusVRevoked
	return &v
}