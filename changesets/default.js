//changeset n.wagensonner:userUniqueUsernam
db.user.remove({});

db.user.createIndex(
    {username : 1}, 
    {unique: true}
);