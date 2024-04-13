import datetime as dt
import hashlib as hl


class Block:
    def __init__(self, index, timestamp, data, prevhash) -> None:
        self.index = index
        self.timestamp = timestamp
        self.data = data
        self.prevhash = prevhash
        self.hash = self.hash_block()

    def hash_block(self):
        block_encryption = hl.sha256()
        block_encryption.update(str(self.index).encode('utf-8') +
                                str(self.timestamp).encode('utf-8') +
                                str(self.data).encode('utf-8') +
                                str(self.prevhash).encode('utf-8'))
        return block_encryption.hexdigest()
    
    @staticmethod
    def genesisblock():
        return Block(0, dt.datetime.now(), "Genesis Block", "")
    
    @staticmethod
    def newblock(last_block):
        index = last_block.index + 1
        timestamp = dt.datetime.now()
        hashblock = last_block.hash
        data = f"Transaction {index}"
        return Block(index, timestamp, data, hashblock)

if __name__ == "__main__":
    blockchain = [Block.genesisblock()]
    previous_block = blockchain[0]
    num_of_blocks = 20
    print(f"Block ID {previous_block.index}")
    print(f"Timestamp: {previous_block.timestamp}")
    print(f"Hash: {previous_block.hash}")
    print(f"Previous Hash: {previous_block.prevhash}")
    print(f"Data: {previous_block.data}\n")

    for i in range(0, num_of_blocks):
        new_block = Block.newblock(previous_block)
        blockchain.append(new_block)
        previous_block = new_block
        print(f"Block ID {new_block.index}")
        print(f"Timestamp: {new_block.timestamp}")
        print(f"Hash: {new_block.hash}")
        print(f"Previous Hash: {new_block.prevhash}")
        print(f"Data: {new_block.data}\n")