#!/usr/bin/env bash

echo "COPY nessesary files"
cp -r /usr/src/apkbuild/* /home/packager/

echo "Create ~/.abuild folder"

mkdir -p /home/packager/.abuild
touch /home/packager/.abuild/abuild.conf

echo "Import external key"

if [ "${EXTERNAL_KEY}" == "true" ]; then

    if [ -f $KEYNAME ]; then
        PUB_KEYFILE = $PUBKEY 
    else
        PUB_KEYFILE = /usr/share/apkbuild/$PUBKEY
    fi

    if [ -z $PUB_KEYFILE ]; then
        echo "Unable to import Key pub key does not exist"
    fi

    if [ -f $PRIVKEY ]; then
        PRIV_KEYFILE = $PRIVKEY 
    else
        PRIV_KEYFILE = /usr/share/apkbuild/$PRIVKEY
    fi


    if [ -z $PRIV_KEYFILE ]; then
        echo "Unable to import Key pub key does not exist"
    fi
    
    cp $PRIV_KEYFILE /home/packager/.abuild
    cp $PUB_KEYFILE /home/packager/.abuild
    cp $PUB_KEYFILE /home/packager/.abuild

    echo "PACKAGER_PRIVKEY=/home/packager/.abuild/$PRIVKEY" >  /home/packager/.abuild/abuild.conf
fi

echo "FIX Permissions on /home/packager"
chown -R packager:packager /home/packager/.abuild
chown -R packager:packager /home/packager/*

su packager -c "$*"
