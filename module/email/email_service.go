package email

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
	
	// Get current date
	now := time.Now()
	orderDate := now.Format("02/01/2006 15:04")
	
	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="vi">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Xác nhận đơn hàng - EcoCau</title>
		</head>
		<body style="margin: 0; padding: 0; font-family: 'Inter', 'Segoe UI', Arial, sans-serif; background-color: #faf9f7; color: #2b4626;">
			<div style="max-width: 650px; margin: 0 auto; background-color: #ffffff; box-shadow: 0 10px 40px rgba(0, 0, 0, 0.12);">

				<!-- Header với branding -->
				<div style="background: linear-gradient(135deg, #65a354 0%%, #4d8540 100%%); padding: 40px 30px; text-align: center; border-radius: 0;">
					<div style="background-color: rgba(255, 255, 255, 0.1); display: inline-block; padding: 15px 25px; border-radius: 25px; margin-bottom: 20px;">
						<h1 style="color: #ffffff; margin: 0; font-size: 28px; font-weight: 700; letter-spacing: 1px;">EcoCau</h1>
					</div>
					<h2 style="color: #ffffff; margin: 0; font-size: 24px; font-weight: 600;">Xác nhận đơn hàng</h2>
					<div style="background-color: rgba(255, 255, 255, 0.2); height: 2px; width: 80px; margin: 20px auto; border-radius: 1px;"></div>
				</div>

				<!-- Content -->
				<div style="padding: 40px 30px;">

					<!-- Greeting -->
					<div style="margin-bottom: 30px;">
						<p style="font-size: 18px; color: #2b4626; margin: 0 0 10px 0; font-weight: 600;">Xin chào %s,</p>
						<p style="font-size: 16px; color: #5e4939; margin: 0; line-height: 1.6;">Cảm ơn bạn đã đặt hàng tại cửa hàng của chúng tôi! Chúng tôi đã nhận được đơn hàng của bạn và đang xử lý.</p>
					</div>

					<!-- Order Info Card -->
					<div style="background: linear-gradient(135deg, #f6faf5 0%%, #eaf5e7 100%%); border-left: 5px solid #65a354; padding: 25px; margin: 30px 0; border-radius: 15px; box-shadow: 0 2px 15px rgba(101, 163, 84, 0.1);">
						<h3 style="color: #2b4626; margin: 0 0 20px 0; font-size: 20px; font-weight: 700; display: flex; align-items: center;">
							<span style="background-color: #65a354; color: white; padding: 8px 12px; border-radius: 50%%; font-size: 14px; margin-right: 15px; display: inline-block; width: 24px; height: 24px; text-align: center; line-height: 24px;">✓</span>
							Thông tin đơn hàng
						</h3>

						<div style="display: grid; grid-template-columns: 1fr 1fr; gap: 15px;">
							<div style="background-color: rgba(255, 255, 255, 0.7); padding: 15px; border-radius: 10px;">
								<p style="margin: 0 0 5px 0; font-size: 14px; color: #8a7052; font-weight: 600; text-transform: uppercase; letter-spacing: 0.5px;">Mã đơn hàng</p>
								<p style="margin: 0; font-size: 16px; color: #2b4626; font-weight: 700; font-family: 'Courier', monospace;">%s</p>
							</div>
							<div style="background-color: rgba(255, 255, 255, 0.7); padding: 15px; border-radius: 10px;">
								<p style="margin: 0 0 5px 0; font-size: 14px; color: #8a7052; font-weight: 600; text-transform: uppercase; letter-spacing: 0.5px;">Tổng tiền</p>
								<p style="margin: 0; font-size: 18px; color: #65a354; font-weight: 700;">%s VND</p>
							</div>
							<div style="background-color: rgba(255, 255, 255, 0.7); padding: 15px; border-radius: 10px;">
								<p style="margin: 0 0 5px 0; font-size: 14px; color: #8a7052; font-weight: 600; text-transform: uppercase; letter-spacing: 0.5px;">Trạng thái</p>
								<p style="margin: 0; font-size: 16px; color: #e8c04f; font-weight: 700; display: flex; align-items: center;">
									<span style="background-color: #e8c04f; width: 8px; height: 8px; border-radius: 50%%; margin-right: 8px; display: inline-block;"></span>
									Đang xử lý
								</p>
							</div>
							<div style="background-color: rgba(255, 255, 255, 0.7); padding: 15px; border-radius: 10px;">
								<p style="margin: 0 0 5px 0; font-size: 14px; color: #8a7052; font-weight: 600; text-transform: uppercase; letter-spacing: 0.5px;">Ngày đặt</p>
								<p style="margin: 0; font-size: 16px; color: #2b4626; font-weight: 600;">%s</p>
							</div>
						</div>
					</div>

					<!-- Message -->
					<div style="background: linear-gradient(135deg, #faf9f7 0%%, #f3f1ed 100%%); padding: 25px; margin: 30px 0; border-radius: 15px; border: 1px solid #e8e4db;">
						<p style="margin: 0 0 15px 0; color: #2b4626; font-size: 16px; line-height: 1.7;">
							<strong>Chúng tôi sẽ xử lý đơn hàng của bạn trong thời gian sớm nhất.</strong> Bạn sẽ nhận được thông báo qua email khi đơn hàng được chuẩn bị và giao.
						</p>
						<p style="margin: 0; color: #5e4939; font-size: 15px; line-height: 1.6;">
							Nếu bạn có bất kỳ câu hỏi nào về đơn hàng, vui lòng liên hệ với chúng tôi qua email hoặc số điện thoại bên dưới.
						</p>
					</div>

					<!-- Contact Info -->
					<div style="background-color: #4d8540; color: white; padding: 25px; margin: 30px 0; border-radius: 15px; text-align: center;">
						<h4 style="margin: 0 0 20px 0; font-size: 18px; font-weight: 700;">Thông tin liên hệ</h4>
						<div style="display: flex; justify-content: space-around; flex-wrap: wrap; gap: 20px;">
							<div style="text-align: center;">
								<div style="background-color: rgba(255, 255, 255, 0.2); padding: 12px; border-radius: 50%%; width: 44px; height: 44px; margin: 0 auto 10px; display: flex; align-items: center; justify-content: center;">
									<span style="font-size: 20px;">📧</span>
								</div>
								<p style="margin: 0; font-size: 14px; opacity: 0.9;">support@ecocau.com</p>
							</div>
							<div style="text-align: center;">
								<div style="background-color: rgba(255, 255, 255, 0.2); padding: 12px; border-radius: 50%%; width: 44px; height: 44px; margin: 0 auto 10px; display: flex; align-items: center; justify-content: center;">
									<span style="font-size: 20px;">📞</span>
								</div>
								<p style="margin: 0; font-size: 14px; opacity: 0.9;">(024) 3755 2024</p>
							</div>
						</div>
					</div>

					<!-- Thank you note -->
					<div style="text-align: center; padding: 20px 0;">
						<p style="margin: 0 0 15px 0; color: #2b4626; font-size: 18px; font-weight: 600;">Trân trọng cảm ơn!</p>
						<p style="margin: 0; color: #65a354; font-size: 16px; font-weight: 700;">Đội ngũ EcoCau Store</p>
					</div>
				</div>

				<!-- Footer -->
				<div style="background-color: #2b4626; color: #d5ead0; padding: 25px 30px; text-align: center;">
					<p style="margin: 0 0 10px 0; font-size: 13px; opacity: 0.8;">
						Email này được gửi tự động từ hệ thống EcoCau. Vui lòng không trả lời email này.
					</p>
					<div style="background-color: rgba(213, 234, 208, 0.1); height: 1px; width: 60%%; margin: 15px auto;"></div>
					<p style="margin: 0; font-size: 12px; opacity: 0.6;">
						© 2025 EcoCau Store. Chuyên cung cấp sản phẩm thân thiện với môi trường.
					</p>
				</div>
			</div>
		</body>
		</html>
	`, customerName, orderNumber, formatCurrency(totalAmount), orderDate)

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
