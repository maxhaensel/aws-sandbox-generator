from typing import List, Dict, TypedDict
from enum import Enum


class Enviroments(Enum):
    prod = "prod"
    test = "test"


class Sandbox(TypedDict):
    name: str
    id: str


sandboxes: List[Sandbox] = [
    {"name": "Sandbox-4", "id": "815837829183"},
    # {"name": "test2", "id": "815837829183"},
    # {"name": "test3", "id": "815837829183"},
    # {"name": "test4", "id": "815837829183"},
]

sso_account = str("172920935848")
root_account = str("063661473261")
region = str("eu-central-1")
