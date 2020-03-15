
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
   * Type in `Name` to `Python Blogger CLI` (it can be any string) and click create.
   * You will see an entry with name `Python Blogger CLI` under `OAuth 2.0 Client IDs`.
   * Download the credential files by clicking on a *down arrow* button.
   * This will be a `json` file and we need it to authenticate it with google `OAuth2.0` services.

* Login to blogger site *https://www.blogger.com*
   * Sign-in to blogger.
   * Enter a display name you like.
   * Create a `NEW BLOG` and give it a name you like in `Ttile`
   * Type in address. It has to be a unique e.g somethingrandom23455.blogspot.com and click `create blog`
   * On next screen, look at the address bar of browser, it will be like *https://www.blogger.com/blogger.g?blogID=3342324243435#allposts*
   * Note down blog ID (a string after **blogID=**, excluding #allposts) in above URL. We need this ID to put in python script.

## Python environment creation
I am using following python version.

```bash
$ python --version
Python 3.7.0
```

We will be creating a python virtual environment.
```bash
mkdir BloggerCli
cd BloggerCli
python -m venv python-2.7.0
source python-2.7.0/bin/activate
```

### Install required libraries.
```bash
pip install google-api-python-client
pip install --upgrade google-auth-oauthlib
```

### Copy credential script.
```bash
mkdir python-cli
```
Copy above downloaded credential json file in `python-cli` folder and rename it to `OAuth2.0_secret.json`.

**Note:** You can rename it any name.

### Create a test file to be uploaded.
```bash
echo 'This is a first post' > sometext.txt
```
**Note**: You can create an HTML file as well.

### Run the script.
```bash
$ ./bloggerCli.py -f sometext.txt -t "My First Post" -b 34456769666334 -l label1 -l label2 -l label3
Please visit this URL: https://accounts.google.com/o/oauth2/auth?response_type=code&client_id=26652434353-e9u77isb2sdsd4343sdsd343434.apps.googleusercontent.com&redirect_uri=http%3A%2F%2Flocalhost%3A9090%2F&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fblogger&state=BkaXVyiW15sdsdsdADH7sadsH8hS&access_type=offline
No of posts in blog = 0
====== Inserting a new post =======
No of posts in blog = 1
```
**Note-1**: While running above script, this will open a browser and wil ask you to login to your google account.

**Note-2**: You may need to change the *FLOW_SERVER_PORT* in above script.

### References used:
* https://github.com/googleapis/google-api-python-client/blob/master/docs/README.md
* https://github.com/googleapis/google-api-python-client/blob/master/docs/oauth-server.md
* https://github.com/googleapis/google-api-python-client/blob/master/docs/oauth-installed.md
* http://googleapis.github.io/google-api-python-client/docs/dyn/blogger_v3.html
* http://googleapis.github.io/google-api-python-client/docs/dyn/blogger_v3.posts.html
* https://developers.google.com/identity/protocols/googlescopes
* https://console.developers.google.com/apis/credentials

