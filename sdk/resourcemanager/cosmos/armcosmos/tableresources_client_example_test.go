//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1ad29756bd141a47cac770140105a706d065ae1b/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBTableList.json
func ExampleTableResourcesClient_NewListTablesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTableResourcesClient().NewListTablesPager("rgName", "ddb1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.TableListResult = armcosmos.TableListResult{
		// 	Value: []*armcosmos.TableGetResults{
		// 		{
		// 			Name: to.Ptr("tableName"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/tables"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/tables/tableName"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcosmos.TableGetProperties{
		// 				Resource: &armcosmos.TableGetPropertiesResource{
		// 					ID: to.Ptr("tableName"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1ad29756bd141a47cac770140105a706d065ae1b/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBTableGet.json
func ExampleTableResourcesClient_GetTable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTableResourcesClient().GetTable(ctx, "rg1", "ddb1", "tableName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TableGetResults = armcosmos.TableGetResults{
	// 	Name: to.Ptr("tableName"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/tables"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/tables/tableName"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.TableGetProperties{
	// 		Resource: &armcosmos.TableGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			ID: to.Ptr("tableName"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1ad29756bd141a47cac770140105a706d065ae1b/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBTableCreateUpdate.json
func ExampleTableResourcesClient_BeginCreateUpdateTable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTableResourcesClient().BeginCreateUpdateTable(ctx, "rg1", "ddb1", "tableName", armcosmos.TableCreateUpdateParameters{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.TableCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{},
			Resource: &armcosmos.TableResource{
				ID: to.Ptr("tableName"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TableGetResults = armcosmos.TableGetResults{
	// 	Name: to.Ptr("tableName"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/tables"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/tables/tableName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcosmos.TableGetProperties{
	// 		Resource: &armcosmos.TableGetPropertiesResource{
	// 			ID: to.Ptr("tableName"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1ad29756bd141a47cac770140105a706d065ae1b/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBTableDelete.json
func ExampleTableResourcesClient_BeginDeleteTable() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTableResourcesClient().BeginDeleteTable(ctx, "rg1", "ddb1", "tableName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1ad29756bd141a47cac770140105a706d065ae1b/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBTableThroughputGet.json
func ExampleTableResourcesClient_GetTableThroughput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTableResourcesClient().GetTableThroughput(ctx, "rg1", "ddb1", "tableName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ThroughputSettingsGetResults = armcosmos.ThroughputSettingsGetResults{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/tables/throughputSettings"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/tables/tableName/throughputSettings/default"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 		Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			InstantMaximumThroughput: to.Ptr("10000"),
	// 			MinimumThroughput: to.Ptr("400"),
	// 			OfferReplacePending: to.Ptr("true"),
	// 			SoftAllowedMaximumThroughput: to.Ptr("1000000"),
	// 			Throughput: to.Ptr[int32](400),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1ad29756bd141a47cac770140105a706d065ae1b/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBTableThroughputUpdate.json
func ExampleTableResourcesClient_BeginUpdateTableThroughput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTableResourcesClient().BeginUpdateTableThroughput(ctx, "rg1", "ddb1", "tableName", armcosmos.ThroughputSettingsUpdateParameters{
		Location: to.Ptr("West US"),
		Tags:     map[string]*string{},
		Properties: &armcosmos.ThroughputSettingsUpdateProperties{
			Resource: &armcosmos.ThroughputSettingsResource{
				Throughput: to.Ptr[int32](400),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ThroughputSettingsGetResults = armcosmos.ThroughputSettingsGetResults{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/tables/throughputSettings"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/tables/tableName/throughputSettings/default"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 		Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			MinimumThroughput: to.Ptr("400"),
	// 			OfferReplacePending: to.Ptr("true"),
	// 			Throughput: to.Ptr[int32](400),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1ad29756bd141a47cac770140105a706d065ae1b/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBTableMigrateToAutoscale.json
func ExampleTableResourcesClient_BeginMigrateTableToAutoscale() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTableResourcesClient().BeginMigrateTableToAutoscale(ctx, "rg1", "ddb1", "tableName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ThroughputSettingsGetResults = armcosmos.ThroughputSettingsGetResults{
	// 	Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 		Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			AutoscaleSettings: &armcosmos.AutoscaleSettingsResource{
	// 				MaxThroughput: to.Ptr[int32](4000),
	// 			},
	// 			MinimumThroughput: to.Ptr("4000"),
	// 			OfferReplacePending: to.Ptr("false"),
	// 			Throughput: to.Ptr[int32](400),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1ad29756bd141a47cac770140105a706d065ae1b/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBTableMigrateToManualThroughput.json
func ExampleTableResourcesClient_BeginMigrateTableToManualThroughput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTableResourcesClient().BeginMigrateTableToManualThroughput(ctx, "rg1", "ddb1", "tableName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ThroughputSettingsGetResults = armcosmos.ThroughputSettingsGetResults{
	// 	Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 		Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 			Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			Rid: to.Ptr("PD5DALigDgw="),
	// 			Ts: to.Ptr[float32](1459200611),
	// 			MinimumThroughput: to.Ptr("400"),
	// 			OfferReplacePending: to.Ptr("true"),
	// 			Throughput: to.Ptr[int32](400),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1ad29756bd141a47cac770140105a706d065ae1b/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBTableBackupInformation.json
func ExampleTableResourcesClient_BeginRetrieveContinuousBackupInformation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTableResourcesClient().BeginRetrieveContinuousBackupInformation(ctx, "rgName", "ddb1", "tableName1", armcosmos.ContinuousBackupRestoreLocation{
		Location: to.Ptr("North Europe"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BackupInformation = armcosmos.BackupInformation{
	// 	ContinuousBackupInformation: &armcosmos.ContinuousBackupInformation{
	// 		LatestRestorableTimestamp: to.Ptr("2021-02-05T02:40:50Z"),
	// 	},
	// }
}
