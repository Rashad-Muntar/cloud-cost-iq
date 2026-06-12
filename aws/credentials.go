package aws

import (
    "context"
    "fmt"
    
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/aws"
)

// CredentialsManager - handles AWS authentication
// Think of this as a secure vault for your AWS keys
type CredentialsManager struct {
    staticCreds *AWSCredentials  // Optional: for static credentials
    region      string
}

func NewCredentialsManager(region string, creds *AWSCredentials) *CredentialsManager {
    return &CredentialsManager{
        staticCreds: creds,
        region:      region,
    }
}

// GetAWSCredentials - returns AWS SDK compatible credentials
// This abstracts whether we use static creds or assume roles
func (m *CredentialsManager) GetAWSCredentials(ctx context.Context, roleARN string) (aws.Config, error) {
    // If we're assuming a role in another account
    if roleARN != "" {
        return m.getRoleCredentials(ctx, roleARN)
    }
    
    // Otherwise use static credentials for the management account
    return m.getStaticCredentials()
}

// getStaticCredentials - for the management account
func (m *CredentialsManager) getStaticCredentials() (aws.Config, error) {
    if m.staticCreds == nil {
        return aws.Config{}, fmt.Errorf("no static credentials provided")
    }
    
    credProvider := credentials.NewStaticCredentialsProvider(
        m.staticCreds.AccessKeyID,
        m.staticCreds.SecretAccessKey,
        m.staticCreds.SessionToken,
    )
    
    cfg, err := config.LoadDefaultConfig(context.Background(),
        config.WithRegion(m.region),
        config.WithCredentialsProvider(credProvider),
    )
    
    return cfg, err
}

// getRoleCredentials - assume a role in another account
// This is how we securely access child accounts without storing their keys
func (m *CredentialsManager) getRoleCredentials(ctx context.Context, roleARN string) (aws.Config, error) {
    // First get credentials for management account
    baseCfg, err := m.getStaticCredentials()
    if err != nil {
        return aws.Config{}, err
    }
    
    // Then assume the target role
    stsClient := NewSTSClient(baseCfg)
    
    input := AssumeRoleInput{
        RoleARN:         roleARN,
        SessionName:     "CloudCostIQ",
        DurationSeconds: 3600, // 1 hour is safe
    }
    
    creds, err := stsClient.AssumeRole(ctx, input)
    if err != nil {
        return aws.Config{}, fmt.Errorf("failed to assume role %s: %w", roleARN, err)
    }
    
    // Create new config with assumed role credentials
    cfg, err := config.LoadDefaultConfig(ctx,
        config.WithRegion(m.region),
        config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
            creds.AccessKeyID,
            creds.SecretAccessKey,
            creds.SessionToken,
        )),
    )
    
    return cfg, err
}