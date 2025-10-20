import redis
import os

redis_host = os.getenv('REDIS_HOST', 'cache')

def get_redis():
    return redis.Redis(host=redis_host, decode_responses=True)