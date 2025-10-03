package email

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	host     string
	port     int
	username string
	password string
	from     string
}

type EmailData struct {
	To      string
	Subject string
	Body    string
}

func NewEmailService() *EmailService {
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if port == 0 {
		port = 587 // Default SMTP port
	}

	return &EmailService{
		host:     os.Getenv("SMTP_HOST"),
		port:     port,
		username: os.Getenv("SMTP_USERNAME"),
		password: os.Getenv("SMTP_PASSWORD"),
		from:     os.Getenv("SMTP_FROM"),
	}
}

func (e *EmailService) SendEmail(data EmailData) error {
	if e.host == "" || e.username == "" || e.password == "" {
		log.Println("Email configuration not found, skipping email send")
		return nil
	}

	m := gomail.NewMessage()
	m.SetHeader("From", e.from)
	m.SetHeader("To", data.To)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", data.Body)

	d := gomail.NewDialer(e.host, e.port, e.username, e.password)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Failed to send email: %v", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Printf("Email sent successfully to %s", data.To)
	return nil
}

func (e *EmailService) SendOrderConfirmationEmail(customerEmail, customerName, orderNumber string, totalAmount float64) error {
	subject := fmt.Sprintf("Xác nhận đơn hàng #%s", orderNumber)
	
	body := fmt.Sprintf(`
		<html>
		<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
			<div style="max-width: 600px; margin: 0 auto; padding: 20px;">
				<h2 style="color: #2c3e50; border-bottom: 2px solid #3498db; padding-bottom: 10px;">
					Xác nhận đơn hàng
				</h2>
				
				<p>Xin chào <strong>%s</strong>,</p>
				
				<p>Cảm ơn bạn đã đặt hàng tại cửa hàng của chúng tôi!</p>
				
				<div style="background-color: #f8f9fa; padding: 20px; border-radius: 5px; margin: 20px 0;">
					<h3 style="color: #2c3e50; margin-top: 0;">Thông tin đơn hàng:</h3>
					<p><strong>Mã đơn hàng:</strong> %s</p>
					<p><strong>Tổng tiền:</strong> %s VND</p>
					<p><strong>Trạng thái:</strong> Đang xử lý</p>
				</div>
				
				<p>Chúng tôi sẽ xử lý đơn hàng của bạn trong thời gian sớm nhất. Bạn sẽ nhận được thông báo khi đơn hàng được giao.</p>
				
				<p>Nếu bạn có bất kỳ câu hỏi nào, vui lòng liên hệ với chúng tôi.</p>
				
				<div style="margin-top: 30px; padding-top: 20px; border-top: 1px solid #eee; color: #666; font-size: 14px;">
					<p>Trân trọng,<br>Đội ngũ Mocau Store</p>
				</div>
			</div>
		</body>
		</html>
	`, customerName, orderNumber, formatCurrency(totalAmount))

	return e.SendEmail(EmailData{
		To:      customerEmail,
		Subject: subject,
		Body:    body,
	})
}

func formatCurrency(amount float64) string {
	// Format số tiền với dấu phẩy ngăn cách hàng nghìn
	return fmt.Sprintf("%.0f", amount)
}
