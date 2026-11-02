# Device configuration

You can configure net devices via:
- CLI, SNMP (traditional)
- network specific APIs
	- NETCONF - SSH/HTTP, IETF
	- RESTCONF - SSH/HTTP, IETF
	- gNMI - gRPC, OpenConfig (Google), data models + protocols
- multipurpose APIs (HTTP, Linux shell)

# Data modeling

- **data formatting** (e.g. YAML, JSON, protobuf) is for for serializing and
  deserializing data
- data formats provide basic type system so a program knows whether the given
  bytes represent a string, integer, or boolean
- but here's the problem - what happens when we send this to a network device:

	{ "hostname": "" }

- JSON is correct but the device might have problems setting the hostname
- the device software could do something like

	req = json.loads(json_str)
	if req["hostname"] == "":
		raise Exception("Hostname field must not be emtpy")

- but if there are many fields and many checks needed, **data modeling** is a
  better approach
- it's a set of tools and techniques that provide constraints to the structured
  data like JSON

- the canonical example is the *database schema*
- it describes the organization and structure of data within a DB
- for example, primary key is a special designation describing a column of
  values that can be used to uniquely identify the row to which those values
  belong

- in general, *schema* is just a way to describe the structure of data
- data model =~ schema

Data modeling technologies
- YANG
- JSON schema
- XML schema definition

## JSON schema

JSON document to model:

	{
		"hostname": "sw01",
		"vlans": [
			{
				"id": 100,
				"name": "VLAN_100"
			},
			{
				"id": 200,
				"name": "VLAN_200"
			}
		],
		"nameservers": [
			"1.1.1.1",
			"8.8.8.8"
		]
	}

JSON Schema definition (the model):

	{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"title": "Config",
		"description": "A bit of configuration data for a network device",
		"type": "object",
		"properties": {
			"hostname": {
				"type": "string",
				"minlenght": 1,
				"maxlength": 20
			},
			...omitted for brevity...
		},
		"required": ["hostname", "vlans", "nameservers"]
	}

---

From: Network Programmability and Automation, 2nd Edition
