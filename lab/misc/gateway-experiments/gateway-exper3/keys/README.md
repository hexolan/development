Generation of RS256:

ssh-keygen -t rsa -b 4096 -m PEM -C "" -f key.pem

---

Public Key Generation:

* Printing Out:
  * ssh-keygen -f key.pem.pub -e -m pkcs8
  * *(or just from private key directly:)* ssh-keygen -f key.pem -e -m pkcs8

* Saved Copy:
  * openssl rsa -in key.pem -pubout -outform PEM -out key.pem.pub
