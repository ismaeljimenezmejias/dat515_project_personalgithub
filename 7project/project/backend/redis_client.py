import redis
import os

def get_redis():
    """Return a Redis client if REDIS_URL/REDIS_HOST is configured; otherwise None."""
    url = os.getenv('REDIS_URL', '').strip()
    if url:
        try:
            return redis.from_url(url, decode_responses=True)
        except Exception:
            return None
    host = os.getenv('REDIS_HOST', '').strip()
    if not host or host.lower() in ('none', 'false', '0'):
        return None
    port = int(os.getenv('REDIS_PORT', '6379'))
    password = os.getenv('REDIS_PASSWORD') or None
    return redis.Redis(host=host, port=port, password=password, decode_responses=True)