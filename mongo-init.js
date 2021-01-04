db.createUser(
        {
            user: "testuser",
            pwd: "mongo",
            roles: [
                {
                    role: "readWrite",
                    db: "testapp"
                }
            ]
        }
);