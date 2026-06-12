package aws

import (
    "context"
    
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/organizations"
)

// OrganizationsClient - discover all AWS accounts in an organization
// Like reading the company directory to find all employees
type OrganizationsClient struct {
    client *organizations.Client
}

func NewOrganizationsClient(cfg aws.Config) *OrganizationsClient {
    return &OrganizationsClient{
        client: organizations.NewFromConfig(cfg),
    }
}

// ListAllAccounts - get every account in the AWS Organization
// Handles pagination automatically (AWS returns 20 at a time)
func (c *OrganizationsClient) ListAllAccounts(ctx context.Context) ([]AccountInfo, error) {
    var accounts []AccountInfo
    var nextToken *string
    
    for {
        input := &organizations.ListAccountsInput{
            NextToken: nextToken,
        }
        
        output, err := c.client.ListAccounts(ctx, input)
        if err != nil {
            return nil, err
        }
        
        for _, account := range output.Accounts {
            accounts = append(accounts, AccountInfo{
                ID:          *account.Id,
                Name:        *account.Name,
                Email:       *account.Email,
                Status:      string(account.Status),
                JoinedMethod: string(account.JoinedMethod),
                JoinedAt:    *account.JoinedTimestamp,
            })
        }
        
        nextToken = output.NextToken
        if nextToken == nil {
            break
        }
    }
    
    return accounts, nil
}

// GetAccount - get a single account by ID
func (c *OrganizationsClient) GetAccount(ctx context.Context, accountID string) (*AccountInfo, error) {
    input := &organizations.DescribeAccountInput{
        AccountId: aws.String(accountID),
    }
    
    output, err := c.client.DescribeAccount(ctx, input)
    if err != nil {
        return nil, err
    }
    
    return &AccountInfo{
        ID:          *output.Account.Id,
        Name:        *output.Account.Name,
        Email:       *output.Account.Email,
        Status:      string(output.Account.Status),
        JoinedMethod: string(output.Account.JoinedMethod),
        JoinedAt:    *output.Account.JoinedTimestamp,
    }, nil
}