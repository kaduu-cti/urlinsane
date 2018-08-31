# URLInsane

Multilingual domain typo permutation engine used to perform or detect typosquatting, 
brandjacking, URL hijacking, fraud, phishing attacks, corporate espionage and 
threat intelligence.





Table of contents
=================

<!--ts-->
   * [Table of contents](#table-of-contents)
   * [Introduction](#introduction)
   * [Installation](#installation)
   * [Usage](#usage)
   * [Features](#features)
   * [Languages](#languages)
      * [English](#english)
   * [Algorithms](#algorithms)
   * [Extra Functions](#extra-functions)
      * [TODO](#todo)
   * [Authors](#authors)
   * [License](#license)
<!--te-->



## Introduction

The engine is designed to execute concurrent typo algorithms then additional 
concurrent functions for each domain variation. The additional functions can 
check DNS records and much more. Its also designed for extensibility, allowing 
developers to add functionality and support for additional languages. See 
[URLInsane](https://rangertaha.github.io/urlinsane/) for more details.


## Installation

Create the binary executable with the make command or 
[download](https://github.com/rangertaha/urlinsane/releases/tag/0.2.0) one of the 
pre-built release binaries. 

```
make
```

## Usage
Generates variations for `google.com` using the character omission **(CO)** 
algorithm.
```
urlinsane google.com -t co

 _   _  ____   _      ___
| | | ||  _ \ | |    |_ _| _ __   ___   __ _  _ __    ___
| | | || |_) || |     | | | '_ \ / __| / _' || '_ \  / _ \
| |_| ||  _ < | |___  | | | | | |\__ \| (_| || | | ||  __/
 \___/ |_| \_\|_____||___||_| |_||___/ \__,_||_| |_| \___|

 Version: 0.1.0

  LIVE | TYPE |   TYPO    | SUFFIX |   IDNA     
+------+------+-----------+--------+-----------+
       | CO   | oogle.com | com    | oogle.com  
       | CO   | gogle.com | com    | gogle.com  
       | CO   | gogle.com | com    | gogle.com  
       | CO   | goole.com | com    | goole.com  
       | CO   | googe.com | com    | googe.com  
       | CO   | googl.com | com    | googl.com

```



Generates variations for `google.com` using the character omission **(CO)** 
algorithm and check for **ip** addresses. 
```
urlinsane google.com -t co -x ip

 _   _  ____   _      ___
| | | ||  _ \ | |    |_ _| _ __   ___   __ _  _ __    ___
| | | || |_) || |     | | | '_ \ / __| / _' || '_ \  / _ \
| |_| ||  _ < | |___  | | | | | |\__ \| (_| || | | ||  __/
 \___/ |_| \_\|_____||___||_| |_||___/ \__,_||_| |_| \___|

 Version: 0.1.0

   LIVE  | TYPE |   TYPO    | SUFFIX |      IPV4      |            IPV6              
+--------+------+-----------+--------+----------------+-----------------------------+
  ONLINE | CO   | oogle.com | com    | <nil>          | 2400:cb00:2048:1::681c:1ca2  
  ONLINE | CO   | gogle.com | com    | <nil>          | 2607:f8b0:4006:804::2004     
  ONLINE | CO   | gogle.com | com    | <nil>          | 2607:f8b0:4006:804::2004     
  ONLINE | CO   | goole.com | com    | 217.160.0.201  | 217.160.0.201                
  ONLINE | CO   | googe.com | com    | 162.243.10.151 | 162.243.10.151               
  ONLINE | CO   | googl.com | com    | <nil>          | 2607:f8b0:4006:81a::2004   
```

Generates variations for `google.com` using the character omission **(CO)** 
algorithm. Also execute extra functions to get the **ip** addresses, **idna** 
format and check for **ns** records. 
```
urlinsane google.com -t co -x ip -x idna -x ns

 _   _  ____   _      ___
| | | ||  _ \ | |    |_ _| _ __   ___   __ _  _ __    ___
| | | || |_) || |     | | | '_ \ / __| / _' || '_ \  / _ \
| |_| ||  _ < | |___  | | | | | |\__ \| (_| || | | ||  __/
 \___/ |_| \_\|_____||___||_| |_||___/ \__,_||_| |_| \___|

 Version: 0.1.0

   LIVE  | TYPE |   TYPO    | SUFFIX |      IPV4      |            IPV6             |   IDNA    |        NS         
+--------+------+-----------+--------+----------------+-----------------------------+-----------+------------------+
  ONLINE | CO   | oogle.com | com    | <nil>          | 2400:cb00:2048:1::681c:1da2 | oogle.com | mx2.zoho.com      
  ONLINE | CO   | gogle.com | com    | <nil>          | 2607:f8b0:4006:813::2004    | gogle.com |                   
  ONLINE | CO   | gogle.com | com    | <nil>          | 2607:f8b0:4006:813::2004    | gogle.com |                   
  ONLINE | CO   | goole.com | com    | 217.160.0.201  | 217.160.0.201               | goole.com | mx01.1and1.co.uk  
  ONLINE | CO   | googe.com | com    | 162.243.10.151 | 162.243.10.151              | googe.com |                   
  ONLINE | CO   | googl.com | com    | <nil>          | 2607:f8b0:4006:804::2004    | googl.com |  
  
```

For more details look at the **-h --help** output.
```
urlinsane -h

    
Generates domain typos and variations to detect and perform typo squatting, URL hijacking, phishing, and corporate espionage.

USAGE:
  urlinsane [domains] [flags]

OPTIONS:
  -c, --concurrency int         Number of concurrent workers (default 50)
  -f, --file string             Output filename
  -o, --format string           Output format (json, csv, text) (default "text")
  -x, --funcs stringArray       Extra functions for retrieving additional data (default [idna])
  -h, --help                    help for urlinsane
  -k, --keyboards stringArray   Keyboards/layouts ID to use (default [en1])
  -t, --typos stringArray       Types of typos to perform (default [all])
  -v, --verbose                 Output additional details
  
  ...
  
  
```





## Features

* Binary executable, written in Go with no dependencies. 
* Will have all the functionally of URLCrazy and DNSTwist. 
* Contains 19 typosquatting algorithms and 7 extra functions to retrieve additional data
* Modular architecture for language, keyboard, typo algorithm, and functions extensibility.
* Supports multiple keyboard layouts found in English, Spanish, Russian, Finish, and Arabic.
* Supports multiple languages with the ability to add more languages with ease.
* Concurrent function (**-x --funcs**) workers to retrieve additional info on each record.
* Concurrent typo squatting workers.



### English

* Over 8000 common misspellings
* Over 500 common homophones
* Registered graphemes, vowels, homoglyphs, and numerals
* Common keyboard layouts (qwerty, azerty, qwertz, dvorak)

See [Languages](https://rangertaha.github.io/urlinsane/#languages) for details 
on other languages.

## Algorithms

The modular architecture for code extensibility allows developers to add new 
typosquatting algorithms with ease. Currently we have implements 19 
typosquatting algorithms. See [Typo Algorithms](https://rangertaha.github.io/urlinsane/#algorithms) for details.


## Extra Functions

- **IDNA**  Show international domain name (Default)
- **MX**    Checking for DNS's MX records
- **TXT**   Checking for DNS's TXT records
- **IP**    Checking for IP address
- **NS**    Checks DNS NS records
- **CNAME** Checks DNS CNAME records
- **SIM**   Show domain similarity % using fuzzy hashing with ssdeep


### TODO 

* GeoIp Lookup.
* Estimate popularity of a domain variant via google search
* Lookup whois record
* Check for redirects and get new domain
* Emoji domains
* Grabs HTTP and SMTP service banners


## Authors

* [Rangertaha](https://github.com/rangertaha)


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
