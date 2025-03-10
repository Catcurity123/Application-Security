
# 1. Normal Case
+ `mvn clean compile exec:java`: compile and run the app

+ `curl -i -d '{"name": "test space", "owner": "demo"}' http://localhost:4567/spaces`: run the API

+ `curl -i -d "{\"name\": \"test'space\", \"owner\": \"demo\"}" http://localhost:4567/spaces`:  use escaped double quotes to properly format JSON string.

+ `curl -i -d '{"name": "test space", "owner": "demo"}' http://localhost:4567/spaces`: What would this translate to

=> INSERT INTO spaces (space_id, name, owner) VALUES (1, 'test space' ,'demo')

# 2. Case of ill-formatted entry that crashes both the database and the app itself.
+ `curl -i -d '{"name": "test'space", "owner": "demo"}' http://localhost:4567/spaces`
=> INSERT INTO spaces (space_id, name, owner) VALUES (2, 'test' space' ,'demo')
==> space' become syntax error

# 3. Case when user submit input as code and drop the fucking database
+ `curl -i -d "{\"name\": \"test\",\"owner\": \"'); DROP TABLE spaces; --\"}" http://localhost:4567/spaces`
=> INSERT INTO spaces (space_id, name, owner) VALUES (3, 'test', ''); DROP TABLE spaces; -- ');

+ `curl -i -d "{\"name\": \"', ''); DROP TABLE spaces; --\", \"owner\": \"\"}" http://localhost:4567/spaces`
=> INSERT INTO spaces (space_id, name, owner) VALUES (4, "', ''); DROP TABLE spaces; --", "")

+ `curl -d '{"name":"test", "owner":"a really long username that is more than 30 characters long"}' http://localhost:4567/spaces -i`: Error handling

