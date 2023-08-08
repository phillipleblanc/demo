import duckdb
import spicepy
from web3 import Web3
import os

def process(context: dict, 
            duckdb: duckdb.DuckDBPyConnection, 
            spice_client: spicepy.Client):
  api_key = os.environ["SPICE_API_KEY"]
  w3 = Web3(Web3.HTTPProvider(f"https://data.spiceai.io/eth?api_key={api_key}"))
  print(f"The latest block number from the node is {w3.eth.get_block_number()}")

  df = spice_client.query("SELECT MAX(number) as max_num FROM eth.recent_blocks").read_pandas()
  print(f"The latest block number from Spice is {df['max_num'].iloc[0]}")

  duckdb.sql(f"INSERT INTO output.hello_world VALUES ({context['block_number']}, 'Hello!')")

  print("Hello, World!")
