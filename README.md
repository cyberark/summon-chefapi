# summon-chefapi

For many, Chef encrypted data bags are difficult to work with. This Summon provider allows you to use
[Summon + secrets.yml](http://conjurinc.github.io/summon/) to improve your development workflow with encrypted data bags.

## Example

Create an encrypted data bag with a PostgreSQL password.

```sh-session
$ knife data bag create passwords postgres --secret-file encrypted_data_bag_secret
```

```json
{
    "id": "postgres",
    "value": "mysecretpassword"
}
```

Install [Summon](https://github.com/conjurinc/summon) and this provider.

Create a [secrets.yml](https://conjurinc.github.io/summon/#secrets.yml) file.

```yaml
POSTGRES_PASSWORD: !var passwords/postgres/value
```

Now you can inject the password into any process as an environment variable. Instead of dealing with the Data Bag API
in your Chef recipe, you can just use `ENV['POSTGRES_PASSWORD']`.

```sh-session
$ summon chef-client --once
```

Once `chef-client` finishes, the password is gone, not left on your system.

## Getting Started 

### Using summon-chefapi with Conjur OSS 

Are you using this project with [Conjur OSS](https://github.com/cyberark/conjur)? Then we 
**strongly** recommend choosing the version of this project to use from the latest [Conjur OSS 
suite release](https://docs.conjur.org/Latest/en/Content/Overview/Conjur-OSS-Suite-Overview.html). 
Conjur maintainers perform additional testing on the suite release versions to ensure 
compatibility. When possible, upgrade your Conjur version to match the 
[latest suite release](https://docs.conjur.org/Latest/en/Content/ReleaseNotes/ConjurOSS-suite-RN.htm); 
when using integrations, choose the latest suite release that matches your Conjur version. For any 
questions, please contact us on [Discourse](https://discuss.cyberarkcommons.org/c/conjur/5).

### Install

1. Install the [latest release of Summon](https://github.com/conjurinc/summon#install).
2. Download the [latest release of this provider](https://github.com/conjurinc/summon-chefapi/releases)
and extract it to `/usr/local/lib/summon/`.

If you have more than one provider installed, select this one with `summon -p summon-chefapi ...`.

## Configure

Configuration of this provider is through environment variables:

* `CHEF_NODE_NAME`: The name of the node. (`node_name` in knife.rb)
* `CHEF_CLIENT_KEY_PATH`: The location of the file that contains the client key. (`client_key` in knife.rb)
* `CHEF_SERVER_URL`: The URL for the Chef server. (`chef_server_url` in knife.rb)
* `CHEF_DECRYPTION_KEY_PATH`: The location of the file that contains the decryption key.
* `CHEF_SKIP_SSL`: Skip SSL verification (for self-signed certs). Set to "1" to activate.

---

## Contributing

We welcome contributions of all kinds to this repository. For instructions on how to get started and descriptions of our development workflows, please see our [contributing
guide][contrib].

[contrib]: CONTRIBUTING.md

