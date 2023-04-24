# secp128r1
### 128-bit prime field Weierstrass curve.

A randomly generated curve. [SEC2v1](https://www.secg.org/SEC2-Ver-1.0.pdf) states 'E was chosen verifiably at random as specified in ANSI X9.62 from the seed'.

Digital Signature:

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

ECDH:
```
           Alice                         Bob
           -----                         ---
           choose private key         
             d_A                         d_B
              |                           |
              v                           v
           compute public key:        compute public key:
           Q_A = d_A * G              Q_B = d_B * G
              |                           |
              v                           v
     ----- Begin Key Exchange Phase -----
              |                           |
              v                           v
           compute shared secret:     compute shared secret:
           S_A = d_A * Q_B            S_B = d_B * Q_A
              |                           |
              v                           v
     ----- End Key Exchange Phase -----
              |                           |
              v                           v
            (S_A)                       (S_B)
```
