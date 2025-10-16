// Timestamp: 2025-10-16 01:07:25

import re
def is_valid_email(email: str) -> bool:
    pattern = r'^[\\w\\.-]+@[\\w\\.-]+\\.\\w+$'
    return bool(re.match(pattern, email))

