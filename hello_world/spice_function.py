import duckdb
import spicepy

def process(context: dict, 
            duckdb: duckdb.DuckDBPyConnection, 
            spice_client: spicepy.Client):
  # Temporary step
  duckdb.sql("""CREATE TABLE IF NOT EXISTS output.hello_world (
    block_number BIGINT,
    greeting TEXT
    )""")

  duckdb.sql(f"INSERT INTO output.hello_world VALUES ({context['block_number']}, 'Hello!'")
