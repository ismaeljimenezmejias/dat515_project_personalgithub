import redis
import os

def get_redis():
    """Return a Redis client if REDIS_URL/REDIS_HOST is configured; otherwise None.
    Note: decode_responses must be False for Flask-Session because it stores
    pickled binary data; forcing UTF-8 decoding causes UnicodeDecodeError (0x80).
    """
    url = os.getenv('REDIS_URL', '').strip()
    if url:
        try:
            # Do NOT set decode_responses=True; session values are binary (pickle)
            return redis.from_url(url)
        except Exception:
            return None
    host = os.getenv('REDIS_HOST', '').strip()
    if not host or host.lower() in ('none', 'false', '0'):
        return None
    port = int(os.getenv('REDIS_PORT', '6379'))
    password = os.getenv('REDIS_PASSWORD') or None
    # Do NOT set decode_responses=True; session values are binary (pickle)
    return redis.Redis(host=host, port=port, password=password)