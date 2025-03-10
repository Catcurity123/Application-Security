package com.manning.apisecurityinaction.controller;

import org.dalesbred.Database;
import org.json.JSONObject;
import spark.Request;
import spark.Response;

import java.sql.SQLException;

public class SpaceController {
    private final Database database;

    public SpaceController(Database database) {
        this.database = database;
    }

    public JSONObject createSpace(Request request, Response response) throws SQLException {
        // Parse the JSON request body to extract space details
        var json = new JSONObject(request.body());

        // Extract and validate space name
        var spaceName = json.getString("name");
        if (spaceName.length() > 255) {
            throw new IllegalArgumentException("space name too long");
        }

        // Extract and validate owner name
        var owner = json.getString("owner");
        if (!owner.matches("[a-zA-Z][a-zA-Z0-9]{1,29}")) {
            throw new IllegalArgumentException("invalid username: " + owner);
        }

        return database.withTransaction(tx -> {
            // Generate a unique ID for the new space
            var spaceId = database.findUniqueLong("SELECT NEXT VALUE FOR space_id_seq;");

            // Insert the new space into the database
            database.updateUnique(
                "INSERT INTO spaces(space_id, name, owner) VALUES (?, ?, ?);",
                spaceId, spaceName, owner
            );

            // Set the response status and header
            response.status(201);  // HTTP 201 Created
            response.header("Location", "/spaces/" + spaceId);

            // Return a JSON object with the new space details
            return new JSONObject()
                    .put("name", spaceName)
                    .put("uri", "/spaces/" + spaceId);
        });
    }
}
