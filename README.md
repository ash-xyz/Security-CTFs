# HackTrinity 2020 Solutions

Write up's for HackTrinity 2020.
This is a bit brief but I'll add more as we go along.

I started preparing for this CTF around 3-4 days prior and this was my first CTF so it was certainly difficult.

There were a few helpful resources that helped me. Primarily [Conor Mac Amhlaoibh](https://github.com/conormccauley1999/HackTrinity19Writeups/blob/master/README.md)'s write up for last years HackTrinity which helped me get a good feel for competition.

## Solved Problems


### Stegosaurus
Topic(s) | Points | Difficulty (in my opinion)
---------|--------|-----------
`forensics` | `388` | ![Hard](https://img.shields.io/badge/-Hard-orange.svg)
#### Challenge:
    My friend Wu said this was a better stegosaurus picture. Personally I can't see any DIFFERENCE, from the Previous Visual Depictions. Can you?
![flag.png](images/Stegosaurus/flag.png)
#### Solution: 
This problem was really interesting and took me a while to figure out. 

The key was figuring out the few hints in the challenge description. Wu isn't some random person, he's the author of a paper on [Pixel Value Differencing](https://people.cs.nctu.edu.tw/~whtsai/Journal%20Paper%20PDFs/Wu_&_Tsai_PRL_2003.pdf).

A quick read of the paper specifies an algorithm to encode and decode these images but I wasn't arsed to implement it myself so I looked one up on github and found one with nearly the exact same challenge description by [zst-ctf](https://github.com/zst-ctf/tjctf-2019-writeups/blob/master/Writeups/Planning_Virtual_Distruction/README.md).

I quickly tried his python script only to find that it didn't work.
I read over the paper again and found another set of range widths [2, 2, 4, 4, 4,8, 8, 16, 16, 32, 32, 64, and 64] and adjusted the if statements to suit. [Adjusted Python Script](src/Stegosaurus/Stegosaurus.py)

Running the program outputted the following:
![flag.png](images/Stegosaurus/output.png)
With the flag being: HackTrinity{Th3_H4ck3r_M4n1f35t0}

### Locked Out 1
Topic(s) | Points | Difficulty
---------|--------|-----------
`Reversing` | `288` | ![Medium](https://img.shields.io/badge/-Medium-yellow.svg)

##### Challenge:
    So I found this program on a hacking forum. The OP attached it to the thread saying he'd found a way to share his flag and you can't read it without knowing the password. I have no idea what the password is, but there were a lot of others laughing at OP saying they got the flag. Can you get it for me?
[Link](src/LockedOut/locked_out)
#### Solution:
This challenge actually has 2 solutions and the only reason I'm giving it a medium is because I took so long to find the reversing one. The second solution is discussed in [Locked Out 2](#Locked-Out-2)

The solution is rather simple and I think it's best described by a [LiveOverflow Video](https://www.youtube.com/watch?v=LyNyf3UM9Yc)

But all we're doing is going into Binary Ninja and patching the if statement so that it always outputs the flag. Save it is locked_out_cracked, run it and there's the flag.
##### Binary Ninja
![binaryninja](images/LockedOut/BinaryNinja.png)
##### Flag
![flag](images/LockedOut/flag.png)
### Locked Out 2
Topic(s) | Points | Difficulty
---------|--------|-----------
`Exploitation` | `394` | ![Easy](https://img.shields.io/badge/-Easy-green.svg)
#### Challenge:
    So back on the hacking forum, OP sick of being laughed at, said he's secured his flag by running his code on a remote server - he hasn't changed the code though.
#### Solution:
This challenge was a lot handier because I had already had the idea of abusing bufferoverflows to retrieve the flag.

The code was the exact same except this locked out file was located on a server and when you connect to it, there's no room for commands. This could only be a buffer overflow attack.

Since we had the code we just needed to figure out how many 'A's I'd need to overflow the buffer. I basically brute forced this by getting a really high number and a low number and found the number of A's needed .
##### Overflow
![overflow](images/LockedOut2/overflow.png)

Next I needed to find the location of the flag. I found a function called print_flag using gdb and then found its location.

##### print_flag location
![print_flag](images/LockedOut2/printFlagLocation.png)

After that I used python to prepare our overflow attack and execute the program

The struct module was handy because it converted our function location and transferred it instantly
```
python -c "import struct; print 'A'*264 + struct.pack('<I',0x000000004000139b)" | ./locked_out
```
![test](images/LockedOut2/test.png)

Now that I verified it worked for my test, I was ready to do it remotely. 

Place the following into the terminal, and our flag is revealed

```
python -c "import struct; print 'A'*264 + struct.pack('<I',0x000000004000139b)" | nc 192.168.146.1 1337
```
Flag: HackTrinity{bonfire_vet_TV_misplacement}
