
# 示例：提取用户活动频率
data['user'] = data['log'].str.extract(r'user (\w+)')
user_activity = data.groupby('user').size()


