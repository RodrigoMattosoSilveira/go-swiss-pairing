# go-swiss-pairing
A tool for tournament directors to organize tournaments using the Swiss pairing method, written in Go

# What I want:
- [x] To use the Domain Driven Design approach and folder structure from go-swiss-pairing;
- To use the React / gRPC integration mechanism in react-grpc-go, but using the folder nexjs-ts-materio-template folder structure;
- To use the React Material components in nexjs-ts-materio-template

# The path to get what I want:
-  [x] Refactor the go-swiss-pairing Makefile to use dependencies, see react-grpc-go for reference
-  [ ] Refactor the go-swiss-pairing club-member service name to member;
-  [ ] Add a ping rpc handler to MemberService ;
-  [ ] Replace the go-swiss-pairing gRPC server (not service) with the react-grpc-go, service, including certificates, etc
-  [ ] Add the react-grpc-go UI to go-swiss-pairing, replacing pingpong PingPong with the Member ping call
-  [ ] Refactor the UI to use the nexjs-ts-materio-template folder structure
-  [ ] Replace the original simple react-grpc-go UI page with a simple nexjs-ts-materio-template component
-  [ ] Design the landing page, the members page, including logic to CRUD members

# Links and References
* Link to the [Recursive Wildcard Function](https://blog.jgc.org/2011/07/gnu-make-recursive-wildcard-function.html) I used to build a list of dependencies
