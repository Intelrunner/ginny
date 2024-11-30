const express = require("express");
const cors = require("cors");
const app = express();

app.use(cors());
app.use(express.json());

// Sample products API
app.get("/api/products", (req, res) => {

    res.json([
        { id: 0, request: req },
        { id: 1, name: "Product 1", price: 10 },
        { id: 2, name: "Product 2", price: 20 },
    ]);
});

const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
    console.log(`Backend running on port ${PORT}`);
});
