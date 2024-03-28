package resources_test

import (
	"fmt"
	"testing"

	"github.com/pingidentity/pingctl/internal/connector"
	"github.com/pingidentity/pingctl/internal/connector/pingone/platform/resources"
	"github.com/pingidentity/pingctl/internal/testutils"
)

func TestNotificationTemplateContentExport(t *testing.T) {
	// Get initialized apiClient and resource
	sdkClientInfo := testutils.GetPingOneSDKClientInfo(t)
	resource := resources.NotificationTemplateContent(sdkClientInfo)

	// Defined the expected ImportBlocks for the resource
	expectedImportBlocks := []connector.ImportBlock{
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Voice_en_Test Duplication on Device Pairing",
			ResourceID:   fmt.Sprintf("%s/device_pairing/d4ed6d8d-1b54-4903-970f-1c9896eed55d", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Email_en_Test Duplication on Device Pairing",
			ResourceID:   fmt.Sprintf("%s/device_pairing/2acfe36d-065c-465e-be21-cb95e46cee45", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_SMS_en_Test Duplication on Device Pairing",
			ResourceID:   fmt.Sprintf("%s/device_pairing/f67b076d-bb78-4cbd-b945-f721be9c88f6", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Voice_en",
			ResourceID:   fmt.Sprintf("%s/device_pairing/d02693ae-8809-4a7f-a7f9-da9f272c8096", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Email_en",
			ResourceID:   fmt.Sprintf("%s/device_pairing/625d98de-9f2d-4e1b-8417-d0ba139d36b2", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_SMS_en",
			ResourceID:   fmt.Sprintf("%s/device_pairing/d4ca6154-bf87-4201-825b-6a1fecbd66ac", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Voice_th",
			ResourceID:   fmt.Sprintf("%s/device_pairing/f123677b-c496-7700-1921-8b5d1f4fe213", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Voice_de",
			ResourceID:   fmt.Sprintf("%s/device_pairing/921bb61b-ea0e-786e-037a-76dabd6f7943", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Voice_ko",
			ResourceID:   fmt.Sprintf("%s/device_pairing/16d551c6-d1f1-7602-0245-4764f0812223", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Voice_zh",
			ResourceID:   fmt.Sprintf("%s/device_pairing/e9027287-dfd4-7431-157b-f46d3c5ab34d", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Email_fr-CA",
			ResourceID:   fmt.Sprintf("%s/device_pairing/732dd359-d5bd-7611-0005-4581d56c4b54", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_SMS_fr",
			ResourceID:   fmt.Sprintf("%s/device_pairing/c21bda2c-64b4-7025-2c83-d04d0f72077f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Voice_ru",
			ResourceID:   fmt.Sprintf("%s/device_pairing/41b1ab2d-2f8d-7db2-2c1a-de6aa7ca5e7c", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_SMS_ru",
			ResourceID:   fmt.Sprintf("%s/device_pairing/eb4149fe-4017-7f5d-3f3c-332e469d6587", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Email_es",
			ResourceID:   fmt.Sprintf("%s/device_pairing/73c5d0dd-62ba-7be7-31c6-b152edfff05c", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Email_nl",
			ResourceID:   fmt.Sprintf("%s/device_pairing/d51516ba-d0c4-7844-01bf-7d189165349a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_SMS_es",
			ResourceID:   fmt.Sprintf("%s/device_pairing/1cb17177-c38d-7895-174f-403ab707ef5b", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_SMS_ja",
			ResourceID:   fmt.Sprintf("%s/device_pairing/4a58be8a-2efa-7ca7-32cf-32318b1deea5", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Email_ru",
			ResourceID:   fmt.Sprintf("%s/device_pairing/bd0ccb52-596d-7b9a-1e9b-f8c289fea7cc", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_SMS_zh",
			ResourceID:   fmt.Sprintf("%s/device_pairing/454d0d8b-64c2-7ce7-190c-3df9b81434c6", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Email_pt",
			ResourceID:   fmt.Sprintf("%s/device_pairing/b691df72-4e8e-78dd-3174-646ea0a1f150", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_SMS_fr-CA",
			ResourceID:   fmt.Sprintf("%s/device_pairing/32ee5112-a7fc-7981-38b5-86d5dcbf7b05", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_Email_th",
			ResourceID:   fmt.Sprintf("%s/device_pairing/c72308f2-cc88-7318-2ea1-be12efdb96f0", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_SMS_th",
			ResourceID:   fmt.Sprintf("%s/device_pairing/6f64c300-86bd-7d79-0b83-682c87a280f2", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "device_pairing_SMS_ko",
			ResourceID:   fmt.Sprintf("%s/device_pairing/1add1335-bba1-7d10-063d-d7fdb6708aba", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "email_verification_admin_Email_en",
			ResourceID:   fmt.Sprintf("%s/email_verification_admin/b130f9a6-a422-72c0-3afa-105d5f8fbb88", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "email_verification_user_Email_en",
			ResourceID:   fmt.Sprintf("%s/email_verification_user/5eda6f7b-59c6-7c22-3348-9821179c2b37", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Email_ru",
			ResourceID:   fmt.Sprintf("%s/general/b412ca53-835f-7789-1e21-341c502c95ac", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_SMS_zh",
			ResourceID:   fmt.Sprintf("%s/general/22b77442-1c33-7155-2291-650b72285009", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Email_es",
			ResourceID:   fmt.Sprintf("%s/general/170217ff-f947-7b8c-183f-814ab265e30f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Voice_fr",
			ResourceID:   fmt.Sprintf("%s/general/831e9b77-5a05-7ed1-0fa6-c8cb637b5904", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_SMS_fr",
			ResourceID:   fmt.Sprintf("%s/general/63501c32-723c-7d4c-1f93-4e3c8c0cb292", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Voice_ja",
			ResourceID:   fmt.Sprintf("%s/general/aaceac3d-b0c7-7dd9-271c-0882864a08da", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Voice_th",
			ResourceID:   fmt.Sprintf("%s/general/fe9ef310-2989-7f6c-243d-8470115ab29f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Voice_es",
			ResourceID:   fmt.Sprintf("%s/general/641a1ef3-3df6-7f51-1d42-46502083e4bf", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Voice_tr",
			ResourceID:   fmt.Sprintf("%s/general/cd381b47-30c5-7378-144e-5b57f550922a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_SMS_pt",
			ResourceID:   fmt.Sprintf("%s/general/a8eed216-4cfd-7e28-0c7a-2790a9d8e9f8", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_SMS_th",
			ResourceID:   fmt.Sprintf("%s/general/f2edf0b7-d63f-7330-0d35-c5b49520dd53", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Voice_ru",
			ResourceID:   fmt.Sprintf("%s/general/8cb84ea8-fd49-7f83-29d5-ee3af090dc00", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_SMS_nl",
			ResourceID:   fmt.Sprintf("%s/general/fb0b491b-171a-771a-1e1c-512ec778dfa0", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_SMS_tr",
			ResourceID:   fmt.Sprintf("%s/general/0cb0d6c3-ae2d-7234-34bb-242ad05e3a5f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Voice_zh",
			ResourceID:   fmt.Sprintf("%s/general/21dfbdfb-19dc-7f99-1098-3bae6933795f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Email_zh",
			ResourceID:   fmt.Sprintf("%s/general/b9a031c7-d7a8-72ca-2160-a0e5978ee8f6", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_SMS_it",
			ResourceID:   fmt.Sprintf("%s/general/911546c3-b469-7b7b-1e91-dc26754dd586", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_SMS_ja",
			ResourceID:   fmt.Sprintf("%s/general/830325bd-53e4-7b1c-054b-9d1bff645e4b", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Email_fr-CA",
			ResourceID:   fmt.Sprintf("%s/general/3ba6eeec-c9cf-7f37-0721-0ff449888253", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_SMS_en",
			ResourceID:   fmt.Sprintf("%s/general/1dd4c1a3-802b-70c0-3d10-5524eb9defc7", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_SMS_ko",
			ResourceID:   fmt.Sprintf("%s/general/386d9e50-0032-73af-2599-71e3a1e90c43", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Voice_it",
			ResourceID:   fmt.Sprintf("%s/general/aa8739bc-0e80-717f-3cd1-ad309021ad75", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Voice_nl",
			ResourceID:   fmt.Sprintf("%s/general/6f64a454-985d-7387-024c-ccd6556da5ea", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_SMS_fr-CA",
			ResourceID:   fmt.Sprintf("%s/general/57871df8-08d3-7cf1-1bab-37c6428156bb", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "general_Email_ja",
			ResourceID:   fmt.Sprintf("%s/general/149a1660-40da-75e5-1ac3-321032293974", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_hu",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/cf3ddcfa-49c0-7f91-01f9-49f54f5fa46b", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_ja",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/a6f53259-ee85-7b29-1b65-db7797768a0b", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_SMS_it",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/b9f01eb1-cd75-738f-13b9-b854386d7bd7", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_zh",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/2555ddb2-d6e9-72f4-2326-81928f04e3a8", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_SMS_hu",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/37b26ed0-8e73-7f65-2df5-e15c2eeac195", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_SMS_th",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/5f6fbcf2-0d18-7e4a-2d15-1778ba0d724d", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_SMS_fr",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/a5dacd1c-c395-74ab-216f-a17037b22cf6", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_pl",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/ca145c95-43d7-7a48-13f8-d439da68e0d1", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_cs",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/10e267ed-0b7a-7203-304a-117624380015", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_SMS_ru",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/f1ccff08-4ca4-74fa-3374-f3949db8703b", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_SMS_zh",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/d14abb31-9f9e-7c4f-0ba3-1a898b1367ec", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_en",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/995558d3-39a9-72bf-32a6-e3c1e395aa1f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_es",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/f09287f6-0a4b-7174-09f8-e492b8d90f1b", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_SMS_ja",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/4b712b00-cd89-74db-1821-535f5a7b1055", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_pt",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/636cbcc4-2e1a-7657-1c68-7b3a9a0b29cc", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_it",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/8b55a1b4-9b64-7c37-21e5-fb2311f80cfd", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_de",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/60017a3c-c879-7293-24a1-e87978149dae", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_SMS_es",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/2d5d4be6-d4f9-70cb-1ff9-55263d786727", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_fr-CA",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/ee45f4cd-c78e-76ed-22d8-b9650ae03a78", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_nl",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/fbf6ac2f-ba5e-7809-0f2b-9929f761d323", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_th",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/d6f3d742-01f0-7554-28e3-f63a6efd12f0", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_SMS_fr-CA",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/1f23255e-8e02-7bdf-2bc3-1e6c97949f1c", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_tr",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/aa60b302-779f-70c8-087b-1a52a0d6964e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_SMS_ko",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/6169f085-bce7-7297-1137-d6ccc5cf426c", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "new_device_paired_Email_ru",
			ResourceID:   fmt.Sprintf("%s/new_device_paired/cc3d944f-5aac-7788-163d-3d05908799b2", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Voice_en",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/b5664ec8-a329-4ced-92ce-bbab388e7329", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Email_en",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/d1235c66-48c6-46ae-ae6d-599513ab26d7", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_SMS_en",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/41054e31-dacd-4591-a8c8-f44cbec6313f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Push_en",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/c6b2f1e9-fcde-4b64-b473-f5370219da76", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Email_tr",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/f78ffacb-ee5e-7284-3135-00ab200ab5e6", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Email_ja",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/e3bf81cb-0517-766c-3f2c-f2b46837f6c3", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_SMS_nl",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/8e15f0a6-6e1d-7e2c-3d53-8fd0e4852964", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Voice_es",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/5ee39c01-7ecb-7176-2db9-deaccc3f1395", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_SMS_fr",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/807cd1a1-f3f8-7440-10f5-5f9cf944abb3", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Voice_nl",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/b877c6c2-cfaa-7e26-02fe-33482f2b6cdb", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_SMS_ja",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/b5005a0e-913a-7c86-0bc3-3a11eb67b87b", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Voice_fr-CA",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/32284ea4-57a9-7b42-29e8-bc49951f61a4", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_SMS_pt",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/1b4fa5a9-8f29-7e8c-3e61-40457540dcd4", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Email_ru",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/995afe73-43ba-7375-101c-b85aa6c86f98", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Email_de",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/fc86c47c-6443-7735-1b89-7bcc5bc713f9", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Voice_th",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/e682a1f9-64af-7c47-3969-04cf83661da6", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Push_fr",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/d3d66f1b-b748-7afc-2d4b-a1daffd50a77", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Push_es",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/68ecf587-cc1b-737d-035b-00bc6473ab51", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Voice_it",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/5dbedf02-569f-7587-31b4-e20e5bf112a0", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Email_it",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/6be1ebc0-2579-736c-0737-8fe14e4abf6a", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Push_nl",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/88509f3a-51e5-7618-173d-ad8811f7b296", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_SMS_it",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/51a48e16-b73b-7b79-35f8-1764c5f5bade", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Push_de",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/c3cb8aea-118f-7c45-359a-7c9abea00e0e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_Email_zh",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/037f3874-fd32-7ac9-222a-419c2469b275", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "strong_authentication_SMS_fr-CA",
			ResourceID:   fmt.Sprintf("%s/strong_authentication/9158b297-2f88-7a1b-1591-07b449306c13", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_SMS_fr",
			ResourceID:   fmt.Sprintf("%s/transaction/10458132-7361-7d6b-3e42-04128ae31625", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Voice_ko",
			ResourceID:   fmt.Sprintf("%s/transaction/da316ad8-d09f-7717-0a1b-cbbdcc1d26b3", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_SMS_de",
			ResourceID:   fmt.Sprintf("%s/transaction/bbc3d144-2a78-7cca-0ea3-9376c950ff50", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Email_nl",
			ResourceID:   fmt.Sprintf("%s/transaction/21dd2f95-9ef7-731c-357e-19cb7bb75f40", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Email_th",
			ResourceID:   fmt.Sprintf("%s/transaction/13e0c37d-857a-74b9-31b5-8a1bcf8ba003", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Push_pt",
			ResourceID:   fmt.Sprintf("%s/transaction/a56d3760-d5ce-7ae1-03e2-c54cfbda642f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Email_fr-CA",
			ResourceID:   fmt.Sprintf("%s/transaction/8988bd44-b5b8-7b70-2fff-f207c607ef04", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Voice_ja",
			ResourceID:   fmt.Sprintf("%s/transaction/c43ecaf7-6400-710a-355a-3a02f5764092", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Push_ko",
			ResourceID:   fmt.Sprintf("%s/transaction/4de1a3dc-48be-74b8-3e08-422ebd1df3fd", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Push_de",
			ResourceID:   fmt.Sprintf("%s/transaction/8829a91b-c2d1-7027-3d12-e9840d8979a9", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Push_fr-CA",
			ResourceID:   fmt.Sprintf("%s/transaction/720367d3-3940-796d-0bcf-043338dfdd2c", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Email_ru",
			ResourceID:   fmt.Sprintf("%s/transaction/2a8272a6-c7b3-796d-36dc-b121bea8fe66", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Push_th",
			ResourceID:   fmt.Sprintf("%s/transaction/96238774-3058-7dc9-1464-46d7450178b8", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Voice_fr-CA",
			ResourceID:   fmt.Sprintf("%s/transaction/2e2630e1-670e-721a-0c0e-76caab0724dd", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Email_ko",
			ResourceID:   fmt.Sprintf("%s/transaction/c284f995-c294-7b91-0768-9407a27bd858", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_SMS_it",
			ResourceID:   fmt.Sprintf("%s/transaction/986af7e6-2c6d-7687-0105-ba1bd560e4c6", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_SMS_zh",
			ResourceID:   fmt.Sprintf("%s/transaction/2aaade7c-4ee3-7342-30e6-a805fb2f31e7", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Email_ja",
			ResourceID:   fmt.Sprintf("%s/transaction/59f60d16-29aa-70c9-1907-09f80f2e538f", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Voice_tr",
			ResourceID:   fmt.Sprintf("%s/transaction/332a628f-bdd5-7604-13df-f4ef7a0f989b", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_SMS_tr",
			ResourceID:   fmt.Sprintf("%s/transaction/060be0cd-4b54-7952-065a-d1df6e94d307", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Voice_nl",
			ResourceID:   fmt.Sprintf("%s/transaction/a3e34f26-93f9-7fd8-252e-774d05f55226", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Voice_pt",
			ResourceID:   fmt.Sprintf("%s/transaction/9e629a0d-264d-7519-236f-4187bfa602a0", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Email_it",
			ResourceID:   fmt.Sprintf("%s/transaction/74e46659-5609-7b1c-19c7-cec573cb44cf", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_Push_ru",
			ResourceID:   fmt.Sprintf("%s/transaction/41e9229c-6f86-7013-1941-31513da15cfd", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "transaction_SMS_ko",
			ResourceID:   fmt.Sprintf("%s/transaction/ce65e44d-8a4d-75f1-21ef-f0056fc4ae18", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_ja",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/b39ce15b-b391-7e76-0d65-6d5908a76070", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_zh",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/c23d1cfe-9d93-7c65-3b36-21bc57b15ac5", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_fr-CA",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/480a9958-88b2-748b-3e3c-0408c0302ea1", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_tr",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/d2ac7e7f-9b0d-7ab0-1361-5ddba89058c7", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_de",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/944164e5-5306-7ef8-138b-06c458337926", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_it",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/460dc781-943f-7578-177b-1fcce3d59849", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_th",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/256c1ad8-19d3-73ab-3f2d-92bec76832f7", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_es",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/74892957-f733-7dbc-0aef-e7137bd0aa9c", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_ru",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/5818998e-4209-772e-3e5b-a61ec890bea0", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_fr",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/03bdf108-c71d-74fb-28e8-143f22b98125", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_ko",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/0b38211b-779f-7404-0cae-4eb313c3a7bf", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_nl",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/900c409c-77d4-7910-3fb8-032a9251da2e", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_pt",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/ebb4b4dd-4d9f-71df-01ab-6cc3d552a5e2", testutils.GetEnvironmentID()),
		},
		{
			ResourceType: "pingone_notification_template_content",
			ResourceName: "verification_code_template_Email_en",
			ResourceID:   fmt.Sprintf("%s/verification_code_template/93688f61-e554-736d-227d-ac8ee610c254", testutils.GetEnvironmentID()),
		},
	}

	expectedImportBlocksMap := map[string]connector.ImportBlock{}
	for _, importBlock := range expectedImportBlocks {
		expectedImportBlocksMap[importBlock.ResourceName] = importBlock
	}

	testutils.ValidateImportBlocks(t, resource, &expectedImportBlocksMap)
}
