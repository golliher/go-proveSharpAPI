# Hypothesis

The "newly designed for 2015" television I purchased does not yet fully implement the Sharp IP CONTROL protocol.  It does partially support IP CONTROL, but in a manner that is crippled and might suggest the firmware
was shipped before IP CONTROL was finished and tested.

# Goal

Demonstrate a defect in the current firmware and influence Sharp Inc. to include a fix in the next version of the firmware.  

Further, I hope to determine the likelihood of this happening in a timely manner and if I should return the new television to the retailer as unsuitable to my needs.

# Evidence

## Observed problems
None of the several IP remote applications I tried worked with my new television, but all worked with my old television.   This was observed after verifying no network issues and no configuration issues with the two television sets tested.

Further, I discovered that some applications *partially* worked.

Further still, I discovered that the presence of a configured password for the remote control endpoint had no effect on a application working or not working.   The apps that partially worked did so without a password even if I configured the TV to require a password.

These are the software packages I tried:
- https://github.com/golliher/go-sharptv   (This is my own software)
- https://github.com/benburkhart1/sharp-aquos-remote-control
* Sharp Remote app (legacy app?)
* Roomie remote


## Reported problems
I found a least one report on the Internet indicating a lack of IP CONTROL support in all 2015 Sharp TV models.  Further, that source recommends **NOT BUYING** Sharp TVs, and instead buying Sony TVs.

Source: http://www.roomieremote.com/forums/users/admin/replies/


# Experimental testing

With the knowledge that some, but not all functionality is present in the new TV, I set out
to understand in more detail what works and what does not work.

In order to clearly demonstrate the difference between a functional firmware and a dis-functional firmware
I developed a small test program to send commands to two television sets and report the results.  Details follow
including how to obtain the test software.

### "OLD TV"

- Status:  Fully functional
- Model: LC-70C6400U
- IP address: 192.168.4.11
- Protocol Version: 0100
- Firmware Version: 222U1302091

### "NEW TV"

- Status:  Partially functional / significantly crippled
- Model: LC-43LE653U
- IP address: 192.168.4.21
- Protocol Version:    UNKNOWN
- Firmware Version:   SW 2.13.1
- ULI Software Version 1
- ULI Module Version 14



# Test results

        ----------OLD TV---------------               ----------NEW TV---------------

         SENT          RECEIVED                        SENT          RECEIVED
         ========      ================                ========      ================
         POWR1         OK                              POWR1         OK
         MUTE0         OK                              MUTE0         OK
         SWVN1         222U1302091                     SWVN1         ERR
         MNRD1         C6400UA                         MNRD1         ERR
         POWR?         1                               POWR?         ERR
         MUTE?         1                               MUTE?         ERR
         MUTE1         OK                              MUTE1         ERR
         VOLM?         1                               VOLM?         ERR
         IAVD?         4                               IAVD?         ERR
         MUTE2         OK                              MUTE2         ERR
         IAVD?         4                               IAVD?         ERR
         VOLM1         OK                              VOLM1         ERR
         RCKY33        OK                              RCKY33        OK
         RCKY32        OK                              RCKY32        OK
         RCKY39        OK                              RCKY39        OK
         RCKY46        OK                              RCKY46        OK
         RCKY36        OK                              RCKY36        OK

# Conclusion

Only what I call "one-way" commands are implemented in the new TV.  These are sufficient to duplicate the functionality of a handheld remote control, but insufficient for automation.

The more useful commands for automation that are able to directly set a value are **NOT working**.

Examples:

- Directly set the input to input #4 without sending the input key four times.    (i.e.   "IAVD4")
- Directly set the volume level to 25 without sending the volume up keypress 25 times. (i.e. "VOLM25")

Moreover, a second class of useful commands also does **NOT work**.  These are the commands that return information when called.

Examples of useful status commands that do NOT work:

- Get the current power state on/off state (i.e. "POWR?")
- Get the current volume level (i.e. "VOLM?")
- Get the current input source (i.e. "IAVD?")



# Testing software

The software and source code used for this test is freely available for download at the following location should you wish to duplicate my experiment.

https://github.com/golliher/go-proveSharpAPI/


# Reference documentation

Source of documentation for expected behavior of IP CONTROL protocol is pages 55 & 56 of the Manual for Model LC-70C6400U.

http://files.sharpusa.com/Downloads/ForHome/HomeEntertainment/LCDTVs/Manuals/mon_man_LC70LE640U_LC60LE640U_LC52LE640U_LC70C6400U_LC60C6400U_LC52C6400U_LC80LE633U.pdf

# Other observations

This section is not germane to my main point, and I include it only as supplemental feedback to the developers of the firmware.

The non-functional firmware issues a Username and Password prompt that are seemingly non-functional.  If one simply ignores their presence and sends commands, they will be accepted or result in an error -- no password required.

While the TV has configuration options to set Username and Password, no documentation that I am aware of details how a client program should provide this authentication information.

There is a web server running on port 80 with a "A simple html+cgi example" being served.  This appears to be an oversight and not what one would expect to find from a top brand.   Better would be a blank page, a page with the sharp logo, a link to the eManual or no web server at all.

# Reported by

- Darrell Golliher, customer
- Background and contact info available at http://golliher.net
