import redis
import os

def get_redis():
    """Return a Redis client if REDIS_HOST is configured; otherwise None."""
    host = os.getenv('REDIS_HOST', '').strip()
    if not host or host.lower() in ('none', 'false', '0'):  # treat empty/false-ish as disabled
        return None
    return redis.Redis(host=host, decode_responses=True)