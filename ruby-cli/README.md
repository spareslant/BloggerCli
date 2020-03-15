## Account Preparation
* You need to have a `google` account (gmail) (Paid account NOT required).
* Once you have a google account, Login to *https://console.developers.google.com/apis/credentials* and create a new `Project` here.
* Click on `Dashboard` on left hand column, and click on `+ ENABLE APIS AND SERVICES`.
   * Type in `blogger` in search bar, and select `Blogger API v3` and then click `ENABLE`.`
* Click on `OAuth consent screen` on left hand column, and select `User Type` to `External` and click `CREATE`.
   * On next screen type in `Application Name` as `Blogger CLI` (It can be any string).
   * Click on `Add Scope` and select both the `Blogger API v3` options and click `ADD` and click `Save`.
* Click on `Credentials` on the left hand side and you will be presented with following two options of creating credentials for Blogger API.
   * API Keys
   * OAuth 2.0 Client IDs
   * Service Accounts
   * Click on `+ CREATE CREDENTAILS` on the top and select `OAuth Client ID` from drop down menu.
   * Select `Application Type` to `other`.
   * Type in `Name` to `CLI Utility` (it can be any string) and click create.
   * You will see an entry with name `CLI Utility` under `OAuth 2.0 Client IDs`.
   * Download the credential files by clicking on a *down arrow* button.
   * This will be a `json` file and we need it to authenticate it with google `OAuth2.0` services.

* Login to blogger site *https://www.blogger.com*
   * Sign-in to blogger.
   * Enter a display name you like.
   * Create a `NEW BLOG` and give it a name you like in `Ttile`
   * Type in address. It has to be a unique e.g somethingrandom23455.blogspot.com and click `create blog`
   * On next screen, look at the address bar of browser, it will be like *https://www.blogger.com/blogger.g?blogID=3342324243435#allposts*
   * Note down blog ID (a string after **blogID=**, excluding #allposts) in above URL. We need this ID to put in ruby script.

## Ruby environment creation
I am using following ruby version.

```bash
$ ruby --version
ruby 2.7.0p0 (2019-12-25 revision 647ee6f091) [x86_64-darwin19]
```

We will be creating a python virtual environment.
```bash
mkdir BloggerCli
cd BloggerCli
mkdir ruby-cli
```

### Install required libraries.
Create a file `ruby-cli/Gemfile` with following contents.

```bash
source 'https://rubygems.org'
gem 'google-api-client'
gem 'googleauth'
gem 'launchy'
gem 'slop'
```


Run following commands to install above gems.
```bash
cd ruby-cli
bundler
```

### Copy credential script.
Copy above downloaded credential json file in `ruby-cli` folder and rename it to `OAuth2.0_secret.json`.

**Note:** You can rename it any name.

### Create a test file to be uploaded.
```bash
echo 'This is a first post' > sometext.txt
```
**Note**: You can create an HTML file as well.

### Run the script.
```bash
$ ruby bloggerCli.rb -f sometext.txt -t "My First Post" -e "<userAccount>@gmail.com" -b "3434533535534" -l label1,label2
enter the code you got from browser here and press Enter: 6/xQEBsJRhTZfthVP9ZSNsasaa7Y78HSDj2sd7aG
Total posts in this blog = 0
========== Inserting a new Post ========
Total posts in this blog = 1
```
**Note-1**: While running above script, this will open a browser and wil ask you to login to your google account.

**Note-2**: Above run will create a `tokens.yaml` file. On subsequent runs, browser window will not open due to presence of tokens.yaml file. If you remove this file, then you will have to authenticate again in the browser.

### References used:
* https://github.com/googleapis/google-api-ruby-client
* https://github.com/googleapis/google-api-ruby-client/blob/master/docs/oauth-server.md
* https://github.com/googleapis/google-api-ruby-client/blob/master/docs/oauth-installed.md
* https://github.com/googleapis/google-auth-library-ruby
* https://googleapis.dev/ruby/google-api-client/latest/Google/Apis/BloggerV3/BloggerService.html
* https://developers.google.com/identity/protocols/googlescopes
* https://console.developers.google.com/apis/credentials
* https://www.rubydoc.info/gems/slop/4.8.0

