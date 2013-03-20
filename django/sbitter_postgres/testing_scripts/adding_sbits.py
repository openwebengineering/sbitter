#
# Pre-Alpha Testing Script for Sbitter
# Please Note: This is not a real testing script, as in something to be used
#              for getting results that actually measure the perfomance of this
#              application.
#
# Author: AJ Bahnken, OpenWeb Engineering

# Add a bunch of random sbits through a GET request with JSON. Time it for fun.

import json
import requests
import time
import sys

from string import ascii_lowercase
from random import choice

if __name__ == '__main__':
    #First build 500 sbits, don't time this.
    print "building messages...\n"
    msgs = []
    for l in xrange(500):
        msgs.append(''.join(choice(ascii_lowercase) for x in xrange(10)))

    print "building sbit JSON objects...\n"
    sbits = []
    for msg in msgs:
        sbits.append(json.dumps({"message": msg, "user": "ajvb"})) #change user

    #Start timer, just for fun.
    print "POSTing sbits. This may take a while...\n"
    start_time = time.time()
    for sbit in sbits:
        r = requests.post('http://sbitterpy.openwebengineering.com/post_sbit_json/',
                          data=sbit)
        if r.status_code != 200:
            print "Something went wrong...\n"
            print r.status_code
            print r.content[3000:6000]
            sys.exit()

    end_time = time.time() - start_time

    print "Finished adding 500 sbits. Took: %s" % end_time
