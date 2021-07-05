<?php

include("chilkat_9_5_0.php");

$crypt = new CkCrypt2();

print "This is Chilkat version " . $crypt->version() . "\n";

$success = $crypt->UnlockComponent('Anything for 30-day trial');
if ($success != true) {
    print $crypt->lastErrorText() . "\n";
    exit;
}

//  AES is also known as Rijndael.
$crypt->CryptAlgorithm = 'aes';

//  CipherMode may be "ecb" or "cbc"
$crypt->CipherMode = 'cbc';

//  KeyLength may be 128, 192, 256
$crypt->KeyLength = 256;

//  The padding scheme determines the contents of the bytes
//  that are added to pad the result to a multiple of the
//  encryption algorithm's block size.  AES has a block
//  size of 16 bytes, so encrypted output is always
//  a multiple of 16.
$crypt->PaddingScheme = 0;

//  EncodingMode specifies the encoding of the output for
//  encryption, and the input for decryption.
//  It may be "hex", "url", "base64", or "quoted-printable".
$crypt->EncodingMode = 'hex';

//  An initialization vector is required if using CBC mode.
//  ECB mode does not use an IV.
//  The length of the IV is equal to the algorithm's block size.
//  It is NOT equal to the length of the key.
$ivHex = '000102030405060708090A0B0C0D0E0F';
$crypt->SetEncodedIV($ivHex,'hex');

//  The secret key must equal the size of the key.  For
//  256-bit encryption, the binary secret key is 32 bytes.
//  For 128-bit encryption, the binary secret key is 16 bytes.
$keyHex = '000102030405060708090A0B0C0D0E0F101112131415161718191A1B1C1D1E1F';
$crypt->SetEncodedKey($keyHex,'hex');

//  Encrypt a string...
//  The input string is 44 ANSI characters (i.e. 44 bytes), so
//  the output should be 48 bytes (a multiple of 16).
//  Because the output is a hex string, it should
//  be 96 characters long (2 chars per byte).
$encStr = $crypt->encryptStringENC('The quick brown fox jumps over the lazy dog.');
print $encStr . "\n";

//  Now decrypt:
$decStr = $crypt->decryptStringENC($encStr);
print $decStr . "\n";

?>

