// Copyright 2012 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
tsuru is a command line tool for application developers.

It provides some commands that allow a developer to register
himself/herself, manage teams, apps and services.

Usage:

	% tsuru <command> [args]

The currently available commands are (grouped by subject):

	target            changes or retrive the current tsuru server
	version           displays current tsuru version

	user-create       creates a new user
	login             authenticates the user with tsuru server
	logout            finishes the session with tsuru server
	key-add           adds a public key to tsuru deploy server
	key-remove        removes a public key from tsuru deploy server

	team-create       creates a new team (adding the current user to it automatically)
	team-list         list teams that the user is member
	team-user-add     adds a user to a team
	team-user-remove  removes a user from a team

	app-create        creates an app
	app-remove        removes an app
	app-list          lists apps that the user has access (see app-grant and team-user-add)
	app-info          displays information about an app
	app-grant         allows a team to have access to an app
	app-revoke        revokes access to an app from a team
	log               shows log for an app
	run               runs a command in all units of an app
	restart           restarts the app's application server

	env-get           display environment variables for an app
	env-set           set environment variable(s) to an app
	env-unset         unset environment variable(s) from an app

	bind              binds an app to a service instance
	unbind            unbinds an app from a service instance

	service-list      list all services, and instances of each service
	service-add       creates a new instance of a service
	service-remove    removes a instance of a service
	service-status    checks the status of a service instance
	service-info      list instances of a service, and apps binded to each instance
	service-doc       displays documentation for a service

Use "tsuru help <command>" for more information about a command.


Change or retrieve remote tsuru server

Usage:

	% tsuru target [target]

This command should be used to set current tsuru target, or retrieve current
target.

The target is the tsuru server to which all operations will be directed to.


Check current version

Usage:

	% tsuru version

This command returns the current version of tsuru command.


Create a user

Usage:

	% tsuru user-create <email>

user-create creates a user within tsuru remote server. It will ask for the
password before issue the request.


Authenticate within remote tsuru server

Usage:

	% tsuru login <email>

Login will ask for the password and check if the user is successfully
authenticated. If so, the token generated by the tsuru server will be stored in
${HOME}/.tsuru_token.

All tsuru actions require the user to be authenticated (except login and
user-create, obviously).


Logout from remote tsuru server

Usage:

	% tsuru logout

Logout will delete the token file and terminate the session within tsuru
server.


Add SSH public key to tsuru's git server

Usage:

	% tsuru key-add [${HOME}/.ssh/id_rsa.pub]

key-add sends your public key to tsuru's git server. By default, it will try
send a public RSA key, located at ${HOME}/.ssh/id_rsa.pub. If you want to send
other file, you can call it with the path to the file. For example:

	% tsuru key-add /etc/my-keys/id_dsa.pub

The key will be added to the current logged in user.


Remove SSH public key from tsuru's git server

Usage:

	% tsuru key-remove [${HOME}/.ssh/id_rsa.pub]

key-remove removes your public key from tsuru's git server. By default, it will
try to remove a key that match you public RSA key located at
${HOME}/.ssh/id_rsa.pub. If you want to remove a key located somewhere else,
you can pass it as parameter to key-remove:

	% tsuru key-remove /etc/my-keys/id_dsa.pub

The key will be removed from the current logged in user.


Create a new team for the user

Usage:

	% tsuru team-create <team-name>

team-create will create a team for the user. Tsuru requires a user to be a
member of at least one team in order to create an app or a service instance.

When you create a team, you're automatically member of this team.


List teams that the user is member of

Usage:

	% tsuru team-list

team-list will list all teams that you are member of.


Add a user to a team

Usage:

	% tsuru team-user-add <team-name> <user@email>

team-user-add adds a user to a team. You need to be a member of the team to be
able to add another user to it.


Remove a user from a team

Usage:

	% tsuru team-user-remove <team-name> <user@email>

team-user-remove removes a user from a team. You need to be a member of the
team to be able to remove a user from it.

A team can never have 0 users. If you are the last member of a team, you can't
remove yourself from it.


Create an app

Usage:

	% tsuru app-create <app-name> <platform>

app-create will create a new app using the given name and platform. For tsuru,
a platform is a Juju charm. To check the available platforms/charms, check this
URL: https://github.com/globocom/charms/tree/master/precise.

In order to create an app, you need to be member of at least one team. All
teams that you are member (see "tsuru team-list") will be able to access the
app.


Remove an app

Usage:

	% tsuru app-remove <app-name>

app-remove removes an app. If the app is binded to any service instance, it
will be unbinded before be removed (see "tsuru unbind"). You need to be a
member of a team that has access to the app to be able to remove it (you are
able to remove any app that you see in "tsuru app-list").


List apps that you have access to

Usage:

	% tsuru app-list

app-list will list all apps that you have access to. App access is controlled
by teams. If your team has access to an app, then you have access to it.


Display information about an app

Usage:

	% tsuru app-info <app-name>

app-info will display some informations about an specific app (its state,
platform, git repository, etc.). You need to be a member of a team that access
to the app to be able to see informations about it.


Allow a team to access an app

Usage:

	% tsuru app-grant <app-name> <team-name>

app-grant will allow a team to access an app. You need to be a member of a team
that has access to the app to allow another team to access it.


Revoke from a team access to an app

Usage:

	% tsuru app-revoke <app-name> <team-name>

app-revoke will revoke the permission to access an app from a team. You need to
have access to the app to revoke access from a team.

An app cannot be orphaned, so it will always have at least one authorized team.


See app's logs

Usage:

	% tsuru log <app-name>

