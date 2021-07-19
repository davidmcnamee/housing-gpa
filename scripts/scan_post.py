import asyncio
import os

import openai
from dotenv import load_dotenv

load_dotenv()

async def main():
  openai.api_key = os.getenv("OPENAPI_API_KEY")
  prompt = ["todo"]
  response = openai.Completion.create(engine="babbage", prompt=prompt, max_tokens=5)
  

  pass

asyncio.get_event_loop().run_until_complete(main())


