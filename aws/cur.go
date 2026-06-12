package aws

import (
    "context"
    "encoding/json"
    "fmt"
	"io"
    "strings"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

// CURClient - reads Cost and Usage Report data from S3
// This is our main data source for billing information
type CURClient struct {
    s3Client *s3.Client
}

func NewCURClient(cfg aws.Config) *CURClient {
    return &CURClient{
        s3Client: s3.NewFromConfig(cfg),
    }
}

// GetManifest - read the manifest file that tells us where cost data lives
// The manifest is like a README for the cost report
func (c *CURClient) GetManifest(ctx context.Context, bucket, key string) (*CURManifest, error) {
    input := &s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    }
    
    output, err := c.s3Client.GetObject(ctx, input)
    if err != nil {
        return nil, fmt.Errorf("failed to get manifest: %w", err)
    }
    defer output.Body.Close()
    
    var manifest struct {
        ReportName string `json:"reportName"`
        Compression string `json:"compression"`
        Columns     []string `json:"columns"`
        Bucket      string `json:"bucket"`
        Prefix      string `json:"prefix"`
    }
    
    if err := json.NewDecoder(output.Body).Decode(&manifest); err != nil {
        return nil, fmt.Errorf("failed to parse manifest: %w", err)
    }
    
    return &CURManifest{
        Bucket:          manifest.Bucket,
        Key:             key,
        ReportName:      manifest.ReportName,
        Compression:     manifest.Compression,
        Columns:         manifest.Columns,
        RefreshInterval: "DAILY", // Default, could parse from manifest
    }, nil
}

// ListCostFiles - list all available cost data files for a date range
// CUR generates files like: report-name/2024/01/01/report-name_00001.parquet
func (c *CURClient) ListCostFiles(ctx context.Context, manifest *CURManifest, year, month int) ([]string, error) {
    prefix := fmt.Sprintf("%s/%04d/%02d/", manifest.ReportName, year, month)
    
    input := &s3.ListObjectsV2Input{
        Bucket: aws.String(manifest.Bucket),
        Prefix: aws.String(prefix),
    }
    
    var files []string
    paginator := s3.NewListObjectsV2Paginator(c.s3Client, input)
    
    for paginator.HasMorePages() {
        page, err := paginator.NextPage(ctx)
        if err != nil {
            return nil, err
        }
        
        for _, obj := range page.Contents {
            // Only include data files, skip manifest and schema files
            if *obj.Key != manifest.Key && !strings.HasSuffix(*obj.Key, ".json") {
                files = append(files, *obj.Key)
            }
        }
    }
    
    return files, nil
}

// ReadCostFile - stream and parse cost data from S3
// For now, we'll return raw bytes. The ETL layer will parse based on format
func (c *CURClient) ReadCostFile(ctx context.Context, bucket, key string) ([]byte, error) {
    input := &s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    }
    
    output, err := c.s3Client.GetObject(ctx, input)
    if err != nil {
        return nil, fmt.Errorf("failed to get cost file: %w", err)
    }
    defer output.Body.Close()
    
    // In production, we'd stream this instead of reading all at once
    // But for simplicity in this layer, we'll read fully
    data, err := io.ReadAll(output.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read cost file: %w", err)
    }
    
    return data, nil
}