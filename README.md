# SHA256TREE

SHA256TREE is a cryptographic hash function that has the following
properties:

- It uses the SHA-256 block cipher.
- For inputs larger than 1024 bytes, it uses a Merkle tree structure,
  just like BLAKE3.
- Unlike BLAKE3, it does not use a "chunk counter" when hashing the
  input, allowing repetitive parts of the input (e.g., long spans of
  zero bytes) to only be hashed once.

Like BLAKE3, it is possible to compute SHA256TREE hashes in parallel
using SIMD. That said, by using the more conventional SHA-256 block
cipher, it is also possible to use dedicated SHA-256 hardware
instructions that are provided by certain CPUs. This makes SHA256TREE
more energy efficient than BLAKE3.

This hash function was designed for use in combination with
Bazel's Remote Execution protocol. Please refer to pull request
[remote-apis#235](https://github.com/bazelbuild/remote-apis/pull/235)
for a more thorough specification of the hash function.
