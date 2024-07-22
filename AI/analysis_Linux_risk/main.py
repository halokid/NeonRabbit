
import pandas as pd
import re
from sklearn.ensemble import IsolationForest

# 提取日期时间部分
def extract_datetime(s):
  match = re.match(r'^[A-Za-z]{3}\s+\d{1,2}\s+\d{2}:\d{2}:\d{2}', s)
  if match:
    return match.group(0)
  else:
    return None

def data_prepare():
  # 读取系统日志文件
  # data = pd.read_csv('./data/syslog.txt', delimiter='\t', header=None, names=['timestamp', 'log'])
  data = pd.read_csv('./data/syslog.txt', delimiter=' ', header=None, names=['timestamp', 'log'])
  print('-->>> data')
  print(data['timestamp'])
  print('-----------------------------')
  print(data['log'])
  # print(data[0])
  # print(data[1])

  # data['timestamp'] = data['timestamp'].apply(extract_datetime)
  # 将字符串转换为日期时间格式
  # data['timestamp'] = pd.to_datetime(data['timestamp'], format='%b %d %H:%M:%S', errors='coerce')

  # 转换时间戳
  data['timestamp'] = pd.to_datetime(data['timestamp'])
  return data

def feature_extraction(data):
  # 示例：提取用户活动频率
  data['user'] = data['log'].str.extract(r'user (\w+)')
  user_activity = data.groupby('user').size()
  return user_activity

def train_model(data):
  # 示例：使用孤立森林检测异常行为
  model = IsolationForest(contamination=0.01)
  features = data[['user', 'log_length']]  # 假设已经提取了相关特征
  model.fit(features)

  # 检测异常
  data['anomaly'] = model.predict(features)
  return data

def send_alert(message):
  print('send alert......')
  """
  msg = MIMEText(message)
  msg['Subject'] = 'Server Alert'
  msg['From'] = 'your_email@example.com'
  msg['To'] = 'admin@example.com'

  s = smtplib.SMTP('localhost')
  s.sendmail(msg['From'], [msg['To']], msg.as_string())
  s.quit()
  """

def monitor(data):
  anomalies = data[data['anomaly'] == -1]
  if not anomalies.empty:
    send_alert(f"Detected {len(anomalies)} anomalies on the server.")

if __name__ == '__main__':
  data = data_prepare()
  data = feature_extraction(data)
  data = train_model(data)
  monitor(data)






