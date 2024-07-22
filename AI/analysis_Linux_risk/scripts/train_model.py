
from sklearn.ensemble import IsolationForest

# 示例：使用孤立森林检测异常行为
model = IsolationForest(contamination=0.01)
features = data[['user', 'log_length']]  # 假设已经提取了相关特征
model.fit(features)

# 检测异常
data['anomaly'] = model.predict(features)


