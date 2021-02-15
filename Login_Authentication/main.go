// JWTs, or  JSON Web Tokens allow you to transmit information from a client to the server in a stateless, but secure way.
//The JWT standard uses either a secret, using the HMAC algorithm, or a public/private key pair using RSA or ECDSA.
//These are heavily used within Single-Page Applications (SPAs) as a means of secure communications as they allow us to do two key things
//[1] Authentication - The most commonly used practice. Once a user logs in to your application, or authenticates in some manner,
//every request that is then sent from the client on behalf of the user will contain the JWT.
//[2] Information Exchange - The second use for JWTs is to securely transmit information between different systems. These JWTs
//can be signed using public/private key pairs so you can verify each system in this transaction in a secure manner and JWTs
//contain an anti-tamper mechanism as they are signed based off the header and the payload.

//we can use a JWT that has been signed with a secure key that both our client and server will have knowledge off.

//When our client goes to hit our server API, it will include this JWT as part of the request.
// CREATING SIMPLE REST API

//DOES NOT STORE PASSWORD IN JWT BECAUSE IT IS SHOWING PUBLICLY
//GO TO BROWSER AND CHECK
//http://localhost:8081/
package main

import (
	"log"
	"net/http"
)

func main() {
	//DEFINE HANDLERS
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)

	//START THE SERVER
	log.Fatal(http.ListenAndServe(":8080", nil))
}
