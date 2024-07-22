
import pandas as pd

# 读取系统日志文件
data = pd.read_csv('syslog.txt', delimiter='\t', header=None, names=['timestamp', 'log'])

# 转换时间戳
data['timestamp'] = pd.to_datetime(data['timestamp'])
