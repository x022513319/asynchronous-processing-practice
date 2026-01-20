import redis


def new_redis_client(host: str, port: int, db: int) -> redis.Redis:
    # decode_responses: 自動把 Redis 的 bytes 轉成 Python 字串
    return redis.Redis(host=host, port=port, db=db, decode_responses=True)
