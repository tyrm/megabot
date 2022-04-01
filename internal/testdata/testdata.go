package testdata

import (
	"github.com/tmthrgd/go-hex"
	"github.com/tyrm/megabot/internal/chatbot"
	"github.com/tyrm/megabot/internal/models"
)

// TestUsers contains a set of test users
var TestUsers = []*models.User{
	{
		ID:                1,
		Email:             "test@example.com",
		EncryptedPassword: "$2a$14$iU.0NmiiQ5vdQefC77RTMeWpEbBUFsmyOOddo0srZHqXJF7oVC7ye",
	},
	{
		ID:                2,
		Email:             "test2@example.com",
		EncryptedPassword: "$2a$14$gleBixsHuNkr/TJGYbkTiOrci1J33778f/Nq39EAn7mlirR87XIx.",
	},
}

// TestGroupMembership contains a set of test group memberships
var TestGroupMembership = []*models.GroupMembership{
	{
		ID:      1,
		UserID:  TestUsers[0].ID,
		GroupID: models.GroupSuperAdmin(),
	},
}

// TestChatbotServices contains a set of test chatbot services.
var TestChatbotServices = []*models.ChatbotService{
	{
		ID:          1,
		Description: "Chatbot Service Mock 1",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("4bb2813f3e33ce35b881deb4d4009e6383c3b8ba0d53374aacb7a1f65cd8534c441baccdaf419aa450b4f88d"),
	},
	{
		ID:          2,
		Description: "Chatbot Service Mock 2",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("83fce20396e7128e740413e6cf66502d8049ecde3950fa6444f77bda4d793e01413e391266e19768b8cc49b2"),
	},
	{
		ID:          3,
		Description: "Chatbot Service Mock 3",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("307dc6bc1228e0e1c42d89fc1a1fbbfd0e0afe10bae5eb19017df4306154e1a060c39df98daa9211efc1b92b"),
	},
	{
		ID:          4,
		Description: "Chatbot Service Mock 4",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("92e272a1a51390491935496f54dfbdf1eade8c48fcc3a75cc6ec26a11032de69dcfd263a0334217fd9fab134"),
	},
	{
		ID:          5,
		Description: "Chatbot Service Mock 5",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("f9ee157e4a5740bf0cb4a69f020381d005b112af8a38a50c5d4f4771ddaefefc4e206220509f5987fc885786"),
	},
	{
		ID:          6,
		Description: "Chatbot Service Mock 6",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("7475ae95324594baa531131e9508e1a1cda9737eae562d74a021b94020114033c1ff3b9f44c5ac93e0887120"),
	},
	{
		ID:          7,
		Description: "Chatbot Service Mock 7",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("9b519855fd9ae14e8bf72b0a3b810848fab655739718c065e037c502359022ba3ddd7483795e5c691fd46511"),
	},
	{
		ID:          8,
		Description: "Chatbot Service Mock 8",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("8f9c91179355c59954c4e61aeb3944941360801b195bd25ca89e5b3ce4289993faa94daa55fb1406c79143a7"),
	},
	{
		ID:          9,
		Description: "Chatbot Service Mock 9",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("1007ebb7b8d76c5f3cba2ba455947fba51fe416dc002e3cf09cdea090880d0248889183955f727acb68e67f6"),
	},
	{
		ID:          10,
		Description: "Chatbot Service Mock 10",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("249e6687d1198ee689bcdf15c5b02b9e7be15259ca7b47610840670a352374cb609b1e5d33074477df5f1ae6d8"),
	},
	{
		ID:          11,
		Description: "Chatbot Service Mock 11",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("99f000a534facfbd9f1b64c63945d1ea2570df3ad2562413a77c6ca137c0e87bd75b3abd4c25468aab8cce124b"),
	},
	{
		ID:          12,
		Description: "Chatbot Service Mock 12",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("c2bf9b86c8873235daf92a58e620d7d5f5c128d3149de918a26493d234d56dae116e2ddb438f4af1e09759d895"),
	},
	{
		ID:          13,
		Description: "Chatbot Service Mock 13",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("e4228730400c3bd7bd4cfaa3e1923402e5a1dbf985f29f280ba316a1db9c552c4494f58f6b9a14818dc54dbbe7"),
	},
	{
		ID:          14,
		Description: "Chatbot Service Mock 14",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("3af424dec6dbeea99e2488009f01c8f9a42fd3cd0a01fe38a99809b78f67f9ac8c38aab9d9db11b200c6ef2243"),
	},
	{
		ID:          15,
		Description: "Chatbot Service Mock 15",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("8200c65e4fd8ae591e4c15dbf0212f96fbdea2ba4ca45711ef70daf1716dac22a0ddff87856a9c8b7ea224c95d"),
	},
	{
		ID:          16,
		Description: "Chatbot Service Mock 16",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("c859bd11cc9d303c773aeb82587bb54975849ac419caac88fc02c7342aad1d5386b11d8b2b74150c97bba36144"),
	},
	{
		ID:          17,
		Description: "Chatbot Service Mock 17",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("6167049ddc5db01984d7eb81427f0429aaa31a6cfc17dfc6c78bdfb1a7d68979bc2df6778808680e2ad087e9a1"),
	},
	{
		ID:          18,
		Description: "Chatbot Service Mock 18",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("940d44957f1098991e2989170e08ac769b28421b0ba0ffc9fa51a357406110e72d8c864e3a1168c0a97573fe6e"),
	},
	{
		ID:          19,
		Description: "Chatbot Service Mock 19",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("5b5c589d9819e3217087ad71d4c327a51b412d8bec1fdfcc1c20466fd7f76c3d8cdd26bd60d843a156a9b50c9c"),
	},
	{
		ID:          20,
		Description: "Chatbot Service Mock 20",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("6a4dd7d48f8197b620124d8c55276651b8e0c5eb065b74faec3e9c729e89ea3548fc336dbab9c0db19e1249937"),
	},
	{
		ID:          21,
		Description: "Chatbot Service Mock 21",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("7eff4d7cae7e457c89073eacb5fa55e0e9073dae4d9e4a629ef6545b265e71023c1463acbe74bb0c3f4a6cf23a"),
	},
	{
		ID:          22,
		Description: "Chatbot Service Mock 22",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("4021bc94394536c298c6eeb410b0d016b84fc8c4411c83b3bb08e4d0f5c83b63f63455f3021cb98f206697b8f2"),
	},
	{
		ID:          23,
		Description: "Chatbot Service Mock 23",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("d6596aee3762634eb44df95ef54eafef7eec5d5944f1a112b4108a19efa67c464971c78f7bfe0535adeb3678f2"),
	},
	{
		ID:          24,
		Description: "Chatbot Service Mock 24",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("4b2af9727bc1fbf46e637c41373e95e055bfd085dc3072a6f5dd220fc12ef694be21794c7f714f9c69109a2571"),
	},
	{
		ID:          25,
		Description: "Chatbot Service Mock 25",
		ServiceType: chatbot.ServiceMock,
		Config:      hex.MustDecodeString("110db2b23d4beda426ebfa8cc9bb7250b773eb14653fcbbf2be41cb724538818af680db040ab3e92212749b06e"),
	},
}

var TestChatbotServicesConfigs = []string{
	"{\"key\":\"test-1\"}",
	"{\"key\":\"test-2\"}",
	"{\"key\":\"test-3\"}",
	"{\"key\":\"test-4\"}",
	"{\"key\":\"test-5\"}",
	"{\"key\":\"test-6\"}",
	"{\"key\":\"test-7\"}",
	"{\"key\":\"test-8\"}",
	"{\"key\":\"test-9\"}",
	"{\"key\":\"test-10\"}",
	"{\"key\":\"test-11\"}",
	"{\"key\":\"test-12\"}",
	"{\"key\":\"test-13\"}",
	"{\"key\":\"test-14\"}",
	"{\"key\":\"test-15\"}",
	"{\"key\":\"test-16\"}",
	"{\"key\":\"test-17\"}",
	"{\"key\":\"test-18\"}",
	"{\"key\":\"test-19\"}",
	"{\"key\":\"test-20\"}",
	"{\"key\":\"test-21\"}",
	"{\"key\":\"test-22\"}",
	"{\"key\":\"test-23\"}",
	"{\"key\":\"test-24\"}",
	"{\"key\":\"test-25\"}",
}
