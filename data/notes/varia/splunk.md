Basics
------

The Splunk *index* is a directory on a filesystem for data ingested by Splunk software. Events are stored in the index as two groups of files:

* Rawdata (.gz) - raw data in a crompressed form
* Index files (.tsidx) - metadata files for fast search that point to the raw data

.. these files reside in sets of directories, called *buckets*. An index tipycally consists of many buckets, organized by age: hot (stored in `db` - see below), warm (`db`), cold (`colddb`), frozen (data in this bucket is not searchable but can be thawed), thawed (`thaweddb`, buckets restored from an archive). Hot bucket is being written to, and has not necessarily been optimized.

Splunk stores raw data it indexed and its indexes within *flat files* in a structured directory (`$SPLUNK_HOME/var/lib/splunk`), meaning it doesn't require any database software running in the background. For example, the default index looks like this:

    /opt/splunk/var/lib/splunk/defaultdb $ tree -sC
    .
    ├── [       4096]  colddb
    ├── [       4096]  datamodel_summary
    ├── [       4096]  db
    │   ├── [         10]  CreationTime
    │   ├── [       4096]  GlobalMetaData
    │   └── [       4096]  hot_v1_0
    │       ├── [      53152]  1466256608-1466256608-4107774829233827835.tsidx
    │       ├── [      68997]  1466256608-1466256608-4107774830220961558.tsidx
    │       ├── [         67]  bucket_info.csv
    │       ├── [        105]  Hosts.data
    │       ├── [       4096]  rawdata
    │       │   └── [      94207]  0
    │       ├── [        111]  Sources.data
    │       ├── [        106]  SourceTypes.data
    │       ├── [         78]  splunk-autogen-params.dat
    │       └── [         41]  Strings.data
    └── [       4096]  thaweddb
    
    7 directories, 10 files

Splunk breaks data into events based on the timestamps it identifies.

Event data - all IT data that has been added to software indexes. The individual pieces of data are called *events*.

Splunk is designed as a platform extensible via Apps and Add-ons. Both are
packaged *sets of configuration*. "Search" is the default App.

* Apps - generally offer user interface
* Add-ons - generally enable digesting of particular type of data

Instance types (components)
---------------------------

Splunk data pipeline:

1. INPUT - receipt of raw data from log files, ports or scripts
2. PARSING (analyzing) - raw data split into events, time parsing, running
   transforms, setting base metadata, ...
3. INDEXING - data storage and optimization of indexes
4. SEARCH - running of queries and results presentation

These four stages of processing are generally split across two to four layers.

![Splunk network diagram](https://answers.splunk.com/storage/attachments/369-splunk-common-network-ports-ver1.5.jpg)

Forwarders (INPUT/PARSING)

* consume data (generally on the machines where the data originates) and then forward it to an indexer
* 1) Universal Forwarder - Splunk minus indexing and searching
* 2) full installation of Splunk configured as:
 * light forwarder (deprecated in 6.0) - no parsing, just sending of raw data to indexer
 * Heavy Forwarder - parse events and send them to the indexers
* important config files: inputs.conf, outputs.conf, props.conf,
  default-mode.conf, limits.conf

Indexers (PARSING/INDEXING/SEARCHING)

* do the heavy lifting; parse and index the data and run searches
* needs direct access to fast disks (local, SAN); NFS is not recommended
* each indexer just indexes data and performs searches *across its own indexes*
* important config files: inputs.conf, indexes.conf

Search heads (SEARCH)

* search management, i.e. coordinate searches across the set of indexers, consolidating the results and presenting them to the user
* configuration mostly managed via web interface: Settings => Distributed
  Search

Deployment server

* tool for distributing configurations, apps, and content *updates* to groups of components (classes) - forwarders, non-clustered indexers, and search heads
* can't be used for initial or upgrade installation

Sizing indexers
---------------

The following indexer should handle 100GB of raw logs per day (using AutoLB feature of Splunk forwarders) and four concurrent searches (including both interactive and [saved](https://docs.splunk.com/Splexicon:Savedsearch) ones):

* 8GB of RAM
* 8 fast physical CPUs
* Storage doing 800 IOPS (use bonnie++ to measure this)

If you have 200GB of logs per day you should have two such indexers (you should have two of them anyway because of high availability).

.conf files
-----------

Everything in Splunk is controlled by configuration files sitting in the filesystem of each instance of Splunk. Configuration changes done via the web interfaces end up in these .conf files.

$SPLUNK_HOME/etc/

* system/
 * default - defaults shipped with Splunk; never edit these!
 * local - global overrides specific to this host; very few configs need to live here
* **apps/$app_name/**
 * default - default app's configs
 * **local** - most configs should live here; all the non-private configs created via web interface will be placed here
* users/$user_name/$app_name/local - private configs; once the permissions are changed will be moved to $app_name/local

inputs.conf

* after leaving this stage data has some basic metadata associated with it:
  host, source, sourcetype, index

        # --- Cisco PIX ---
        [monitor:///logs/cisco_pix/*/*.log]
        
        # which parsing rules in props.conf to apply to these events
        # NOTE: important to set, otherwise explosion of sourcetypes will happen
        sourcetype = cisco:pix
        
        index = firewall_emea
        
        # set host from 3rd directory path segment
        host_segment = 3

props.conf

* which events to match based on host, source, and sourcetype

        [cisco:pix]
        TIME_PREFIX = ^
        TIME_FORMAT = %Y-%m-%dT%H:%M:%S
        MAX_TIMESTAMP_LOOKAHEAD = 25
        LINE_BREAKER = ([\r\n]+)
        SHOULD_LINEMERGE = false
        TZ = UTC
        
serverclass.conf

* mapping of apps to deployment clients

        [serverClass:<className>]
        # this stanza maps a host to a server class
        # options that should be applied to all apps in this class
        whitelist.0 = *.example.com
        whitelist.1 = 10.0.0.1
        [serverClass:<className>:app:<appName>]
        # this stanza maps an application to a server class
        # options that should be applied only to this app in this serverclass

Data sources
------------

Moniroting logs on servers (UF)

* preferred way
* highly optimized, events are usually searchable within few seconds

Monitoring logs on shared drive

Consuming logs in batch

* only copy *complete* logs to the watched directory
* use batch (vs. monitor) stanza in `inputs.conf` so Splunk deletes logs after indexing them
* copy sets of logs to different forwarders but don't copy the same logs to multiple machines

Receiving syslog events

Consuming logs from a database

Using scripts to gather data

Searching
=========

     search terms (keywords, "quoted phrases", fields, wildcard*, booleans, comparisons)
                         +
                         |                                    clause
                         |                                      +
                         |                                      |
                         v                                      v
    [search]  sourcetype=access_* status=503 | stats sum(price) as lost_revenue
                                                ^     ^    ^
                                                |     |    |
                                                +     |    +
                                           command    |   argument
                                                      +
                                                 function

* field - searchable name/value pair in event data

Performance tips
----------------

As events are stored by time, time is *always* the most efficient filter. After time, the most powerful keywords are `host`, `source`, `sourcetype`.

The more you tell Splunk, the better the chance for good results.

Field extraction is one of the most costly parts of a search. `fields [+] <wc-field-list>` - include *only* the specified fields; occurs before field extraction => improved performance.

Debugging
---------
    
https://answers.splunk.com/answers/4075/whats-the-best-way-to-track-down-props-conf-problems.html

    $SPLUNK_HOME/bin/splunk cmd btool props list <sourcetype>

Resources

* Implementing Splunk (Packt Publishing, 2015)
* http://docs.splunk.com/Splexicon
* http://docs.splunk.com/Documentation
