package notify

import (
	"fmt"

	"github.com/usual2970/certimate/internal/domain"
	"github.com/usual2970/certimate/internal/pkg/core/notifier"
	pBark "github.com/usual2970/certimate/internal/pkg/core/notifier/providers/bark"
	pDingTalk "github.com/usual2970/certimate/internal/pkg/core/notifier/providers/dingtalkbot"
	pEmail "github.com/usual2970/certimate/internal/pkg/core/notifier/providers/email"
	pGotify "github.com/usual2970/certimate/internal/pkg/core/notifier/providers/gotify"
	pLark "github.com/usual2970/certimate/internal/pkg/core/notifier/providers/larkbot"
	pMattermost "github.com/usual2970/certimate/internal/pkg/core/notifier/providers/mattermost"
	pPushover "github.com/usual2970/certimate/internal/pkg/core/notifier/providers/pushover"
	pPushPlus "github.com/usual2970/certimate/internal/pkg/core/notifier/providers/pushplus"
	pServerChan "github.com/usual2970/certimate/internal/pkg/core/notifier/providers/serverchan"
	pTelegram "github.com/usual2970/certimate/internal/pkg/core/notifier/providers/telegrambot"
	pWebhook "github.com/usual2970/certimate/internal/pkg/core/notifier/providers/webhook"
	pWeCom "github.com/usual2970/certimate/internal/pkg/core/notifier/providers/wecombot"
	maputil "github.com/usual2970/certimate/internal/pkg/utils/map"
)

// Deprecated: v0.4.x 将废弃
func createNotifierProviderUseGlobalSettings(channel domain.NotifyChannelType, channelConfig map[string]any) (notifier.Notifier, error) {
	/*
	  注意：如果追加新的常量值，请保持以 ASCII 排序。
	  NOTICE: If you add new constant, please keep ASCII order.
	*/
	switch channel {
	case domain.NotifyChannelTypeBark:
		return pBark.NewNotifier(&pBark.NotifierConfig{
			DeviceKey: maputil.GetString(channelConfig, "deviceKey"),
			ServerUrl: maputil.GetString(channelConfig, "serverUrl"),
		})

	case domain.NotifyChannelTypeDingTalk:
		return pDingTalk.NewNotifier(&pDingTalk.NotifierConfig{
			WebhookUrl: "https://oapi.dingtalk.com/robot/send?access_token=" + maputil.GetString(channelConfig, "accessToken"),
			Secret:     maputil.GetString(channelConfig, "secret"),
		})

	case domain.NotifyChannelTypeEmail:
		return pEmail.NewNotifier(&pEmail.NotifierConfig{
			SmtpHost:        maputil.GetString(channelConfig, "smtpHost"),
			SmtpPort:        maputil.GetInt32(channelConfig, "smtpPort"),
			SmtpTls:         maputil.GetOrDefaultBool(channelConfig, "smtpTLS", true),
			Username:        maputil.GetOrDefaultString(channelConfig, "username", maputil.GetString(channelConfig, "senderAddress")),
			Password:        maputil.GetString(channelConfig, "password"),
			SenderAddress:   maputil.GetString(channelConfig, "senderAddress"),
			ReceiverAddress: maputil.GetString(channelConfig, "receiverAddress"),
		})

	case domain.NotifyChannelTypeGotify:
		return pGotify.NewNotifier(&pGotify.NotifierConfig{
			ServerUrl: maputil.GetString(channelConfig, "url"),
			Token:     maputil.GetString(channelConfig, "token"),
			Priority:  maputil.GetOrDefaultInt64(channelConfig, "priority", 1),
		})

	case domain.NotifyChannelTypeLark:
		return pLark.NewNotifier(&pLark.NotifierConfig{
			WebhookUrl: maputil.GetString(channelConfig, "webhookUrl"),
		})

	case domain.NotifyChannelTypeMattermost:
		return pMattermost.NewNotifier(&pMattermost.NotifierConfig{
			ServerUrl: maputil.GetString(channelConfig, "serverUrl"),
			ChannelId: maputil.GetString(channelConfig, "channelId"),
			Username:  maputil.GetString(channelConfig, "username"),
			Password:  maputil.GetString(channelConfig, "password"),
		})

	case domain.NotifyChannelTypePushover:
		return pPushover.NewNotifier(&pPushover.NotifierConfig{
			Token: maputil.GetString(channelConfig, "token"),
			User:  maputil.GetString(channelConfig, "user"),
		})

	case domain.NotifyChannelTypePushPlus:
		return pPushPlus.NewNotifier(&pPushPlus.NotifierConfig{
			Token: maputil.GetString(channelConfig, "token"),
		})

	case domain.NotifyChannelTypeServerChan:
		return pServerChan.NewNotifier(&pServerChan.NotifierConfig{
			ServerUrl: maputil.GetString(channelConfig, "url"),
		})

	case domain.NotifyChannelTypeTelegram:
		return pTelegram.NewNotifier(&pTelegram.NotifierConfig{
			BotToken: maputil.GetString(channelConfig, "apiToken"),
			ChatId:   maputil.GetInt64(channelConfig, "chatId"),
		})

	case domain.NotifyChannelTypeWebhook:
		return pWebhook.NewNotifier(&pWebhook.NotifierConfig{
			WebhookUrl:               maputil.GetString(channelConfig, "url"),
			AllowInsecureConnections: maputil.GetBool(channelConfig, "allowInsecureConnections"),
		})

	case domain.NotifyChannelTypeWeCom:
		return pWeCom.NewNotifier(&pWeCom.NotifierConfig{
			WebhookUrl: maputil.GetString(channelConfig, "webhookUrl"),
		})
	}

	return nil, fmt.Errorf("unsupported notifier channel '%s'", channelConfig)
}
