// ---------------------------------------------------------
// BLOG-AGGREGATOR
// ---------------------------------------------------------

Blog-Aggregator is a robust, command-line interface application designed to continuously scrape and aggregate RSS feeds. It allows users to register accounts, subscribe to their favorite feeds, and view the latest posts directly in their terminal. The application runs a background worker that intelligently fetches the most outdated feeds on a customizable interval.

The backend architecture is built entirely in Go (1.25+). It utilizes PostgreSQL for persistent data storage. Database migrations are managed via Goose, ensuring safe schema updates. The data access layer is generated using SQLC, providing type-safe and efficient database queries.

// ---------------------------------------------------------
// PREREQUISITES
// ---------------------------------------------------------

To run this application, you must have the following installed on your system:

Go (version 1.25.5 or higher): Required to compile and run the application.

PostgreSQL: A running Postgres database server is required for data storage. You will need to create a database and execute the included Goose migrations (found in the sql/schema directory) to set up the necessary tables before running the program.

// ---------------------------------------------------------
// INSTALLATION
// ---------------------------------------------------------

You can install the application globally on your system using the go install command. Run the following in your terminal:

go install github.com/Psyduck000054/Blog-Aggregator@latest

Ensure that your Go binary directory (typically ~/go/bin) is added to your system's PATH variable so that you can execute the command from any terminal window.

// ---------------------------------------------------------
// SETUP AND USAGE
// ---------------------------------------------------------

Before running any commands, you must configure the application by creating a file named .gatorconfig.json in your user's home directory. This file stores your database connection string and tracks your active session. Add the following JSON structure to the file, replacing the URL with your actual database credentials:

{
"db_url": "postgres://username:password@localhost:5432/your_database_name?sslmode=disable"
}

Once the configuration file is in place and the database schema is built, you can start using the CLI. The executable name will match the repository name (Blog-Aggregator).

To see a full list of available commands and how to use them, run the built-in help manual:

Blog-Aggregator help

(Note: If you run the program without any arguments, the help menu will display automatically to guide you.)