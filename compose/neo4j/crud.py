# neo4j crud operations

from neo4j import GraphDatabase

driver = GraphDatabase.driver("bolt://localhost:7687", auth=("neo4j", "password"))

# create
with driver.session() as session:
    session.run("CREATE (:Person {name: 'Alice', age: 30})")

# read
with driver.session() as session:
    result = session.run("MATCH (p:Person {name: 'Alice'}) RETURN p")

    for record in result:
        print(record["p"])
    
# update
with driver.session() as session:
    session.run("MATCH (p:Person {name: 'Alice'}) SET p.age = 31")

# read after update
with driver.session() as session:
    result = session.run("MATCH (p:Person {name: 'Alice'}) RETURN p")

    for record in result:
        print(record["p"])