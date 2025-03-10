package com.manning.apisecurityinaction;

import java.nio.file.*;
import com.manning.apisecurityinaction.controller.SpaceController;
import org.dalesbred.Database;
import org.dalesbred.result.EmptyResultException;
import org.h2.jdbcx.JdbcConnectionPool;
import org.json.JSONException;
import org.json.JSONObject;
import spark.Request;
import spark.Response;
import static spark.Spark.*;

public class Main {

    public static void main(String... args) throws Exception {
        // Create an in-memory H2 database connection pool
        var datasource = JdbcConnectionPool.create("jdbc:h2:mem:natter", "natter", "password");
        
        // Create a Database object using the connection pool
        var database = Database.forDataSource(datasource);
        
        // Set up tables in the database
        createTables(database);

        // Construct the SpaceController and pass it the Database object
        var spaceController = new SpaceController(database);
        
        // Route to handle POST requests to /spaces by calling createSpace on the controller
        post("/spaces", spaceController::createSpace);
        
        // Add a response filter to set all output as JSON
        after((request, response) -> response.type("application/json"));
        
        // Set up error handling for 500 and 404 status codes
        internalServerError(new JSONObject().put("error", "internal server error").toString());
        notFound(new JSONObject().put("error", "not found").toString());

        // Exception handling for specific errors
        exception(IllegalArgumentException.class, Main::badRequest);
        exception(JSONException.class, Main::badRequest);
        exception(EmptyResultException.class, (e, request, response) -> response.status(404));
    }

    private static void badRequest(Exception ex, Request request, Response response) {
        response.status(400);
        response.body(new JSONObject().put("error", ex.getMessage()).toString());
    }

    private static void createTables(Database database) throws Exception {
        // Load the schema.sql file from the resources directory
        var path = Paths.get(Main.class.getResource("/schema.sql").toURI());
        
        // Execute SQL commands to create tables in the database
        database.update(Files.readString(path));
    }
}
