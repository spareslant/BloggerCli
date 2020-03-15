#! /Users/gpal/.rbenv/shims/ruby

# https://github.com/googleapis/google-api-ruby-client
# https://github.com/googleapis/google-api-ruby-client/blob/master/docs/oauth-server.md
# https://github.com/googleapis/google-api-ruby-client/blob/master/docs/oauth-installed.md
# https://github.com/googleapis/google-auth-library-ruby
# https://googleapis.dev/ruby/google-api-client/latest/Google/Apis/BloggerV3/BloggerService.html
# https://developers.google.com/identity/protocols/googlescopes
# https://console.developers.google.com/apis/credentials
# https://www.rubydoc.info/gems/slop/4.8.0

# Note: This script does NOT check the duplicate title of posts while inserting the post in Blog.

require 'googleauth'
require 'google/apis/blogger_v3'
require 'googleauth/stores/file_token_store'
require 'launchy'
require 'slop'

opts = Slop::Options.new
opts.banner = "usage: #{File.basename(__FILE__)} -f <html file> -t <post title> -l <label_string1> -l <label_string2> -l .. -l .."
opts.string '-f', '--uploadfile', 'html file', required: true
opts.string '-t', '--title', 'Post Title', required: true
opts.string '-e', '--email', 'Account email id', required: true
opts.string '-b', '--blogid', 'Blog id', required: true
opts.array '-l', '--labels', 'comma separated labels', delimiter: ','

parser = Slop::Parser.new(opts)
begin
    result = parser.parse(ARGV)
rescue Slop::MissingRequiredOption, Slop::MissingRequiredOption
    puts opts
    raise
end

BLOG_ID = result[:blogid]
ACCOUNT_EMAIL = result[:email]
FILE_TO_POST = result[:uploadfile]
POST_TITLE = result[:title]
LABELS_FOR_POST = result[:labels]

SCOPE = 'https://www.googleapis.com/auth/blogger'
OOB_URI = 'urn:ietf:wg:oauth:2.0:oob'
SERVICE_ACCOUNT_FILE = 'blogger-api-credentials.json'
OAUTH2_ACCOUNT_FILE = 'OAuth2.0_secret.json'


# create an API client from service account
def create_client_from_serviceAccount(scope, service_account_cred)
    authorizer = Google::Auth::ServiceAccountCredentials.make_creds(
                    json_key_io: File.open(service_account_cred), scope: scope)

    authorizer.fetch_access_token!
    blogger = Google::Apis::BloggerV3::BloggerService.new
    blogger.authorization = authorizer
    return blogger
end

# create an API client from an AOUTH2.0 account.
# Following function will create a new file called tokens.yaml
def create_client_from_outhAccount(scope, oob_uri, user_id, oauth_cred_file)
    #oob_uri = 'urn:ietf:wg:oauth:2.0:oob'
    #user_id = 'eyemole@gmail.com'
    client_id = Google::Auth::ClientId.from_file(oauth_cred_file)
    token_store = Google::Auth::Stores::FileTokenStore.new(:file => 'tokens.yaml')
    authorizer = Google::Auth::UserAuthorizer.new(client_id, scope, token_store)
    credentials = authorizer.get_credentials(user_id)
    if credentials.nil?
        url = authorizer.get_authorization_url(base_url: oob_uri )
        #Launchy.open(url)
        puts "Open this URL in Browser and enter the code you got from browser below"
        puts "URL: #{url}"
        print "enter the code you got from browser here and press Enter: "
        # code = gets
        code = STDIN.gets.chomp
        credentials = authorizer.get_and_store_credentials_from_code(user_id: user_id, code: code, base_url: oob_uri)
    end
    blogger = Google::Apis::BloggerV3::BloggerService.new
    blogger.authorization = credentials
    return blogger
end

# From a service account you cannot create a new post in blogger. It might need some special permissions that I am not aware of.
# However, Service account can read posts and fetch them.
# blogger = create_client_from_serviceAccount(scope, SERVICE_ACCOUNT_FILE)

# Below will open a browser window to authenticate yourself.
blogger = create_client_from_outhAccount(SCOPE, OOB_URI, ACCOUNT_EMAIL, OAUTH2_ACCOUNT_FILE)

blog = blogger.get_blog(BLOG_ID)
# puts "blog URL = #{blog.url}"
puts "Total posts in this blog = #{blog.posts.total_items}"

puts "========== Inserting a new Post ========"

contents = File.read(FILE_TO_POST)

new_post = Google::Apis::BloggerV3::Post.new
new_post.kind = "blogger#post"
new_post.title = POST_TITLE
new_post.content = contents
new_post.labels = LABELS_FOR_POST
blogger.insert_post(BLOG_ID, new_post)

blog = blogger.get_blog(BLOG_ID)
puts "Total posts in this blog = #{blog.posts.total_items}"

#posts.items.each do |post|
#    puts post.to_json
#end
