   ####                      #####
  ##  ##                    ##   ##
 ##        ####             #         ####    ######    ####    ######
 ##       ##  ##   ######    #####   ##  ##    ##  ##      ##    ##  ##
 ##  ###  ##  ##                 ##  ##        ##       #####    ##  ##
  ##  ##  ##  ##            ##   ##  ##  ##    ##      ##  ##    #####
   #####   ####              #####    ####    ####      #####    ##
                                                                ####
_______________________________________________________________________
                                  ABOUT
_______________________________________________________________________

Description: 
Program: GoScrap | Web-Scraper
Current version: 1.0
Languages: Golang 1.23.5
Tested on: Linux 6.11.2 
Author: scarlet-oni
Dependencies: colly

_______________________________________________________________________
                                DOCUMENTATION
_______________________________________________________________________
Compilation:
    Makefile:
        $ cd src
        $ make
    Or go-build:
        $ cd src
        $ go build GoScrap.go
    Binary: 
        Watch in release

Launch:
    ./GoScrap <website> <timeout> <proxy>
        Required parameters:
            * <website> - website from which you need to do web scraping
        Minor parameters:
            * <timeout> - request timeout in seconds
            * <proxy> - proxy server in format <IP>:<PORT>
    
_________________________________________________________________________
                            LEGAL STATEMENT
_________________________________________________________________________
By downloading, modifying, redistributing, and/or executing GoScrap, the
user agrees to the contained LEGAL.txt statement found in this repository.

I, scarlet-oni, the creator, take no legal responsibility for unlawful actions
caused/stemming from this program. 

Use responsibly and ethically!
