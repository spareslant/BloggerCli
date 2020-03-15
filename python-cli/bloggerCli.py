#! /Users/gpal/BloggerCli/python-2.7.0/bin/python

# https://github.com/googleapis/google-api-python-client/blob/master/docs/README.md
# https://github.com/googleapis/google-api-python-client/blob/master/docs/oauth-server.md
# https://github.com/googleapis/google-api-python-client/blob/master/docs/oauth-installed.md
# http://googleapis.github.io/google-api-python-client/docs/dyn/blogger_v3.html
# http://googleapis.github.io/google-api-python-client/docs/dyn/blogger_v3.posts.html
# https://developers.google.com/identity/protocols/googlescopes
# https://console.developers.google.com/apis/credentials
# pip install --upgrade google-auth-oauthlib

# Note: This script does NOT check the duplicate title of posts while inserting the post in Blog.

from google.oauth2 import service_account
import googleapiclient.discovery
from google_auth_oauthlib.flow import InstalledAppFlow
import os
import sys
import argparse


usage = f"{os.path.basename(__file__)} -f <html file> -t <post title> -l <label_string1> -l <label_string2> -l .. -l .."

parser = argparse.ArgumentParser(prog=os.path.basename(__file__), usage=usage, description='Upload a post to Blogger')
parser.add_argument('-f', '--uploadfile', action='store', dest='fileToPost', help='html file to post', required=True)
parser.add_argument('-l', '--labels', action='append', dest='labels', default=[], help='-l <label1> -l <label2>')
parser.add_argument('-t', '--title', action='store', dest='title', help='-t <Post Title string>', required=True)
parser.add_argument('-b', '--blogid', action='store', dest='blogid', help='-b <blogid string>', required=True)
arguments = parser.parse_args()

if len(sys.argv) == 1:
    parser.print_help()
    parser.error("You must specify command line flags as mentioned above.")


FILE_TO_POST = arguments.fileToPost
LABELS_FOR_POST = arguments.labels
POST_TITLE = arguments.title
BLOG_ID = arguments.blogid

SCOPES = ['https://www.googleapis.com/auth/blogger']
SERVICE_ACCOUNT_FILE = 'blogger-api-credentials.json'
OAUTH2_ACCOUNT_FILE = 'OAuth2.0_secret.json'
FLOW_SERVER_PORT = 9090
API_SERVICE_NAME = 'blogger'
API_SERVICE_VERSION = 'v3'

# create a API client from service account
def create_client_from_serviceAccount():
    credentials = service_account.Credentials.from_service_account_file( SERVICE_ACCOUNT_FILE, scopes=SCOPES)

    blogger = googleapiclient.discovery.build(API_SERVICE_NAME, API_SERVICE_VERSION, credentials=credentials)
    return blogger

# create a client from an AOUTH2.0 account.
def create_client_from_outhAccount():
    flow = InstalledAppFlow.from_client_secrets_file(OAUTH2_ACCOUNT_FILE, scopes=SCOPES)
    credentials = flow.run_local_server(host='localhost', port=FLOW_SERVER_PORT, authorization_prompt_message='Please visit this URL: {url}', 
                        success_message='The auth flow is complete; you may close this window.', open_browser=True)
    blogger = googleapiclient.discovery.build(API_SERVICE_NAME, API_SERVICE_VERSION, credentials=credentials)
    return blogger

def print_number_of_posts(postList):
    if not 'items' in postList:
        print("No of posts in blog = 0")
    else:    
        print("No of posts in blog = {}".format(len(postList['items'])))

# From a service account you cannot create a new post in blogger. It might need some special permissions that I am not aware of.
# However, Service account can read posts and fetch them.
# blogger = create_client_from_serviceAccount()

# Below will open a browser window to authenticate yourself.
blogger = create_client_from_outhAccount()

postList = blogger.posts().list(blogId=BLOG_ID).execute()
print_number_of_posts(postList)

print("====== Inserting a new post =======")

with open(FILE_TO_POST) as f:
    contents = f.read()

blogBody = {
    "content": contents,
    "kind": "blogger#post",
    "title": POST_TITLE,
    "labels": LABELS_FOR_POST,
        }

result = blogger.posts().insert(blogId=BLOG_ID, body=blogBody).execute()
postList = blogger.posts().list(blogId=BLOG_ID).execute()
print_number_of_posts(postList)
