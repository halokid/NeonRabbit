from flask import Flask

# 创建 Flask 应用程序实例
app = Flask(__name__)

# 定义路由，当访问根 URL 时返回 "Hello, World!"
@app.route('/hello-post', methods=['POST'])
def hello_post():
    return 'hello post'


@app.route('/hello-get', methods=['GET'])
def hello_get():
    return 'hello get'


@app.route('/health')
def health():
    return 'ping'


# 运行 Flask 应用程序
if __name__ == '__main__':
    app.run(host='0.0.0.0', debug=True)

