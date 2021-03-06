package main

var tealLight = "int 1"

var tealNormal = `txn Receiver
addr 7JOPVEP3ABJUW5YZ5WFIONLPWTZ5MYX5HFK4K7JLGSIAG7RRB42MNLQ224
==
arg 0
len
int 32
==
&&
byte base64 iZWMx72KvU6Bw6sPAWQFL96YH+VMrBA0XKWD9XbZOZI=
sha256
byte base64 iZWMx72KvU6Bw6sPAWQFL96YH+VMrBA0XKWD9XbZOZI=
==
&&
txn Receiver
addr 7JOPVEP3ABJUW5YZ5WFIONLPWTZ5MYX5HFK4K7JLGSIAG7RRB42MNLQ224
==
txn FirstValid
int 0
>
&&
||
txn Fee
int 1000000
<
&&
int 1
||
`

var tealHeavy = `byte base64 iZWMx72KvU6Bw6sPAWQFL96YH+VMrBA0XKWD9XbZOZI=
byte base64 if8ooA+32YZc4SQBvIDDY8tgTatPoq4IZ8Kr+We1t38LR2RuURmaVu9D4shbi4VvND87PUqq5/0vsNFEGIIEDA==
addr 7JOPVEP3ABJUW5YZ5WFIONLPWTZ5MYX5HFK4K7JLGSIAG7RRB42MNLQ224
ed25519verify
byte base64 iZWMx72KvU6Bw6sPAWQFL96YH+VMrBA0XKWD9XbZOZI=
byte base64 if8ooA+32YZc4SQBvIDDY8tgTatPoq4IZ8Kr+We1t38LR2RuURmaVu9D4shbi4VvND87PUqq5/0vsNFEGIIEDA==
addr 7JOPVEP3ABJUW5YZ5WFIONLPWTZ5MYX5HFK4K7JLGSIAG7RRB42MNLQ224
ed25519verify
&&
byte base64 iZWMx72KvU6Bw6sPAWQFL96YH+VMrBA0XKWD9XbZOZI=
byte base64 if8ooA+32YZc4SQBvIDDY8tgTatPoq4IZ8Kr+We1t38LR2RuURmaVu9D4shbi4VvND87PUqq5/0vsNFEGIIEDA==
addr 7JOPVEP3ABJUW5YZ5WFIONLPWTZ5MYX5HFK4K7JLGSIAG7RRB42MNLQ224
ed25519verify
&&
byte base64 iZWMx72KvU6Bw6sPAWQFL96YH+VMrBA0XKWD9XbZOZI=
byte base64 if8ooA+32YZc4SQBvIDDY8tgTatPoq4IZ8Kr+We1t38LR2RuURmaVu9D4shbi4VvND87PUqq5/0vsNFEGIIEDA==
addr 7JOPVEP3ABJUW5YZ5WFIONLPWTZ5MYX5HFK4K7JLGSIAG7RRB42MNLQ224
ed25519verify
&&
byte base64 iZWMx72KvU6Bw6sPAWQFL96YH+VMrBA0XKWD9XbZOZI=
byte base64 if8ooA+32YZc4SQBvIDDY8tgTatPoq4IZ8Kr+We1t38LR2RuURmaVu9D4shbi4VvND87PUqq5/0vsNFEGIIEDA==
addr 7JOPVEP3ABJUW5YZ5WFIONLPWTZ5MYX5HFK4K7JLGSIAG7RRB42MNLQ224
ed25519verify
&&
byte base64 iZWMx72KvU6Bw6sPAWQFL96YH+VMrBA0XKWD9XbZOZI=
byte base64 if8ooA+32YZc4SQBvIDDY8tgTatPoq4IZ8Kr+We1t38LR2RuURmaVu9D4shbi4VvND87PUqq5/0vsNFEGIIEDA==
addr 7JOPVEP3ABJUW5YZ5WFIONLPWTZ5MYX5HFK4K7JLGSIAG7RRB42MNLQ224
ed25519verify
&&
byte base64 iZWMx72KvU6Bw6sPAWQFL96YH+VMrBA0XKWD9XbZOZI=
byte base64 if8ooA+32YZc4SQBvIDDY8tgTatPoq4IZ8Kr+We1t38LR2RuURmaVu9D4shbi4VvND87PUqq5/0vsNFEGIIEDA==
addr 7JOPVEP3ABJUW5YZ5WFIONLPWTZ5MYX5HFK4K7JLGSIAG7RRB42MNLQ224
ed25519verify
&&
byte base64 iZWMx72KvU6Bw6sPAWQFL96YH+VMrBA0XKWD9XbZOZI=
byte base64 if8ooA+32YZc4SQBvIDDY8tgTatPoq4IZ8Kr+We1t38LR2RuURmaVu9D4shbi4VvND87PUqq5/0vsNFEGIIEDA==
addr 7JOPVEP3ABJUW5YZ5WFIONLPWTZ5MYX5HFK4K7JLGSIAG7RRB42MNLQ224
ed25519verify
&&
byte base64 iZWMx72KvU6Bw6sPAWQFL96YH+VMrBA0XKWD9XbZOZI=
byte base64 if8ooA+32YZc4SQBvIDDY8tgTatPoq4IZ8Kr+We1t38LR2RuURmaVu9D4shbi4VvND87PUqq5/0vsNFEGIIEDA==
addr 7JOPVEP3ABJUW5YZ5WFIONLPWTZ5MYX5HFK4K7JLGSIAG7RRB42MNLQ224
ed25519verify
&&
byte base64 iZWMx72KvU6Bw6sPAWQFL96YH+VMrBA0XKWD9XbZOZI=
byte base64 if8ooA+32YZc4SQBvIDDY8tgTatPoq4IZ8Kr+We1t38LR2RuURmaVu9D4shbi4VvND87PUqq5/0vsNFEGIIEDA==
addr 7JOPVEP3ABJUW5YZ5WFIONLPWTZ5MYX5HFK4K7JLGSIAG7RRB42MNLQ224
ed25519verify
&&
int 1
||`
