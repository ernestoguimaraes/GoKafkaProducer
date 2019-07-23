# Send messages to Kafka using GO

This small sample shows how to send messages to a Kafka Topic Using GOLang. To accomplish that in your DEV environament, I recommend to use the Spotify´ Kafka image which contains in one image the KAFKA and ZOOKEEPER.

https://hub.docker.com/r/spotify/kafka/

The code is very simple. It uses the https://github.com/Shopify/sarama reference to accomplish the communication with Kafka. This package is very powerfull. Feel Free to create a Consumer either!

Note for Windows 10 Users: I needed to install a GCC compiler becuase you will notice that when running the application, GO Terminal will thrown a GCC error.  https://mingw-w64.org/doku.php

Have Fun.