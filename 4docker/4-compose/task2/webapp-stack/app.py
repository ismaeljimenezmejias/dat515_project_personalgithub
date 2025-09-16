from flask import Flask, jsonify, request
import mysql.connector
import redis
import os
import time
from datetime import datetime

app = Flask(__name__)

# Database configuration
db_config = {
    'host': os.getenv('DB_HOST', 'database'),
    'user': os.getenv('DB_USER', 'webuser'),
    'password': os.getenv('DB_PASSWORD', 'webpass'),
    'database': os.getenv('DB_NAME', 'webapp')
}

# Redis configuration
redis_host = os.getenv('REDIS_HOST', 'cache')

@app.route('/')
def home():
    return jsonify({
        'message': 'Welcome to the Advanced Web App',
        'timestamp': datetime.now().isoformat(),
        'version': '1.0'
    })

@app.route('/health')
def health():
    try:
        # Check database
        conn = mysql.connector.connect(**db_config)
        conn.close()
        db_status = 'healthy'
    except:
        db_status = 'unhealthy'

    try:
        # Check Redis
        r = redis.Redis(host=redis_host, decode_responses=True)
        r.ping()
        cache_status = 'healthy'
    except:
        cache_status = 'unhealthy'

    return jsonify({
        'status': 'healthy' if db_status == 'healthy' and cache_status == 'healthy' else 'unhealthy',
        'database': db_status,
        'cache': cache_status,
        'timestamp': datetime.now().isoformat()
    })

@app.route('/counter')
def counter():
    try:
        r = redis.Redis(host=redis_host, decode_responses=True)
        count = r.incr('page_views')
        return jsonify({'page_views': count})
    except Exception as e:
        return jsonify({'error': str(e)}), 500

@app.route('/data', methods=['GET', 'POST'])
def data():
    try:
        conn = mysql.connector.connect(**db_config)
        cursor = conn.cursor()

        if request.method == 'POST':
            cursor.execute("INSERT INTO messages (content, created_at) VALUES (%s, %s)",
                         (request.json.get('message', 'Hello'), datetime.now()))
            conn.commit()

        cursor.execute("SELECT * FROM messages ORDER BY created_at DESC LIMIT 10")
        messages = [{'id': row[0], 'content': row[1], 'created_at': row[2].isoformat()}
                   for row in cursor.fetchall()]

        conn.close()
        return jsonify({'messages': messages})

    except Exception as e:
        return jsonify({'error': str(e)}), 500

if __name__ == '__main__':
    # Wait for database to be ready
    for i in range(30):
        try:
            conn = mysql.connector.connect(**db_config)
            conn.close()
            break
        except:
            time.sleep(1)

    app.run(host='0.0.0.0', port=5000)
