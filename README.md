# secp128r1
### 128-bit prime field Weierstrass curve.

A randomly generated curve. [SEC2v1](https://www.secg.org/SEC2-Ver-1.0.pdf) states 'E was chosen verifiably at random as specified in ANSI X9.62 from the seed'.

```
                  +---------------------------+
                  |        Private Key        |
                  +---------------------------+
                             |
                             | Private Key Operations
                             |
                             v
                  +---------------------------+
                  |     Signing Algorithm     |
                  +---------------------------+
                             |
                             | Signature
                             |
                             v
                  +---------------------------+
                  |        Signature          |
                  +---------------------------+
                             |
                             | Public Key Operations
                             |
                             v
                  +---------------------------+
                  |   Verification Algorithm  |
                  +---------------------------+
                             |
                             | Verification Result
                             |
                             v
                  +---------------------------+
                  |    True or False Output   |
                  +---------------------------+
```
