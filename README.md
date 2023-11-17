# Wurzel

Wurzel is an ETL system designed to combine data from different sources and, optionally, serve it back in a semi standardised way. 
The power of this is from mashing the various source creating only configuration files and that the UI runs standard queries returning standard objects and so can reuse code.

Wurzel is composed of two parts the Harvester which collects the data, and the Combiner which stores and serves it. 

## The Harvester
defines it's data sources using feed defintion files, this means you just need to create a defintion file to
read a feed rather than create a new program or service. Currently the definition files support JSON, XML, Google Sheets and CSVs,
support is best for JSON. The definiton file specifies the location of the feed, how often it is fetched, what the feed looks like
and what data is to be extracted. The extracted data is published to a redis channel.

The Harvester can run independently and you can write your own program to deal with the data it produces, or it can be run with the Combiner.If it is run with the combiner it will need the entries in the feed defintion file to produce the standard objects that the combiner understands there are examples in the harvesters feeds.d directory.

## The Combiner
consists of a consumer and a server. The consumer waits for data to published to the redis channel, transforms the paths iand references acording to it's own definition file and saves the objects to the database (currently mongodb). The server returns data from the database using standard queries as standard JSOn objects. The Openapi/Swagger defintion is in the server/gen/http directory. The server should also allow the direct uploading of files to support data sources that can't be pulled from the internet, such as internal spreadsheets.

## Standard Objects

There are three types of objects in Wurzel: Pois, Events and Organisms objects can all reference each other so that they can be interconnected.  The different types all have standard fields but also a Properties field which can hold a hash of strings and so support any extra data needed.

### Standard Fields
All object types support :
- ID
- Title
- Description (Optional)
- Deactivated flag
- Path this is where the object lives in the database
- Refs a list of paths to other objects
- Properties

### Organisms
These are things witch are neither geographic nor time based, the contain just the standard fields.

### Events
are time based objects in addition to the standard fields thet have a start and an (optional) end.

### Pois
are geographically based objects, they are a bit of a work in progress but currently support :
- Location, a text location perhaps an address
- GeoJson, a GeoJson string
- GeoType, a GeoJson type, currently only Point, LineString, Polygon are supported
- GeoLoaction, a list of Geolocation objects each with a latitude and longitude attribute

We probably need to settle on just GeoJson or WKT as a string and map that into the db.

