package api

type AlertChannelSlackConfig struct {
	WebhookURL string `json:"webhook_url"`
}

type AlertChannelEmailConfig struct {
	To       []string `json:"to"`
	From     string   `json:"from"`
	CC       []string `json:"cc"`
	BCC      []string `json:"bcc"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Host     string   `json:"host"`
	Port     int      `json:"port"`
}

type AlertChannelPagerDutyConfig struct {
	ServiceKey string `json:"service_key"`
}

type AlertChannelWebhookConfig struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}
