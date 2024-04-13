const bcrypt = require('bcrypt')

class Block{ 
    constructor(blockid, previousHash, data){
        this.blockid = blockid;
        this.timestamp = Date.now();
        this.blockhash = this.getHash();
        this.previousHash = previousHash;
        this.data = data;
    }
    getHash(){
        return bcrypt.hashSync(String(this.blockid + this.timestamp + this.blockhash + this.previousHash + JSON.stringify(this.data)), 10)
    };
}

class Blockchain{
    constructor(){
        this.chain = [];
    }
    addBlock(data){
        let blockid = this.chain.length;
        let previousHash = this.chain.length !== 0 ? this.chain[this.chain.length - 1].blockhash : '';
        let block = new Block(blockid, previousHash, data);
        this.chain.push(block);
    }
}

const Myfirstblockchain = new Blockchain();

Myfirstblockchain.addBlock({sender: "Bruce wayne", receiver: "Tony Stark", amount: 24034});
Myfirstblockchain.addBlock({sender: "Harrison Wells", receiver: "Han Solo", amount: 32032});
Myfirstblockchain.addBlock({sender: "Tony Stark", receiver: "Bruce wayne", amount: 20993});

console.log(JSON.stringify(Myfirstblockchain, null, 6));