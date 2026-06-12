package aws

import (
    "time"
)

// AWSCredentials - minimal representation of AWS access
// Think of this as the key to your AWS account
type AWSCredentials struct {
    AccessKeyID     string
    SecretAccessKey string
    SessionToken    string  // Only for temporary credentials
    Region          string
}

// AssumeRoleInput - what we need to assume a role in another account
// Like having a VIP pass that lets you into different departments
type AssumeRoleInput struct {
    RoleARN         string        // aws:arn:iam::123456789012:role/OrganizationAccountAccessRole
    SessionName     string        // What we call this session (e.g., "cost-analyzer")
    DurationSeconds int           // How long the pass is valid (max 3600 for most roles)
    ExternalID      string        // Optional: extra security for cross-account access
}

// AccountInfo - minimal AWS account info from Organizations
// This is AWS's view of an account, not ours
type AccountInfo struct {
    ID          string    // 12-digit AWS account ID
    Name        string    // Human-readable name
    Email       string    // Contact email
    Status      string    // ACTIVE, SUSPENDED, etc.
    JoinedMethod string   // INVITED or CREATED
    JoinedAt    time.Time
}

// CURManifest - points to where the actual cost data lives
// Like a table of contents for your bill
type CURManifest struct {
    Bucket          string    // S3 bucket name
    Key             string    // Path to manifest file
    ReportName      string    // Name of the CUR report
    Compression     string    // PARQUET, CSV, etc.
    Columns         []string  // What data fields are available
    RefreshInterval string    // HOURLY, DAILY, etc.
}

// CURRecord - one line of cost data
// We only care about the fields we actually use
type CURRecord struct {
    IdentityTimeInterval string    // When the cost was incurred
    BillInvoiceID        string    
    LineItemUsageStartDate time.Time
    LineItemUsageEndDate   time.Time
    
    LineItemAccountID      string   // Which AWS account
    LineItemServiceCode    string   // Which service (EC2, S3, etc.)
    LineItemResourceID     string   // Specific resource (i-12345, bucket-name)
    LineItemUsageType      string   
    LineItemOperation      string   
    LineItemAvailabilityZone string // Region or AZ
    
    PricingUnit            string   // Hours, GB, Requests
    PricingPublicOnDemandRate float64
    
    ReservationARN         string   // If using reserved instances
    
    ProductProductName     string   // Human readable service name
    ProductRegion          string
    
    CostBeforeTax          float64
    Cost                   float64
    Currency               string
    
    UsageQuantity          float64
}