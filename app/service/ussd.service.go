package service

func StartUSSD(sessionId, serviceCode, phoneNumber, text string) string {
	if text == "" {
		// This is the first request. Note how we start the response with CON
		return `CON What would you like to check
        1. My account
        2. My phone number`
	} else if text == "1" {
		// Business logic for first level response
		return `CON Choose account information you want to view
        1. Account number`
	} else if text == "1*1" {
		return "Oi Oi"
	}

	return "Hello World"
}
