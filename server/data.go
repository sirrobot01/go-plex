package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirrobot01/goplex_debrid/config"
	"github.com/sirrobot01/goplex_debrid/schemas"
)

func GetProviderData(server *schemas.Server) fiber.Map {
	data := fiber.Map{
		"MediaContainer": fiber.Map{
			"size":                          1,
			"allowCameraUpload":             true,
			"allowChannelAccess":            true,
			"allowMediaDeletion":            true,
			"allowSharing":                  true,
			"allowSync":                     true,
			"allowTuners":                   true,
			"backgroundProcessing":          true,
			"certificate":                   true,
			"companionProxy":                true,
			"countryCode":                   "deu",
			"diagnostics":                   "logs,databases,streaminglogs",
			"eventStream":                   true,
			"friendlyName":                  server.Name,
			"livetv":                        7,
			"machineIdentifier":             server.Hash,
			"myPlex":                        true,
			"myPlexMappingState":            "mapped",
			"myPlexSigninState":             "ok",
			"myPlexSubscription":            true,
			"myPlexUsername":                config.CONFIG.PlexEmail,
			"offlineTranscode":              1,
			"ownerFeatures":                 "002c9f1a-2fc0-4812-b85b-0e6140f21a0f,00cc618e-eb08-4e0e-9221-82b4835dd89b,044a1fac-6b55-47d0-9933-25a035709432,04d7d794-b76c-49ef-9184-52f8f1f501ee,05690239-443e-43fb-bc1a-95b5d916ca63,05826f20-284b-4bcb-b45c-2367e5c0ea72,06d14b9e-2af8-4c2b-a4a1-ea9d5c515824,075954ad-56ef-4f5e-9519-9cfb0ed05827,07f804e6-28e6-4beb-b5c3-f2aefc88b938,0de6151c-e0dd-47c8-a81e-1acb977c7f0f,0eee866d-782b-4dfd-b42b-3bbe8eb0af16,13056a62-9bd2-47cf-aba9-bab00095fd08,1417df52-986e-4e4b-8dcd-3997fbc5c976,1844737f-1a87-45c3-ab20-01435959e63c,1b3a63e4-c2f4-4011-a181-2343d3a97ef7,1b870b8e-f1a7-497c-80b2-857d45f3123f,1dd846ed-7cde-4dc5-8ef6-53d3ce8c4e9d,1df3cd16-faf2-4d37-8349-1fcf3713bf1d,1f952ea5-0837-44cb-8539-a69a14a75d4a,222020fb-1504-492d-af33-a0b80a49558a,228a6439-ee2f-4a9b-b0fc-1bfcd48b5095,22b27e12-472e-4383-92ea-2ec3976d8e72,24b4cf36-b296-4002-86b7-f1adb657e76a,2797e341-b062-46ed-862f-0acbba5dd522,298a11d3-9324-4104-8047-0ac10df4a8a6,2ea0e464-ea4f-4be2-97c1-ce6ed4b377dd,300231e0-69aa-4dce-97f4-52d8c00e3e8c,32cc8bf5-b425-4582-a52d-71b4f1cf436b,34e182bd-2f62-4678-a9e9-d13b3e25019d,39dbdd84-8339-4736-96a1-0eb105cc2e08,3a2b0cb6-1519-4431-98e2-823c248c70eb,3ae06d3a-a76b-435e-8cef-2d2008610ba2,3bfd3ccf-8c63-4dbb-8f87-9b21b402c82b,3c376154-d47e-4bbf-9428-2ea2592fd20a,4742780c-af9d-4b44-bf5b-7b27e3369aa8,4b522f91-ae89-4f62-af9c-76f44d8ef61c,4ca03b04-54c1-4f9f-aea2-f813ae48f317,4cd4dc0e-6cbe-456c-9988-9f073fadcd73,547514ab-3284-46e5-af77-bbaff247e3fc,55b9f6ed-5d26-4d2d-a436-68882a9901b5,567033ef-ffee-44fb-8f90-f678077445f9,5b6190a9-77a4-477e-9fbc-c8118e35a4c1,5c1951bf-ccf1-4821-8ee7-e50f51218ae7,5d819d02-5d04-4116-8eec-f49def4e2d6f,5e2a89ec-fb26-4234-b66e-14d37f35dff2,62b1e357-5450-41d8-9b60-c7705f750849,6380e085-02fe-43b5-8bff-380fa4f2423c,644c4466-05fa-45e0-a478-c594cf81778f,65152b75-13a9-408a-bd30-dbd23a259183,65685ff8-4375-4e4c-a806-ec1f0b4a8b7f,67c80530-eae3-4500-a9fa-9b6947d0f6d1,68747f3a-ce13-46ce-9274-1e0544c9f500,6b85840c-d79d-40c2-8d8f-dfc0b7d26776,6d7be725-9a96-42c7-8af4-01e735138822,6f82ca43-6117-4e55-ae0e-5ea3b3e99a96,7b392594-6949-4736-9894-e57a9dfe4037,7ee1495c-2798-4288-94e2-9cd98e67d441,81c8d5fa-8d90-4833-aa10-a31a51310e2f,82999dd3-a2be-482e-9f44-357879b4f603,849433b0-ef60-4a71-9dd9-939bc01f5362,84a754b0-d1ca-4433-af2d-c949bf4b4936,850f3d1e-3f38-44c1-9c0c-e3c9127b8b5a,8536058d-e1dd-4ae7-b30f-e8b059b7cc17,85ebfb7b-77fb-4afd-bb1a-2fe2fefdddbe,86da2200-58db-4d78-ba46-f146ba25906b,88aba3a3-bd62-42a5-91bb-0558a4c1db57,8e8dd5c8-14a4-4208-97d4-623e09191774,8fd37970-6e4e-4f00-a64a-e70b52f18e94,926bc176-58ca-47da-b8e3-080ed14ea6ba,93bf35b9-3b62-4a8a-b09b-5c85437fa67b,95149521-f64b-46ea-825c-9114e56afd2c,96cac76e-c5bc-4596-87eb-4fdfef9aaa11,9a67bff2-cb80-4bf9-81c6-9ad2f4c78afd,9c982beb-c676-4d6f-a777-ff5d37ec3081,9dc1df45-fb45-4be1-9ab2-eb23eb57f082,9e93f8a8-7ccd-4d15-99fa-76a158027660,a19d495a-1cef-4f7c-ab77-5186e63e17f7,a3d2d5c4-46a0-436e-a2d6-80d26f32b369,a536a6e1-0ece-498a-bf64-99b53c27de3a,a548af72-b804-4d05-8569-52785952d31d,a6e0a154-4735-4cbb-a6ec-7a0a146c8216,a6f3f9b3-c10c-4b94-ad59-755e30ac6c90,abd37b14-706c-461f-8255-fa9563882af3,adaptive_bitrate,af866200-1116-4502-be15-f71457ce9d27,b227c158-e062-4ff1-95d8-8ed11cecafb1,b2403ac6-4885-4971-8b96-59353fd87c72,b3b87f19-5ccd-4b14-bb62-b9d7b982392e,b46d16ae-cbd6-4226-8ee9-ab2b27e5dd42,b5874ecb-6610-47b2-8906-1b5a897acb02,b58d7f28-7b4a-49bb-97a7-152645505f28,b612f571-83c3-431a-88eb-3f05ce08da4a,b77e6744-c18d-415a-8e7c-7aac5d7a7750,b83c8dc9-5a01-4b7a-a7c9-5870c8a6e21b,b8cf9f40-4f8a-4de4-b203-5bbcf8b09f5a,bb50c92f-b412-44fe-8d8a-b1684f212a44,bbf73498-4912-4d80-9560-47c4fe212cec,bc8d1fca-deb0-4d0a-a6f4-12cfd681002d,bec2ba97-4b25-472b-9cfc-674f5c68c2ae,c2409baa-d044-45c7-b1f4-e9e7ccd2d128,c36a6985-eee3-4400-a394-c5787fad15b5,c55d5900-b546-416d-a8c5-45b24a13e9bc,c5adf9dc-af13-4a85-a24b-98de6fa2f595,c7ae6f8f-05e6-48bb-9024-c05c1dc3c43e,c92d4903-bc06-4715-8ce4-4a22674abac8,c987122a-a796-432f-af00-953821c127bb,c9d9b7ee-fdd9-474e-b143-5039c04e9b9b,camera_upload,cb151c05-1943-408a-b37c-06f7d409d6bb,cbae4949-1643-46f3-a488-71836b025d63,cc9bea3b-11ab-4402-a222-4958bb129cab,ccef9d3a-537a-43d9-8161-4c7113c6e2bb,ce8f644e-87ce-4ba5-b165-fadd69778019,collections,content_filter,d14556be-ae6d-4407-89d0-b83953f4789a,d1477307-4dac-4e57-9258-252e5b908693,d20f9af2-fdb1-4927-99eb-a2eb8fbff799,d413fb56-de7b-40e4-acd0-f3dbb7c9e104,d85cb60c-0986-4a02-b1e1-36c64c609712,d8810b38-ec9b-494c-8555-3df6e365dfbd,d9f42aea-bc9d-47db-9814-cd7a577aff48,dab501df-5d99-48ef-afc2-3e839e4ddc9a,db965785-ca5c-46fd-bab6-7b3d29c18492,dbdc0575-9fc7-4706-9e2d-fca98d10ad71,download_certificates,dvr,e45bc5ae-1c3a-4729-922b-c69388c571b7,e66aa31c-abdd-483d-93bc-e17485d8837f,e703655b-ee05-4e24-97e3-a138da62c425,e8230c74-0940-4b91-9e20-6571eb068086,e954ef21-08b4-411e-a1f0-7551f1e57b11,ea442c16-044a-4fa7-8461-62643f313c62,ec64b6f6-e804-4ef3-b114-9d5c63e1a941,ee352392-2934-4061-ba35-5f3189f19ab4,f0f40559-a43a-4b8f-85ef-bdb1de1a912a,f1ac7a53-c524-4311-9a27-713562fc24fa,f3235e61-c0eb-4718-ac0a-7d6eb3d8ff75,f3a99481-9671-4274-a0d3-4c06a72ef746,f83450e2-759a-4de4-8b31-e4a163896d43,f8463032-28f1-447b-a76c-8b57a071acad,f87f382b-4a41-4951-b4e4-d5822c69e4c6,f8ea4f37-c554-476a-8852-1cbd2912f3f6,f8f68869-07a5-4fc3-9359-d0b5ba9c487c,fb34e64d-cd89-47b8-8bae-a6d20c542bae,fd6683b9-1426-4b00-840f-cd5fb0904a6a,fec722a0-a6d4-4fbd-96dc-4ffb02b072c5,federated-auth,hardware_transcoding,home,hwtranscode,item_clusters,kevin-bacon,livetv,loudness,lyrics,music_videos,pass,photosV6-edit,photosV6-tv-albums,premium_music_metadata,radio,server-manager,session_bandwidth_restrictions,session_kick,shared-radio,sync,trailers,tuner-sharing,type-first,unsupportedtuners,webhooks",
			"photoAutoTag":                  true,
			"platform":                      "Linux",
			"platformVersion":               "22.04.2 LTS (Jammy Jellyfish)",
			"pluginHost":                    true,
			"pushNotifications":             false,
			"readOnlyLibraries":             false,
			"streamingBrainABRVersion":      3,
			"streamingBrainVersion":         2,
			"sync":                          true,
			"transcoderActiveVideoSessions": 0,
			"transcoderAudio":               true,
			"transcoderLyrics":              true,
			"transcoderSubtitles":           true,
			"transcoderVideo":               true,
			"transcoderVideoBitrates":       "64,96,208,320,720,1500,2000,3000,4000,8000,10000,12000,20000",
			"transcoderVideoQualities":      "0,1,2,3,4,5,6,7,8,9,10,11,12",
			"transcoderVideoResolutions":    "128,128,160,240,320,480,768,720,720,1080,1080,1080,1080",
			"updatedAt":                     1693327244,
			"updater":                       true,
			"version":                       "1.32.5.7349-8f4248874",
			"voiceSearch":                   true,
			"MediaProvider": []fiber.Map{
				{
					"identifier": "com.plexapp.plugins.library",
					"title":      "Library",
					"types":      "video,audio,photo",
					"protocols":  "stream,download",
					"Feature": []fiber.Map{
						{
							"key":  "/library/sections",
							"type": "content",
							"Directory": []fiber.Map{
								{
									"hubKey": "/hubs",
									"title":  "Home",
								},
								{
									"agent":      "tv.plex.agents.movie",
									"language":   "en-US",
									"refreshing": false,
									"scanner":    "Plex Movie",
									"uuid":       "016f2df0-980c-4d70-84b6-843163c85b2c",
									"id":         "1",
									"key":        "/library/sections/1",
									"hubKey":     "/hubs/sections/1",
									"type":       "movie",
									"title":      "Movies",
									"updatedAt":  1667326439,
									"scannedAt":  1698686166,
									"Pivot": []fiber.Map{
										{
											"id":      "recommended",
											"key":     "/hubs/sections/1",
											"type":    "hub",
											"title":   "Recommended",
											"context": "content.discover",
											"symbol":  "star",
										},
										{
											"id":      "library",
											"key":     "/library/sections/1/all?type=1",
											"type":    "list",
											"title":   "Library",
											"context": "content.library",
											"symbol":  "library",
										},
										{
											"id":      "categories",
											"key":     "/library/sections/1/categories",
											"type":    "list",
											"title":   "Categories",
											"context": "content.categories",
											"symbol":  "stack",
										},
									},
								},
								{
									"agent":      "tv.plex.agents.series",
									"language":   "en-US",
									"refreshing": false,
									"scanner":    "Plex TV Series",
									"uuid":       "a4215914-9cf7-4092-b08f-6705bb70cc5c",
									"id":         "5",
									"key":        "/library/sections/5",
									"hubKey":     "/hubs/sections/5",
									"type":       "show",
									"title":      "TV Shows",
									"updatedAt":  1689872928,
									"scannedAt":  1698688886,
									"Pivot": []fiber.Map{
										{
											"id":      "recommended",
											"key":     "/hubs/sections/5",
											"type":    "hub",
											"title":   "Recommended",
											"context": "content.discover",
											"symbol":  "star",
										},
										{
											"id":      "library",
											"key":     "/library/sections/5/all?type=2",
											"type":    "list",
											"title":   "Library",
											"context": "content.library",
											"symbol":  "library",
										},
										{
											"id":      "categories",
											"key":     "/library/sections/5/categories",
											"type":    "list",
											"title":   "Categories",
											"context": "content.categories",
											"symbol":  "stack",
										},
									},
								},
							},
						},
						{
							"key":  "/hubs/search",
							"type": "search",
						},
						{
							"key":  "/library/matches",
							"type": "match",
						},
						{
							"key":  "/library/metadata",
							"type": "metadata",
						},
						{
							"key":  "/:/rate",
							"type": "rate",
						},
						{
							"key":  "/photo/:/transcode",
							"type": "imagetranscoder",
						},
						{
							"key":  "/hubs/promoted",
							"type": "promoted",
						},
						{
							"key":  "/hubs/continueWatching",
							"type": "continuewatching",
						},
						{
							"key":  "/actions",
							"type": "actions",
							"Action": []fiber.Map{
								{
									"id":  "removeFromContinueWatching",
									"key": "/actions/removeFromContinueWatching",
								},
							},
						},
						{
							"flavor": "universal",
							"key":    "/playlists",
							"type":   "playlist",
						},
						{
							"flavor": "universal",
							"key":    "/playQueues",
							"type":   "playqueue",
						},
						{
							"key":  "/library/collections",
							"type": "collection",
						},
						{
							"scrobbleKey":   "/:/scrobble",
							"unscrobbleKey": "/:/unscrobble",
							"key":           "/:/timeline",
							"type":          "timeline",
						},
						{
							"type": "manage",
						},
						{
							"type": "queryParser",
						},
						{
							"flavor": "download",
							"type":   "subscribe",
						},
					},
				},
			},
		},
	}
	return data
}

