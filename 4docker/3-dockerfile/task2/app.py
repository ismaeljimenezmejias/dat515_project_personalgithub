from http.server import HTTPServer, SimpleHTTPRequestHandler
import os
import json
from datetime import datetime

class MyHandler(SimpleHTTPRequestHandler):
    def do_GET(self):
        if self.path == '/':
            self.send_response(200)
            self.send_header('Content-type', 'text/html')
            self.end_headers()

            html = f"""
            <!DOCTYPE html>
            <html>
            <head><title>Advanced Docker App</title></head>
            <body>
                <h1>Advanced Docker Application</h1>
                <p>Environment: {os.environ.get('APP_ENV', 'development')}</p>
                <p>Version: {os.environ.get('APP_VERSION', '1.0')}</p>
                <p>Current time: {datetime.now()}</p>
                <p>Working directory: {os.getcwd()}</p>
            </body>
            </html>
            """
            self.wfile.write(html.encode())
        else:
            super().do_GET()

if __name__ == '__main__':
    port = int(os.environ.get('PORT', 8000))
    server = HTTPServer(('0.0.0.0', port), MyHandler)
    print(f"Server starting on port {port}")
    server.serve_forever()
