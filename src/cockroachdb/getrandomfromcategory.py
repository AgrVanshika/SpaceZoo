import os
import psycopg2
from getpass import getpass
from random import randint
from sys import argv
from cockroachcreds import CDBUSER, CDBPASS

def getRandomFromCategory(category):
  
  conn_string = "postgresql://{}:{}@free-tier.gcp-us-central1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full&sslrootcert=$HOME/.postgresql/root.crt&options=--cluster%3Dzany-yak-4847".format(CDBUSER, CDBPASS)
  global conn
  conn = psycopg2.connect(os.path.expandvars(conn_string))

  #given a category, get a random row and display the name, description, and source
  with conn.cursor() as cur:
     cur.execute("SELECT COUNT(category) FROM spacezoo.spacezoo WHERE category='{}';".format(category))
     categorySize = cur.fetchone()[0]
     #assigns row numbers to every row in category=category and selects name, description, source from a random one with randint(1, categorySize)
     cur.execute("SELECT name, description, source FROM (SELECT ROW_NUMBER() OVER (ORDER BY name) AS TecNr, name, description, source FROM spacezoo.spacezoo WHERE category='{}') WHERE TecNr={};".format(category, randint(1, categorySize)))
     #saves results as a tuple 'technique' structured (name, description, category)
     technique = cur.fetchone()
    #prints in format (name,description,source)
     return(technique)

'''     
def main():
  #takes the cockroach connection string and formats it with user and password
  conn_string = "postgresql://{}:{}@free-tier.gcp-us-central1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full&sslrootcert=$HOME/.postgresql/root.crt&options=--cluster%3Dzany-yak-4847".format(CDBUSER, CDBPASS)
  global conn
  conn = psycopg2.connect(os.path.expandvars(conn_string))
  getRandomFromCategory(print(getRandomFromCategory(argv[1])))

  # Close communication with the database.
  conn.close()
'''



if __name__ == "__main__":
    main()
