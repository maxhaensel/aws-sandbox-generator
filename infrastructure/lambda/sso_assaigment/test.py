import os
from distutils import util

x = util.strtobool(os.getenv("dry_run", False).lower())

# x = bool(os.getenv("dry_run", False).lower() in ("true", "1", "t"))

if x:
    print("wahr")
else:
    print("falsch")
