## PHP Example Explanation

Here is an example of a `PasswordReminder` that connects to a MySQL database:

```php
class MySQLConnection
{
    public function connect()
    {
        // handle the database connection
        return 'Database connection';
    }
}

class PasswordReminder
{
    private $dbConnection;

    public function __construct(MySQLConnection $dbConnection)
    {
        $this->dbConnection = $dbConnection;
    }
}
```

First, the `MySQLConnection` is the low-level module while the `PasswordReminder` is high level, but according to the definition of D in SOLID, which states to Depend on abstraction, not on concretions. This snippet above violates this principle as the `PasswordReminder` class is being forced to depend on the `MySQLConnection` class.

Later, if you were to change the database engine, you would also have to edit the PasswordReminder class, and this would violate the open-close principle.

The `PasswordReminder` class should not care what database your application uses. To address these issues, you can code to an interface since high-level and low-level modules should depend on abstraction:

```php
interface DBConnectionInterface
{
    public function connect();
}
```

The interface has a connect method and the `MySQLConnection` class implements this interface. Also, instead of directly type-hinting `MySQLConnection` class in the constructor of the `PasswordReminder`, you instead type-hint the `DBConnectionInterface` and no matter the type of database your application uses, the `PasswordReminder` class can connect to the database without any problems and open-close principle is not violated.

```php
class MySQLConnection implements DBConnectionInterface
{
    public function connect()
    {
        // handle the database connection
        return 'Database connection';
    }
}

class PasswordReminder
{
    private $dbConnection;

    public function __construct(DBConnectionInterface $dbConnection)
    {
        $this->dbConnection = $dbConnection;
    }
}
```

This code establishes that both the high-level and low-level modules depend on abstraction.
