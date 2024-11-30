const express = require("express");
const cors = require("cors");
const app = express();

app.use(cors());
app.use(express.json());

// print the request body
app.post("/api/test", (req, res) => {
    console.log(req.body);
    res.json({ message: "Request received!" });
});

// Sample products API
app.get("/api/products", (req, res) => {

    res.json([
        { id: 1, name: "Product 1", price: 10 },
        { id: 2, name: "Product 2", price: 20 },
    ]);
});

const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
    console.log(`Backend running on port ${PORT}`);
});
