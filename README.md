# PasswordManagerApi

The PasswordManagerApi is a REST Api that handles CRUD operations for the [PasswordManager](https://github.com/Legitzx/PasswordManager).

## Functionality

***Register*** ``Method: POST``
 - Description: Registers a user.
 - Input: User Object with id & email
 
***Login*** ``Method: POST``
 - Description: Takes in the users id and checks to see if they are in the database, if so they are authenticated using a JSON Web Token (JWT).
 - Input: User Object with id
 
***Get Vault*** ``Method: GET``
 - Description: Sends back User Objects to authenticated users.
 - Input: User Object with id
 
***Update Vault*** ``Method: PUT``
 - Description: Updates a users Vault. Must be authenticated.
 - Input: User Object with id
 
## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
