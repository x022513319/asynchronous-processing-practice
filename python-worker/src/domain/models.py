from dataclasses import dataclass
import json
import os
from typing import Optional
from datetime import datetime


@dataclass
class TaskMessage:
    """對應 Go 的 task_message（從 Redis 讀取的格式）"""

    id: str
    type: str
    payload: dict  # Go 的 json.RawMessage 在 Python 是 dict

    @classmethod
    def from_json(cls, json_str: str):
        """從 Redis 取出的 JSON 字串建立物件"""
        data = json.loads(json_str)
        return cls(**data)


@dataclass
class TaskExecution:
    """Python Worker 執行時使用的擴展模型"""

    task_message: TaskMessage

    retry_count: int = 0
    max_retries: int = int(os.getenv("MAX_RETRIES", 3))

    idempotency_key: Optional[str] = None
    resource_id: Optional[str] = None

    started_at: Optional[datetime] = None
    error_message: Optional[str] = None

    def __post_init__(self):
        if not self.idempotency_key:
            self.idempotency_key = self.task_message.id

        if not self.resource_id and isinstance(self.task_message.payload, dict):
            self.resource_id = self.task_message.payload.get("resource_id")
