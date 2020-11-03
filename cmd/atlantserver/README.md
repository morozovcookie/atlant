Atlantserver
=============

Atlantserver is a main part of the current system. It responsible for processing requests from clients.


# Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)


# Installation

```bash
$ go get github.com/morozovcookie/atlant/cmd/atlantserver
```


# Configuration

Atlantserver use environment variables as configuration parameters. Below you can see list of available options:

<table>
<thead>
    <tr>
        <th>Name</th>
        <th>Type</th>
        <th>Description</th>
        <th>Sample</th>
        <th>Default value</th>
    </tr>
</thead>
<tbody>
<tr>
    <td colspan="5"><b>Kafka configuration parameters</b></td>
</tr>
<tr>
    <td>KAFKA_PRODUCT_PRODUCER_SERVERS</td>
    <td>string array</td>
    <td>List of Kafka broker's addresses</td>
    <td>127.0.0.1:29092,127.0.0.1:29093,127.0.0.1:29094</td>
    <td></td>
</tr>

<tr>
    <td colspan="5"><b>RPC server configuration parameters</b></td>
</tr>
<tr>
    <td>RPC_SERVER_HOST</td>
    <td>string</td>
    <td>Address which will be used for serving requests</td>
    <td>127.0.0.1:8080</td>
    <td></td>
</tr>

<tr>
    <td colspan="5"><b>MongoDB configuration parameters</b></td>
</tr>
<tr>
    <td>MONGODB_URI</td>
    <td>string</td>
    <td>MongoDB atlant database address</td>
    <td>mongodb://127.0.0.1:27017/atlant</td>
    <td></td>
</tr>
</tbody>
</table>
