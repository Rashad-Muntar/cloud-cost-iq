package aws

import (
    "context"
    
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/sts"
)

// STSClient - simple wrapper for AWS STS (Security Token Service)
// STS is how you get temporary credentials for cross-account access
type STSClient struct {
    client *sts.Client
}

func NewSTSClient(cfg aws.Config) *STSClient {
    return &STSClient{
        client: sts.NewFromConfig(cfg),
    }
}

// AssumeRole - get temporary credentials for another AWS account
func (c *STSClient) AssumeRole(ctx context.Context, input AssumeRoleInput) (*AWSCredentials, error) {
    duration := int32(input.DurationSeconds)
    if duration == 0 {
        duration = 3600 // Default to 1 hour
    }
    
    assumeInput := &sts.AssumeRoleInput{
        RoleArn:         aws.String(input.RoleARN),
        RoleSessionName: aws.String(input.SessionName),
        DurationSeconds: &duration,
    }
    
    if input.ExternalID != "" {
        assumeInput.ExternalId = aws.String(input.ExternalID)
    }
    
    result, err := c.client.AssumeRole(ctx, assumeInput)
    if err != nil {
        return nil, err
    }
    
    return &AWSCredentials{
        AccessKeyID:     *result.Credentials.AccessKeyId,
        SecretAccessKey: *result.Credentials.SecretAccessKey,
        SessionToken:    *result.Credentials.SessionToken,
    }, nil
}