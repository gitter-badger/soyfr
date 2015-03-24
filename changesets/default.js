//mongeez formatted javascript

//changeset nwagensonner:userUniqueUsername
db.user.remove({});

db.user.createIndex(
    {username : 1}, 
    {unique: true}
);