from src.utils.redis_client import new_redis_client

host = "localhost"
redis_port = 6379
db = 0


def main():
    # create connection
    # r = new_redis_client(host, redis_port, db)
    new_redis_client(host, redis_port, db)

    # TODO: start Worker
    # start Worker
