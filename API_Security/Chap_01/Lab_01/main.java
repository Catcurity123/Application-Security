package com.manning.apisecurityinaction;

import java.nio.file.*;
import org.dalesbred.Database;
import org.h2.jdbcx.JdbcConnectionPool;
import org.json.JSONObject;
import com.manning.apisecurityinaction.controller.SpaceController;
import static spark.Spark.*;

public class Main {
    public static void main(String... args) throws Exception {
        var datasource = JdbcConnectionPool.create(
            "jdbc:h2:mem:natter", "natter", "password"
        );

        var database = Database.forDataSource(datasource);
        createTables(database);

        var spaceController = new SpaceController(database);

        post("/spaces", spaceController::createSpace);

        after((request, response) -> {
            response.type("application/json");
        });

        internalServerError(new JSONObject()
            .put("error", "internal server error")
            .toString()
        );

        notFound(new JSONObject()
            .put("error", "not found")
            .toString()
        );
    }

    private static void createTables(Database database) throws Exception {
        // Load the schema.sql file from resources
        var path = Paths.get(
            Main.class.getResource("/schema.sql").toURI()
        );

        // Execute the SQL schema script
        database.update(Files.readString(path));
    }
}