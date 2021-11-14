#!/bin/bash


file="cockroachdb/creds.py"

for file in {"cockroach/cockroachcreds.py","twilio-sms-api/creds.py"}
do
    if [ ! -f "$file" ]; then
        touch $file
        cat > $file << "END"
    CDBUSER = ""
    CDBPASS = ""
END
    fi
done