func GetMobileProviderData(server *schemas.Server) string {
	str := `<?xml version="1.0" encoding="UTF-8"?>\n<MediaContainer size="1" allowCameraUpload="1" allowChannelAccess="1" allowMediaDeletion="1" allowSharing="1" allowSync="1" allowTuners="1" backgroundProcessing="1" certificate="1" companionProxy="1" countryCode="deu" diagnostics="logs,databases,streaminglogs" eventStream="1" friendlyName="%s" livetv="7" machineIdentifier="%s" musicAnalysis="2" myPlex="1" myPlexMappingState="mapped" myPlexSigninState="ok" myPlexSubscription="1" myPlexUsername="%s" offlineTranscode="1" ownerFeatures="002c9f1a-2fc0-4812-b85b-0e6140f21a0f,044a1fac-6b55-47d0-9933-25a035709432,04d7d794-b76c-49ef-9184-52f8f1f501ee,05690239-443e-43fb-bc1a-95b5d916ca63,06d14b9e-2af8-4c2b-a4a1-ea9d5c515824,07f804e6-28e6-4beb-b5c3-f2aefc88b938,0a348865-4f87-46dc-8bb2-f37637975724,0de6151c-e0dd-47c8-a81e-1acb977c7f0f,0eee866d-782b-4dfd-b42b-3bbe8eb0af16,13056a62-9bd2-47cf-aba9-bab00095fd08,1417df52-986e-4e4b-8dcd-3997fbc5c976,16d69c53-4c40-4821-b9f3-57ca690b2d4d,1844737f-1a87-45c3-ab20-01435959e63c,1b3a63e4-c2f4-4011-a181-2343d3a97ef7,1dd846ed-7cde-4dc5-8ef6-53d3ce8c4e9d,1df3cd16-faf2-4d37-8349-1fcf3713bf1d,222020fb-1504-492d-af33-a0b80a49558a,228a6439-ee2f-4a9b-b0fc-1bfcd48b5095,22b27e12-472e-4383-92ea-2ec3976d8e72,22d52c96-9e2b-45c0-9e2a-1d6c66ad3474,24b4cf36-b296-4002-86b7-f1adb657e76a,2797e341-b062-46ed-862f-0acbba5dd522,298a11d3-9324-4104-8047-0ac10df4a8a6,2ea0e464-ea4f-4be2-97c1-ce6ed4b377dd,300231e0-69aa-4dce-97f4-52d8c00e3e8c,32cc8bf5-b425-4582-a52d-71b4f1cf436b,34e182bd-2f62-4678-a9e9-d13b3e25019d,39dbdd84-8339-4736-96a1-0eb105cc2e08,3a2b0cb6-1519-4431-98e2-823c248c70eb,3ae06d3a-a76b-435e-8cef-2d2008610ba2,3bfd3ccf-8c63-4dbb-8f87-9b21b402c82b,3c376154-d47e-4bbf-9428-2ea2592fd20a,4742780c-af9d-4b44-bf5b-7b27e3369aa8,4b522f91-ae89-4f62-af9c-76f44d8ef61c,4ca03b04-54c1-4f9f-aea2-f813ae48f317,4cd4dc0e-6cbe-456c-9988-9f073fadcd73,547514ab-3284-46e5-af77-bbaff247e3fc,55b9f6ed-5d26-4d2d-a436-68882a9901b5,567033ef-ffee-44fb-8f90-f678077445f9,5b6190a9-77a4-477e-9fbc-c8118e35a4c1,5c1951bf-ccf1-4821-8ee7-e50f51218ae7,5d819d02-5d04-4116-8eec-f49def4e2d6f,5e2a89ec-fb26-4234-b66e-14d37f35dff2,62b1e357-5450-41d8-9b60-c7705f750849,6380e085-02fe-43b5-8bff-380fa4f2423c,644c4466-05fa-45e0-a478-c594cf81778f,65152b75-13a9-408a-bd30-dbd23a259183,65685ff8-4375-4e4c-a806-ec1f0b4a8b7f,67c80530-eae3-4500-a9fa-9b6947d0f6d1,68747f3a-ce13-46ce-9274-1e0544c9f500,6b85840c-d79d-40c2-8d8f-dfc0b7d26776,6d7be725-9a96-42c7-8af4-01e735138822,6f82ca43-6117-4e55-ae0e-5ea3b3e99a96,78643fe5-d192-40c7-8e93-5ccf04c0b767,7e7596aa-6e2c-41d1-a460-1e13cf0b62f2,7ee1495c-2798-4288-94e2-9cd98e67d441,82999dd3-a2be-482e-9f44-357879b4f603,849433b0-ef60-4a71-9dd9-939bc01f5362,84a754b0-d1ca-4433-af2d-c949bf4b4936,850f3d1e-3f38-44c1-9c0c-e3c9127b8b5a,8536058d-e1dd-4ae7-b30f-e8b059b7cc17,85ebfb7b-77fb-4afd-bb1a-2fe2fefdddbe,86da2200-58db-4d78-ba46-f146ba25906b,88aba3a3-bd62-42a5-91bb-0558a4c1db57,8e8dd5c8-14a4-4208-97d4-623e09191774,8fd37970-6e4e-4f00-a64a-e70b52f18e94,95149521-f64b-46ea-825c-9114e56afd2c,96cac76e-c5bc-4596-87eb-4fdfef9aaa11,98872b06-2ff3-4b71-96bc-039e2ebe7adc,9a67bff2-cb80-4bf9-81c6-9ad2f4c78afd,9c982beb-c676-4d6f-a777-ff5d37ec3081,9dc1df45-fb45-4be1-9ab2-eb23eb57f082,a19d495a-1cef-4f7c-ab77-5186e63e17f7,a3d2d5c4-46a0-436e-a2d6-80d26f32b369,a536a6e1-0ece-498a-bf64-99b53c27de3a,a548af72-b804-4d05-8569-52785952d31d,a6e0a154-4735-4cbb-a6ec-7a0a146c8216,a6f3f9b3-c10c-4b94-ad59-755e30ac6c90,abd37b14-706c-461f-8255-fa9563882af3,adaptive_bitrate,b20d91ca-1b2f-45a2-a115-c1ad24c66ac5,b227c158-e062-4ff1-95d8-8ed11cecafb1,b2403ac6-4885-4971-8b96-59353fd87c72,b46d16ae-cbd6-4226-8ee9-ab2b27e5dd42,b5874ecb-6610-47b2-8906-1b5a897acb02,b58d7f28-7b4a-49bb-97a7-152645505f28,b612f571-83c3-431a-88eb-3f05ce08da4a,b77e6744-c18d-415a-8e7c-7aac5d7a7750,b83c8dc9-5a01-4b7a-a7c9-5870c8a6e21b,b8cf9f40-4f8a-4de4-b203-5bbcf8b09f5a,bb50c92f-b412-44fe-8d8a-b1684f212a44,bbf73498-4912-4d80-9560-47c4fe212cec,bc8d1fca-deb0-4d0a-a6f4-12cfd681002d,bfeaee4e-965a-4d24-b163-020c3c57d936,c2409baa-d044-45c7-b1f4-e9e7ccd2d128,c55d5900-b546-416d-a8c5-45b24a13e9bc,c5adf9dc-af13-4a85-a24b-98de6fa2f595,c7ae6f8f-05e6-48bb-9024-c05c1dc3c43e,c92d4903-bc06-4715-8ce4-4a22674abac8,c9d9b7ee-fdd9-474e-b143-5039c04e9b9b,camera_upload,cc9bea3b-11ab-4402-a222-4958bb129cab,cloudsync,collections,content_filter,d14556be-ae6d-4407-89d0-b83953f4789a,d1477307-4dac-4e57-9258-252e5b908693,d20f9af2-fdb1-4927-99eb-a2eb8fbff799,d413fb56-de7b-40e4-acd0-f3dbb7c9e104,d85cb60c-0986-4a02-b1e1-36c64c609712,d8810b38-ec9b-494c-8555-3df6e365dfbd,d9f42aea-bc9d-47db-9814-cd7a577aff48,dab501df-5d99-48ef-afc2-3e839e4ddc9a,db965785-ca5c-46fd-bab6-7b3d29c18492,download_certificates,dvr,e45bc5ae-1c3a-4729-922b-c69388c571b7,e66aa31c-abdd-483d-93bc-e17485d8837f,e8230c74-0940-4b91-9e20-6571eb068086,e954ef21-08b4-411e-a1f0-7551f1e57b11,ea442c16-044a-4fa7-8461-62643f313c62,ec64b6f6-e804-4ef3-b114-9d5c63e1a941,ee352392-2934-4061-ba35-5f3189f19ab4,f0f40559-a43a-4b8f-85ef-bdb1de1a912a,f3235e61-c0eb-4718-ac0a-7d6eb3d8ff75,f3a99481-9671-4274-a0d3-4c06a72ef746,f83450e2-759a-4de4-8b31-e4a163896d43,f87f382b-4a41-4951-b4e4-d5822c69e4c6,f8ea4f37-c554-476a-8852-1cbd2912f3f6,fb34e64d-cd89-47b8-8bae-a6d20c542bae,fd6683b9-1426-4b00-840f-cd5fb0904a6a,fec722a0-a6d4-4fbd-96dc-4ffb02b072c5,federated-auth,hardware_transcoding,home,hwtranscode,item_clusters,kevin-bacon,livetv,loudness,lyrics,music-analysis,music_videos,pass,photo_autotags,photos-v5,photosV6-edit,photosV6-tv-albums,premium_music_metadata,radio,server-manager,session_bandwidth_restrictions,session_kick,shared-radio,sync,trailers,tuner-sharing,type-first,unsupportedtuners,webhooks" photoAutoTag="1" platform="Windows" platformVersion="10.0 (Build 19042)" pluginHost="1" pushNotifications="0" readOnlyLibraries="0" streamingBrainABRVersion="3" streamingBrainVersion="2" sync="1" transcoderActiveVideoSessions="0" transcoderAudio="1" transcoderLyrics="1" transcoderSubtitles="1" transcoderVideo="1" transcoderVideoBitrates="64,96,208,320,720,1500,2000,3000,4000,8000,10000,12000,20000" transcoderVideoQualities="0,1,2,3,4,5,6,7,8,9,10,11,12" transcoderVideoResolutions="128,128,160,240,320,480,768,720,720,1080,1080,1080,1080" updatedAt="1655653966" updater="1" version="1.27.0.5897-3940636f2" voiceSearch="1">\n<MediaProvider identifier="com.plexapp.plugins.library" title="Library" types="video,audio,photo" protocols="stream,download">\n<Feature key="/library/sections" type="content">\n<Directory hubKey="/hubs" title="Home">\n</Directory>\n<Directory agent="tv.plex.agents.movie" language="en-US" refreshing="0" scanner="Plex Movie" uuid="4f860389-974a-488a-9f0e-70c8afc7d8e4" id="10" key="/library/sections/10" hubKey="/hubs/sections/10" type="movie" title="Movies" updatedAt="1655532971" scannedAt="1655532971">\n<Preferences>\n<Setting id="hidden" label="Visibility" summary="Restrict where content from this library should appear." type="int" default="0" value="0" hidden="0" advanced="0" group="" enumValues="0:Include in home screen and global search|1:Exclude from home screen|2:Exclude from home screen and global search" />\n<Setting id="enableCinemaTrailers" label="Enable Cinema Trailers" summary="Allow Cinema Trailers to play before items in this library." type="bool" default="True" value="True" hidden="0" advanced="0" group="" />\n<Setting id="country" label="Certification Country" summary="This will influence which content rating system is used. Changing this setting will require refreshing the metadata for the changes to take effect." type="text" default="US" value="US" hidden="0" advanced="0" group="" enumValues="AF:Afghanistan|AX:&#197;land Islands|AL:Albania|DZ:Algeria|AS:American Samoa|AD:Andorra|AO:Angola|AI:Anguilla|AQ:Antarctica|AG:Antigua &amp; Barbuda|AR:Argentina|AM:Armenia|AW:Aruba|AU:Australia|AT:Austria|AZ:Azerbaijan|BS:Bahamas|BH:Bahrain|BD:Bangladesh|BB:Barbados|BY:Belarus|BE:Belgium|BZ:Belize|BJ:Benin|BM:Bermuda|BT:Bhutan|BO:Bolivia|BA:Bosnia &amp; Herzegovina|BW:Botswana|BV:Bouvet Island|BR:Brazil|IO:British Indian Ocean Territory|VG:British Virgin Islands|BN:Brunei|BG:Bulgaria|BF:Burkina Faso|BI:Burundi|KH:Cambodia|CM:Cameroon|CA:Canada|IC:Canary Islands|CV:Cape Verde|BQ:Caribbean Netherlands|KY:Cayman Islands|CF:Central African Republic|EA:Ceuta &amp; Melilla|TD:Chad|CL:Chile|CN:China|CX:Christmas Island|CC:Cocos (Keeling) Islands|CO:Colombia|KM:Comoros|CG:Congo - Brazzaville|CD:Congo - Kinshasa|CK:Cook Islands|CR:Costa Rica|CI:C&#244;te d&#8217;Ivoire|HR:Croatia|CU:Cuba|CW:Cura&#231;ao|CY:Cyprus|CZ:Czechia|DK:Denmark|DG:Diego Garcia|DJ:Djibouti|DM:Dominica|DO:Dominican Republic|EC:Ecuador|EG:Egypt|SV:El Salvador|GQ:Equatorial Guinea|ER:Eritrea|EE:Estonia|SZ:Eswatini|ET:Ethiopia|FK:Falkland Islands|FO:Faroe Islands|FJ:Fiji|FI:Finland|FR:France|GF:French Guiana|PF:French Polynesia|TF:French Southern Territories|GA:Gabon|GM:Gambia|GE:Georgia|DE:Germany|GH:Ghana|GI:Gibraltar|GR:Greece|GL:Greenland|GD:Grenada|GP:Guadeloupe|GU:Guam|GT:Guatemala|GG:Guernsey|GN:Guinea|GW:Guinea-Bissau|GY:Guyana|HT:Haiti|HM:Heard &amp; McDonald Islands|HN:Honduras|HK:Hong Kong SAR China|HU:Hungary|IS:Iceland|IN:India|ID:Indonesia|IR:Iran|IQ:Iraq|IE:Ireland|IM:Isle of Man|IL:Israel|IT:Italy|JM:Jamaica|JP:Japan|JE:Jersey|JO:Jordan|KZ:Kazakhstan|KE:Kenya|KI:Kiribati|XK:Kosovo|KW:Kuwait|KG:Kyrgyzstan|LA:Laos|LV:Latvia|LB:Lebanon|LS:Lesotho|LR:Liberia|LY:Libya|LI:Liechtenstein|LT:Lithuania|LU:Luxembourg|MO:Macao SAR China|MG:Madagascar|MW:Malawi|MY:Malaysia|MV:Maldives|ML:Mali|MT:Malta|MH:Marshall Islands|MQ:Martinique|MR:Mauritania|MU:Mauritius|YT:Mayotte|MX:Mexico|FM:Micronesia|MD:Moldova|MC:Monaco|MN:Mongolia|ME:Montenegro|MS:Montserrat|MA:Morocco|MZ:Mozambique|MM:Myanmar (Burma)|NA:Namibia|NR:Nauru|NP:Nepal|NL:Netherlands|NC:New Caledonia|NZ:New Zealand|NI:Nicaragua|NE:Niger|NG:Nigeria|NU:Niue|NF:Norfolk Island|KP:North Korea|MK:North Macedonia|MP:Northern Mariana Islands|NO:Norway|OM:Oman|PK:Pakistan|PW:Palau|PS:Palestinian Territories|PA:Panama|PG:Papua New Guinea|PY:Paraguay|PE:Peru|PH:Philippines|PN:Pitcairn Islands|PL:Poland|PT:Portugal|PR:Puerto Rico|QA:Qatar|RE:R&#233;union|RO:Romania|RU:Russia|RW:Rwanda|WS:Samoa|SM:San Marino|ST:S&#227;o Tom&#233; &amp; Pr&#237;ncipe|SA:Saudi Arabia|SN:Senegal|RS:Serbia|SC:Seychelles|SL:Sierra Leone|SG:Singapore|SX:Sint Maarten|SK:Slovakia|SI:Slovenia|SB:Solomon Islands|SO:Somalia|ZA:South Africa|GS:South Georgia &amp; South Sandwich Islands|KR:South Korea|SS:South Sudan|ES:Spain|LK:Sri Lanka|BL:St. Barth&#233;lemy|SH:St. Helena|KN:St. Kitts &amp; Nevis|LC:St. Lucia|MF:St. Martin|PM:St. Pierre &amp; Miquelon|VC:St. Vincent &amp; Grenadines|SD:Sudan|SR:Suriname|SJ:Svalbard &amp; Jan Mayen|SE:Sweden|CH:Switzerland|SY:Syria|TW:Taiwan|TJ:Tajikistan|TZ:Tanzania|TH:Thailand|TL:Timor-Leste|TG:Togo|TK:Tokelau|TO:Tonga|TT:Trinidad &amp; Tobago|TN:Tunisia|TR:Turkey|TM:Turkmenistan|TC:Turks &amp; Caicos Islands|TV:Tuvalu|UG:Uganda|UA:Ukraine|AE:United Arab Emirates|GB:United Kingdom|US:United States|UY:Uruguay|UM:US Outlying Islands|VI:US Virgin Islands|UZ:Uzbekistan|VU:Vanuatu|VA:Vatican City|VE:Venezuela|VN:Vietnam|WF:Wallis &amp; Futuna|EH:Western Sahara|YE:Yemen|ZM:Zambia|ZW:Zimbabwe" />\n<Setting id="originalTitles" label="Use original titles" summary="Use the original titles for all items regardless of the library language." type="bool" default="False" value="False" hidden="0" advanced="0" group="" />\n<Setting id="localizedArtwork" label="Prefer artwork based on library language" summary="Use localized posters when available. This is determined by the library language setting." type="bool" default="True" value="True" hidden="0" advanced="0" group="" />\n<Setting id="useLocalAssets" label="Use local assets" summary="When scanning this library, use local posters and artwork if present." type="bool" default="True" value="True" hidden="0" advanced="0" group="" />\n<Setting id="respectTags" label="Prefer local metadata" summary="When scanning this library, prefer embedded tags and local files if present." type="bool" default="False" value="False" hidden="0" advanced="0" group="" />\n<Setting id="useExternalExtras" label="Find extras" summary="Find trailers and extras automatically (Plex Pass required)" type="bool" default="True" value="True" hidden="0" advanced="0" group="" />\n<Setting id="skipNonTrailerExtras" label="Only show trailers" summary="Skip extras which aren&#39;t trailers" type="bool" default="False" value="False" hidden="0" advanced="0" group="" />\n<Setting id="useRedbandTrailers" label="Allow red band trailers" summary="Use red band (restricted audiences) trailers when available" type="bool" default="False" value="False" hidden="0" advanced="0" group="" />\n<Setting id="includeExtrasWithLocalizedSubtitles" label="Localized subtitles" summary="Include extras with subtitles in Library language" type="bool" default="False" value="False" hidden="0" advanced="0" group="" />\n<Setting id="includeAdultContent" label="Include adult content" summary="Allow matching adult titles and returning adult titles in fix match results." type="bool" default="False" value="False" hidden="0" advanced="0" group="" />\n<Setting id="autoCollectionThreshold" label="Minimum automatic collection size" summary="Automatically create collections when there are at least the selected number of items for an available collection. Changing this value will have no effect on any existing collections." type="int" default="0" value="0" hidden="0" advanced="0" group="" enumValues="0:Disabled|1:1|2:2|3:3|4:4" />\n<Setting id="ratingsSource" label="Ratings Source" summary="Select a primary source for ratings." type="text" default="rottentomatoes" value="rottentomatoes" hidden="0" advanced="0" group="" enumValues="rottentomatoes:Rotten Tomatoes|imdb:IMDb|themoviedb:The Movie Database" />\n<Setting id="enableBIFGeneration" label="Enable video preview thumbnails" summary="Generate video preview thumbnails for items in this library when enabled in server settings." type="bool" default="True" value="True" hidden="0" advanced="0" group="" />\n<Setting id="collectionMode" label="Collections" summary="How to display collections." type="int" default="2" value="2" hidden="0" advanced="0" group="" enumValues="0:Disabled|1:Hide items which are in collections|2:Show collections and their items" />\n</Preferences>\n<Pivot id="recommended" key="/hubs/sections/10" type="hub" title="Recommended" context="content.discover" symbol="star" />\n<Pivot id="library" key="/library/sections/10/all?type=1" type="list" title="Library" context="content.library" symbol="library" />\n</Directory>\n<Directory agent="tv.plex.agents.series" language="en-US" refreshing="0" scanner="Plex TV Series" uuid="bba1ee3b-9416-439d-8063-7f65d2124fdb" id="11" key="/library/sections/11" hubKey="/hubs/sections/11" type="show" title="TV Shows" updatedAt="1655532984" scannedAt="1655532984">\n<Preferences>\n<Setting id="hidden" label="Visibility" summary="Restrict where content from this library should appear." type="int" default="0" value="0" hidden="0" advanced="0" group="" enumValues="0:Include in home screen and global search|1:Exclude from home screen|2:Exclude from home screen and global search" />\n<Setting id="episodeSort" label="Episode sorting" summary="How to sort the episodes for this show." type="text" default="-1" value="-1" hidden="0" advanced="0" group="" enumValues="-1:Library default|0:Oldest first|1:Newest first" />\n<Setting id="country" label="Certification Country" summary="This will influence which content rating system is used. Changing this setting will require refreshing the metadata for the changes to take effect." type="text" default="US" value="US" hidden="0" advanced="0" group="" enumValues="AF:Afghanistan|AX:&#197;land Islands|AL:Albania|DZ:Algeria|AS:American Samoa|AD:Andorra|AO:Angola|AI:Anguilla|AQ:Antarctica|AG:Antigua &amp; Barbuda|AR:Argentina|AM:Armenia|AW:Aruba|AU:Australia|AT:Austria|AZ:Azerbaijan|BS:Bahamas|BH:Bahrain|BD:Bangladesh|BB:Barbados|BY:Belarus|BE:Belgium|BZ:Belize|BJ:Benin|BM:Bermuda|BT:Bhutan|BO:Bolivia|BA:Bosnia &amp; Herzegovina|BW:Botswana|BV:Bouvet Island|BR:Brazil|IO:British Indian Ocean Territory|VG:British Virgin Islands|BN:Brunei|BG:Bulgaria|BF:Burkina Faso|BI:Burundi|KH:Cambodia|CM:Cameroon|CA:Canada|IC:Canary Islands|CV:Cape Verde|BQ:Caribbean Netherlands|KY:Cayman Islands|CF:Central African Republic|EA:Ceuta &amp; Melilla|TD:Chad|CL:Chile|CN:China|CX:Christmas Island|CC:Cocos (Keeling) Islands|CO:Colombia|KM:Comoros|CG:Congo - Brazzaville|CD:Congo - Kinshasa|CK:Cook Islands|CR:Costa Rica|CI:C&#244;te d&#8217;Ivoire|HR:Croatia|CU:Cuba|CW:Cura&#231;ao|CY:Cyprus|CZ:Czechia|DK:Denmark|DG:Diego Garcia|DJ:Djibouti|DM:Dominica|DO:Dominican Republic|EC:Ecuador|EG:Egypt|SV:El Salvador|GQ:Equatorial Guinea|ER:Eritrea|EE:Estonia|SZ:Eswatini|ET:Ethiopia|FK:Falkland Islands|FO:Faroe Islands|FJ:Fiji|FI:Finland|FR:France|GF:French Guiana|PF:French Polynesia|TF:French Southern Territories|GA:Gabon|GM:Gambia|GE:Georgia|DE:Germany|GH:Ghana|GI:Gibraltar|GR:Greece|GL:Greenland|GD:Grenada|GP:Guadeloupe|GU:Guam|GT:Guatemala|GG:Guernsey|GN:Guinea|GW:Guinea-Bissau|GY:Guyana|HT:Haiti|HM:Heard &amp; McDonald Islands|HN:Honduras|HK:Hong Kong SAR China|HU:Hungary|IS:Iceland|IN:India|ID:Indonesia|IR:Iran|IQ:Iraq|IE:Ireland|IM:Isle of Man|IL:Israel|IT:Italy|JM:Jamaica|JP:Japan|JE:Jersey|JO:Jordan|KZ:Kazakhstan|KE:Kenya|KI:Kiribati|XK:Kosovo|KW:Kuwait|KG:Kyrgyzstan|LA:Laos|LV:Latvia|LB:Lebanon|LS:Lesotho|LR:Liberia|LY:Libya|LI:Liechtenstein|LT:Lithuania|LU:Luxembourg|MO:Macao SAR China|MG:Madagascar|MW:Malawi|MY:Malaysia|MV:Maldives|ML:Mali|MT:Malta|MH:Marshall Islands|MQ:Martinique|MR:Mauritania|MU:Mauritius|YT:Mayotte|MX:Mexico|FM:Micronesia|MD:Moldova|MC:Monaco|MN:Mongolia|ME:Montenegro|MS:Montserrat|MA:Morocco|MZ:Mozambique|MM:Myanmar (Burma)|NA:Namibia|NR:Nauru|NP:Nepal|NL:Netherlands|NC:New Caledonia|NZ:New Zealand|NI:Nicaragua|NE:Niger|NG:Nigeria|NU:Niue|NF:Norfolk Island|KP:North Korea|MK:North Macedonia|MP:Northern Mariana Islands|NO:Norway|OM:Oman|PK:Pakistan|PW:Palau|PS:Palestinian Territories|PA:Panama|PG:Papua New Guinea|PY:Paraguay|PE:Peru|PH:Philippines|PN:Pitcairn Islands|PL:Poland|PT:Portugal|PR:Puerto Rico|QA:Qatar|RE:R&#233;union|RO:Romania|RU:Russia|RW:Rwanda|WS:Samoa|SM:San Marino|ST:S&#227;o Tom&#233; &amp; Pr&#237;ncipe|SA:Saudi Arabia|SN:Senegal|RS:Serbia|SC:Seychelles|SL:Sierra Leone|SG:Singapore|SX:Sint Maarten|SK:Slovakia|SI:Slovenia|SB:Solomon Islands|SO:Somalia|ZA:South Africa|GS:South Georgia &amp; South Sandwich Islands|KR:South Korea|SS:South Sudan|ES:Spain|LK:Sri Lanka|BL:St. Barth&#233;lemy|SH:St. Helena|KN:St. Kitts &amp; Nevis|LC:St. Lucia|MF:St. Martin|PM:St. Pierre &amp; Miquelon|VC:St. Vincent &amp; Grenadines|SD:Sudan|SR:Suriname|SJ:Svalbard &amp; Jan Mayen|SE:Sweden|CH:Switzerland|SY:Syria|TW:Taiwan|TJ:Tajikistan|TZ:Tanzania|TH:Thailand|TL:Timor-Leste|TG:Togo|TK:Tokelau|TO:Tonga|TT:Trinidad &amp; Tobago|TN:Tunisia|TR:Turkey|TM:Turkmenistan|TC:Turks &amp; Caicos Islands|TV:Tuvalu|UG:Uganda|UA:Ukraine|AE:United Arab Emirates|GB:United Kingdom|US:United States|UY:Uruguay|UM:US Outlying Islands|VI:US Virgin Islands|UZ:Uzbekistan|VU:Vanuatu|VA:Vatican City|VE:Venezuela|VN:Vietnam|WF:Wallis &amp; Futuna|EH:Western Sahara|YE:Yemen|ZM:Zambia|ZW:Zimbabwe" />\n<Setting id="showOrdering" label="Episode ordering" summary="How the episodes are named on disk." type="text" default="tmdbAiring" value="tmdbAiring" hidden="0" advanced="0" group="" enumValues="tmdbAiring:The Movie Database|aired:TheTVDB" />\n<Setting id="useSeasonTitles" label="Use season titles" summary="Use season titles when available." type="bool" default="False" value="False" hidden="0" advanced="0" group="" />\n<Setting id="originalTitles" label="Use original titles" summary="Use the original titles for all items regardless of the library language." type="bool" default="False" value="False" hidden="0" advanced="0" group="" />\n<Setting id="localizedArtwork" label="Prefer artwork based on library language" summary="Use localized posters when available. This is determined by the library language setting." type="bool" default="True" value="True" hidden="0" advanced="0" group="" />\n<Setting id="useLocalAssets" label="Use local assets" summary="When scanning this library, use local posters and artwork if present." type="bool" default="True" value="True" hidden="0" advanced="0" group="" />\n<Setting id="respectTags" label="Prefer local metadata" summary="When scanning this library, prefer embedded tags and local files if present." type="bool" default="False" value="False" hidden="0" advanced="0" group="" />\n<Setting id="useExternalExtras" label="Find extras" summary="Find trailers and extras automatically (Plex Pass required)" type="bool" default="True" value="True" hidden="0" advanced="0" group="" />\n<Setting id="skipNonTrailerExtras" label="Only show trailers" summary="Skip extras which aren&#39;t trailers" type="bool" default="False" value="False" hidden="0" advanced="0" group="" />\n<Setting id="useRedbandTrailers" label="Allow red band trailers" summary="Use red band (restricted audiences) trailers when available" type="bool" default="False" value="False" hidden="0" advanced="0" group="" />\n<Setting id="includeExtrasWithLocalizedSubtitles" label="Localized subtitles" summary="Include extras with subtitles in Library language" type="bool" default="False" value="False" hidden="0" advanced="0" group="" />\n<Setting id="includeAdultContent" label="Include adult content" summary="Allow matching adult titles and returning adult titles in fix match results." type="bool" default="False" value="False" hidden="0" advanced="0" group="" />\n<Setting id="enableBIFGeneration" label="Enable video preview thumbnails" summary="Generate video preview thumbnails for items in this library when enabled in server settings." type="bool" default="True" value="True" hidden="0" advanced="0" group="" />\n<Setting id="collectionMode" label="Collections" summary="How to display collections." type="int" default="2" value="2" hidden="0" advanced="0" group="" enumValues="0:Disabled|1:Hide items which are in collections|2:Show collections and their items" />\n<Setting id="flattenSeasons" label="Seasons" summary="Choose whether to display seasons." type="int" default="0" value="0" hidden="0" advanced="0" group="" enumValues="0:Show|2:Hide for single-season series|1:Hide" />\n<Setting id="enableIntroMarkerGeneration" label="Enable intro detection" summary="Generate intro detection for items in this library when enabled in server settings." type="bool" default="True" value="True" hidden="0" advanced="0" group="" />\n</Preferences>\n<Pivot id="recommended" key="/hubs/sections/11" type="hub" title="Recommended" context="content.discover" symbol="star" />\n<Pivot id="library" key="/library/sections/11/all?type=2" type="list" title="Library" context="content.library" symbol="library" />\n</Directory>\n<Directory id="playlists" key="/playlists" type="playlist" title="Playlists">\n<Pivot id="playlists.video" key="/playlists?playlistType=video" type="list" title="Video" context="content.playlists.video" symbol="playlist" />\n</Directory>\n</Feature>\n<Feature key="/hubs/search" type="search">\n</Feature>\n<Feature key="/library/matches" type="match">\n</Feature>\n<Feature key="/library/metadata" type="metadata">\n</Feature>\n<Feature key="/:/rate" type="rate">\n</Feature>\n<Feature key="/photo/:/transcode" type="imagetranscoder">\n</Feature>\n<Feature key="/hubs/promoted" type="promoted">\n</Feature>\n<Feature key="/hubs/continueWatching" type="continuewatching">\n</Feature>\n<Feature key="/actions" type="actions">\n<Action id="removeFromContinueWatching" key="/actions/removeFromContinueWatching" />\n</Feature>\n<Feature flavor="universal" key="/playlists" type="playlist">\n</Feature>\n<Feature flavor="universal" key="/playQueues" type="playqueue">\n</Feature>\n<Feature key="/library/collections" type="collection">\n</Feature>\n<Feature scrobbleKey="/:/scrobble" unscrobbleKey="/:/unscrobble" key="/:/timeline" type="timeline">\n</Feature>\n<Feature type="manage">\n</Feature>\n<Feature type="queryParser">\n</Feature>\n<Feature flavor="download" type="subscribe">\n</Feature>\n</MediaProvider>\n</MediaContainer>\n`
	return fmt.Sprintf(str, server.Name, server.Hash, config.CONFIG.PlexEmail)
}

func PreferenceData(server *schemas.Server) PrefMediaResponse {
	return PrefMediaResponse{
		MediaContainer: PrefMediaContainer{
			Size: 1,
			Settings: []PrefSetting{
				{
					ID:       "FriendlyName",
					Label:    "Friendly name",
					Summary:  "This name will be used to identify this media server to other computers on your network. If you leave it blank, your computer's name will be used instead.",
					Type:     "text",
					Default:  "",
					Value:    server.Name,
					Hidden:   false,
					Advanced: false,
					Group:    "general",
				},
			},
		},
	}

}

func GetMetadataData(server *schemas.Server) fiber.Map {
	return fiber.Map{}
}
