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
		<!doctype html>
<html lang="vi">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Xác nhận đơn hàng - EcoCau</title>
  </head>
  <body
    style="
      margin: 0;
      padding: 0;
      background: #f5f7f4;
      font-family: 'Inter', 'Segoe UI', Arial, sans-serif;
      color: #2b4626;
    "
  >
    <div style="max-width: 620px; margin: 0 auto; background: #ffffff">
      <!-- Header -->
      <div
        style="
          padding: 36px 30px;
          text-align: center;
          background: linear-gradient(135deg, #66a355 0%, #4c7f3f 100%);
          color: #ffffff;
        "
      >
        <div style="display: inline-flex; align-items: center; gap: 12px; margin-bottom: 12px">
          <div
            style="
              width: 44px;
              height: 44px;
              border-radius: 50%;
              background: rgba(255, 255, 255, 0.18);
              display: flex;
              align-items: center;
              justify-content: center;
              font-weight: 700;
              font-size: 18px;
            "
          >
            E
          </div>
          <span style="font-size: 28px; font-weight: 700; letter-spacing: 0.5px">EcoCau</span>
        </div>
        <h1 style="margin: 0; font-size: 22px; font-weight: 600">Xác nhận đơn hàng</h1>
        <p style="margin: 10px 0 0; font-size: 14px; opacity: 0.9">
          Cảm ơn bạn đã đồng hành cùng sản phẩm thân thiện với môi trường của chúng tôi.
        </p>
      </div>

      <!-- Body -->
      <div style="padding: 32px 28px 36px">
        <!-- Greeting -->
        <p style="margin: 0 0 16px; font-size: 16px; font-weight: 600">
          Xin chào {{customer_name}},
        </p>
        <p style="margin: 0 0 26px; font-size: 15px; line-height: 1.6; color: #476040">
          Cảm ơn bạn đã đặt hàng tại EcoCau! Đơn hàng của bạn đang được xử lý. Chúng tôi sẽ thông
          báo ngay khi đơn được chuẩn bị xong để giao đến bạn.
        </p>

        <!-- Order summary -->
        <div
          style="
            border: 1px solid #e3eadf;
            border-radius: 14px;
            padding: 22px 24px;
            margin-bottom: 28px;
            background: #f9fbf8;
          "
        >
          <h2 style="margin: 0 0 18px; font-size: 18px; font-weight: 600; color: #2b4626">
            Thông tin đơn hàng
          </h2>
          <div style="display: flex; flex-wrap: wrap; gap: 18px 24px; font-size: 15px">
            <div style="flex: 1 1 240px">
              <div style="text-transform: uppercase; font-size: 12px; color: #829277; margin-bottom: 6px">
                Mã đơn hàng
              </div>
              <div style="font-weight: 600; letter-spacing: 0.5px">{{order_number}}</div>
            </div>
            <div style="flex: 1 1 240px">
              <div style="text-transform: uppercase; font-size: 12px; color: #829277; margin-bottom: 6px">
                Tổng tiền
              </div>
              <div style="font-weight: 600; color: #4d8540">{{total_amount}} VND</div>
            </div>
            <div style="flex: 1 1 240px">
              <div style="text-transform: uppercase; font-size: 12px; color: #829277; margin-bottom: 6px">
                Trạng thái
              </div>
              <div>
                <span
                  style="
                    display: inline-block;
                    padding: 4px 10px;
                    border-radius: 999px;
                    background: rgba(232, 192, 79, 0.16);
                    color: #b68b2d;
                    font-weight: 600;
                    font-size: 13px;
                  "
                  >Đang xử lý</span
                >
              </div>
            </div>
          </div>
        </div>

        <!-- Shipping -->
        <div
          style="
            border: 1px solid #e3eadf;
            border-radius: 14px;
            padding: 22px 24px;
            margin-bottom: 28px;
            background: #ffffff;
          "
        >
          <h2 style="margin: 0 0 16px; font-size: 18px; font-weight: 600; color: #2b4626">
            Thông tin giao hàng
          </h2>
          <div style="font-size: 15px; line-height: 1.7; color: #476040">
            <p style="margin: 0 0 6px"><strong style="color: #2b4626">Họ tên:</strong> {{shipping_name}}</p>
            <p style="margin: 0 0 6px">
              <strong style="color: #2b4626">Điện thoại:</strong> {{shipping_phone}}
            </p>
            <p style="margin: 0 0 6px">
              <strong style="color: #2b4626">Email:</strong> {{customer_email}}
            </p>
            <p style="margin: 0">
              <strong style="color: #2b4626">Địa chỉ:</strong> {{shipping_address}}
            </p>
          </div>
        </div>

        <!-- Message -->
        <div
          style="
            border-radius: 14px;
            padding: 22px 24px;
            background: #f1f5ef;
            border: 1px solid #e1e8e0;
            margin-bottom: 30px;
          "
        >
          <p style="margin: 0 0 12px; font-size: 15px; color: #2b4626; line-height: 1.7">
            <strong>Nhắc bạn một chút:</strong> đội ngũ EcoCau sẽ xử lý đơn hàng của bạn trong thời
            gian sớm nhất. Bạn sẽ nhận thêm email khi đơn được chuẩn bị và giao cho đơn vị vận
            chuyển.
          </p>
          <p style="margin: 0; font-size: 14px; color: #5f715d; line-height: 1.6">
            Nếu cần hỗ trợ hoặc muốn thay đổi/hủy đơn hàng, bạn chỉ cần trả lời lại email này hoặc
            liên hệ với chúng tôi qua thông tin bên dưới.
          </p>
        </div>

        <!-- Contact -->
        <div
          style="
            border-radius: 14px;
            padding: 20px 24px;
            background: #ffffff;
            border: 1px solid #e3eadf;
            margin-bottom: 32px;
          "
        >
          <h3 style="margin: 0 0 14px; font-size: 16px; font-weight: 600; color: #2b4626">
            Thông tin liên hệ
          </h3>
          <div style="display: flex; flex-wrap: wrap; gap: 12px 24px; font-size: 14px; color: #476040">
            <div style="flex: 1 1 220px">
              <div style="font-weight: 600; color: #2b4626">Email</div>
              <div>ecocauviet@gmail.com</div>
            </div>
            <div style="flex: 1 1 220px">
              <div style="font-weight: 600; color: #2b4626">Hotline</div>
              <div>076 348 4202</div>
            </div>
          </div>
        </div>

        <div style="text-align: center">
          <p style="margin: 0 0 10px; font-size: 16px; font-weight: 600">Trân trọng cảm ơn!</p>
          <p style="margin: 0; font-size: 15px; color: #4d8540; font-weight: 600">
            Đội ngũ EcoCau Store
          </p>
        </div>
      </div>

      <!-- Footer -->
      <div style="background: #2b4626; color: #dce8d8; text-align: center; padding: 22px 28px">
        <p style="margin: 0 0 10px; font-size: 12px; opacity: 0.75">
          Email này được gửi tự động từ hệ thống EcoCau. Vui lòng không trả lời email này.
        </p>
        <p style="margin: 0; font-size: 12px; opacity: 0.6">
          © {{current_year}} EcoCau Store — Đồng hành vì sản phẩm thân thiện với môi trường.
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
