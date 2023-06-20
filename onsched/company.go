package onsched

type Company struct {
	Object                          string `json:"object"`
	ID                              string `json:"id"`
	Name                            string `json:"name"`
	RegistrationDate                string `json:"registrationDate"`
	RegistrationEmail               string `json:"registrationEmail"`
	DeletedStatus                   bool   `json:"deletedStatus"`
	DeletedTime                     string `json:"deletedTime"`
	AddressLine1                    string `json:"addressLine1"`
	AddressLine2                    string `json:"addressLine2"`
	City                            string `json:"city"`
	State                           string `json:"state"`
	PostalCode                      string `json:"postalCode"`
	Country                         string `json:"country"`
	Phone                           string `json:"phone"`
	Fax                             string `json:"fax"`
	Email                           string `json:"email"`
	Website                         string `json:"website"`
	TimezoneID                      string `json:"timezoneId"`
	TimezoneName                    string `json:"timezoneName"`
	NotificationFromEmailAddress    string `json:"notificationFromEmailAddress"`
	NotificationFromName            string `json:"notificationFromName"`
	BookingWebhookURL               string `json:"bookingWebhookUrl"`
	CustomerWebhookURL              string `json:"customerWebhookUrl"`
	ReminderWebhookURL              string `json:"reminderWebhookUrl"`
	ResourceWebhookURL              string `json:"resourceWebhookUrl"`
	WebhookSignatureHash            string `json:"webhookSignatureHash"`
	DisableEmailAndSmsNotifications bool   `json:"disableEmailAndSmsNotifications"`
}
