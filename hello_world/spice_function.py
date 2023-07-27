import duckdb
import spicepy
import web3

def process(context: dict, 
            duckdb: duckdb.DuckDBPyConnection, 
            spice_client: spicepy.Client):
  # Temporary step
  duckdb.sql("""CREATE TABLE IF NOT EXISTS output.hello_world (
    block_number BIGINT,
    greeting TEXT
    )""")
  api_key = os.environ["SPICE_API_KEY"]
  w3 = Web3(Web3.HTTPProvider(f"https://data.spiceai.io/eth?api_key={api_key}"))
  print(f"The latest block number is {w3.eth.get_block_number()}")

  duckdb.sql(f"INSERT INTO output.hello_world VALUES ({context['block_number']}, 'Hello!')")

  print("Hello, World!")
