## URL Shortener

modify InitClient(), replace Redis server address with yours.

```
go get github.com/gin-gonic/gin
go get github.com/go-redis/redis
go get github.com/itchyny/base58-go
```

```
go run main.go route.go
```

### hash function 

The `sha256.New()` function is used to create a new SHA-256 hash object, and `algorithm.Write([]byte(input))` is used to write the input string into the hash object. Finally, `algorithm.Sum(nil)` computes the final hash value and returns it as a byte slice.

### [binary to text encoding](https://en.wikipedia.org/wiki/Binary-to-text_encoding?ref=eddywm.com) algorithm

This binary to text encoding will be used to provide the final output of the process.
People familiar with binary to text encodings might ask why we are not using **[Base64](https://en.wikipedia.org/wiki/Base64?ref=eddywm.com)**, which is probably one of the most popular encoding scheme of this class of algorithms.

> The reasons are mostly related to the user experience, **Base58** reduces confusion in character output.  **[Satoshi Nakamoto](https://en.wikipedia.org/wiki/Satoshi_Nakamoto?ref=eddywm.com)** the anonymous early developer of the Bitcoin protocol and creator of the encoding scheme gave [some good reasons](https://en.bitcoin.it/wiki/Base58Check_encoding?ref=eddywm.com) on why Base58.

Those reasons are still valid  even today :

- The characters **0,O, I, l** are highly confusing when used in certain fonts and are even quite harder to differentiate for people with visuality issues .
- Removing ponctuations characters prevent confusion for line breakers.
- Double-clicking selects the whole number as one word if it's all alphanumeric.