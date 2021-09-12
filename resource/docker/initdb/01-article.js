db = db.getSiblingDB("golang_grpc_mongodb");

// db.createCollection("articles");

db.articles.insertMany([
  {
    author_id: "1",
    title: "List of The Best Programming Languange in The World",
    content: "Golang is on of the best programming languange in the world",
  },
  {
    author_id: "2",
    title: "List of The Best Database Driver in The World",
    content: "MongoDB is on of the best database driver in the world",
  },
]);
