https://notes.salrahman.com/generate-es256-es384-es512-private-keys/

SEEMS EXTREMELY USEFUL:
https://medium.com/@scottbrady91/creating-elliptical-curve-keys-using-openssl-4f155b35709e

prime256v1 IS secp256r1 (ergo a NIST and SECG curve - safe for production use - gov audited)

openssl ecparam -name $CURVE -genkey -noout -out private.ec.key

---

https://www.scottbrady91.com/jose/jwts-which-signing-algorithm-should-i-use
elyptic curve is faster for signing, but may be slower than rsa for verification
although key size is much smaller