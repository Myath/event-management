<h3>Project Name: Event Management</h3>
<h4>Features</h4>
<ol>
<li>Admin can create and manage event types</li>
<li>Admin can create and manage events</li>
<li>Admin can manage users</li>
<li>Guest users can register as users</li>
<li>The user can login </li>
<li>The user can participate in event</li>
<li>The user can comment on any events</li>
<li>Show all events in a calendar </li>
<li>Guest and login user booths are can see the all events and filter by event type and search events</li>
</ol>
<h4>Installation</h4>
<h5>HRM Config</h5>

<p><b>[Server]</b></p>
<p>host =127.0.0.1</p>
<p>port =3000</p>
<p><b>[database]</b></p>
<p>host =127.0.0.1</p>
<p>port =5432</p>
<p>dbname =event_hrm</p>
<p>user =postgres</p>
<p>password =secret</p>
<p>sslmode =disabled</p>
<h5>CMS Config</h5>
<p><b>[server]</b></p>
<p>baseurl=http://127.0.0.1:4000</p>
<p>port =4000</p>
<p><b>[session]</b></p>
<p>lifetime =24</p>
<p>idletime =20</p>
<p><b>[database]</b></p>
<p>host =127.0.0.1</p>
<p>Port =5432</p>
<p>dbname =event_session</p>
<p>user =postgress</p>
<p>password =secret</p>
<p>sslmode =disables</p>
<p><b>[hrm]</b></p>
<p>url =127.0.0.1:3000</p>

<h4>First To Need</h4>
<ul>
<li>Go should be istalled</li>
<li>Postgress should be installed</li>
<li>Browser should be istalled to serve localhost</li>
<li>A Edittor should be installed to run the project .Ex: Visual studio</li>
</ul>
<h4>Tables</h4>
<h5>Users table</h5>
<table>
        <thead>
            <tr>
                <th>Coulmn name</th>
                <th>Data type</th>
                <th>Default</th>
                <th>key</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>Id</td>
                <td>(int)</td>
                <td>not null</td>
                <td>primary key</td>
            </tr>
            <tr>
                <td>First name </td>
                <td>(varchar(20))</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Last name</td>
                <td>(varchar(20))</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Username</td>
                <td>(varchar(20))</td>
                <td>not null</td>
                <td>unique</td>
            </tr>
            <tr>
                <td>Email</td>
                <td>(varchar(50))</td>
                <td>not null</td>
                <td>unique</td>
            </tr>
            <tr>
                <td>Password</td>
                <td>(varchar(100))</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Is admin</td>
                <td>(bool)</td>
                <td>default false</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Is_active</td>
                <td>(bool)</td>
                <td>default active</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Created at</td>
                <td>(timestamp)</td>
                <td>default current timestamp </td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Updated at</td>
                <td>(timestamp)</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Deleted at</td>
                <td>(timestamp)</td>
                <td>default null</td>
                <td>n/a</td>
            </tr>
        </tbody>
    </table>
<h4>Event type table</h4>
<table>
        <thead>
            <tr>
                <th>Coulmn name</th>
                <th>Data type</th>
                <th>Default</th>
                <th>key</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>id</td>
                <td>int</td>
                <td>not null</td>
                <td>Primary key</td>
            </tr>
            <tr>
                <td>name</td>
                <td>varchar(20)</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>created_at</td>
                <td>timestamp</td>
                <td>Current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>updated_at</td>
                <td>timestamp</td>
                <td>Current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>deleted_at</td>
                <td>timestamp</td>
                <td>Default null</td>
                <td>n/a</td>
            </tr>
        </tbody>
    </table>
<h4>Event table</h4>
<table>
        <thead>
            <tr>
                <th>Coulmn name</th>
                <th>Data type</th>
                <th>Default</th>
                <th>key</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>Id</td>
                <td>int</td>
                <td>not null</td>
                <td>primary key</td>
            </tr>
            <tr>
                <td>Event type id</td>
                <td>int</td>
                <td>not null</td>
                <td>foreign key</td>
            </tr>
            <tr>
                <td>Name</td>
                <td>(text)</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Description</td>
                <td>(text)</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Location</td>
                <td>varchar(50)</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Start at</td>
                <td>(timestamp)</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>End at</td>
                <td>(timestamp)</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Published at</td>
                <td>(timestamp)</td>
                <td>default null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Created at</td>
                <td>(timestamp)</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Updated at</td>
                <td>(timestamp)</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Deleted at</td>
                <td>(timestamp)</td>
                <td>default null  </td>
                <td>n/a</td>
            </tr>
        </tbody>
    </table>
<h4>Comment table</h4>
<table>
        <thead>
            <tr>
                <th>Coulmn name</th>
                <th>Data type</th>
                <th>Default</th>
                <th>key</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>Id</td>
                <td>int</td>
                <td>not null</td>
                <td>primary key</td>
            </tr>
            <tr>
                <td>User id</td>
                <td>int</td>
                <td>not null</td>
                <td>foreign key</td>
            </tr>
            <tr>
                <td>Event id</td>
                <td>int</td>
                <td>not null</td>
                <td>foreign key</td>
            </tr>
            <tr>
                <td>comment</td>
                <td>(text)</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Created at</td>
                <td>timestamp</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Updated at</td>
                <td>timestamp</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
        </tbody>
    </table>

<h4>User event table</h4>
<table>
        <thead>
            <tr>
                <th>Coulmn name</th>
                <th>Data type</th>
                <th>Default</th>
                <th>key</th>
            </tr>
        </thead>
        <tbody>
            <tr>
                <td>User Id</td>
                <td>(int)</td>
                <td>not null</td>
                <td>foreign key</td>
            </tr>
            <tr>
                <td>Event id</td>
                <td>(int)</td>
                <td>not null</td>
                <td>foreign key</td>
            </tr>
            <tr>
                <td>Status</td>
                <td>(small int)</td>
                <td>not null</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Created at</td>
                <td>(timestamp)</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
            <tr>
                <td>Updated at</td>
                <td>(timestamp)</td>
                <td>default current timestamp</td>
                <td>n/a</td>
            </tr>
        </tbody>
    </table>
<h4>Requirements:</h4>
<ul>
<li>GRPC</li>
<li>Postgres</li>
<li>Golang</li>
<li>Microservice Architecture</li>
</ul>