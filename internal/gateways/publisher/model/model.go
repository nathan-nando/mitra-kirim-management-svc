package model

type PublisherEmail struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

/*
Type Email
1. Incoming
	From visitors
	- Email Service as system@mitrakirim.co.id will send 2 email:
		1. to visitor with thx message
		2. forwarded the original message to all@mitrakirim.co.id

2. Outgoing
	- Email service as admin@mitrakirim.co.id will send 2 email:
		1. to visitor with admin message
		2. forwarded the admin message to all@mitrakirim.co.id
*/
