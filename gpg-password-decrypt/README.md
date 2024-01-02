# About

This tool addresses the challenge of decrypting GPG-encrypted files on Windows and Mac OS without requiring additional software installation, catering to users who lack IT literacy or software installation capabilities.

![](images/situation.drawio.svg)

# How to Use

A sample procedure to decrypt gpg encrypted file with password.

Build the program

```sh
# on Mac
make build
# you can find bin/main.exe for Windows program
```

Create a gpg encrypted file on Linux or Mac OS.

```sh
# cd to your empty working directory
cd tmp
docker run -it --rm --workdir /home -v $(pwd):/home/ alpine sh

apk update && apk add gpg gpg-agent openssl
```

For demonstration, let's create below sample text file.

Please replace this file with your own file.

```sh
echo 'secret' > file.txt
```

copy generated string to clipboard

```sh
openssl rand -base64 20
# FUOMRTTqmXid1fhlMBn0VhcTylw=
```

```sh
gpg --symmetric --output file-aes256.txt.gpg --cipher-algo AES256 file.txt
# input the above random password

# Please consult with your security officer for recommended password length and cipher algorithm.

# /home # cat file-aes256.txt.gpg
# �       p�      �1y���D���1*�K������y���[<$�QoX��/aC
# �)�V�'�j�/��
```

Decrypt it

Again, use this program because the other end can't install any additional software.

For Windows, please share this program to the other end as well as the target file via email.

It's recommended to share passphrase via SMS or some other method than the above.

After decryption, you will find decrypted file name in the same directory as the source file.

Here is an example of executing the script on Windows command prompt.

```text
# pre-requisite
# download this program and target file to any working folder
# open cmd
# move to the download folder

# below sample, saved this program to sample folder and placed gpg file under sample/tmp/
z:\sample>dir tmp

 Directory of z:\sample\tmp

01/30/2023  07:15 AM                90 file-aes256.txt.gpg
               1 File(s)             90 bytes
               2 Dir(s)   7,641,030,656 bytes free

z:\sample>main.exe tmp/file-aes256.txt.gpg <password>

z:\sample>dir tmp

 Directory of z:\sample\tmp

01/30/2023  07:15 AM                90 file-aes256.txt.gpg
01/30/2023  07:17 AM                12 file-aes256.txt
               2 File(s)            102 bytes
               2 Dir(s)   7,636,475,904 bytes free
```

# Troubleshooting

If you give this program to Mac user, the program may not be allowed to open on Mac due to Mac security. Please refer to this [Stackoverflow article](https://stackoverflow.com/questions/4833052/how-do-i-remove-the-extended-attributes-on-a-file-in-mac-os-x).

What you need to do is to remove some attribute.

```shell
xattr -d com.apple.quarantine main-mac
```
