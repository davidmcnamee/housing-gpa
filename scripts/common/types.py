from dataclasses import dataclass
from typing import List

@dataclass
class FacebookPost:
  """Represents one facebook marketplace post"""
  link: str
  pics: List[str]
  main_text: str
  bottom_text: str
