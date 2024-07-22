
import smtplib
from email.mime.text import MIMEText

def send_alert(message):
  msg = MIMEText(message)
  msg['Subject'] = 'Server Alert'
  msg['From'] = 'your_email@example.com'
  msg['To'] = 'admin@example.com'

  s = smtplib.SMTP('localhost')
  s.sendmail(msg['From'], [msg['To']], msg.as_string())
  s.quit()

# 示例：发送报警
anomalies = data[data['anomaly'] == -1]
if not anomalies.empty:
  send_alert(f"Detected {len(anomalies)} anomalies on the server.")


