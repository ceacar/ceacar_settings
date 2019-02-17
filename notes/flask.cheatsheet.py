#BAREBONE APP
from flask import flask

app = flask(__name__)

@app.route('/hello/<stringname>')
def hello(name):
    return ("hello, world", 200)


#ROUTING
@app.route('/test') #default, only allow get method
@app.route('/test', method=['GET', 'POST']) #default, only allow get method
@app.route('/test', method=['PUT']) #allows only put

#CONFIGURATION
#direct access to config
app.config['CONFIG_NAME'] = 'config value'
#import an env var with a path to a config file
app.config.from_envar['ENV_VAR_NAME']


#TEMPALTES
from flask import render_template

@app.route('/')
def index():
    return render_template('template_file.html', var1 = value1, ...)

#JSON RESPONSE
import jsonify
@app.route('/returnstuff')
def returnstuff():
    return jsonify({'output':num_dict})

#ACCESS REQUEST DATA
requests.args['name'] # query string argument
request.form['name'] #get form data with fieldname name
request.method #request type
request.cookies.get('cookie_name') #cookies
request.files['name'] #files

#REDIRECT
from flask import url_for, redirect
#normal template
@app.route('/home/<var1>')
def index(var):
    return render_template('home.html')
#redirect template
@app.route('/redirect')
def redirect_example():
    return redirect(url_for('index'), variable = "sam") #redirect user to /home which defined by def index()


#ABORT
from falsk import abort
@app.route('/home')
def index():
    abort(404)

#SET COOKIE

@app.route('/home')
def index():
    resp = make_response(render_template('index.html'))
    resp.set_cookie('cookie_name', 'cookie_value')
    return resp

#SESSION HANDLING
import session
app.config['SECRET_KEY'] = 'any random string'

#set session
@app.route('/login_success')
def login_success():
    session['key_name'] = 'key_value' #stores a secure cookie in browser
    return redirect(url_for('index'))

#read session
@app.route('/')
def index():
    if 'key_name' in session: #session exist and has key
        session_var = session['key_value']
    else:#session does not exist
        return ("session not found", 401)

#USEFUL PLUGIN
flask-pymongo
flask-sqlalchemy
flask-wtf #a form helper for flask
flask-mail
flask_restful
flask-uploads
flask-user
flask-login