Log will show log entries for an app. These logs are not related to the code of
the app itself, but to actions of the app in tsuru server (deployments,
restarts, etc.).


Run an arbitrary command in the app machine

Usage:

	% tsuru run <app-name> <command> [commandarg1] [commandarg2] ... [commandargn]

Run will run an arbitrary command in the app machine. Base directory for all
commands is the root of the app. For example, in a Django app, "tsuru run" may
show the following output:


	% tsuru run polls ls
	app.conf
	brogui
	deploy
	foo
	__init__.py
	__init__.pyc
	main.go
	manage.py
	settings.py
	settings.pyc
	templates
	urls.py
	urls.pyc


Restart the app's application server

Usage:

	% tsuru restart <app-name>

Restart will restart the application server (as defined in Procfile) of the
application.


Display environment variables of an application

Usage:

	% tsuru env-get <app-name> [variable-names]

env-get will display the name and the value of environment variables exported
in the application's environment. If none name is given, it will display the
value of all environment variables exported in the app via tsuru. It omits the
value of private environment variables (exported by service binding, see bind
command for more details). Examples of use:

	% tsuru env-get myapp MYSQL_DATABASE_NAME MYSQL_PASSWORD
	MYSQL_DATABASE_NAME=myapp_sql
	MYSQL_PASSWORD=*** (private variable)
	% tsuru env-get myapp
	MYSQL_DATABASE_NAME=myapp_sql
	MYSQL_USER=secret
	MYSQL_HOST=remote.mysql.com
	MYSQL_PASSWORD=*** (private variable)
	% tsuru env-get myapp SOMETHING_UNKNOWN

The first command retrieves only the specified variables, while the second
command retrieves all variables. In the last command, we ask for an undefined
variable, and env-get fails silently. All environment variable related commands
fail silently.


Define the value of one or more environment variables

Usage:

	% tsuru env-set <app-name> <NAME_1=VALUE_1> [NAME_2=VALUE_2] ... [NAME_N=VALUE_N]

env-set will (re)define environment variables for your app.  You can specify
one or more environment variables to (re)define. env-set cannot redefine
private variables, and all variables defined using env-set will be public (its
value will be displayed in env-get). env-set does not restart the application
after exporting the variables, for doing that, see restart command. Examples of
use:

	% tsuru env-set myapp MYSQL_DATABASE_NAME=myapp_sql2 MYSQL_PASSWORD=1234
	% tsuru env-get myapp MYSQL_DATABASE_NAME MYSQL_PASSWORD
	MYSQL_DATABASE_NAME=myapp_sql
	MYSQL_PASSWORD=*** (private variable)

Notice that env-set will fail silently to redefine private variables.


Undefine an environment variable

Usage:

	% tsuru env-unset <app-name> <NAME_1> [NAME_2] ... [NAME_N]

env-unset will undefine environments variables in your app.  You can specify
one or more environment variables to undefine.  env-unset cannot remove private
variables. Examples of use:

	% tsuru env-unset myapp MYSQL_DATABASE_NAME MYSQL_PASSWORD
	% tsuru env-get myapp MYSQL_DATABASE_NAME MYSQL_PASSWORD
	MYSQL_PASSWORD=*** (private variable)

Notice that env-unset will fail silently to undefine private variables.


Bind an application to a service instance

Usage:

	% tsuru bind <instance-name> <app-name>

Bind will bind an application to a service instance (see service-add for more
details on how to create a service instance).

When binding an application to a service instance, tsuru will add new
environment variables to the app. All environment variables exported by bind
will be private (not accessible via env-get).


Unbind an application from a service instance

Usage:

	% tsuru unbind <instance-name> <app-name>

Unbind will unbind an application from a service instance.  After unbinding,
the instance will not be available anymore.  For example, when unbinding an
application from a MySQL service, the app would lose access to the database.


List available services and instances

Usage:

	% tsuru service-list

service-list will retrieve and display a list of services that the user has
access to. If the user has any instance of services, it will be displayed by
this command too.


Create a new service instance

Usage:

	% tsuru service-add <service-name> <instance-name>

service-add will create a new service instance. After listing services with
"service-list", you may want to create a new service instance.

Example of use:

	% tsuru service-list
	+----------+-----------+
	| Services | Instances |
	+----------+-----------+
	| mysql    |           |
	+----------+-----------+
	% tsuru service-add mysql newmysql
	Service successfully added.
	% tsuru service-list
	+----------+-----------+
	| Services | Instances |
	+----------+-----------+
	| mysql    | newmysql  |
	+----------+-----------+


Remove a service instance

Usage:

	% tsuru service-remove <instance-name>

service-remove will destroy a service instance. It can't remove a service
instance binded to an app, so before remove a service instance, make sure there
is no apps binded to it (see "service-info" command).


Display information about a service

Usage:

	% tsuru service-info <service-name>

service-info will display a list of all instances of a given service (that the
user has access to), and apps binded to these instances.

Example of use:

	% tsuru service-info mysql
	Info for "mysql"
	+-----------+-------+
	| Instances | Apps  |
	+-----------+-------+
	| newmysql  |       |
	+-----------+-------+
	% tsuru bind newmysql myapp
	...
	% tsuru service-info mysql
	Info for "mysql"
	+-----------+-------+
	| Instances | Apps  |
	+-----------+-------+
	| newmysql  | myapp |
	+-----------+-------+


Check if a service instance is up

Usage:

	% tsuru service-status <instance-name>

service-status will display the status of the given service instance. For now,
it checks only if the instance is "up" (receiving connections) or "down"
(refusing connections).


Display the documentation of a service

Usage:

	% tsuru service-doc <service-name>

service-doc will display the documentation of a service.
*/
package documentation