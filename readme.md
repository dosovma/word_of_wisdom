### Assignment

Test task for Server Engineer

Design and implement “Word of Wisdom” tcp server.

- TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work),
  the challenge-response protocol should be used.
- The choice of the POW algorithm should be explained.
- After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other
  collection of the quotes.
- Docker file should be provided both for the server and for the client that solves the POW challenge

### Prerequisites

Installed Docker https://docs.docker.com/engine/install/

### Local launch

- git clone https://github.com/dosovma/word_of_wisdom.git
- docker compose build
- docker compose up

You will see logs and quote from world of wisdom.

Client will finish automatically. Server should be stopped by executing CTRL+C.

### Project description

#### The choice of the POW algorithm should be explained

It was the most difficult issue. I spent a few hours to look into options and to compare its.

I would like to say that I read and understood all options. But I didn't.

My choice is Hashcash:

- It was invented to prevent denial-of-service attack. It's our target in assignment!
- It has clear algorithm and many descriptions.
- I started looking into the others and realised that I would spend hours just looking for information.

Finally, I see that it's a good fit for assignment

#### Algorithm and implementation

I hadn't had an experience with PoW, so I spent a few hours to understand what PoW is, and to look into implementations.

I project based on two articles which seemed comprehensive to me:

- https://en.wikipedia.org/wiki/Hashcash
- https://www.mdpi.com/1999-4893/16/10/462

You can find a lot of ideas from these articles in code.

#### Process

Getting random quote is 7 steps process:

1 - Client sends a request to get a token.

2 - Server responds with a challenge.

3 - Client solves challenge and sends a solution to Server.

4 - Server checks if the solution is suitable (there are many different solution for one challenge).

5 - Server responds with token.

6 - Client requests a quote including token into the request.

7 - Server validates token; and a random quote is sent to Client.

#### Code decisions that might be unclear

- MasterKey. I decided to use it to ensure server stateless. It's seemed to me as a good alternative of storing have
  gotten requests in database.
- Insufficient number of unit tests. It takes a time, so I wrote a few test to show that I am skilled in that.
- Message format. It's just for fun and to structure client-server communication. Please, don't be rigorous.

#### Message format

Each message has a predefined format.

```text
    "START:"
    "X-Command":[CommandType]
    Payload
    "END:"
```

Payload consists of strings with headers:

```text
	"X-Solution:"
	"X-Challenge:"
	"X-Token:"
	"X-Quote:"
	"X-Request-id:"
	"X-Request-time:"
```

#### Features that I would implement in real work

- To couple server load with difficulty. It provides evenly distribution of requests.
- Use hash tree https://www.mdpi.com/1999-4893/16/10/462 Merkle tree https://en.wikipedia.org/wiki/Merkle_tree to access
  to multiple quotes
- Do a refactor of mapping challenge and solution strings into structure (new entity).
- Add graceful degradation in Client and Server: return default value for errors from the other side.
- Increase test coverage
- Implement Retry logic in case failure of writing data
- Improve error handling: add more data, use a context
- Add auth token invalidation logic
- Implement handling a closed connection issue in client: now, when server closed connection due to timeout, client attempts to write data to closed connection.   
- Add shutdown logic to Server
- Use a context to perform read and write timeouts in client.
- Do access denied logic. I skipped this case in project on purpose.
- Move /pkg and const that describes message format to a separate github repository; make it public; and start using it
  as `go get github/.../tcp_communicator` in Client and Server.
- Move a few const to envs: default difficulty, masterKey, timeout.
- Read and write data into connection using []byte instead of strings. Strings make code more readable, but we shouldn't
  use it in production.  
