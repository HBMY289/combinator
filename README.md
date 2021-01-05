# Combinator

## Scope

This program was written for an Iota user who was in a special situation with only having parts of his seed.
The program aimes at finding the correct seed to access the funds.

## Background
The user had written down a seed that had to be manipulated further to yield his actual seed. Unfortunately he only remembered part of how the seed had to be manipulated and thus did not have access to his funds anymore.
Known steps:
* base seed with 81 characters is available
* two words that are used to manipulate the seed are known as well
* two chunks of the length of each corresponding word had to be cut from the base seed
* the position where to cut is not known anymore
* the two words had to be inserted somewhere in the remaining seed to yield back 81 characters
* the position where to insert is not known anymore
* the user has written down the checksum of the final seed
* also an address of the final seed is known

## How the combinator works
When starting the program the user is first asked for all known input:
* base seed
* word 1
* word 2
* checksum of final seed
* address generated from final seed

The program then generates all possible combinations for cutting out chunks in the length of word 1 and word 2. It then iterates over these combinations and generates all possible seeds for each one where word 1 and word 2 are inserted at all possible positions.
The checksum of each of these seeds is compared to the known one. If it matches the first 10 addresses of that seed are generated and compared to the known address. If a match is found the program reports the matching seed.

## Speed
While there are surely ways to improve the code with regards to processing speed it proved to be fast enough for the challenge. For the given word lengths the combinator generated about 20 million possible seeds. With about 40.000 combinations tested per second the combinator could test all seeds in about 10 minutes.
