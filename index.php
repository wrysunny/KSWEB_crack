<?php
// php 7.0.*


// The version number (9_5_0) should match version of the Chilkat extension used, omitting the micro-version number.
// For example, if using Chilkat v9.5.0.48, then include as shown here:
include("./chilkat-9.5.0.86.2-php-7.0-x86_64-linux/chilkat_9_5_0.php");

// This example assumes the Chilkat API to have been previously unlocked.
// See Global Unlock Sample for sample code.

$rsa = new CkRsa();

// First load a public key object with a public key.
// In this case, we'll load it from a file.
$pubkey = new CkPublicKey();
$success = $pubkey->LoadFromFile('ksweb-public.key');
if ($success != true) {
    print $pubkey->lastErrorText() . "\n";
    exit;
}

// RSA encryption is limited to small amounts of data. The limit
// is typically a few hundred bytes and is based on the key size and
// padding (OAEP vs. PKCS1_5).  RSA encryption is typically used for
// encrypting hashes or symmetric (bulk encryption algorithm) secret keys.
$plainText = 'pass;2';

// Import the public key to be used for encrypting.
$success = $rsa->ImportPublicKeyObj($pubkey);

// To get OAEP padding, set the OaepPadding property:
$rsa->put_OaepPadding(true);

// To use SHA1 or SHA-256, set the OaepHash property
$rsa->put_OaepHash('sha256');
// for SHA1 --
$rsa->put_OaepHash('sha1');

// Indicate we'll want hex output
$rsa->put_EncodingMode('hex');

// Encrypt..
$usePrivateKey = false;
$encryptedStr = $rsa->encryptStringENC($plainText,$usePrivateKey);
print $encryptedStr . "\n";

// -------------------------------------------------
// Now decrypt with the matching private key.
$rsa2 = new CkRsa();

$privKey = new CkPrivateKey();
$success = $privKey->LoadEncryptedPem('ksweb-private.key','');
if ($success != true) {
    print $privKey->lastErrorText() . "\n";
    exit;
}

$success = $rsa2->ImportPrivateKeyObj($privKey);

// Make sure we have the same settings used for encryption.
$rsa2->put_OaepPadding(true);
$rsa2->put_EncodingMode('hex');
$rsa2->put_OaepHash('sha1');

$originalStr = $rsa2->decryptStringENC($encryptedStr,true);

print $originalStr . "\n";

?>
