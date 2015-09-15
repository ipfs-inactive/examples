# Fgallery - Static image gallery

Fgallery comprises of some perl code, that can generate an image gallery built completely staticly with no server side code. This can then be easily hosted using ipfs.

http://www.thregr.org/~wavexx/software/fgallery/

## Dependencies

Fgallery reqires a few dependencies to work, on Ubuntu 14.04 I needed:

```
Perl, with the following modulesi (which can be installed from CPAN):
   Image::ExifTool
   JSON

imagemagick
liblcms2-utils
exiftran
```

## Building the gallery

Fistly, you'll need to prepare a directory with your images in. Make a note of where it's located on the file system.

The latest version of fgallery can be downloaded from their website, unzip the archive and change directory into the extracted folder.

Run:

```
./fgallery /directory/to/images /output/location
```

Output location can be anywhere on your system you want. Fgallery will then start processing images and building all the static files in the output directory:

```
reading completed
processing completed
generating archive...
completed
```

## Hosting with IPFS

Putting the images online is done with just one command:

```
ipfs add -r /output/location
```

Obviously replacing /output/location with the location you asked fgallery to output the files to.

This will add all of the files:

```
added QmYby7AidsKyM29z5TgVH6CGFkj5k9mauzJSSdRXmqoqXZ gallery/back.png
added Qma6c2Cz4Wp9ZsdtMUxzmc1rfTN1QewSdDkjdkqYkoYKJZ gallery/blurs/01O0FMD.jpg
added QmemKBx9awiTz6GubFuW6dQ3M7RnM7wiXaN6sA3NVXEAHM gallery/blurs/03Kj6JO.jpg
added QmZxRcxLzEvMiqEPxXECz8d2XdwCL9KvS9zKjVYHkw4t3P gallery/blurs/04CVyjY.jpg
added QmSTHdib78JDtPyWR498CbYnYWExVJLzD1298utjgMZkJj gallery/blurs/05vYu5k.jpg
added QmWKJo49HxhFnvTEW8r9D8vV7ezytn2n3u9XR2QhsUNAe2 gallery/blurs/06YYa0i.jpg
added Qmbc9U7yJKVRGWarSzDqFsFPZ5U4oFr3rnu7CMzxPSLj9R gallery/blurs/06uDGyB.jpg
added QmU3ski5amcbK1wQZVLHAhYpJRyv1vWfwVEfcJB5iRrzQ2 gallery/blurs/07v5H70.jpg
added QmWFVgR6i81FUXTS9HxppL8y1hy4DmqyS3Xse2uiGLmosj gallery/blurs/07vVDJI.jpg
added QmVmZBm5j6HZcAdYnPNJaSFL1FRxkyPkzJ7ixno1mGUNi4 gallery/blurs/097mGzX.jpg
added QmQMbgdCgNmSXJSoqm46nDq5QUjAduZ2RRRhgeB8pFnfmU gallery/blurs/0AKuTh0.jpg
added QmYjMJSrg3eWLEkH94Js2i8EgHWEgbYFV775BnPtPWMkdf gallery/blurs/GdRkiQe.jpg
added QmWhZjvDXM9W4Es2ZFeqYUSZgfB7B5URxCgXHcnjs2d9gk gallery/blurs
added QmPGmwDDBvbh3WwdPmwomBaVkRPcwKxjBqhBwuixoKQ5Ax gallery/cut-left.png
added QmUTzLj5XovXtrVKxNCEV2gewn3uHKuUt2CvfBVMn3aS6F gallery/cut-mov.png
added Qme591h16sr4d1GAd9DTQj2MzS7SGCDddaTjW9X8pmtdQP gallery/cut-right.png
added Qmdh4iSEzJ4JJ8TxpJqMWEFrEJGTXnibMycjFEi5fA2ga2 gallery/cut-top.png
added QmNZhsytZxPJYkYLFsR7p1xgusaDUq4XqPo7JxVjeokMAA gallery/data.json
added Qmc6QSF5vfAKAvRUmeozG39BnomLcjyGCcREbRFCxCGzbx gallery/download.png
added QmVwtBtRCngQJrPcFeq3DSa96FyLqBi7bfMQcSKxcDqLhU gallery/eye.png
added QmNbsd4SMzxSRNUm9J3a5sgBJg8H3tJUGU5bjJVRKkYugD gallery/files/album.zip
added Qma7ZL21cVBTYpMPFx9Z3TKiX3FWgwGog6LqWV8TzMXdrz gallery/files
added QmYksDd1jEtz6kcYb2A7WXiHd2e9iNjtRWhytbQ9pS9SeU gallery/imgs/01O0FMD.jpg
added QmTtYSk393PnukPUGtJtQiiT88ToMsq1rTi6DmZaT4NNiT gallery/imgs/03Kj6JO.jpg
added QmTSJAEzzrTxZc7QCYKquTnxYy93mMeBcxtzxFeVFfZfUU gallery/imgs/04CVyjY.jpg
added QmSbKdZoQfuuXCfkcvJHLVJEAxSf3e3ynpiB9tERJ4otgW gallery/imgs/05vYu5k.jpg
added Qme3prmgJwo3Y2Ao8H1XEcaRq5ZGCgySR4xskJJXBXsfty gallery/imgs/06YYa0i.jpg
added QmQZZjigYvdWdCigVh8icbRNiYueihHHRrVTc3A6RBQpx9 gallery/imgs/06uDGyB.jpg
added Qmbr6nyNEQR3STZ1WFg73Qh7HjjqQfQkJMeBaqiu1BzezW gallery/imgs/07v5H70.jpg
added QmcBkLJWpsG75yZbQnmcKrS2gvNKho44R9rS5mKuEi9uqo gallery/imgs/07vVDJI.jpg
added QmceXN5f82R5Rd8MCvhSUgR8ZA4SgNBSya1vg3hByVKwHu gallery/imgs/097mGzX.jpg
added QmQh7L4vTGGPuV7CbHVBgx7S6Y173oAUXDgE6bVTa6mdjW gallery/imgs/0AKuTh0.jpg
added QmQrryUqtM4UYYVLXiKbhyxBy9i75p46tHMHABrBQDqPWf gallery/imgs/GdRkiQe.jpg
added QmTQzSHt6B92KPzhtdtGE6DP7dvVJzCP2PuRrxZXJeekPX gallery/imgs
added QmUpYBMVqDrhfBCK2ghHjoaX62fQtgrrCxfabaAwYjeDSe gallery/index.css
added QmQeGFPJ6FrjG4GsiTZswZJac7dmLo6G21x9zoFZ1naczr gallery/index.html
added QmRVVLVVcvQnr4mwiNU266SQ49z3BosaiumjT9aydL54HC gallery/index.js
added QmQ1ERFyV2MxhFC5ng1ttVdg86dafivgQyLkwtWbmyHKjV gallery/left.png
added QmQKrKStaD7ijv2A4LP5weMwyXn9mZ6FJEpd4kax85W69h gallery/mootools-core-1.4.js
added Qmd9w91XMVfjHud9mTtGaAKJsxwCEZB2GaoVUA9j7s2oJz gallery/mootools-idle.js
added QmcJ7TBvQHCeQgBkJnidzZWoWDyDErq6BtrVCTAyCsmy3Y gallery/mootools-mooswipe.js
added QmWtHAixSGya2jToEWvAxLeKhiT1dkcXQr6kg7f6nPFdSq gallery/mootools-more-1.4.js
added QmTq42T1LBXiubNNEiGN8howjGKFWYPGUuiZQU68ffqvaq gallery/noise.png
added QmRziDBMHQZybCWF5s7kbe8MdQxugbmXR1WCgaLLGya4jb gallery/right.png
added QmQvs1n39aE9Ey7fFiBkKfNyJvST5LqTjMSUz9MWvKPMeb gallery/throbber.gif
added QmedcnTL7w34755XxnBMCf5zdXu4s57hkPsNxb7fDcVBJj gallery/thumbs/01O0FMD.jpg
added QmNvz7ZjBRo5TRiKgmw3JuhAvmvGibYmVHs7tdGmdKDF2H gallery/thumbs/03Kj6JO.jpg
added QmSopgp2tbbPh1TvTVjutWaAyAf4wfUTiwh65VHR56ZRVi gallery/thumbs/04CVyjY.jpg
added QmfDUXqikZQtx984nwkgf2uvm1H9qCzSkHkeaDvhKLZvoL gallery/thumbs/05vYu5k.jpg
added QmaknQgxfX2MS3We17EzVRssCmr53DzKmbz94bzLXmpKEE gallery/thumbs/06YYa0i.jpg
added QmU4YBZGAG5A7kmQmYZGJMc6z51kdnjtJY2RUfHoZrMw37 gallery/thumbs/06uDGyB.jpg
added QmY2K9Zd2kPntvwenjizWt7fGQvptLia94vkS5ejewkiKH gallery/thumbs/07v5H70.jpg
added QmVofJdrSKXeALbnDi6wesCciCs5weg51aXzD3AR9CU1ug gallery/thumbs/07vVDJI.jpg
added QmRCHjy36Dh9WwaNi6Tj2wAJvqF4cKaaJJwMku5micERX1 gallery/thumbs/097mGzX.jpg
added QmR2YruXg2mujXax6cKuQTHKSJ5zZvLCsCKX6MRnkGVKao gallery/thumbs/0AKuTh0.jpg
added QmQM49GgBY3WNEtT9xTiQrBEuVvYtAgcsG4scL3mQbaCRL gallery/thumbs/GdRkiQe.jpg
added QmUYC9AuQYAQPNCannRUyiNg94MbRHXUZYktMNRqwJbfvF gallery/thumbs
added QmedGgNBaGUDsc8anZ3LZ6HNjfKcoZGDpjoPxXwwKWjK2p gallery
```

Take the hash of the folder, in this case "QmedGgNBaGUDsc8anZ3LZ6HNjfKcoZGDpjoPxXwwKWjK2p", and add it to http://ipfs.io/ipfs/ to see the gallery live:

https://ipfs.io/ipfs/QmedGgNBaGUDsc8anZ3LZ6HNjfKcoZGDpjoPxXwwKWjK2p
